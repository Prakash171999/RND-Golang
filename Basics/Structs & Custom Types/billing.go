package main

type billing struct {
	name  string
	items map[string]float64
	tip   float64
}

//Create new bill
func newBill(name string, tip float64) billing {
	b := billing{
		name:  name,
		items: map[string]float64{},
		tip:   tip,
	}

	return b
}
