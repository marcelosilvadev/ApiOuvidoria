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

//InsertSegmento ...
func InsertSegmento(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var s model.Segmento
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = s.InsertSegmento(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Segmento", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, s, 0, 0)
}

//UpdateSegmento ...
func UpdateSegmento(w http.ResponseWriter, r *http.Request) {
	var s model.Segmento
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateSegmento] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&s); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	s.CodigoSegmento = int64(id)
	if err := s.UpdateSegmento(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, s, 0, 0)
}

//DeleteSegmento ...
func DeleteSegmento(w http.ResponseWriter, r *http.Request) {
	var s model.Segmento
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteSegmento -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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

	s.CodigoSegmento = int64(id)
	if err := s.DeleteSegmento(db); err != nil {
		log.Printf("[handler/DeleteSegmento -  Erro ao tentar deletar segmento. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, s, 0, 0)
}

//GetSegmento ...
func GetSegmento(w http.ResponseWriter, r *http.Request) {
	var s model.Segmento
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

	s.CodigoSegmento = int64(id)
	err = s.GetSegmento(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSegmento -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetSegmento -  Erro ao tentar buscar segmento. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, s, 0, 0)
}

//GetSegmentos ...
func GetSegmentos(w http.ResponseWriter, r *http.Request) {
	var s model.Segmento
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetSegmentos] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	descricao := r.FormValue("descricao")

	s.CodigoSegmento = int64(id)
	s.DescricaoSegmento = descricao

	segmentos, err := s.GetSegmentos(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetSegmentos -  Não há segmento com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há segmento cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetSegmentost -  Erro ao tentar buscar segmentos. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, segmentos, 0, 0)
}
