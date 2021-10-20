package main

import "fmt"

type Bag struct {
	items []int
}

func Insert(b *Bag, itemId int) {
	fmt.Printf("address of *b: %p\n", b)
	b.items = append(b.items, itemId)
}

func InsertValue(b Bag, itemId int) Bag {
	fmt.Printf("address of b: %p\n", &b)
	b.items = append(b.items, itemId)
	return b
}

func main() {
	bag := new(Bag)
	fmt.Printf("address of bag: %p\n", bag)
	fmt.Println("新增元素前給ptr: ", bag)
	Insert(bag, 1000)

	fmt.Println("新增元素後給ptr: ", bag)

	bagValue := Bag{}
	fmt.Printf("address of bagValue: %p\n", bag)
	fmt.Println("新增元素前給實例前: ", bagValue)
	InsertValue(bagValue, 1001)
	fmt.Println("新增元素後, 但沒賦值回去: ", bagValue)
	bagValue = InsertValue(bagValue, 1001)
	fmt.Println("新增元素後, 有沒賦值回去: ", bagValue)
}