package main

import (
	"encoding/json"
	"example/unmarshal_interface/animal"
	"example/unmarshal_interface/establishment"
	"fmt"
)

// type Something interface{}

// type Something1 struct {
// 	Aaa, Bbb string
// }

// type Something2 struct {
// 	Ccc, Ddd string
// }

// type Container struct {
// 	Type  string    `json:"type"`
// 	Value Something `json:"value"`
// }

// func (c *Container) UnmarshalJSON(data []byte) error {
// 	m := map[string]interface{}{}
// 	json.Unmarshal(data, &m)

// 	c.Type = m["type"].(string)

// 	value, err := UnmarshalCustomValue(data, "type", "value", map[string]reflect.Type{
// 		"something1": reflect.TypeOf(Something1{}),
// 		"something2": reflect.TypeOf(Something2{}),
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	c.Value = value
// 	return nil
// }

// func UnmarshalCustomValue(data []byte, typeField, valueField string, customTypes map[string]reflect.Type) (interface{}, error) {
// 	m := map[string]interface{}{}
// 	json.Unmarshal(data, &m)

// 	typeString, ok := m[typeField].(string)

// 	if !ok {
// 		return nil, fmt.Errorf("field type %s doesn't exist or can't converted into string", typeField)
// 	}

// 	t, typeExists := customTypes[typeString]

// 	if !typeExists {
// 		return nil, fmt.Errorf("type for %s doesn't exists", typeString)
// 	}

// 	s := reflect.New(t).Interface().(Something)
// 	data, _ = json.Marshal(m[valueField])
// 	err := json.Unmarshal(data, &s)

// 	return &s, err
// }

// func panicIfErr(err error) {
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// func testUnmarshalling(jsonString string) {
// 	var container Container
// 	err := json.Unmarshal([]byte(jsonString), &container)
// 	panicIfErr(err)
// 	fmt.Printf("container=%+v\n", container)
// 	fmt.Printf("value=%+v\n", container.Value)
// }

// func main() {
// 	testUnmarshalling(`{"type":"something1","value":{"Aaa": "a"}}`)
// 	testUnmarshalling(`{"type":"something2","value":{"Ccc": "a"}}`)
// }

func main() {
	var zoo establishment.Zoo

	jsonString := `
	{
		"name": "Zooland",
		"address": "Ocean ave. #543",
		"animals": [
			{
				"type": "wolf",
				"name": "wolfie"
			},
			{
				"type": "lion",
				"name": "Kimba"
			}
		]
	}
	`

	if err := json.Unmarshal([]byte(jsonString), &zoo); err != nil {
		panic(err.Error())
	}

	fmt.Println(zoo.Name)
	fmt.Println(zoo.GetAddress())

	for _, a := range zoo.Animals {
		var atype string

		switch t := a.(type) {
		case *animal.Wolf:
			atype = "Wolf"
			t.Howl()
		case *animal.Lion:
			atype = "Lion"
		case *animal.Elephant:
			atype = "Elephant"
		default:
			atype = fmt.Sprintf("unknown type %T", t)
		}

		fmt.Printf("%s: %s\n", atype, a.GetNoise())
	}
}
