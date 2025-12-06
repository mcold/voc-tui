package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/go-vgo/robotgo"
	"github.com/rivo/tview"
)

type itemType struct {
	itemKey   string
	pageNum   int
	attachKey string
}

type pageVocType struct {
	lVoc             *tview.List
	vTrans           *tview.TextView
	vLink            *tview.TextView
	fTrans           *tview.Flex
	mPosTrans        map[int]string
	mPosItems        map[int]itemType
	vocItems         []vocItem
	*tview.Flex
}

type vocItem struct {
	primaryText string
	comment     string
	shortcut    rune
}

var pageVoc pageVocType

func (pageVoc *pageVocType) build() {

	pageVoc.mPosTrans = make(map[int]string)
	pageVoc.mPosItems = make(map[int]itemType)
	pageVoc.vocItems = make([]vocItem, 0)

	pageVoc.lVoc = tview.NewList()
	pageVoc.lVoc.SetBorderPadding(2, 2, 2, 2).
		SetBorderColor(tcell.ColorBlack)

	pageVoc.lVoc.SetSelectedFunc(func(pos int, _ string, _ string, _ rune) {
		pageVoc.vTrans.SetText(pageVoc.mPosTrans[pos])
		pageVoc.vLink.SetText(fmt.Sprintf("zotero://open-pdf/library/items/%s?page=%d&annotation=%s", pageVoc.mPosItems[pos].attachKey, pageVoc.mPosItems[pos].pageNum, pageVoc.mPosItems[pos].itemKey))
	})

	pageVoc.vTrans = tview.NewTextView()
	pageVoc.vTrans.SetBorderPadding(1, 1, 1, 1).
		SetBorderColor(tcell.ColorBlack)

	pageVoc.vLink = tview.NewTextView()
	pageVoc.vLink.SetBorderPadding(1, 1, 1, 1).
		SetBorderColor(tcell.ColorBlack)

	pageVoc.fTrans = tview.NewFlex()
	pageVoc.fTrans.SetBorderPadding(1, 1, 1, 1)

	pageVoc.fTrans.SetDirection(tview.FlexRow)

	pageVoc.fTrans.AddItem(pageVoc.vTrans, 0, 8, true).
		//AddItem(pageVoc.vLink, 0, 1, false).
		SetBorderColor(tcell.ColorBlack)

	pageVoc.Flex = tview.NewFlex()
	pageVoc.Flex.SetBorderPadding(2, 2, 2, 2)

	pageVoc.Flex.AddItem(pageVoc.lVoc, 0, 3, true).
		AddItem(pageVoc.fTrans, 0, 8, false).
		SetBorderColor(tcell.ColorBlack)

	pageVoc.Flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '`' && event.Modifiers() == tcell.ModAlt {
			err := OpenLinkInBrowser(pageVoc.vLink.GetText(true))
			if err != nil {
			}
		}
		if event.Rune() == '1' && event.Modifiers() == tcell.ModAlt {
			err := clipboard.WriteAll(pageVoc.vLink.GetText(true))
			check(err)
		}
		if event.Rune() == 'h' && event.Modifiers() == tcell.ModAlt {
			pageVoc.toggleTranslations()
		}
		if event.Key() == tcell.KeyDown || event.Key() == tcell.KeyUp {
			err := robotgo.KeyTap("Enter")
			check(err)
		}

		return event

	})

	pageMain.Pages.AddPage("voc", pageVoc.Flex, true, true)
	
	// Initialize the list with current comment display state
	pageVoc.refreshVocList(showComments)
}

func (pageVoc *pageVocType) show() {
	pageMain.Pages.SwitchToPage("voc")
	pageVoc.lVoc.Clear()
	pageVoc.vocItems = make([]vocItem, 0) // Clear existing items
	setVoc()
	pageVoc.refreshVocList(showComments) // Apply current comment display state
	app.SetFocus(pageMain.Pages)
}

func setVoc() {
	// Clear existing voc items
	pageVoc.vocItems = make([]vocItem, 0)

	query := `select iian.itemID as annID
				, iian.KEY as annItemKEY
				, ipar."key" as attachKey
				, json_extract(ian.position, '$.pageIndex')+1 pageNum
				, ian."text" as text
				, ian.comment as trans
		  from items i
		  join itemData idat on idat.itemID = i.itemID and idat.fieldID = 7
		  join itemDataValues ival on ival.valueID = idat.valueID
		  join itemAttachments ia on ia.parentItemID = i.itemID
		  join itemAnnotations ian on ian.parentItemID = ia.itemID
		  join items iian on iian.itemID = ian.itemID
		  join items ipar on ipar.itemID = ian.parentItemID
		  where ival.value is not null
		    and ian.type = 5
			and ival.value = '` + os.Args[1] + `'
		  order by lower(ian."text")`

	log.Println(query)

	words, err := zoteroDB.Query(query)
	check(err)

	posNum := 0
	var lastFirstLetter rune = 0

	var id, pageNum sql.NullInt64
	var text, comment, attachKey, itemKey sql.NullString

	for words.Next() {
		err := words.Scan(&id, &itemKey, &attachKey, &pageNum, &text, &comment)
		check(err)

		var currentFirstLetter rune = 0
		if len(strings.ToLower(text.String)) > 0 {
			currentFirstLetter = unicode.ToLower(rune(strings.ToLower(text.String)[0]))
		}

		var displayRune rune = 0
		if currentFirstLetter != lastFirstLetter {
			displayRune = currentFirstLetter
			lastFirstLetter = currentFirstLetter
		}

		// Store voc item for comment toggle functionality
		pageVoc.vocItems = append(pageVoc.vocItems, vocItem{
			primaryText: strings.ToLower(text.String),
			comment:     comment.String,
			shortcut:    displayRune,
		})

		pageVoc.mPosItems[posNum] = itemType{itemKey: itemKey.String, pageNum: int(pageNum.Int64), attachKey: attachKey.String}
		pageVoc.mPosTrans[posNum] = comment.String

		posNum++
	}

	err = words.Close()
	check(err)
}

func OpenLinkInBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "windows": // Windows
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux": // Linux
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("Unsupported platform")
	}

	return cmd.Run()
}

// toggleTranslations hides/shows all translations
func (pageVoc *pageVocType) toggleTranslations() {
	// Toggle the global state
	showComments = !showComments

	// Refresh the list with new state
	pageVoc.refreshVocList(showComments)
}

// refreshVocList rebuilds the vocabulary list with current comment display state
func (pageVoc *pageVocType) refreshVocList(showComments bool) {
	pageVoc.lVoc.Clear()

	for i, voc := range pageVoc.vocItems {
		secondaryText := ""
		if showComments {
			secondaryText = voc.comment
		}

		pageVoc.lVoc.AddItem(voc.primaryText, secondaryText, voc.shortcut, func() {})
		pageVoc.mPosTrans[i] = voc.comment
		// mPosItems remains the same
	}
}
