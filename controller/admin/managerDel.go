package adminController

import (
	"net/http"
	"api/services"
	"api/services/admin"
	"api/controller"
	"github.com/gorilla/mux"
)

func ManagerDel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 获取参数
	defer r.Body.Close()
	var response controller.ApiResponse

	ResultCode,result :=adminServices.DelManager(vars["member_no"])
	response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
	switch(ResultCode){
		case 0:
			services.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}

}
