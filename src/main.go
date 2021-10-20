package main

import (
    "fmt"
    "time"
    "net/http"
    "strings"
    "log"
    "regexp"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
)
func main() {
    http.HandleFunc("/", sayhelloName) //設定存取的路由
    err := http.ListenAndServe(":9090", nil) //設定監聽的埠
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析參數，預設是不會解析的
    fmt.Println(r.Form)  //這些資訊是輸出到伺服器端的列印資訊
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    s := strings.Join(r.Form["name"], "")
    re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    s = re.ReplaceAllString(s, "")
    fmt.Fprintf(w,s)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://35.236.156.68:27017"))
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("My_first_mongo").Collection("end")
    res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
    if err != nil {
        log.Fatal(err)
    }
    id := res.InsertedID
    fmt.Println(id)
    // id := res.InsertedID
    // fmt.Fprintf(w,id)



}
