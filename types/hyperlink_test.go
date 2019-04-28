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

	link.Set(
		Hyperlink.ToUrl("http://google.com#anchor"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID:      "http://google.com",
			Location: "anchor",
		},
		linkType: hyperlinkTypeWeb,
	}, link)
	require.Nil(t, link.Validate())

	//loo large
	link.Set(
		Hyperlink.ToUrl(fmt.Sprintf("http://g%sgle.com", strings.Repeat("o", internal.UrlLimit))),
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

	//XLSX file
	link = NewHyperlink(
		Hyperlink.ToFile(`c:\temp\foo.xlsx`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `file:///c:\temp\foo.xlsx`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.NotNil(t, link.Validate())

	//XLS file
	link = NewHyperlink(
		Hyperlink.ToFile(`c:\temp\foo.xls`),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			RID: `file:///c:\temp\foo.xls`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.NotNil(t, link.Validate())

	//loo large
	link.Set(
		Hyperlink.ToFile(fmt.Sprintf("c:\temp\f%s.doc", strings.Repeat("o", internal.FilePathLimit))),
	)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkOption_ToRef(t *testing.T) {
	//#A1 - without sheet
	link := NewHyperlink(
		Hyperlink.ToRef(`A1`, ""),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#A1`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet1'!A1 - with sheet without space in name
	link = NewHyperlink(
		Hyperlink.ToRef(`A1`, "Sheet1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#'Sheet1'!A1`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet 1'!A2' - with sheet with space in name
	link = NewHyperlink(
		Hyperlink.ToRef(`A1`, "Sheet 1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#'Sheet 1'!A1`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.Nil(t, link.Validate())

	//#'Sheet\'1'!A1 - with sheet with single quote in name
	link = NewHyperlink(
		Hyperlink.ToRef(`A1`, "Sheet'1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: `#'Sheet\'1'!A1`,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.Nil(t, link.Validate())

	//BAD - with sheet, but without ref
	link = NewHyperlink(
		Hyperlink.ToRef(``, "Sheet'1"),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: ``,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.NotNil(t, link.Validate())

	//BAD - without sheet and without ref
	link = NewHyperlink(
		Hyperlink.ToRef(``, ""),
	)

	require.IsType(t, &HyperlinkInfo{}, link)
	require.Equal(t, &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{
			Location: ``,
		},
		linkType: hyperlinkTypeWorkbook,
	}, link)
	require.NotNil(t, link.Validate())
}

func TestHyperlinkInfo_Formatting(t *testing.T) {
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
