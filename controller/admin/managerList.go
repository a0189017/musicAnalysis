package adminController

import (
	"net/http"
	"api/services"
	"api/services/admin"
	"api/controller"
)

func ManagerList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var response controller.ApiResponse

	ResultCode,result :=adminServices.GetManagerList()
	response = controller.ApiResponse{ResultCode:ResultCode, ResultMessage:result}
	switch(ResultCode){
		case 0:
			services.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			services.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}

}
