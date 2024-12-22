package utils

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/kasshii28/go-cli/static"
)

func  PrintAaFromText(fp string) {
	text := readFile(filepath.Join("aa", fp))
	fmt.Print(text)
}

func AaFromText(fp string) string {
	text := readFile(filepath.Join("aa", fp))
	return text
}

func readFile(fp string) string {
	file, err := aa.Aa.Open(fp)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}