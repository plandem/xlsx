package primitives_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoundsList(t *testing.T) {
	type Entity struct {
		Attribute primitives.BoundsList `xml:"attribute,attr"`
	}

	//empty
	entity := Entity{Attribute: primitives.BoundsList{}}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	//encode
	b := primitives.BoundsListFromRefs("A1:B2", "C3:F10", "Z1")
	require.Equal(t, primitives.RefList("A1:B2 C3:F10 Z1"), b.ToRefList())
	entity = Entity{Attribute: b}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="A1:B2 C3:F10 Z1"></Entity>`, string(encoded))

	//decode
	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)

	//methods
	require.Equal(t, true, primitives.BoundsList{}.IsEmpty())
	require.Equal(t, primitives.RefList("A1:B2 C3:F10 Z1"), decoded.Attribute.ToRefList())
	require.Equal(t, false, decoded.Attribute.IsEmpty())

	decoded.Attribute.Add("X11")
	require.Equal(t, primitives.RefList("A1:B2 C3:F10 Z1 X11"), decoded.Attribute.ToRefList())
	require.Equal(t, "A1:B2 C3:F10 Z1 X11", decoded.Attribute.String())
}
