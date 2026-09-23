package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func GetDefaultEditor() (string, error) {
	editor, exists := os.LookupEnv("EDITOR")

	if !exists {
		return editor, fmt.Errorf("No EDITOR found in environment")
	}

	return editor, nil
}

func EditFile(editor string, prefix string, ext string, content string) (string, error) {
	file, err := os.CreateTemp(os.TempDir(), prefix+"_*."+ext)
	if err != nil {
		return content, err
	}

	path := file.Name()

	_, err = file.WriteString(content)
	if err != nil {
		return content, err
	}

	err = file.Close()
	if err != nil {
		return content, err
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return content, err
	}

	updated, err := os.ReadFile(path)
	if err != nil {
		return content, err
	}

	return string(updated), nil
}
