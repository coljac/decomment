package commentdelimiter

import (
	"strings"
)

type CommentDelimiter struct {
	Start string
	End   string
	LineStartOnly bool
}

var cStyle = []CommentDelimiter{{Start: "//", End: "\n"}, {Start: "/*", End: "*/"}}
var hashStyle = []CommentDelimiter{{Start: "#", End: "\n"}}

var fileExtensionDelimiters = map[string][]CommentDelimiter{
	".go":   cStyle,
	".c":    cStyle,
	".h":    cStyle,
	".cpp":  cStyle,
	".hpp":  cStyle,
	".rs":   {{Start: "//", End: "\n"}},
	".java": cStyle,
	".kt":   cStyle,
	".js":   cStyle,
	".ts":   cStyle,
	".py":   hashStyle,
	".sh":   hashStyle,
	".pl":   hashStyle,
	".php":  cStyle,
	".rb":   {{Start: "#", End: "\n"}, {Start: "=begin", End: "=end"}},
	".sql":  {{Start: "--", End: "\n"}},
	".html": {{Start: "<!--", End: "-->"}},
	".css":  {{Start: "/*", End: "*/"}},
	".tex":  {{Start: "%", End: "\n"}},
	".vim":  {{Start: "\"", End: "\n", LineStartOnly: true}},
	".lua":  {{Start: "--", End: "\n"}},
}

func GetCommentDelimiters(ext string) []CommentDelimiter {
	if delimiters, ok := fileExtensionDelimiters[ext]; ok {
		return delimiters
	}

	// Default to sh-style
	return hashStyle
}

func GuessCommentDelimiters(content string) []CommentDelimiter {
	if strings.HasPrefix(content, "#!/") {
		return hashStyle
	}
	if strings.HasPrefix(content, "/*") {
		return cStyle
	}
	return hashStyle
}

type StringSliceFlag []string

func (s *StringSliceFlag) String() string {
	return strings.Join(*s, ", ")
}

func (s *StringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}
