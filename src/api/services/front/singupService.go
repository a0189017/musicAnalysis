package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "log"
    "api/services"
    "time"
)
type MemberSingup struct {
    Email   string
    Password string
    Name string
    Birthday string
    Country string
    State string
    City string
    Gender string
}
func SetMemberSingup(data MemberSingup)(ResultCode int,result string){
	conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    //find
    var resultIDExist MemberSingup
    error_staus :=0
    filter := bson.M{"email":data.Email}
    conn.FindOne(context.TODO(), filter).Decode(&resultIDExist)
    if resultIDExist.Email!=""{
        result="帳號重複！"
        error_staus=1
        ResultCode=5
    }

    //insert
    if(error_staus==0){
        member_no :=memberNoIncrement()
        _, err := conn.InsertOne(context.TODO(), bson.D{{"member_no", member_no}, {"email", data.Email},{"password",services.PasswordMd5(data.Password)},{"name",data.Name},{"birthday",data.Birthday},{"country",data.Country},{"state",data.State},{"city",data.City},{"gender",data.Gender},{"create_time",time.Now().Add(time.Hour * 8)}})
        if err != nil{
            result="格式錯誤！"
            log.Println(err)
            ResultCode=5
        }else{
            result="success"
            ResultCode=0
        }
    }
    return ResultCode,result

}
func memberNoIncrement() int{
    conn,client :=services.DbConnect("project", "auto_increment")
    defer client.Disconnect(context.TODO())
    update := bson.M{
        "$inc":bson.M{"member_no":1},
    }
    conn.UpdateOne(context.TODO(), bson.M{"collection":"member"},update)

    type MemberNo struct {
        Member_no   int
    }
    var result MemberNo
    filter := bson.M{"collection":"member"}
    err := conn.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Println(err)
    }
    return result.Member_no
}