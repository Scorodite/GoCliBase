package cli

import (
	"fmt"
	"reflect"
	"slices"
)

func (c *CliEngine) WrapMethodToCallback(target any, methodName string) (callback CliCallback, args []string) {
	method := reflect.ValueOf(target).MethodByName(methodName)

	ins := slices.Collect(method.Type().Ins())
	insConvs := make([]func(string) (any, error), len(ins))

	args = make([]string, len(ins))
	for i, in := range ins {
		args[i] = in.Name()
	}

	for idx, in := range ins {
		switch in.Kind() {
		case reflect.String:
			insConvs[idx] = func(v string) (any, error) {
				return v, nil
			}
		case reflect.Int8:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[int8](c, idx, s)
			}
		case reflect.Bool:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[bool](c, idx, s)
			}
		case reflect.Int16:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[int16](c, idx, s)
			}
		case reflect.Int32:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[int32](c, idx, s)
			}
		case reflect.Int64:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[int8](c, idx, s)
			}
		case reflect.Int:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[int](c, idx, s)
			}
		case reflect.Uint8:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[uint8](c, idx, s)
			}
		case reflect.Uint16:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[uint16](c, idx, s)
			}
		case reflect.Uint32:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[uint32](c, idx, s)
			}
		case reflect.Uint64:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[uint8](c, idx, s)
			}
		case reflect.Uint:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[uint](c, idx, s)
			}
		case reflect.Float32:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[float32](c, idx, s)
			}
		case reflect.Float64:
			insConvs[idx] = func(s string) (any, error) {
				return ParseArgN[float64](c, idx, s)
			}
		default:
			wasStruct := in.Kind() == reflect.Struct
			if wasStruct {
				in = reflect.PointerTo(in)
			}

			if !in.AssignableTo(reflect.TypeFor[ParseableCliArgument]()) ||
				in.Kind() != reflect.Pointer ||
				in.Elem().Kind() != reflect.Struct {
				panic("Unsupported method argument type")
			}
			insConvs[idx] = func(s string) (any, error) {
				argR := reflect.New(in.Elem())
				arg := argR.Interface().(ParseableCliArgument)
				e := arg.ParseCli(s, target)
				if e != nil {
					return nil, e
				}
				if wasStruct {
					return argR.Elem().Interface(), nil
				}
				return arg, nil
			}
		}
	}

	outs := slices.Collect(method.Type().Outs())

	var outConv func([]reflect.Value) (string, error)

	switch len(outs) {
	case 0:
		outConv = func(v []reflect.Value) (string, error) { return "", nil }
	case 1:
		if outs[0].String() == "error" {
			outConv = func(v []reflect.Value) (string, error) {
				e := v[0].Interface()
				if e == nil {
					return "", nil
				}
				return "", e.(error)
			}
		} else if outs[0].String() == "bool" {
			outConv = func(v []reflect.Value) (string, error) {
				b := v[0].Interface().(bool)
				if b {
					return "1", nil
				}
				return "0", nil
			}
		} else {
			outConv = func(v []reflect.Value) (string, error) {
				return fmt.Sprint(v[0].Interface()), nil
			}
		}
	case 2:
		if outs[0].Kind() != reflect.String ||
			outs[1].String() != "error" {
			panic("Unsupported method return type")
		}
		outConv = func(v []reflect.Value) (string, error) {
			r := fmt.Sprint(v[0].Interface())
			e := v[1].Interface()
			if e == nil {
				return r, nil
			}
			return r, e.(error)
		}
	default:
		panic("Unsupported method return type")
	}

	callback = func(args []string) (string, error) {
		if len(args) != len(ins) {
			return "", NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		}
		vals := make([]reflect.Value, len(args))
		for i, arg := range args {
			val, e := insConvs[i](arg)
			if e != nil {
				return "", e
			}
			vals[i] = reflect.ValueOf(val)
		}
		return outConv(method.Call(vals))
	}

	return
}
