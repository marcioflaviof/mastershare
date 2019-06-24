package control

import (
	"encoding/json"
	"log"
	"net/http"
	"project/database"
	"project/models"
	"project/configs"
	"project/validation"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Instanciar o modelo a ser usado
	var user models.User

	// Pegar o JSON da requisição
	body := r.Body

	// Converter o JSON no modelo criado
	err := JsonStruct(&user, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(user)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Inserir o usuário no banco de dados
	err = database.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o cliente com o usuário inserido
	response, err := json.Marshal(user)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Pegar cookie da requisição
	c, err := r.Cookie("token")
	if err != nil {
		log.Printf(configs.COOKIE_ERROR+"%v\n",err)
		w.Write([]byte(configs.RESPONSE_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Verificar se usuario está logado
	s, err := ValidateLogin(c)
	if err != nil {
		// For any other type of error, return a bad request status
		log.Printf(configs.LOGIN_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_LOGIN))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Instanciar o modelo de atualização a ser usado
	var updateUser models.UpdatableUser

	// Pegar o JSON da requisição
	body:= r.Body

	// Converter o JSON no modelo criado
	err = JsonStruct(&updateUser, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(updateUser)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verificar se a alteração será no usuário logado
	if s != updateUser.Filter.Email {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("[ERROR] User dont match")
		return
	}

	// Alterar o usuário no banco de dados
	err = database.UpdateUser(updateUser.Filter, updateUser.Update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o Usuário Atualizado
	response, err := json.Marshal(updateUser.Update)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Pegar cookie da requisição
	c, err := r.Cookie(configs.COOKIE_NAME)
	if err != nil {
		log.Printf(configs.COOKIE_ERROR+"%v\n",err)
		w.Write([]byte(configs.RESPONSE_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Verificar se usuário está logado
	s, err := ValidateLogin(c)
	if err != nil {
		// For any other type of error, return a bad request status
		log.Printf(configs.LOGIN_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_LOGIN))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Instanciar o modelo de atualização a ser usado
	var deleteuser models.SecureUser

	//Pegar o JSON da requsição
	body := r.Body

	//Converter o JSON no modelo criado
	err = JsonStruct(&deleteuser, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	//Validar o modelo preenchido
	err = validation.Validator.Struct(deleteuser)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verificar se o usuário a ser deletado é o que está logado
	if s != deleteuser.Email {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Deletar o usuário no banco de dados
	resp, err := database.DeleteUser(deleteuser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o Usuário deletado
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func SearchUser(w http.ResponseWriter, r *http.Request) {

	// Identificar qual usuário buscar
	query := r.URL.Query()

	// Buscar o usuário no banco de dados
	resp, err := database.SearchUserEmail(query.Get("email"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o resultado da busca
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write(response)
}
