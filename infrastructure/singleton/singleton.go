package singleton

import (
	"sync"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/redis/go-redis/v9"
	"go-service/infrastructure/integrations"
	jobInterfaces "go-service/infrastructure/job/interfaces"
	schedulerInterfaces "go-service/infrastructure/scheduler/interfaces"
	"gorm.io/gorm"
)

var (
	once             sync.Once
	httpClient       *integrations.HttpClient
	db               *gorm.DB
	redisClient      *redis.Client
	opensearchClient *opensearchapi.Client
	jobRegistry      jobInterfaces.JobRegistryInterface
	jobPool          jobInterfaces.JobPoolInterface
	scheduler        schedulerInterfaces.SchedulerInterface
)

func Init(
	httpClientInstance *integrations.HttpClient,
	dbInstance *gorm.DB,
	redisClientInstance *redis.Client,
	opensearchClientInstance *opensearchapi.Client,
	jobRegistryInstance jobInterfaces.JobRegistryInterface,
	jobPoolInstance jobInterfaces.JobPoolInterface,
	schedulerInstance schedulerInterfaces.SchedulerInterface,
) {
	once.Do(func() {
		httpClient = httpClientInstance
		db = dbInstance
		redisClient = redisClientInstance
		opensearchClient = opensearchClientInstance
		jobRegistry = jobRegistryInstance
		jobPool = jobPoolInstance
		scheduler = schedulerInstance
	})
}
