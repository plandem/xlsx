package xlsx

import (
	"github.com/plandem/xlsx/types"
)

type comments struct {
	sheet          *sheetInfo
}

//newComments creates an object that implements comments functionality
func newComments(sheet *sheetInfo) *comments {
	return &comments{sheet: sheet}
}

func (c *comments) initIfRequired() {
	////attach hyperlinks object if required
	//if h.sheet.ml.Hyperlinks == nil {
	//	var links []*ml.Hyperlink
	//	h.sheet.ml.Hyperlinks = &links
	//}
}

func (c *comments) Add(bounds types.Bounds, comment interface{}) error {
	c.initIfRequired()

	return nil
}

//Remove removes comment info for bounds
func (c *comments) Remove(bounds types.Bounds) {
	c.initIfRequired()
}

