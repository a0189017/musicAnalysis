package adminServices

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "api/services"
    "time"
    "strconv"
)
type memberInfo struct {
    Id string `bson:"_id"` 
    Member_no int
    Email string
    Name string
    Birthday string
    State string
    City string
    Gender string
    Create_time time.Time
}
type memberRecord struct{
    Type string
    Playmin int
    Create_time time.Time
}
type memberFeedback struct{
    Score int
    Feedback string
    Create_time time.Time
}
type MemberNo struct{
    Member_no int
}
type memberDetail struct{
    Info []*memberInfo
    Record []*memberRecord 
    Feedback []*memberFeedback 
}
func GetMemberList()(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())

    var resultMemberList []*memberInfo

    sortStage := bson.D{{"$sort", bson.D{{"member_no", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{sortStage})
    resultInfo.All(context.TODO(), &resultMemberList)
    ResultCode=0

    return ResultCode,resultMemberList
}
func GetMemberDetail(data map[string]string)(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "member")
    conn_reord,client :=services.DbConnect("project", "member_record")
    conn_feedback,client :=services.DbConnect("project", "member_feedback")
    defer client.Disconnect(context.TODO())
    member_no, _ := strconv.ParseInt(data["member_no"], 10, 64) 
    matchStage := bson.D{{"$match", bson.D{{"member_no", member_no}}}}
    //memberInfo
    var resultMemberInfo []*memberInfo
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
    resultInfo.All(context.TODO(), &resultMemberInfo)
    //memberRecord
    var resultMemberRecord []*memberRecord
    sortStage := bson.D{{"$sort", bson.D{{"create_time", -1}}}}
    resultInfo,_ =conn_reord.Aggregate(context.TODO(), mongo.Pipeline{matchStage,sortStage})
    resultInfo.All(context.TODO(), &resultMemberRecord)
    //memberFeedback
    var resultMemberFeedback []*memberFeedback
    sortStage = bson.D{{"$sort", bson.D{{"create_time", -1}}}}
    resultInfo,_ =conn_feedback.Aggregate(context.TODO(), mongo.Pipeline{matchStage,sortStage})
    resultInfo.All(context.TODO(), &resultMemberFeedback)

    ResultCode=0
    result =memberDetail{resultMemberInfo,resultMemberRecord,resultMemberFeedback}

    return ResultCode,result

}
