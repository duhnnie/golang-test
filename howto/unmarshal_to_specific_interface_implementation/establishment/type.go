package establishment

import (
	"encoding/json"
	"example/unmarshal_interface/animal"
	"fmt"
	"reflect"
)

var knownAnimals = map[string]reflect.Type{
	"lion":     reflect.TypeOf(animal.Lion{}),
	"wolf":     reflect.TypeOf(animal.Wolf{}),
	"elephant": reflect.TypeOf(animal.Elephant{}),
}

type Establishment interface {
	GetAddress() string
}

type Zoo struct {
	Name    string
	Address string
	Animals []animal.Animal
}

func (z *Zoo) GetAddress() string {
	return z.Address
}

func (z *Zoo) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Animals []json.RawMessage
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	z.Name = aux.Name
	z.Address = aux.Address

	for _, animalData := range aux.Animals {
		m := map[string]interface{}{}
		json.Unmarshal(animalData, &m)

		var a animal.Animal

		if animalType, exists := m["type"]; !exists {
			return fmt.Errorf("missing property: type")
		} else if animalType, ok := animalType.(string); !ok {
			return fmt.Errorf("property \"type\" needs to be string")
		} else if animalStructType, exists := knownAnimals[animalType]; !exists {
			return fmt.Errorf("unknown animal: %s", animalType)
		} else {
			// TODO: handle possible error
			a = reflect.New(animalStructType).Interface().(animal.Animal)
		}

		if err := json.Unmarshal(animalData, &a); err != nil {
			return err
		}

		z.Animals = append(z.Animals, a)
	}

	return nil
}
