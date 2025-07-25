package mi_paquete

import "fmt"

type Calculadora struct {
	Resultado int
}

// Ejemplo de función exportada
func Sumar(a, b int) int {
	return a + b
}

func (c *Calculadora) Restar(x, y int) {
	c.Resultado = x - y
}

func (c *Calculadora) Multiplicar(x, y int) {
	c.Resultado = x * y
}

func Dividir(x, y float64) (float64, error) {
	if y == 0 || y == 0.0 {

		return 0, fmt.Errorf("No se puede dividir por 0")
	}
	return x / y, nil
}
