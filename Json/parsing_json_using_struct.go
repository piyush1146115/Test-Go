package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myJSON2 struct {
	IntValue        int        `json:"intValue"`
	BoolValue       bool       `json:"boolValue"`
	StringValue     string     `json:"stringValue"`
	DateValue       time.Time  `json:"dateValue"`
	ObjectValue     *myObject2 `json:"objectValue"`
	NullStringValue *string    `json:"nullStringValue,omitempty"`
	NullIntValue    *int       `json:"nullIntValue"`
	EmptyString     string     `json:"emptyString,omitempty"`
}

type myObject2 struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	jsonData := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null,
			"extraValue":4321
		}
	`

	var data *myJSON2
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json struct: %#v\n", data)
	fmt.Printf("dateValue: %#v\n", data.DateValue)
	fmt.Printf("objectValue: %#v\n", data.ObjectValue)
}
