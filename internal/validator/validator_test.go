package validator_test

import (
	"github.com/plandem/xlsx/internal/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"spam@spam.it", true},
		{"s@s.s", true},
		{"spam@spam.com.it", true},
		{"spam+spam@spam.it", true},
		{"spam@spam.spam", true},
		{"spam@spam.spam..spam", false},
		{"spam@spam.it?spam", false},
		{"spam@spam.it?spam=spam", false},
		{"spam@spam.spam.spam", true},
		{"spam@spam.中中中", true},
		{"spam@", false},
		{"spam.com", false},
		{"@spam.com", false},
		{"spam|spam@s中pam.com", true},
		{"spam@s中pam.com", true},
		{"spam.s中pam@spam.com", true},
		{"SpAm.sPaM@spam.iT", true},
		{"SPAM.SPAM@SPAM.SPAM.IT", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validator.IsEmail(test.param), "IsEmail(%q) should be %v", test.param, test.expected)
	}
}

func TestIsUrl(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"http://foobar.org/", true},
		{"http://foobar.ORG", true},
		{"http://foobar.org:8080/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://foobar.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://user:pass@foo_bar_bar.bar_foo.com", true},
		{".com", false},
		{"rtmp://foobar.com", false},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foo bar.org", false},
		{"http://foo.bar.org", true},
		{"http://www.foo.bar.org", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validator.IsURL(test.param), "IsURL(%q) should be %v", test.param, test.expected)
	}
}

func TestIsMailTo(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"spam@spam.it", false},
		{"mailto:spam@spam.it?", true},
		{"mailto:spam@spam.it?spam", true},
		{"mailto:spam@spam.it?spam=", true},
		{"mailto:spam@spam.it?subject", true},
		{"mailto:spam@spam.it?subject=", true},
		{"mailto:spam@spam.it?subject=spam", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validator.IsMailTo(test.param), "IsMailTo(%q) should be %v", test.param, test.expected)
	}
}