package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var application applicationType
var showComments bool = true // Default to showing comments/translations

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No lang name sent")
		// TODO: change to query from DB
		fmt.Println("Please mention one from list: ENG, ESP, DEU")
		fmt.Println("Usage: voc <lang> [flag]")
		fmt.Println("  flag: no_comments (hide comments/translations)")
		os.Exit(1)
	}

	// Parse optional flag to hide comments/translations
	if len(os.Args) >= 3 {
		if os.Args[2] == "no_comments" {
			showComments = false
		} else {
			fmt.Println("Unknown flag:", os.Args[2])
			fmt.Println("Supported flag: no_comments (hide comments/translations)")
			os.Exit(1)
		}
	}

	application.init()
}

func check(err interface{}) {
	if err != nil {
		_, fileName, lineNo, _ := runtime.Caller(1) // Получаем информацию о вызывающем файле
		log.Printf("%s: %d\n", filepath.Base(fileName), lineNo)
		log.Println(err)

		panic(err)
	}
}
