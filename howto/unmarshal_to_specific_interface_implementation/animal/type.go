package animal

import "fmt"

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
	Name string
}

func (w Wolf) GetNoise() string {
	return fmt.Sprintf("wooooo! my name is %s", w.Name)
}

func (w Wolf) Howl() {
	fmt.Println("woooooooo")
}

type Elephant struct {
	Name string
}

func (e Elephant) GetNoise() string {
	return fmt.Sprintf("uhmmmm! my name is %s", e.Name)
}
