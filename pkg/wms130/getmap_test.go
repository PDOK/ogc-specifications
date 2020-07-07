package wms130

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

func sp(s string) *string {
	return &s
}

func TestGetMapType(t *testing.T) {
	dft := GetMap{}
	if dft.Type() != `GetMap` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetMap`, dft.Type())
	}
}

func TestBuildBoundingBox(t *testing.T) {
	var tests = []struct {
		boundingbox string
		bbox        ows.BoundingBox
	}{
		0: {boundingbox: "0,0,100,100", bbox: ows.BoundingBox{LowerCorner: [2]float64{0, 0}, UpperCorner: [2]float64{100, 100}}},
		1: {boundingbox: "0,0,-100,-100", bbox: ows.BoundingBox{LowerCorner: [2]float64{0, 0}, UpperCorner: [2]float64{-100, -100}}}, // while this isn't correct, this will be 'addressed' in the validation step
		2: {boundingbox: "0,0,100", bbox: ows.BoundingBox{}},
		3: {boundingbox: ",,,", bbox: ows.BoundingBox{}},
		4: {boundingbox: ",,,100", bbox: ows.BoundingBox{}},
		5: {boundingbox: "number,,,100", bbox: ows.BoundingBox{}},
	}

	for k, test := range tests {
		bbox := buildBoundingBox(test.boundingbox)

		if bbox != test.bbox {
			t.Errorf("test: %d, expected: %+v \ngot: %+v", k, test.bbox, bbox)
		}
	}
}

func TestBuildStyledLayerDescriptor(t *testing.T) {
	var tests = []struct {
		layers []string
		styles []string
		sld    StyledLayerDescriptor
		Error  ows.Exception
	}{
		0: {layers: []string{"layer1", "layer2"}, styles: []string{"style1", "style2"}, sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1", NamedStyle: &NamedStyle{Name: "style1"}}, {Name: "layer2", NamedStyle: &NamedStyle{Name: "style2"}}}}},
		1: {layers: []string{"layer1", "layer2"}, styles: []string{"style1", "style2", "style3"}, Error: StyleNotDefined()},
		2: {layers: []string{"layer1", "layer2"}, styles: []string{"style1"}, Error: StyleNotDefined()},
		3: {layers: []string{"layer1", "layer2"}, sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1"}, {Name: "layer2"}}}},
	}

	for k, test := range tests {
		result, err := buildStyledLayerDescriptor(test.layers, test.styles)

		if err != nil {
			if test.Error == nil || err != test.Error {
				t.Errorf("test: %d, expected: %+v \ngot: %+v", k, test.Error, err)
			}
		}

		for s := range result.NamedLayer {
			rl := &result.NamedLayer[s]
			tl := &test.sld.NamedLayer[s]
			if rl.Name != tl.Name {
				t.Errorf("test: %d, expected: %+v \ngot: %+v", k, tl.Name, rl.Name)
			}
			if rl.NamedStyle != nil && tl.NamedStyle != nil {
				if rl.NamedStyle.Name != tl.NamedStyle.Name {
					t.Errorf("test: %d, expected: %+v \ngot: %+v", k, tl.NamedStyle.Name, rl.NamedStyle.Name)
				}
			}
		}
	}
}

func TestParseBodyGetFeature(t *testing.T) {
	var tests = []struct {
		Body     []byte
		Excepted GetMap
		Error    ows.Exception
	}{
		// GetMap schemas.opengis.net/sld/1.1.0/example_getmap.xml example request
		0: {Body: []byte(`<GetMap xmlns="http://www.opengis.net/sld" xmlns:gml="http://www.opengis.net/gml" xmlns:ogc="http://www.opengis.net/ogc" xmlns:ows="http://www.opengis.net/ows" 
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
	</GetMap>`),
			Excepted: GetMap{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
					Attr: ows.XMLAttribute{
						xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/sld"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ogc"}, Value: "http://www.opengis.net/ogc"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "se"}, Value: "http://www.opengis.net/se"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "wms"}, Value: "http://www.opengis.net/wms"},
						xml.Attr{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
						xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetMap.xsd"},
					}},
				GetMapCore: GetMapCore{
					StyledLayerDescriptor: StyledLayerDescriptor{
						Version: "1.1.0",
						NamedLayer: []NamedLayer{
							{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
							{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
							{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
						}},
					CRS: "EPSG:4326",
					BoundingBox: ows.BoundingBox{
						Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
						LowerCorner: [2]float64{-180.0, -90.0},
						UpperCorner: [2]float64{180.0, 90.0},
					},
					Output: Output{
						Size:        Size{Width: 1024, Height: 512},
						Format:      "image/jpeg",
						Transparent: sp("false")},
				},
				Exceptions: sp("XML"),
			},
		},
		1: {Body: []byte(``), Error: ows.MissingParameterValue()},
		2: {Body: []byte(`<UnknownTag/>`), Excepted: GetMap{}},
	}
	for k, n := range tests {
		var gm GetMap
		err := gm.ParseBody(n.Body)
		if err != nil {
			if err.Error() != n.Error.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
			}
		} else {
			compareGetMapObject(gm, n.Excepted, t, k)
		}
	}
}

func TestGetLayerQueryParameter(t *testing.T) {
	var tests = []struct {
		StyledLayerDescriptor StyledLayerDescriptor
		Excepted              string
	}{
		0: {StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers"},
				{Name: "Roads"},
				{Name: "Houses"},
			}}, Excepted: "Rivers,Roads,Houses"},
		1: {StyledLayerDescriptor: StyledLayerDescriptor{}, Excepted: ""},
		2: {StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers"},
			}}, Excepted: "Rivers"},
	}

	for k, n := range tests {
		result := n.StyledLayerDescriptor.getLayerQueryParameter()
		if n.Excepted != result {
			t.Errorf("test Exceptions: %d, expected: %v+ ,\n got: %v+", k, n.Excepted, result)
		}
	}
}

func TestGetStyleQueryParameter(t *testing.T) {
	var tests = []struct {
		StyledLayerDescriptor StyledLayerDescriptor
		Excepted              string
	}{
		0: {StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}}, Excepted: "CenterLine,CenterLine,Outline"},
		1: {StyledLayerDescriptor: StyledLayerDescriptor{}, Excepted: ""},
		2: {StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads"},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}}, Excepted: "CenterLine,,Outline"},
		// 4. sould fail in the validation step
		4: {StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{NamedStyle: &NamedStyle{Name: "Outline"}},
			}}, Excepted: ""},
	}

	for k, n := range tests {
		result := n.StyledLayerDescriptor.getStyleQueryParameter()
		if n.Excepted != result {
			t.Errorf("test Exceptions: %d, expected: %v+ ,\n got: %v+", k, n.Excepted, result)
		}
	}
}

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		Query    url.Values
		Excepted GetMap
		Error    ows.Exception
	}{
		0: {Query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version}}, Excepted: GetMap{XMLName: xml.Name{Local: getmap}, BaseRequest: BaseRequest{Version: Version, Service: Service}}},
		1: {Query: url.Values{}, Excepted: GetMap{}},
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers,Roads,Houses&STYLES=CenterLine,CenterLine,Outline&CRS=EPSG:4326&BBOX=-180.0,-90.0,180.0,90.0&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&TRANSPARENT=FALSE&EXCEPTIONS=XML
		2: {Query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:      {`Rivers,Roads,Houses`},
			STYLES:      {`CenterLine,CenterLine,Outline`},
			CRS:         {`EPSG:4326`},
			BBOX:        {`-180.0,-90.0,180.0,90.0`},
			WIDTH:       {`1024`},
			HEIGHT:      {`512`},
			FORMAT:      {`image/jpeg`},
			TRANSPARENT: {`FALSE`},
			EXCEPTIONS:  {`XML`},
		},
			Excepted: GetMap{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
				},
				GetMapCore: GetMapCore{
					StyledLayerDescriptor: StyledLayerDescriptor{
						NamedLayer: []NamedLayer{
							{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
							{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
							{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
						}},
					CRS: "EPSG:4326",
					BoundingBox: ows.BoundingBox{
						LowerCorner: [2]float64{-180.0, -90.0},
						UpperCorner: [2]float64{180.0, 90.0},
					},
					Output: Output{
						Size:        Size{Width: 1024, Height: 512},
						Format:      "image/jpeg",
						Transparent: sp("false")},
				},
				Exceptions: sp("XML"),
			}},
		3: {Query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			BGCOLOR: {`0x7F7F7F`},
		}, Excepted: GetMap{
			BaseRequest: BaseRequest{
				Version: "1.3.0",
			},
			GetMapCore: GetMapCore{Output: Output{BGcolor: sp(`0x7F7F7F`)}},
		}},
	}
	for k, n := range tests {
		var gm GetMap
		err := gm.ParseQuery(n.Query)
		if err != nil {
			if err.Error() != n.Error.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
			}
		} else {
			compareGetMapObject(gm, n.Excepted, t, k)
		}
	}
}

func TestBuildQuery(t *testing.T) {
	var tests = []struct {
		Object   GetMap
		Excepted url.Values
		Error    ows.Exception
	}{
		0: {Object: GetMap{
			XMLName: xml.Name{Local: "GetMap"},
			BaseRequest: BaseRequest{
				Version: "1.3.0",
				Service: "WMS",
			},
			GetMapCore: GetMapCore{
				StyledLayerDescriptor: StyledLayerDescriptor{
					NamedLayer: []NamedLayer{
						{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
					}},
				CRS: "EPSG:4326",
				BoundingBox: ows.BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Output: Output{
					Size:        Size{Width: 1024, Height: 512},
					Format:      "image/jpeg",
					Transparent: sp("false")},
			},
			Exceptions: sp("XML"),
		}, Excepted: map[string][]string{
			LAYERS:      {`Rivers,Roads,Houses`},
			STYLES:      {`CenterLine,CenterLine,Outline`},
			CRS:         {`EPSG:4326`},
			BBOX:        {`-180.000000,-90.000000,180.000000,90.000000`},
			EXCEPTIONS:  {`XML`},
			FORMAT:      {`image/jpeg`},
			HEIGHT:      {`512`},
			WIDTH:       {`1024`},
			TRANSPARENT: {`false`},
			VERSION:     {`1.3.0`},
			REQUEST:     {`GetMap`},
			SERVICE:     {`WMS`},
		}},
		1: {Object: GetMap{
			GetMapCore: GetMapCore{CRS: "EPSG:4326",
				BoundingBox: ows.BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
			},
		},
			Excepted: map[string][]string{
				LAYERS:  {``},
				STYLES:  {``},
				CRS:     {`EPSG:4326`},
				BBOX:    {`-180.000000,-90.000000,180.000000,90.000000`},
				FORMAT:  {``},
				HEIGHT:  {`0`},
				WIDTH:   {`0`},
				VERSION: {``},
				REQUEST: {``},
				SERVICE: {``},
			}},
	}

	for k, n := range tests {
		url := n.Object.BuildQuery()
		if len(n.Excepted) != len(url) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, n.Excepted, url)
		} else {
			for _, rid := range url {
				found := false
				for _, erid := range n.Excepted {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, n.Excepted, url)
				}
				found = false
			}
		}
	}
}

func TestBuildBody(t *testing.T) {
	var tests = []struct {
		gm     GetMap
		result string
	}{
		0: {gm: GetMap{},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetMap service="" version="">
 <StyledLayerDescriptor version=""></StyledLayerDescriptor>
 <CRS></CRS>
 <BoundingBox>
  <LowerCorner>0.000000 0.000000</LowerCorner>
  <UpperCorner>0.000000 0.000000</UpperCorner>
 </BoundingBox>
 <Output>
  <Size>
   <Width>0</Width>
   <Height>0</Height>
  </Size>
  <Format></Format>
 </Output>
</GetMap>`},
	}

	for k, v := range tests {
		body := v.gm.BuildBody()

		if string(body) != v.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, v.result, string(body))
		}
	}

}

