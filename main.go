package main

import (
	"fmt"
	"github.com/mongo-generic-dao/dao"
	"github.com/mongo-generic-dao/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DATABASE = "myFirstDatabase"
	COLLECTION = "myMongoCollection"
)

func main() {
	fmt.Println("Teste de interação com MongoBD via DAO genérico")
	fmt.Println("\nInicializando DAO e conectando ao MongoDB")

	usuarioDAO := dao.DAO{ Database: DATABASE, Collection: COLLECTION }
	usuarioDAO.Connect()
	defer usuarioDAO.Disconnect()

	fmt.Println("\nRecuperando todos os documentos...")
	getAllUsers(usuarioDAO)

	fmt.Println("\nPegando um documento por ID...")
	id, usuario := getUserById(usuarioDAO)

	fmt.Println("\nAtualizando um documento...")
	updateUser(usuario, usuarioDAO, id)

	fmt.Println("\nCriando novo documento...")
	id = createUser(usuario, usuarioDAO, id)

	fmt.Println("\nRemovendo um documento...")
	deleteUser(usuarioDAO, id)
}

func createUser(usuario model.Usuario, usuarioDAO dao.DAO, id string) string {
	usuario = model.Usuario{Nome: "Otávio", Idade: 42, Bio: "Inserindo novo usuário", Foto: "https://pbs.twimg.com/profile_images/767051707389992963/w0x4-LIs_400x400.jpg"}
	objectID := usuarioDAO.Insert(usuario)

	id = objectID.(primitive.ObjectID).Hex()
	fmt.Printf("Inseriu novo usuário %s\n", id)

	return id
}

func getAllUsers(usuarioDAO dao.DAO) {
	var usuarios []model.Usuario
	usuarioDAO.GetAll(&usuarios)

	for _, usuario := range usuarios {
		fmt.Printf("%+v\n", usuario)
	}
}

func getUserById(usuarioDAO dao.DAO) (string, model.Usuario) {
	var id = "610060037ce3bd450438e71c"
	var usuario model.Usuario
	usuarioDAO.GetById(id, &usuario)

	fmt.Printf("%+v\n", usuario)

	return id, usuario
}

func updateUser(usuario model.Usuario, usuarioDAO dao.DAO, id string) {
	usuario.Bio = "Testando MongoDB com Go e NodeJS"
	result := usuarioDAO.Update(id, usuario)

	fmt.Printf("%+v\n", result)
}

func deleteUser(usuarioDAO dao.DAO, id string) {
	qtdDeleted := usuarioDAO.DeleteById(id)
	fmt.Printf("\nDeletados %d registros", qtdDeleted)
}
