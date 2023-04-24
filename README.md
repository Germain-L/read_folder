# README

## Overview

This program is a simple command-line utility written in Go that recursively walks through a specified directory and displays the content of each file. The program also provides options to ignore specific folders, files, and file types.

## Usage

```sh
go run main.go [folderPath] [--ignore-folder FOLDER1,FOLDER2,...] [--ignore-file FILE1,FILE2,...] [--ignore-filetype .EXT1,.EXT2,...]
```


### Arguments
 - folderPath: The path to the directory you want to traverse. This is a required argument.
 - --ignore-folder: An optional argument followed by a comma-separated list of folder names to ignore while traversing the directory tree.
 - --ignore-file: An optional argument followed by a comma-separated list of file names to ignore while traversing the directory tree.
 - --ignore-filetype: An optional argument followed by a comma-separated list of file extensions to ignore while traversing the directory tree.

### Examples

#### Example 1

This command will traverse the directory at /path/to/folder and display the content of each file without any filtering.
```sh
go run main.go /path/to/folder
```

#### Example 2

This command will traverse the directory at /path/to/folder and display the content of each file, excluding files in the node_modules folder.
```sh
go run main.go /path/to/folder --ignore-folder node_modules
```

#### Example 3

This command will traverse the directory at /path/to/folder and display the content of each file, excluding files named .gitignore and .editorconfig.
```sh
go run main.go /path/to/folder --ignore-file .gitignore,.editorconfig
```

#### Example 4

This command will traverse the directory at /path/to/folder and display the content of each file, excluding files with the .json and .txt extensions.

```sh
go run main.go /path/to/folder --ignore-filetype .json,.txt
```

#### Example 5

This command will traverse the directory at /path/to/folder and display the content of each file, excluding files in the test_data and logs folders, files named README.md and LICENSE, and files with the .json and .xml extensions.

```sh
go run main.go /path/to/folder --ignore-folder test_data,logs --ignore-file README.md,LICENSE --ignore-filetype .json,.xml
```

## Dependencies

The program uses only standard Go packages:

 - fmt
 - io/ioutil
 - os
 - path/filepath
 - strings