func compareGetMapObject(result, expected GetMap, t *testing.T, k int) {
	if result.BaseRequest.Version != expected.BaseRequest.Version {
		t.Errorf("test Version: %d, expected: %s ,\n got: %s", k, expected.Version, result.Version)
	}

	if len(expected.BaseRequest.Attr) == len(result.BaseRequest.Attr) {
		c := false
		for _, expectedAttr := range expected.BaseRequest.Attr {
			for _, result := range result.BaseRequest.Attr {
				if result.Name.Local == expectedAttr.Name.Local && result.Value == expectedAttr.Value {
					c = true
				}
			}
			if !c {
				t.Errorf("testBaseRequest.Attr : %d, expected: %s ,\n got: %s", k, expected.BaseRequest.Attr, result.BaseRequest.Attr)
			}
			c = false
		}
	} else {
		t.Errorf("test BaseRequest.Attr: %d, expected: %s ,\n got: %s", k, expected.BaseRequest.Attr, result.BaseRequest.Attr)
	}
	if len(expected.StyledLayerDescriptor.NamedLayer) == len(result.StyledLayerDescriptor.NamedLayer) {
		c := false
		for _, expected := range expected.StyledLayerDescriptor.NamedLayer {
			for _, result := range result.StyledLayerDescriptor.NamedLayer {
				if result.Name == expected.Name {
					if *&result.NamedStyle.Name == *&expected.NamedStyle.Name {
						c = true
					}
				}
			}
			if !c {
				t.Errorf("test StyledLayerDescriptor.NamedLayer: %d, expected: %v+ ,\n got: %v+", k, expected, result.StyledLayerDescriptor.NamedLayer)
			}
			c = false
		}
	} else {
		t.Errorf("test StyledLayerDescriptor: %d, expected: %v+ ,\n got: %v+", k, expected.StyledLayerDescriptor, result.StyledLayerDescriptor)
	}
	if expected.CRS != result.CRS {
		t.Errorf("test CRS: %d, expected: %v+ ,\n got: %v+", k, expected.CRS, result.CRS)
	}
	if expected.BoundingBox != result.BoundingBox {
		t.Errorf("test BoundingBox: %d, expected: %v+ ,\n got: %v+", k, expected.BoundingBox, result.BoundingBox)
	}
	if expected.Output.Size != result.Output.Size {
		t.Errorf("test Output.Size: %d, expected: %v+ ,\n got: %v+", k, expected.Output.Size, result.Output.Size)
	}
	if expected.Exceptions != nil {
		if *expected.Exceptions != *result.Exceptions {
			t.Errorf("test Exceptions: %d, expected: %v+ ,\n got: %v+", k, *expected.Exceptions, *result.Exceptions)
		}
	}
	if expected.Output.BGcolor != nil {
		if *expected.Output.BGcolor != *result.Output.BGcolor {
			t.Errorf("test BGcolor: %d, expected: %v+ ,\n got: %v+", k, *expected.Output.BGcolor, *result.Output.BGcolor)
		}
	}
}
