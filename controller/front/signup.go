package frontController

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"api/services"
	"api/services/front"
	"api/controller"
)

func MemberSignup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var memberInfo frontServices.SignupData
	body =services.InputCheck(string(body))
	json.Unmarshal(body, &memberInfo) //轉為json
	defer r.Body.Close()
	var response controller.ApiResponse
	if(len(memberInfo.Email)==0 || len(memberInfo.Password)==0 || len(memberInfo.Name)==0 || len(memberInfo.Birthday)==0 || len(memberInfo.State)==0 || len(memberInfo.City)==0 || len(memberInfo.Gender)==0 || len(memberInfo.Country)==0){
		response = controller.ApiResponse{ResultCode:2, ResultMessage:controller.ErrorMessage["error2"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=frontServices.SetMemberSignup(memberInfo)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusCreated, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}


}
