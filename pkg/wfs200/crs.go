package wfs200

import (
	"regexp"
	"strconv"
)

const (
	codeSpace = `urn:ogc:def:crs:EPSG::`
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

// ParseString build CRS struct from input string
func (c *CRS) ParseString(s string) {
	c.parseString(s)
}

func (c *CRS) parseString(s string) {
	regex := regexp.MustCompile(`(^.*):([0-9]+)`)
	code := regex.FindStringSubmatch(s)
	if len(code) == 3 { // code[0] is the full match, the other the parts
		c.Namespace = codeSpace

		// the regex already checks if it [0-9]
		i, _ := strconv.Atoi(code[2])
		c.Code = i
	}
}
