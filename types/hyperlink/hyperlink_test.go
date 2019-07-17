// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hyperlink

import (
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHyperlinkOption_ToMail(t *testing.T) {
	//unknown type
	link := New()
	require.NotNil(t, link.Validate())

	link.Set(
		ToMail("spam@spam.it", "My subject"),
	)
	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it?subject=My subject",
		},
		linkType: hyperlinkTypeEmail,
	}, link)
	require.Nil(t, link.Validate())

	link.Set(
		ToMail("spam@spam.it", ""),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it",
		},
		linkType: hyperlinkTypeEmail,
	}, link)
	require.Nil(t, link.Validate())

	//loo large
	link.Set(
		ToMail(fmt.Sprintf("spa%sm@spam.it", strings.Repeat("a", internal.UrlLimit)), ""),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToUrl(t *testing.T) {
	link := New(
		ToUrl("http://google.com"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "http://google.com",
		},
		linkType: hyperlinkTypeWeb,
	}, link)
	require.Nil(t, link.Validate())

	//url with pound sign(#)
	link.Set(
		ToUrl("http://google.com/#anchor"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "http://google.com/%23anchor",
		},
		linkType: hyperlinkTypeWeb,
	}, link)
	require.Nil(t, link.Validate())

	//too large
	link.Set(
		ToUrl(fmt.Sprintf("http://g%sgle.com", strings.Repeat("o", internal.UrlLimit))),
	)
	require.NotNil(t, link.Validate())

	//too short
	link.Set(
		ToUrl("a.b"),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToFile(t *testing.T) {
	//DOS style
	link := New(
		ToFile(`c:\temp\foo.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `file:///c:\temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//UNIX style
	link.Set(
		ToFile(`/Users/andrey/foo.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `\Users\andrey\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//UNC share
	link.Set(
		ToFile(`\\NET\share\foo.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `file:///\\NET\share\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//relative DOS file
	link.Set(
		ToFile(`.\temp\foo.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//relative UNIX file
	link.Set(
		ToFile(`./temp/foo.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//filename with pound sign(#)
	link.Set(
		ToFile(`temp\foo#bar.txt`),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo%23bar.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//too large
	link.Set(
		ToFile(fmt.Sprintf(`c:\temp\f%s.doc`, strings.Repeat("o", internal.FilePathLimit))),
	)
	require.NotNil(t, link.Validate())

	//too short
	link.Set(
		ToFile("a.b"),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToRef(t *testing.T) {
	//#A1 - without sheet
	link := New(
		ToRef("A1", ""),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet1'!A1 - with sheet without space in name
	link = New(
		ToRef("A1", "Sheet1"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "'Sheet1'!A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet 1'!A2' - with sheet with space in name
	link = New(
		ToRef("A1", "Sheet 1"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "'Sheet 1'!A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet\'1'!A1 - with sheet with single quote in name
	link = New(
		ToRef("A1", "Sheet'1"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: `'Sheet\'1'!A1`,
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//BAD - with sheet, but without ref
	link = New(
		ToRef("", "Sheet'1"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.NotNil(t, link.Validate())

	//BAD - without sheet and without ref
	link = New(
		ToRef("", ""),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_Formatting(t *testing.T) {
	link := New(
		ToMail("spam@spam.it", "My subject"),
	)
	require.Equal(t, nil, link.format)

	link.Set(Styles(styles.DirectStyleID(1)))
	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it?subject=My subject",
		},
		format:   styles.DirectStyleID(1),
		linkType: hyperlinkTypeEmail,
	}, link)

	link.Set(Styles(styles.New(styles.Font.Bold)))
	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it?subject=My subject",
		},
		format:   styles.New(styles.Font.Bold),
		linkType: hyperlinkTypeEmail,
	}, link)
}

func TestHyperlinkOption_ToBookmark(t *testing.T) {
	//#Bookmark - without space
	link := New(
		ToBookmark("Bookmark"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "'Bookmark'",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Bookmark Name' - with space in name
	link = New(
		ToBookmark("Bookmark Name"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: "'Bookmark Name'",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Bookmark\'1' - with single quote in name
	link = New(
		ToBookmark("Bookmark'1"),
	)

	require.IsType(t, &Info{}, link)
	require.Equal(t, &Info{
		hyperlink: &ml.Hyperlink{
			Location: `'Bookmark\'1'`,
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())
}

func TestHyperlinkInfo_String(t *testing.T) {
	//#'Bookmark' - bookmark without space
	link := New(
		ToBookmark(`Bookmark`),
	)
	require.Equal(t, `#'Bookmark'`, link.String())

	//#'Bookmark Name' - bookmark with space
	link = New(
		ToBookmark(`Bookmark Name`),
	)
	require.Equal(t, `#'Bookmark Name'`, link.String())

	//#A1 - ref without sheet
	link = New(
		ToRef(`A1`, ``),
	)
	require.Equal(t, `#A1`, link.String())

	//#'Sheet Name'!A1 - ref with sheet
	link = New(
		ToRef(`A1`, `Sheet Name`),
	)
	require.Equal(t, `#'Sheet Name'!A1`, link.String())

	//mailto:spam@spam.it - without subject
	link = New(
		ToMail(`spam@spam.it`, ``),
	)
	require.Equal(t, `mailto:spam@spam.it`, link.String())

	//mailto:spam@spam.it?subject=topic - with subject
	link = New(
		ToMail(`spam@spam.it`, `topic`),
	)
	require.Equal(t, `mailto:spam@spam.it?subject=topic`, link.String())

	//http://google.com - without hash
	link = New(
		ToUrl(`http://google.com`),
	)
	require.Equal(t, `http://google.com`, link.String())

	//http://google.com/%23hash - with hash
	link = New(
		ToUrl(`http://google.com/#hash`),
	)
	require.Equal(t, `http://google.com/%23hash`, link.String())

	//http://google.com/file.xlsx#bookmark' - with bookmark
	link = New(
		ToUrl(`http://google.com/file.xlsx`),
		ToBookmark(`bookmark`),
	)
	require.Equal(t, `http://google.com/file.xlsx#'bookmark'`, link.String())

	//http://google.com/%23hash#'bookmark' - with hash and bookmark
	link = New(
		ToUrl(`http://google.com/#hash`),
		ToBookmark(`bookmark`),
	)
	require.Equal(t, `http://google.com/%23hash#'bookmark'`, link.String())

	//\Users\abc\test.xlsx - without hash
	link = New(
		ToFile(`/Users/abc/test.xlsx`),
	)
	require.Equal(t, `\Users\abc\test.xlsx`, link.String())

	//\Users\abc\test%23hash.xlsx - with hash
	link = New(
		ToFile(`/Users/abc/test#hash.xlsx`),
	)
	require.Equal(t, `\Users\abc\test%23hash.xlsx`, link.String())

	//\Users\abc\test%23hash.xlsx#'bookmark' - with hash and bookmark
	link = New(
		ToFile(`/Users/abc/test#hash.xlsx`),
		ToBookmark(`bookmark`),
	)
	require.Equal(t, `\Users\abc\test%23hash.xlsx#'bookmark'`, link.String())

	//file:///C:\Users\abc\test.xlsx - without hash
	link = New(
		ToFile(`C:\Users\abc\test.xlsx`),
	)
	require.Equal(t, `file:///C:\Users\abc\test.xlsx`, link.String())

	//file:///C:\Users\abc\test%23hash.xlsx - with hash
	link = New(
		ToFile(`C:\Users\abc\test#hash.xlsx`),
	)
	require.Equal(t, `file:///C:\Users\abc\test%23hash.xlsx`, link.String())

	//file:///C:\Users\abc\test%23hash.xlsx#'bookmark' - with hash and bookmark
	link = New(
		ToFile(`C:\Users\abc\test#hash.xlsx`),
		ToBookmark(`bookmark`),
	)
	require.Equal(t, `file:///C:\Users\abc\test%23hash.xlsx#'bookmark'`, link.String())

	//file:///\\Users\abc\test.xlsx - without hash
	link = New(
		ToFile(`\\Users\abc\test.xlsx`),
	)
	require.Equal(t, `file:///\\Users\abc\test.xlsx`, link.String())

	//file:///\\Users\abc\test%23hash.xlsx - with hash
	link = New(
		ToFile(`\\Users\abc\test#hash.xlsx`),
	)
	require.Equal(t, `file:///\\Users\abc\test%23hash.xlsx`, link.String())

	//file:///\\Users\abc\test%23hash.xlsx#'bookmark' - with hash and bookmark
	link = New(
		ToFile(`\\Users\abc\test#hash.xlsx`),
		ToBookmark(`bookmark`),
	)
	require.Equal(t, `file:///\\Users\abc\test%23hash.xlsx#'bookmark'`, link.String())
}

func TestHyperlinkOption_ToTarget(t *testing.T) {
	//t.Parallel()

	var tests = []struct {
		target   string
		expected *Info
	}{
		{`#A1`, &Info{
			hyperlink: &ml.Hyperlink{
				Location: `A1`,
			},
			linkType: hyperlinkTypeUnknown,
		}},
		{`#SheetName!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				Location: `SheetName!A1`,
			},
			linkType: hyperlinkTypeUnknown,
		}},
		{`#'Sheet Name'!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				Location: `'Sheet Name'!A1`,
			},
			linkType: hyperlinkTypeUnknown,
		}},
		{`D:\Folder\File.docx`, &Info{
			hyperlink: &ml.Hyperlink{
				RID: `file:///D:\Folder\File.docx`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`D:\Folder\File.docx#Bookmark`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.docx`,
				Location: `Bookmark`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`D:\Folder\File.xlsx#SheetName!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.xlsx`,
				Location: `SheetName!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`D:\Folder\File.xlsx#'Sheet Name'!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.xlsx`,
				Location: `'Sheet Name'!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[D:\Folder\File.docx]`, &Info{
			hyperlink: &ml.Hyperlink{
				RID: `file:///D:\Folder\File.docx`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[D:\Folder\File.docx]Bookmark`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.docx`,
				Location: `Bookmark`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[D:\Folder\File.xlsx]SheetName!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.xlsx`,
				Location: `SheetName!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[D:\Folder\File.xlsx]'Sheet Name'!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///D:\Folder\File.xlsx`,
				Location: `'Sheet Name'!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`\\SERVER\Folder\File.doc`, &Info{
			hyperlink: &ml.Hyperlink{
				RID: `file:///\\SERVER\Folder\File.doc`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`\\SERVER\Folder\File.xlsx#SheetName!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///\\SERVER\Folder\File.xlsx`,
				Location: `SheetName!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},

		{`\\SERVER\Folder\File.xlsx#'Sheet Name'!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///\\SERVER\Folder\File.xlsx`,
				Location: `'Sheet Name'!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[\\SERVER\Folder\File.xlsx]SheetName!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///\\SERVER\Folder\File.xlsx`,
				Location: `SheetName!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`[\\SERVER\Folder\File.xlsx]'Sheet Name'!A1`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `file:///\\SERVER\Folder\File.xlsx`,
				Location: `'Sheet Name'!A1`,
			},
			linkType: hyperlinkTypeFile,
		}},
		{`https://www.spam.it`, &Info{
			hyperlink: &ml.Hyperlink{
				RID: `https://www.spam.it`,
			},
			linkType: hyperlinkTypeWeb,
		}},
		{`https://www.spam.it/#bookmark`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `https://www.spam.it`,
				Location: `bookmark`,
			},
			linkType: hyperlinkTypeWeb,
		}},
		{`[https://www.spam.it]bookmark`, &Info{
			hyperlink: &ml.Hyperlink{
				RID:      `https://www.spam.it`,
				Location: `bookmark`,
			},
			linkType: hyperlinkTypeWeb,
		}},
		{"spam@spam.it", &Info{
			hyperlink: &ml.Hyperlink{
				RID: `mailto:spam@spam.it`,
			},
			linkType: hyperlinkTypeEmail,
		}},
		{`mailto:spam@spam.it?subject=topic`, &Info{
			hyperlink: &ml.Hyperlink{
				RID: `mailto:spam@spam.it?subject=topic`,
			},
			linkType: hyperlinkTypeEmail,
		}},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, New(ToTarget(test.target)), "ToTarget(%q) should be %v", test.target, test.expected)
	}
}
