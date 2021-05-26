package wmts100

import "github.com/pdok/ogc-specifications/pkg/wsc110"

// CannotLockAllFeatures exception
func TileOutOfRange() wsc110.Exception {
	return exception{
		ExceptionCode: "TileOutOfRange",
		ExceptionText: "TileRow or TileCol out of rangeName",
		LocatorCode:   "", // TODO parse the right parameter TileRow or TileCol
	}
}
