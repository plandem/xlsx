package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//FontCharsetType is alias of original primitives.FontCharsetType type because more can be used by system and to:
// 1) make it public
// 2) forbid usage of integers directly
type FontCharsetType = primitives.FontCharsetType

//List of all possible values for FontCharsetType
const (
	FontCharsetANSI        FontCharsetType = 0
	FontCharsetDEFAULT     FontCharsetType = 1
	FontCharsetSYMBOL      FontCharsetType = 2
	FontCharsetMAC         FontCharsetType = 77
	FontCharsetSHIFTJIS    FontCharsetType = 128
	FontCharsetHANGUL      FontCharsetType = 129
	FontCharsetJOHAB       FontCharsetType = 130
	FontCharsetGB2312      FontCharsetType = 134
	FontCharsetCHINESEBIG5 FontCharsetType = 136
	FontCharsetGREEK       FontCharsetType = 161
	FontCharsetTURKISH     FontCharsetType = 162
	FontCharsetVIETNAMESE  FontCharsetType = 163
	FontCharsetHEBREW      FontCharsetType = 177
	FontCharsetARABIC      FontCharsetType = 178
	FontCharsetBALTIC      FontCharsetType = 186
	FontCharsetRUSSIAN     FontCharsetType = 204
	FontCharsetTHAI        FontCharsetType = 222
	FontCharsetEASTEUROPE  FontCharsetType = 238
	FontCharsetOEM         FontCharsetType = 255
)
