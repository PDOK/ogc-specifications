package main

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/wfs200"
	"github.com/pdok/ogc-specifications/pkg/wms130"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func main() {
	wmsReport := wms130.Exceptions{wms130.MissingParameterValue(wms130.Version), wms130.OperationNotSupported(`Unknown Operation`)}

	wfsReport := wsc110.Exceptions{wsc110.MissingParameterValue(wfs200.SERVICE), wfs200.FeaturesNotLocked()}

	fmt.Println(string(wmsReport.ToReport().ToBytes()))

	fmt.Println(string(wfsReport.ToReport(wfs200.Version).ToBytes()))
}
