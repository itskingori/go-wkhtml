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

const (
	imageConverterBinary = "wkhtmltoimage"
)

// ImageFlagSet represents key-value pairs of image converter flags
type ImageFlagSet flagSet

// ImageOptions represents wkhtmlimage attributes
type ImageOptions struct {
	CacheDir                *string      `json:"cache_dir,omitempty"`                 // Web cache directory
	Cookie                  *[]CookieSet `json:"cookie,omitempty"`                    // Set an additional cookie with URL encoded values
	CropH                   *int         `json:"crop_h,omitempty"`                    // Set height for cropping
	CropW                   *int         `json:"crop_w,omitempty"`                    // Set width for cropping
	CropX                   *int         `json:"crop_x,omitempty"`                    // Set x coordinate for cropping
	CropY                   *int         `json:"crop_y,omitempty"`                    // Set y coordinate for cropping
	CustomHeader            *[]HeaderSet `json:"custom_header,omitempty"`             // Set an additional HTTP header
	CustomHeaderPropagation *bool        `json:"custom_header_propagation,omitempty"` // Add HTTP headers specified by CustomHeader for each resource request
	DebugJavascript         *bool        `json:"debug_javascript,omitempty"`          // Show javascript debugging output
	Encoding                *string      `json:"encoding,omitempty"`                  // Set the default text encoding, for input
	Format                  *string      `json:"format,omitempty"`                    // Output file format
	Height                  *int         `json:"height,omitempty"`                    // Set screen height
	Images                  *bool        `json:"images,omitempty"`                    // Load or print images
	Javascript              *bool        `json:"javascript,omitempty"`                // Allow web pages to run javascript
	JavascriptDelay         *int         `json:"javascript_delay,omitempty"`          // Milliseconds to wait for javascript to finish
	LoadErrorHandling       *string      `json:"load_error_handling,omitempty"`       // Specify how to handle pages that fail to load
	LoadMediaErrorHandling  *string      `json:"load_media_error_handling,omitempty"` // Specify how to handle media files that fail to load
	MinimumFontSize         *int         `json:"minimum_font_size,omitempty"`         // Minimum font size
	Password                *string      `json:"password,omitempty"`                  // HTTP Authentication password
	Quality                 *int         `json:"quality,omitempty"`                   // Output image quality
	SmartWidth              *bool        `json:"smart_width,omitempty"`               // Extend width to fit unbreakable content or use the specified width (even if it is not large enough for the content)
	StopSlowScripts         *bool        `json:"stop_slow_scripts,omitempty"`         // Stop slow running javascripts
	Transparent             *bool        `json:"transparent,omitempty"`               // Make the background transparent in PNGs
	UseXServer              *bool        `json:"use_xserver,omitempty"`               // Use the X server
	Username                *string      `json:"username,omitempty"`                  // HTTP Authentication username
	Width                   *int         `json:"width,omitempty"`                     // Set screen width, as a guide (needs SmartWidth disabled to enforce)
	Zoom                    *float64     `json:"zoom,omitempty"`                      // Use this zoom factor
}

