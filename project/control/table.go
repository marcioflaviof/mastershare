package control

import (
	"encoding/json"
	"log"
	"project/database"
	"project/models"
	"project/validation"
	"project/configs"
	"net/http"
	"github.com/gorilla/mux"
)

//Create Table
func RegisterTable(w http.ResponseWriter, r *http.Request) {

	//instanciando o modelo para uso
	var table models.Table

	//Convertendo o JSON no nosso modelo
	body := r.Body
	err := JsonStruct(&table, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	//Validando estrutura
	err = validation.Validator.Struct(table)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//(opcional) Já dividir as contas do modelo para a inserção
	//ShareTable(&table)

	//Inserir a mesa no banco de dados 
	err = database.InsertTable(table)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder a requisição com o usuário inserido
	response, err := json.Marshal(table)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write([]byte(response))

}

func UpdateTable(w http.ResponseWriter, r *http.Request) {

	// Identificar a mesa a ser atualizada
	params := mux.Vars(r)

	// Criar o modelo da mesa de atualização
	var updateTable models.Table

	// Pegar o JSON da requisição
	body := r.Body

	//Converter o JSON no modelo criado
	err := JsonStruct(&updateTable, body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Realizar as validadções no modelo
	err = validation.Validator.Struct(updateTable)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Atualizar no banco de dados a mesa
	err = database.UpdateTable(params["master"], updateTable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com a mesa atualizada
	response, err := json.Marshal(updateTable)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	}  
	w.Write([]byte(response))

}

func DeleteTable(w http.ResponseWriter, r *http.Request) {

	// Identificar a mesa que será deletada
	params := mux.Vars(r) 

	// Deletar mesa no banco de dados
	resp, err := database.DeleteTable(params["master"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responder o usuário com a mesa deletada
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	}   
	w.Write([]byte(response))

}

func SearchTable(w http.ResponseWriter, r *http.Request) {

	//Identificar a mesa que deseja buscar
	params := mux.Vars(r) 

	// Realizar a busca no banco de dados
	resp, err := database.SearchTable(params["master"])
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

func SearchTables(w http.ResponseWriter, r *http.Request) { 

	// Buscar TODAS as mesas no servidor
	resp, err := database.SearchTables()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Responde o usuário com o resultado da busca
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	} 
	w.Write(response)
}

func TableShare(w http.ResponseWriter, r *http.Request) {
	
	//Identificar a mesa que deseja dividir a conta
	params := mux.Vars(r) 

	// Realizar a busca no banco de dados
	resp, err := database.SearchOnlyTable(params["master"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Dividir a conta para cada mesa da buscar
	ShareTable(&resp)

	// Atualizar o banco de dados com as mesas alteradas
	err = database.UpdateTable(params["master"], resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}


	// Responder o usuário com as mesas alteradas
	response, err := json.Marshal(resp)
	if err != nil{
		log.Printf(configs.MARSHAL_ERROR+"%v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_MARSHAL))
		return
	}  
	w.Write(response)
}

func ShareTable(table *models.Table) {

	//pegar a soma dos valores dos produtos da mesa
	p := table.TotalValue()

	//pegar o total de usuários na mesa
	u := float64(len(table.Users))

	c := p / u
	
	//dividir a conta igualmente para todos os usuários
	table.ShareAllBills(c)
}
