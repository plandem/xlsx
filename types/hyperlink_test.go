package types

import (
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHyperlinkOption_ToMail(t *testing.T) {
	//unknown type
	link := NewHyperlink()
	require.NotNil(t, link.Validate())

	link.Set(
		Hyperlink.ToMail("spam@spam.it", "My subject"),
	)
	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it?subject=My subject",
		},
		linkType: hyperlinkTypeEmail,
	}, link)
	require.Nil(t, link.Validate())

	link.Set(
		Hyperlink.ToMail("spam@spam.it", ""),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it",
		},
		linkType: hyperlinkTypeEmail,
	}, link)
	require.Nil(t, link.Validate())

	//loo large
	link.Set(
		Hyperlink.ToMail(fmt.Sprintf("spa%sm@spam.it", strings.Repeat("a", internal.UrlLimit)), ""),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToUrl(t *testing.T) {
	link := NewHyperlink(
		Hyperlink.ToUrl("http://google.com"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: "http://google.com",
		},
		linkType: hyperlinkTypeWeb,
	}, link)
	require.Nil(t, link.Validate())

	//url with pound sign(#)
	link.Set(
		Hyperlink.ToUrl("http://google.com/#anchor"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID:      "http://google.com/%23anchor",
		},
		linkType: hyperlinkTypeWeb,
	}, link)
	require.Nil(t, link.Validate())

	//too large
	link.Set(
		Hyperlink.ToUrl(fmt.Sprintf("http://g%sgle.com", strings.Repeat("o", internal.UrlLimit))),
	)
	require.NotNil(t, link.Validate())

	//too short
	link.Set(
		Hyperlink.ToUrl("a.b"),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToFile(t *testing.T) {
	//DOS style
	link := NewHyperlink(
		Hyperlink.ToFile(`c:\temp\foo.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `file:///c:\temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//UNIX style
	link.Set(
		Hyperlink.ToFile(`/Users/andrey/foo.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `file:///\Users\andrey\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//UNC share
	link.Set(
		Hyperlink.ToFile(`\\NET\share\foo.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `file:///\\NET\share\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//relative DOS file
	link.Set(
		Hyperlink.ToFile(`.\temp\foo.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//relative UNIX file
	link.Set(
		Hyperlink.ToFile(`./temp/foo.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//filename with pound sign(#)
	link.Set(
		Hyperlink.ToFile(`./temp/foo#bar.txt`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `temp\foo%23bar.txt`,
		},
		linkType: hyperlinkTypeFile,
	}, link)
	require.Nil(t, link.Validate())

	//too large
	link.Set(
		Hyperlink.ToFile(fmt.Sprintf("c:\temp\f%s.doc", strings.Repeat("o", internal.FilePathLimit))),
	)
	require.NotNil(t, link.Validate())

	//too short
	link.Set(
		Hyperlink.ToFile("a.b"),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToRef(t *testing.T) {
	//#A1 - without sheet
	link := NewHyperlink(
		Hyperlink.ToRef("A1", ""),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "#A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet1'!A1 - with sheet without space in name
	link = NewHyperlink(
		Hyperlink.ToRef("A1", "Sheet1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "#'Sheet1'!A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet 1'!A2' - with sheet with space in name
	link = NewHyperlink(
		Hyperlink.ToRef("A1", "Sheet 1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "#'Sheet 1'!A1",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet\'1'!A1 - with sheet with single quote in name
	link = NewHyperlink(
		Hyperlink.ToRef("A1", "Sheet'1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#'Sheet\'1'!A1`,
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//BAD - with sheet, but without ref
	link = NewHyperlink(
		Hyperlink.ToRef("", "Sheet'1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.NotNil(t, link.Validate())

	//BAD - without sheet and without ref
	link = NewHyperlink(
		Hyperlink.ToRef("", ""),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_Formatting(t *testing.T) {
	link := NewHyperlink(
		Hyperlink.ToMail("spam@spam.it", "My subject"),
	)
	require.Equal(t, format.DefaultDirectStyle, link.Formatting())

	link.Set(Hyperlink.Formatting(1))
	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: "mailto:spam@spam.it?subject=My subject",
		},
		styleID:  1,
		linkType: hyperlinkTypeEmail,
	}, link)

	require.Equal(t, format.DirectStyleID(1), link.Formatting())
}

func TestHyperlinkOption_ToBookmark(t *testing.T) {
	//#Bookmark - without space
	link := NewHyperlink(
		Hyperlink.ToBookmark("Bookmark"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "#'Bookmark'",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Bookmark Name' - with space in name
	link = NewHyperlink(
		Hyperlink.ToBookmark("Bookmark Name"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: "#'Bookmark Name'",
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())

	//#'Bookmark\'1' - with single quote in name
	link = NewHyperlink(
		Hyperlink.ToBookmark("Bookmark'1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#'Bookmark\'1'`,
		},
		linkType: hyperlinkTypeUnknown,
	}, link)
	require.Nil(t, link.Validate())
}

func TestHyperlinkInfo_String(t *testing.T) {
	//#'Bookmark' - bookmark without space
	link := NewHyperlink(
		Hyperlink.ToBookmark("Bookmark"),
	)
	require.Equal(t, "#'Bookmark'", link.String())

	//#'Bookmark Name' - bookmark with space
	link = NewHyperlink(
		Hyperlink.ToBookmark("Bookmark Name"),
	)
	require.Equal(t, "#'Bookmark Name'", link.String())

	//#A1 - ref without sheet
	link = NewHyperlink(
		Hyperlink.ToRef("A1", ""),
	)
	require.Equal(t, "#A1", link.String())

	//#'Sheet Name'!A1 - ref with sheet
	link = NewHyperlink(
		Hyperlink.ToRef("A1", "Sheet Name"),
	)
	require.Equal(t, "#'Sheet Name'!A1", link.String())

	//mailto:spam@spam.it - without subject
	link = NewHyperlink(
		Hyperlink.ToMail("spam@spam.it", ""),
	)
	require.Equal(t, "mailto:spam@spam.it", link.String())


	//mailto:spam@spam.it?subject=topic - with subject
	link = NewHyperlink(
		Hyperlink.ToMail("spam@spam.it", "topic"),
	)
	require.Equal(t, "mailto:spam@spam.it?subject=topic", link.String())

	//http://google.com - without hash
	link = NewHyperlink(
		Hyperlink.ToUrl("http://google.com"),
	)
	require.Equal(t, "http://google.com", link.String())

	//http://google.com/%23hash - with hash
	link = NewHyperlink(
		Hyperlink.ToUrl("http://google.com/#hash"),
	)
	require.Equal(t, "http://google.com/%23hash", link.String())

	//http://google.com/%23hash#'bookmark' - with hash and bookmark
	link = NewHyperlink(
		Hyperlink.ToUrl("http://google.com/#hash"),
		Hyperlink.ToBookmark("bookmark"),
	)
	require.Equal(t, "http://google.com/%23hash#'bookmark'", link.String())

	//file://Users\abc\test.xls - without hash
	link = NewHyperlink(
		Hyperlink.ToFile("/Users/abc/test.xlsx"),
	)
	require.Equal(t, `file:///\Users\abc\test.xlsx`, link.String())

	//file://Users\abc\test%23hash.xls - with hash
	link = NewHyperlink(
		Hyperlink.ToFile("/Users/abc/test#hash.xlsx"),
	)
	require.Equal(t, `file:///\Users\abc\test%23hash.xlsx`, link.String())

	//file://Users\abc\test%23hash.xls#'bookmark' - with hash and bookmark
	link = NewHyperlink(
		Hyperlink.ToFile("/Users/abc/test#hash.xlsx"),
		Hyperlink.ToBookmark("bookmark"),
	)
	require.Equal(t, `file:///\Users\abc\test%23hash.xlsx#'bookmark'`, link.String())
}
