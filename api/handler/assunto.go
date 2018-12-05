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

//InsertAssunto ...
func InsertAssunto(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var a model.Assunto
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = a.InsertAssunto(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Assunto", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}

//UpdateAssunto ...
func UpdateAssunto(w http.ResponseWriter, r *http.Request) {
	var a model.Assunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateAssunto] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&a); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	a.CodigoAssunto = int64(id)
	if err := a.UpdateAssunto(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}

//DeleteAssunto ...
func DeleteAssunto(w http.ResponseWriter, r *http.Request) {
	var a model.Assunto
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteAssunto -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	a.CodigoAssunto = int64(id)
	if err := a.DeleteAssunto(db); err != nil {
		log.Printf("[handler/DeleteAssunto -  Erro ao tentar deletar Assunto. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}

//GetAssunto ...
func GetAssunto(w http.ResponseWriter, r *http.Request) {
	var a model.Assunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetAssunto] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	a.CodigoAssunto = int64(id)
	err = a.GetAssunto(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetAssunto -  Não há Assunto com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há assunto com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetAssunto -  Erro ao tentar buscar assunto. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}

//GetAssuntos ...
func GetAssuntos(w http.ResponseWriter, r *http.Request) {
	var a model.Assunto
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetAssuntos] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")
	codSegmento, _ := strconv.Atoi(r.FormValue("codSegmento"))

	a.CodigoAssunto = int64(id)
	a.DescricaoAssunto = descricao
	a.CodigoSegmento = int64(codSegmento)

	assuntos, err := a.GetAssuntos(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetAssuntos -  Não há Assunto com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Demanda cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetAssuntos -  Erro ao tentar buscar Assunto. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, assuntos, 0, 0)
}
