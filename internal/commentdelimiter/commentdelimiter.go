package commentdelimiter

import (
	"path/filepath"
	"strings"
)

type CommentDelimiter struct {
	Start string
	End   string
}

var cStyle = []CommentDelimiter{{Start: "//", End: "\n"}, {Start: "/*", End: "*/"}}
var hashStyle = []CommentDelimiter{{Start: "#", End: "\n"}}

var fileExtensionDelimiters = map[string][]CommentDelimiter{
	".go":   cStyle,
	".c":    cStyle,
	".cpp":  cStyle,
	".rs":   {{Start: "//", End: "\n"}},
	".java": cStyle,
	".kt":   cStyle,
	".js":   cStyle,
	".ts":   cStyle,
	".py":   hashStyle,
	".sh":   hashStyle,
	".pl":   hashStyle,
	".rb":   {{Start: "#", End: "\n"}, {Start: "=begin", End: "=end"}},
	".sql":  {{Start: "--", End: "\n"}},
	".php":  {{Start: "//", End: "\n"}, {Start: "/*", End: "*/"}},
	".html": {{Start: "<!--", End: "-->"}},
	".css":  {{Start: "/*", End: "*/"}},
}

func GetCommentDelimiters(filename string) []CommentDelimiter {
	ext := strings.ToLower(filepath.Ext(filename))
	if delimiters, ok := fileExtensionDelimiters[ext]; ok {
		return delimiters
	}

	// Default to C-style comments if extension is not recognized
	return nil
}

func GuessCommentDelimiters(content string) []CommentDelimiter {
	if strings.HasPrefix(content, "#!/") {
		// If it's a script with a shebang, use '#' as the comment delimiter
		return hashStyle
	}

	// Default to C-style comments if no shebang is found
	return cStyle
}
