# gls

golang implementation of the [ls](https://www.maths.cam.ac.uk/computing/linux/unixinfo/ls) command

## Installation
```console
$ go get github.com/Sudhindra3000/gls
```

## Usage
```console
$ gls -h
golang implementation of the ls command

Usage:
gls [flags]

Flags:
-a, --all       List all files including hidden files (files with names beginning with a dot)
-h, --help      help for gls
-r, --reverse   List the files in the Reverse of the order that they would otherwise have been listed in
```

## Lists files in sub-directories or directory with a given path

```console
$ gls
List files in current directory
$ gls subdir
List files in sub directory with given name
$ gls C:\Users\JohnDoe\Downloads
List files in directory with given path
```

## File Patterns
Accepts file patterns as arguments to list files based on patterns in names

### Supported Patterns
- `*pat*` (files with names containing given pattern)
- `*pat` (files with names ending with given pattern)
- `pat*` (files with names starting with given pattern)
