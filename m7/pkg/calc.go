package сalc

import "errors"

type calculator struct {
}

func Create() *calculator {
	return &calculator{}
}

func (c *calculator) Calculate(a, b float64, action string) (float64, error) {
	if action == "+" {
		return c.sum(a, b), nil
	}
	if action == "-" {
		return c.subtr(a, b), nil
	}
	if action == "*" {
		return c.mult(a, b), nil
	}
	if action == "/" {
		result, err := c.div(a, b)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	return 0, errors.New("передан неверный оператор")
}

func (c *calculator) sum(a, b float64) float64 {
	return a + b
}

func (c calculator) subtr(a, b float64) float64 {
	return a - b
}

func (c calculator) mult(a, b float64) float64 {
	return a * b
}

func (c calculator) div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ноль в делении использовать нельзя")
	}
	return a / b, nil
}
