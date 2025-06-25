package wms130

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUnmarshalYAMLPosition(t *testing.T) {

	var tests = []struct {
		positionstring   []byte
		expectedposition Position
	}{
		0: {positionstring: []byte(`2.52712538742158 50.2128625669452`), expectedposition: Position{2.52712538742158, 50.2128625669452}},
		1: {positionstring: []byte(`7.37402550506231 55.7211602557705`), expectedposition: Position{7.37402550506231, 55.7211602557705}},
		2: {positionstring: []byte(`7.37402550506231 55.7211602557705 0 1 2 3`), expectedposition: Position{7.37402550506231, 55.7211602557705}},
	}

	for k, test := range tests {
		var pos Position
		err := yaml.Unmarshal(test.positionstring, &pos)
		if err != nil {
			t.Errorf("test: %d, yaml.UnMarshal failed with '%s'\n", k, err)
		} else {
			if pos != test.expectedposition {
				t.Errorf("test: %d, expected: %v+,\n got: %v+", k, test.expectedposition, pos)
			}
		}
	}
}

func TestMarshalYAMLPosition(t *testing.T) {
	var tests = []struct {
		positionstring []byte
		position       Position
	}{
		0: {positionstring: []byte("2.52712538742158 50.2128625669452\n"), position: Position{2.52712538742158, 50.2128625669452}},
		1: {positionstring: []byte("7.37402550506231 55.7211602557705\n"), position: Position{7.37402550506231, 55.7211602557705}},
	}

	for k, test := range tests {
		pos, err := yaml.Marshal(test.position)
		if err != nil {
			t.Errorf("test: %d, yaml.Marshal failed with '%s'\n", k, err)
		} else {
			if string(pos) != string(test.positionstring) {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.positionstring, pos)
			}
		}
	}
}
