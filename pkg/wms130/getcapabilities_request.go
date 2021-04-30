package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilities) ParseXML(body []byte) common.Exceptions {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return common.Exceptions{common.MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return common.Exceptions{common.MissingParameterValue("REQUEST")}
	}
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		default:
			n = append(n, a)
		}
	}

	gc.Attr = common.StripDuplicateAttr(n)
	return nil
}

// ParseKVP builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) ParseKVP(query url.Values) common.Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		return common.Exceptions{common.MissingParameterValue(SERVICE), common.MissingParameterValue(REQUEST)}
	}

	gckvp := GetCapabilitiesKVP{}
	if err := gckvp.ParseKVP(query); err != nil {
		return err
	}

	if err := gc.ParseOperationRequestKVP(&gckvp); err != nil {
		return err
	}

	return nil
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gc *GetCapabilities) ParseOperationRequestKVP(orkvp common.OperationRequestKVP) common.Exceptions {
	gckvp := orkvp.(*GetCapabilitiesKVP)

	gc.XMLName.Local = gckvp.Request
	gc.BaseRequest.Build(gckvp.Service, gckvp.Version)
	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gc *GetCapabilities) BuildKVP() url.Values {
	gckvp := GetCapabilitiesKVP{}
	gckvp.ParseOperationRequest(gc)

	kvp := gckvp.BuildKVP()
	return kvp
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilities) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}
