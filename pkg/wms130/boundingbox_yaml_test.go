package wms130

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUnmarshalYAMLPosition(t *testing.T) {

	var tests = []struct {
		positionString   []byte
		expectedPosition Position
	}{
		0: {positionString: []byte(`2.52712538742158 50.2128625669452`), expectedPosition: Position{2.52712538742158, 50.2128625669452}},
		1: {positionString: []byte(`7.37402550506231 55.7211602557705`), expectedPosition: Position{7.37402550506231, 55.7211602557705}},
		2: {positionString: []byte(`7.37402550506231 55.7211602557705 0 1 2 3`), expectedPosition: Position{7.37402550506231, 55.7211602557705}},
	}

	for k, test := range tests {
		var pos Position
		err := yaml.Unmarshal(test.positionString, &pos)
		if err != nil {
			t.Errorf("test: %d, yaml.UnMarshal failed with '%s'\n", k, err)
		} else if pos != test.expectedPosition {
			t.Errorf("test: %d, expected: %v+,\n got: %v+", k, test.expectedPosition, pos)
		}
	}
}

func TestMarshalYAMLPosition(t *testing.T) {
	var tests = []struct {
		positionString []byte
		position       Position
	}{
		0: {positionString: []byte("2.52712538742158 50.2128625669452\n"), position: Position{2.52712538742158, 50.2128625669452}},
		1: {positionString: []byte("7.37402550506231 55.7211602557705\n"), position: Position{7.37402550506231, 55.7211602557705}},
	}

	for k, test := range tests {
		pos, err := yaml.Marshal(test.position)
		if err != nil {
			t.Errorf("test: %d, yaml.Marshal failed with '%s'\n", k, err)
		} else if string(pos) != string(test.positionString) {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, test.positionString, pos)
		}
	}
}
