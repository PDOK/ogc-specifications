package wms130

import (
	"encoding/xml"
	"net/url"
	"strings"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

func TestGetFeatureInfoBuildKVP(t *testing.T) {
	var tests = []struct {
		Object    GetFeatureInfoRequest
		Excepted  url.Values
		Exception Exceptions
	}{
		0: {Object: GetFeatureInfoRequest{
			XMLName: xml.Name{Local: `GetFeatureInfo`},
			BaseRequest: BaseRequest{
				Service: Service,
				Version: Version},
			StyledLayerDescriptor: StyledLayerDescriptor{
				NamedLayer: []NamedLayer{
					{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
				}},
			CRS: "EPSG:4326",
			BoundingBox: BoundingBox{
				LowerCorner: [2]float64{-180.0, -90.0},
				UpperCorner: [2]float64{180.0, 90.0},
			},
			Size:         Size{Width: 1024, Height: 512},
			QueryLayers:  []string{`CenterLine`},
			I:            1,
			J:            1,
			InfoFormat:   `application/json`,
			FeatureCount: ip(8),
			Exceptions:   sp(`xml`),
		},
			Excepted: map[string][]string{
				VERSION:      {Version},
				SERVICE:      {Service},
				REQUEST:      {`GetFeatureInfo`},
				BBOX:         {`-180.000000,-90.000000,180.000000,90.000000`},
				"CRS":        {`EPSG:4326`},
				LAYERS:       {`Rivers,Roads,Houses`},
				STYLES:       {`CenterLine,CenterLine,Outline`},
				QUERYLAYERS:  {`CenterLine`},
				WIDTH:        {`1024`},
				HEIGHT:       {`512`},
				I:            {`1`},
				J:            {`1`},
				INFOFORMAT:   {`application/json`},
				FEATURECOUNT: {`8`},
				EXCEPTIONS:   {`xml`},
			},
		},
	}

	for k, n := range tests {
		url := n.Object.ToQueryParameters()
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
			}
		}
	}
}

