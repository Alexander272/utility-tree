package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for i, file := range files {
		if printFiles && !file.IsDir() {
			info, _ := file.Info()
			size := info.Size()
			var sizeStr string
			if size == 0 {
				sizeStr = "empty"
			} else {
				sizeStr = fmt.Sprintf("%db", size)
			}
			if i == len(files)-1 {
				fmt.Fprintf(out, "└───%s (%s)\n", file.Name(), sizeStr)
			} else {
				fmt.Fprintf(out, "├───%s (%s)\n", file.Name(), sizeStr)
			}
		}
		if file.IsDir() {
			if i == len(files)-1 {
				fmt.Fprintf(out, "└───%s\n", file.Name())
			} else {
				fmt.Fprintf(out, "├───%s\n", file.Name())
			}
			if i == len(files) {
				printFile(out, filepath.Join(path, file.Name()), printFiles, "\t")
			} else {
				printFile(out, filepath.Join(path, file.Name()), printFiles, "│\t")
			}
		}

	}

	return nil
}

func printFile(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for i, file := range files {
		if printFiles && !file.IsDir() {
			info, _ := file.Info()
			size := info.Size()
			var sizeStr string
			if size == 0 {
				sizeStr = "empty"
			} else {
				sizeStr = fmt.Sprintf("%db", size)
			}
			if i == len(files)-1 {
				fmt.Fprintf(out, "%s└───%s (%s)\n", prefix, file.Name(), sizeStr)
			} else {
				fmt.Fprintf(out, "%s├───%s (%s)\n", prefix, file.Name(), sizeStr)
			}
		}
		if file.IsDir() {
			if i == len(files)-1 {
				fmt.Fprintf(out, "%s└───%s\n", prefix, file.Name())
			} else {
				fmt.Fprintf(out, "%s├───%s\n", prefix, file.Name())
			}
			if i == len(files)-1 {
				printFile(out, filepath.Join(path, file.Name()), printFiles, prefix+"\t")
			} else {
				printFile(out, filepath.Join(path, file.Name()), printFiles, prefix+"│\t")
			}
		}
	}

	return nil
}
