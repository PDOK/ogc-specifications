package wms130

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

func TestBuildStyledLayerDescriptor(t *testing.T) {
	var tests = []struct {
		layers    []string
		styles    []string
		sld       StyledLayerDescriptor
		Exception Exceptions
	}{
		0: {layers: []string{"layer1", "layer2"}, styles: []string{"style1", "style2"}, sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1", NamedStyle: &NamedStyle{Name: "style1"}}, {Name: "layer2", NamedStyle: &NamedStyle{Name: "style2"}}}}},
		1: {layers: []string{"layer1", "layer2"}, styles: []string{"style1", "style2", "style3"}, Exception: StyleNotDefined().ToExceptions()},
		2: {layers: []string{"layer1", "layer2"}, styles: []string{"style1"}, Exception: StyleNotDefined().ToExceptions()},
		3: {layers: []string{"layer1", "layer2"}, sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1"}, {Name: "layer2"}}}},
	}

	for k, test := range tests {
		result, exceptions := buildStyledLayerDescriptor(test.layers, test.styles)

		if exceptions != nil {
			if test.Exception == nil || exceptions[0] != test.Exception[0] {
				t.Errorf("test: %d, expected: %+v \ngot: %+v", k, test.Exception, exceptions)
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

func TestValidateStyledLayerDescriptor(t *testing.T) {
	var tests = []struct {
		capabilities Capabilities
		sld          StyledLayerDescriptor
		exceptions   Exceptions
	}{
		0: {
			capabilities: Capabilities{
				WMSCapabilities: WMSCapabilities{
					Layer: []Layer{
						{Name: sp(`layer1`)},
						{Name: sp(`layer2`), Style: []*Style{{Name: `styleone`}}},
					},
				},
			},
			sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1", NamedStyle: &NamedStyle{Name: ``}}, {Name: "layer2", NamedStyle: &NamedStyle{Name: `styleone`}}}},
		},
		1: {
			capabilities: Capabilities{
				WMSCapabilities: WMSCapabilities{
					Layer: []Layer{
						{Name: sp(`layer2`), Style: []*Style{{Name: `styleone`}}},
						{Name: sp(`layer3`)},
					},
				},
			},
			sld:        StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1"}, {Name: "layer2", NamedStyle: &NamedStyle{Name: `styletwo`}}}},
			exceptions: Exceptions{LayerNotDefined(`layer1`), StyleNotDefined(`styletwo`, `layer2`)},
		},
	}

	for k, test := range tests {
		exceptions := test.sld.Validate(test.capabilities)
		if len(exceptions) > 0 {
			for _, exception := range exceptions {
				found := false
				for _, testexception := range test.exceptions {
					if exception == testexception {
						found = true
					}
				}
				if !found {
					t.Errorf("test exception: %d, expected one of: %s ,\n got: %s", k, test.exceptions, exception.Error())
				}
			}
		}
	}
}

func TestGetMapParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		excepted  GetMapRequest
		exception exception
	}{
		// GetMap http://schemas.opengis.net/sld/1.1.0/example_getmap.xml example request
		0: {body: []byte(`<GetMap xmlns="http://www.opengis.net/sld" xmlns:gml="http://www.opengis.net/gml" xmlns:ogc="http://www.opengis.net/ogc" xmlns:ows="http://www.opengis.net/ows" 
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
			excepted: GetMapRequest{
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
			},
		},
		1: {body: []byte(``),
			exception: MissingParameterValue()},
		2: {body: []byte(`<UnknownTag/>`), excepted: GetMapRequest{}},
	}
	for k, test := range tests {
		var gm GetMapRequest
		exceptions := gm.ParseXML(test.body)
		if exceptions != nil {
			if exceptions[0].Error() != test.exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.exception, exceptions)
			}
		} else {
			compareGetMapObject(gm, test.excepted, t, k)
		}
	}
}

func TestGetLayerParameterValue(t *testing.T) {
	var tests = []struct {
		styledLayerDescriptor StyledLayerDescriptor
		excepted              string
	}{
		0: {styledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers"},
				{Name: "Roads"},
				{Name: "Houses"},
			}},
			excepted: "Rivers,Roads,Houses"},
		1: {styledLayerDescriptor: StyledLayerDescriptor{},
			excepted: ""},
		2: {styledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers"},
			}},
			excepted: "Rivers"},
	}

	for k, test := range tests {
		result := test.styledLayerDescriptor.getLayerParameterValue()
		if test.excepted != result {
			t.Errorf("test Exceptions: %d, expected: %v+ ,\n got: %v+", k, test.excepted, result)
		}
	}
}

func TestGetStyleParameterValue(t *testing.T) {
	var tests = []struct {
		styledLayerDescriptor StyledLayerDescriptor
		excepted              string
	}{
		0: {styledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
			excepted: "CenterLine,CenterLine,Outline"},
		1: {styledLayerDescriptor: StyledLayerDescriptor{},
			excepted: ""},
		2: {styledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads"},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
			excepted: "CenterLine,,Outline"},
		// 4. This needs to fail in the validation step
		4: {styledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
			excepted: ""},
	}

	for k, test := range tests {
		result := test.styledLayerDescriptor.getStyleParameterValue()
		if test.excepted != result {
			t.Errorf("test Exceptions: %d, expected: %v+ ,\n got: %v+", k, test.excepted, result)
		}
	}
}

func TestGetMapParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query     url.Values
		excepted  GetMapRequest
		exception exception
	}{
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers,Roads,Houses&STYLES=CenterLine,CenterLine,Outline&CRS=EPSG:4326&BBOX=invalid&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&TRANSPARENT=FALSE&EXCEPTIONS=XML
		0: {query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:      {`Rivers,Roads,Houses`},
			STYLES:      {`CenterLine,CenterLine,Outline`},
			"CRS":       {`EPSG:4326`},
			BBOX:        {`invalid`},
			WIDTH:       {`1024`},
			HEIGHT:      {`512`},
			FORMAT:      {`image/jpeg`},
			TRANSPARENT: {`FALSE`},
			EXCEPTIONS:  {`XML`},
			BGCOLOR:     {`0x7F7F7F`},
		},
			exception: InvalidParameterValue(`invalid`, `boundingbox`),
		},
		1: {query: url.Values{},
			exception: MissingParameterValue(VERSION)},
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers,Roads,Houses&STYLES=CenterLine,CenterLine,Outline&CRS=EPSG:4326&BBOX=-180.0,-90.0,180.0,90.0&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&TRANSPARENT=FALSE&EXCEPTIONS=XML
		2: {query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
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
		},
			excepted: GetMapRequest{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
				},
				StyledLayerDescriptor: StyledLayerDescriptor{
					NamedLayer: []NamedLayer{
						{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
					}},
				CRS: CRS{Namespace: "EPSG", Code: 4326},
				BoundingBox: BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Output: Output{
					Size:        Size{Width: 1024, Height: 512},
					Format:      "image/jpeg",
					Transparent: bp(false),
					BGcolor:     sp(`0x7F7F7F`)},
				Exceptions: sp("XML"),
			}},
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers,Roads,Houses&STYLES=CenterLine,CenterLine,Outline&CRS=EPSG:4326&BBOX=-180.0,-90.0,180.0,90.0&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&TRANSPARENT=zzzz&EXCEPTIONS=XML
		3: {query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:      {`Rivers,Roads,Houses`},
			STYLES:      {`CenterLine,CenterLine,Outline`},
			"CRS":       {`EPSG:4326`},
			BBOX:        {`-180.0,-90.0,180.0,90.0`},
			WIDTH:       {`1024`},
			HEIGHT:      {`512`},
			FORMAT:      {`image/jpeg`},
			TRANSPARENT: {`zzzz`},
			EXCEPTIONS:  {`XML`},
			BGCOLOR:     {`0x7F7F7F`},
		},
			exception: InvalidParameterValue(`zzzz`, TRANSPARENT),
			},
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers,Roads,Houses&STYLES=CenterLine,CenterLine,Outline&CRS=EPSG:4326&BBOX=-180.0,-90.0,180.0,90.0&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&EXCEPTIONS=XML
		4: {query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:      {`Rivers,Roads,Houses`},
			STYLES:      {`CenterLine,CenterLine,Outline`},
			"CRS":       {`EPSG:4326`},
			BBOX:        {`-180.0,-90.0,180.0,90.0`},
			WIDTH:       {`1024`},
			HEIGHT:      {`512`},
			FORMAT:      {`image/jpeg`},
			EXCEPTIONS:  {`XML`},
			BGCOLOR:     {`0x7F7F7F`},
		},
			excepted: GetMapRequest{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
				},
				StyledLayerDescriptor: StyledLayerDescriptor{
					NamedLayer: []NamedLayer{
						{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
					}},
				CRS: CRS{Namespace: "EPSG", Code: 4326},
				BoundingBox: BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Output: Output{
					Size:        Size{Width: 1024, Height: 512},
					Format:      "image/jpeg",
					Transparent: bp(false),
					BGcolor:     sp(`0x7F7F7F`)},
				Exceptions: sp("XML"),
			}},
		//REQUEST=GetMap&SERVICE=WMS&VERSION=1.3.0&LAYERS=Rivers&STYLES=&CRS=EPSG:4326&BBOX=-180.0,-90.0,180.0,90.0&WIDTH=1024&HEIGHT=512&FORMAT=image/jpeg&EXCEPTIONS=XML
		5: {query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:      {`Rivers`},
			STYLES:      {``},
			"CRS":       {`EPSG:4326`},
			BBOX:        {`-180.0,-90.0,180.0,90.0`},
			WIDTH:       {`1024`},
			HEIGHT:      {`512`},
			FORMAT:      {`image/jpeg`},
			EXCEPTIONS:  {`XML`},
			BGCOLOR:     {`0x7F7F7F`},
		},
			excepted: GetMapRequest{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
				},
				StyledLayerDescriptor: StyledLayerDescriptor{
					NamedLayer: []NamedLayer{
						{Name: "Rivers"},
					}},
				CRS: CRS{Namespace: "EPSG", Code: 4326},
				BoundingBox: BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Output: Output{
					Size:        Size{Width: 1024, Height: 512},
					Format:      "image/jpeg",
					Transparent: bp(false),
					BGcolor:     sp(`0x7F7F7F`)},
				Exceptions: sp("XML"),
			}},
	}
	for k, test := range tests {
		var gm GetMapRequest
		exceptions := gm.ParseQueryParameters(test.query)
		if exceptions != nil {
			if exceptions[0].Error() != test.exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.exception, exceptions)
			}
		} else {
			compareGetMapObject(gm, test.excepted, t, k)
		}
	}
}

func TestGetMapToQueryParameters(t *testing.T) {
	var tests = []struct {
		object    GetMapRequest
		excepted  url.Values
		exception exception
	}{
		0: {object: GetMapRequest{
			XMLName: xml.Name{Local: "GetMap"},
			BaseRequest: BaseRequest{
				Version: "1.3.0",
				Service: "WMS",
			},
			StyledLayerDescriptor: StyledLayerDescriptor{
				NamedLayer: []NamedLayer{
					{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
				}},
			CRS: CRS{Namespace: "EPSG", Code: 4326},
			BoundingBox: BoundingBox{
				LowerCorner: [2]float64{-180.0, -90.0},
				UpperCorner: [2]float64{180.0, 90.0},
			},
			Output: Output{
				Size:        Size{Width: 1024, Height: 512},
				Format:      "image/jpeg",
				Transparent: bp(false)},
			Exceptions: sp("XML"),
		},
			excepted: map[string][]string{
				VERSION:     {Version},
				LAYERS:      {`Rivers,Roads,Houses`},
				STYLES:      {`CenterLine,CenterLine,Outline`},
				"CRS":       {`EPSG:4326`},
				BBOX:        {`-180.000000,-90.000000,180.000000,90.000000`},
				EXCEPTIONS:  {`XML`},
				FORMAT:      {`image/jpeg`},
				HEIGHT:      {`512`},
				WIDTH:       {`1024`},
				TRANSPARENT: {`false`},
				REQUEST:     {`GetMap`},
				SERVICE:     {`WMS`},
			}},
		1: {object: GetMapRequest{
			CRS: CRS{Namespace: "EPSG", Code: 4326},
			BoundingBox: BoundingBox{
				LowerCorner: [2]float64{-180.0, -90.0},
				UpperCorner: [2]float64{180.0, 90.0},
			},
			Exceptions: sp(`XML`),
		},
			excepted: map[string][]string{
				LAYERS:     {``},
				STYLES:     {``},
				"CRS":      {`EPSG:4326`},
				BBOX:       {`-180.000000,-90.000000,180.000000,90.000000`},
				FORMAT:     {``},
				HEIGHT:     {`0`},
				WIDTH:      {`0`},
				VERSION:    {Version},
				REQUEST:    {`GetMap`},
				SERVICE:    {`WMS`},
				EXCEPTIONS: {`XML`},
			}},
	}

	for k, test := range tests {
		url := test.object.ToQueryParameters()
		if len(test.excepted) != len(url) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.excepted, url)
		} else {
			for _, rid := range url {
				found := false
				for _, erid := range test.excepted {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.excepted, url)
				}
			}
		}
	}
}

func TestGetMapToXML(t *testing.T) {
	var tests = []struct {
		gm     GetMapRequest
		result string
	}{
		0: {gm: GetMapRequest{},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetMap service="" version="">
 <StyledLayerDescriptor version=""></StyledLayerDescriptor>
 <CRS>
  <Namespace></Namespace>
  <Code>0</Code>
 </CRS>
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

	for k, test := range tests {
		body := test.gm.ToXML()

		if string(body) != test.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, test.result, string(body))
		}
	}

}

func TestGetNamedStyles(t *testing.T) {
	var tests = []struct {
		sld    StyledLayerDescriptor
		styles []string
	}{
		0: {sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1", NamedStyle: &NamedStyle{Name: "style1"}}, {Name: "layer2", NamedStyle: &NamedStyle{Name: "style2"}}}},
			styles: []string{"style1", "style2"}},
		1: {sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1"}, {Name: "layer2"}}},
			styles: []string{"", ""}},
		2: {sld: StyledLayerDescriptor{NamedLayer: []NamedLayer{{Name: "layer1", NamedStyle: &NamedStyle{Name: "style1"}}, {Name: "layer2"}, {Name: "layer3", NamedStyle: &NamedStyle{Name: "style3"}}}},
			styles: []string{"style1", "", "style3"}},
	}

	for k, test := range tests {
		styleslist := test.sld.getNamedStyles()
		if len(styleslist) != len(test.styles) {
			t.Errorf("test: %d, Expected %s but was not \n got: %s", k, test.styles, styleslist)
		} else {
			for _, style := range styleslist {
				found := false
				for _, expected := range test.styles {
					if expected == style {
						found = true
					}
				}
				if !found {
					t.Errorf("test: %d, Expected %s but was not \n got: %s", k, test.styles, style)
				}
			}
		}
	}
}

func TestCheckCRS(t *testing.T) {
	definedCrs := []CRS{{Namespace: `CRS`, Code: 84}, {Namespace: `EPSG`, Code: 4326}, {Namespace: `EPSG`, Code: 3857}}
	var tests = []struct {
		crs       CRS
		exception Exceptions
	}{
		0: {crs: CRS{Namespace: `CRS`, Code: 84}},
		1: {crs: CRS{Namespace: `UNKNOWN`}, exception: InvalidCRS(`UNKNOWN`).ToExceptions()},
	}

	for k, test := range tests {
		exception := checkCRS(test.crs, definedCrs)
		if exception != nil {
			if test.exception != nil {
				if exception[0].Code() != test.exception[0].Code() {
					t.Errorf("test: %d, Expected one of %s but was not \n got: %s", k, test.exception[0].Code(), exception[0].Code())
				}
			} else {
				t.Errorf("test: %d, Expected one of %v but was not \n got: %s", k, definedCrs, test.crs.String())
			}
		}
	}
}

func compareGetMapObject(result, expected GetMapRequest, t *testing.T, k int) {
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
				t.Errorf("test BaseRequest.Attr : %d, expected: %s ,\n got: %s", k, expected.BaseRequest.Attr, result.BaseRequest.Attr)
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
					if result.NamedStyle == nil {
						if expected.NamedStyle == nil {
							c = true
						}
					} else {
						if result.NamedStyle.Name == expected.NamedStyle.Name {
							c = true
						}
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

// ----------
// Validation
// ----------

func TestGetMapValidate(t *testing.T) {
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
		OptionalConstraints: OptionalConstraints{LayerLimit: 1, MaxWidth: 2048, MaxHeight: 2048},
	}

	var tests = []struct {
		gm         GetMapRequest
		exceptions Exceptions
	}{
		0: {gm: GetMapRequest{
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
		}},
	}

	for k, test := range tests {
		getmapexceptions := test.gm.Validate(capabilities)
		if getmapexceptions != nil {
			t.Errorf("test Validation: %d, expected: %v+ ,\n got: %v+", k, test.exceptions, getmapexceptions)
		}
	}
}
