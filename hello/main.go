package main

// "strings"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func option(b bill) {
	for {
		reader := bufio.NewReader(os.Stdin)
		opt, _ := getInput("a) add item\n b)display bill\n c)update tip \n", reader)
		switch opt {
		case "a":
			name1, _ := getInput("Item Name:	", reader)
			price, _ := getInput("Item price:	", reader)
			fmt.Println(name1, "  ", price)
			// b.additems(name1, (price))
		case "b":
			fmt.Println(b)
		case "c":
			tip, _ := getInput("Enter  the tip amount\n", reader)
			fmt.Println(tip)

		default:
			fmt.Println("invalid choices\n ")
		}
	}

}
func createbill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter the name  \n", reader)
	return newbill(name)

}

func main() {
	// name := "augustin"
	// age := 23

	// var name string = "hello   World"
	// // fmt.Println("My name is ", name, "and my age is ", age)
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("The value  is %v \n ", i)
	// }
	// name := []string{"raj", "shiv", "rohit", "sam", "joy", "ajay"}
	// for index, value := range name {
	// 	fmt.Printf("The value at the index %v is %v \n", index, value)
	// }

	// 	menu := map[string]float64{
	// 		"soup":   3.5,
	// 		"pie":    5.25,
	// 		"snacks": 7.99,
	// 		"tea":    6.99,
	// 	}

	// 	fmt.Println(menu)
	// 	for k, v := range menu {
	// 		fmt.Println(k, "--", v)

	// }
	// bill := newbill("marios bill")
	bill := createbill()
	// bill.update(10)
	// bill.additems("snacks", 7.99)
	// bill.additems("coffe", 2.50)

	// fmt.Println(bill.format())
	option(bill)
}