// NewImageFlagSetFromOptions generates a FlagSet from ImageOptions
func NewImageFlagSetFromOptions(opts *ImageOptions) ImageFlagSet {
	ifs := make(ImageFlagSet)

	if opts.CacheDir != nil {
		ifs.SetCacheDir(*opts.CacheDir)
	}

	if opts.Cookie != nil {
		ifs.SetCookie(*opts.Cookie)
	}

	if opts.CropH != nil {
		ifs.SetCropH(*opts.CropH)
	}

	if opts.CropW != nil {
		ifs.SetCropW(*opts.CropW)
	}

	if opts.CropX != nil {
		ifs.SetCropX(*opts.CropX)
	}

	if opts.CropY != nil {
		ifs.SetCropY(*opts.CropY)
	}

	if opts.CustomHeader != nil {
		ifs.SetCustomHeader(*opts.CustomHeader)
	}

	if opts.CustomHeaderPropagation != nil {
		ifs.SetCustomHeaderPropagation(*opts.CustomHeaderPropagation)
	}

	if opts.DebugJavascript != nil {
		ifs.SetDebugJavascript(*opts.DebugJavascript)
	}

	if opts.Encoding != nil {
		ifs.SetEncoding(*opts.Encoding)
	}

	if opts.Format != nil {
		ifs.SetFormat(*opts.Format)
	}

	if opts.Height != nil {
		ifs.SetHeight(*opts.Height)
	}

	if opts.Images != nil {
		ifs.SetImages(*opts.Images)
	}

	if opts.Javascript != nil {
		ifs.SetJavascript(*opts.Javascript)
	}

	if opts.JavascriptDelay != nil {
		ifs.SetJavascriptDelay(*opts.JavascriptDelay)
	}

	if opts.LoadErrorHandling != nil {
		ifs.SetLoadErrorHandling(*opts.LoadErrorHandling)
	}

	if opts.LoadMediaErrorHandling != nil {
		ifs.SetLoadMediaErrorHandling(*opts.LoadMediaErrorHandling)
	}

	if opts.MinimumFontSize != nil {
		ifs.SetMinimumFontSize(*opts.MinimumFontSize)
	}

	if opts.Password != nil {
		ifs.SetPassword(*opts.Password)
	}

	if opts.Quality != nil {
		ifs.SetQuality(*opts.Quality)
	}

	if opts.SmartWidth != nil {
		ifs.SetSmartWidth(*opts.SmartWidth)
	}

	if opts.StopSlowScripts != nil {
		ifs.SetStopSlowScripts(*opts.StopSlowScripts)
	}

	if opts.Transparent != nil {
		ifs.SetTransparent(*opts.Transparent)
	}

	if opts.UseXServer != nil {
		ifs.SetUseXServer(*opts.UseXServer)
	}

	if opts.Username != nil {
		ifs.SetUsername(*opts.Username)
	}

	if opts.Width != nil {
		ifs.SetWidth(*opts.Width)
	}

	if opts.Zoom != nil {
		ifs.SetZoom(*opts.Zoom)
	}

	return ifs
}

// Flags generates a String slice from an ImageFlagSet
func (ifs *ImageFlagSet) Flags() []string {
	var flags []string

	for flagKey, flagValue := range *ifs {
		switch flagValue.(type) {
		case int:
			evaluateIntFlag(&flags, flagKey, flagValue.(int))
		case string:
			evaluateStringFlag(&flags, flagKey, flagValue.(string))
		case float64:
			evaluateFloat64Flag(&flags, flagKey, flagValue.(float64))
		case []string:
			evaluateStringSliceFlag(&flags, flagKey, flagValue.([]string))
		case []CookieSet:
			evaluateCookieSetSliceFlag(&flags, flagKey, flagValue.([]CookieSet))
		case []HeaderSet:
			evaluateHeaderSetSliceFlag(&flags, flagKey, flagValue.([]HeaderSet))
		case bool:
			// Positive --XXX, Negative --no-XXX
			type1Flags := []string{
				"custom-header-propagation",
				"debug-javascript",
				"images",
				"stop-slow-scripts",
			}

			// Positive --enable-XXX, Negative --disable-XXX
			type2Flags := []string{
				"javascript",
				"smart-width",
			}

			// Positive --XXX, Negative (absent)
			type3Flags := []string{
				"transparent",
				"use-xserver",
			}

			if checkStringSliceContains(type1Flags, flagKey) {
				evaluateBoolType1Flag(&flags, flagKey, flagValue.(bool))
			}

			if checkStringSliceContains(type2Flags, flagKey) {
				evaluateBoolType2Flag(&flags, flagKey, flagValue.(bool))
			}

			if checkStringSliceContains(type3Flags, flagKey) {
				evaluateBoolType3Flag(&flags, flagKey, flagValue.(bool))
			}
		}
	}

	return flags
}

