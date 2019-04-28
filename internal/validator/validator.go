package validator

import (
	"github.com/plandem/xlsx/internal"
	"net/url"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// IsFilePath check is a str is Win or Unix file
func IsFilePath(str string) bool {
	if regWinPath.MatchString(str) {
		//check windows path limit see:
		if len(str[3:]) > internal.FilePathLimit {
			return false
		}

		return true
	} else if regUnixPath.MatchString(str) {
		return true
	}

	return false
}

// IsURL check if the str is an URL
func IsURL(str string) bool {
	if str == "" || utf8.RuneCountInString(str) >= internal.UrlLimit || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}

	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}

	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}

	if strings.HasPrefix(u.Host, ".") {
		return false
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return regURL.MatchString(str)
}

// IsEmail check if the str is an email
func IsEmail(str string) bool {
	return regEmail.MatchString(str)
}

// IsExcelFile check if str is path to excel file
func IsExcelFile(str string) bool {
	if !IsFilePath(str)	{
		return false
	}

	if ext := filepath.Ext(str); ext == ".xlsx" || ext == ".xls" {
		return true
	}

	return false
}

// IsMailTo check if the str is an mailto format
func IsMailTo(str string) bool {
	return IsURL(str) && strings.HasPrefix(str, "mailto:")
}
