# README

## Overview

This Go-based command-line utility allows users to recursively navigate through a specified directory and display the content of each file. The utility offers options for ignoring specific folders, files, and file types during traversal.

## Usage

```sh
go run main.go [folderPath] [--ignore-folder FOLDER1,FOLDER2,...] [--ignore-file FILE1,FILE2,...] [--ignore-filetype .EXT1,.EXT2,...]
```

### Arguments
- `folderPath`: The path of the directory to traverse (required).
- `--ignore-folder`: An optional argument, followed by a comma-separated list of folder names to exclude during directory tree traversal.
- `--ignore-file`: An optional argument, followed by a comma-separated list of file names to exclude during directory tree traversal.
- `--ignore-filetype`: An optional argument, followed by a comma-separated list of file extensions to exclude during directory tree traversal.

### Examples

#### Example 1

Traverse the directory at `/path/to/folder` and display the content of each file without any filtering:
```sh
go run main.go /path/to/folder
```

#### Example 2

Traverse the directory at `/path/to/folder` and display the content of each file, excluding those within the `node_modules` folder:
```sh
go run main.go /path/to/folder --ignore-folder node_modules
```

#### Example 3

Traverse the directory at `/path/to/folder` and display the content of each file, excluding those named `.gitignore` and `.editorconfig`:
```sh
go run main.go /path/to/folder --ignore-file .gitignore,.editorconfig
```

#### Example 4

Traverse the directory at `/path/to/folder` and display the content of each file, excluding those with `.json` and `.txt` extensions:
```sh
go run main.go /path/to/folder --ignore-filetype .json,.txt
```

#### Example 5

Traverse the directory at `/path/to/folder` and display the content of each file, excluding files within `test_data` and `logs` folders, files named `README.md` and `LICENSE`, and files with `.json` and `.xml` extensions:
```sh
go run main.go /path/to/folder --ignore-folder test_data,logs --ignore-file README.md,LICENSE --ignore-filetype .json,.xml
```

## Dependencies

The program relies on the following standard Go packages:

- fmt
- io/ioutil
- os
- path/filepath
- strings