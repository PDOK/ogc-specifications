package wms130

import (
	"encoding/xml"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// ----------
// Benchmarks
// ----------

func BenchmarkGetMapToQueryParameters(b *testing.B) {
	gm := GetMapRequest{
		BaseRequest: BaseRequest{
			Version: "1.3.0",
			Attr: utils.XMLAttribute{
				xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/sld"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ogc"}, Value: "http://www.opengis.net/ogc"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "se"}, Value: "http://www.opengis.net/se"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "wms"}, Value: "http://www.opengis.net/wms"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
				xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetMap.xsd"},
			}},
		StyledLayerDescriptor: StyledLayerDescriptor{
			Version: "1.1.0",
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: CRS{Namespace: "EPSG", Code: 4326},
		BoundingBox: BoundingBox{
			Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Output: Output{
			Size:        Size{Width: 1024, Height: 512},
			Format:      "image/jpeg",
			Transparent: bp(false)},
		Exceptions: sp("XML"),
	}
	for i := 0; i < b.N; i++ {
		gm.ToQueryParameters()
	}
}

func BenchmarkGetMapToXML(b *testing.B) {
	gm := GetMapRequest{
		BaseRequest: BaseRequest{
			Version: "1.3.0",
			Attr: utils.XMLAttribute{
				xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/sld"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ogc"}, Value: "http://www.opengis.net/ogc"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "se"}, Value: "http://www.opengis.net/se"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "wms"}, Value: "http://www.opengis.net/wms"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
				xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetMap.xsd"},
			}},
		StyledLayerDescriptor: StyledLayerDescriptor{
			Version: "1.1.0",
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: CRS{Namespace: "EPSG", Code: 4326},
		BoundingBox: BoundingBox{
			Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Output: Output{
			Size:        Size{Width: 1024, Height: 512},
			Format:      "image/jpeg",
			Transparent: bp(false)},
		Exceptions: sp("XML"),
	}
	for i := 0; i < b.N; i++ {
		gm.ToXML()
	}
}

func BenchmarkGetMapParseQueryParameters(b *testing.B) {
	mpv := map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
		LAYERS:      {`Rivers,Roads,Houses`},
		STYLES:      {`CenterLine,CenterLine,Outline`},
		"CRS":       {`EPSG:4326`},
		BBOX:        {`-180.0,-90.0,180.0,90.0`},
		WIDTH:       {`1024`},
		HEIGHT:      {`512`},
		FORMAT:      {`image/jpeg`},
		TRANSPARENT: {`FALSE`},
		EXCEPTIONS:  {`XML`},
		BGCOLOR:     {`0x7F7F7F`},
	}

	for i := 0; i < b.N; i++ {
		gm := GetMapRequest{}
		gm.ParseQueryParameters(mpv)
	}
}

func BenchmarkGetMapParseXML(b *testing.B) {
	doc := []byte(`<GetMap xmlns="http://www.opengis.net/sld" xmlns:gml="http://www.opengis.net/gml" xmlns:ogc="http://www.opengis.net/ogc" xmlns:ows="http://www.opengis.net/ows" 
	xmlns:se="http://www.opengis.net/se" xmlns:wms="http://www.opengis.net/wms" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/sld GetMap.xsd" version="1.3.0">
	<StyledLayerDescriptor version="1.1.0">
		<NamedLayer>
			<se:Name>Rivers</se:Name>
			<NamedStyle>
				<se:Name>CenterLine</se:Name>
			</NamedStyle>
		</NamedLayer>
		<NamedLayer>
			<se:Name>Roads</se:Name>
			<NamedStyle>
				<se:Name>CenterLine</se:Name>
			</NamedStyle>
		</NamedLayer>
		<NamedLayer>
			<se:Name>Houses</se:Name>
			<NamedStyle>
				<se:Name>Outline</se:Name>
			</NamedStyle>
		</NamedLayer>
	</StyledLayerDescriptor>
	<CRS>EPSG:4326</CRS>
	<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326">					
		<ows:LowerCorner>-180.0 -90.0</ows:LowerCorner>
		<ows:UpperCorner>180.0 90.0</ows:UpperCorner>
	</BoundingBox>
	<Output>
		<Size>
			<Width>1024</Width>
			<Height>512</Height>
		</Size>
		<wms:Format>image/jpeg</wms:Format>
		<Transparent>false</Transparent>
	</Output>
	<Exceptions>XML</Exceptions>
</GetMap>`)

	for i := 0; i < b.N; i++ {
		gm := GetMapRequest{}
		gm.ParseXML(doc)
	}
}

