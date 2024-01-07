package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Monteiro712/go-rest-api/db"
	"github.com/Monteiro712/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request){
	
	var p []models.Personalidade
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request){
	//pegar o id na url
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	db.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request){
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	db.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)	
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	db.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars ["id"]
	var personalidade models.Personalidade
	db.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	db.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}