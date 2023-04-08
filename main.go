package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get the command line arguments
	var ignoreFlag string
	flag.StringVar(&ignoreFlag, "ignore", "", "comma-separated list of files or folders to ignore")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run script.go [-ignore folder1,file2,...] folder_to_search")
		os.Exit(1)
	}

	// The first argument is the folder to look at
	folder := flag.Arg(0)

	// The remaining arguments are files or folders to ignore
	var ignores []string
	if ignoreFlag != "" {
		ignores = strings.Split(strings.TrimSpace(ignoreFlag), ",")
	}
	fmt.Println("ignoreFlag:", ignoreFlag)
	fmt.Println("ignores:", ignores)

	// Walk through all files in the folder, including subfolders
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Ignore directories and files in the ignore list
		if info.IsDir() || shouldIgnore(path, ignores) {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()

		contents, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Print the file's contents with markdown formatting and file extension
		ext := strings.TrimLeft(filepath.Ext(path), ".")
		fmt.Printf("## `%s`\n\n```%s\n%s\n```\n\n", info.Name(), ext, string(contents))

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Helper function to check if a file or directory should be ignored
func shouldIgnore(path string, ignores []string) bool {
	fileName := filepath.Base(path)
	directory := filepath.Dir(path)

	// path looks like this : venv/no.txt
	// we need to split it and check if any of the items has a match in ignores
	for _, ignore := range ignores {
		if fileName == ignore || directory == ignore {
			return true
		}
	}
	return false
}
