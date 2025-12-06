package main

import (
	"fmt"
	"os"
	"strings"
)

var application applicationType
var showComments bool = true // Default to showing comments/translations

func main() {

	// Get available languages from database first
	var availableLanguages []string
	var fallbackLanguages = []string{"ENG", "ESP", "DEU"}

	err := zoteroDB.Connect()
	if err == nil {
		availableLanguages, _ = zoteroDB.GetLanguages()
	}

	// If no languages found in DB, use fallback
	if len(availableLanguages) == 0 {
		availableLanguages = fallbackLanguages
	}

	if len(os.Args) < 2 {
		fmt.Println("No lang name sent")
		fmt.Printf("Please mention one from list: %s\n", strings.Join(availableLanguages, ", "))
		fmt.Println("Usage: voc <lang> [flag]")
		fmt.Println("  flag: no_trans (hide comments/translations)")
		os.Exit(1)
	}

	// Validate language argument
	requestedLang := strings.ToUpper(os.Args[1])
	isValidLang := false
	for _, lang := range availableLanguages {
		if strings.ToUpper(lang) == requestedLang {
			isValidLang = true
			break
		}
	}

	if !isValidLang {
		fmt.Printf("Invalid language: %s\n", os.Args[1])
		fmt.Printf("Available languages: %s\n", strings.Join(availableLanguages, ", "))
		os.Exit(1)
	}

	// Parse optional flag to hide comments/translations
	if len(os.Args) >= 3 {
		if os.Args[2] == "no_trans" {
			showComments = false
		} else {
			fmt.Println("Unknown flag:", os.Args[2])
			fmt.Println("Supported flag: no_trans (hide comments/translations)")
			os.Exit(1)
		}
	}

	application.init()
}

func check(err interface{}) {
	if err != nil {
		panic(err)
	}
}
