package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTreeWalk(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	files, err := ioutil.ReadDir(path + "/")

	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	return nil

}

func dirTreeWalk(out io.Writer, mypath string, printFiles bool) error {

	err := filepath.Walk(mypath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.Name() == "." {
			return nil
		}

		var count int = len(strings.Split(path, "\\"))

		if info.IsDir() {
			fmt.Printf("%+v \n", printLevel(count, "")+info.Name())
		} else {
			fmt.Printf("%+v \n", printLevel(count, "")+info.Name()+" ("+fmt.Sprintf("%v", info.Size())+")")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", err)
		return nil
	}

	return nil
}

// Если уровень = deep - 1 - тогда ---, иначе "   " уровень - 1 и повторить
func printLevel(deep int, pathPic string) (result string) {

	if deep-1 == 0 {
		result = pathPic + "├───"
	} else {
		result = printLevel(deep-1, pathPic+"│   ")
	}

	return result
}
