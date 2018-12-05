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

//InsertCidadao ...
func InsertCidadao(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var dm model.Cidadao
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = dm.InsertCidadao(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//UpdateCidadao ...
func UpdateCidadao(w http.ResponseWriter, r *http.Request) {
	var dm model.Cidadao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateCidadao] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&dm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	dm.CodigoCidadao = int64(id)
	if err := dm.UpdateCidadao(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//DeleteCidadao ...
func DeleteCidadao(w http.ResponseWriter, r *http.Request) {
	var dm model.Cidadao
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteCidadao -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	dm.CodigoCidadao = int64(id)
	if err := dm.DeleteCidadao(db); err != nil {
		log.Printf("[handler/DeleteCidadao -  Erro ao tentar deletar Cidadão. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetCidadao ...
func GetCidadao(w http.ResponseWriter, r *http.Request) {
	var dm model.Cidadao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemanda] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	dm.CodigoCidadao = int64(id)
	err = dm.GetCidadao(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetCidadao -  Não há Cidadao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cidadão com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetCidadao -  Erro ao tentar buscar Cidadao. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetCidadaos ...
func GetCidadaos(w http.ResponseWriter, r *http.Request) {
	var dm model.Cidadao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemandas] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")

	dm.CodigoCidadao = int64(id)
	dm.Nome = nome

	cidadaos, err := dm.GetCidadaos(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetCidadaos -  Não há Cidadão com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Cidadão cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetCidadaos -  Erro ao tentar buscar Cidadão. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cidadaos, 0, 0)
}
