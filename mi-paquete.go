package mi_paquete

// Ejemplo de función exportada
func Sumar(a, b int) int {
	return a + b
}

// Ejemplo de tipo exportado
type Calculadora struct {
	Resultado int
}

func (c *Calculadora) Multiplicar(x, y int) {
	c.Resultado = x * y
}
