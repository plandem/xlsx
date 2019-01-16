package ml

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Workbook is a direct mapping of XSD CT_Workbook
type Workbook struct {
	XMLName             ml.Name               `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main workbook"`
	RIDName             ml.RIDName            `xml:",attr"`
	FileVersion         *FileVersion          `xml:"fileVersion,omitempty"`
	FileSharing         *ml.Reserved          `xml:"fileSharing,omitempty"`
	WorkbookPr          *WorkbookPr           `xml:"workbookPr,omitempty"`
	WorkbookProtection  *ml.Reserved          `xml:"workbookProtection,omitempty"`
	BookViews           *[]*BookView          `xml:"bookViews>workbookView"` //we HAVE TO remove 'bookViews' if there is no 'bookView'
	Sheets              []*Sheet              `xml:"sheets>sheet"`
	FunctionGroups      *ml.Reserved          `xml:"functionGroups,omitempty"`
	ExternalReferences  *[]*ExternalReference `xml:"externalReferences>externalReference"` //we HAVE TO remove 'externalReferences' if there is no 'externalReference'
	DefinedNames        *ml.Reserved          `xml:"definedNames,omitempty"`
	CalcPr              *ml.Reserved          `xml:"calcPr,omitempty"`
	OleSize             *ml.Reserved          `xml:"oleSize,omitempty"`
	CustomWorkbookViews *ml.Reserved          `xml:"customWorkbookViews,omitempty"`
	PivotCaches         *ml.Reserved          `xml:"pivotCaches,omitempty"`
	SmartTagPr          *ml.Reserved          `xml:"smartTagPr,omitempty"`
	SmartTagTypes       *ml.Reserved          `xml:"smartTagTypes,omitempty"`
	WebPublishing       *ml.Reserved          `xml:"webPublishing,omitempty"`
	FileRecoveryPr      *ml.Reserved          `xml:"fileRecoveryPr,omitempty"`
	WebPublishObjects   *ml.Reserved          `xml:"webPublishObjects,omitempty"`
	ExtLst              *ml.Reserved          `xml:"extLst,omitempty"`
	Conformance         string                `xml:"conformance,attr,omitempty"`
}

//FileVersion is a direct mapping of XSD CT_FileVersion
type FileVersion struct {
	AppName      string `xml:"appName,attr,omitempty"`
	CodeName     string `xml:"codeName,attr,omitempty"`
	LastEdited   string `xml:"lastEdited,attr,omitempty"`
	LowestEdited string `xml:"lowestEdited,attr,omitempty"`
	RupBuild     string `xml:"rupBuild,attr,omitempty"`
}

//WorkbookPr is a direct mapping of XSD CT_WorkbookPr
type WorkbookPr struct {
	Date1904                   bool                       `xml:"date1904,attr,omitempty"`
	ShowObjects                primitives.ObjectsType     `xml:"showObjects,attr,omitempty"`
	ShowBorderUnselectedTables bool                       `xml:"showBorderUnselectedTables,attr,omitempty"`
	FilterPrivacy              bool                       `xml:"filterPrivacy,attr,omitempty"`
	PromptedSolutions          bool                       `xml:"promptedSolutions,attr,omitempty"`
	ShowInkAnnotation          bool                       `xml:"showInkAnnotation,attr,omitempty"`
	BackupFile                 bool                       `xml:"backupFile,attr,omitempty"`
	SaveExternalLinkValues     bool                       `xml:"saveExternalLinkValues,attr,omitempty"`
	UpdateLinks                primitives.UpdateLinksType `xml:"updateLinks,attr,omitempty"`
	CodeName                   string                     `xml:"codeName,attr,omitempty"`
	HidePivotFieldList         bool                       `xml:"hidePivotFieldList,attr,omitempty"`
	ShowPivotChartFilter       bool                       `xml:"showPivotChartFilter,attr,omitempty"`
	AllowRefreshQuery          bool                       `xml:"allowRefreshQuery,attr,omitempty"`
	PublishItems               bool                       `xml:"publishItems,attr,omitempty"`
	CheckCompatibility         bool                       `xml:"checkCompatibility,attr,omitempty"`
	AutoCompressPictures       bool                       `xml:"autoCompressPictures,attr,omitempty"`
	RefreshAllConnections      bool                       `xml:"refreshAllConnections,attr,omitempty"`
	DefaultThemeVersion        uint                       `xml:"defaultThemeVersion,attr,omitempty"`
}

//BookView is a direct mapping of XSD CT_BookView
type BookView struct {
	ExtLst                 *ml.Reserved              `xml:"extLst,omitempty"`
	Visibility             primitives.VisibilityType `xml:"visibility,attr,omitempty"`
	Minimized              bool                      `xml:"minimized,attr,omitempty"`
	ShowHorizontalScroll   bool                      `xml:"showHorizontalScroll,attr,omitempty"`
	ShowVerticalScroll     bool                      `xml:"showVerticalScroll,attr,omitempty"`
	ShowSheetTabs          bool                      `xml:"showSheetTabs,attr,omitempty"`
	XWindow                int                       `xml:"xWindow,attr,omitempty"`
	YWindow                int                       `xml:"yWindow,attr,omitempty"`
	WindowHeight           uint                      `xml:"windowHeight,attr,omitempty"`
	WindowWidth            uint                      `xml:"windowWidth,attr,omitempty"`
	TabRatio               uint                      `xml:"tabRatio,attr,omitempty"`
	FirstSheet             uint                      `xml:"firstSheet,attr,omitempty"`
	ActiveTab              int                       `xml:"activeTab,attr,omitempty"`
	AutoFilterDateGrouping bool                      `xml:"autoFilterDateGrouping,attr,omitempty"`
}

//Sheet is a direct mapping of XSD CT_Sheet
type Sheet struct {
	Name    string                    `xml:"name,attr"`
	SheetID uint                      `xml:"sheetId,attr"`
	State   primitives.VisibilityType `xml:"state,attr,omitempty"`
	RID     ml.RID                    `xml:"id,attr"`
}

//ExternalReference is a direct mapping of XSD CT_ExternalReference
type ExternalReference struct {
	RID ml.RID `xml:"id,attr,omitempty"`
}
