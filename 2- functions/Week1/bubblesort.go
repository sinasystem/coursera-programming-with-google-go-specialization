package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func swap(sli []int, index int) {
	swapTemp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = swapTemp
}
func bubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}
func main() {
	dataSli := make([]int, 0, 10)
	var tempData int
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter an integer,(%d remaining)\n", 10-i)
		fmt.Scan(&tempData)
		dataSli = append(dataSli, tempData)
	}
	bubbleSort(dataSli)
	fmt.Printf("Sorted slice: %v", dataSli)

}
