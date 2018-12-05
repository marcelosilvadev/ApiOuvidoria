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

//InsertPrioridade ...
func InsertPrioridade(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var pm model.Prioridade
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = pm.InsertPrioridade(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Prioridade", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, pm, 0, 0)
}

//UpdatePrioridade ...
func UpdatePrioridade(w http.ResponseWriter, r *http.Request) {
	var pm model.Prioridade
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdatePrioridade] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&pm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	pm.CodigoPrioridade = int64(id)
	if err := pm.UpdatePrioridade(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, pm, 0, 0)
}

//DeletePrioridade ...
func DeletePrioridade(w http.ResponseWriter, r *http.Request) {
	var pm model.Prioridade
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeletePrioridade -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	pm.CodigoPrioridade = int64(id)
	if err := pm.DeletePrioridade(db); err != nil {
		log.Printf("[handler/DeletePrioridade -  Erro ao tentar deletar Prioridade. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, pm, 0, 0)
}

//GetPrioridade ...
func GetPrioridade(w http.ResponseWriter, r *http.Request) {
	var pm model.Prioridade
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetPrioridade] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	pm.CodigoPrioridade = int64(id)
	err = pm.GetPrioridade(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetPrioridade -  Não há Prioridade com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Prioridade com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetPrioridade -  Erro ao tentar buscar Prioridade. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, pm, 0, 0)
}

//GetPrioridades ...
func GetPrioridades(w http.ResponseWriter, r *http.Request) {
	var pm model.Prioridade
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetPrioridades] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	pm.CodigoPrioridade = int64(id)
	pm.DescricaoPrioridade = descricao

	prioridades, err := pm.GetPrioridades(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetPrioridades -  Não há Prioridade com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Prioridade cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetPrioridades -  Erro ao tentar buscar Prioridade. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, prioridades, 0, 0)
}
