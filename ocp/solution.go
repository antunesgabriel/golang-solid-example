package main

import "log"

// of course all this would be in separate files

type Food interface {
	Eat() string
}

type Recipe interface {
	Prepare() Food
}

type Coxinha struct{}
type CoxinhaRecipe struct{}

func (c *Coxinha) Eat() string {
	return "eat coxinha"
}

func (cr *CoxinhaRecipe) Prepare() Food {
	return &Coxinha{}
}

type Pudding struct{}
type PuddingRecipe struct{}

func (p *Pudding) Eat() string {
	return "eat pudding"
}

func (pr *PuddingRecipe) Prepare() Food {
	return &Pudding{}
}

type FoodMachine struct {
}

func (fm *FoodMachine) PrepareRecipe(recipe Recipe) Food {
	return recipe.Prepare()
}

func main() {
	foodMachine := new(FoodMachine)
	coxinhaRecipe := CoxinhaRecipe{}
	puddingRecipe := PuddingRecipe{}

	coxinha := foodMachine.PrepareRecipe(&coxinhaRecipe)
	pudding := foodMachine.PrepareRecipe(&puddingRecipe)

	log.Println("Coxinha:", coxinha.Eat())
	log.Println("Pudding:", pudding.Eat())
}
