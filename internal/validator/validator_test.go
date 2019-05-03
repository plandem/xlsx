package validator_test

import (
	"github.com/plandem/xlsx/internal/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUrl(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"https", false},
		{"https://", false},
		{"/absolute-path", false},
		{"./relative-path", false},
		{"testing-path", false},
		{"alskjff#?asf//dfas", false},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"http://foobar.org/", true},
		{"http://foobar.ORG", true},
		{"http://foobar.org:8080/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://foobar.com/?q=%2F", true},
		{"http://localhost", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://user:pass@foo_bar_bar.bar_foo.com", true},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", false},
		{"http://foobar.com/#baz=qux", true},
		{"http://foo bar.org", false},
		{"http://foo.bar.org", true},
		{"http://www.foo.bar.org", true},
		{"http:::/not.valid/a//a??a?b=&&c#hi", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validator.IsURL(test.param), "IsURL(%q) should be %v", test.param, test.expected)
	}
}

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

func TestIsMailTo(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
		result map[string]string
	}{
		{"", false, map[string]string{}},
		{"spam@spam.it", false, map[string]string{}},
		{"mailto:spam@spam.it?", false, map[string]string{}},
		{"mailto:spam@spam.it?spam", false, map[string]string{}},
		{"mailto:spam@spam.it?spam=", false, map[string]string{}},
		{"mailto:spam@spam.it?subject", false, map[string]string{}},
		{"mailto:spam@spam.it", true, map[string]string{"email": "spam@spam.it", "subject":""}},
		{"mailto:spam@spam.it?subject=", true, map[string]string{"email": "spam@spam.it", "subject":""}},
		{"mailto:spam@spam.it?subject=the_spam", true, map[string]string{"email": "spam@spam.it", "subject":"the_spam"}},
	}

	for _, test := range tests {
		r1, r2 := validator.IsMailTo(test.param)
		assert.Equal(t, test.expected, r1, "IsMailTo(%q) should be %v", test.param, test.expected)
		assert.Equal(t, test.result, r2, "IsMailTo(%q) should be %v", test.param, test.result)
	}
}