package domain

type AlumnoRepository interface {
    Save(alumno *Alumno) error
    FindByID(id string) (*Alumno, error)
    FindAll() ([]Alumno, error)
    Delete(id string) error
}
