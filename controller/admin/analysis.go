package adminController

import (
	"net/http"
	"api/services"
	"api/services/admin"
	"api/controller"

)

func AnalysisUsetime(w http.ResponseWriter, r *http.Request) {
	token :=r.Header.Get("Authorization")
	param := r.URL.Query()
	data :=services.InputCheckUrl(param)
	defer r.Body.Close()
	var response controller.ApiResponse

	tokenStatus,_ := services.TokenCheck(token)
	if(tokenStatus=="error"){
		response = controller.ApiResponse{ResultCode:3, ResultMessage:controller.ErrorMessage["error3"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else if(len(data["startdate"])==0 || len(data["enddate"])==0){
		response = controller.ApiResponse{2, data}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=adminServices.GetAnalysisUsetime(data)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusOK, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}
}
func AnalysisState(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var response controller.ApiResponse

	ResultCode,result :=adminServices.GetAnalysisState()
	response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
	switch(ResultCode){
		case 0:
			services.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}
func AnalysisFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var response controller.ApiResponse

	ResultCode,result :=adminServices.GetAnalysisFeedback()
	response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
	switch(ResultCode){
		case 0:
			services.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}
func AnalysisAge(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var response controller.ApiResponse

	ResultCode,result :=adminServices.GetAnalysisAge()
	response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
	switch(ResultCode){
		case 0:
			services.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}


