package cli

import (
	"strings"
)

type stdLibStrings struct{}

func (c *CliEngine) RegisterLibStrings() {
	lib := &stdLibStrings{}

	c.RegisterMethod(lib, "Str.Count", "Counts substrings in value", "Value", "Substring")
	c.RegisterMethod(lib, "Str.CountEls", "Counts elements in string with separators", "Value", "Separator")
	c.RegisterMethod(lib, "Str.At", "Splits value by separator and returns substring at index", "Value", "Separator", "Index")
	c.RegisterMethod(lib, "Str.Replace", "Replaces substrings in value", "Value", "From", "To")
	c.RegisterMethod(lib, "Str.Repeat", "Repeats string", "Value", "Count")
	c.RegisterMethod(lib, "Str.Contains", "Checks if value contains substring", "Value", "Substring")
	c.RegisterMethod(lib, "Str.StartsWith", "Checks if value starts with substring", "Value", "Substring")
	c.RegisterMethod(lib, "Str.EndsWith", "Checks if value ends with substring", "Value", "Substring")
	c.RegisterMethod(lib, "Str.IndexOf", "Returns index of substring in value", "Value", "Substring")
	c.RegisterMethod(lib, "Str.LastIndexOf", "Returns index of last occurensce of substring in value", "Value", "Substring")
	c.RegisterMethod(lib, "Str.Lower", "Lowers all characters in string", "Value")
	c.RegisterMethod(lib, "Str.Upper", "Uppers all characters in string", "Value")
	c.RegisterMethod(lib, "Str.Trim", "Removes characters from cutset from value", "Value", "Cutset")
	c.RegisterMethod(lib, "Str.TrimLeft", "Removes characters from cutset from value in the beginning", "Value", "Cutset")
	c.RegisterMethod(lib, "Str.TrimRight", "Removes characters from cutset from value in the endian", "Value", "Cutset")
	c.RegisterMethod(lib, "Str.Sub", "Returns substring of value with specific beginning and endian", "Value", "Begin", "End")
	c.RegisterMethod(lib, "Str.BitLength", "Gets binary length of string value", "Value")
	c.RegisterMethod(lib, "Str.Length", "Gets length of string value", "Value")
	c.RegisterMethod(lib, "Str.CharAt", "Gets character at index from value", "Value", "Index")

	c.Register("Str.Join", []string{"Separator", "Args..."}, "Concatenates all arguments in single string with separator", func(args []string) (string, error) {
		if len(args) < 2 {
			return "", nil
		}
		return strings.Join(args[1:], args[0]), nil
	})

	c.Register("Str.Concat", []string{"Args..."}, "Concatenates all arguments in single string", func(args []string) (string, error) {
		return strings.Join(args, ""), nil
	})
}

func (c *stdLibStrings) Count(value, substr string) int {
	return strings.Count(value, substr)
}

func (c *stdLibStrings) CountEls(value, sep string) int {
	if value == "" {
		return 0
	}
	return strings.Count(value, sep) + 1
}

func (c *stdLibStrings) At(value, sep string, index int) string {
	split := strings.Split(value, sep)
	if index < 0 || index >= len(split) {
		return ""
	}
	return split[index]
}

func (c *stdLibStrings) Replace(value, from, to string) string {
	return strings.Replace(value, from, to, 0)
}

func (c *stdLibStrings) Repeat(value string, count int) string {
	return strings.Repeat(value, count)
}

func (c *stdLibStrings) Contains(value, substring string) bool {
	return strings.Contains(value, substring)
}

func (c *stdLibStrings) StartsWith(value, substring string) bool {
	return strings.HasPrefix(value, substring)
}

func (c *stdLibStrings) EndsWith(value, substring string) bool {
	return strings.HasSuffix(value, substring)
}

func (c *stdLibStrings) IndexOf(value, substring string) int {
	idx := strings.Index(value, substring)
	return len([]rune(value[:idx]))
}

func (c *stdLibStrings) LastIndexOf(value, substring string) int {
	idx := strings.LastIndex(value, substring)
	return len([]rune(value[:idx]))
}

func (c *stdLibStrings) Lower(value string) string {
	return strings.ToLower(value)
}

func (c *stdLibStrings) Upper(value string) string {
	return strings.ToUpper(value)
}

func (c *stdLibStrings) Trim(value, cutset string) string {
	return strings.Trim(value, cutset)
}

func (c *stdLibStrings) TrimLeft(value, cutset string) string {
	return strings.TrimLeft(value, cutset)
}

func (c *stdLibStrings) TrimRight(value, cutset string) string {
	return strings.TrimRight(value, cutset)
}

func (c *stdLibStrings) Sub(value string, begin, end int) string {
	runes := []rune(value)
	if begin < 0 {
		begin = 0
	} else if begin > len(runes) {
		begin = len(runes)
	}
	if end < 0 {
		end = 0
	} else if end > len(runes) {
		end = len(runes)
	}
	if begin > end {
		begin = end
	}
	return string(runes[begin:end])
}

func (c *stdLibStrings) BitLength(value string) int {
	return len(value)
}

func (c *stdLibStrings) Length(value string) int {
	return len([]rune(value))
}

func (c *stdLibStrings) CharAt(value string, index int) string {
	return string([]rune{[]rune(value)[index]})
}
