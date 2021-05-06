package wsc110

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
