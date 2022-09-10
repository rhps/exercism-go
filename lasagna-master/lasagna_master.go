package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgtime int) int {
	if avgtime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgtime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesneed := 0
	sauces := 0.
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesneed = noodlesneed + 50

		} else if layer == "sauce" {
			sauces = sauces + 0.2
		}
	}
	return noodlesneed, float64(sauces)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var scaled []float64
	for _, quantity := range quantities {
		scaled = append(scaled, quantity*(float64(portion)/2))
	}

	return scaled
}
