package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 2.99, "tea": 5.25},
		tip:   1,
	}
	return b
}

func (b bill) format() string {
	fs := "Breakdown of the Bill\n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...........$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%v...........$%v\n", "Tip :", b.tip)
	fs += fmt.Sprintf("%v...........$%v\n", "Total", total+b.tip)
	return fs

}
func (b *bill) additems(name string, price float64) {
	b.items[name] = price
}
func (b *bill) update(tip float64) {
	b.tip += tip
}
