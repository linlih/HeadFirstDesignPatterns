package main

type PizzaStore struct {
	simplePizzaFactory *SimplePizzaFactory
}

func NewPizzaStore(simplePizzaFactory *SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{simplePizzaFactory: simplePizzaFactory}
}

func (p *PizzaStore) orderPizza(typ string) *Pizza {
	pizza := p.simplePizzaFactory.CreatePizza(typ)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
