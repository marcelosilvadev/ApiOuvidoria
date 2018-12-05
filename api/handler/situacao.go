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

//InsertSituacao ...
func InsertSituacao(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var si model.Situacao
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&si); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = si.InsertSituacao(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//UpdateSituacao ...
func UpdateSituacao(w http.ResponseWriter, r *http.Request) {
	var si model.Situacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateSituacao] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&si); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	si.CodigoSituacao = int64(id)
	if err := si.UpdateSituacao(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//DeleteSituacao ...
func DeleteSituacao(w http.ResponseWriter, r *http.Request) {
	var si model.Situacao
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteSituacao -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	si.CodigoSituacao = int64(id)
	if err := si.DeleteSituacao(db); err != nil {
		log.Printf("[handler/DeleteSituacao -  Erro ao tentar deletar Situacao. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//GetSituacao ...
func GetSituacao(w http.ResponseWriter, r *http.Request) {
	var si model.Situacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSituacao] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	si.CodigoSituacao = int64(id)
	err = si.GetSituacao(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSituacao -  Não há Situacao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetSituacao -  Erro ao tentar buscar Situacao. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//GetSituacoes ...
func GetSituacoes(w http.ResponseWriter, r *http.Request) {
	var si model.Situacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSituacaos] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	si.CodigoSituacao = int64(id)
	si.DescricaoSituacao = descricao

	situacoes, err := si.GetSituacoes(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSituacoes -  Não há Situacao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Situacao cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetSituacoes -  Erro ao tentar buscar Situacoes. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, situacoes, 0, 0)
}
