# VOC - Vocabulary Learning Tool

**VOC** (Vocabulary) is a terminal application (TUI) for learning vocabulary based on Zotero annotations.

## Description

The application extracts vocabulary items from annotations in your Zotero database and provides a convenient interface for studying them. Supports multiple languages and various display modes.


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
# Show words with translations (default)
./bin/main.exe ENG

# Hide translations
./bin/main.exe ENG no_trans
```

### Command Line Arguments

```
Usage: voc <lang> [flag]
  lang: Language for word extraction (ENG, ESP, DEU)
  flag: no_trans (hide comments/translations)
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

## Configuration

### Variables

- `showComments` - global variable for controlling comment display
- Automatically set via command line arguments

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

## Support

For questions and suggestions, create Issues in the project repository.

---

**Note**: Requires presence of annotations in Zotero PDF documents with corresponding language labels for proper operation.