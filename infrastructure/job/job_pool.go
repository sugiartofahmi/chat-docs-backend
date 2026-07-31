package job

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"go-service/entities"
	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/job/interfaces"

	"gorm.io/gorm"
)

type JobPool struct {
	registry interfaces.JobRegistryInterface
	db       *gorm.DB
}

func NewJobPool(registry interfaces.JobRegistryInterface, db *gorm.DB) interfaces.JobPoolInterface {
	return &JobPool{registry: registry, db: db}
}

func (p *JobPool) Start(numWorkers int) {
	for range numWorkers {
		go p.poll()
	}
}

func (p *JobPool) Delegate(jobName string, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		panic(*exceptions.ServerErrorException(err))
	}

	entry := &entities.JobEntity{
		JobName: jobName,
		Source:  "pool",
		Payload: string(body),
		Status:  "pending",
	}
	if err := p.db.Create(entry).Error; err != nil {
		panic(*exceptions.ServerErrorException(err))
	}
}

func (p *JobPool) poll() {
	ticker := time.NewTicker(PollInterval)
	defer ticker.Stop()

	for range ticker.C {
		p.claimAndRun()
	}
}

func (p *JobPool) claimAndRun() {
	var entry entities.JobEntity

	if err := p.db.Where("status = ?", "pending").Order("created_at").First(&entry).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("job claim failed: %v", err)
		}
		return
	}

	now := time.Now()
	if err := p.db.Model(&entry).Updates(map[string]any{
		"status":     "running",
		"started_at": now,
	}).Error; err != nil {
		log.Printf("job claim update failed: %v", err)
		return
	}

	p.safeRun(entry)
}

func (p *JobPool) safeRun(entry entities.JobEntity) {
	defer func() {
		if r := recover(); r != nil {
			p.finish(&entry, "failed", fmt.Sprintf("%v", r))
			log.Printf("job %s panicked: %v", entry.JobName, r)
		}
	}()

	handler, ok := p.registry.Get(entry.JobName)
	if !ok {
		message := fmt.Sprintf("job %s not registered", entry.JobName)
		p.finish(&entry, "failed", message)
		log.Println(message)
		return
	}

	if err := handler.Handle([]byte(entry.Payload)); err != nil {
		p.finish(&entry, "failed", err.Error())
		log.Printf("job %s failed: %v", entry.JobName, err)
		return
	}

	p.finish(&entry, "success", "")
}

func (p *JobPool) finish(entry *entities.JobEntity, status string, errMessage string) {
	finishedAt := time.Now()
	updates := map[string]any{
		"status":      status,
		"finished_at": finishedAt,
	}
	if errMessage != "" {
		updates["error"] = errMessage
	}
	if err := p.db.Model(entry).Updates(updates).Error; err != nil {
		log.Printf("job update failed for %s: %v", entry.JobName, err)
	}
}
