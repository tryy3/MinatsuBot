package permissionmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	file string
}

func (c *config) load(m interface{}) error {
	file, err := c.create()

	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(b, m)
	return err
}

func (c *config) save(m interface{}) error {
	file, err := c.create()

	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(m, "", "    ")

	if err != nil {
		return err
	}

	file.Write(bytes)
	return nil
}

func (c *config) create() (*os.File, error) {
	file, err := os.Open(c.file)
	if os.IsNotExist(err) {
		file, err = os.Create(c.file)
		if err != nil {
			return nil, err
		}
	}
	return file, nil
}
