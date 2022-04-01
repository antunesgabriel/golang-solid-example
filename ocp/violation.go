//go:build violation
// +build violation

package ocp

type FoodMachine struct {
}

func (fm *FoodMachine) PrepareRecipe(recipe any) any {
	if recipe.Name == "pudding" {
		// prepare pudding and return pudding
	}

	if food == "coxinha" {
		// prepare coxinha and return coxinha
	}

	// how far could that go with each new recipe?
}
