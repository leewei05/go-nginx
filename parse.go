package nginx

import (
	"io/fs"
)

func ParseYaml(fileSystem fs.FS, name string) []byte {
	data, _ := fs.ReadFile(fileSystem, name)
	return data
}
