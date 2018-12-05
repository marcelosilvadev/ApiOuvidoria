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

//InsertTipoManifestacao ...
func InsertTipoManifestacao(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var tm model.TipoManifestacao
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = tm.InsertTipoManifestacao(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Tipo Manifestação", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, tm, 0, 0)
}

//UpdateTipoManifestacaoo ...
func UpdateTipoManifestacaoo(w http.ResponseWriter, r *http.Request) {
	var tm model.TipoManifestacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateTipoManifestacaoo] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&tm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	tm.CodigoTipoManifestacao = int64(id)
	if err := tm.UpdateTipoManifestacao(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, tm, 0, 0)
}

//DeleteTipoManifestacao ...
func DeleteTipoManifestacao(w http.ResponseWriter, r *http.Request) {
	var tm model.TipoManifestacao
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteTipoManifestacao -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	tm.CodigoTipoManifestacao = int64(id)
	if err := tm.DeleteTipoManifestacao(db); err != nil {
		log.Printf("[handler/DeleteTipoManifestacao -  Erro ao tentar deletar Tipo Manifestacao. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, tm, 0, 0)
}

//GetTipoManifestacao ...
func GetTipoManifestacao(w http.ResponseWriter, r *http.Request) {
	var tm model.TipoManifestacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetTipoManifestacao] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	tm.CodigoTipoManifestacao = int64(id)
	err = tm.GetTipoManifestacao(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetTipoManifestacao -  Não há Tipo Manifestacao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Tipo Manifestacao com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetTipoManifestacao -  Erro ao tentar buscar Tipo Manifestacao. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, tm, 0, 0)
}

//GetTipoManifestacoes ...
func GetTipoManifestacoes(w http.ResponseWriter, r *http.Request) {
	var tm model.TipoManifestacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetTipoManifestacoes] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	tm.CodigoTipoManifestacao = int64(id)
	tm.DescricaoTipoManifestacao = descricao

	manifestacoes, err := tm.GetTipoManifestacoes(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetTipoManifestacoes -  Não há Tipo Manifestacao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Tipo Manifestacao cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetTipoManifestacoes -  Erro ao tentar buscar Tipo Manifestacao. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, manifestacoes, 0, 0)
}
