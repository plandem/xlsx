// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validator

import (
	"github.com/plandem/xlsx/internal"
	"net/url"
	"regexp"
	"unicode/utf8"
)

//FindNamedMatches uses FindStringSubmatch to match and transform result into map
func FindNamedMatches(pattern *regexp.Regexp, str string) map[string]string {
	match := pattern.FindStringSubmatch(str)
	results := make(map[string]string)

	for i, value := range match {
		if name := pattern.SubexpNames()[i]; name != "" {
			results[name] = value
		}
	}

	return results
}

// IsWinPath check is a str is Win file
func IsWinPath(str string) bool {
	if regWinPath.MatchString(str) {
		//check windows path limit see:
		if len(str) > internal.FilePathLimit {
			return false
		}

		return true
	}

	return false
}

// IsUnixPath check is a str is Unix file
func IsUnixPath(str string) bool {
	return regUnixPath.MatchString(str)
}

// IsFilePath check is a str is Win or Unix file
func IsFilePath(str string) bool {
	return IsWinPath(str) || IsUnixPath(str)
}

// IsURL check if the str is an URL
func IsURL(str string) bool {
	if utf8.RuneCountInString(str) >= internal.UrlLimit {
		return false
	}

	//for XLSX we need more strict rules for url
	u, err := url.ParseRequestURI(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// IsEmail check if the str is an email
func IsEmail(str string) bool {
	return regEmail.MatchString(str)
}

// IsMailTo check if the str is an mailto format
func IsMailTo(str string) (bool, map[string]string) {
	results := FindNamedMatches(regMailTo, str)
	return len(results) > 0, results
}
