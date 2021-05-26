package wms130

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	codeSpace = `urn:ogc:def:crs:EPSG::`
	EPSG      = `EPSG`
)

// CRS struct with namespace/authority/registry and code
type CRS struct {
	Namespace string //TODO maybe AuthorityType is a better name...?
	Code      int
}

// String of the EPSGCode
func (c *CRS) String() string {
	return c.Namespace + `:` + strconv.Itoa(c.Code)
}

// Identifier returns the EPSG
func (c *CRS) Identifier() string {
	return codeSpace + strconv.Itoa(c.Code)
}

func (c *CRS) parseString(s string) {
	regex := regexp.MustCompile(`(^.*):([0-9]+)`)
	code := regex.FindStringSubmatch(s)
	if len(code) == 3 { // code[0] is the full match, the other the parts
		f := strings.Index(code[1], EPSG)
		if f > -1 {
			c.Namespace = EPSG
		} else {
			c.Namespace = code[1]
		}

		// the regex already checks if it [0-9]
		i, _ := strconv.Atoi(code[2])
		c.Code = i
	}
}

// MarshalXML Position
func (c *CRS) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s = ``
	if c.Namespace != `` {
		s = fmt.Sprintf("%s:%d", c.Namespace, c.Code)
	}

	return e.EncodeElement(s, start)
}

// UnmarshalXML Position
func (c *CRS) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var crs CRS
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch el := token.(type) {
		case xml.CharData:
			crs.parseString(string([]byte(el)))
		case xml.EndElement:
			if el == start.End() {
				*c = crs
				return nil
			}
		}
	}
}
