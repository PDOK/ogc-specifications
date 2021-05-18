package wmts100

import (
	"encoding/xml"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func sp(s string) *string {
	return &s
}

func bp(b bool) *bool {
	return &b
}

// Example based on 7.1.1.3 ServiceMetadata document example from OGC WMTS 1.0.0 spec
var contentsWithLegend = Contents{
	Layer: []Layer{
		{
			Title: "etopo2",
			Abstract: `ETOPO2 - 2 minute Worldwide Bathymetry/Topography
		Data taken from National Geophysical Data Center(NGDC),
		ETOPO2 Global 2' Elevations, September 2001...`,
			WGS84BoundingBox: wsc110.WGS84BoundingBox{LowerCorner: wsc110.Position{-180, -90}, UpperCorner: wsc110.Position{180, 90}},
			Identifier:       "etopo2",
			Metadata:         &Metadata{Href: "http://www.maps.bob/etopo2/ metadata.htm"},
			Style: []Style{
				{
					Title:      sp(`default`),
					Identifier: `default`,
					LegendURL: []*LegendURL{
						{
							Format: "image/png",
							Href:   "http://www.maps.bob/etopo2/legend.png",
						},
					},
					IsDefault: bp(true),
				},
			},
			Format: "image/png",
			TileMatrixSetLink: []TileMatrixSetLink{
				{
					TileMatrixSet: "WholeWorld_CRS_84",
				},
			},
		},
	},
}

func TestBuildStyleWithLegend(t *testing.T) {
	expected := `<Style isDefault="true">
  <ows:Identifier>default</ows:Identifier>
  <ows:Title>default</ows:Title>
  <LegendURL format="image/png" xlink:href="http://www.maps.bob/etopo2/legend.png"></LegendURL>
</Style>`
	output, _ := xml.MarshalIndent(contentsWithLegend.Layer[0].Style, "", "  ")
	if string(output) != expected {
		t.Errorf("test: %d, expected: %s ,\n got: %s", 1, expected, string(output))
	}
}

// Example based on 7.1.1.3 ServiceMetadata document example from OGC WMTS 1.0.0 spec
var contentsWithoutLegend = Contents{
	Layer: []Layer{
		{
			Title: "etopo2",
			Abstract: `ETOPO2 - 2 minute Worldwide Bathymetry/Topography
		Data taken from National Geophysical Data Center(NGDC),
		ETOPO2 Global 2' Elevations, September 2001...`,
			WGS84BoundingBox: wsc110.WGS84BoundingBox{LowerCorner: wsc110.Position{-180, -90}, UpperCorner: wsc110.Position{180, 90}},
			Identifier:       "etopo2",
			Metadata:         &Metadata{Href: "http://www.maps.bob/etopo2/ metadata.htm"},
			Style: []Style{
				{
					Title:      sp(`default`),
					Identifier: `default`,
					IsDefault:  bp(true),
				},
			},
			Format: "image/png",
			TileMatrixSetLink: []TileMatrixSetLink{
				{
					TileMatrixSet: "WholeWorld_CRS_84",
				},
			},
		},
	},
}

func TestBuildStyleWithoutLegend(t *testing.T) {
	expected := `<Style isDefault="true">
  <ows:Identifier>default</ows:Identifier>
  <ows:Title>default</ows:Title>
</Style>`
	output, _ := xml.MarshalIndent(contentsWithoutLegend.Layer[0].Style, "", "  ")
	if string(output) != expected {
		t.Errorf("test: %d, expected: %s ,\n got: %s", 1, expected, string(output))
	}
}
