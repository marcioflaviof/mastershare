package service

import (
	"project/configs"
	"project/control"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createRouter() (r *mux.Router) {

	r = mux.NewRouter()
	//Users routes
	r.HandleFunc(configs.USER_PATH, control.RegisterUser).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.USER_PATH, control.UpdateUser).Methods(http.MethodPut).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.USER_PATH, control.DeleteUser).Methods(http.MethodDelete).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.USER_PATH, control.SearchUser).Methods(http.MethodGet)
	r.HandleFunc(configs.USER_PATH+"login/", control.Login).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	//Tables routes
	r.HandleFunc(configs.TABLE_PATH, control.RegisterTable).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.TABLE_PATH+configs.TABLE_ID, control.UpdateTable).Methods(http.MethodPut).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.TABLE_PATH+configs.TABLE_ID, control.DeleteTable).Methods(http.MethodDelete).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.TABLE_PATH+configs.TABLE_ID, control.SearchTable).Methods(http.MethodGet)
	r.HandleFunc(configs.TABLE_PATH, control.SearchTables).Methods(http.MethodGet)
	r.HandleFunc(configs.TABLE_PATH+configs.TABLE_ID+"share/", control.TableShare).Methods(http.MethodGet)


	//Product routes
	r.HandleFunc(configs.PRODUCT_PATH, control.RegisterProduct).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.PRODUCT_PATH, control.UpdateProduct).Methods(http.MethodPut).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.PRODUCT_PATH, control.DeleteProduct).Methods(http.MethodDelete).Headers("Content-Type", "application/json")
	r.HandleFunc(configs.PRODUCT_PATH, control.SearchProduct).Methods(http.MethodGet)
	r.HandleFunc(configs.PRODUCT_PATH+"all/", control.SearchProducts).Methods(http.MethodGet)

	r.HandleFunc("/refresh", control.Refresh)

	return
}

func createCORS() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})
}