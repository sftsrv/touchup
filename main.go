package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sftsrv/touchup/lib"
)

const usage = `touchup

A pipe for passing files to your $EDITOR without manually creating intermediate files

## Usage

You can use the '--help' flag to view usage information:

'''sh
touchup --help
'''

'touchup' can be used to interactively edit some text content via your configured $EDITOR

'''sh
cat my-file.txt | touchup
'''

It will take in the content of the input file and print out the result of editing the file

`

func main() {
	defaultEditor, err := lib.GetDefaultEditor()

	helpFlag := flag.Bool("help", false, "show usage info")
	editorFlag := flag.String("editor", defaultEditor, "editor to edit file paths with")
	prefixFlag := flag.String("prefix", "touchup_", "prefix to use for temp files")
	extFlag := flag.String("ext", "txt", "extension to use for temp files")

	flag.Parse()

	if *helpFlag {
		fmt.Print(usage)
		flag.Usage()
		return
	}

	if *editorFlag == "" && err != nil {
		panic(fmt.Errorf("Could not determine editor. Provide the --editor flag or ensure that the EDITOR environment variable is set"))
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	output, err := lib.EditFile(*editorFlag, *prefixFlag, *extFlag, string(input))
	if err != nil {
		panic(err)
	}

	os.Stdout.WriteString(output)
}
