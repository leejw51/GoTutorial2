package apple

import (
	"encoding/json"
	"fmt"
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

var jsonStr = `
{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World",
  "ok": {"1":"one","2":2}
}
`

func TestJson() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)
}