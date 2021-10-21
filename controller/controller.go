package controller
import (
	"net/http"
	"api/services"
)
var ErrorMessage = map[string]string{"error1":"Invalid api.", "error2":"Invalid parameters.", "error3":"Token is Not Valid."}
type ApiResponse struct {
	ResultCode    int
	ResultMessage interface{}
}
type Token struct{
	AccessToken    string
}

func OptionRespons(w http.ResponseWriter, r *http.Request){
	services.ResponseWithJson(w, http.StatusOK, 0)
}