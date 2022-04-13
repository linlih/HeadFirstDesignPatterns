package main

type ICaffeineBeveragePrepare interface {
	Brew()
	AddCondiments()
	PrepareRecipe(prepare ICaffeineBeveragePrepare)
}

type ICaffeineBeveragePrepareWithHook interface {
	Brew()
	AddCondiments()
	PrepareRecipe(prepare ICaffeineBeveragePrepareWithHook)
	CustomerWantsCondiments() bool
}
