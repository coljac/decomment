package processor

/* TODO
- What about lines with just whitespace and comments
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/coljac/decomment/internal/commentdelimiter"
	"github.com/coljac/decomment/internal/quotedetector"
)

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func Process(reader io.Reader, filename string, inPlace bool, fileType string, commentDelimiters []commentdelimiter.CommentDelimiter, quoteDelimiters []quotedetector.QuoteDelimiter) {
	scanner := bufio.NewScanner(reader)
	// var commentDelimiters []commentdelimiter.CommentDelimiter
	// var quoteDelimiters []quotedetector.QuoteDelimiter
	var inComment bool = false
	var inQuote bool = false
	var nextDelimiter string = ""
	var quoteEnd string = ""
	var output strings.Builder

	if scanner.Scan() {
		firstLine := scanner.Text()
		if strings.HasPrefix(firstLine, "#!") {
			output.WriteString(firstLine + "\n")
			if commentDelimiters == nil {
				commentDelimiters = commentdelimiter.GuessCommentDelimiters(firstLine)
			}
		} else {
			output.WriteString(processLine(firstLine, commentDelimiters, quoteDelimiters, &inComment, &inQuote, &nextDelimiter, &quoteEnd))
		}

		for scanner.Scan() {
			line := scanner.Text()
			output.WriteString(processLine(line, commentDelimiters, quoteDelimiters, &inComment, &inQuote, &nextDelimiter, &quoteEnd))
		}
	} else {
		return
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

	if inPlace && filename != "" {
		err := os.WriteFile(filename, []byte(output.String()), 0644)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", filename, err)
		}
	} else {
		fmt.Print(output.String())
	}
}

func processLine(line string, commentDelimiters []commentdelimiter.CommentDelimiter, quoteDelimiters []quotedetector.QuoteDelimiter, inComment *bool, inQuote *bool, commentEnd *string, quoteEnd *string) string {

	toPrint := ""
	ln := len(line)

	rest:
	for i := 0; i < len(line); i++ {
		rest := line[i:]

		if *inComment {
			if strings.HasPrefix(rest, *commentEnd) {
				*inComment = false
				i += len(*commentEnd) - 1
				continue
			}
		} else if *inQuote {
			if strings.HasPrefix(rest, *quoteEnd) {
				*inQuote = false
				toPrint += *quoteEnd
				i += len(*quoteEnd) - 1
				continue
			}
		} else {
			for _, cDelim := range commentDelimiters {
				if cDelim.LineStartOnly && i > 0 {
					break
				}
				if strings.HasPrefix(rest, cDelim.Start) {
					*inComment = true
					*commentEnd = cDelim.End
					i += len(cDelim.Start) - 1
					break
					// continue rest
				}
			}

			for _, qDelim := range quoteDelimiters {
				if qDelim.Regex {
					m, l := quotedetector.MatchPatternStrLen("^"+qDelim.Start, rest)
					if l > 0 {
						*inQuote = true
						*quoteEnd = m
						i += l - 1
						toPrint += rest[:l]
						// break
						continue rest
					}
				} else {
					if strings.HasPrefix(rest, qDelim.Start) {
						*inQuote = true
						*quoteEnd = qDelim.End
						i += len(qDelim.Start) - 1
						toPrint += qDelim.Start
						// break
						continue rest
					}
				}
			}
		}

		if !*inComment {
			toPrint += string(line[i])
		}
	}

	if *inComment && *commentEnd == "\n" {
		*inComment = false
	}

	lenNonSpace := len(strings.TrimSpace(toPrint))
	if ln == len(toPrint) || (lenNonSpace > 0 || (lenNonSpace == ln)) || ln == 0 {
		return trimRightSpace(toPrint) + "\n"
	}
	return ""
}
