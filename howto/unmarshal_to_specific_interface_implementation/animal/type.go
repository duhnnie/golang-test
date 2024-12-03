package animal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Animal interface {
	GetNoise() string
}

type Lion struct {
	Name string
}

func (l Lion) GetNoise() string {
	return fmt.Sprintf("roarrr! my name is %s", l.Name)
}

type Wolf struct {
	Name string `json:"name" required:"true"`
	Age  int    `json:"age"`
}

func (w Wolf) GetNoise() string {
	return fmt.Sprintf("wooooo! my name is %s", w.Name)
}

func (w Wolf) Howl() {
	fmt.Printf("woooooooo! my age is %d\n", w.Age)
}

func (w *Wolf) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	vp := reflect.ValueOf(w)
	tp := reflect.TypeOf(w)
	v := vp.Elem()
	t := tp.Elem()

	for i := 0; i < t.NumField(); i += 1 {
		field := t.Field(i)
		tag := field.Tag
		required := tag.Get("required") == "true"
		jsonField := tag.Get("json")

		if fieldValue, exists := m[jsonField]; !exists && required {
			return fmt.Errorf("\"%s\" field is required", jsonField)
		} else if exists {
			if field.Type.Kind() == reflect.Int && reflect.TypeOf(fieldValue).Kind() == reflect.Float64 {
				v.Field(i).SetInt(int64(fieldValue.(float64)))
			} else {
				v.Field(i).Set(reflect.ValueOf(fieldValue))
			}
		}
	}

	return nil
}

type Elephant struct {
	Name string
}

func (e Elephant) GetNoise() string {
	return fmt.Sprintf("uhmmmm! my name is %s", e.Name)
}
