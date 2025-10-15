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

func ParalelCalculate(a, b float64, op string, ChanelCalculate chan struct {
	Value float64
	Err   error
}) {
	var c float64

	switch op {
	case "+":
		c = a + b
		ChanelCalculate <- struct {
			Value float64
			Err   error
		}{c, nil}
	case "-":
		c = a - b
		ChanelCalculate <- struct {
			Value float64
			Err   error
		}{c, nil}
	case "*":
		c = a * b
		ChanelCalculate <- struct {
			Value float64
			Err   error
		}{c, nil}
	case "/":
		if b == 0 {
			ChanelCalculate <- struct {
				Value float64
				Err   error
			}{0, fmt.Errorf("division by zero")}
			return
		}
		c = a / b
		ChanelCalculate <- struct {
			Value float64
			Err   error
		}{c, nil}
	default:
		ChanelCalculate <- struct {
			Value float64
			Err   error
		}{0, fmt.Errorf("unsupported operation: %s", op)}
	}
}
