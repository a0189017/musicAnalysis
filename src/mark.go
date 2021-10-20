package main

import (
	"fmt"
	"net/http"

)
type Element interface{} //空interface可以存任何型態變數
type List [] Element //宣告為 Element 型別

func myWeb(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //它還將請求主體解析爲表單，獲得POST Form表單數據，必須先調用這個函數

	for k, v := range r.URL.Query() {
		fmt.Fprintln(w,"key:", k, ", value:", v[0])
	}

	for k, v := range r.PostForm {
		fmt.Fprintln(w,"key:", k, ", value:", v[0])
	}
	var c  = "123"
	//output: (5+5i)
	fmt.Fprintln(w, c)
}
func max(a int, b int) int{
    if a > b {
        return a
    }
    return b
}
func main() {
	//http.HandleFunc("/", myWeb)

	// fmt.Println("服務器即將開啓，訪問地址 http://localhost:8080")

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("服務器開啓錯誤: ", err)
	// }

	//二維陣列
	yourMap := make(map[int]map[string]string)
	yourMap[0] = make(map[string]string)
	yourMap[0]["test"]="abc"

	//foreach
	for _,v :=range yourMap{
		fmt.Println(v["test"])
	}
	//switch
	integer := 6
	switch integer {
	case 4:
	    fmt.Println("The integer was <= 4")
	    fallthrough //不使用break
	case 5:
	    fmt.Println("The integer was <= 5")
	case 6:
	    fmt.Println("The integer was <= 6")
	case 7:
	    fmt.Println("The integer was <= 7")
	case 8:
	    fmt.Println("The integer was <= 8")
	default:
	    fmt.Println("default case")
	}
	//呼叫函數
	x := 3
    y := 4
    max_xy := max(x, y) //呼叫函式 max(x, y)
    fmt.Println(max_xy)
    //defer 後指定的函式會在函式退出前呼叫
	for i := 0; i < 5; i++ {
    	defer fmt.Printf("%d ", i)
	}

}
