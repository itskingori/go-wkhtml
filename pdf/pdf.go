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

package pdf

import (
	"encoding/json"
)

// FlagSets are key-value pair of flags
type FlagSets map[string]interface{}

// Options represents a wkhtmlpdf attributes
type Options struct {
	GlobalOptions
}

// GlobalOptions represents wkhtmltopdf global options
type GlobalOptions struct {
	MarginTop    *int `json:"margin_top,omitempty"`    // Set the page top margin
	MarginBottom *int `json:"margin_bottom,omitempty"` // Set the page bottom margin
	MarginLeft   *int `json:"margin_left,omitempty"`   // Set the page left margin
	MarginRight  *int `json:"margin_right,omitempty"`  // Set the page right margin

	PageHeight *int    `json:"page_height,omitempty"` // Height of the page
	PageWidth  *int    `json:"page_width,omitempty"`  // Width of the page
	PageSize   *string `json:"page_size,omitempty"`   // Size of the page
}

// NewOptionsFromJSON create pdf options from JSON
func NewOptionsFromJSON(data []byte) (*Options, error) {
	opts := &Options{}

	err := json.Unmarshal(data, &opts)
	if err != nil {
		return opts, err
	}

	return opts, nil
}

// Flags generates flag key-value-map pairs from Options
func (opts *Options) Flags() FlagSets {
	fss := make(FlagSets)

	if opts.MarginTop != nil {
		fss["margin-top"] = *opts.MarginTop
	}

	if opts.MarginBottom != nil {
		fss["margin-bottom"] = *opts.MarginBottom
	}

	if opts.MarginLeft != nil {
		fss["margin-left"] = *opts.MarginLeft
	}

	if opts.MarginRight != nil {
		fss["margin-right"] = *opts.MarginRight
	}

	if opts.PageHeight != nil {
		fss["page-height"] = *opts.PageHeight
	}

	if opts.PageWidth != nil {
		fss["page-width"] = *opts.PageWidth
	}

	if opts.PageSize != nil {
		fss["page-size"] = *opts.PageSize
	}

	return fss
}
