package main

import (
    "fmt"
    //"time"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "go.mongodb.org/mongo-driver/bson"
)
func main() {
    // ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // defer cancel()
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://35.236.156.68:27017"))
    defer client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("project").Collection("auto_increment")
    // res, err := collection.InsertOne(context.TODO(), bson.D{{"member_no", 5}, {"status", "waiting"}})
    // if err != nil {
    //     log.Fatal(err)
    // }
    // id := res.InsertedID
    // fmt.Println(id)

    // update := bson.M{
    //     "$set":bson.M{"email":"a@123"},
    // }
    // updateResult, err := collection.UpdateMany(context.TODO(), bson.M{"member_no":bson.M{"$ne": "0"}},update)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Fatal(updateResult)

    // type sunshareboy struct {
    //     Name string
    //     Age  int
    //     City string
    // }
    // dongdong:=sunshareboy{"張鼕鼕",29,"成都"}
    // huazai:=sunshareboy{"華仔",28,"深圳"}
    // suxin:=sunshareboy{"素心",24,"甘肅"}
    // god:=sunshareboy{"劉大仙",24,"杭州"}
    // qiaoke:=sunshareboy{"喬克",29,"重慶"}
    // jiang:=sunshareboy{"姜總",24,"上海"}
    // //插入多條資料要用到切片
    // boys:=[]interface{}{dongdong,huazai,suxin,god,qiaoke,jiang}
    // insertMany,err:= collection.InsertMany(ctx,boys)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println("Inserted multiple documents: ", insertMany.InsertedIDs)
    type Member struct {
        member_no   int
    }
    var result Member
    filter := bson.D{{"collection","member"}}
    err = collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found a single document: %+v\n", result)

}
