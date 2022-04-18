package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Billing struct {
	Name              string
	Phone, PostalCode string
	PaymentDue        int
}

func main() {
	// name := "Scott"
	//make memory a pointer to name, memory address to name
	// memory := &name
	//astrict = value change of memory address (name = mark)
	// memory = "Mark"
	// fmt.Println(name)

	b := []Billing{}

	allJsonFiles, err := filepath.Glob(`\\qumulo\BLIS\ScottGetzFiles\OVER--1000\*.json`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, js := range allJsonFiles {
		fileContents, err := os.ReadFile(js)
		if err != nil {
			fmt.Println(err)
			continue
		}

		structBill := Billing{}

		err = json.Unmarshal(fileContents, &structBill)
		if err != nil {
			fmt.Println(err)
			continue
		}

		b = append(b, structBill)
	}

	var allPipes string

	for _, bill := range b {
		if bill.PaymentDue < 1000 {
			continue
		}

		allPipes = allPipes + fmt.Sprintf("%v|%v|%v\n", bill.Name, bill.Phone, bill.PaymentDue)
	}

	fmt.Println(allPipes)

	// fileData, err := os.ReadFile(`\\qumulo\BLIS\ScottGetzFiles\OVER--1000\AaronMccarty.json`)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// err = json.Unmarshal(fileData, &b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
