package json

import (
	"encoding/json"
	"log"
)

func Struct2json(s interface{}) string {
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Fatal("fail to marshal json", err)
	}
	return string(jsonData)
}

func Json2struct(j string, s interface{}) error {
	err := json.Unmarshal([]byte(j), s)
	if err != nil {
		log.Fatal("fail to unmarshal json", err)
	}
	return err
}
