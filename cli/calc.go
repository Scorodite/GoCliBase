package cli

import (
	"fmt"
	"strconv"
)

type operatorCallback func(a, b string) (string, error)

type operator struct {
	Symbol   string
	Order    int
	Callback operatorCallback
}

var operators = (func() map[string]operator {
	ops := []map[string]operatorCallback{
		{
			"..": func(a, b string) (string, error) {
				return a + b, nil
			},
		},
		{
			"*": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					r = fmt.Sprint(x * y)
				}
				return
			},
			"/": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					r = fmt.Sprint(x / y)
				}
				return
			},
		},
		{
			"+": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					r = fmt.Sprint(x + y)
				}
				return
			},
			"-": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					r = fmt.Sprint(x - y)
				}
				return
			},
		},
		{
			">": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					if x > y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
			"<": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					if x < y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
			">=": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					if x >= y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
			"<=": func(a, b string) (r string, e error) {
				var x, y float64
				x, y, e = parseFloats(a, b)
				if e == nil {
					if x <= y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
		},
		{
			"==": func(a, b string) (r string, e error) {
				if a == b {
					r = "1"
				} else {
					r = "0"
				}
				return
			},
			"!=": func(a, b string) (r string, e error) {
				if a == b {
					r = "0"
				} else {
					r = "1"
				}
				return
			},
		},
		{
			"&&": func(a, b string) (r string, e error) {
				var x, y bool
				x, y, e = parseBools(a, b)
				if e == nil {
					if x && y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
			"||": func(a, b string) (r string, e error) {
				var x, y bool
				x, y, e = parseBools(a, b)
				if e == nil {
					if x || y {
						r = "1"
					} else {
						r = "0"
					}
				}
				return
			},
		},
	}
	opMap := map[string]operator{}

	for order, callbacks := range ops {
		for symbol, callback := range callbacks {
			opMap[symbol] = operator{
				Symbol:   symbol,
				Order:    order,
				Callback: callback,
			}
		}
	}

	return opMap
})()

func calc(args []string) (string, error) {
	if len(args)%2 != 1 {
		return "", fmt.Errorf("invalid argument count")
	}

	for len(args) > 1 {
		i := -1
		mostOrder := 10000
		mostOrderI := 0

		for i < len(args)-2 {
			i += 2

			op, ok := operators[args[i]]
			if !ok {
				return "", fmt.Errorf("invalid operator %s", args[i])
			}

			if mostOrder <= op.Order {
				continue
			}

			mostOrderI = i
			mostOrder = op.Order
		}

		op := operators[args[mostOrderI]]
		res, e := op.Callback(args[mostOrderI-1], args[mostOrderI+1])
		if e != nil {
			return "", e
		}
		args[mostOrderI-1] = res
		args = append(args[:mostOrderI], args[mostOrderI+2:]...)
	}

	return args[0], nil
}

func parseFloats(a, b string) (x, y float64, e error) {
	x, e = strconv.ParseFloat(a, 64)
	if e != nil {
		return
	}
	y, e = strconv.ParseFloat(b, 64)
	return
}

func parseBools(a, b string) (x, y bool, e error) {
	x, e = strconv.ParseBool(a)
	if e != nil {
		return
	}
	y, e = strconv.ParseBool(b)
	return
}
