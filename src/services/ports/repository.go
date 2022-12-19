package ports

import "github.com/oswaldom-code/affiliate-tracker/src/domain/models"

type Repository interface {
	// system_services ports
	TestDb() error
	Save(referred models.Referred) error
	GetAll() ([]models.Referred, error)
	GetById(id int64) (models.Referred, error)
	GetByAgentId(agentId string) ([]models.Referred, error)
}
