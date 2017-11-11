# go-wkhtml

[![GoDoc](https://godoc.org/github.com/itskingori/go-wkhtml?status.svg)](https://godoc.org/github.com/itskingori/go-wkhtml)
[![Build Status](https://travis-ci.org/itskingori/go-wkhtml.svg?branch=master)](https://travis-ci.org/itskingori/go-wkhtml)

Golang wrapper for `wkhtmltoimage` and `wkhtmltopdf`.

## Usage

### Image

To render an image: pass in the options, generate a flag-set then call generate
on the flag-set ...

```go
import (
	"fmt"
	"path/filepath"
	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

format := "png"
width := 640
height := 480
opts := wkhtmltox.ImageOptions{
	Format: &format,
	Width:  &width,
	Height: &height,
}
ifs := opts.FlagSet()

inputURL := "http://duckduckgo.com"
outputFile = filepath.Join("/some/path", "file.png")
outputLogs, _ := ifs.Generate(inputURL, outputFile)

fmt.Println(outputLogs)
```

### PDF

To render a pdf: pass in the options, generate a flag-set then call generate
on the flag-set ...


```go
import (
	"fmt"
	"path/filepath"
	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

pageSize := "A4"
opts := wkhtmltox.PDFOptions{
	PageSize: &pageSize,
}
pfs := opts.FlagSet()

inputURL := "http://duckduckgo.com"
outputFile = filepath.Join("/some/path", "file.pdf")
outputLogs, _ := pfs.Generate(inputURL, outputFile)

fmt.Println(outputLogs)
```

## Development

### Testing

1. Install the Go testing tools via `make testing_dependencies`.
2. Run linter using `make lint` and test using `make test`.
