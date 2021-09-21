package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Files struct {
	level  int
	parent string
	files  []string
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	flag := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	//     panic(err.Error())
	// }
	//     F.Dirs = make(map[string][]string)
	log.Print(flag, "FFF")
	dirTree(out, path, flag, "")
}

func dirTree(out *os.File, path string, flag bool, prefix string) {

	// todo: flag -> show with files, else show only folder
	// todo2:  show size file
	// last file set └───

	//if level && folderName == , parent Set, child - append - files ^=& folders
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	var directories []os.FileInfo

	//find all directories
	for _, file := range fileInfo {
		if !flag && !file.IsDir() {
			continue
		}
		directories = append(directories, file)
	}

	//if flag ->  get files, else only folder
	for i, file := range fileInfo {
		// newLevel := Files{level:level, parent: path}
		currentPrefix := prefix

		//accum previus prefix , like accum, level+= tab +\t
		//if now folder - add "   "else add pipeline
		if i != len(directories)-1 {
			currentPrefix = currentPrefix + "│   "
		} else {
			currentPrefix = currentPrefix + "    "
		}

		if file.IsDir() {
			_, err = fmt.Fprintln(out, prefix+file.Name())
			// fs = append(fs, newLevel)
			dirTree(out, path+string(os.PathSeparator)+file.Name(), flag, currentPrefix)
		} else {
			if flag {
				//if file index == len folders, add "└───", else
				if i == len(directories)-1 {
					_, err = fmt.Fprintln(out, prefix+"└───"+file.Name())
				} else {
					_, err = fmt.Fprintln(out, prefix+"├───"+file.Name())
				}
			}
			// newLevel.files = append(newLevel.files, file.Name())
			// fs = append(fs, newLevel)
		}
	}
}
