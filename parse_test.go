package nginx

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestReadFileFromYaml(t *testing.T) {
	fs := fstest.MapFS{
		"nginx.yaml": {Data: []byte("map[a:Easy! b:map[c:2 d:[3 4]]]")},
	}

	data := ParseYaml(fs, "nginx.yaml")
	expectedData := fs["nginx.yaml"].Data

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("got %d wanted %d", data, expectedData)
	}
}
