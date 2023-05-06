package main

import (
	"markets/wholefoods"
	"markets/traderjoes"
	"markets/shoppers"
	"github.com/pkg/errors"
)

func getIngredients() ([]Ingredient, error) {

}

func main() {
	ingredients, err := getIngredients()

	if err != nil {
		panic(err)
	}

	makeSandwich(ingredients)
}
