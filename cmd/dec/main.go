package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/coljac/decomment/internal/commentdelimiter"
	"github.com/coljac/decomment/internal/processor"
	"github.com/coljac/decomment/internal/quotedetector"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help information")
	fileType := flag.String("f", "", "Specify file or input type")
	inPlace := flag.Bool("i", false, "Change file in place")

	var commentDelimiters []commentdelimiter.CommentDelimiter
	var quoteDelimiters []quotedetector.QuoteDelimiter
	var cDelim commentdelimiter.StringSliceFlag
	var qDelim quotedetector.StringSliceFlag

	flag.Var(&cDelim, "c", "Add a comment delimiter(s); format=<start>,[<end>] (can be used multiple times)")
	flag.Var(&qDelim, "q", "Set quote delimiter(s) (can be used multiple times)")
	flag.Parse()

	// If filetype doesn't start with a ".", prepend it
	if len(*fileType) > 0 && !strings.HasPrefix(*fileType, ".") {
		*fileType = "." + *fileType
	}

	if len(cDelim) > 0 {
		commentDelimiters = make([]commentdelimiter.CommentDelimiter, len(cDelim))
		for i, delim := range cDelim {
			tokens := strings.Split(delim, ",")
			start := tokens[0]
			end := "\n"
			if len(tokens) > 1 {
				end = tokens[1]
			}
			commentDelimiters[i] = commentdelimiter.CommentDelimiter{Start: start, End: end}
		}
	}

	if *helpFlag {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [FILE...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Remove comments from files or stdin.\n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Use remaining arguments as input files
	args := flag.Args()

	if len(args) > 0 {
		for _, path := range args {
			file, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", path, err)
				continue
			}
			if *fileType == "" {
				*fileType = strings.ToLower(filepath.Ext(path))
			}
			if commentDelimiters == nil {
				commentDelimiters = commentdelimiter.GetCommentDelimiters(*fileType)
			}
			if quoteDelimiters == nil {
				quoteDelimiters = quotedetector.GetQuoteDelimiters(*fileType)
			}
			defer file.Close()
			processor.Process(file, path, *inPlace, *fileType, commentDelimiters, quoteDelimiters)
		}
	} else {
		if *fileType == "" {
			*fileType = ".sh"
		}
		if commentDelimiters == nil {
			commentDelimiters = commentdelimiter.GetCommentDelimiters(*fileType)
		}
		processor.Process(os.Stdin, "", false, *fileType, commentDelimiters, quoteDelimiters)
	}
}