// GetCacheDir retrieves the CacheDir from an ImageFlagSet
func (ifs *ImageFlagSet) GetCacheDir() (string, bool) {
	dir, exists := (*ifs)["cache-dir"]

	return dir.(string), exists
}

// GetCookie retrieves the Cookie from an ImageFlagSet
func (ifs *ImageFlagSet) GetCookie() ([]CookieSet, bool) {
	cookies, exists := (*ifs)["cookie"]

	return cookies.([]CookieSet), exists
}

// GetCropH retrieves the CropH from an ImageFlagSet
func (ifs *ImageFlagSet) GetCropH() (int, bool) {
	height, exists := (*ifs)["crop-h"]

	return height.(int), exists
}

// GetCropW retrieves the CropW from an ImageFlagSet
func (ifs *ImageFlagSet) GetCropW() (int, bool) {
	width, exists := (*ifs)["crop-w"]

	return width.(int), exists
}

// GetCropX retrieves the CropX from an ImageFlagSet
func (ifs *ImageFlagSet) GetCropX() (int, bool) {
	xc, exists := (*ifs)["crop-x"]

	return xc.(int), exists
}

// GetCropY retrieves the CropY from an ImageFlagSet
func (ifs *ImageFlagSet) GetCropY() (int, bool) {
	yc, exists := (*ifs)["crop-y"]

	return yc.(int), exists
}

// GetCustomHeader retrieves the CustomHeader from an ImageFlagSet
func (ifs *ImageFlagSet) GetCustomHeader() ([]HeaderSet, bool) {
	headers, exists := (*ifs)["custom-header"]

	return headers.([]HeaderSet), exists
}

// GetCustomHeaderPropagation retrieves the CustomHeaderPropagation from an ImageFlagSet
func (ifs *ImageFlagSet) GetCustomHeaderPropagation() (bool, bool) {
	value, exists := (*ifs)["custom-header-propagation"]

	return value.(bool), exists
}

// GetDebugJavascript retrieves the DebugJavascript from an ImageFlagSet
func (ifs *ImageFlagSet) GetDebugJavascript() (bool, bool) {
	value, exists := (*ifs)["debug-javascript"]

	return value.(bool), exists
}

// GetEncoding retrieves the Encoding from an ImageFlagSet
func (ifs *ImageFlagSet) GetEncoding() (string, bool) {
	encoding, exists := (*ifs)["encoding"]

	return encoding.(string), exists
}

// GetFormat retrieves the Format from an ImageFlagSet
func (ifs *ImageFlagSet) GetFormat() (string, bool) {
	format, exists := (*ifs)["format"]

	return format.(string), exists
}

// GetHeight retrieves the Height from an ImageFlagSet
func (ifs *ImageFlagSet) GetHeight() (int, bool) {
	height, exists := (*ifs)["height"]

	return height.(int), exists
}

// GetImages retrieves the Images from an ImageFlagSet
func (ifs *ImageFlagSet) GetImages() (bool, bool) {
	value, exists := (*ifs)["images"]

	return value.(bool), exists
}

// GetJavascript retrieves the Javascript from an ImageFlagSet
func (ifs *ImageFlagSet) GetJavascript() (bool, bool) {
	value, exists := (*ifs)["javascript"]

	return value.(bool), exists
}

// GetJavascriptDelay retrieves the JavascriptDelay from an ImageFlagSet
func (ifs *ImageFlagSet) GetJavascriptDelay() (int, bool) {
	milliseconds, exists := (*ifs)["javascript-delay"]

	return milliseconds.(int), exists
}

// GetLoadErrorHandling retrieves the LoadErrorHandling from an ImageFlagSet
func (ifs *ImageFlagSet) GetLoadErrorHandling() (string, bool) {
	handling, exists := (*ifs)["load-error-handling"]

	return handling.(string), exists
}

