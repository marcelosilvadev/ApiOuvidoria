package model

import (
	"database/sql"
	"strings"
)

//Situacao struct
type Situacao struct {
	CodigoSituacao    int64  `json:"codigo"`
	DescricaoSituacao string `json:"descricao"`
	StatusSituacao    int64  `json:"status"`
}

//InsertSituacao ...
func (si *Situacao) InsertSituacao(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO SITUACAO (DESCRICAO, STATUS)
								VALUES
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(si.DescricaoSituacao, si.StatusSituacao)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	si.CodigoSituacao = id
	return nil
}

//UpdateSituacao ...
func (si *Situacao) UpdateSituacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SITUACAO SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(si.DescricaoSituacao, si.CodigoSituacao)

	return err
}

//GetSituacao ...
func (si *Situacao) GetSituacao(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM SITUACAO
					WHERE CODIGO =  ?`, si.CodigoSituacao).Scan(&si.CodigoSituacao, &si.DescricaoSituacao, &si.StatusSituacao)
	if err != nil {
		return err
	}

	return err
}

//GetSituacoes ...
func (si *Situacao) GetSituacoes(db *sql.DB) ([]Situacao, error) {
	var values []interface{}
	var where []string

	if si.CodigoSituacao != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, si.CodigoSituacao)
	}

	if si.DescricaoSituacao != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, si.DescricaoSituacao)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM SITUACAO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	situacoes := []Situacao{}
	defer rows.Close()
	for rows.Next() {
		var st Situacao
		if err = rows.Scan(&st.CodigoSituacao, &st.DescricaoSituacao, &st.StatusSituacao); err != nil {
			return nil, err
		}
		situacoes = append(situacoes, st)
	}
	return situacoes, nil
}

//DeleteSituacao ...
func (si *Situacao) DeleteSituacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SITUACAO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, si.CodigoSituacao)

	return err
}
