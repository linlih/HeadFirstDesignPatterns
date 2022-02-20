package ingredient

type IIngredientFactory interface {
	CreateDough() string
	CreateSauce() string
	CreateCheese() string
	CreateClams() string
	CreateVeggie() []string
	CreatePepperoni() string
}
