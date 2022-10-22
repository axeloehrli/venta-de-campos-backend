package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/axeloehrli/venta-de-campos-backend/util"
)

func createRandomUsuario(t *testing.T) Usuario {
	arg := CreateUsuarioParams{
		NombreUsuario: util.RandomNombreUsuario(),
		Nombre:        util.RandomNombre(),
		Apellido:      util.RandomApellido(),
		Email:         util.RandomEmail(),
	}

	usuario, err := CreateUsuario(context.Background(), arg, database)
	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}
	if (Usuario{}) == usuario {
		t.Fatalf("EMPTY USUARIO")
	}
	if arg.NombreUsuario != usuario.NombreUsuario {
		t.Fatalf("DIFFERENT NOMBRE USUARIO")
	}
	if arg.Nombre != usuario.Nombre {
		t.Fatalf("DIFFERENT NOMBRE")
	}
	if arg.Apellido != usuario.Apellido {
		t.Fatalf("DIFFERENT APELLIDO")
	}
	if arg.Email != usuario.Email {
		t.Fatalf("DIFFERENT EMAILS")
	}
	if usuario.ID == 0 {
		t.Fatalf("INVALID ID")
	}
	if usuario.FechaCreacion.IsZero() {
		t.Fatalf("INVALID FECHA CREACION")
	}
	return usuario
}
func TestCreateUsuario(t *testing.T) {
	createRandomUsuario(t)
}

func TestGetUsuario(t *testing.T) {
	usuario1 := createRandomUsuario(t)
	usuario2, err := GetUsuario(context.Background(), usuario1.ID, database)
	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}
	if (Usuario{}) == usuario2 {
		t.Fatalf("EMPTY USUARIO")
	}
	if usuario2.NombreUsuario != usuario1.NombreUsuario {
		t.Fatalf("DIFFERENT NOMBRE USUARIO")
	}
	if usuario2.Nombre != usuario1.Nombre {
		t.Fatalf("DIFFERENT NOMBRE")
	}
	if usuario2.Apellido != usuario1.Apellido {
		t.Fatalf("DIFFERENT APELLIDO")
	}
	if usuario2.Email != usuario1.Email {
		t.Fatalf("DIFFERENT EMAIL")
	}
	if usuario2.FechaCreacion != usuario1.FechaCreacion {
		t.Fatalf("DIFFERENT FECHA CREACION")
	}
}

func TestDeleteUsuario(t *testing.T) {
	usuario1 := createRandomUsuario(t)
	err := DeleteUsuario(context.Background(), usuario1.ID, database)

	if err != nil {
		t.Fatalf("ERROR DELETING USUARIO: %v", err)
	}

	usuario2, err := GetUsuario(context.Background(), usuario1.ID, database)

	if err == nil {
		t.Fatalf("ACCOUNT NOT SUCCESFULLY DELETED: %v", err)
	}

	if err != sql.ErrNoRows {
		t.Fatalf("UNKNOWN ERROR: %v", err)
	}

	if (Usuario{}) != usuario2 {
		t.Fatalf("USUARIO MUST BE EMPTY")
	}
}

func TestGetUsuarios(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUsuario(t)
	}
	arg := GetUsuariosParams{
		Limit:  5,
		Offset: 5,
	}
	usuarios, err := GetUsuarios(context.Background(), arg, database)

	if err != nil {
		t.Fatalf("THERE WAS AN ERROR: %v", err)
	}

	if len(usuarios) == 0 {
		t.Fatalf("USUARIO LIST IS EMPTY")
	}

	for _, usuario := range usuarios {
		if (Usuario{}) == usuario {
			t.Fatalf("USUARIO IS EMPTY")
		}
	}
}
