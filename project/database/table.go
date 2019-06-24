package database

import (
	"context"
	"log"
	"project/models"
	"project/configs"

	"go.mongodb.org/mongo-driver/bson"
)

//////////////////////////////
// Single Resultd functions //
//////////////////////////////

func InsertTable(u models.Table) (err error) {

	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.TABLE_COLLECTION)

	// Inserir no banco de dados
	resp, err := c.InsertOne(context.TODO(), u)
	if err != nil{
		log.Printf("[ERROR] Error in insert table: %v\n",err)
	}

	log.Printf("[INFO] Table of %v was inserted\n", resp)

	return
}

func DeleteTable(master string) (u models.Table, err error) {

	// Selecionar a collection que vamos usar
	c := DB.Collection(configs.TABLE_COLLECTION)

	// Criar filtro para busca
	filter := bson.D{{"master", master}}

	// Buscar e deletar mesa da busca
	err = c.FindOneAndDelete(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Printf("[ERROR] Error in remove table: %v\n", err)
	}

	log.Printf("[INFO] Table of %v was deleted\n", u.Master)

	return

}

func SearchOnlyTable(name string) (table models.Table, err error) {

	// Selecionar a collection que iremos usar
	collection := DB.Collection(configs.TABLE_COLLECTION)

	// Criar filtro para a busca
	filter := bson.D{{"master", name}}

	// Buscar apenas uma mesa com o mestre identificado
	err = collection.FindOne(context.TODO(), filter).Decode(&table)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
	}

	log.Printf("[INFO] Table of %v was searched\n", table.Master)

	return
}


func UpdateTable(s string, newTable models.Table) (err error) {

	// Selecionar a collection que vamos usar
	c := DB.Collection(configs.TABLE_COLLECTION)

	// Definir BSON de busca
	filter := bson.D{{"master", s}}

	//Definir BSON de atualização
	filter2 := bson.D{{"users", newTable.Users}, {"products", newTable.Products}}

	// Atualizar uma uníca mesa
	resp, err := c.UpdateOne(context.TODO(), filter, bson.D{{"$set", filter2}})
	if err != nil {
		log.Printf("[ERROR] probleming searching Table: %v\n", err)
		return
	}

	log.Printf("[INFO] Table of %v was updated\n", resp)

	return

}
////////////////////////////////
// Multiple Results functions //
////////////////////////////////

func SearchTable(name string) (tables []models.Table, err error) {

	// Selecionar a colletion que iremos usar
	collection := DB.Collection(configs.TABLE_COLLECTION)

	// Criar filtro para a busca
	filter := bson.D{{"master", name}}

	// Buscar o conjunto de mesas de um mesmo mestre
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	// Devolver as mesas encontradas
	cur.All(context.TODO(), &tables)

	log.Printf("[INFO] The many tables of %v were searched\n", name)

	return
}

func SearchTables() (tables []models.Table, err error) {

	// Selecionar a collection que iremos usar
	collection := DB.Collection(configs.TABLE_COLLECTION)

	// Realizar busca completa
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	// Devolver as mesas encontradas
	cur.All(context.TODO(), &tables)

	log.Println("[INFO] All tables were searched")
	
	return
}
