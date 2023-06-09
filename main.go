package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printHelp() {
	fmt.Println("Usage: go run main.go <folderPath> [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --ignore-folder <folder1,folder2,...>    Ignore specified folders (comma-separated)")
	fmt.Println("  --ignore-file <file1,file2,...>          Ignore specified files (comma-separated)")
	fmt.Println("  --ignore-filetype <ext1,ext2,...>        Ignore specified file extensions (comma-separated)")
	fmt.Println("  --help                                   Show this help message")
}

func main() {
	// Get command line arguments
	var folderPath string
	var ignoreFolders []string
	var ignoreFiles []string
	var ignoreFileTypes []string

	// Check if arguments exist
	if len(os.Args) > 1 {
		// Check for help argument
		for _, arg := range os.Args {
			if arg == "--help" {
				printHelp()
				return
			}
		}

		folderPath = os.Args[1]

		// Check for optional arguments
		for i := 2; i < len(os.Args); i++ {
			arg := os.Args[i]

			switch arg {
			case "--ignore-folder":
				if i+1 < len(os.Args) {
					ignoreFolders = strings.Split(os.Args[i+1], ",")
				}
			case "--ignore-file":
				if i+1 < len(os.Args) {
					ignoreFiles = strings.Split(os.Args[i+1], ",")
				}
			case "--ignore-filetype":
				if i+1 < len(os.Args) {
					ignoreFileTypes = strings.Split(os.Args[i+1], ",")
				}
			}
		}
	} else {
		printHelp()
		return
	}

	// Walk through directory tree
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		// Check for errors
		if err != nil {
			return err
		}

		// always skip .git folder
		if strings.Contains(path, ".git") && info.IsDir() {
			return filepath.SkipDir
		}

		// Check if folder should be ignored
		for _, folder := range ignoreFolders {
			if info.IsDir() && info.Name() == folder {
				return filepath.SkipDir
			}
		}

		// Check if file should be ignored
		for _, file := range ignoreFiles {
			if !info.IsDir() && info.Name() == file {
				return nil
			}
		}

		// Check if file extension should be ignored
		for _, extension := range ignoreFileTypes {
			if !info.IsDir() && filepath.Ext(info.Name()) == extension {
				return nil
			}
		}

		// Display file contents
		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			} else {
				fileExt := filepath.Ext(info.Name())
				if len(fileExt) > 0 {
					fileExt = fileExt[1:]
				} else {
					fileExt = "txt"
				}
				fmt.Printf("```%s\n%s\n```\n\n", fileExt, string(data))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
