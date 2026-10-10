package cli

import (
	"errors"
	"fmt"
	"strings"
)

type CliErrorType int

const (
	CliErrorEOF CliErrorType = iota
	CliErrorInvalidExpression
	CliErrorDepthLimit
	CliErrorUndefinedCommand
	CliErrorMultipleExpressions
	CliErrorArgumentParseFailed
	CliErrorInvalidArgumentCount
	CliErrorInternal
	CliErrorUndefinedAstType
	CliErrorBreak
	CliErrorContinue
)

type CliError struct {
	Message  string
	Location Location
	Type     CliErrorType
	Base     error
}

func (e CliError) Error() string {
	message := e.Message
	if e.Base != nil {
		baseMessage := e.Base.Error()
		if baseMessage != message {
			message = fmt.Sprintf("%s (%s)", message, baseMessage)
		}
	}
	if e.Location.IsEmpty() {
		return message
	}
	return fmt.Sprintf("%s [%s]", message, e.Location)
}

func (e CliError) Unwrap() []error {
	if e.Base == nil {
		return nil
	}
	return []error{e.Base}
}

func NewCliError(message string, ttype CliErrorType, location Location) CliError {
	return CliError{Message: message, Type: ttype, Location: location}
}

func NewCliErrorWr(message string, ttype CliErrorType, base error) CliError {
	return CliError{Message: message, Type: ttype, Base: base}
}

func IsCliErrorOfType(e error, ttype CliErrorType) bool {
	var cliError CliError
	if !errors.As(e, &cliError) {
		return false
	}
	return cliError.Type == ttype
}

func PrintCliError(code string, e error) {
	fmt.Printf("\tError: %s\n\n", e)
	cliE, ok := e.(CliError)
	if !ok {
		println("\t", e.Error())
		return
	}
	if cliE.Location.IsEmpty() || code == "" {
		return
	}
	lines := strings.Split(code, "\n")
	if cliE.Location.Line < 0 || cliE.Location.Line >= len(lines) {
		return
	}
	line := lines[cliE.Location.Line]

	offsIdx := 0
	for _, l := range lines[:cliE.Location.Line] {
		offsIdx += len(l) + 1
	}
	lineIdx := cliE.Location.Index - offsIdx
	lineLen := cliE.Location.Length

	if lineIdx < 0 {
		lineIdx = 0
	} else if lineIdx > len(line) {
		lineIdx = len(line)
	}
	if lineLen > len(line)-lineIdx {
		lineLen = 0
	}

	linePos := fmt.Sprint(cliE.Location.Line + 1)

	fmt.Printf("\t%s| %s\n", linePos, line)

	if lineLen <= 0 {
		return
	}

	fmt.Println("\t" +
		strings.Repeat(" ", lineIdx+len(linePos)+2) +
		strings.Repeat("^", lineLen))
}
