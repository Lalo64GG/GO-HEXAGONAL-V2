package infraestructure

import (
	"database/sql"
	"errors"
	"go-hexagonal-v2/src/alumno/domain"
	"strconv"
)

type MySQLAlumnoRepository struct {
    db *sql.DB
}

func NewMySQLAlumnoRepository(db *sql.DB) *MySQLAlumnoRepository {
    return &MySQLAlumnoRepository{db}
}

func (r *MySQLAlumnoRepository) Save(alumno *domain.Alumno) error {
    query := "INSERT INTO alumnos (id, name, age, grup) VALUES (?, ?, ?, ?)"
    result, err := r.db.Exec(query, alumno.ID, alumno.Name, alumno.Age, alumno.Grup)
    if err != nil {
        return err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    alumno.ID = strconv.FormatInt(id, 10)
    return nil
}

func (r *MySQLAlumnoRepository) FindByID(id string) (*domain.Alumno, error) {
    query := "SELECT id, name, age, grup FROM alumnos WHERE id = ?"
    row := r.db.QueryRow(query, id)
    var alumno domain.Alumno
    if err := row.Scan(&alumno.ID, &alumno.Name, &alumno.Age, &alumno.Grup); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &alumno, nil
}

func (r *MySQLAlumnoRepository) FindAll() ([]domain.Alumno, error) {
    query := "SELECT id, name, age, grup FROM alumnos"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var alumnos []domain.Alumno
    for rows.Next() {
        var alumno domain.Alumno
        if err := rows.Scan(&alumno.ID, &alumno.Name, &alumno.Age, &alumno.Grup); err != nil {
            return nil, err
        }
        alumnos = append(alumnos, alumno)
    }

    return alumnos, nil
}

func (r *MySQLAlumnoRepository) Delete(id string) error {
    query := "DELETE FROM alumnos WHERE id = ?"
    _, err := r.db.Exec(query, id)
    return err
}
