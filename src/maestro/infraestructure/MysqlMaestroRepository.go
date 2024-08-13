package infraestructure

import (
	"database/sql"
	"errors"
	"go-hexagonal-v2/src/maestro/domain"
	"log"
)

type MySQLMaestroRepository struct {
    db *sql.DB
}

func NewMySQLMaestroRepository(db *sql.DB) *MySQLMaestroRepository {
	return &MySQLMaestroRepository{db}
}

func (r *MySQLMaestroRepository) Create(maestro *domain.Maestro) error {
	query := "INSERT INTO maestros (id, name, email, password) VALUES (?, ?, ?, ?)"
    _, err := r.db.Exec(query, maestro.ID, maestro.Name, maestro.Email, maestro.Password) 
	if err != nil {
        log.Printf("Error inserting maestro: %v", err)
        return err
    }

	log.Printf("Inserted Maestro ID: %s", maestro.ID) 
	return nil
}

func (r *MySQLMaestroRepository) FindByID(id string) (*domain.Maestro, error) {
	query := "SELECT id, name, email FROM maestros WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var maestro domain.Maestro
	if err := row.Scan(&maestro.ID, &maestro.Name, &maestro.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Error scanning row: %v", err) // Log the error
		return nil, err
	}
	return &maestro, nil
}

func (r *MySQLMaestroRepository) FindAll() ([]domain.Maestro, error) {
	query := "SELECT id, name, email FROM maestros"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying rows: %v", err) // Log the error
		return nil, err
	} 

	defer rows.Close()

	var maestros []domain.Maestro

	for rows.Next() {
		var maestro domain.Maestro
        if err := rows.Scan(&maestro.ID, &maestro.Name, &maestro.Email); err != nil {
            log.Printf("Error scanning row: %v", err) // Log the error
            return nil, err
        }
        maestros = append(maestros, maestro)
	}

	return maestros, nil
}

func (r *MySQLMaestroRepository) Delete(id string) error {
	query := "DELETE FROM maestros WHERE id = ?"
    _, err := r.db.Exec(query, id)
    if err != nil {
        log.Printf("Error deleting row: %v", err) // Log the error
    }
    return err
}

func (repo *MySQLMaestroRepository) FindByEmailAndPassword(email, password string) (*domain.Maestro, error) {
    var maestro domain.Maestro
    err := repo.db.QueryRow("SELECT id, name, email FROM maestros WHERE email = ? AND password = ?", email, password).Scan(&maestro.ID, &maestro.Name, &maestro.Email)
    if err != nil {
        log.Printf("Error finding maestro by email and password: %v", err) // Log the error
        return nil, err
    }
    return &maestro, nil
}