// GetLoadMediaErrorHandling retrieves the LoadMediaErrorHandling from an ImageFlagSet
func (ifs *ImageFlagSet) GetLoadMediaErrorHandling() (string, bool) {
	handling, exists := (*ifs)["load-media-error-handling"]

	return handling.(string), exists
}

// GetMinimumFontSize retrieves the MinimumFontSize from an ImageFlagSet
func (ifs *ImageFlagSet) GetMinimumFontSize() (int, bool) {
	size, exists := (*ifs)["minimum-font-size"]

	return size.(int), exists
}

// GetPassword retrieves the Password from an ImageFlagSet
func (ifs *ImageFlagSet) GetPassword() (string, bool) {
	password, exists := (*ifs)["password"]

	return password.(string), exists
}

// GetQuality retrieves the Quality from an ImageFlagSet
func (ifs *ImageFlagSet) GetQuality() (int, bool) {
	quality, exists := (*ifs)["quality"]

	return quality.(int), exists
}

// GetSmartWidth retrieves the SmartWidth from an ImageFlagSet
func (ifs *ImageFlagSet) GetSmartWidth() (bool, bool) {
	value, exists := (*ifs)["smart-width"]

	return value.(bool), exists
}

// GetStopSlowScripts retrieves the StopSlowScripts from an ImageFlagSet
func (ifs *ImageFlagSet) GetStopSlowScripts() (bool, bool) {
	value, exists := (*ifs)["stop-slows-cripts"]

	return value.(bool), exists
}

// GetTransparent retrieves the Transparent from an ImageFlagSet
func (ifs *ImageFlagSet) GetTransparent() (bool, bool) {
	value, exists := (*ifs)["transparent"]

	return value.(bool), exists
}

// GetUseXServer retrieves the UseXServer from an ImageFlagSet
func (ifs *ImageFlagSet) GetUseXServer() (bool, bool) {
	value, exists := (*ifs)["use-xserver"]

	return value.(bool), exists
}

// GetUsername retrieves the Username from an ImageFlagSet
func (ifs *ImageFlagSet) GetUsername() (string, bool) {
	username, exists := (*ifs)["username"]

	return username.(string), exists
}

// GetWidth retrieves the Width from an ImageFlagSet
func (ifs *ImageFlagSet) GetWidth() (int, bool) {
	width, exists := (*ifs)["width"]

	return width.(int), exists
}

// GetZoom retrieves the Zoom from an ImageFlagSet
func (ifs *ImageFlagSet) GetZoom() (float64, bool) {
	zoom, exists := (*ifs)["zoom"]

	return zoom.(float64), exists
}

// SetCacheDir sets the CacheDir of an ImageFlagSet
func (ifs *ImageFlagSet) SetCacheDir(dir string) {
	(*ifs)["cache-dir"] = dir
}

// SetCookie sets the Cookie of an ImageFlagSet
func (ifs *ImageFlagSet) SetCookie(cookies []CookieSet) {
	(*ifs)["cookie"] = cookies
}

// SetCropH sets the CropH of an ImageFlagSet
func (ifs *ImageFlagSet) SetCropH(height int) {
	(*ifs)["crop-h"] = height
}

// SetCropW sets the CropW of an ImageFlagSet
func (ifs *ImageFlagSet) SetCropW(width int) {
	(*ifs)["crop-w"] = width
}

// SetCropX sets the CropX of an ImageFlagSet
func (ifs *ImageFlagSet) SetCropX(xc int) {
	(*ifs)["crop-x"] = xc
}

// SetCropY sets the CropY of an ImageFlagSet
func (ifs *ImageFlagSet) SetCropY(yc int) {
	(*ifs)["crop-y"] = yc
}

// SetCustomHeader sets the CustomHeader of an ImageFlagSet
func (ifs *ImageFlagSet) SetCustomHeader(headers []HeaderSet) {
	(*ifs)["custom-header"] = headers
}

// SetCustomHeaderPropagation sets the CustomHeaderPropagation of an ImageFlagSet
func (ifs *ImageFlagSet) SetCustomHeaderPropagation(value bool) {
	(*ifs)["custom-header-propagation"] = value
}

