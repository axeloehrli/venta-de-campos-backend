package db

import (
	"context"
	"database/sql"
)

const createUsuario = `
INSERT INTO usuarios (
  nombre_usuario,
  nombre,
  apellido,
	email
) VALUES (
  $1, $2, $3, $4
) RETURNING *
`

type CreateUsuarioParams struct {
	NombreUsuario string `json:"nombre_usuario"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
}

func CreateUsuario(ctx context.Context, arg CreateUsuarioParams, db *sql.DB) (Usuario, error) {
	row := db.QueryRowContext(ctx, createUsuario, arg.NombreUsuario, arg.Nombre, arg.Apellido, arg.Email)
	var u Usuario
	err := row.Scan(
		&u.ID,
		&u.NombreUsuario,
		&u.Nombre,
		&u.Apellido,
		&u.Email,
		&u.FechaCreacion,
	)
	return u, err
}

const getUsuario = `
SELECT * FROM usuarios
WHERE id = $1 LIMIT 1
`

func GetUsuario(ctx context.Context, id int64, db *sql.DB) (Usuario, error) {
	row := db.QueryRowContext(ctx, getUsuario, id)
	var u Usuario
	err := row.Scan(
		&u.ID,
		&u.NombreUsuario,
		&u.Nombre,
		&u.Apellido,
		&u.Email,
		&u.FechaCreacion,
	)
	return u, err
}

const getUsuarios = `
SELECT * FROM usuarios
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetUsuariosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func GetUsuarios(ctx context.Context, arg GetUsuariosParams, db *sql.DB) ([]Usuario, error) {
	rows, err := db.QueryContext(ctx, getUsuarios, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Usuario{}
	for rows.Next() {
		var i Usuario
		if err := rows.Scan(
			&i.ID,
			&i.NombreUsuario,
			&i.Nombre,
			&i.Apellido,
			&i.Email,
			&i.FechaCreacion,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteUsuario = `
DELETE FROM usuarios
WHERE id = $1
`

func DeleteUsuario(ctx context.Context, id int64, db *sql.DB) error {
	_, err := db.ExecContext(ctx, deleteUsuario, id)
	return err
}
