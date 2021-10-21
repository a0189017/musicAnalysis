package frontController

import (
	"net/http"
	"api/services"
	"api/services/front"
	"api/controller"
)

func MemberLogout(w http.ResponseWriter, r *http.Request) {
	token :=r.Header.Get("Authorization")
	defer r.Body.Close()
	var response controller.ApiResponse

	tokenStatus,_ := services.TokenCheck(token)
	if(tokenStatus=="error"){
		response = controller.ApiResponse{ResultCode:3, ResultMessage:controller.ErrorMessage["error3"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=frontServices.SetMemberLogout(token)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusCreated, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}


}
