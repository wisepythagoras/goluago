package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
)

func main() {
	var filePath string
	var showVersion bool

	flag.StringVar(&filePath, "path", "", "The path to the file to run")
	flag.BoolVar(&showVersion, "version", false, "Output the version")

	flag.Parse()

	if showVersion {
		fmt.Printf("%s v%s\n", NAME, VERSION)
		fmt.Println("Usage:\n\tgoluago -path file/to/script.lua")
		os.Exit(0)
	} else if filePath == "" {
		fmt.Println("A file is required")
		os.Exit(1)
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileName := "main.lua"
	var dir fs.DirEntry
	var main fs.DirEntry

	if fileInfo.IsDir() {
		dir = fs.FileInfoToDirEntry(fileInfo)

		mainFile, err := os.Open(path.Join(filePath, fileName))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		mainFileInfo, _ := mainFile.Stat()
		main = fs.FileInfoToDirEntry(mainFileInfo)
	} else {
		mainFile, err := os.Open(filePath)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fileName = path.Base(filePath)
		filePath = path.Dir(filePath)

		baseDir, err := os.Open(path.Join(filePath, fileName))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		baseDirInfo, _ := baseDir.Stat()
		dir = fs.FileInfoToDirEntry(baseDirInfo)

		mainFileInfo, _ := mainFile.Stat()
		main = fs.FileInfoToDirEntry(mainFileInfo)
	}

	if dir == nil || main == nil {
		fmt.Println("Path is incorrect")
		os.Exit(1)
	}

	runtime := Runtime{
		Path: filePath,
		Dir:  dir,
		Main: main,
	}

	runtime.Init()
}
