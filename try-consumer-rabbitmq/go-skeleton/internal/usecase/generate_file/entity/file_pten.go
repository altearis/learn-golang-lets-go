package entity

import "encoding/json"

type FilePten struct {
	Filename  string `json:"filename"`
	Container string `json:"container"`
}

func (c *FilePten) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}
