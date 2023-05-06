package main

import (
	"markets/wholefoods"
	"markets/traderjoes"
	"markets/shoppers"
	"github.com/pkg/errors"
)

// func BuyAvocados() (Ingredient, error)

// func BuyEggs() (Ingredient, error)

// func BuyBread() (Ingredient, error)

// type error interface {
// 	Error() string
// }

// func getIngredients() ([]Ingredient, error) {
// 	avocados, err := wholefoods.BuyAvocados()

// 	boiledEggs, err := traderjoes.BuyEggs()

// 	bread, err := shoppers.BuyBread()

// 	return []Ingredient{ avocados, boiledEggs, bread }, nil
// }

// func getIngredients() ([]Ingredient, error) {
// 	avocados, err := wholefoods.BuyAvocados()

// 	if err != nil {
// 		return nil, err
// 	}

// 	boiledEggs, err := traderjoes.BuyEggs()

// 	if err != nil {
// 		return nil, err
// 	}

// 	bread, err := shoppers.BuyBread()

// 	if err != nil {
// 		return nil, err
// 	}

// 	return []Ingredient{ avocados, boiledEggs, bread }, nil
// }

func getIngredients() ([]Ingredient, error) {
	avocados, err := wholefoods.BuyAvocados()

	if err != nil {
		return nil, errors.Wrap(err, "could not buy avocados")
	}

	boiledEggs, err := traderjoes.BuyEggs()

	if err != nil {
		return nil, errors.Wrap(err, "could not buy eggs")
	}

	bread, err := shoppers.BuyBread()

	if err != nil {
		return nil, errors.Wrap(err, "could not buy bread")
	}

	return []Ingredient{ avocados, boiledEggs, bread }, nil
}

func main() {
	ingredients, err := getIngredients()

	if err != nil {
		panic(err)
	}

	makeSandwich(ingredients)
}

// return fmt.Errorf("unique error message: %w", err)

// import "github.com/pkg/errors"

// return errors.Wrap(err, "unique error message")
