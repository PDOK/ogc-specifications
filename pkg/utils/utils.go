package utils

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

const (
	REQUEST = `REQUEST`
)

// KeysToUpper convert all the keys to UpperCase
func KeysToUpper(q url.Values) url.Values {
	r := url.Values{}
	for k, v := range q {
		r[strings.ToUpper(k)] = append(r[strings.ToUpper(k)], v...)
	}
	return r
}

// KeysToLower convert all the keys to LowerCase
func KeysToLower(q url.Values) url.Values {
	r := url.Values{}
	for k, v := range q {
		r[strings.ToLower(k)] = append(r[strings.ToLower(k)], v...)
	}
	return r
}

type identify struct {
	XMLName xml.Name
}

// func (i *identify) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	i.XMLName = start.Name
// 	for {
// 		token, _ := d.Token()
// 		switch el := token.(type) {
// 		case xml.EndElement:
// 			if el == start.End() {
// 				return nil
// 			}
// 		}
// 	}
// }

func IdentifyRequest(doc []byte) (string, wsc110.Exceptions) {
	var i identify

	if err := xml.Unmarshal(doc, &i); err != nil {
		return ``, wsc110.Exceptions{wsc110.MissingParameterValue()}
	} else {
		return i.XMLName.Local, nil
	}
}

func IdentifyRequestKVP(query url.Values) (string, wsc110.Exceptions) {
	if query[REQUEST] != nil {
		return query[REQUEST][0], nil
	}

	return ``, wsc110.Exceptions{wsc110.MissingParameterValue()}
}
