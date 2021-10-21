package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "api/services"
)
type MemberLogin struct {
    Email   string
    Password string
}
type memberLoginInfo struct {
    Member_no int
}
func GetMemberLogin(data MemberLogin)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    //find
    var resultMemberInfo memberLoginInfo
    //error_staus :=0
    filter := bson.M{"email":data.Email,"password":services.PasswordMd5(data.Password)}
    conn.FindOne(context.TODO(), filter).Decode(&resultMemberInfo)
    if (memberLoginInfo{}!=resultMemberInfo){
        today :=time.Now()
        result=services.PasswordMd5(data.Email+today.String()+"project@#$S_ecrEt")
        ResultCode=0
        conn_token,client_token :=services.DbConnect("project", "member_token")
        defer client_token.Disconnect(context.TODO())
        conn_token.InsertOne(context.TODO(), bson.M{"token":result,"member_no":resultMemberInfo.Member_no,"expire_time":today.AddDate(0, 0, 7).Unix()})
    }else{
        result="帳號/密碼錯誤！"
        ResultCode=5
    }
    return ResultCode,result

}