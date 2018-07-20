package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
)

func TestUpdateLinks(t *testing.T) {
	type Entity struct {
		Attribute types.UpdateLinksType `xml:"attribute,attr"`
	}

	list := map[string] types.UpdateLinksType{
		"userSet": types.UpdateLinksTypeUserSet,
		"never": types.UpdateLinksTypeNever,
		"always": types.UpdateLinksTypeAlways,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T){
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
