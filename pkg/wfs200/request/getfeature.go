package request

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wfs200/exception"
)

// Contains the GetFeature struct and specific functions for building a GetFeature request

//
const (
	getfeature = `GetFeature`

	// table5
	STARTINDEX = `STARTINDEX`
	COUNT      = `COUNT`
	RESULTTYPE = `RESULTTYPE`
	// table6
	RESOLVE        = `RESOLVE`
	RESOLVEDEPTH   = `RESOLVEDEPTH`
	RESOLVETIMEOUT = `RESOLVETIMEOUT`
	// table7
	NAMESPACES = `NAMESPACES`
	// table8
	TYPENAMES      = `TYPENAMES`
	ALIASES        = `ALIASES`
	SRSNAME        = `SRSNAME`
	FILTER         = `FILTER`
	FILTERLANGUAGE = `FILTER_LANGUAGE`
	RESOURCEID     = `RESOURCEID`
	BBOX           = `BBOX` // OGC 06-121r3
	SORTBY         = `SORTBY`
	// table10
	STOREDQUERYID = `STOREDQUERY_ID`
)

// Type returns GetFeature
func (gf *GetFeature) Type() string {
	return getfeature
}

// WFS tables as map[string]bool, where the key (string) is the TOKEN and the bool if its a mandatory (true) or optional (false) attribute
var table5 = map[string]bool{STARTINDEX: false, COUNT: false, OUTPUTFORMAT: false, RESULTTYPE: false}

//var table6 = map[string]bool{RESOLVE: false, RESOLVEDEPTH: false, RESOLVETIMEOUT: false}
var table7 = map[string]bool{NAMESPACES: false} //VSPs (<- vendor specific parameters)
var table8 = map[string]bool{TYPENAMES: true, ALIASES: false, SRSNAME: false, FILTER: false, FILTERLANGUAGE: false, RESOURCEID: false, BBOX: false, SORTBY: false}

//var table10 = map[string]bool{STOREDQUERYID: true} //storedquery_parameter=value

// ParseXML builds a GetCapabilities object based on a XML document
func (gf *GetFeature) ParseXML(doc []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return ows.NoApplicableCode("Could not process XML, is it XML?")
	}
	xml.Unmarshal(doc, &gf) //When object can be Unmarshalled -> XMLAttributes, it can be Unmarshalled -> GetFeature
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		case STARTINDEX:
		case COUNT:
		case OUTPUTFORMAT:
		default:
			n = append(n, a)
		}
	}
	gf.BaseRequest.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// ParseKVP builds a GetCapabilities object based on the available query parameters
