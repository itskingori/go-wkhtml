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
	"os/exec"
)

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

func checkStringSliceContains(ss []string, str string) bool {
	for _, v := range ss {
		if v == str {
			return true
		}
	}

	return false
}

func evaluateIntFlag(flags *[]string, flagKey string, flagValue int) {
	*flags = append(*flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%d", flagValue))
}

func evaluateStringFlag(flags *[]string, flagKey string, flagValue string) {
	*flags = append(*flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%s", flagValue))
}

func evaluateFloat64Flag(flags *[]string, flagKey string, flagValue float64) {
	*flags = append(*flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%v", flagValue))
}

func evaluateCookieSetSliceFlag(flags *[]string, flagKey string, flagValue []CookieSet) {
	for _, cs := range flagValue {
		*flags = append(*flags, fmt.Sprintf("--%s", flagKey), cs.Name, cs.Value)
	}
}

func evaluateHeaderSetSliceFlag(flags *[]string, flagKey string, flagValue []HeaderSet) {
	for _, hs := range flagValue {
		*flags = append(*flags, fmt.Sprintf("--%s", flagKey), hs.Name, hs.Value)
	}
}

func evaluateStringSliceFlag(flags *[]string, flagKey string, flagValue []string) {
	for _, flagValueChild := range flagValue {
		*flags = append(*flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%s", flagValueChild))
	}
}

func evaluateBoolType1Flag(flags *[]string, flagKey string, flagValue bool) {
	if flagValue {
		*flags = append(*flags, fmt.Sprintf("--%s", flagKey))
	} else {
		*flags = append(*flags, fmt.Sprintf("--no-%s", flagKey))
	}
}

func evaluateBoolType2Flag(flags *[]string, flagKey string, flagValue bool) {
	if flagValue {
		*flags = append(*flags, fmt.Sprintf("--enable-%s", flagKey))
	} else {
		*flags = append(*flags, fmt.Sprintf("--disable-%s", flagKey))
	}
}

func evaluateBoolType3Flag(flags *[]string, flagKey string, flagValue bool) {
	if flagValue {
		*flags = append(*flags, fmt.Sprintf("--%s", flagKey))
	}
}

func runConversionCommand(binary string, params []string, inputURL *string, outputFile *string) ([]byte, error) {
	var out []byte

	// I'm uncertain if we need to escape parameters ... can't seem to find
	// anything conclusive yet, but so far this seems to be the best find:
	// https://stackoverflow.com/a/8025343/2184155

	params = append(params, *inputURL, *outputFile)
	cmd := exec.Command(binary, params...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, err
}
