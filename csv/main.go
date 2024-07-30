package main

import (
	"fmt"
	"time"
)

func main() {
	in := []struct {
		Name       string    `csv:"name"`
		OtherNames []string  `csv:"other_names"`
		Worker     bool      `csv:"worker"`
		CreatedAt  time.Time `csv:"created_at"`
	}{
		{Name: "Ryan", Worker: true, CreatedAt: time.Now().AddDate(-20, -3, -5), OtherNames: []string{"John", "Jane"}},
		{Name: "John", OtherNames: []string{"Josh", "Ryan"}, Worker: false, CreatedAt: time.Now().AddDate(-4, -7, -5)},
		{OtherNames: []string{"Naomi", "Joseph"}, Worker: true, CreatedAt: time.Now().AddDate(-32, -4, -20), Name: "Emily"},
	}

	b, err := ToCsv(in)
	if err != nil {
		fmt.Printf("convert to csv: %v\n", err)
		return
	}
	fmt.Println(string(b))
}
