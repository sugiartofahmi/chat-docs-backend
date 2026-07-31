package scheduler

import (
	"fmt"
	"log"
	"time"

	"go-service/entities"
	"go-service/infrastructure/scheduler/interfaces"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Scheduler struct {
	cron *cron.Cron
	db   *gorm.DB
}

func NewScheduler(db *gorm.DB) interfaces.SchedulerInterface {
	return &Scheduler{cron: cron.New(), db: db}
}

func (s *Scheduler) Schedule(name string, spec string, callback func()) error {
	_, err := s.cron.AddFunc(spec, func() {
		s.runLogged(name, callback)
	})
	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) runLogged(name string, callback func()) {
	startedAt := time.Now()
	logEntry := &entities.JobEntity{
		JobName:   name,
		Source:    "scheduler",
		Status:    "running",
		StartedAt: &startedAt,
	}
	if err := s.db.Create(logEntry).Error; err != nil {
		log.Printf("job_logs insert failed for %s: %v", name, err)
	}

	defer func() {
		if r := recover(); r != nil {
			s.finish(logEntry, "failed", fmt.Sprintf("%v", r))
			log.Printf("scheduled job %s panicked: %v", name, r)
			return
		}
		s.finish(logEntry, "success", "")
	}()

	callback()
}

func (s *Scheduler) finish(logEntry *entities.JobEntity, status string, errMessage string) {
	finishedAt := time.Now()
	updates := map[string]any{
		"status":      status,
		"finished_at": finishedAt,
	}
	if errMessage != "" {
		updates["error"] = errMessage
	}
	if err := s.db.Model(logEntry).Updates(updates).Error; err != nil {
		log.Printf("job_logs update failed for %s: %v", logEntry.JobName, err)
	}
}
