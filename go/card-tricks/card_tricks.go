package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index <= len(slice) && index >= 0 {
        return slice[index]
    } else {
        return int(-1)
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
        slice[index] = value
        return slice
    } else {
        slice = append(slice, value)
        return slice
    }
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if values != nil {
        slice = append(values,slice...)
        return slice
    } else {
        return slice
    }
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index > len(slice) || index < 0 {
        return slice
    } else {
    	if index < len(slice)-1 {
        	slice = append(slice[:index], slice[index+1:]...)
            return slice
    	} else {
        	slice = slice[:index]
            return slice
    	}
    }
}
