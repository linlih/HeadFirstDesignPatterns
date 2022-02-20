package pizza

type PizzaStore struct {
}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{}
}

func (p *PizzaStore) OrderPizza(pizza IPizza) IPizza {
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
