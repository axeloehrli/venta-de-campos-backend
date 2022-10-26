package db

import "time"

type Usuario struct {
	ID                int64     `json:"id"`
	NombreUsuario     string    `json:"nombre_usuario"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Nombre            string    `json:"nombre"`
	Apellido          string    `json:"apellido"`
	Email             string    `json:"email"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
}

type Campo struct {
	ID                int64     `json:"id"`
	IDUsuario         int64     `json:"id_usuario"`
	Titulo            string    `json:"titulo"`
	Tipo              string    `json:"tipo"`
	Hectareas         int64     `json:"hectareas"`
	PrecioPorHectarea int64     `json:"precio_por_hectarea"`
	Ciudad            string    `json:"ciudad"`
	Provincia         string    `json:"provincia"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
}
