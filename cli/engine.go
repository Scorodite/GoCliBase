package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type CliEngine struct {
	commands       map[string]CliCommand
	currentCommand CliCommand
}

func NewCliEngine() *CliEngine {
	res := &CliEngine{
		commands: make(map[string]CliCommand),
	}

	res.Register("If", []string{"Condition: bool", "TrueValue", "FalseValue"}, "Returns one of two values desired by condition. Does not execute unused value", func(args []string) (string, error) {
		return "", fmt.Errorf("attempted to call 'If' command directly")
	})
	res.Register("Loop", []string{"Body"}, "Repeats body until 'Break' command called", func(args []string) (string, error) {
		return "", fmt.Errorf("attempted to call 'Loop' command directly")
	})
	res.Register("While", []string{"Condition: bool", "Body"}, "Repeats body while condition is true or until 'Break' command called", func(args []string) (string, error) {
		return "", fmt.Errorf("attempted to call 'While' command directly")
	})
	res.Register("Yield", []string{"Value"}, "Returns provided argument", func(args []string) (string, error) {
		if len(args) != 1 {
			return "", fmt.Errorf("'Yield' takes exactly one argument")
		}
		return args[0], nil
	})
	res.Register("Try", []string{"Expression"}, "If expression succeded, returns value, otherwise returns error with 'error:' prefix", func(args []string) (string, error) {
		if len(args) != 1 {
			return "", fmt.Errorf("'Try' takes exactly one argument")
		}
		return args[0], nil
	})
	res.Register("Break", []string{"Value?"}, "Breaks out of loop, able to make loop return the provided value", func(args []string) (string, error) {
		if len(args) > 1 {
			return "", NewCliError("'Break' takes 1 or no arguments", CliErrorInvalidExpression, Location{})
		}
		var val string
		if len(args) > 0 {
			val = args[0]
		}
		return val, NewCliError("used 'break' outside the loop", CliErrorBreak, Location{})
	})
	res.Register("Continue", nil, "Continues loop from the beginning", func(args []string) (string, error) {
		if len(args) != 0 {
			return "", NewCliError("'Continue' takes no arguments", CliErrorInvalidExpression, Location{})
		}
		return "", NewCliError("used 'continue' outside the loop", CliErrorContinue, Location{})
	})

	return res
}

func (c *CliEngine) GetCommands() map[string]CliCommand {
	cp := map[string]CliCommand{}
	for k, v := range c.commands {
		cp[k] = v
	}
	return cp
}

func (c *CliEngine) Register(name string, args []string, about string, callback CliCallback) {
	if name == "" {
		panic("Command name is empty")
	}
	if callback == nil {
		panic("Callback is nil")
	}
	_, exists := c.commands[name]
	if exists {
		panic(fmt.Sprintf("Command with name '%s' already exists", name))
	}
	cmd := CliCommand{
		name:      name,
		arguments: args,
		callback:  callback,
		about:     about,
	}
	c.commands[name] = cmd
}

func (c *CliEngine) RegisterMethod(target any, command string, about string, args ...string) {
	methodName := command
	if strings.Contains(command, ".") {
		_, methodName, _ = strings.CutLast(command, ".")
	}
	callback, argTypes := c.WrapMethodToCallback(target, methodName)
	if len(argTypes) != len(args) {
		panic("Invalid argument count")
	}
	argsMod := make([]string, len(args))
	for i, arg := range args {
		argsMod[i] = arg
		if argTypes[i] == "string" {
			continue
		}
		argsMod[i] += ": " + argTypes[i]
	}
	c.Register(command, argsMod, about, callback)
}

func (c *CliEngine) Eval(expr string) (string, error) {
	tokenReader := TokenReader{Source: expr}
	astReader := AstReader{Tokens: tokenReader.ReadToEnd()}
	ast, e := astReader.Next()
	if e != nil {
		return "", e
	}
	if !astReader.IsEOF() {
		return "", NewCliError("multiple expressions", CliErrorMultipleExpressions, astReader.peekToken().Location)
	}

	return c.EvalAst(ast)
}

func (c *CliEngine) Exec(code string) (r string, e error) {
	tokenReader := TokenReader{Source: code}
	astReader := AstReader{Tokens: tokenReader.ReadToEnd()}
	var asts []Ast
	asts, e = astReader.ReadToEnd()
	if e != nil {
		return
	}

	for _, ast := range asts {
		r, e = c.EvalAst(ast)
		if e != nil {
			return
		}
	}

	return
}

