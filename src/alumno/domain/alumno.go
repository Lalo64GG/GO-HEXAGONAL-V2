package domain

type Alumno struct {
    ID     string   `json:"id"`
    Name string `json:"nombre" validate:"required"`
    Age   int    `json:"edad" validate:"required"`
    Grup  string `json:"grupo" `
}
