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

func CalculateManyVar(a []float64, op string) (float64, error) {
	var Result float64
	var err error

	for _, v := range a {
		switch op {
		case "+":
			Result += v
		case "-":
			Result -= v
		case "*":
			Result *= v
		case "/":
			if v == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			Result /= v
		default:
			return 0, fmt.Errorf("unsupported operation: %s", op)
		}
	}

	return Result, err

}

func ParalelCalculateManyVar(a []float64, op string, ChanelCalculate chan struct {
	Value float64
	Err   error
}) {
	var Result float64

	for _, v := range a {
		switch op {
		case "+":
			Result += v
		case "-":
			Result -= v
		case "*":
			Result *= v
		case "/":
			if v == 0 {
				ChanelCalculate <- struct {
					Value float64
					Err   error
				}{0, fmt.Errorf("division by zero")}
				return
			}
			Result /= v
		default:
			ChanelCalculate <- struct {
				Value float64
				Err   error
			}{0, fmt.Errorf("unsupported operation: %s", op)}
			return
		}
	}

	ChanelCalculate <- struct {
		Value float64
		Err   error
	}{Result, nil}

}
