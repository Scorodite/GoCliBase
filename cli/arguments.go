package cli

import "strconv"

func ParseArgs1[T1 any](c *CliEngine, args []string) (r1 T1, e error) {
	if len(args) != 1 {
		e = NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		return
	}

	r1, e = ParseArgN[T1](c, 0, args[0])

	return
}

func ParseArgs2[T1, T2 any](c *CliEngine, args []string) (r1 T1, r2 T2, e error) {
	if len(args) != 2 {
		e = NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		return
	}

	r1, e = ParseArgN[T1](c, 0, args[0])
	if e != nil {
		return
	}

	r2, e = ParseArgN[T2](c, 1, args[1])

	return
}

func ParseArgs3[T1, T2, T3 any](c *CliEngine, args []string) (r1 T1, r2 T2, r3 T3, e error) {
	if len(args) != 3 {
		e = NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		return
	}

	r1, e = ParseArgN[T1](c, 0, args[0])
	if e != nil {
		return
	}

	r2, e = ParseArgN[T2](c, 1, args[1])
	if e != nil {
		return
	}

	r3, e = ParseArgN[T3](c, 2, args[2])

	return
}

func ParseArgs4[T1, T2, T3, T4 any](c *CliEngine, args []string) (r1 T1, r2 T2, r3 T3, r4 T4, e error) {
	if len(args) != 4 {
		e = NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		return
	}

	r1, e = ParseArgN[T1](c, 0, args[0])
	if e != nil {
		return
	}

	r2, e = ParseArgN[T2](c, 1, args[1])
	if e != nil {
		return
	}

	r3, e = ParseArgN[T3](c, 2, args[2])
	if e != nil {
		return
	}

	r4, e = ParseArgN[T4](c, 3, args[3])

	return
}

func ParseArgs5[T1, T2, T3, T4, T5 any](c *CliEngine, args []string) (r1 T1, r2 T2, r3 T3, r4 T4, r5 T5, e error) {
	if len(args) != 5 {
		e = NewCliError("invalid argument count", CliErrorInvalidArgumentCount, Location{})
		return
	}

	r1, e = ParseArgN[T1](c, 0, args[0])
	if e != nil {
		return
	}

	r2, e = ParseArgN[T2](c, 1, args[1])
	if e != nil {
		return
	}

	r3, e = ParseArgN[T3](c, 2, args[2])
	if e != nil {
		return
	}

	r4, e = ParseArgN[T4](c, 3, args[3])
	if e != nil {
		return
	}

	r5, e = ParseArgN[T5](c, 4, args[4])

	return
}

func ParseArgN[T any](c *CliEngine, i int, value string) (r T, e error) {
	var (
		pI int64
		pU uint64
		pF float64
		pB bool
	)

	var name string = "?"
	if i >= 0 && i < len(c.currentCommand.arguments) {
		name = c.currentCommand.arguments[i]
	}

	switch any(r).(type) {
	case ParseableCliArgument:
		var p ParseableCliArgument = any(r).(ParseableCliArgument)
		e = p.ParseCli(value, p)
		if e == nil {
			r = p.(T)
		}
	case string:
		r = any(value).(T)
	case bool:
		pB, e = strconv.ParseBool(value)
		if e != nil {
			e = NewCliErrorWr("failed to parse boolean argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(bool(pB)).(T)
	case int:
		pI, e = strconv.ParseInt(value, 10, 64)
		if e != nil {
			e = NewCliErrorWr("failed to parse integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(int(pI)).(T)
	case int8:
		pI, e = strconv.ParseInt(value, 10, 8)
		if e != nil {
			e = NewCliErrorWr("failed to parse integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(int8(pI)).(T)
	case int16:
		pI, e = strconv.ParseInt(value, 10, 16)
		if e != nil {
			e = NewCliErrorWr("failed to parse integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(int16(pI)).(T)
	case int32:
		pI, e = strconv.ParseInt(value, 10, 32)
		if e != nil {
			e = NewCliErrorWr("failed to parse integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(int32(pI)).(T)
	case int64:
		pI, e = strconv.ParseInt(value, 10, 64)
		if e != nil {
			e = NewCliErrorWr("failed to parse integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(int64(pI)).(T)
	case uint:
		pU, e = strconv.ParseUint(value, 10, 64)
		if e != nil {
			e = NewCliErrorWr("failed to parse unsigned-integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(uint(pU)).(T)
	case uint8:
		pU, e = strconv.ParseUint(value, 10, 8)
		if e != nil {
			e = NewCliErrorWr("failed to parse unsigned-integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(uint8(pU)).(T)
	case uint16:
		pU, e = strconv.ParseUint(value, 10, 16)
		if e != nil {
			e = NewCliErrorWr("failed to parse unsigned-integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(uint16(pU)).(T)
	case uint32:
		pU, e = strconv.ParseUint(value, 10, 32)
		if e != nil {
			e = NewCliErrorWr("failed to parse unsigned-integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(uint32(pU)).(T)
	case uint64:
		pU, e = strconv.ParseUint(value, 10, 64)
		if e != nil {
			e = NewCliErrorWr("failed to parse unsigned-integer argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(uint64(pU)).(T)
	case float32:
		pF, e = strconv.ParseFloat(value, 32)
		if e != nil {
			e = NewCliErrorWr("failed to parse float argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(float32(pF)).(T)
	case float64:
		pF, e = strconv.ParseFloat(value, 64)
		if e != nil {
			e = NewCliErrorWr("failed to parse float argument: "+name, CliErrorArgumentParseFailed, e)
			return
		}
		r = any(float64(pF)).(T)
	default:
		e = NewCliError("unsupported type", CliErrorArgumentParseFailed, Location{})
	}
	return
}
