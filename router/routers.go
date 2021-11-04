package routes

import (
	"net/http"
	frontController "api/controller/front"
	adminController "api/controller/admin"
	"api/controller"
	"api/services"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route
var routesManager []Route

func init() {
	/* front  */
	//POST
	register("POST", "/signup", frontController.MemberSignup, nil)
	register("POST", "/login", frontController.MemberLogin, nil)
	register("POST", "/feedBack", frontController.MemberFeedBack, nil)
	register("POST", "/record", frontController.MemberRecord, nil)
	//GET
	register("GET", "/memberInfo", frontController.MemberInfo, nil)
	register("GET", "/playAnalysis", frontController.MemberPlayAnalysis, nil)
	//DELETE
	register("DELETE", "/logout", frontController.MemberLogout, nil)

	/* admin  */
	//POST
	register("POST", "/admin/login", adminController.AdminLogin, nil)
	registerManager("POST", "/managerAdd", adminController.ManagerAdd, adminMiddleware)
	//GET
	registerManager("GET", "/manager", adminController.Manager, adminMiddleware)
	registerManager("GET", "/managerList", adminController.ManagerList, adminMiddleware)
	registerManager("GET", "/memberList", adminController.MemberList, adminMiddleware)
	registerManager("GET", "/memberDetail", adminController.MemberDetail, adminMiddleware)
	registerManager("GET", "/analysis_usetime", adminController.AnalysisUsetime, adminMiddleware)
	registerManager("GET", "/analysis_state", adminController.AnalysisState, adminMiddleware)
	registerManager("GET", "/analysis_feedback", adminController.AnalysisFeedback, adminMiddleware)
	registerManager("GET", "/analysis_age", adminController.AnalysisAge, adminMiddleware)
	//DELETE
	registerManager("DELETE", "/managerDel/{member_no}", adminController.ManagerDel, adminMiddleware)

}
func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Do stuff here

		token :=r.Header.Get("Authorization")
		tokenStatus,member_no := services.TokenCheck(token)
		if(r.Method=="OPTIONS"){
			next.ServeHTTP(w, r)
		}else if(tokenStatus=="error"){
			http.Error(w, "404 not found", http.StatusBadRequest)
		}else{
			ResultCode,_ :=services.GetManager(member_no)
			if(ResultCode==0){
				next.ServeHTTP(w, r)
			}else{
				http.Error(w, "404 not found", http.StatusBadRequest)
			}
		}
	})
}
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/").Subrouter()
	for _, route := range routes {
		apiRouter.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			apiRouter.Use(route.Middleware)
		}
	}

	// mux := mux.NewRouter()
	adminRouter := r.PathPrefix("/admin/").Subrouter()
	for _, route := range routesManager {
		adminRouter.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)
		if route.Middleware != nil {
			adminRouter.Use(route.Middleware)
		}
	}

	return r
}
func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
	routes = append(routes, Route{"OPTIONS", pattern, controller.OptionRespons, nil})
}
func registerManager(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routesManager = append(routesManager, Route{method, pattern, handler, middleware})
	routesManager = append(routesManager, Route{"OPTIONS", pattern, controller.OptionRespons, nil})
}

