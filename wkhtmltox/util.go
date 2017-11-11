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
	"strings"

	"os/exec"
)

// LookupConverter checks if converter executable exists and returns
// it's path and version.
func LookupConverter(name string) (string, string, error) {
	var version string

	path, err := exec.LookPath(name)
	if err != nil {
		return path, version, err
	}

	cmd := exec.Command(name, "--version")
	out, err := cmd.CombinedOutput()
	version = strings.TrimRight(string(out), "\r\n")

	return path, version, err
}
