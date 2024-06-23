package main

import (
	"mod1tarea2/calculadora"
	"mod1tarea2/payments"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDadoDosEnteros_CuandoSuma_EntoncesOk(t *testing.T) {
	casosDePrueba := []struct {
		nombre   string
		entero1  int
		entero2  int
		esperado int
	}{
		{"Suma positivos", 4, 2, 6},
		{"Suma negativos", -1, -1, -2},
		{"Suma mixto", -1, 1, 0},
		{"Suma Ejemplo Video 1", 2, 4, 6},
		{"Suma Ejemplo Video 2", 1, 8, 9},
		{"Suma Ejemplo Video 3", 2, 5, 7},
		{"Suma Ejemplo Video 4", 7, 3, 10},
		{"Suma Ejemplo Video 5", 22, 4, 26},
		{"Suma Ejemplo Video 6", 12, 14, 26},
	}

	cal := calculadora.CalculadoraReal{} // Inicialización de la instancia de CalculadoraReal

	for _, caso := range casosDePrueba {
		t.Run(caso.nombre, func(t *testing.T) {
			resultado := cal.Sumar(caso.entero1, caso.entero2)
			if resultado != caso.esperado {
				t.Errorf("Sumar(%d, %d) = %d; lo esperado es: %d", caso.entero1, caso.entero2, resultado, caso.esperado)
			}
		})
	}
}

func TestDadoDosEnteros_CuandoResta_EntoncesOk(t *testing.T) {
	entero1 := 4
	entero2 := 2
	resultado := calculadora.Restar(entero1, entero2)

	esperado := 2
	if resultado != esperado {
		t.Errorf("Restar(%d, %d) = %d; lo esperado es: %d", entero1, entero2, resultado, esperado)
	}
}

func TestDadoDosEnteros_CuandoDivide_EntoncesOk(t *testing.T) {
	entero1 := 4
	entero2 := 2
	resultado, err := calculadora.Dividir(entero1, entero2)

	if err != nil {
		t.Fatalf("Se esperaba que no hubiera un error, pero se obtuvo uno: %v", err)
	}

	esperado := 2
	if resultado != esperado {
		t.Errorf("Dividir(%d, %d) = %d; lo esperado es: %d", entero1, entero2, resultado, esperado)
	}
}

func TestDadoDosEnteros_CuandoDividePorCero_EntoncesError(t *testing.T) {
	entero1 := 4
	entero2 := 0
	_, err := calculadora.Dividir(entero1, entero2)

	if err == nil {
		t.Errorf("Se esperaba un error al dividir por cero, pero no se obtuvo ninguno")
	}
}
func TestDadoDosEnteros_CuandoMultiplica_EntoncesOk(t *testing.T) {
	entero1 := 4
	entero2 := 2
	resultadoChan := make(chan int)
	go func() {
		resultado := calculadora.Multiplicar(entero1, entero2)
		resultadoChan <- resultado
	}()

	select {
	case resultado := <-resultadoChan:
		esperado := 8
		if resultado != esperado {
			t.Errorf("Multiplicar(%d, %d) = %d; lo esperado es: %d", entero1, entero2, resultado, esperado)
		}
	case <-time.After(15 * time.Second):
		t.Error("La función Multiplicar tardó más de 5 segundos en responder")
	}
}

// ********************************************************************
// *********************** FUNCIONES DE PRUEBA CON MOCKS **************
// ********************************************************************
func TestDadoDosEnteros_CuandoSuma_EntoncesOk_ConMock(t *testing.T) {
	mockCalc := new(calculadora.CalculadoraMock)

	// Configura el mock para esperar una llamada específica y devolver un resultado específico
	mockCalc.On("Sumar", 4, 2).Return(6)

	// Usa el mock en tu prueba
	resultado := mockCalc.Sumar(4, 2)

	// Verifica que el resultado es el esperado
	assert.Equal(t, 6, resultado)

	// Asegúrate de que todas las expectativas sobre el mock se cumplieron
	mockCalc.AssertExpectations(t)
}

// TestMakePaymentOK
func Test_CuandoMakePaymentOK_ConMock(t *testing.T) {
	mockGateway := new(payments.MockPaymentGateway)
	processor := payments.PaymentProcessor{Gateway: mockGateway}
	// Configura el mock para devolver un estado OK cuando se llame a RequestPayment.
	mockGateway.On("RequestPayment", mock.AnythingOfType("PaymentRequest")).Return(payments.PaymentResponse{PaymentStatus: payments.OK}, nil)
	result := processor.MakePayment(100.0) // Realiza la prueba
	// Verifica que el resultado sea verdadero y que el mock haya sido llamado correctamente.
	assert.True(t, result)
	mockGateway.AssertExpectations(t)
}

func Test_CuandoMakePaymentERROR_ConMock(t *testing.T) {
	mockGateway := new(payments.MockPaymentGateway)
	processor := payments.PaymentProcessor{Gateway: mockGateway}
	// Configura el mock para devolver un estado BAD cuando se llame a RequestPayment.
	mockGateway.On("RequestPayment", mock.AnythingOfType("PaymentRequest")).Return(payments.PaymentResponse{PaymentStatus: payments.ERROR}, nil)
	result := processor.MakePayment(100.0) // Realiza la prueba
	// Verifica que el resultado sea falso y que el mock haya sido llamado correctamente.
	assert.False(t, result)
	mockGateway.AssertExpectations(t)
}
