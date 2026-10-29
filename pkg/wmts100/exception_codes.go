package wmts100

import (
	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// TileOutOfRange exception
func TileOutOfRange() wsc110.Exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: "TileOutOfRange",
		ExceptionText: "TileRow or TileCol out of rangeName",
		LocatorCode:   "", // TODO parse the right parameter TileRow or TileCol
	}}
}
