package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "log"
    "time"
    "api/services"
)
type MemberRecord struct {
    Type string
    Playsec int
}
func SetMemberRecord(data MemberRecord,member_no int)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member_record")
    defer client.Disconnect(context.TODO())
    _,err := conn.InsertOne(context.TODO(), bson.M{"member_no":member_no,"type":data.Type,"playmin":data.Playsec,"date":time.Now().Format("2006-01-02"),"create_time":time.Now().Add(time.Hour * 8)})
 
    if (err==nil){
        result="success"
        ResultCode=0
    }else{
        result="error"
        ResultCode=5
    }
    log.Println(err)
    return ResultCode,result
}