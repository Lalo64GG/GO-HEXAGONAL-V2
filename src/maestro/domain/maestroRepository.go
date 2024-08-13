package domain

type MaestroRepository interface {
    Create(maestro *Maestro) error
    FindByID(id string) (*Maestro, error)
    FindAll() ([]Maestro, error)
    Delete(id string) error
    FindByEmailAndPassword(email, password string) (*Maestro, error) 
}