// All the keys from the query url.Values need to be UpperCase, this is done during the execution of the operations.ValidRequest()
func (gf *GetFeature) ParseKVP(query url.Values) ows.Exception {

	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return ows.MissingParameterValue(VERSION)
	}

	q := utils.KeysToUpper(query)

	// Base
	if len(q[REQUEST]) > 0 {
		gf.XMLName.Local = q[REQUEST][0]
	}

	var br BaseRequest
	if err := br.ParseKVP(q); err != nil {
		return err
	}
	gf.BaseRequest = br

	// Table 5
	for k, m := range table5 {
		if len(q[k]) > 0 {
			switch k {
			case STARTINDEX:
				i, _ := strconv.Atoi(q[k][0])
				gf.Startindex = &i
			case COUNT:
				i, _ := strconv.Atoi(q[k][0])
				gf.Count = &i
			case OUTPUTFORMAT:
				gf.OutputFormat = &q[k][0]
			case RESULTTYPE:
				gf.ResultType = &q[k][0]
			default:
				if m {
					//TODO add return error, missing mandatory key... or accept for now and check during validation
				}
			}
		}
	}

	// Table 7
	for k, m := range table7 {
		if len(q[k]) > 0 {
			switch k {
			case NAMESPACES:
				gf.BaseRequest.Attr = procesNamespaces(q[k][0])
			default:
				if m {
					//TODO add return error, missing mandatory key... or accept for now and check during validation
				}
			}
		}
	}

	// Table 8
	for k, m := range table8 {
		if len(q[k]) > 0 {
			switch k {
			case TYPENAMES:
				gf.Query.TypeNames = q[k][0]
			case ALIASES:
				// TODO
				// 7.9.2.4.3 aliases parameter
				// fes:AbstractAdhocQueryExpressionType type (see ISO 19143, 6.3.2)
			case SRSNAME:
				gf.Query.SrsName = &q[k][0]
			case FILTER:
				var filter Filter
				if err := xml.Unmarshal([]byte(q[k][0]), &filter); err != nil {
					// TODO what if the filter is corrupt
					// Now it won't unmarshal resulting in a empty/corrupt (but maybe valid) filter object
					// Validation of the content is handled further downstream
				}
				if gf.Query.Filter != nil {
					// We are at this point only interressed in the RESOURCEID's
					// When none are found the filter will we overwritten with this one from the FILTER=
					if gf.Query.Filter.ResourceID != nil {
						mergedRids := mergeResourceIDGroups(*gf.Query.Filter.ResourceID, *filter.ResourceID)
						filter.ResourceID = &mergedRids
					}
				}
				gf.Query.Filter = &filter
			case FILTERLANGUAGE:
				// TODO
				// See ISO 19143:2010, 6.3.3 seems to behind a pay wall...
				// For now we are gonna skip it
			case RESOURCEID:
				// Resourceid's are
				ids := strings.Split(q[k][0], `,`)
				var resourceids []ResourceID
				for _, id := range ids {
					resourceids = append(resourceids, ResourceID{Rid: id})
				}
				if gf.Query.Filter != nil {
					mergedRids := mergeResourceIDGroups(*gf.Query.Filter.ResourceID, resourceids)
					gf.Query.Filter.ResourceID = &mergedRids
				} else {
					var filter Filter
					filter.ResourceID = &resourceids
					gf.Query.Filter = &filter
				}
			case BBOX:
				var geobbox GEOBBOX
				geobbox.UnmarshalText(q[k][0])
				if gf.Query.Filter != nil {
					gf.Query.Filter.BBOX = &geobbox
				} else {
					var filter Filter
					filter.BBOX = &geobbox
					gf.Query.Filter = &filter
				}
			case SORTBY:
			default:
				if m {
					//TODO add return error, missing mandatory key... or accept for now and check during validation
				}
			}
		}
	}
	return nil
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
// TODO: In the Filter>Query>... the content of the GeometryOperand (Point,Line,Polygon,...) is the raw xml (text)
func (gf *GetFeature) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gf, "", " ")
	return append([]byte(xml.Header), si...)
}

// BuildKVP builds a new query string that will be proxied
func (gf *GetFeature) BuildKVP() url.Values {
	querystring := make(map[string][]string)
	// base
	querystring[REQUEST] = []string{gf.XMLName.Local}
	querystring[SERVICE] = []string{gf.BaseRequest.Service}
	querystring[VERSION] = []string{gf.BaseRequest.Version}

	// Table 5
	for k, v := range gf.BaseGetFeatureRequest.BuildQueryString() {
		querystring[k] = v
	}

	// Table 7
	// Table 8
	for k, v := range gf.Query.BuildQueryString() {
		querystring[k] = v
	}
	return querystring
}

func mergeResourceIDGroups(rids ...[]ResourceID) []ResourceID {
	var mergedRids []ResourceID
	for _, grp := range rids {
		for _, rid := range grp {
			mergedRids = append(mergedRids, rid)
		}
	}
	return mergedRids
}

