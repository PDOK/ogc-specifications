package wms130

import (
	"fmt"
	"strconv"
)

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

// MarshalYAML Position
func (p Position) MarshalYAML() (interface{}, error) {
	return fmt.Sprintf("%s %s", strconv.FormatFloat(p[0], 'f', -1, 64), strconv.FormatFloat(p[1], 'f', -1, 64)), nil
}
