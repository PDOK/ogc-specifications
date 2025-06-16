package wms130

// UnmarshalYAML CRS
func (c *CRS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var crs CRS
	crs.parseString(s)

	*c = crs

	return nil
}

func (c CRS) MarshalYAML() (interface{}, error) {
	return c.Identifier(), nil
}
