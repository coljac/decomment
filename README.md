# Decommenter

This utility removes comments from a file. It figures out the comment syntax from the file extension, and outputs a file that is functionally identical but stripped of any comments.

It's smart enough to keep "comments" inside quotes, e.g. `"This will still #be in the output."`

If your source file is malformed, the output is undefined; e.g. an unclosed string terminator.

`#!` on the first line is preserved.

# Installation

Install with `go get github.com/coljac/decomment`, or download a binary from the Releases page.

# Known issues

- Mixed HTML and Javascript will not remove javascript comments.
- If the final line of the source file has no newline, one will be inserted.
- Trailing whitespace will be removed

# ToDo

Version 0.1.0 will have:

- Help/man page
- Specify file type with a flag, i.e. `decomment -f java ...`
- Custom delimiters from the command line (`decomment -c "/*,*/" my_file.foo`)
- In-place comment removal (`decomment -i query.sql`) 
- An expanded list of source file types
