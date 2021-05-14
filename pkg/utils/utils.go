package utils

import (
	"encoding/xml"
	"errors"
	"net/url"
	"strings"
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

func IdentifyRequest(doc []byte) (string, error) {
	var i identify

	if err := xml.Unmarshal(doc, &i); err != nil {
		return ``, errors.New(`unknown REQUEST parameter`) // error string can be used in a OGC type exception
	} else {
		return i.XMLName.Local, nil
	}
}

func IdentifyRequestKVP(query url.Values) (string, error) {
	if query[REQUEST] != nil {
		return query[REQUEST][0], nil
	}

	// error string can be used in a OGC type exception
	return ``, errors.New(`unknown REQUEST parameter`)
}
