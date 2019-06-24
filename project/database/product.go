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

func InsertProduct(p models.Product) (err error) {

	// Selecionar a colletion que vamos usar
	c := DB.Collection(configs.PRODUCT_COLLECTION)

	// Inserir no banco de dados
	resp, err := c.InsertOne(context.TODO(), p)
	if err != nil{
		log.Printf("[ERROR] Error in insert product: %v\n",err)
	}

	log.Printf("[INFO] Inserted Product: %v\n", resp)

	return
}

func DeleteProduct(d models.SecureProducts) (p models.Product, err error) {
	
	// Selecionar a colletion que vamos usar
	c := DB.Collection(configs.PRODUCT_COLLECTION)
	
	// Instanciar filtro de busca
	filter := bson.D{{"name", d.Name}}

	// Buscar e deletar mesa da busca
	err = c.FindOneAndDelete(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Printf("[ERROR] Error in remove Product: %v\n", err)
	}

	log.Printf("[INFO] Deleted Product: %v\n", p)

	return

}

func UpdateProduct(s models.SecureProducts, newProduct models.Product) (err error) {
	
	// Selecionar a colletion que vamos usar
	c := DB.Collection(configs.PRODUCT_COLLECTION)
	
	// Instanciar filtro de busca
	filter := bson.D{{"name", s.Name}}

	// Instanciar filtro de atualização
	filter2 := bson.D{{"name", newProduct.Name}, {"price", newProduct.Price}}
	
	// Atualizar um unico produto
	resp, err := c.UpdateOne(context.TODO(), filter, bson.D{{"$set", filter2}})
	if err != nil {
		log.Printf("[ERROR] probleming searching Product: %v\n", err)
		return
	}

	log.Printf("[INFO] Updated Product: %v\n", resp)

	return

}

////////////////////////////////
// Multiple Results functions //
////////////////////////////////

func SearchProduct(name string) (products []models.Product, err error) {

	// Selecionar a colletion que vamos usar
	collection := DB.Collection(configs.PRODUCT_COLLECTION)

	// Instanciar filtro de busca
	filter := bson.D{{"name", name}}

	// Buscar o conjunto de produtos com um mesmo nome
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	// Devolver os produtos encontrados
	cur.All(context.TODO(), &products)

	log.Println("[INFO] Searched Products: sucess")

	return
}

func SearchProducts() (products []models.Product, err error) {

	// Selecionar a colletion que vamos usar
	collection := DB.Collection(configs.PRODUCT_COLLECTION)

	// Instanciar filtro de busca
	filter := bson.D{{}}

	// Buscar o conjunto de produtos com um mesmo nome
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	// Devolver os produtos encontrados
	cur.All(context.TODO(), &products)

	log.Println("[INFO] Searched Products: sucess")

	return
}