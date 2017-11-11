// Copyright © 2017 Job King'ori Maina <j@kingori.co>
//
// This file is part of go-wkhtml.
//
// go-wkhtml is free software: you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// go-wkhtml is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with go-wkhtml. If not, see <http://www.gnu.org/licenses/>.

package wkhtmltox

import (
	"fmt"

	"encoding/json"
	"os/exec"
)

const (
	pdfConverterBinary = "wkhtmltopdf"
)

// PDFFlagSet are key-value pair of PDF converter flags
type PDFFlagSet FlagSet

// PDFOptions represents a wkhtmlpdf attributes
type PDFOptions struct {
	// Global
	MarginTop    *int    `json:"margin_top,omitempty"`    // Set the page top margin
	MarginBottom *int    `json:"margin_bottom,omitempty"` // Set the page bottom margin
	MarginLeft   *int    `json:"margin_left,omitempty"`   // Set the page left margin
	MarginRight  *int    `json:"margin_right,omitempty"`  // Set the page right margin
	PageHeight   *int    `json:"page_height,omitempty"`   // Height of the page
	PageWidth    *int    `json:"page_width,omitempty"`    // Width of the page
	PageSize     *string `json:"page_size,omitempty"`     // Size of the page
}

// NewPDFFlagSetFromJSON create pdf options from JSON
func NewPDFFlagSetFromJSON(data []byte) (PDFFlagSet, error) {
	opts := &PDFOptions{}

	err := json.Unmarshal(data, &opts)
	pfs := opts.FlagSet()
	if err != nil {
		return pfs, err
	}

	return pfs, nil
}

// FlagSet generates a FlagSet from PDFOptions
func (opts *PDFOptions) FlagSet() PDFFlagSet {
	pfs := make(PDFFlagSet)

	if opts.MarginTop != nil {
		pfs["margin-top"] = *opts.MarginTop
	}

	if opts.MarginBottom != nil {
		pfs["margin-bottom"] = *opts.MarginBottom
	}

	if opts.MarginLeft != nil {
		pfs["margin-left"] = *opts.MarginLeft
	}

	if opts.MarginRight != nil {
		pfs["margin-right"] = *opts.MarginRight
	}

	if opts.PageHeight != nil {
		pfs["page-height"] = *opts.PageHeight
	}

	if opts.PageWidth != nil {
		pfs["page-width"] = *opts.PageWidth
	}

	if opts.PageSize != nil {
		pfs["page-size"] = *opts.PageSize
	}

	return pfs
}

// Flags converts a FlagSet to a String slice
func (pfs *PDFFlagSet) Flags() []string {
	var flags []string

	for flagKey, flagValue := range *pfs {
		flags = append(flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%v", flagValue))
	}

	return flags
}

// GetMarginTop retrieves the MarginTop from PDFFlagSet
func (pfs *PDFFlagSet) GetMarginTop() (int, bool) {
	value, exists := (*pfs)["margin-top"]

	return value.(int), exists
}

// GetMarginBottom retrieves the MarginBottom from PDFFlagSet
func (pfs *PDFFlagSet) GetMarginBottom() (int, bool) {
	value, exists := (*pfs)["margin-bottom"]

	return value.(int), exists
}

// GetMarginLeft retrieves the MarginLeft from PDFFlagSet
func (pfs *PDFFlagSet) GetMarginLeft() (int, bool) {
	value, exists := (*pfs)["margin-left"]

	return value.(int), exists
}

// GetMarginRight retrieves the MarginRight from PDFFlagSet
func (pfs *PDFFlagSet) GetMarginRight() (int, bool) {
	value, exists := (*pfs)["margin-right"]

	return value.(int), exists
}

// GetPageHeight retrieves the PageHeight from PDFFlagSet
func (pfs *PDFFlagSet) GetPageHeight() (int, bool) {
	value, exists := (*pfs)["page-height"]

	return value.(int), exists
}

// GetPageWidth retrieves the PageWidth from PDFFlagSet
func (pfs *PDFFlagSet) GetPageWidth() (int, bool) {
	value, exists := (*pfs)["page-width"]

	return value.(int), exists
}

// GetPageSize retrieves the PageSize from PDFFlagSet
func (pfs *PDFFlagSet) GetPageSize() (string, bool) {
	value, exists := (*pfs)["page-size"]

	return value.(string), exists
}

// Generate creates an pdf based on the parameters passed in
func (pfs *PDFFlagSet) Generate(inputURL string, outputFile string) ([]byte, error) {
	var out []byte

	params := pfs.Flags()
	params = append(params, inputURL, outputFile)
	cmd := exec.Command(pdfConverterBinary, params...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, err
}