// SetDebugJavascript sets the DebugJavascript of an ImageFlagSet
func (ifs *ImageFlagSet) SetDebugJavascript(value bool) {
	(*ifs)["debug-javascript"] = value
}

// SetEncoding sets the Encoding of an ImageFlagSet
func (ifs *ImageFlagSet) SetEncoding(encoding string) {
	(*ifs)["encoding"] = encoding
}

// SetFormat sets the Format of an ImageFlagSet
func (ifs *ImageFlagSet) SetFormat(format string) {
	(*ifs)["format"] = format
}

// SetHeight sets the Height of an ImageFlagSet
func (ifs *ImageFlagSet) SetHeight(height int) {
	(*ifs)["height"] = height
}

// SetImages sets the Images of an ImageFlagSet
func (ifs *ImageFlagSet) SetImages(value bool) {
	(*ifs)["images"] = value
}

// SetJavascript sets the Javascript of an ImageFlagSet
func (ifs *ImageFlagSet) SetJavascript(value bool) {
	(*ifs)["javascript"] = value
}

// SetJavascriptDelay sets the JavascriptDelay of an ImageFlagSet
func (ifs *ImageFlagSet) SetJavascriptDelay(milliseconds int) {
	(*ifs)["javascript-delay"] = milliseconds
}

// SetLoadErrorHandling sets the LoadErrorHandling of an ImageFlagSet
func (ifs *ImageFlagSet) SetLoadErrorHandling(handling string) {
	(*ifs)["load-error-handling"] = handling
}

// SetLoadMediaErrorHandling sets the LoadMediaErrorHandling of an ImageFlagSet
func (ifs *ImageFlagSet) SetLoadMediaErrorHandling(handling string) {
	(*ifs)["load-media-error-handling"] = handling
}

// SetMinimumFontSize sets the MinimumFontSize of an ImageFlagSet
func (ifs *ImageFlagSet) SetMinimumFontSize(size int) {
	(*ifs)["minimum-font-size"] = size
}

// SetPassword sets the Password of an ImageFlagSet
func (ifs *ImageFlagSet) SetPassword(password string) {
	(*ifs)["password"] = password
}

// SetQuality sets the Quality of an ImageFlagSet
func (ifs *ImageFlagSet) SetQuality(quality int) {
	(*ifs)["quality"] = quality
}

// SetSmartWidth sets the SmartWidth of an ImageFlagSet
func (ifs *ImageFlagSet) SetSmartWidth(value bool) {
	(*ifs)["smart-width"] = value
}

// SetStopSlowScripts sets the StopSlowScripts of an ImageFlagSet
func (ifs *ImageFlagSet) SetStopSlowScripts(value bool) {
	(*ifs)["stop-slows-cripts"] = value
}

// SetTransparent sets the Transparent of an ImageFlagSet
func (ifs *ImageFlagSet) SetTransparent(value bool) {
	(*ifs)["transparent"] = value
}

// SetUseXServer sets the UseXServer of an ImageFlagSet
func (ifs *ImageFlagSet) SetUseXServer(value bool) {
	(*ifs)["use-xserver"] = value
}

// SetUsername sets the Username of an ImageFlagSet
func (ifs *ImageFlagSet) SetUsername(username string) {
	(*ifs)["username"] = username
}

// SetWidth sets the Width of an ImageFlagSet
func (ifs *ImageFlagSet) SetWidth(width int) {
	(*ifs)["width"] = width
}

// SetZoom sets the Zoom of an ImageFlagSet
func (ifs *ImageFlagSet) SetZoom(zoom float64) {
	(*ifs)["zoom"] = zoom
}

// Generate performs the image conversion and saves the file to disk
func (ifs *ImageFlagSet) Generate(inputURL string, outputFile string) ([]byte, error) {
	out, err := runConversionCommand(imageConverterBinary, ifs.Flags(), &inputURL, &outputFile)

	return out, err
}
