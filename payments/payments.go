package payments

import (
	"github.com/stretchr/testify/mock"
)

// ********************************************************************
// *********************** INTERFACE Y ESTRUCTURA *********************
// ********************************************************************
type PaymentGateway interface {
	RequestPayment(req PaymentRequest) (PaymentResponse, error)
}

type PaymentRequest struct { // Simula la clase del ejemplo en java
	Amount float64
}

type PaymentResponse struct { // Simula la clase del ejemplo en java
	PaymentStatus
}

type PaymentStatus string // Simula el ENUM del ejemplo en java
const (
	OK    PaymentStatus = "OK-TAREA2"
	ERROR PaymentStatus = "ERROR-TAREA2"
)

// PaymentProcessor define la estructura para procesar pagos.
type PaymentProcessor struct {
	Gateway PaymentGateway // Gateway de pago utilizado para procesar el pago
}

func (p *PaymentProcessor) MakePayment(amount float64) bool {
	req := PaymentRequest{
		Amount: amount,
	}

	response, err := p.Gateway.RequestPayment(req)
	if err != nil {
		return response.PaymentStatus == ERROR
	}

	return response.PaymentStatus == OK
}

// ********************************************************************
// *********************** CREACION DE FUNCIONES MOCK *****************
// ********************************************************************

// MockPaymentGateway es un mock de PaymentGateway.
type MockPaymentGateway struct {
	mock.Mock
}

// RequestPayment simula la implementación de RequestPayment del PaymentGateway.
func (m *MockPaymentGateway) RequestPayment(req PaymentRequest) (PaymentResponse, error) {
	args := m.Called(req)
	return args.Get(0).(PaymentResponse), args.Error(1)
}
