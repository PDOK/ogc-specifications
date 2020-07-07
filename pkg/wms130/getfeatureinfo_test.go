package wms130

import "testing"

func TestGetFeatureInfoType(t *testing.T) {
	dft := GetFeatureInfo{}
	if dft.Type() != `GetFeatureInfo` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetFeatureInfo`, dft.Type())
	}
}
