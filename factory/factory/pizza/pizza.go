package pizza

import "fmt"

type IPizza interface {
	GetName() string
	Prepare()
	Bake()
	Cut()
	Box()
	String() string
}

type Pizza struct {
	Name    string
	Dough   string
	Sauce   string
	Topping []string
}

func NewPizza(name string, dough string, sauce string, topping []string) *Pizza {
	return &Pizza{Name: name, Dough: dough, Sauce: sauce, Topping: topping}
}

func (p *Pizza) GetName() string {
	return p.Name
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
	display += p.Dough + "\n"
	display += p.Sauce + "\n"
	for i := 0; i < len(p.Topping); i++ {
		display += p.Topping[i] + "\n"
	}
	return display
}