func (c *CliEngine) EvalAst(ast Ast) (r string, e error) {
	r, e = c.evalAst(ast)

	if e != nil {
		cliE, ok := e.(CliError)
		if ok && cliE.Location.IsEmpty() {
			cliE.Location = ast.Location
			e = cliE
		}
		if !ok {
			cliE = NewCliErrorWr("internal error", CliErrorInternal, e)
			cliE.Location = ast.Location
			e = cliE
		}
	}

	return
}
func (c *CliEngine) evalAst(ast Ast) (string, error) {
	switch ast.Type {
	case AstTypeText:
		return ast.Name, nil
	case AstTypeCommand:
		return c.evalCommand(ast)
	case AstTypeIf:
		return c.evalIf(ast)
	case AstTypeMultiple:
		return c.evalMultiple(ast)
	case AstTypeTry:
		return c.evalTry(ast)
	case AstTypeLoop:
		return c.evalLoop(ast)
	case AstTypeWhile:
		return c.evalWhile(ast)
	default:
		return "", NewCliError("undefined AST type", CliErrorUndefinedAstType, ast.Location)
	}
}

func (c *CliEngine) evalCommand(ast Ast) (string, error) {
	com, ok := c.commands[ast.Name]
	if !ok {
		return "", NewCliError("undefined command: "+ast.Name, CliErrorUndefinedCommand, ast.Location)
	}

	if len(ast.Args) == 0 {
		return com.callback(nil)
	}

	argsEvaled := make([]string, 0, len(ast.Args))
	for _, argAst := range ast.Args {
		argEvaled, e := c.EvalAst(argAst)
		if e != nil {
			return "", e
		}
		argsEvaled = append(argsEvaled, argEvaled)
	}

	c.currentCommand = com
	res, e := com.callback(argsEvaled)

	if e != nil {
		cliError, ok := e.(CliError)

		if !ok {
			cliError = NewCliErrorWr(e.Error(), CliErrorInternal, e)
		}

		if cliError.Location.IsEmpty() {
			cliError.Location = ast.Location
		}

		e = cliError
	}
	return res, e
}

func (c *CliEngine) evalIf(ast Ast) (string, error) {
	if len(ast.Args) != 2 && len(ast.Args) != 3 {
		return "", NewCliError("'If' takes 2 or 3 arguments", CliErrorInvalidExpression, ast.Location)
	}

	cond, e := c.EvalAst(ast.Args[0])
	if e != nil {
		return "", e
	}

	condBool, e := strconv.ParseBool(cond)
	if e != nil {
		return "", e
	}

	if condBool {
		return c.EvalAst(ast.Args[1])
	}

	if len(ast.Args) < 3 {
		return "", nil
	}

	return c.EvalAst(ast.Args[2])
}

func (c *CliEngine) evalYield(ast Ast) (string, error) {
	if len(ast.Args) != 1 {
		return "", NewCliError("'Yield' takes exactly 1 argument", CliErrorInvalidExpression, ast.Location)
	}
	if ast.Args[0].Type == AstTypeText {
		return ast.Args[0].Name, nil
	}
	return c.EvalAst(ast.Args[0])
}

func (c *CliEngine) evalMultiple(ast Ast) (r string, e error) {
	for _, arg := range ast.Args {
		r, e = c.EvalAst(arg)
		if e != nil {
			break
		}
	}
	return
}

func (c *CliEngine) evalTry(ast Ast) (string, error) {
	if len(ast.Args) != 1 {
		return "", NewCliError("'Try' takes exactly 1 argument", CliErrorInvalidExpression, ast.Location)
	}
	r, e := c.EvalAst(ast.Args[0])
	if e != nil {
		return fmt.Sprintf("error:%s", e), nil
	}
	return r, e
}

func (c *CliEngine) evalLoop(ast Ast) (string, error) {
	if len(ast.Args) != 1 {
		return "", NewCliError("'Loop' takes exactly 1 argument", CliErrorInvalidExpression, ast.Location)
	}
	body := ast.Args[0]
	for {
		r, e := c.EvalAst(body)
		if e != nil {
			if IsCliErrorOfType(e, CliErrorBreak) {
				return r, nil
			}
			if IsCliErrorOfType(e, CliErrorContinue) {
				continue
			}
			return "", e
		}
	}
}

func (c *CliEngine) evalWhile(ast Ast) (string, error) {
	if len(ast.Args) != 2 {
		return "", NewCliError("'While' takes exactly 2 arguments", CliErrorInvalidExpression, ast.Location)
	}
	condition, body := ast.Args[0], ast.Args[1]
	for {
		condRes, e := c.EvalAst(condition)
		if e != nil {
			return "", e
		}
		condBool, e := strconv.ParseBool(condRes)
		if e != nil {
			return "", e
		}
		if !condBool {
			return "", nil
		}

		r, e := c.EvalAst(body)
		if e != nil {
			if IsCliErrorOfType(e, CliErrorBreak) {
				return r, nil
			}
			if IsCliErrorOfType(e, CliErrorContinue) {
				continue
			}
			return "", e
		}
	}
}
