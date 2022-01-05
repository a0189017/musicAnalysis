package adminServices

import (
	"api/services"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MemberLogin struct {
	Email    string
	Password string
}
type memberLoginInfo struct {
	Member_no int
}

func GetAdminLogin(data MemberLogin) (ResultCode int, result interface{}) {
	conn, client := services.DbConnect("project", "member")
	conn_manager, client_manager := services.DbConnect("project", "manager")
	defer client.Disconnect(context.TODO())
	defer client_manager.Disconnect(context.TODO())
	//find
	var resultMemberInfo memberLoginInfo
	var resuletManager memberLoginInfo
	//memeber
	filter := bson.M{"email": data.Email, "password": services.PasswordMd5(data.Password)}
	conn.FindOne(context.TODO(), filter).Decode(&resultMemberInfo)
	//manager
	filter = bson.M{"member_no": resultMemberInfo.Member_no}
	conn_manager.FindOne(context.TODO(), filter).Decode(&resuletManager)

	if (memberLoginInfo{} != resuletManager) {
		today := time.Now()
		today_string := time.Now().Format("2006-01-02 15:04:05")
		result = services.PasswordMd5(data.Email + today_string + "project@#$S_ecrEt")
		ResultCode = 0
		conn_token, client_token := services.DbConnect("project", "member_token")
		defer client_token.Disconnect(context.TODO())
		conn_token.InsertOne(context.TODO(), bson.M{"token": result, "member_no": resultMemberInfo.Member_no, "expire_time": today.AddDate(0, 0, 7).Unix()})
	} else {
		result = "帳號/密碼錯誤！"
		ResultCode = 5
	}
	return ResultCode, result

}
