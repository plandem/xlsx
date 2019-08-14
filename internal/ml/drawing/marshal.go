// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/ml"
)

const (
	anchorAbsolute      = "absoluteAnchor"
	anchorOneCellAnchor = "oneCellAnchor"
	anchorTwoCellAnchor = "twoCellAnchor"
	errorUnknownAnchor  = "unknown type of anchor: %s"
)

//UnmarshalXML unmarshal Anchor
func (a *AnchorList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case anchorAbsolute:
		anchor := &AbsoluteAnchor{}
		if err := d.DecodeElement(anchor, &start); err != nil {
			return err
		}

		a.Add(anchor)
	case anchorOneCellAnchor:
		anchor := &OneCellAnchor{}
		if err := d.DecodeElement(anchor, &start); err != nil {
			return err
		}

		a.Add(anchor)
	case anchorTwoCellAnchor:
		anchor := &TwoCellAnchor{}
		if err := d.DecodeElement(anchor, &start); err != nil {
			return err
		}

		a.Add(anchor)
	default:
		return fmt.Errorf(errorUnknownAnchor, start)
	}

	return nil
}

func (a *AnchorList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error

	for _, anchor := range *a {
		switch at := anchor.(type) {
		case AbsoluteAnchor, *AbsoluteAnchor:
			err = e.EncodeElement(at, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, xml.Name{Local: anchorAbsolute})})
		case OneCellAnchor, *OneCellAnchor:
			err = e.EncodeElement(at, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, xml.Name{Local: anchorOneCellAnchor})})
		case TwoCellAnchor, *TwoCellAnchor:
			err = e.EncodeElement(at, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, xml.Name{Local: anchorTwoCellAnchor})})
		default:
			err = fmt.Errorf(errorUnknownAnchor, at)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

//MarshalXML marshal Drawing
func (w *Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, xml.Name{Local: "wsDr"})
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceDrawingExcel,
		ml.NamespaceDrawing,
	)...)

	return e.EncodeElement(*w, start)
}

func (n *Point) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n *Size) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n *Marker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n Coordinate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(string(n), xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n id) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(n), xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}

func (n *GraphicFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedAttributes.ResolveNamespacePrefixes()
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
}
