package quotedetector

// Random number generator import

import (
	// "math/rand"
	"path/filepath"
	"strings"
	"errors"
	"regexp"
	// "strings"

	// "github.com/google/uuid"
	"github.com/dlclark/regexp2"
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
		End: "=cut",
		Regex: true,
	}, singleQuote, doubleQuote},
	".py":    {tripleDoubleQuote, singleQuote, doubleQuote},
	".c":     {doubleQuote},
	".cpp":   {doubleQuote},
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
}

// var singleQuote = QuoteDelimiter{Start: "'", End: "'"}
// var doubleQuote = QuoteDelimiter{Start: "\"", End: "\""}
// var backtick = QuoteDelimiter{Start: "`", End: "`"}
// var tripleSingleQuote = QuoteDelimiter{Start: "'''", End: "'''"}
// var tripleDoubleQuote = QuoteDelimiter{Start: "\"\"\"", End: "\"\"\""}

// Use regular expression to match the start and end of the here document
// var hereDocPerl = QuoteDelimiter{Start: "<<", End: "\n"}
//
// var simple = []QuoteDelimiter{singleQuote, doubleQuote, backtick}

func MatchPatternStrLen(r string, input string) (string, int) {
	re := regexp.MustCompile(r)
    match := re.FindStringSubmatchIndex(input)
    if len(match) == 4 {
        return input[match[2]:match[3]], match[1]
    }
    return "", 0
}

func MatchPatternStr(r string, input string) (string, error) {
    // re := regexp.MustCompile(`<< *([A-Za-z]+)`)
    re := regexp.MustCompile(r)
    match := re.FindStringSubmatch(input)
    if len(match) > 1 {
        return match[1], nil
    }
	return "", errors.New("")
}

func FindAllMatchesStr(r string, s string) []string {
	re := regexp2.MustCompile(r, 0)
	return FindAllMatches(re, s)
}

func FindAllMatches(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

// func replaceHereDocs(input string, language string) (string, string) {
// 	formats := hereDocFormats[language]

// 	// Keep a map of idString to the original here document for post-replacement
// 	hereDocMatches := []string{}

// 	for _, variant := range(formats) {
// 		re := regexp2.Compile(variant)
// 		hereDocMatches = regexp2FindAllString(re, input)
// 		for _, x := range(hereDocMatches) {
// 			idString := uuid.Must(uuid.NewRandom()).String()
// 			// return strings.ReplaceAll(input, token, replacement), idString
// 			input = strings.ReplaceAll(input, x, idString)
// 			hereDocMatches[idString] = x
// 		}
// 	}
// 	return input, hereDocMatches
// }

// func findHeredocToken(language, input string) (string, bool) {
// 	var re *regexp.Regexp

// 	switch language {
// 	case "perl", "ruby", "shell":
// 		re = regexp.MustCompile(`<<[-~]?\s*(?:'(\w+)'|"(\w+)"|(\w+))`)
// 	case "php":
// 		re = regexp.MustCompile(`<<<\s*(?:'(\w+)'|"(\w+)"|(\w+))`)
// 	case "python", "java":
// 		re = regexp.MustCompile(`('{3}|"{3})(\w*)`)
// 	case "lua":
// 		re = regexp.MustCompile(`\[(=*)\[\s*(\w*)`)
// 	default:
// 		return "", false
// 	}

// 	match := re.FindStringSubmatch(input)
// 	if match == nil {
// 		return "", false
// 	}

// 	// Find the non-empty capture group (the token)
// 	for _, group := range match[1:] {
// 		if group != "" {
// 			return group, true
// 		}
// 	}

// 	return "", false
// }

// var fileExtensionDelimiters = map[string][]QuoteDelimiter{
// 	".go":   simple,
// 	".c":    simple,
// 	".cpp":  simple,
// 	".rs":   simple,
// 	".java": simple,
// 	".kt":   simple,
// 	".js":   simple,
// 	".py":   simple,
// 	".sh":   simple,
// 	".pl":   simple,
// 	".rb":   simple,
// 	".sql":  {singleQuote},
// 	".php":  {},
// 	".html": {},
// 	".css":  {},
// }

func GetQuoteDelimiters(filename string) []QuoteDelimiter {
	ext := strings.ToLower(filepath.Ext(filename))
	if delimiters, ok := quoteFormats[strings.ToLower(ext)]; ok {
		return delimiters
	}

	// Default to C-style comments if extension is not recognized
	return nil
}

// func GuessQuoteDelimiters(content string) []QuoteDelimiter {
// 	if strings.HasPrefix(content, "#!/") {
// 		// If it's a script with a shebang, use '#' as the comment delimiter
// 		return hashStyle
// 	}

// 	// Default to C-style comments if no shebang is found
// 	return cStyle
// }
