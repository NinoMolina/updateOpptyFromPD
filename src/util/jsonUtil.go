package util

import "encoding/json"

func Unmarshal(jsonAsString string) map[string]interface{} {
	in := []byte(jsonAsString)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	return raw
}


func Marshal(jsonMap map[string]interface{}) string {
	out, _ := json.Marshal(jsonMap)
	return string(out)
}

func ToJson(p interface{}) []byte {
	bytes, err := json.Marshal(p)
	CheckErr(err, "")
	return bytes
}

func FromJson(jsonAsString []byte) interface{}{
	var raw interface{}
	err := json.Unmarshal(jsonAsString, &raw)
	CheckErr(err, "")
	return raw
}
