package adminServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "api/services"
    "strconv"

)
type managerNo struct {
    Member_no   int
}
type managerInfo struct {
    Member_no int
    Email   string
}
type ManagerEmail struct {
    Email   string
}
func GetManagerList()(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "manager")
    conn_member,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())

    var resultManagerNo []managerNo
    var resultManagerInfo managerInfo
    var resultManagerList []managerInfo

    sortStage := bson.D{{"$sort", bson.D{{"member_no", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{sortStage})
    resultInfo.All(context.TODO(), &resultManagerNo)

    for _,v:=range resultManagerNo{
        filter := bson.M{"member_no":v.Member_no}
        conn_member.FindOne(context.TODO(), filter).Decode(&resultManagerInfo)
        resultManagerList = append(resultManagerList, resultManagerInfo)
    }
 
    ResultCode=0
    return ResultCode,resultManagerList
}
func DelManager(data string)(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "manager")
    defer client.Disconnect(context.TODO())

    member_no, _ := strconv.ParseInt(data, 10, 64) 
    filter := bson.M{"member_no":member_no}
    err :=conn.FindOneAndDelete(context.TODO(), filter).Err()
    if (err==nil){
        ResultCode=0
        result="success"
    }else{
        result="error"
        ResultCode=5
    }

    return ResultCode,result

}
func SetManager(email ManagerEmail)(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "manager")
    conn_member,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    var memberNo managerNo
    var err interface{}
    filter := bson.M{"email":email.Email}
    conn_member.FindOne(context.TODO(), filter).Decode(&memberNo)
    if(memberNo.Member_no!=0){
        _,err = conn.InsertOne(context.TODO(), bson.M{"member_no":memberNo.Member_no})
    }
    if (err==nil && memberNo.Member_no!=0){
        result="success"
        ResultCode=0
    }else{
        result="error"
        ResultCode=5
    }
    return ResultCode,result
}