// the use of a map make that with dublicate namespaces prefixes the last match is used
func procesNamespaces(namespace string) []xml.Attr {
	regex := regexp.MustCompile(`xmlns\((.*?)\)`)
	namespacematches := regex.FindAllStringSubmatch(namespace, -1)
	attributemap := make(map[string]string)
	for _, match := range namespacematches {
		n := strings.Split(match[1], ",")[0]
		v := strings.Split(match[1], ",")[1]
		attributemap[n] = v
	}

	var attributes []xml.Attr
	for k, v := range attributemap {
		attributes = append(attributes, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}

	return attributes
}

// BaseGetFeatureRequest struct used by GetFeature
type BaseGetFeatureRequest struct {
	OutputFormat *string `xml:"outputFormat,attr" yaml:"outputformat"`
	Count        *int    `xml:"count,attr" yaml:"count"`
	Startindex   *int    `xml:"startindex,attr" yaml:"startindex"`
	ResultType   *string `xml:"resultType,attr" yaml:"resulttype"`
}

// BuildQueryString for BaseGetFeatureRequest struct
func (b *BaseGetFeatureRequest) BuildQueryString() url.Values {
	querystring := make(map[string][]string)

	for k := range table5 {
		switch k {
		case STARTINDEX:
			if b.Startindex != nil {
				querystring[STARTINDEX] = []string{strconv.Itoa(*b.Startindex)}
			}
		case COUNT:
			if b.Count != nil {
				querystring[COUNT] = []string{strconv.Itoa(*b.Count)}
			}
		case OUTPUTFORMAT:
			if b.OutputFormat != nil {
				querystring[OUTPUTFORMAT] = []string{*b.OutputFormat}
			}
		case RESULTTYPE:
			if b.ResultType != nil {
				querystring[RESULTTYPE] = []string{*b.ResultType}
			}
		}
	}
	return querystring
}

// BuildQueryString for Query struct
func (q *Query) BuildQueryString() url.Values {
	querystring := make(map[string][]string)

	for k := range table8 {
		switch k {
		case TYPENAMES:
			// TODO
			// typenames is now a string -> []string
			if len(q.TypeNames) > 0 {
				querystring[TYPENAMES] = []string{q.TypeNames}
			}
		case ALIASES:
			// TODO
			// 7.9.2.4.3 aliases parameter
			// fes:AbstractAdhocQueryExpressionType type (see ISO 19143, 6.3.2)
		case SRSNAME:
			if q.SrsName != nil {
				querystring[TYPENAMES] = []string{*q.SrsName}
			}
		case FILTER:
			if q.Filter != nil {
				for k, v := range q.Filter.BuildQueryString() {
					querystring[k] = v
				}
			}
		case FILTERLANGUAGE:
			// TODO
			// See ISO 19143:2010, 6.3.3 seems to behind a pay wall...
			// For now we are gonna skip it
		case RESOURCEID:
			// Will be in Filter object
		case BBOX:
			// Will be in Filter object
		case SORTBY:
		}
	}
	return querystring
}

// Query struct for parsing the WFS filter xml
type Query struct {
	TypeNames    string    `xml:"typeNames,attr" yaml:"typenames"`
	SrsName      *string   `xml:"srsName,attr" yaml:"srsname"`
	Filter       *Filter   `xml:"Filter" yaml:"filter"`
	SortBy       *SortBy   `xml:"SortBy" yaml:"sortby"`
	PropertyName *[]string `xml:"PropertyName" yaml:"propertyname"`
}

// BuildQueryString for Filter struct
func (f *Filter) BuildQueryString() url.Values {
	querystring := make(map[string][]string)
	si, _ := xml.Marshal(f)
	if len(si) > 0 {
		querystring[FILTER] = []string{url.QueryEscape(string(si))}
	}
	return querystring
}

// Filter struct for Query
type Filter struct {
	AND        *AND          `xml:"AND" yaml:"and"`
	OR         *OR           `xml:"OR" yaml:"or"`
	NOT        *NOT          `xml:"NOT" yaml:"not"`
	ResourceID *[]ResourceID `xml:"ResourceId" yaml:"resourceid"`
	ComparisonOperator
	SpatialOperator
}

// AND struct for Filter
type AND struct {
	AND *AND `xml:"AND" yaml:"and"`
	OR  *OR  `xml:"OR" yaml:"or"`
	NOT *NOT `xml:"NOT" yaml:"not"`
	ComparisonOperator
	SpatialOperator
}

// OR struct for Filter
type OR struct {
	AND *AND `xml:"AND" yaml:"and"`
	OR  *OR  `xml:"OR" yaml:"or"`
	NOT *NOT `xml:"NOT" yaml:"not"`
	ComparisonOperator
	SpatialOperator
}

// NOT struct for Filter
type NOT struct {
	AND *AND `xml:"AND" yaml:"and"`
	OR  *OR  `xml:"OR" yaml:"or"`
	NOT *NOT `xml:"NOT" yaml:"not"`
	ComparisonOperator
	SpatialOperator
}

// ResourceID struct for Filter
type ResourceID struct {
	Rid string `xml:"rid,attr" yaml:"rid"`
}

// ComparisonOperator struct for Filter
type ComparisonOperator struct {
	PropertyIsEqualTo              *[]PropertyIsEqualTo              `xml:"PropertyIsEqualTo" yaml:"propertyisequalto"`
	PropertyIsNotEqualTo           *[]PropertyIsNotEqualTo           `xml:"PropertyIsNotEqualTo" yaml:"propertyisnotequalto"`
	PropertyIsLessThan             *[]PropertyIsLessThan             `xml:"PropertyIsLessThan" yaml:"propertyislessthan"`
	PropertyIsGreaterThan          *[]PropertyIsGreaterThan          `xml:"PropertyIsGreaterThan" yaml:"propertyisgreaterthan"`
	PropertyIsLessThanOrEqualTo    *[]PropertyIsLessThanOrEqualTo    `xml:"PropertyIsLessThanOrEqualTo" yaml:"propertyislessthanorequalto"`
	PropertyIsGreaterThanOrEqualTo *[]PropertyIsGreaterThanOrEqualTo `xml:"PropertyIsGreaterThanOrEqualTo" yaml:"propertyisgreaterthanorequalto"`
	PropertyIsBetween              *[]PropertyIsBetween              `xml:"PropertyIsBetween" yaml:"propertyisbetween"`
	PropertyIsLike                 *[]PropertyIsLike                 `xml:"PropertyIsLike" yaml:"propertyislike"`
}

// ComparisonOperatorAttribute struct for the ComparisonOperators
type ComparisonOperatorAttribute struct {
	MatchCase      *string `xml:"matchCase,attr" yaml:"matchcase"`
	PropertyName   *string `xml:"PropertyName" yaml:"propertyname"`     // property i.e: id
	ValueReference *string `xml:"ValueReference" yaml:"valuereference"` // path to a property i.e: the/path/to/a/id/in/a/document or just a id ...
	Literal        string  `xml:"Literal" yaml:"literal"`
}

// PropertyIsEqualTo for ComparisonOperator
type PropertyIsEqualTo struct {
	ComparisonOperatorAttribute
}

// PropertyIsNotEqualTo for ComparisonOperator
type PropertyIsNotEqualTo struct {
	ComparisonOperatorAttribute
}

// PropertyIsLessThan for ComparisonOperator
type PropertyIsLessThan struct {
	ComparisonOperatorAttribute
}

// PropertyIsGreaterThan for ComparisonOperator
type PropertyIsGreaterThan struct {
	ComparisonOperatorAttribute
}

// PropertyIsLessThanOrEqualTo for ComparisonOperator
type PropertyIsLessThanOrEqualTo struct {
	ComparisonOperatorAttribute
}

// PropertyIsGreaterThanOrEqualTo for ComparisonOperator
type PropertyIsGreaterThanOrEqualTo struct {
	ComparisonOperatorAttribute
}

// PropertyIsLike for ComparisonOperator
// wildcard='*' singleChar='.' escape='!'>
type PropertyIsLike struct {
	Wildcard   string `xml:"wildcard,attr" yaml:"wildcard"`
	SingleChar string `xml:"singleChar,attr" yaml:"singlechar"`
	Escape     string `xml:"escape,attr" yaml:"escape"`
	ComparisonOperatorAttribute
}

// PropertyIsBetween for ComparisonOperator
type PropertyIsBetween struct {
	PropertyName  string `xml:"PropertyName" yaml:"propertyname"`
	LowerBoundary string `xml:"LowerBoundary" yaml:"lowerboundary"`
	UpperBoundary string `xml:"UpperBoundary" yaml:"upperboundary"`
}

// GeometryOperand struct for Filter
type GeometryOperand struct {
	Point           *Point           `xml:"Point" yaml:"point"`
	MultiPoint      *MultiPoint      `xml:"MultiPoint" yaml:"multipoint"`
	LineString      *LineString      `xml:"LineString" yaml:"linestring"`
	MultiLineString *MultiLineString `xml:"MultiLineString" yaml:"multiLinestring"`
	Curve           *Curve           `xml:"Curve" yaml:"curve"`
	MultiCurve      *MultiCurve      `xml:"MultiCurve" yaml:"multicurve"`
	Polygon         *Polygon         `xml:"Polygon" yaml:"polygon"`
	MultiPolygon    *MultiPolygon    `xml:"MultiPolygon" yaml:"multipolygon"`
	Surface         *Surface         `xml:"Surface" yaml:"surface"`
	MultiSurface    *MultiSurface    `xml:"MultiSurface" yaml:"multisurface"`
	Box             *Box             `xml:"Box" yaml:"box"`
	Envelope        *Envelope        `xml:"Envelope" yaml:"envelope"`
}

// Geometry struct for GeometryOperand geometries
type Geometry struct {
	SrsName string `xml:"srsName,attr" yaml:"srname"`
	Content string `xml:",innerxml"`
}

// Point struct for GeometryOperand
type Point struct {
	Geometry
}

// MultiPoint struct for GeometryOperand
type MultiPoint struct {
	Geometry
}

// LineString struct for GeometryOperand
type LineString struct {
	Geometry
}

// MultiLineString struct for GeometryOperand
type MultiLineString struct {
	Geometry
}

// Curve struct for GeometryOperand
type Curve struct {
	Geometry
}

// MultiCurve struct for GeometryOperand
type MultiCurve struct {
	Geometry
}

// Polygon struct for GeometryOperand
type Polygon struct {
	Geometry
}

// MultiPolygon struct for GeometryOperand
type MultiPolygon struct {
	Geometry
}

// Surface struct for GeometryOperand
type Surface struct {
	Geometry
}

// MultiSurface struct for GeometryOperand
type MultiSurface struct {
	Geometry
}

// Box struct for GeometryOperand
type Box struct {
	Geometry
}

// Envelope struct for GeometryOperand
type Envelope struct {
	LowerCorner ows.Position `xml:"lowerCorner" yaml:"lowercorner"`
	UpperCorner ows.Position `xml:"upperCorner" yaml:"uppercorner"`
}

// SpatialOperator struct for Filter
type SpatialOperator struct {
	Equals     *Equals     `xml:"Equals" yaml:"equals"`
	Disjoint   *Disjoint   `xml:"Disjoint" yaml:"disjoint"`
	Touches    *Touches    `xml:"Touches" yaml:"touches"`
	Within     *Within     `xml:"Within" yaml:"within"`
	Overlaps   *Overlaps   `xml:"Overlaps" yaml:"overlaps"`
	Crosses    *Crosses    `xml:"Crosses" yaml:"crosses"`
	Intersects *Intersects `xml:"Intersects" yaml:"intersects"`
	Contains   *Contains   `xml:"Contains" yaml:"contains"`
	DWithin    *DWithin    `xml:"DWithin" yaml:"dwithin"`
	Beyond     *Beyond     `xml:"Beyond" yaml:"beyond"`
	BBOX       *GEOBBOX    `xml:"BBOX" yaml:"bbox"`
}

// Equals for SpatialOperator
type Equals struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Disjoint for SpatialOperator
type Disjoint struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Touches for SpatialOperator
type Touches struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Within for SpatialOperator
type Within struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Overlaps for SpatialOperator
type Overlaps struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Crosses for SpatialOperator
type Crosses struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Intersects for SpatialOperator
type Intersects struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// Contains for SpatialOperator
type Contains struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
}

