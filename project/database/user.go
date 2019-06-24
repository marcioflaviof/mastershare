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

func InsertUser(u models.User) (err error) {

	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Inserir no banco de dados
	resp, err := c.InsertOne(context.TODO(), u)
	if err != nil {
		log.Printf("[ERROR]  insert user: %v\n", err)
		return
	}

	log.Printf("[INFO] Inserted user: %v\n", resp)

	return
}

func DeleteUser(d models.SecureUser) (u models.User, err error) {
	
	// Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Criar filtro de busca 
	filter := bson.D{{"email", d.Email}, {"pass", d.Pass}}

	// Buscar e deletar usuario da busca
	err = c.FindOneAndDelete(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Printf("[ERROR] Error in remove User: %v\n", err)
		return
	}

	log.Printf("[INFO] Deleted User: %v\n", u)

	return

}

func SearchUser(name string) (user models.User, err error) {

	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Criar filtro de busca com nome
	filter := bson.D{{"name", name}}

	// Buscar um unico usuario
	err = c.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	log.Printf("[INFO] Searched User: %v\n", user)

	return
}

func SearchUserEmail(email string) (user models.User, err error) {

	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Criar filtro de busca com email
	filter := bson.D{{"email", email}}

	//Buscar um unico usuario no banco de dados
	err = c.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	log.Printf("[INFO] Searched User: %v\n", user)
	return
}

func UpdateUser(s models.SecureUser, newUser models.User) (err error) {
	
	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)
	
	// Criar filtro de busca 
	filter := bson.D{{"email", s.Email}, {"pass", s.Pass}}

	// Criar conjunto de atualização 
	updater := bson.D{{"name", newUser.Name}, {"pass", newUser.Pass}, {"bill", newUser.Bill}}
	
	// Atualizar um unico usuario
	resp, err := c.UpdateOne(context.TODO(), filter, bson.D{{"$set", updater}})
	if err != nil {
		log.Printf("[ERROR] probleming searching user: %v\n", err)
		return
	}

	log.Printf("[INFO] Updated User: %v\n", resp)

	return

}

////////////////////////////////
// Multiple Results functions //
////////////////////////////////

func SearchUsers(name string) (users []models.User, err error) {

	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Criar filtro de busca com nome
	filter := bson.D{{"name", name}}

	// Buscar o conjunto de mesas de um mesmo mestre
	cur, err := c.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("[ERROR] Search error: %v\n", err)
		return
	}

	// Devolver as mesas encontradas
	cur.All(context.TODO(), &users)

	log.Println("[INFO] Searched Users: sucess")

	return
}

func LoginUser(l models.SecureUser) (u models.User, err error) {
	
	//Selecionar a collection que vamos usar
	c := DB.Collection(configs.USER_COLLECTION)

	// Criar filtro de busca 
	filter := bson.D{{"email", l.Email}, {"pass", l.Pass}}

	// Buscar um unico usuario no banco de dados
	err = c.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Printf("[ERROR] probleming searching user: %v\n", err)
		return
	}

	log.Printf("[INFO] Searched User: %v\n", u)

	return
}
