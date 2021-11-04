package services

import (
	"encoding/json"
	"encoding/hex"
	"net/http"
    "net/url"
	"crypto/md5"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "regexp"
    "strings"
    "time"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,authorization")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Expose-Headers", "Cache-Control,Content-Language,Content-Type, Expires, Last-Modified")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func DbConnect(db string, collection string)(conn *mongo.Collection ,client *mongo.Client) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://172.31.0.2:27017"))
    if err != nil {
    	return nil,nil
	}
    conn = client.Database(db).Collection(collection)

    return conn,client

}
func PasswordMd5(str string) string  {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}
func InputCheck(src string) []byte {
    //Remove SCRIPT
    re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    src = re.ReplaceAllString(src, "")
    //Remove STYLE
    re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src = re.ReplaceAllString(src, "")
    src = strings.TrimSpace(src)
    return []byte(src)
}
func InputCheckUrl(data url.Values) map[string]string{
    paramMap := make(map[string]string, 0)
    for k, v := range data {
        if len(v) == 1 && len(v[0]) !=0 {
            paramMap[k] = v[0]
            //Remove SCRIPT
            re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
            paramMap[k] = re.ReplaceAllString(paramMap[k], "")
            //Remove STYLE
            re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
            paramMap[k] = re.ReplaceAllString(paramMap[k], "")
            paramMap[k] = strings.TrimSpace(paramMap[k])
        } else {
            break
        }
    }
    return paramMap
}
func TokenCheck(accessToken string)(restult string, member_no int){
	conn,client :=DbConnect("project", "member_token")
    defer client.Disconnect(context.TODO())

    type token struct{
    	Token string
    	Member_no int
    }
    var tokenExist token
    var result string

    filter := bson.M{"token":accessToken,"expire_time": bson.M{"$gte": time.Now().Unix()}}
    conn.FindOne(context.TODO(), filter).Decode(&tokenExist)
    if tokenExist.Token!=""{
    	conn.UpdateOne(context.TODO(), bson.M{"token":accessToken},bson.M{"$set":bson.M{"expire_time":time.Now().AddDate(0, 0, 7).Unix()}})
    	result ="success"
    }else{
    	result ="error"
    }
    return result,tokenExist.Member_no

}
func GetManager(member_no int)(ResultCode int,result interface{}){
    type ManagerInfo struct {
        Member_no int
    }   
    conn,client :=DbConnect("project", "manager")
    defer client.Disconnect(context.TODO())

    var resultManager ManagerInfo

    filter := bson.M{"member_no":member_no}
    conn.FindOne(context.TODO(), filter).Decode(&resultManager)
 
    if (ManagerInfo{}!=resultManager){
        ResultCode=0
        result="success"
    }else{
        result="沒有權限！"
        ResultCode=5
    }
    return ResultCode,result
}