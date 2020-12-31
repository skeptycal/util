package gob

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Encode(e *gob.Encoder, data map[string]interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	m := make(map[string]interface{})
	m["foo"] = "bar"

	if err := enc.Encode(m); err != nil {
		log.Fatal(err)
	}
	return nil
}
