package helpers

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"reflect"
	"unsafe"
)

func FromHyperlink(i *types.HyperlinkInfo) (hyperlink *ml.Hyperlink, styleID format.DirectStyleID, err error) {
	if err = i.Validate(); err != nil {
		return
	}

	v := reflect.ValueOf(i).Elem()

	styleID = format.DirectStyleID(v.FieldByName("styleID").Int())
	hyperlink = (*ml.Hyperlink)(unsafe.Pointer(v.FieldByName("hyperlink").Pointer()))

	return
}
