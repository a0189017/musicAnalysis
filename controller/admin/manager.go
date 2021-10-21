package adminController

import (
	"net/http"
	"api/services"
)

func Manager(w http.ResponseWriter, r *http.Request) {

	services.ResponseWithJson(w, http.StatusOK, 0)


}
