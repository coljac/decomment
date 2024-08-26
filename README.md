# Decommenter

This utility removes comments from a file. It figures out the comment syntax from the file extension, and outputs a file that is functionally identical but stripped of any comments.

It's smart enough to keep "comments" inside quotes, e.g. `"This will still #be in the output."`

If your source file is malformed, the output is undefined; e.g. an unclosed string terminator.

`#!` on the first line is preserved.

This is an alpha version of the software, use with caution.

# Installation

Install with `go install github.com/coljac/decomment/cmd/dec@v0.0.2`, or download a binary from the Releases page.

# Usage

When passed one more files as arguments, `dec` will figure out the syntax from the file extension (see below). This can be overridden with the `-f` flag, passing a format matching a usual file extension (e.g. `-f java` for Java). If no legal format is specified `dec` will check for a C-style comment at the first line, otherwise will default to `#` comments.

To specify a file format manually:

`dec -c "/*,*/" -c "//" myfile` - C-style comments. When the second token is ommitted, a newline is assumed. The same syntax with `-q` applies to quote delimiters; no regex, and two tokens are required.

## Examples

To decomment a file and write to stdout:

`dec source.c`

To edit in-place

`dec -i *.java`

From stdin:

`cat myfile.c | dec -f c > out.c`

Some weird file format:

`dec -c "~~,~~" -q "','" myfile.foo`

# Known issues

- Mixed HTML and Javascript will not remove javascript comments.
- If the final line of the source file has no newline, one will be inserted.
- Trailing whitespace will be removed at the end of lines.

# ToDo

- An expanded list of source file types
- More robust testing

# Built-in formats

`dec` knows about commments and strings in:

| syntax     | File extension     |
| ---------- | ------------------ |
| C/C++      | .c, .h, .cpp, .hpp |
| Java       | .java              |
| Rust       | .rs                |
| Python     | .py                |
| (z\|ba)?sh | .sh                |
| Go         | .go                |
| Kotlin     | .kt                |
| JavaScript | .js                |
| TypeScript | .ts                |
| Perl       | .pl                |
| Ruby       | .rb                |
| SQL        | .sql               |
| PHP        | .php               |
| HTML       | .html              |
| CSS        | .css               |
