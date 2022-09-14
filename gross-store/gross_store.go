package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unts := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return unts
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	rtnValue := false
	if val, ok := units[unit]; ok {
		if valItem, ok := bill[item]; ok {
			bill[item] = valItem + val
			rtnValue = true
		} else {
			bill[item] = val
			rtnValue = true
		}
	}
	return rtnValue
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	rtnValue := true

	upieces, uexists := units[unit]
	bpieces, bexists := bill[item]

	if !uexists || !bexists || upieces > bpieces {
		rtnValue = false
	} else if upieces == bpieces {
		delete(bill, item)
	} else {
		bill[item] -= upieces
	}

	return rtnValue
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	bpieces, bexists := bill[item]
	return bpieces, bexists
}
