package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "api/services"
)
type MemberInfo struct {
    Member_no int `json:"member_no"`
    Email string `json:"email"`
    Name string `json:"name"`
    Birthday string `json:"birthday"`
    State string `json:"state"`
    City string `json:"city"`
    Gender string `json:"gender"`
}
func GetMemberInfo(member_no int)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    //find
    var resultMemberInfo MemberInfo

    filter := bson.M{"member_no":member_no}
    conn.FindOne(context.TODO(), filter).Decode(&resultMemberInfo)
    if (MemberInfo{}!=resultMemberInfo){
        ResultCode=0
        result=resultMemberInfo
    }else{
        result="資料錯誤！"
        ResultCode=5
    }
    return ResultCode,result

}