// DWithin for SpatialOperator
type DWithin struct {
	PropertyName string `xml:"PropertyName" yaml:"propertyname"`
	GeometryOperand
	Distance Distance `xml:"Distance" yaml:"distance"`
}

// Beyond for SpatialOperator
type Beyond struct {
	Units string `xml:"unit,attr" yaml:"unit"`
	GeometryOperand
	Distance Distance `xml:"Distance" yaml:"distance"`
}

// Distance for DWithin and Beyond
type Distance struct {
	Units string `xml:"units,attr" yaml:"unit"`
	Text  string `xml:",chardata"`
}

// GEOBBOX for SpatialOperator
type GEOBBOX struct {
	Units          *string  `xml:"unit,attr" yaml:"unit"` // unit or units..
	SrsName        *string  `xml:"srsName,attr" yaml:"srsname"`
	ValueReference *string  `xml:"ValueReference" yaml:"valuereference"`
	Envelope       Envelope `xml:"Envelope" yaml:"envelope"`
	// Text           string   `xml:",chardata"`
	// <fes:BBOX>
	// 	<fes:ValueReference>/RS1/geometry</fes:ValueReference>
	// 	<gml:Envelope srsName="urn:ogc:def:crs:EPSG::1234">
	// 		<gml:lowerCorner>10 10</gml:lowerCorner>
	// 		<gml:upperCorner>20 20</gml:upperCorner>
	// 	</gml:Envelope>
	// </fes:BBOX>
}

