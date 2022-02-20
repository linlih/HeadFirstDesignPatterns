package pizza

import "fmt"

type IPizza interface {
	GetName() string
	SetName(name string)
	Prepare()
	Bake()
	Cut()
	Box()
	String() string
}

type Pizza struct {
	Name   string
	Dough  string
	Sauce  string
	Veggie []string
	Cheese string
	Clams  string
}

func (p *Pizza) GetName() string {
	return p.Name
}

func (p *Pizza) SetName(name string) {
	p.Name = name
}

func (p *Pizza) Prepare() {
	fmt.Println("Preparing " + p.Name)
}

func (p *Pizza) Bake() {
	fmt.Println("Baking " + p.Name)
}

func (p *Pizza) Cut() {
	fmt.Println("Cutting " + p.Name)
}

func (p *Pizza) Box() {
	fmt.Println("Boxing " + p.Name)
}

func (p *Pizza) String() string {
	var display string
	display += "----- " + p.Name + " -----\n"
	if p.Dough != "" {
		display += p.Dough + "\n"
	}
	if p.Sauce != "" {
		display += p.Sauce + "\n"
	}
	if p.Cheese != "" {
		display += p.Cheese + "\n"
	}
	if len(p.Veggie) != 0 {
		for i := 0; i < len(p.Veggie); i++ {
			display += p.Veggie[i] + "\n"
		}
	}
	if p.Clams != "" {
		display += p.Clams + "\n"
	}
	return display
}
