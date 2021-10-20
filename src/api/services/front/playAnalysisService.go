package frontServices

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
    "api/services"
)
type MemberAnalysis struct {
    AccessToken string
    StartDate string
    EndDate string
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
type analysisResult struct{
    RecordSum []*memberRecordSum `json:"recordSum"`
    RecordEvery []*memberRecordEvery `json:"recordEvery"`
}
func GetMemberAnalysis(data map[string]string,member_no int)(ResultCode int,result interface{}){
	conn,client :=services.DbConnect("project", "member_record")
    defer client.Disconnect(context.TODO())
    //find
    var dbResultSum []*memberRecordSum
    var dbResultEvery []*memberRecordEvery

    startdate, _ := time.Parse(time.RFC3339, data["startdate"]+"T00:00:00.000Z")
    enddate, _ := time.Parse(time.RFC3339, data["enddate"]+"T00:00:00.000Z")

    //everyday
    matchStage := bson.D{{"$match", bson.D{{"member_no", member_no},{"create_time",bson.M{"$gt":startdate}},{"create_time",bson.M{"$lte":enddate}}}}}
    groupStage := bson.D{{"$group", bson.D{{"_id", "$date"}, {"sum", bson.D{{"$sum", "$playmin"}}}, {"count", bson.D{{"$sum", 1}}}}}}
    sortStage := bson.D{{"$sort", bson.D{{"_id", 1}}}}
    resultInfo,_ :=conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage,sortStage})
    resultInfo.All(context.TODO(), &dbResultEvery)
    //sum
    matchStage = bson.D{{"$match", bson.D{{"member_no", member_no},{"create_time",bson.M{"$gt":startdate}},{"create_time",bson.M{"$lte":enddate}}}}}
    groupStage = bson.D{{"$group", bson.D{{"_id", "$type"}, {"sum", bson.D{{"$sum", "$playmin"}}}, {"count", bson.D{{"$sum", 1}}}}}}
    resultInfo,_ =conn.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
    resultInfo.All(context.TODO(), &dbResultSum)

    result =analysisResult{dbResultSum,dbResultEvery}

    return ResultCode,result


}