package main

import (
	"flag"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

const sniffLen = 512

func detectContentExt(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	var buf [sniffLen]byte
	n, _ := io.ReadFull(file, buf[:])
	kind, _ := filetype.Match(buf[:n])
	if kind == filetype.Unknown {
		return ""
	}

	return kind.Extension
}

func main() {
	var cwd = ""
	var defaultExt = ""

	flag.StringVar(&defaultExt, "default", "", "Default ext when failed to detect automatically, default empty string")
	flag.String("", "", "Working Directory, default is current working directory")

	flag.Parse()

	workingDir := flag.Arg(0)
	if workingDir == "" {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		cwd = workingDir
	}

	files, err := os.ReadDir(cwd)
	if err != nil {
		panic(err)
	}

	noneExtFiles := make([]string, 0, len(files))
	for _, file := range files {
		name := file.Name()
		if file.IsDir() || strings.Contains(name, ".") {
			continue
		}
		noneExtFiles = append(noneExtFiles, name)
	}

	if len(noneExtFiles) == 0 {
		log.Printf("No none-ext file found in %s", cwd)
		return
	}

	tab := table.NewWriter()
	tab.SetOutputMirror(os.Stdout)
	tab.AppendHeader(table.Row{"#", "File", "Ext"})
	for index, file := range noneExtFiles {
		name := path.Join(cwd, file)
		ext := detectContentExt(name)
		if ext == "" {
			ext = defaultExt
		}
		tab.AppendRow([]any{index + 1, file, ext})
		if ext == "" {
			continue
		}
		err := os.Rename(name, fmt.Sprintf("%s.%s", name, ext))
		if err != nil {
			panic(err)
		}
	}
	tab.Render()
}
