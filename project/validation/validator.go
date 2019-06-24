package validation

import (
	val "gopkg.in/go-playground/validator.v9"
	//"net/http"
)

var Validator *val.Validate

//var ResponseError map[string]int

func CreateValidator() {
	Validator = val.New()
	//createResponse()
}

//func createResponse (){
//	ResponseError = make(map[string]int)
//	ResponseError["Failed Token"] = http.StatusInternalServerError
//	ResponseError["Claim without email"] = http.StatusUnauthorized
//	ResponseError["Token not valid"] = http.StatusUnauthorized

//}

//func ResponseStatus(s error) (i int){
//	return ResponseError[s+""]
//}
