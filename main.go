package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	intValue := []int{5, 33, 15, 90, 66, 124, 55, 666}
	sort.Ints(intValue)
	fmt.Println("Sorted ints:   ", intValue)

	floatValue := []float64{7.5, 8.5, 5.5, 22.5, 44.5, 88.5, 89.9, 89.8, 89.95}
	sort.Float64s(floatValue)
	fmt.Println("Sorted floats:   ", floatValue)

	stringValue := []string{"C", "D", "A", "B", "F", "Z"}
	sort.Strings(stringValue)
	fmt.Println("Sorted strings:", stringValue)

	myMap := map[string]int{"variable A": 1, "variable B": 2}
	json, _ := json.Marshal(myMap)
	fmt.Println(string(json))

}
