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
	pdfConverterBinary = "wkhtmltopdf"
)

// PDFFlagSet represents key-value pairs of PDF converter flags
type PDFFlagSet flagSet

// PDFOptions represents wkhtmlpdf attributes
type PDFOptions struct {
	CacheDir                *string      `json:"cache_dir,omitempty"`                 // Web cache directory
	Cookie                  *[]CookieSet `json:"cookies,omitempty"`                   // Set an additional cookie with url encoded values
	CustomHeader            *[]HeaderSet `json:"custom_headers,omitempty"`            // Set an additional HTTP header
	CustomHeaderPropagation *bool        `json:"custom_header_propagation,omitempty"` // Add HTTP headers specified by CustomHeader for each resource request
	DebugJavascript         *bool        `json:"debug_javascript,omitempty"`          // Show javascript debugging output
	DPI                     *int         `json:"dpi,omitempty"`                       // Change the DPI explicitly
	Encoding                *string      `json:"encoding,omitempty"`                  // Set the default text encoding, for input
	ExternalLinks           *bool        `json:"external_links,omitempty"`            // Make links to remote web pages
	Forms                   *bool        `json:"forms,omitempty"`                     // Turn HTML form fields into pdf form fields
	Grayscale               *bool        `json:"grayscale,omitempty"`                 // Generate the PDF in grayscale
	Images                  *bool        `json:"images,omitempty"`                    // Load or print images
	ImageDPI                *int         `json:"image_dpi,omitempty"`                 // Scale down images to this DPI when embedding images
	ImageQuality            *int         `json:"image_quality,omitempty"`             // JPEG compress images to this quality
	InternalLinks           *bool        `json:"internal_links,omitempty"`            // Make local links
	Javascript              *bool        `json:"javascript,omitempty"`                // Allow web pages to run javascript
	JavascriptDelay         *int         `json:"javascript_delay,omitempty"`          // Milliseconds to wait for javascript to finish
	LoadErrorHandling       *string      `json:"load_error_handling,omitempty"`       // Specify how to handle pages that fail to load
	LoadMediaErrorHandling  *string      `json:"load_media_error_handling,omitempty"` // Specify how to handle media files that fail to load
	LowQuality              *bool        `json:"lowquality,omitempty"`                // Generates lower quality PDF/PS
	MarginBottom            *int         `json:"margin_bottom,omitempty"`             // Set the page bottom margin
	MarginLeft              *int         `json:"margin_left,omitempty"`               // Set the page left margin
	MarginRight             *int         `json:"margin_right,omitempty"`              // Set the page right margin
	MarginTop               *int         `json:"margin_top,omitempty"`                // Set the page top margin
	MinimumFontSize         *int         `json:"minimum_font_size,omitempty"`         // Minimum font size
	NoPDFCompression        *bool        `json:"no_pdf_compression,omitempty"`        // Do not use lossless compression on PDF objects
	Orientation             *string      `json:"orientation,omitempty"`               // Set orientation to landscape or portrait
	PageHeight              *int         `json:"page_height,omitempty"`               // Height of the page
	PageSize                *string      `json:"page_size,omitempty"`                 // Size of the page
	PageWidth               *int         `json:"page_width,omitempty"`                // Width of the page
	Password                *string      `json:"password,omitempty"`                  // HTTP Authentication password
	SmartShrinking          *bool        `json:"smart_width,omitempty"`               // Enable the intelligent shrinking strategy used by WebKit that makes the pixel/dpi ratio none constant
	StopSlowScripts         *bool        `json:"stop_slow_scripts,omitempty"`         // Stop slow running javascripts
	Title                   *string      `json:"title,omitempty"`                     // The title of the generated PDF file
	UseXServer              *bool        `json:"use_xserver,omitempty"`               // Use the X server
	Username                *string      `json:"username,omitempty"`                  // HTTP Authentication username
	Zoom                    *float64     `json:"zoom,omitempty"`                      // Use this zoom factor
}