// UnmarshalText a string to a GEOBBOX object
func (gb *GEOBBOX) UnmarshalText(q string) ows.Exception {
	regex := regexp.MustCompile(`,`)
	result := regex.Split(q, -1)
	if len(result) == 4 || len(result) == 5 {

		var lx, ly, ux, uy float64
		var err error

		if lx, err = strconv.ParseFloat(result[0], 64); err != nil {
			return exception.InvalidValue(BBOX)
		}
		if ly, err = strconv.ParseFloat(result[1], 64); err != nil {
			return exception.InvalidValue(BBOX)
		}
		if ux, err = strconv.ParseFloat(result[2], 64); err != nil {
			return exception.InvalidValue(BBOX)
		}
		if uy, err = strconv.ParseFloat(result[3], 64); err != nil {
			return exception.InvalidValue(BBOX)
		}

		gb.Envelope.LowerCorner = ows.Position{lx, ly}
		gb.Envelope.UpperCorner = ows.Position{ux, uy}
	}
	if len(result) == 5 {
		gb.SrsName = &result[4]
	}
	return nil
}

// MarshalText build a KVP string of a GEOBBOX object
func (gb *GEOBBOX) MarshalText() string {
	regex := regexp.MustCompile(` `)
	var str string
	if len(gb.Envelope.LowerCorner) >= 2 && len(gb.Envelope.UpperCorner) >= 2 && gb.Envelope.LowerCorner != gb.Envelope.UpperCorner {
		str = fmt.Sprintf("%f,%f,%f,%f", gb.Envelope.LowerCorner[0], gb.Envelope.LowerCorner[1], gb.Envelope.UpperCorner[0], gb.Envelope.UpperCorner[1])
	}
	if len(str) > 0 && gb.SrsName != nil {
		str = str + ` ` + *gb.SrsName
	}
	return regex.ReplaceAllString(str, `,`)
}

// SortBy for Query
type SortBy struct {
	SortProperty *[]SortProperty `xml:"SortProperty" yaml:"sortproperty"`
}

// SortProperty for SortBy
type SortProperty struct {
	Content string `xml:",innerxml"`
}

// ProjectionClause based on Table 9 WFS2.0.0 spec
type ProjectionClause struct {
	Propertyname string
}

// StoredQuery based on Table 10 WFS2.0.0 spec
type StoredQuery struct {
	StoredQueryID string
}

// GetFeature struct with the needed parameters/attributes needed for making a GetFeature request
type GetFeature struct {
	XMLName xml.Name `xml:"GetFeature" yaml:"getfeature"`
	BaseRequest
	BaseGetFeatureRequest
	Query Query `xml:"Query" yaml:"query"`
}
