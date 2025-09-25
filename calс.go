package calc

import "fmt"

func Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func CalculateParalel(a, b float64, op string, ChanelFloat64 chan float64, ChanelError chan error) {
	var c float64

	switch op {
	case "+":
		c = a + b
		ChanelFloat64 <- c
	case "-":
		c = a - b
		ChanelFloat64 <- c
	case "*":
		c = a * b
		ChanelFloat64 <- c
	case "/":
		if b == 0 {
			ChanelError <- fmt.Errorf("division by zero")
		}
		c = a / b
		ChanelFloat64 <- c
	default:
		ChanelError <- fmt.Errorf("unsupported operation: %s", op)
	}

}
