package controllers

import (
	"encoding/json"
	"github.com/afagundes/mongo-generic-dao/config"
	"github.com/afagundes/mongo-generic-dao/dao"
	"github.com/afagundes/mongo-generic-dao/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
)

var usuariosDAO dao.DAO

func init() {
	usuariosDAO = dao.DAO{Database: config.Database, Collection: config.Collection}
}

func GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	usuariosDAO.Connect()
	defer usuariosDAO.Disconnect()

	var usuarios []model.Usuario
	usuariosDAO.GetAll(&usuarios)

	writeResponse(w, usuarios, http.StatusOK)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	usuariosDAO.Connect()
	defer usuariosDAO.Disconnect()

	vars := mux.Vars(r)
	id := vars["id"]

	var usuario model.Usuario
	usuariosDAO.GetById(id, &usuario)

	if len(usuario.Nome) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeResponse(w, usuario, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	usuario, finalizado := getBody(w, r)
	if finalizado {
		return
	}

	usuariosDAO.Connect()
	defer usuariosDAO.Disconnect()

	id := usuariosDAO.Insert(usuario)
	usuario.ID = id.(primitive.ObjectID)

	writeResponse(w, usuario, http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	usuario, finalizado := getBody(w, r)
	if finalizado {
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	usuariosDAO.Connect()
	defer usuariosDAO.Disconnect()

	usuariosDAO.Update(id, usuario)
	usuario.ID, _ = primitive.ObjectIDFromHex(id)

	writeResponse(w, usuario, http.StatusOK)
}

func DeleteUser(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	usuariosDAO.Connect()
	defer usuariosDAO.Disconnect()

	usuariosDAO.DeleteById(id)
}

func getBody(w http.ResponseWriter, r *http.Request) (model.Usuario, bool) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return model.Usuario{}, true
	}

	var usuario model.Usuario
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return model.Usuario{}, true
	}

	return usuario, false
}

func writeResponse(w http.ResponseWriter, usuario interface{}, status int) {
	setJsonHeader(w)

	if status > 0 {
		w.WriteHeader(status)
	}

	err := json.NewEncoder(w).Encode(usuario)
	if err != nil {
		log.Fatal(err)
	}
}

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
}
