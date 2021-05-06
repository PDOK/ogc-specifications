package common

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"
)

//
const (
	codeSpace = `urn:ogc:def:crs:EPSG::`
	EPSG      = `EPSG`
)

// Keywords in struct for repeatability
type Keywords struct {
	Keyword []string `xml:"Keyword" yaml:"keyword"`
}

// StripDuplicateAttr removes the duplicate Attributes from a []Attribute
func StripDuplicateAttr(attr []xml.Attr) []xml.Attr {
	attributemap := make(map[xml.Name]string)
	for _, a := range attr {
		attributemap[xml.Name{Space: a.Name.Space, Local: a.Name.Local}] = a.Value
	}

	var strippedAttr []xml.Attr
	for k, v := range attributemap {
		strippedAttr = append(strippedAttr, xml.Attr{Name: k, Value: v})
	}
	return strippedAttr
}

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

// ParseString build CRS struct from input string
func (c *CRS) ParseString(s string) {
	c.parseString(s)
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

func getPositionFromString(position string) []float64 {
	regex := regexp.MustCompile(` `)
	result := regex.Split(position, -1)
	var ps []float64 //slice because length can be 2 or more

	// check if 'strings' are parsable to float64
	// if one is not return nothing
	for _, fs := range result {
		f, err := strconv.ParseFloat(fs, 64)
		if err != nil {
			return nil
		}
		ps = append(ps, f)
	}
	return ps
}
