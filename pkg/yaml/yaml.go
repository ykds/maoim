package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func DecodeFile(filePath string, v interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, v)
}
