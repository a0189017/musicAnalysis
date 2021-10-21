package adminServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/bson/primitive"
    "time"
    "api/services"

)
type memberRecordSum struct {
    Type string `bson:"_id" json:"type"` 
    Count int `bson:"count" json:"count"`
    Sum int `bson:"sum" json:"sum"`
}
type memberRecordEvery struct {
    Date string `bson:"_id" json:"date"`
    Count int `bson:"count" json:"count"`
    Sum int `bson:"sum" json:"sum"`
}
func GetAnalysisUsetime(data map[string]string)(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "member_record")
    defer client.Disconnect(context.TODO())
    type memberRecord struct{
        Type string `bson:"_id" json:"type"`
    }
    type memberRecordSum struct {
        Type string `bson:"_id" json:"type"` 
        Count int `bson:"count" json:"count"`
        Sum int `bson:"sum" json:"sum"`
    }
    type memberRecordEvery struct {
        Date string `bson:"_id" json:"date"`
        Count int `bson:"count" json:"count"`
        Sum int `bson:"sum" json:"sum"`
    }
    type memberRecordEveryType struct {
        Group bson.M `bson:"_id" json:"group"`
        Count int `bson:"count" json:"count"`
        Sum int `bson:"sum" json:"sum"`
    }
    type analysisResult struct{
        OnlineCount int
        ReocrdName []memberRecord
        ResultEvery []memberRecordEvery
        ResultSum []memberRecordSum
        ResultEveryType []memberRecordEveryType
    }
    startdate, _ := time.Parse(time.RFC3339, data["startdate"]+"T00:00:00.000Z")
    enddate, _ := time.Parse(time.RFC3339, data["enddate"]+"T00:00:00.000Z")
    //dbResultRecordName
    var dbResultRecordName []memberRecord
    groupStage := bson.D{{"$group", bson.D{{"_id", "$type"}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
    resultInfo.All(context.TODO(), &dbResultRecordName)
    //onlinecount
    onlineCount :=getOnlieCount()
    //dbResultEveryday
    var dbResultEvery []memberRecordEvery
    matchStage := bson.D{{"$match", bson.D{{"create_time",bson.M{"$gt":startdate}},{"create_time",bson.M{"$lte":enddate}}}}}
    groupStage = bson.D{{"$group", bson.D{{"_id", "$date"}, {"sum", bson.D{{"$sum", "$playmin"}}}, {"count", bson.D{{"$sum", 1}}}}}}
    sortStage := bson.D{{"$sort", bson.D{{"_id", 1}}}}
    resultInfo,_ =conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage,sortStage})
    resultInfo.All(context.TODO(), &dbResultEvery)
    //dbResultSum
    var dbResultSum []memberRecordSum
    matchStage = bson.D{{"$match", bson.D{{"create_time",bson.M{"$gt":startdate}},{"create_time",bson.M{"$lte":enddate}}}}}
    groupStage = bson.D{{"$group", bson.D{{"_id", "$type"}, {"sum", bson.D{{"$sum", "$playmin"}}}, {"count", bson.D{{"$sum", 1}}}}}}
    resultInfo,_ =conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
    resultInfo.All(context.TODO(), &dbResultSum)
    //dbResultEverySum
    var dbResultEveryType []memberRecordEveryType
    matchStage = bson.D{{"$match", bson.D{{"create_time",bson.M{"$gt":startdate}},{"create_time",bson.M{"$lte":enddate}}}}}
    groupStage = bson.D{{"$group", bson.D{{"_id", bson.D{{"type","$type"},{"date", "$date"}}}, {"sum", bson.D{{"$sum", "$playmin"}}}, {"count", bson.D{{"$sum", 1}}}}}}
    sortStage = bson.D{{"$sort", bson.D{{"_id.date", 1}}}}
    resultInfo,_ =conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage,sortStage})
    resultInfo.All(context.TODO(), &dbResultEveryType)

    result =analysisResult{onlineCount,dbResultRecordName,dbResultEvery,dbResultSum,dbResultEveryType}
    return ResultCode,result


}
func GetAnalysisState()(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    onlineCount :=getOnlieCount()
    type memberState struct{
        State string `bson:"_id" json:"state"`
        Count int `bson:"count" json:"count"`
    }
    type analysisResult struct{
        OnlineCount int
        ResultState []memberState
    }
    var dbResultState []memberState

    groupStage := bson.D{{"$group", bson.D{{"_id", "$state"},{"count", bson.D{{"$sum", 1}}}}}}
    sortStage := bson.D{{"$sort", bson.D{{"_id", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{ groupStage,sortStage})
    resultInfo.All(context.TODO(), &dbResultState)

    result =analysisResult{onlineCount,dbResultState}
    return ResultCode,result

}
func GetAnalysisFeedback()(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "member_feedback")
    defer client.Disconnect(context.TODO())
    onlineCount :=getOnlieCount()
    type memberScore struct{
        Score int`bson:"_id" json:"score"`
        Count int `bson:"count" json:"count"`
    }
    type memberFeedback struct{
        Member_no int `bson:"member_no" json:"member_no"`
        Score int `bson:"score" json:"score"`
        Feedback string `bson:"feedback" json:"feedback"`
        Create_time time.Time `bson:"create_time" json:"create_time"`
    }
    type analysisResult struct{
        OnlineCount int
        ResultScore []memberScore
        ResultFeedback []memberFeedback
    }
    //avg score  
    var dbResultScore []memberScore
    groupStage := bson.D{{"$group", bson.D{{"_id", "$score"},{"count", bson.D{{"$sum", 1}}}}}}
    sortStage := bson.D{{"$sort", bson.D{{"_id", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{ groupStage,sortStage})
    resultInfo.All(context.TODO(), &dbResultScore)
    //memberFeedback
    var dbResultFeedback []memberFeedback
    sortStage = bson.D{{"$sort", bson.D{{"create_time", -1}}}}
    resultInfo,_ =conn.Aggregate(context.TODO(), mongo.Pipeline{sortStage})
    resultInfo.All(context.TODO(), &dbResultFeedback)


    result =analysisResult{onlineCount,dbResultScore,dbResultFeedback}
    return ResultCode,result

}
func GetAnalysisAge()(ResultCode int,result interface{}){
    conn,client :=services.DbConnect("project", "member")
    defer client.Disconnect(context.TODO())
    onlineCount :=getOnlieCount()
    type memberBirth struct{
        Birthday string `bson:"birthday" json:"birthday"`
    }
    type analysisResult struct{
        OnlineCount int
        ResultBirth []memberBirth
    }
    var dbResultBirth []memberBirth

    sortStage := bson.D{{"$sort", bson.D{{"birthday", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{sortStage})
    resultInfo.All(context.TODO(), &dbResultBirth)

    result =analysisResult{onlineCount,dbResultBirth}
    return ResultCode,result
}
func getOnlieCount()(result int){
    conn,client :=services.DbConnect("project", "member_token")
    defer client.Disconnect(context.TODO())
    type memberOnline struct{
        Member_no int
    }
    var dbResult []memberOnline
    matchStage := bson.D{{"$match", bson.D{{"expire_time",bson.M{"$gte":time.Now().AddDate(0, 0, 7).Add(-time.Minute*5).Unix()}}}}}
    groupStage := bson.D{{"$group", bson.D{{"_id", "$member_no"}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
    resultInfo.All(context.TODO(), &dbResult)

    return len(dbResult)

}