package ows

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

// UnmarshalYAML Position
func (p *Position) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	position := getPositionFromString(s)
	*p = Position{position[0], position[1]}

	return nil
}
