# VOC - Vocabulary Learning Tool

**VOC** (Vocabulary) is a terminal application (TUI) for learning vocabulary based on Zotero annotations.

## Description

The application extracts vocabulary items from annotations in your Zotero database and provides a convenient interface for studying them. Supports multiple languages and various display modes.

## Features

- 📚 **Zotero Integration**: Automatically extracts words from PDF document annotations
- 🌐 **Multilingual**: Supports ENG, ESP, DEU and other languages
- 💬 **Configurable Display**: Show/hide translations and comments
- ⌨️ **Convenient Navigation**: Hotkeys for quick operation
- 📋 **Quick Actions**: Copy links to clipboard, open in browser
- 🎨 **Beautiful TUI**: Modern interface based on tview/tcell

## Installation

### Requirements

- Go 1.24+
- Zotero with PDF annotations

### Build

```bash
git clone <repository-url>
cd voc
go build -o bin/main.exe ./voc/
```

### Dependencies

The project uses the following libraries:
- `modernc.org/sqlite` - Zotero database operations
- `github.com/gdamore/tcell/v2` + `github.com/rivo/tview` - TUI interface
- `github.com/atotto/clipboard` - clipboard operations
- `github.com/go-vgo/robotgo` - automation

## Usage

### Basic Launch

```bash
# Show words with comments (default)
./bin/main.exe ENG

# Hide comments
./bin/main.exe ENG no_comments
```

### Supported Languages

- `ENG` - English
- `ESP` - Spanish  
- `DEU` - German

### Command Line Arguments

```
Usage: voc <lang> [flag]
  lang: Language for word extraction (ENG, ESP, DEU)
  flag: no_comments (hide comments/translations)
```

## Hotkeys

The following key combinations are available in the application interface:

- `Alt + ` ` (Backtick)` - Open current PDF in Zotero
- `Alt + 1` - Copy annotation link to clipboard
- `Alt + h` - Toggle comment display
- `↑/↓` - Navigate through word list
- `Ctrl+C` - Exit application

## Project Structure

```
voc/
├── voc/                    # Main application code
│   ├── main.go            # Entry point, argument processing
│   ├── app.go             # Main application logic
│   ├── page_voc.go        # Vocabulary learning interface
│   ├── page_main.go       # Main page
│   ├── page_confirm.go    # Exit confirmation dialog
│   └── zoteroDB.go        # Zotero database operations
├── go.mod                 # Go module and dependencies
├── go.sum                 # Dependency checksums
├── .gitignore            # Ignored files
└── README.md             # Documentation
```

## Database

The application connects to the Zotero database at:
```
<UserHome>/Zotero/zotero.sqlite
```

Extracts data from tables:
- `items` - basic item information
- `itemData` + `itemDataValues` - metadata (languages)
- `itemAttachments` - attached files
- `itemAnnotations` - annotations with text and comments

## Workflow Logic

1. **Database Connection**: Opens Zotero database in user's home folder
2. **Data Extraction**: Executes SQL query to get annotations by language
3. **Text Processing**: Splits comments into lines, extracts translations
4. **Grouping**: Groups words by first letter for quick navigation
5. **Display**: Shows words in TUI with mode switching capability

## Configuration

### Variables

- `showComments` - global variable for controlling comment display
- Automatically set via command line arguments

### Logging

The application creates an `app.log` file in the working directory for debugging.

## Development

### Adding New Languages

1. Add new language code to argument validation in `main.go`
2. Ensure Zotero has documents with corresponding annotations

### Interface Customization

Main interface components are located in:
- `page_voc.go` - word list and translation panel
- `app.go` - general application structure

### Database Operations

SQL queries and Zotero integration logic are located in `zoteroDB.go` and `page_voc.go`.

## License

[Specify project license]

## Support

For questions and suggestions, create Issues in the project repository.

---

**Note**: Requires presence of annotations in Zotero PDF documents with corresponding language labels for proper operation.