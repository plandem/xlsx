// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"

	// to link unexported
	_ "unsafe"
)

//go:linkname toRichFont github.com/plandem/xlsx/format/styles.toRichFont
func toRichFont(f *styles.Info) *ml.RichFont

func toRichText(parts ...interface{}) (*ml.StringItem, *styles.Info, error) {
	si := &ml.StringItem{}
	length := 0

	var cellStyles *styles.Info

	if len(parts) > 0 {
		//if last part is format, then use it as cell styles and remove from parts
		if format, ok := parts[len(parts)-1].(*styles.Info); ok {
			cellStyles = format
			parts = parts[:len(parts)-1]
		}

		richText := make([]*ml.RichText, 0, len(parts))
		fontPart := true

		for i, p := range parts {
			switch v := p.(type) {
			case string:
				length += len(v)

				if !fontPart || i == 0 {
					//previous part was string or it's first part - add new block with a string and 'default format'
					richText = append(richText, &ml.RichText{
						Text: primitives.Text(v),
					})
				} else {
					//previous part was a format, so attach a string to prev block
					richText[len(richText)-1].Text = primitives.Text(v)
				}

				fontPart = false

			case *styles.Info:
				if fontPart && i > 0 {
					return nil, nil, errors.New("two styles in row is not allowed")
				}

				richText = append(richText, &ml.RichText{
					Font: toRichFont(v),
				})

				fontPart = true
			}
		}

		if len(richText) == 1 && richText[0].Font == nil {
			si.Text = richText[0].Text
		} else {
			si.RichText = richText
		}
	}

	if length > internal.ExcelCellLimit {
		return nil, nil, fmt.Errorf("text exceeds allowed length for cell value = %d", internal.ExcelCellLimit)
	}

	return si, cellStyles, nil
}

func fromRichText(text *ml.StringItem) (s string) {
	if text == nil {
		return
	}

	s += string(text.Text)

	if text.RichText != nil {
		for _, part := range text.RichText {
			s += string(part.Text)
		}
	}

	return
}
