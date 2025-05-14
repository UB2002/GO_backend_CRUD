package services

import (
	"patient_portal/models"

	"github.com/jinzhu/gorm"
)

type PatientService struct {
    db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
    return &PatientService{db}
}

func (s *PatientService) Create(p *models.Patient) error {
    return s.db.Create(p).Error
}

func (s *PatientService) List() ([]models.Patient, error) {
    var list []models.Patient
    err := s.db.Find(&list).Error
    return list, err
}

func (s *PatientService) GetByID(id uint) (*models.Patient, error) {
    var p models.Patient
    if err := s.db.First(&p, id).Error; err != nil {
        return nil, err
    }
    return &p, nil
}

func (s *PatientService) Update(p *models.Patient) error {
    return s.db.Save(p).Error
}

func (s *PatientService) Delete(id uint) error {
    return s.db.Delete(&models.Patient{}, id).Error
}
