package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func assertEquals(actual, expected int) bool {
	if actual == expected {
		fmt.Println("Si es igual")
		return true
	} else {
		fmt.Printf("NO es igual  Expected: %d, Actual: %d\n", expected, actual)
		return false
	}
}

func main() {
	actual := add(2, 4)
	esperado := 6
	resultadoComparacion := assertEquals(actual, esperado)
	fmt.Println(resultadoComparacion)

	resultadoJunit := assertEquals(6, add(3, 3))
	fmt.Println(resultadoJunit)

}


package calculadora

import "testing"

// Función a probar
func Sumar(a int, b int) int {
    return a + b
}

// Función de prueba
func TestSumar(t *testing.T) {
    resultado := Sumar(2, 3)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Sumar(2, 3) = %d; se esperaba %d", resultado, esperado)
    }
}
