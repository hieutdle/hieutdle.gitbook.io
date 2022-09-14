# 📔 Learning Go: An Idiomatic Approach to Real-World Go Programming

**[You can find all the code for the notes of this book here](https://github.com/ledinhtrunghieu/learning-go)**

## Primitive Types and Declarations

### Literals

**Integer literals** are sequences of numbers; they are normally base ten, but different
prefixes are used to indicate other bases: `0b` for binary (base two), `0o` for octal (base
eight), or `0x` for hexadecimal (base sixteen). You can use either or upper- or lowercase
letters for the prefix. A leading 0 with no letter after it is another way to represent an
**octal literal**. Do not use it, as it is very confusing.

Rune literals represent characters and are surrounded by single quotes. Unlike many
other languages, in Go single quotes and double quotes are **not interchangeable**. Rune
literals can be written as single Unicode characters `('a')`, 8-bit octal numbers
`('\141')`, 8-bit hexadecimal numbers `('\x61')`, 16-bit hexadecimal numbers
`('\u0061')`, or 32-bit Unicode numbers `('\U00000061')`. There are also several back‐
slash escaped rune literals, with the most useful ones being newline `('\n')`, tab
`('\t')`, single quote `('\'')`, double quote `('\"')`, and backslash `('\\').`



