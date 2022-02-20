package pizzastore

import "abstractfactory/pizza"

type PizzaStore struct {
}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{}
}

func (p *PizzaStore) OrderPizza(pizza pizza.IPizza) pizza.IPizza {
	// TODO： 这里没有办法去调用具体工厂的CreatePizza，所以将Pizza传入了
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
