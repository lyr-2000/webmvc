package strutil

import "encoding/json"

func AsJsonIndent(v interface{}) string {
	bs, _ := json.MarshalIndent(v, "", " ")
	return string(bs)
}
