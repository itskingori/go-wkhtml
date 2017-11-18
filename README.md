# go-wkhtml

[![GoDoc](https://godoc.org/github.com/itskingori/go-wkhtml?status.svg)](https://godoc.org/github.com/itskingori/go-wkhtml)
[![Build Status](https://travis-ci.org/itskingori/go-wkhtml.svg?branch=master)](https://travis-ci.org/itskingori/go-wkhtml)

Golang wrapper for `wkhtmltoimage` and `wkhtmltopdf`.

## Usage

### Image Example

```go
import (
	"fmt"
	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

// Attributes
format := "png"
width := 640
height := 480

// Method 1: Construct ImageFlagSet, then generate
ifs := make(wkhtmltox.ImageFlagSet)
ifs.SetFormat(format)
ifs.SetHeight(width)
ifs.SetWidth(height)
outputLogs, _ := ifs.Generate("http://duckduckgo.com", "/some/path/file.png")
fmt.Println(outputLogs)

// Method 2: Construct ImageOptions, extract ImageFlagSet, then generate
opts := wkhtmltox.ImageOptions{Format: &format, Width:  &width, Height: &height}
ifs = wkhtmltox.NewImageFlagSetFromOptions(&opts)
outputLogs, _ = ifs.Generate("http://duckduckgo.com", "/some/path/file.png")
fmt.Println(outputLogs)

// Method 3: Construct ImageOptions from JSON, extract ImageFlagSet, then generate
str := []byte(`{"format":"png","width": 640,"height": 480}`)
opts = wkhtmltox.ImageOptions{}
err := json.NewDecoder(str).Decode(&opts)
if err != nil {
	panic(err)
}
ifs = wkhtmltox.NewImageFlagSetFromOptions(&opts)
outputLogs, _ = ifs.Generate("http://duckduckgo.com", "/some/path/file.png")
fmt.Println(outputLogs)
```

### PDF Example

```go
import (
	"fmt"
	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

// Attributes
pageSize := "A4"

// Method 1: Construct PDFFlagSet, then generate
pfs := make(wkhtmltox.PDFFlagSet)
pfs.SetPageSize(pageSize)
outputLogs, _ := pfs.Generate("http://duckduckgo.com", "/some/path/file.pdf")
fmt.Println(outputLogs)

// Method 2: Construct PDFOptions, extract PDFFlagSet, then generate
opts := wkhtmltox.PDFOptions{PageSize: &pageSize}
pfs = wkhtmltox.NewPDFFlagSetFromOptions(&opts)
outputLogs, _ = pfs.Generate("http://duckduckgo.com", "/some/path/file.pdf")
fmt.Println(outputLogs)

// Method 3: Construct PDFOptions from JSON, extract PDFFlagSet, then generate
str := []byte(`{"page_size":"A4"}`)
opts = wkhtmltox.PDFOptions{}
err := json.NewDecoder(str).Decode(&opts)
if err != nil {
	panic(err)
}
pfs = wkhtmltox.NewPDFFlagSetFromOptions(&opts)
outputLogs, _ = pfs.Generate("http://duckduckgo.com", "/some/path/file.pdf")
fmt.Println(outputLogs)
```

## Development

### Testing

1. Install the Go testing tools via `make testing_dependencies`.
2. Run linter using `make lint` and test using `make test`.
