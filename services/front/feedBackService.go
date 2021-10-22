package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "api/services"
)
type MemberFeedBack struct {
    Score int
    FeedBack string
    Type string
}
func SetMemberFeedBack(data MemberFeedBack,member_no int)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member_feedback")
    defer client.Disconnect(context.TODO())
    _,err := conn.InsertOne(context.TODO(), bson.M{"member_no":member_no,"type":data.Type,"score":data.Score,"feedback":data.FeedBack,"create_time":time.Now().Add(time.Hour * 8)})
 
    if (err==nil){
        result="success"
        ResultCode=0
    }else{
        result="error"
        ResultCode=5
    }
    return ResultCode,result
}