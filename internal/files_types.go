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

	ContentTypeWorkbook      ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	ContentTypeSharedStrings ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	ContentTypeWorksheet     ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	ContentTypeStyles        ml.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
)
