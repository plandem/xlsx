package internal

import (
	"github.com/plandem/ooxml/ml"
)

//List of all supported RelationType and ContentType
const (
	RelationTypeWorkbook      ml.RelationType = ml.NamespaceRelationships + "/officeDocument"
	RelationTypeSharedStrings ml.RelationType = ml.NamespaceRelationships + "/sharedStrings"
	RelationTypeWorksheet     ml.RelationType = ml.NamespaceRelationships + "/worksheet"
	RelationTypeStyles        ml.RelationType = ml.NamespaceRelationships + "/styles"
	RelationTypeHyperlink     ml.RelationType = ml.NamespaceRelationships + "/hyperlink"
	RelationTypeComments      ml.RelationType = ml.NamespaceRelationships + "/comments"
	RelationTypeVmlDrawing    ml.RelationType = ml.NamespaceRelationships + "/vmlDrawing"

	ContentTypeWorkbook      ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	ContentTypeSharedStrings ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	ContentTypeWorksheet     ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	ContentTypeStyles        ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	ContentTypeComments      ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"
)
