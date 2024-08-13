package application

import (
    "go-hexagonal-v2/src/alumno/domain"
    "github.com/google/uuid"
)

type AlumnoService struct {
    repo domain.AlumnoRepository
}

func NewAlumnoService(repo domain.AlumnoRepository) *AlumnoService {
    
    return &AlumnoService{repo}
}

func (s *AlumnoService) CreateAlumno(alumno domain.Alumno) error {
    alumno.ID = uuid.New().String()
    return s.repo.Save(&alumno)
}

func (s *AlumnoService) GetAlumno(id string) (*domain.Alumno, error) {
    return s.repo.FindByID(id)
}

func (s *AlumnoService) GetAllAlumnos() ([]domain.Alumno, error) {
    return s.repo.FindAll()
}

func (s *AlumnoService) DeleteAlumno(id string) error {
    return s.repo.Delete(id)
}
