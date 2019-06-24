package control

import (
	"encoding/json"
	"log"
	"project/database"
	"project/models"
	"project/configs"
	"project/validation"
	"net/http"
)

func RegisterProduct(w http.ResponseWriter, r *http.Request) {

	// Instanciar o modelo a ser usado
	var product models.Product

	// Pegar o JSON da requisição
	body:= r.Body

	// Converter o JSON no modelo criado
	err := JsonStruct(&product, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(product)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Inserir o modelo no banco de dados
	err = database.InsertProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o produto inserido
	response, err := json.Marshal(product)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	// Instanciar o modelo a ser usado
	var updateProduct models.UpdatableProduct

	// Pegar JSON da requisição
	body:= r.Body

	// Converter o JSON no modelo criado
	err := JsonStruct(&updateProduct, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(updateProduct)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Atualizar o modelo no banco de dados
	err = database.UpdateProduct(updateProduct.Filter, updateProduct.Update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o produto atualizado
	response, err := json.Marshal(updateProduct.Update)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	}  
	w.Write([]byte(response))

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	// Instanciar o modelo a ser usado
	var deleteproduct models.SecureProducts

	// Pegar o JSON da requisição
	body := r.Body

	// Converter o JSON no modelo criado
	err := JsonStruct(&deleteproduct, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(deleteproduct)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Deletar o modelo no banco de dados
	resp, err := database.DeleteProduct(deleteproduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com o produto deletar
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func SearchProduct(w http.ResponseWriter, r *http.Request) {

	// Identificar qual produto buscar
	query := r.URL.Query()

	// Buscar o produto no banco de dados
	resp, err := database.SearchProduct(query.Get("name"))
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

func SearchProducts(w http.ResponseWriter, r *http.Request) {

	// Buscar TODOS os produtos no banco de dados
	resp, err := database.SearchProducts()
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


