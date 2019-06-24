package control

import (
	"encoding/json"
	"io/ioutil"
	"io"
)

func JsonStruct(v interface{}, body io.ReadCloser) (err error) {

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, err := ioutil.ReadAll(body)
	//validating parsing
	if err != nil {
		return
	}
	//Matching Json with struct
	err = json.Unmarshal(bytes, v)
	return
}
