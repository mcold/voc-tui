package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var application applicationType

func main() {

	if len(os.Args) != 2 {
		fmt.Println("No lang name sent")
		// TODO: change to query from DB
		fmt.Println("Please mention one from list: ENG, ESP, DEU")
		os.Exit(1)
	} else {
		application.init()
	}
}

func check(err interface{}) {
	if err != nil {
		_, fileName, lineNo, _ := runtime.Caller(1) // Получаем информацию о вызывающем файле
		log.Printf("%s: %d\n", filepath.Base(fileName), lineNo)
		log.Println(err)

		panic(err)
	}
}
