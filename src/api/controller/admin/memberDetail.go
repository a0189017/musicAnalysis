package adminController

import (
	"net/http"
	"api/services"
	"api/services/admin"
	"api/controller"
)

func MemberDetail(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Query()
	data :=services.InputCheckUrl(param)
	defer r.Body.Close()
	var response controller.ApiResponse
	
	if(data["member_no"]==""){
		response = controller.ApiResponse{ResultCode:2, ResultMessage:controller.ErrorMessage["error2"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=adminServices.GetMemberDetail(data)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusOK, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}


}
