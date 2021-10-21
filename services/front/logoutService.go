package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "api/services"
)
func SetMemberLogout(accessToken string)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member_token")
    defer client.Disconnect(context.TODO())

    filter := bson.M{"token":accessToken}
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