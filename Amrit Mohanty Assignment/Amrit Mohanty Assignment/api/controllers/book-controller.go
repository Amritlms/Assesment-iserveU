package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akhil/go-bookstore/api/models"
	"github.com/akhil/go-bookstore/api/utils"
)

var NewBook models.User

func CreateUser(w http.ResponseWriter, r *http.Request){
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u, err:= CreateUser.CreateUser()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := json.NewEncoder(w).Encode(struct { 
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return	
	}
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}