// NewPDFFlagSetFromOptions generates a FlagSet from PDFOptions
func NewPDFFlagSetFromOptions(opts *PDFOptions) PDFFlagSet {
	pfs := make(PDFFlagSet)

	if opts.CacheDir != nil {
		pfs.SetCacheDir(*opts.CacheDir)
	}

	if opts.Cookie != nil {
		pfs.SetCookie(*opts.Cookie)
	}

	if opts.CustomHeader != nil {
		pfs.SetCustomHeader(*opts.CustomHeader)
	}

	if opts.CustomHeaderPropagation != nil {
		pfs.SetCustomHeaderPropagation(*opts.CustomHeaderPropagation)
	}

	if opts.DebugJavascript != nil {
		pfs.SetDebugJavascript(*opts.DebugJavascript)
	}

	if opts.DPI != nil {
		pfs.SetDPI(*opts.DPI)
	}

	if opts.Encoding != nil {
		pfs.SetEncoding(*opts.Encoding)
	}

	if opts.ExternalLinks != nil {
		pfs.SetExternalLinks(*opts.ExternalLinks)
	}

	if opts.Forms != nil {
		pfs.SetForms(*opts.Forms)
	}

	if opts.Grayscale != nil {
		pfs.SetGrayscale(*opts.Grayscale)
	}

	if opts.Images != nil {
		pfs.SetImages(*opts.Images)
	}

	if opts.ImageDPI != nil {
		pfs.SetImageDPI(*opts.ImageDPI)
	}

	if opts.ImageQuality != nil {
		pfs.SetImageQuality(*opts.ImageQuality)
	}

	if opts.InternalLinks != nil {
		pfs.SetInternalLinks(*opts.InternalLinks)
	}

	if opts.Javascript != nil {
		pfs.SetJavascript(*opts.Javascript)
	}

	if opts.JavascriptDelay != nil {
		pfs.SetJavascriptDelay(*opts.JavascriptDelay)
	}

	if opts.LoadErrorHandling != nil {
		pfs.SetLoadErrorHandling(*opts.LoadErrorHandling)
	}

	if opts.LoadMediaErrorHandling != nil {
		pfs.SetLoadMediaErrorHandling(*opts.LoadMediaErrorHandling)
	}

	if opts.LowQuality != nil {
		pfs.SetLowQuality(*opts.LowQuality)
	}

	if opts.MarginBottom != nil {
		pfs.SetMarginBottom(*opts.MarginBottom)
	}

	if opts.MarginLeft != nil {
		pfs.SetMarginLeft(*opts.MarginLeft)
	}

	if opts.MarginRight != nil {
		pfs.SetMarginRight(*opts.MarginRight)
	}

	if opts.MarginTop != nil {
		pfs.SetMarginTop(*opts.MarginTop)
	}

	if opts.MinimumFontSize != nil {
		pfs.SetMinimumFontSize(*opts.MinimumFontSize)
	}

	if opts.NoPDFCompression != nil {
		pfs.SetNoPDFCompression(*opts.NoPDFCompression)
	}

	if opts.Orientation != nil {
		pfs.SetOrientation(*opts.Orientation)
	}

	if opts.PageHeight != nil {
		pfs.SetPageHeight(*opts.PageHeight)
	}

	if opts.PageSize != nil {
		pfs.SetPageSize(*opts.PageSize)
	}

	if opts.PageWidth != nil {
		pfs.SetPageWidth(*opts.PageWidth)
	}

	if opts.Password != nil {
		pfs.SetPassword(*opts.Password)
	}

	if opts.SmartShrinking != nil {
		pfs.SetSmartShrinking(*opts.SmartShrinking)
	}

	if opts.StopSlowScripts != nil {
		pfs.SetStopSlowScripts(*opts.StopSlowScripts)
	}

	if opts.Title != nil {
		pfs.SetTitle(*opts.Title)
	}

	if opts.UseXServer != nil {
		pfs.SetUseXServer(*opts.UseXServer)
	}

	if opts.Username != nil {
		pfs.SetUsername(*opts.Username)
	}

	if opts.Zoom != nil {
		pfs.SetZoom(*opts.Zoom)
	}

	return pfs
}

