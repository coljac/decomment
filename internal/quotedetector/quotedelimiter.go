package quotedetector

import (
	"errors"
	"regexp"
	"strings"
)

type QuoteDelimiter struct {
	Start string
	End   string
	Regex bool
}

var singleQuote = QuoteDelimiter{
	Start: "'",
	End:   "'",
	Regex: false,
}
var doubleQuote = QuoteDelimiter{
	Start: "\"",
	End:   "\"",
	Regex: false,
}
var backtick = QuoteDelimiter{
	Start: "`",
	End:   "`",
	Regex: false,
}
var tripleDoubleQuote = QuoteDelimiter{
	Start: "\"\"\"",
	End:   "\"\"\"",
	Regex: false,
}

var quoteFormats = map[string][]QuoteDelimiter{
	".pl": {QuoteDelimiter{
		Start: "<< ?[\"']([A-z]+)",
		End:   "",
		Regex: true,
	}, QuoteDelimiter{
		Start: "=([A-z]+) ",
		End:   "=cut",
		Regex: true,
	}, singleQuote, doubleQuote},
	".py":    {tripleDoubleQuote, singleQuote, doubleQuote},
	".c":     {doubleQuote},
	".h":     {doubleQuote},
	".cpp":   {doubleQuote},
	".hpp":   {doubleQuote},
	".go":    {doubleQuote, backtick},
	".rs":    {doubleQuote, QuoteDelimiter{Start: "r#", End: "#\""}},
	".java":  {doubleQuote, tripleDoubleQuote},
	".rb":    {singleQuote, doubleQuote, QuoteDelimiter{Start: "<<-? ?([A-z]+)", End: "\\1", Regex: true}},
	".cs":    {doubleQuote},
	".js":    {doubleQuote, singleQuote, backtick},
	".ts":    {doubleQuote, singleQuote, backtick},
	".swift": {doubleQuote, tripleDoubleQuote},
	".kt":    {doubleQuote, tripleDoubleQuote},
	".r":     {doubleQuote},
	".sh":    {doubleQuote, singleQuote, QuoteDelimiter{Start: "<< ?([A-z]+)", End: "\\1", Regex: true}},
	".tex":    {{Start: "\\%", End: " "}},
	".vim":   {},
}

func MatchPatternStrLen(r string, input string) (string, int) {
	re := regexp.MustCompile(r)
	match := re.FindStringSubmatchIndex(input)
	if len(match) == 4 {
		return input[match[2]:match[3]], match[1]
	}
	return "", 0
}

func MatchPatternStr(r string, input string) (string, error) {
	re := regexp.MustCompile(r)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1], nil
	}
	return "", errors.New("")
}

func GetQuoteDelimiters(ext string) []QuoteDelimiter {
	if delimiters, ok := quoteFormats[strings.ToLower(ext)]; ok {
		return delimiters
	}

	return nil
}

type StringSliceFlag []string

func (s *StringSliceFlag) String() string {
	return strings.Join(*s, ", ")
}

func (s *StringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}
