// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestFontCharset(t *testing.T) {
	type Element struct {
		Property primitives.FontCharsetType `xml:"property,omitempty"`
	}

	list := map[styles.FontCharsetType]string{
		styles.FontCharsetANSI:        strconv.Itoa(int(styles.FontCharsetANSI)),
		styles.FontCharsetDEFAULT:     strconv.Itoa(int(styles.FontCharsetDEFAULT)),
		styles.FontCharsetSYMBOL:      strconv.Itoa(int(styles.FontCharsetSYMBOL)),
		styles.FontCharsetMAC:         strconv.Itoa(int(styles.FontCharsetMAC)),
		styles.FontCharsetSHIFTJIS:    strconv.Itoa(int(styles.FontCharsetSHIFTJIS)),
		styles.FontCharsetHANGUL:      strconv.Itoa(int(styles.FontCharsetHANGUL)),
		styles.FontCharsetJOHAB:       strconv.Itoa(int(styles.FontCharsetJOHAB)),
		styles.FontCharsetGB2312:      strconv.Itoa(int(styles.FontCharsetGB2312)),
		styles.FontCharsetCHINESEBIG5: strconv.Itoa(int(styles.FontCharsetCHINESEBIG5)),
		styles.FontCharsetGREEK:       strconv.Itoa(int(styles.FontCharsetGREEK)),
		styles.FontCharsetTURKISH:     strconv.Itoa(int(styles.FontCharsetTURKISH)),
		styles.FontCharsetVIETNAMESE:  strconv.Itoa(int(styles.FontCharsetVIETNAMESE)),
		styles.FontCharsetHEBREW:      strconv.Itoa(int(styles.FontCharsetHEBREW)),
		styles.FontCharsetARABIC:      strconv.Itoa(int(styles.FontCharsetARABIC)),
		styles.FontCharsetBALTIC:      strconv.Itoa(int(styles.FontCharsetBALTIC)),
		styles.FontCharsetRUSSIAN:     strconv.Itoa(int(styles.FontCharsetRUSSIAN)),
		styles.FontCharsetTHAI:        strconv.Itoa(int(styles.FontCharsetTHAI)),
		styles.FontCharsetEASTEUROPE:  strconv.Itoa(int(styles.FontCharsetEASTEUROPE)),
		styles.FontCharsetOEM:         strconv.Itoa(int(styles.FontCharsetOEM)),
		styles.FontCharsetType(25):    strconv.Itoa(25),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: primitives.FontCharsetType(v)}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
				require.Equal(tt, `<Element></Element>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property></Element>`, s), string(encoded))
			}

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
