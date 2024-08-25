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

func Process(reader io.Reader, filename string) {
	scanner := bufio.NewScanner(reader)
	var commentDelimiters []commentdelimiter.CommentDelimiter
	var quoteDelimiters []quotedetector.QuoteDelimiter
	var inComment bool = false
	var inQuote bool = false
	var nextDelimiter string = "" 
	var quoteEnd string = "" 

	// Read the first line to check for shebang
	commentDelimiters = commentdelimiter.GetCommentDelimiters(filename)
	quoteDelimiters = quotedetector.GetQuoteDelimiters(filename)
	if scanner.Scan() {
        firstLine := scanner.Text()
        if strings.HasPrefix(firstLine, "#!") {
            fmt.Println(firstLine)
            if commentDelimiters == nil {
                commentDelimiters = commentdelimiter.GuessCommentDelimiters(firstLine)
            }
		} else {
			fmt.Print(processLine(firstLine, commentDelimiters, quoteDelimiters, &inComment, &inQuote, &nextDelimiter, &quoteEnd))
        }

        for scanner.Scan() {
            line := scanner.Text()
            fmt.Print(processLine(line, commentDelimiters, quoteDelimiters, &inComment, &inQuote, &nextDelimiter, &quoteEnd))
        }
	} else {
        return
    }

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func processLine(line string, commentDelimiters []commentdelimiter.CommentDelimiter, quoteDelimiters []quotedetector.QuoteDelimiter, inComment *bool, inQuote *bool, commentEnd *string, quoteEnd *string) string {
	toPrint := ""
    ln := len(line)

    // rest:
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
                    m, l := quotedetector.MatchPatternStrLen("^" + qDelim.Start, rest)
                    if l >0 {
                        *inQuote = true
                        *quoteEnd = m
                        i += l
                        toPrint += rest[:l]
                        break
                        // continue rest
                    }
				} else {
					if strings.HasPrefix(rest, qDelim.Start) {
						*inQuote = true
                        *quoteEnd = qDelim.End
                        i += len(qDelim.Start)
						toPrint += qDelim.Start
                        break
                        // continue rest
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
