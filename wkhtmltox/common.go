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

type flagSet map[string]interface{}

// CookieSet represents cookie name and value
type CookieSet struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// HeaderSet represents header name and value
type HeaderSet struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
