package main

import "fmt"


func isOdd(slice []int) []int {
    var result []int
    for _, value := range slice {
        if value%2 != 0 {
            result = append(result, value)
        }
    }
    return result
}

func isEven(slice []int ) []int {
    var result []int
    for _, value := range slice {
        if value%2 == 0 {
            result = append(result, value)
        }
    }
    return result
}

// 宣告的函式型別在這個地方當做了一個參數
func main(){
    slice := []int {1, 2, 3, 4, 5, 7}
    fmt.Println("slice = ", slice)
    odd := isOdd(slice)    // 函式當做值來傳遞了
    fmt.Println("Odd elements of slice are: ", odd)
    even := isEven(slice)  // 函式當做值來傳遞了
    fmt.Println("Even elements of slice are: ", even)
}