package db

import (
	"context"
	"database/sql"
)

const createCampo = `
INSERT INTO campos (
  id_usuario,
  titulo,
  tipo,
	hectareas,
	precio_por_hectarea,
	ciudad,
	provincia
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *
`

type CreateCampoParams struct {
	IDUsuario         int64  `json:"id_usuario"`
	Titulo            string `json:"titulo"`
	Tipo              string `json:"tipo"`
	Hectareas         int64  `json:"hectareas"`
	PrecioPorHectarea int64  `json:"precio_por_hectarea"`
	Ciudad            string `json:"ciudad"`
	Provincia         string `json:"provincia"`
}

func CreateCampo(ctx context.Context, arg CreateCampoParams, db *sql.DB) (Campo, error) {
	row := db.QueryRowContext(
		ctx,
		createCampo,
		arg.IDUsuario,
		arg.Titulo,
		arg.Tipo,
		arg.Hectareas,
		arg.PrecioPorHectarea,
		arg.Ciudad,
		arg.Provincia,
	)
	var c Campo
	err := row.Scan(
		&c.ID,
		&c.IDUsuario,
		&c.Titulo,
		&c.Tipo,
		&c.Hectareas,
		&c.PrecioPorHectarea,
		&c.Ciudad,
		&c.Provincia,
		&c.FechaCreacion,
	)
	return c, err
}

const getCampo = `
SELECT * FROM campos
WHERE id = $1 LIMIT 1
`

func GetCampo(ctx context.Context, id int64, db *sql.DB) (Campo, error) {
	row := db.QueryRowContext(ctx, getCampo, id)
	var c Campo
	err := row.Scan(
		&c.ID,
		&c.IDUsuario,
		&c.Titulo,
		&c.Tipo,
		&c.Hectareas,
		&c.PrecioPorHectarea,
		&c.Ciudad,
		&c.Provincia,
		&c.FechaCreacion,
	)
	return c, err
}

const getCampos = `
SELECT * FROM campos
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetCamposParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func GetCampos(ctx context.Context, arg GetCamposParams, db *sql.DB) ([]Campo, error) {
	rows, err := db.QueryContext(ctx, getCampos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Campo{}
	for rows.Next() {
		var c Campo
		if err := rows.Scan(
			&c.ID,
			&c.IDUsuario,
			&c.Titulo,
			&c.Tipo,
			&c.Hectareas,
			&c.PrecioPorHectarea,
			&c.Ciudad,
			&c.Provincia,
			&c.FechaCreacion,
		); err != nil {
			return nil, err
		}
		items = append(items, c)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteCampo = `
DELETE FROM campos
WHERE id = $1
`

func DeleteCampo(ctx context.Context, id int64, db *sql.DB) error {
	_, err := db.ExecContext(ctx, deleteCampo, id)
	return err
}
