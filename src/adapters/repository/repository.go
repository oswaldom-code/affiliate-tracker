package repository

import (
	"fmt"

	"github.com/oswaldom-code/affiliate-tracker/pkg/log"
	"github.com/oswaldom-code/affiliate-tracker/src/domain/models"
	"gorm.io/gorm"
)

var sessionConfig = &gorm.Session{SkipDefaultTransaction: true, FullSaveAssociations: false}

func (r *repository) TestDb() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (r *repository) Save(referred models.Referred) error {
	fmt.Println("Saving: ", referred)
	tx := r.db.Table("referred").Create(&referred)
	if tx.Error != nil {
		log.ErrorWithFields("Error saving referred", log.Fields{"error": tx.Error})
		return tx.Error
	}
	return nil

}

func (r *repository) GetAll() ([]models.Referred, error) {
	referredList := []models.Referred{}
	db := r.db.Session(sessionConfig).Table("referred").Find(&referredList)
	if db.Error != nil {
		return []models.Referred{}, db.Error
	}
	return referredList, nil
}

func (s *repository) GetById(id int64) (models.Referred, error) {
	referred := models.Referred{}
	db := s.db.Session(sessionConfig).Table("referred").Find(&referred, id)
	if db.Error != nil {
		return models.Referred{}, db.Error
	}
	if db.RowsAffected == 0 {
		return models.Referred{}, fmt.Errorf("referred not found with id: %d", id)
	}
	return referred, nil
}

func (s *repository) GetByAgentId(agentId string) ([]models.Referred, error) {
	referredList := []models.Referred{}
	db := s.db.Session(sessionConfig).Table("referred").Where("agent_id = ?", agentId).Find(&referredList)
	if db.Error != nil {
		return []models.Referred{}, db.Error
	}
	return referredList, nil
}
