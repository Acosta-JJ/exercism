package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
    return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen":    6,
        "dozen":              12,
        "small_gross":        120,
        "gross":              144,
        "great_gross":        1728,
    }
}

// NewBill creates a new bill
func NewBill() map[string]int {
    return map[string]int{}
}

// AddItem adds item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
    pieces, exists := units[unit]
    if !exists {
        return false
    }

    bill[item] += pieces
    return true
}

// RemoveItem removes item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    rpieces, rexists := units[unit]
    bpieces, bexists := bill[item]

    if !rexists || !bexists || rpieces > bpieces {
        return false
    } else if rpieces == bpieces {
        delete(bill, item)
    } else {
        bill[item] -= rpieces
    }

    return true
}

// GetItem returns the quantity of item that the customer has in their bill
func GetItem(bill map[string]int, item string) (int, bool) {
    pieces, exists := bill[item]
    return pieces, exists
}


