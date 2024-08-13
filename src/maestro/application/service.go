package application

import (
	"github.com/google/uuid"
	"go-hexagonal-v2/src/maestro/domain"
	"log"
)

type MaestroService struct {
	repo domain.MaestroRepository
}

func NewMaestroService(repo domain.MaestroRepository) *MaestroService {
	return &MaestroService{repo}
}

func (s *MaestroService) CreateMaestro(maestro domain.Maestro) error {
    maestro.ID = uuid.New().String() 
    log.Printf("Generated UUID: %s", maestro.ID) 
    return s.repo.Create(&maestro)
}

func (s *MaestroService) GetMaestro(id string) (*domain.Maestro, error) {
	return s.repo.FindByID(id)
}

func (s *MaestroService) GetAllMaestros() ([]domain.Maestro, error) {
	return s.repo.FindAll()	
}

func (s *MaestroService) DeleteMaestro(id string) error {
	return s.repo.Delete(id)
}

func (s *MaestroService) Authenticate(email, password string) (*domain.Maestro, error) {
    return s.repo.FindByEmailAndPassword(email, password)
}