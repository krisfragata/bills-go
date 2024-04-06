package main

import "fmt"

// a blueprint for any bill. similar to a class
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	// create a var to store string, assign to 'Bill"
	fs := "Bill Breakdown: \n"
	// add a line for bill categories "item ... price"
	fs += fmt.Sprintf("%-25v %v \n", "items:", "price:")
	// store total
	var total float64 = 0
	// iterate through items and have item and price on same line
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%v \n", k+":", v)
		// increment total by appropriate price
		total += v
	}
	// print tip
	fs += fmt.Sprintf("%-25v ...%v\n", "Tip:", b.tip)
	// print total
	fs += fmt.Sprintf("%-25v ...%v", "Total:", total + b.tip)
	// return string
	return fs
}


// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add new items
func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}