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

func MemberRecord(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	token :=r.Header.Get("Authorization")
	var data frontServices.MemberRecord
	body =services.InputCheck(string(body))
	json.Unmarshal(body, &data) //轉為json
	defer r.Body.Close()
	var response controller.ApiResponse

	tokenStatus,member_no := services.TokenCheck(token)
	if(tokenStatus=="error"){
		response = controller.ApiResponse{ResultCode:3, ResultMessage:controller.ErrorMessage["error3"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else if(len(data.Type)==0){
		response = controller.ApiResponse{ResultCode:2, ResultMessage:controller.ErrorMessage["error2"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=frontServices.SetMemberRecord(data,member_no)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusCreated, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}


}
