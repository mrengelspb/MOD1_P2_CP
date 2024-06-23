package calculadora

import (
	"errors"
	"time"

	"github.com/stretchr/testify/mock"
)

// ********************************************************************
// *********************** INTERFACE Y ESTRUCTURA *********************
// ********************************************************************
type Calcu interface {
	Sumar(a, b int) int
}

type CalculadoraReal struct{}

func (c CalculadoraReal) Sumar(a, b int) int {
	return a + b
}

// ********************************************************************
// *********************** FUNCIONES de TESTIFY MOCK ******************
// ********************************************************************

// CalculadoraMock es un mock de la interfaz Calculadora
type CalculadoraMock struct {
	mock.Mock
}

// Sumar es el método mock de Sumar
func (m *CalculadoraMock) Sumar(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

// ********************************************************************
// *********************** FUNCIONES SIN MOCKS ***********************
// ********************************************************************
//func Sumar(a, b int) int {
//	return a + b
//}

func Restar(a, b int) int {
	return a - b
}
func Multiplicar(a, b int) int {
	time.Sleep(10 * time.Second)
	return a * b
}

func Dividir(a, b int) (resultado int, err error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}
