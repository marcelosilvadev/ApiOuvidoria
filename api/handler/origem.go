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

//InsertOrigem ...
func InsertOrigem(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var om model.Origem
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&om); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = om.InsertOrigem(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Origem", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, om, 0, 0)
}

//UpdateOrigem ...
func UpdateOrigem(w http.ResponseWriter, r *http.Request) {
	var om model.Origem
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateOriegm] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&om); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	om.CodigoOrigem = int64(id)
	if err := om.UpdateOrigem(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, om, 0, 0)
}

//DeleteOrigem ...
func DeleteOrigem(w http.ResponseWriter, r *http.Request) {
	var om model.Origem
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteOrigem -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	om.CodigoOrigem = int64(id)
	if err := om.DeleteOrigem(db); err != nil {
		log.Printf("[handler/DeleteOrigem -  Erro ao tentar deletar origem. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, om, 0, 0)
}

//GetOrigem ...
func GetOrigem(w http.ResponseWriter, r *http.Request) {
	var om model.Origem
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetOrigem] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	om.CodigoOrigem = int64(id)
	err = om.GetOrigem(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetOrigem -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetOrigem -  Erro ao tentar buscar origem. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, om, 0, 0)
}

//GetOrigens ...
func GetOrigens(w http.ResponseWriter, r *http.Request) {
	var om model.Origem
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetOrigens] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	om.CodigoOrigem = int64(id)
	om.DescricaoOrigem = descricao

	origens, err := om.GetOrigens(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetOrigens -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há origem cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetClassificacoes -  Erro ao tentar buscar origens. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, origens, 0, 0)
}
