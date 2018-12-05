package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/ApiOuvidoria/api/db"
	"bitbucket.org/ApiOuvidoria/model"
	"bitbucket.org/ApiOuvidoria/util"
	"github.com/gorilla/mux"
)

//InsertSubAssunto ...
func InsertSubAssunto(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var am model.SubAssunto
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&am); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = am.InsertSubAssunto(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir SubAssunto", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, am, 0, 0)
}

//UpdateSubAssunto ...
func UpdateSubAssunto(w http.ResponseWriter, r *http.Request) {
	var am model.SubAssunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateSubAssunto] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&am); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	am.Codigo = int64(id)
	if err := am.UpdateSubAssunto(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, am, 0, 0)
}

//DeleteSubAssunto ...
func DeleteSubAssunto(w http.ResponseWriter, r *http.Request) {
	var am model.SubAssunto
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteSubAssunto -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	am.Codigo = int64(id)
	if err := am.DeleteSubAssunto(db); err != nil {
		log.Printf("[handler/DeleteSubAssunto -  Erro ao tentar deletar origem. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, am, 0, 0)
}

//GetSubAssunto ...
func GetSubAssunto(w http.ResponseWriter, r *http.Request) {
	var am model.SubAssunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSubAssunto] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	am.Codigo = int64(id)
	err = am.GetSubAssunto(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSubAssunto -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetSubAssunto -  Erro ao tentar buscar origem. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, am, 0, 0)
}

//GetSubAssuntos ...
func GetSubAssuntos(w http.ResponseWriter, r *http.Request) {
	var am model.SubAssunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSubAssuntos] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	am.Codigo = int64(id)
	am.DescricaoSubAssunto = descricao

	subAssuntos, err := am.GetSubAssuntos(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSubAssuntos -  Não há SubAssunto com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há SubAssunto cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetSubAssuntos -  Erro ao tentar buscar origens. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, subAssuntos, 0, 0)
}