// Flags generates a String slice from a PDFFlagSet
func (pfs *PDFFlagSet) Flags() []string {
	var flags []string

	for flagKey, flagValue := range *pfs {
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
				"external-links",
				"forms",
				"internal-links",
				"javascript",
				"smart-shrinking",
			}

			// Positive --XXX, Negative (absent)
			type3Flags := []string{
				"grayscale",
				"lowquality",
				"no-pdf-compression",
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

// GetCacheDir retrieves the CacheDir from a PDFFlagSet
func (pfs *PDFFlagSet) GetCacheDir() (string, bool) {
	dir, exists := (*pfs)["cache-dir"]

	return dir.(string), exists
}

// GetCookie retrieves the Cookie from a PDFFlagSet
func (pfs *PDFFlagSet) GetCookie() ([]CookieSet, bool) {
	cookies, exists := (*pfs)["cookie"]

	return cookies.([]CookieSet), exists
}

// GetCustomHeader retrieves the CustomHeader from a PDFFlagSet
func (pfs *PDFFlagSet) GetCustomHeader() ([]HeaderSet, bool) {
	headers, exists := (*pfs)["custom-header"]

	return headers.([]HeaderSet), exists
}

// GetCustomHeaderPropagation retrieves the CustomHeaderPropagation from a PDFFlagSet
func (pfs *PDFFlagSet) GetCustomHeaderPropagation() (bool, bool) {
	value, exists := (*pfs)["custom-header-propagation"]

	return value.(bool), exists
}

// GetDebugJavascript retrieves the DebugJavascript from a PDFFlagSet
func (pfs *PDFFlagSet) GetDebugJavascript() (bool, bool) {
	value, exists := (*pfs)["debug-javascript"]

	return value.(bool), exists
}

// GetDPI retrieves the DPI from a PDFFlagSet
func (pfs *PDFFlagSet) GetDPI() (int, bool) {
	dpi, exists := (*pfs)["dpi"]

	return dpi.(int), exists
}

// GetEncoding retrieves the Encoding from a PDFFlagSet
func (pfs *PDFFlagSet) GetEncoding() (string, bool) {
	encoding, exists := (*pfs)["encoding"]

	return encoding.(string), exists
}

// GetExternalLinks retrieves the ExternalLinks from a PDFFlagSet
func (pfs *PDFFlagSet) GetExternalLinks() (bool, bool) {
	value, exists := (*pfs)["external-links"]

	return value.(bool), exists
}

// GetForms retrieves the Forms from a PDFFlagSet
func (pfs *PDFFlagSet) GetForms() (bool, bool) {
	value, exists := (*pfs)["forms"]

	return value.(bool), exists
}

// GetGrayscale retrieves the Grayscale from a PDFFlagSet
func (pfs *PDFFlagSet) GetGrayscale() (bool, bool) {
	value, exists := (*pfs)["grayscale"]

	return value.(bool), exists
}

// GetImages retrieves the Images from a PDFFlagSet
func (pfs *PDFFlagSet) GetImages() (bool, bool) {
	value, exists := (*pfs)["images"]

	return value.(bool), exists
}

// GetImageDPI retrieves the ImageDPI from a PDFFlagSet
func (pfs *PDFFlagSet) GetImageDPI() (int, bool) {
	dpi, exists := (*pfs)["image-dpi"]

	return dpi.(int), exists
}

// GetImageQuality retrieves the ImageQuality from a PDFFlagSet
func (pfs *PDFFlagSet) GetImageQuality() (int, bool) {
	quality, exists := (*pfs)["image-quality"]

	return quality.(int), exists
}

// GetInternalLinks retrieves the InternalLinks from a PDFFlagSet
func (pfs *PDFFlagSet) GetInternalLinks() (bool, bool) {
	value, exists := (*pfs)["internal-links"]

	return value.(bool), exists
}

// GetJavascript retrieves the Javascript from a PDFFlagSet
func (pfs *PDFFlagSet) GetJavascript() (bool, bool) {
	value, exists := (*pfs)["javascript"]

	return value.(bool), exists
}

// GetJavascriptDelay retrieves the JavascriptDelay from a PDFFlagSet
func (pfs *PDFFlagSet) GetJavascriptDelay() (int, bool) {
	milliseconds, exists := (*pfs)["javascript-delay"]

	return milliseconds.(int), exists
}

// GetLoadErrorHandling retrieves the LoadErrorHandling from a PDFFlagSet
func (pfs *PDFFlagSet) GetLoadErrorHandling() (string, bool) {
	handling, exists := (*pfs)["load-error-handling"]

	return handling.(string), exists
}

// GetLoadMediaErrorHandling retrieves the LoadMediaErrorHandling from a PDFFlagSet
func (pfs *PDFFlagSet) GetLoadMediaErrorHandling() (string, bool) {
	handling, exists := (*pfs)["load-media-error-handling"]

	return handling.(string), exists
}

// GetLowQuality retrieves the LowQuality from a PDFFlagSet
func (pfs *PDFFlagSet) GetLowQuality() (bool, bool) {
	value, exists := (*pfs)["lowquality"]

	return value.(bool), exists
}

// GetMarginBottom retrieves the MarginBottom from a PDFFlagSet
func (pfs *PDFFlagSet) GetMarginBottom() (int, bool) {
	millimetres, exists := (*pfs)["margin-bottom"]

	return millimetres.(int), exists
}

// GetMarginLeft retrieves the MarginLeft from a PDFFlagSet
func (pfs *PDFFlagSet) GetMarginLeft() (int, bool) {
	millimetres, exists := (*pfs)["margin-left"]

	return millimetres.(int), exists
}

// GetMarginRight retrieves the MarginRight from a PDFFlagSet
func (pfs *PDFFlagSet) GetMarginRight() (int, bool) {
	millimetres, exists := (*pfs)["margin-right"]

	return millimetres.(int), exists
}

// GetMarginTop retrieves the MarginTop from a PDFFlagSet
func (pfs *PDFFlagSet) GetMarginTop() (int, bool) {
	millimetres, exists := (*pfs)["margin-top"]

	return millimetres.(int), exists
}

// GetMinimumFontSize retrieves the MinimumFontSize from a PDFFlagSet
func (pfs *PDFFlagSet) GetMinimumFontSize() (int, bool) {
	size, exists := (*pfs)["minimum-font-size"]

	return size.(int), exists
}

// GetNoPDFCompression retrieves the NoPDFCompression from a PDFFlagSet
func (pfs *PDFFlagSet) GetNoPDFCompression() (bool, bool) {
	value, exists := (*pfs)["no-pdf-compression"]

	return value.(bool), exists
}

// GetOrientation retrieves the Orientation from a PDFFlagSet
func (pfs *PDFFlagSet) GetOrientation() (string, bool) {
	orientation, exists := (*pfs)["orientation"]

	return orientation.(string), exists
}

// GetPageHeight retrieves the PageHeight from a PDFFlagSet
func (pfs *PDFFlagSet) GetPageHeight() (int, bool) {
	height, exists := (*pfs)["page-height"]

	return height.(int), exists
}

// GetPageSize retrieves the PageSize from a PDFFlagSet
func (pfs *PDFFlagSet) GetPageSize() (string, bool) {
	size, exists := (*pfs)["page-size"]

	return size.(string), exists
}

// GetPageWidth retrieves the PageWidth from a PDFFlagSet
func (pfs *PDFFlagSet) GetPageWidth() (int, bool) {
	width, exists := (*pfs)["page-width"]

	return width.(int), exists
}

// GetPassword retrieves the Password from a PDFFlagSet
func (pfs *PDFFlagSet) GetPassword() (string, bool) {
	password, exists := (*pfs)["password"]

	return password.(string), exists
}

// GetSmartShrinking retrieves the SmartShrinking from a PDFFlagSet
func (pfs *PDFFlagSet) GetSmartShrinking() (bool, bool) {
	value, exists := (*pfs)["smart-width"]

	return value.(bool), exists
}

// GetStopSlowScripts retrieves the StopSlowScripts from a PDFFlagSet
func (pfs *PDFFlagSet) GetStopSlowScripts() (bool, bool) {
	value, exists := (*pfs)["stop-slow-scripts"]

	return value.(bool), exists
}

// GetTitle retrieves the Title from a PDFFlagSet
func (pfs *PDFFlagSet) GetTitle() (string, bool) {
	title, exists := (*pfs)["title"]

	return title.(string), exists
}

// GetUseXServer retrieves the UseXServer from a PDFFlagSet
func (pfs *PDFFlagSet) GetUseXServer() (bool, bool) {
	value, exists := (*pfs)["use-xserver"]

	return value.(bool), exists
}

// GetUsername retrieves the Username from a PDFFlagSet
func (pfs *PDFFlagSet) GetUsername() (string, bool) {
	username, exists := (*pfs)["username"]

	return username.(string), exists
}

// GetZoom retrieves the Zoom from a PDFFlagSet
func (pfs *PDFFlagSet) GetZoom() (float64, bool) {
	zoom, exists := (*pfs)["zoom"]

	return zoom.(float64), exists
}

// SetCacheDir sets the CacheDir of a PDFFlagSet
func (pfs *PDFFlagSet) SetCacheDir(dir string) {
	(*pfs)["cache-dir"] = dir
}

// SetCookie sets the Cookie of a PDFFlagSet
func (pfs *PDFFlagSet) SetCookie(cookies []CookieSet) {
	(*pfs)["cookie"] = cookies
}

// SetCustomHeader sets the CustomHeader of a PDFFlagSet
func (pfs *PDFFlagSet) SetCustomHeader(headers []HeaderSet) {
	(*pfs)["custom-header"] = headers
}

// SetCustomHeaderPropagation sets the CustomHeaderPropagation of a PDFFlagSet
func (pfs *PDFFlagSet) SetCustomHeaderPropagation(value bool) {
	(*pfs)["custom-header-propagation"] = value
}

// SetDebugJavascript sets the DebugJavascript of a PDFFlagSet
func (pfs *PDFFlagSet) SetDebugJavascript(value bool) {
	(*pfs)["debug-javascript"] = value
}

// SetDPI sets the DPI of a PDFFlagSet
func (pfs *PDFFlagSet) SetDPI(dpi int) {
	(*pfs)["dpi"] = dpi
}

// SetEncoding sets the Encoding of a PDFFlagSet
func (pfs *PDFFlagSet) SetEncoding(encoding string) {
	(*pfs)["encoding"] = encoding
}

// SetExternalLinks sets the ExternalLinks of a PDFFlagSet
func (pfs *PDFFlagSet) SetExternalLinks(value bool) {
	(*pfs)["external-links"] = value
}

// SetForms sets the Forms of a PDFFlagSet
func (pfs *PDFFlagSet) SetForms(value bool) {
	(*pfs)["forms"] = value
}

// SetGrayscale sets the Grayscale of a PDFFlagSet
func (pfs *PDFFlagSet) SetGrayscale(value bool) {
	(*pfs)["grayscale"] = value
}

// SetImages sets the Images of a PDFFlagSet
func (pfs *PDFFlagSet) SetImages(value bool) {
	(*pfs)["images"] = value
}

// SetImageDPI sets the ImageDPI of a PDFFlagSet
func (pfs *PDFFlagSet) SetImageDPI(dpi int) {
	(*pfs)["image-dpi"] = dpi
}

// SetImageQuality sets the ImageQuality of a PDFFlagSet
func (pfs *PDFFlagSet) SetImageQuality(quality int) {
	(*pfs)["image-quality"] = quality
}

// SetInternalLinks sets the InternalLinks of a PDFFlagSet
func (pfs *PDFFlagSet) SetInternalLinks(value bool) {
	(*pfs)["internal-links"] = value
}

// SetJavascript sets the Javascript of a PDFFlagSet
func (pfs *PDFFlagSet) SetJavascript(value bool) {
	(*pfs)["javascript"] = value
}

// SetJavascriptDelay sets the JavascriptDelay of a PDFFlagSet
func (pfs *PDFFlagSet) SetJavascriptDelay(milliseconds int) {
	(*pfs)["javascript-delay"] = milliseconds
}

// SetLoadErrorHandling sets the LoadErrorHandling of a PDFFlagSet
func (pfs *PDFFlagSet) SetLoadErrorHandling(handling string) {
	(*pfs)["load-error-handling"] = handling
}

// SetLoadMediaErrorHandling sets the LoadMediaErrorHandling of a PDFFlagSet
func (pfs *PDFFlagSet) SetLoadMediaErrorHandling(handling string) {
	(*pfs)["load-media-error-handling"] = handling
}

// SetLowQuality sets the LowQuality of a PDFFlagSet
func (pfs *PDFFlagSet) SetLowQuality(value bool) {
	(*pfs)["lowquality"] = value
}

// SetMarginBottom sets the MarginBottom of a PDFFlagSet
func (pfs *PDFFlagSet) SetMarginBottom(millimetres int) {
	(*pfs)["margin-bottom"] = millimetres
}

// SetMarginLeft sets the MarginLeft of a PDFFlagSet
func (pfs *PDFFlagSet) SetMarginLeft(millimetres int) {
	(*pfs)["margin-left"] = millimetres
}

// SetMarginRight sets the MarginRight of a PDFFlagSet
func (pfs *PDFFlagSet) SetMarginRight(millimetres int) {
	(*pfs)["margin-right"] = millimetres
}

// SetMarginTop sets the MarginTop of a PDFFlagSet
func (pfs *PDFFlagSet) SetMarginTop(millimetres int) {
	(*pfs)["margin-top"] = millimetres
}

// SetMinimumFontSize sets the MinimumFontSize of a PDFFlagSet
func (pfs *PDFFlagSet) SetMinimumFontSize(size int) {
	(*pfs)["minimum-font-size"] = size
}

// SetNoPDFCompression sets the NoPDFCompression of a PDFFlagSet
func (pfs *PDFFlagSet) SetNoPDFCompression(value bool) {
	(*pfs)["no-pdf-compression"] = value
}

// SetOrientation sets the Orientation of a PDFFlagSet
func (pfs *PDFFlagSet) SetOrientation(orientation string) {
	(*pfs)["orientation"] = orientation
}

// SetPageHeight sets the PageHeight of a PDFFlagSet
func (pfs *PDFFlagSet) SetPageHeight(height int) {
	(*pfs)["page-height"] = height
}

// SetPageSize sets the PageSize of a PDFFlagSet
func (pfs *PDFFlagSet) SetPageSize(size string) {
	(*pfs)["page-size"] = size
}

// SetPageWidth sets the PageWidth of a PDFFlagSet
func (pfs *PDFFlagSet) SetPageWidth(width int) {
	(*pfs)["page-width"] = width
}

// SetPassword sets the Password of a PDFFlagSet
func (pfs *PDFFlagSet) SetPassword(password string) {
	(*pfs)["password"] = password
}

// SetSmartShrinking sets the SmartShrinking of a PDFFlagSet
func (pfs *PDFFlagSet) SetSmartShrinking(value bool) {
	(*pfs)["smart-width"] = value
}

// SetStopSlowScripts sets the StopSlowScripts of a PDFFlagSet
func (pfs *PDFFlagSet) SetStopSlowScripts(value bool) {
	(*pfs)["stop-slow-scripts"] = value
}

// SetTitle sets the Title of a PDFFlagSet
func (pfs *PDFFlagSet) SetTitle(title string) {
	(*pfs)["title"] = title
}

// SetUseXServer sets the UseXServer of a PDFFlagSet
func (pfs *PDFFlagSet) SetUseXServer(value bool) {
	(*pfs)["use-xserver"] = value
}

// SetUsername sets the Username of a PDFFlagSet
func (pfs *PDFFlagSet) SetUsername(username string) {
	(*pfs)["username"] = username
}

// SetZoom sets the Zoom of a PDFFlagSet
func (pfs *PDFFlagSet) SetZoom(zoom float64) {
	(*pfs)["zoom"] = zoom
}

// Generate performs the PDF conversion and saves the file to disk
func (pfs *PDFFlagSet) Generate(inputURL string, outputFile string) ([]byte, error) {
	out, err := runConversionCommand(pdfConverterBinary, pfs.Flags(), &inputURL, &outputFile)

	return out, err
}
