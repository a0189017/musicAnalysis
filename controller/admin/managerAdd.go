package adminController

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"api/services"
	"api/services/admin"
	"api/controller"
)
func ManagerAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var memberEmail adminServices.ManagerEmail
	body =services.InputCheck(string(body))
	json.Unmarshal(body, &memberEmail) //轉為json
	defer r.Body.Close()
	var response controller.ApiResponse
	
	if(len(memberEmail.Email)==0){
		response = controller.ApiResponse{ResultCode:2, ResultMessage:controller.ErrorMessage["error2"]}
		services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}else{
		ResultCode,result :=adminServices.SetManager(memberEmail)
		response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
		switch(ResultCode){
			case 0:
				services.ResponseWithJson(w, http.StatusCreated, response) //回傳
			case 5:
				services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}


}