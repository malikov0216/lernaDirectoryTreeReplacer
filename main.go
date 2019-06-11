package main

import (
	"fmt"
	"github.com/malikov0216/lerna_Replacing/fileMethods"
	"log"
	"os"
)

func main () {
	source := os.Args[1]
	destination := os.Args[2]
	fmt.Printf(source)


	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <directory> %s source %s destination\n", os.Args[0], os.Args[1], os.Args[2])
	}


	hasFolder, err := fileMethods.Exists(destination)
	if err != nil {
		log.Fatalln(err)
	}
	if hasFolder {
		err := os.MkdirAll(destination, 0777)
		if err != nil {
			log.Fatalln(err)
		}
	}


	fileMethods.CopyDirectory(source, destination)
}

