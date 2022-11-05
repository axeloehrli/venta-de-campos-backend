package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/axeloehrli/venta-de-campos-backend/util"
)

func createRandomCampo(t *testing.T) Campo {
	usuario := createRandomUsuario(t)

	arg := CreateCampoParams{
		IDUsuario:         usuario.ID,
		Titulo:            util.RandomTitulo(),
		Descripcion:       util.RandomString(30),
		Tipo:              util.RandomTipo(),
		Hectareas:         util.RandomHectareas(),
		PrecioPorHectarea: util.RandomPrecioPorHectarea(),
		Ciudad:            util.RandomCiudad(),
		Provincia:         util.RandomProvincia(),
	}

	campo, err := CreateCampo(
		context.Background(),
		arg,
		database,
	)

	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}
	if (Campo{}) == campo {
		t.Fatalf("EMPTY CAMPO")
	}
	if arg.Titulo != campo.Titulo {
		t.Fatalf("DIFFERENT TITULO")
	}
	if arg.Descripcion != campo.Descripcion {
		t.Fatalf("DIFFERENT DESCRIPCION")
	}
	if arg.Tipo != campo.Tipo {
		t.Fatalf("DIFFERENT TIPO")
	}
	if arg.Hectareas != campo.Hectareas {
		t.Fatalf("DIFFERENT HECTAREAS")
	}
	if arg.PrecioPorHectarea != campo.PrecioPorHectarea {
		t.Fatalf("DIFFERENT PRECIO POR HECTAREAS")
	}
	if arg.Ciudad != campo.Ciudad {
		t.Fatalf("DIFFERENT CIUDAD")
	}
	if arg.Provincia != campo.Provincia {
		t.Fatalf("DIFFERENT PROVINCIA")
	}
	if campo.ID == 0 {
		t.Fatalf("INVALID ID")
	}
	if campo.FechaCreacion.IsZero() {
		t.Fatalf("INVALID FECHA CREACION")
	}
	return campo
}
func TestCreateCampo(t *testing.T) {
	createRandomCampo(t)
}

func TestGetCampo(t *testing.T) {
	campo1 := createRandomCampo(t)
	campo2, err := GetCampo(context.Background(), campo1.ID, database)
	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}
	if (Campo{}) == campo2 {
		t.Fatalf("EMPTY CAMPO")
	}
	if campo2.Titulo != campo1.Titulo {
		t.Fatalf("DIFFERENT TITULO")
	}
	if campo2.Descripcion != campo1.Descripcion {
		t.Fatalf("DIFFERENT DESCRIPCION")
	}
	if campo2.Tipo != campo1.Tipo {
		t.Fatalf("DIFFERENT TIPO")
	}
	if campo2.Hectareas != campo1.Hectareas {
		t.Fatalf("DIFFERENT HECTAREAS")
	}
	if campo2.PrecioPorHectarea != campo1.PrecioPorHectarea {
		t.Fatalf("DIFFERENT PRECIO POR HECTAREAS")
	}
	if campo2.Ciudad != campo1.Ciudad {
		t.Fatalf("DIFFERENT CIUDAD")
	}
	if campo2.Provincia != campo1.Provincia {
		t.Fatalf("DIFFERENT PROVINCIA")
	}
	if campo2.FechaCreacion != campo1.FechaCreacion {
		t.Fatalf("DIFFERENT FECHA CREACION")
	}
}

func TestDeleteCampo(t *testing.T) {
	campo1 := createRandomCampo(t)
	err := DeleteCampo(context.Background(), campo1.ID, database)

	if err != nil {
		t.Fatalf("ERROR DELETING CAMPO: %v", err)
	}

	campo2, err := GetCampo(context.Background(), campo1.ID, database)
	if err == nil {
		t.Fatalf("CAMPO NOT SUCCESFULLY DELETED: %v", err)
	}
	if err != sql.ErrNoRows {
		t.Fatalf("UNKNOWN ERROR: %v", err)
	}
	if (Campo{}) != campo2 {
		t.Fatal("CAMPO MUST BE EMPTY")
	}
}

func TestListCampos(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCampo(t)
	}
	arg := ListCamposParams{
		Limit:  5,
		Offset: 5,
	}
	campos, err := ListCampos(context.Background(), arg, database)

	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}

	if len(campos) == 0 {
		t.Fatalf("CAMPO LIST IS EMPTY")
	}

	for _, campo := range campos {
		if (Campo{}) == campo {
			t.Fatalf("CAMPO IS EMPTY")
		}
	}
}
