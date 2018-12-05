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

//InsertDemanda ...
func InsertDemanda(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var dm model.Demanda
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = dm.InsertDemanda(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//UpdateDemanda ...
func UpdateDemanda(w http.ResponseWriter, r *http.Request) {
	var dm model.Demanda
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateDemanda] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	dm.CodigoDemanda = int64(id)
	if err := dm.UpdateDemanda(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//DeleteDemanda ...
func DeleteDemanda(w http.ResponseWriter, r *http.Request) {
	var dm model.Demanda
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteDemanda -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	dm.CodigoDemanda = int64(id)
	if err := dm.DeleteDemanda(db); err != nil {
		log.Printf("[handler/DeleteDemanda -  Erro ao tentar deletar Demanda. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetDemanda ...
func GetDemanda(w http.ResponseWriter, r *http.Request) {
	var dm model.Demanda
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

	dm.CodigoDemanda = int64(id)
	err = dm.GetDemanda(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetDemanda -  Não há Demanda com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetDemanda -  Erro ao tentar buscar Demanda. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetDemandas ...
func GetDemandas(w http.ResponseWriter, r *http.Request) {
	var dm model.Demanda
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
	descricao := r.FormValue("descricao")

	dm.CodigoDemanda = int64(id)
	dm.DescricaoDemanda = descricao

	demandas, err := dm.GetDemandas(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetDemandas -  Não há Demanda com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Demanda cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetDemandas -  Erro ao tentar buscar Demanda. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, demandas, 0, 0)
}