func TestGetFeatureInfoBuildXML(t *testing.T) {
	var tests = []struct {
		gfi    GetFeatureInfoRequest
		result string
	}{
		0: {gfi: GetFeatureInfoRequest{
			XMLName: xml.Name{Local: `GetFeatureInfo`},
			BaseRequest: BaseRequest{
				Service: Service,
				Version: Version},
			StyledLayerDescriptor: StyledLayerDescriptor{
				NamedLayer: []NamedLayer{
					{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
					{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
				}},
			CRS: "EPSG:4326",
			BoundingBox: BoundingBox{
				LowerCorner: [2]float64{-180.0, -90.0},
				UpperCorner: [2]float64{180.0, 90.0},
			},
			Size:        Size{Width: 1024, Height: 512},
			QueryLayers: []string{`CenterLine`},
			InfoFormat:  `application/json`,
			I:           1,
			J:           1,
		},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetFeatureInfo service="WMS" version="1.3.0">
 <StyledLayerDescriptor version="">
  <NamedLayer>
   <Name>Rivers</Name>
   <NamedStyle>
    <Name>CenterLine</Name>
   </NamedStyle>
  </NamedLayer>
  <NamedLayer>
   <Name>Roads</Name>
   <NamedStyle>
    <Name>CenterLine</Name>
   </NamedStyle>
  </NamedLayer>
  <NamedLayer>
   <Name>Houses</Name>
   <NamedStyle>
    <Name>Outline</Name>
   </NamedStyle>
  </NamedLayer>
 </StyledLayerDescriptor>
 <CRS>EPSG:4326</CRS>
 <BoundingBox>
  <LowerCorner>-180.000000 -90.000000</LowerCorner>
  <UpperCorner>180.000000 90.000000</UpperCorner>
 </BoundingBox>
 <Size>
  <Width>1024</Width>
  <Height>512</Height>
 </Size>
 <QueryLayers>CenterLine</QueryLayers>
 <I>1</I>
 <J>1</J>
 <InfoFormat>application/json</InfoFormat>
</GetFeatureInfo>`},
	}

	for k, v := range tests {
		body := v.gfi.ToXML()

		x := strings.Replace(string(body), "\n", ``, -1)
		y := strings.Replace(v.result, "\n", ``, -1)

		if x != y {
			t.Errorf("test: %d, Expected body: \n%s\nbut was not got: \n%s", k, y, x)
		}
	}
}

func TestGetFeatureInfoParseKVP(t *testing.T) {
	var tests = []struct {
		Query      url.Values
		Excepted   GetFeatureInfoRequest
		Exceptions Exceptions
	}{
		0: {Query: map[string][]string{REQUEST: {getfeatureinfo}, SERVICE: {Service}, VERSION: {Version}}, Exceptions: Exceptions{InvalidParameterValue("", `boundingbox`)}},
		1: {Query: url.Values{}, Exceptions: Exceptions{MissingParameterValue(VERSION), MissingParameterValue(REQUEST)}},
		2: {Query: map[string][]string{REQUEST: {getmap}, SERVICE: {Service}, VERSION: {Version},
			LAYERS:       {`Rivers,Roads,Houses`},
			STYLES:       {`CenterLine,,Outline`},
			"CRS":        {`EPSG:4326`},
			BBOX:         {`-180.0,-90.0,180.0,90.0`},
			WIDTH:        {`1024`},
			HEIGHT:       {`512`},
			FORMAT:       {`image/jpeg`},
			EXCEPTIONS:   {`XML`},
			QUERYLAYERS:  {`Rivers`},
			I:            {`101`},
			J:            {`101`},
			INFOFORMAT:   {`application/json`},
			FEATURECOUNT: {`8`},
		},
			Excepted: GetFeatureInfoRequest{
				BaseRequest: BaseRequest{
					Version: "1.3.0",
				},
				StyledLayerDescriptor: StyledLayerDescriptor{
					NamedLayer: []NamedLayer{
						{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &NamedStyle{Name: ""}},
						{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
					}},
				CRS: "EPSG:4326",
				BoundingBox: BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Size:         Size{Width: 1024, Height: 512},
				Exceptions:   sp("XML"),
				QueryLayers:  []string{`Rivers`},
				I:            101,
				J:            101,
				FeatureCount: ip(8),
				InfoFormat:   `application/json`,
			},
		},
		3: {Query: map[string][]string{WIDTH: {`not a number`}, VERSION: {Version}, BBOX: {`-180.0,-90.0,180.0,90.0`}}, Exceptions: Exceptions{MissingParameterValue(WIDTH, `not a number`)}},
		4: {Query: map[string][]string{WIDTH: {`1024`}, HEIGHT: {`not a number`}, VERSION: {Version}, BBOX: {`-180.0,-90.0,180.0,90.0`}}, Exceptions: Exceptions{MissingParameterValue(HEIGHT, `not a number`)}},
		5: {Query: map[string][]string{WIDTH: {`1024`}, HEIGHT: {`1024`}, I: {`not a number`}, J: {`1`}, VERSION: {Version}, BBOX: {`-180.0,-90.0,180.0,90.0`}}, Exceptions: Exceptions{InvalidPoint(`not a number`, `1`)}},
		6: {Query: map[string][]string{WIDTH: {`1024`}, HEIGHT: {`1024`}, I: {`1`}, J: {`not a number`}, VERSION: {Version}, BBOX: {`-180.0,-90.0,180.0,90.0`}}, Exceptions: Exceptions{InvalidPoint(`1`, `not a number`)}},
		7: {Query: map[string][]string{WIDTH: {`1024`}, HEIGHT: {`1024`}, I: {`this in not a number`}, J: {`this is also not a number`}, VERSION: {Version}, BBOX: {`-180.0,-90.0,180.0,90.0`}}, Exceptions: Exceptions{InvalidPoint(`this in not a number`, `this is also not a number`)}},
	}

	for k, test := range tests {
		var gfi GetFeatureInfoRequest
		errs := gfi.ParseQueryParameters(test.Query)
		if errs != nil {
			if len(errs) != len(test.Exceptions) {
				t.Errorf("test: %d, expected: %d exceptions,\n got: %d exceptions", k, len(test.Exceptions), len(errs))
			} else {
				for _, exception := range errs {
					found := false
					for _, expectedeexception := range test.Exceptions {
						if expectedeexception == exception {
							found = true
						}
					}
					if !found {
						t.Errorf("test: %d, expected one of: %s,\n got: %s", k, test.Exceptions, exception)
					}
				}
			}
		} else {
			compareGetFeatureInfoObject(gfi, test.Excepted, t, k)
		}
	}
}

func TestGetFeatureInfoParseXML(t *testing.T) {
	var tests = []struct {
		Body     []byte
		Excepted GetFeatureInfoRequest
		Error    Exceptions
	}{
		// GetFeatureInfo example request
		0: {Body: []byte(`<GetFeatureInfo xmlns="http://www.opengis.net/sld" xmlns:gml="http://www.opengis.net/gml" xmlns:ogc="http://www.opengis.net/ogc" xmlns:ows="http://www.opengis.net/ows" 
		xmlns:se="http://www.opengis.net/se" xmlns:wms="http://www.opengis.net/wms" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/sld GetFeatureInfo.xsd" version="1.3.0">
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
		<Size>
				<Width>1024</Width>
				<Height>512</Height>
			</Size>
	</GetFeatureInfo>`),
			Excepted: GetFeatureInfoRequest{
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
						xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/sld GetFeatureInfo.xsd"},
					}},
				StyledLayerDescriptor: StyledLayerDescriptor{
					Version: "1.1.0",
					NamedLayer: []NamedLayer{
						{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
						{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
					}},
				CRS: "EPSG:4326",
				BoundingBox: BoundingBox{
					Crs:         "http://www.opengis.net/gml/srs/epsg.xml#4326",
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Size:       Size{Width: 1024, Height: 512},
				Exceptions: sp("XML"),
			},
		},
		1: {Body: []byte(``), Error: MissingParameterValue().ToExceptions()},
		2: {Body: []byte(`<UnknownTag/>`), Error: MissingParameterValue("REQUEST").ToExceptions()},
	}
	for k, n := range tests {
		var gm GetFeatureInfoRequest
		err := gm.ParseXML(n.Body)
		if err != nil {
			if err[0].Error() != n.Error[0].Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
			}
		} else {
			compareGetFeatureInfoObject(gm, n.Excepted, t, k)
		}
	}
}

func compareGetFeatureInfoObject(result, expected GetFeatureInfoRequest, t *testing.T, k int) {
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
		for _, sldnamedlayer := range expected.StyledLayerDescriptor.NamedLayer {
			for _, result := range result.StyledLayerDescriptor.NamedLayer {
				if result.Name == sldnamedlayer.Name {
					if result.NamedStyle.Name == sldnamedlayer.NamedStyle.Name {
						c = true
					}
				}
			}
			if !c {
				t.Errorf("test StyledLayerDescriptor.NamedLayer: %d, expected: %v ,\n got: %v", k, expected.StyledLayerDescriptor.NamedLayer, result.StyledLayerDescriptor.NamedLayer)
			}
			c = false
		}
	} else {
		t.Errorf("test StyledLayerDescriptor: %d, expected: %v ,\n got: %v", k, expected.StyledLayerDescriptor, result.StyledLayerDescriptor)
	}
	if expected.CRS != result.CRS {
		t.Errorf("test CRS: %d, expected: %v ,\n got: %v", k, expected.CRS, result.CRS)
	}
	if expected.BoundingBox != result.BoundingBox {
		t.Errorf("test BoundingBox: %d, expected: %v ,\n got: %v", k, expected.BoundingBox, result.BoundingBox)
	}
	if expected.Size != result.Size {
		t.Errorf("test Output.Size: %d, expected: %v ,\n got: %v", k, expected.Size, result.Size)
	}
	if len(expected.QueryLayers) != len(result.QueryLayers) {
		t.Errorf("test QueryLayers: %d, expected: %v ,\n got: %v", k, expected.QueryLayers, result.QueryLayers)
	} else {
		c := false
		for _, eql := range expected.QueryLayers {
			for _, rql := range result.QueryLayers {
				if eql == rql {
					c = true
				}
			}
			if !c {
				t.Errorf("test QueryLayers: %d, expected: %v ,\n got: %v", k, expected.QueryLayers, result.QueryLayers)
			}
			c = false
		}
	}
	if expected.I != result.I {
		t.Errorf("test I: %d, expected: %v ,\n got: %v", k, expected.I, result.I)
	}
	if expected.J != result.J {
		t.Errorf("test J: %d, expected: %v ,\n got: %v", k, expected.J, result.J)
	}

	if expected.InfoFormat != result.InfoFormat {
		t.Errorf("test InfoFormat: %d, expected: %v ,\n got: %v", k, expected.InfoFormat, result.InfoFormat)
	}

	if expected.FeatureCount != nil {
		if *expected.FeatureCount != *result.FeatureCount {
			t.Errorf("test FeatureCount: %d, expected: %v ,\n got: %v", k, *expected.FeatureCount, *result.FeatureCount)
		}
	}

	if expected.Exceptions != nil {
		if *expected.Exceptions != *result.Exceptions {
			t.Errorf("test Exceptions: %d, expected: %v ,\n got: %v", k, *expected.Exceptions, *result.Exceptions)
		}
	}
}

// ----------
// Benchmarks
// ----------

func BenchmarkGetFeatureInfoBuildKVP(b *testing.B) {
	gfi := GetFeatureInfoRequest{
		XMLName: xml.Name{Local: `GetFeatureInfo`},
		BaseRequest: BaseRequest{
			Service: Service,
			Version: Version},
		StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: "EPSG:4326",
		BoundingBox: BoundingBox{
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Size:        Size{Width: 1024, Height: 512},
		QueryLayers: []string{`CenterLine`},
		InfoFormat:  `application/json`,
		I:           1,
		J:           1,
	}
	for i := 0; i < b.N; i++ {
		gfi.ToQueryParameters()
	}
}

func BenchmarkGetFeatureInfoBuildXML(b *testing.B) {
	gfi := GetFeatureInfoRequest{
		XMLName: xml.Name{Local: `GetFeatureInfo`},
		BaseRequest: BaseRequest{
			Service: Service,
			Version: Version},
		StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: "EPSG:4326",
		BoundingBox: BoundingBox{
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Size:        Size{Width: 1024, Height: 512},
		QueryLayers: []string{`CenterLine`},
		InfoFormat:  `application/json`,
		I:           1,
		J:           1,
	}
	for i := 0; i < b.N; i++ {
		gfi.ToXML()
	}
}
