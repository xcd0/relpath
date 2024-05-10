package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func main() {
	//log.Printf("len: %d, args: %v", len(os.Args), os.Args)
	switch len(os.Args) {
	case 2:
		// カレントディレクトリからの相対パスを出力する。
		c, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p, err := filepath.Rel(c, AbsPath(os.Args[1]))
		if err != nil {
			panic(err)
		}
		fmt.Println(filepath.ToSlash(p))
	case 3:
		// $1から$2の相対パスを出力する。
		p, err := filepath.Rel(AbsPath(os.Args[1]), AbsPath(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(filepath.ToSlash(p))
	default:
		program_name := filepath.Base(os.Args[0])
		fmt.Printf(`Usage: %v <path-to>
       %v <path-base> <path-to>

Description:
    '%v' calculates and outputs the relative path between specified paths.

Arguments:
    <path-to>      The destination path for which the relative path is calculated.
                   This path is output as relative from the current directory or from <path-base>.
    <path-base>    The base directory path (optional).
                   If specified, the relative path from this path to <path-to> is calculated.

Examples:
    1. Output the relative path from the current directory to a specified path:
       $ %v /home/user/docs

    2. Output the relative path from one directory to another:
       $ %v /var/log /var/log/syslog

Note:
    If the path does not exist or cannot be read, an error message will be displayed.

`, program_name, program_name, program_name, program_name, program_name)
	}
}

func AbsPath(path string) string {
	absoluteOutputPath, err := filepath.Abs(path)
	if err != nil {
		panic(errors.Errorf("failed to resolve absolute path for output: %w", err))
	}
	absoluteOutputPath = filepath.ToSlash(absoluteOutputPath)
	return absoluteOutputPath
}
