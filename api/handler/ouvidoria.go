package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/ApiOuvidoria/api/db"
	"bitbucket.org/ApiOuvidoria/model"
	"bitbucket.org/ApiOuvidoria/util"
)

//InsertOuvidoria ...
func InsertOuvidoria(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var cm model.Ouvidoria
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = cm.InsertOuvidoria(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Ouvidoria", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cm, 0, 0)
}
