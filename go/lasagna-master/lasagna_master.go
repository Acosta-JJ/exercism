package lasagna

func PreparationTime(layers []string, avg int) int {
	if avg <= 0 {
		avg = 2
	}
	return len(layers) * avg
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(recipe, layers []string) {
	layers[len(layers)-1] = recipe[len(recipe)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	needs := make([]float64, len(amounts))

	for index, amount := range amounts {
		needs[index] = amount * float64(portions) / 2
	}

	return needs
}