// TODO look at the test and the structure
func BenchmarkGetMapValidate(b *testing.B) {
	capabilities := Capabilities{
		WMSCapabilities: WMSCapabilities{
			Request: Request{
				GetMap: RequestType{
					Format:  []string{`image/jpeg`},
					DCPType: DCPType{},
				},
			},
			Layer: []Layer{
				{
					Queryable: ip(1),
					Title:     `Rivers, Roads and Houses`,
					CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
					Layer: []*Layer{
						{
							Queryable: ip(1),
							Name:      sp(`Rivers`),
							Title:     `Rivers`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `CenterLine`,
								},
							},
						},
						{
							Queryable: ip(1),
							Name:      sp(`Roads`),
							Title:     `Roads`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `CenterLine`,
								},
							},
						},
						{
							Queryable: ip(1),
							Name:      sp(`Houses`),
							Title:     `Houses`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `Outline`,
								},
							},
						},
					},
				},
			},
		},
	}

	gm := GetMapRequest{
		BaseRequest: BaseRequest{
			Version: "1.3.0",
			Attr: utils.XMLAttribute{
				xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/sld"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ogc"}, Value: "http://www.opengis.net/ogc"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "se"}, Value: "http://www.opengis.net/se"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "wms"}, Value: "http://www.opengis.net/wms"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
				xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetMap.xsd"},
			}},
		StyledLayerDescriptor: StyledLayerDescriptor{
			Version: "1.1.0",
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: CRS{Namespace: "EPSG", Code: 4326},
		BoundingBox: BoundingBox{
			Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Output: Output{
			Size:        Size{Width: 1024, Height: 512},
			Format:      "image/jpeg",
			Transparent: bp(false)},
		Exceptions: sp("XML"),
	}

	for i := 0; i < b.N; i++ {
		gm.Validate(capabilities)
	}
}

func BenchmarkGetMapParseValidate(b *testing.B) {
	capabilities := Capabilities{
		WMSCapabilities: WMSCapabilities{
			Request: Request{
				GetMap: RequestType{
					Format:  []string{`image/jpeg`},
					DCPType: DCPType{},
				},
			},
			Layer: []Layer{
				{
					Queryable: ip(1),
					Title:     `Rivers, Roads and Houses`,
					CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
					Layer: []*Layer{
						{
							Queryable: ip(1),
							Name:      sp(`Rivers`),
							Title:     `Rivers`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `CenterLine`,
								},
							},
						},
						{
							Queryable: ip(1),
							Name:      sp(`Roads`),
							Title:     `Roads`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `CenterLine`,
								},
							},
						},
						{
							Queryable: ip(1),
							Name:      sp(`Houses`),
							Title:     `Houses`,
							CRS:       []CRS{{Code: 4326, Namespace: `EPSG`}},
							Style: []*Style{
								{
									Name: `Outline`,
								},
							},
						},
					},
				},
			},
		},
	}

	var gm = GetMapRequest{
		BaseRequest: BaseRequest{
			Version: "1.3.0",
			Attr: utils.XMLAttribute{
				xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/sld"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ogc"}, Value: "http://www.opengis.net/ogc"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "se"}, Value: "http://www.opengis.net/se"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "wms"}, Value: "http://www.opengis.net/wms"},
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
				xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetMap.xsd"},
			}},
		StyledLayerDescriptor: StyledLayerDescriptor{
			Version: "1.1.0",
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: CRS{Namespace: "EPSG", Code: 4326},
		BoundingBox: BoundingBox{
			Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Output: Output{
			Size:        Size{Width: 1024, Height: 512},
			Format:      "image/jpeg",
			Transparent: bp(false)},
		Exceptions: sp("XML"),
	}

	for i := 0; i < b.N; i++ {
		gm.Validate(capabilities)
	}
}
