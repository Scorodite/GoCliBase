package cli

type Ast struct {
	Type     AstType
	Name     string
	Args     []Ast
	Location Location
}

type AstReader struct {
	Index  int
	Tokens []Token
}

type AstType int

const (
	AstTypeText AstType = iota
	AstTypeCommand
	AstTypeIf
	AstTypeMultiple
	AstTypeTry
	AstTypeLoop
	AstTypeWhile
)

func (r *AstReader) ReadToEnd() (res []Ast, e error) {
	var ast Ast
	for {
		ast, e = r.Next()

		if e != nil {
			if IsCliErrorOfType(e, CliErrorEOF) {
				e = nil
			}
			break
		}

		res = append(res, ast)
	}
	return
}

func (r *AstReader) Next() (a Ast, e error) {
	return r.next(0, false)
}

func (r *AstReader) next(depth int, isArg bool) (a Ast, e error) {
	t := r.peekToken()

	if depth > 255 {
		e = NewCliError("context depth limit exceeded", CliErrorDepthLimit, t.Location)
		return
	}

	// if isArg && t.Type == TokenText {
	// 	r.Index++
	// 	a.Name = t.Value
	// 	a.Type = AstTypeText
	// 	return
	// }

	for !isArg && t.Type == TokenEOL {
		r.Index++
		t = r.peekToken()
	}

	a.Location = t.Location
	ttype := t.Type

	switch ttype {
	case TokenOpen:
		r.Index++
		var subs []Ast

	subsLoop:
		for {
			var sub Ast
			sub, e = r.next(depth+1, false)
			if e != nil {
				return
			}
			subs = append(subs, sub)

			hadEol := false
		eolLoop:
			for {
				t = r.peekToken()
				switch t.Type {
				case TokenClose:
					r.Index++
					break subsLoop
				case TokenEOL:
					r.Index++
					hadEol = true
					continue eolLoop
				default:
					if hadEol {
						break eolLoop
					}
					e = NewCliError("expected close bracket or end of command", CliErrorInvalidExpression, t.Location)
					return
				}
			}
		}

		if len(subs) == 1 {
			a = subs[0]
		} else {
			a.Type = AstTypeMultiple
			a.Args = subs
		}

	case TokenText, TokenIf, TokenYield, TokenTry,
		TokenLoop, TokenWhile:
		r.Index++

		if isArg {
			a.Name = t.Value
			a.Type = AstTypeText
			return
		}

		a.Name = t.Value
		a.Type = AstTypeCommand

		for {
			at := r.peekToken()

			if at.Type == TokenClose {
				break
			}

			if at.Type == TokenEOF || at.Type == TokenEOL {
				break
			}

			var arg Ast

			arg, e = r.next(depth+1, true)

			if e != nil {
				return
			}

			a.Args = append(a.Args, arg)
		}

		switch ttype {
		case TokenIf:
			a.Type = AstTypeIf
			if len(a.Args) != 2 && len(a.Args) != 3 {
				e = NewCliError("'If' takes 2 or 3 arguments", CliErrorInvalidExpression, t.Location)
			}
		case TokenYield:
			if len(a.Args) != 1 {
				e = NewCliError("'Yield' takes exactly 1 argument", CliErrorInvalidExpression, t.Location)
			}
			a = a.Args[0]
		case TokenTry:
			a.Type = AstTypeTry
			if len(a.Args) != 1 {
				e = NewCliError("'Try' takes exactly 1 argument", CliErrorInvalidExpression, t.Location)
			}
		case TokenLoop:
			a.Type = AstTypeLoop
			if len(a.Args) != 1 {
				e = NewCliError("'Loop' takes exactly 1 argument", CliErrorInvalidExpression, t.Location)
			}
		case TokenWhile:
			a.Type = AstTypeWhile
			if len(a.Args) != 2 {
				e = NewCliError("'While' takes exactly 2 arguments", CliErrorInvalidExpression, t.Location)
			}
		}
	case TokenEOF:
		e = NewCliError("end of file", CliErrorEOF, t.Location)
	default:
		e = NewCliError("invalid expression", CliErrorInvalidExpression, t.Location)
	}
	return
}

func (r *AstReader) peekToken() Token {
	if r.Index >= len(r.Tokens) {
		eof := Token{Type: TokenEOF}

		if len(r.Tokens) > 0 {
			eof.Location = r.Tokens[len(r.Tokens)-1].Location
		}

		return eof
	}
	return r.Tokens[r.Index]
}

func (r *AstReader) IsEOF() bool {
	for r.peekToken().Type == TokenEOL {
		r.Index++
	}
	return r.peekToken().Type == TokenEOF
}
