package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if (index >= len(slice)) || index < 0 {
		return -1
	} else {
		return slice[index]
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if (index >= len(slice)) || index < 0 {
		slice = append(slice, value)
		return slice
	} else {
		slice[index] = value
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) > 0 {
		var slices []int

		for _, value := range values {
			slices = append(slices, value)
		}
		slices = append(slices, slice...)
		return slices
	} else {
		return slice
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
}
