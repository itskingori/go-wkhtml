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

package image

import (
	"encoding/json"
)

// FlagSets are key-value pair of flags
type FlagSets map[string]interface{}

// Options represents a wkhtmlimage attributes
type Options struct {
	GeneralOptions
}

// GeneralOptions represents wkhtmltoimage general options
type GeneralOptions struct {
	CropH *int `json:"crop_h,omitempty"` // Set height for cropping
	CropW *int `json:"crop_w,omitempty"` // Set width for cropping
	CropX *int `json:"crop_x,omitempty"` // Set x coordinate for cropping
	CropY *int `json:"crop_y,omitempty"` // Set y coordinate for cropping

	Format *string `json:"format,omitempty"` // Output file format
	Height *int    `json:"height,omitempty"` // Set screen height
	Width  *int    `json:"width,omitempty"`  // Set screen width
}

// NewOptionsFromJSON create image options from JSON
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

	if opts.CropH != nil {
		fss["crop-h"] = *opts.CropH
	}

	if opts.CropW != nil {
		fss["crop-w"] = *opts.CropW
	}

	if opts.CropX != nil {
		fss["crop-x"] = *opts.CropX
	}

	if opts.CropY != nil {
		fss["crop-y"] = *opts.CropY
	}

	if opts.Format != nil {
		fss["format"] = *opts.Format
	}

	if opts.Height != nil {
		fss["height"] = *opts.Height
	}

	if opts.Width != nil {
		fss["width"] = *opts.Width
	}

	return fss
}
