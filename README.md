# touchup

A little pipe for editing files with your `$EDITOR`

## Installation

Install the application using `go install` as follows:

```sh
go install github.com/sftsrv/touchup
```

## Usage

You can use the `--help` flag to view usage information:

```sh
touchup --help
```

`touchup` can be used to interactively edit some text content via your configured `$EDITOR`

```sh
cat my-file.txt | touchup
```

It will take in the content of the input file and print out the result of editing the file

## Flags

The application also supports the following flags to improve the editing experience:

- `--editor ` - Editor to edit file paths with (default `$EDITOR`)
- `--ext ` - Extension to use for temp files (default `txt`)
- `--prefix ` - Prefix to use for temp files (default `touchup_`)
- `--help` - Show usage info
