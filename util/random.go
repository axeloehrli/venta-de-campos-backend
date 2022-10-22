package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Data for usuarios table

func RandomNombreUsuario() string {
	return fmt.Sprintf("%s_%s", RandomString(4), RandomString(4))
}

func RandomNombre() string {
	return RandomString(6)
}

func RandomApellido() string {
	return RandomString(8)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

// Data for campos table

func RandomTitulo() string {
	return fmt.Sprintf("%s %s", RandomString(4), RandomString(6))
}
func RandomTipo() string {
	tipos := []string{"Ganadero", "Agrícola", "Mixto"}
	n := len(tipos)
	return tipos[rand.Intn(n)]
}
func RandomHectareas() int64 {
	return RandomInt(100, 2000)
}

func RandomPrecioPorHectarea() int64 {
	return RandomInt(1000, 20000)
}

func RandomCiudad() string {
	return RandomString(6)
}

func RandomProvincia() string {
	return RandomString(8)
}
