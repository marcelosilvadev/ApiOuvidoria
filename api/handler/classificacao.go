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

//InsertClassificacao ...
func InsertClassificacao(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var cm model.Classificacao
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
	err = cm.InsertClassificacao(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Classificao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cm, 0, 0)
}

//UpdateClassificacao ...
func UpdateClassificacao(w http.ResponseWriter, r *http.Request) {
	var cm model.Classificacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateClassificacao] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&cm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	cm.CodigoClassificacao = int64(id)
	if err := cm.UpdateClassificacao(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cm, 0, 0)
}

//DeleteClassificacao ...
func DeleteClassificacao(w http.ResponseWriter, r *http.Request) {
	var cm model.Classificacao
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteClassificacao -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	cm.CodigoClassificacao = int64(id)
	if err := cm.DeleteClassificacao(db); err != nil {
		log.Printf("[handler/DeleteClassificacao -  Erro ao tentar deletar segmento. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cm, 0, 0)
}

//GetClassificacao ...
func GetClassificacao(w http.ResponseWriter, r *http.Request) {
	var cm model.Classificacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSegmento] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	cm.CodigoClassificacao = int64(id)
	err = cm.GetClassificacao(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetClassificacao -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetClassificacao -  Erro ao tentar buscar segmento. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, cm, 0, 0)
}

//GetClassificacoes ...
func GetClassificacoes(w http.ResponseWriter, r *http.Request) {
	var cm model.Classificacao
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetClassificacoes] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	cm.CodigoClassificacao = int64(id)
	cm.DescricaoClassificacao = descricao

	classificacoes, err := cm.GetClassificacoes(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetClassificacoes -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há segmento cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetClassificacoes -  Erro ao tentar buscar segmentos. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, classificacoes, 0, 0)
}
