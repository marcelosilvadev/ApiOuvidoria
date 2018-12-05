package model

import (
	"database/sql"
	"strings"
)

//TipoManifestacao struct
type TipoManifestacao struct {
	CodigoTipoManifestacao    int64  `json:"codigo"`
	DescricaoTipoManifestacao string `json:"descricao"`
	StatusTipoManifestacao    int64  `json:"status"`
}

//InsertTipoManifestacao ...
func (tm *TipoManifestacao) InsertTipoManifestacao(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO TIPOMANIFESTACAO (DESCRICAO, STATUS)
								VALUES
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(tm.DescricaoTipoManifestacao, tm.StatusTipoManifestacao)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	tm.CodigoTipoManifestacao = id
	return nil
}

//UpdateTipoManifestacao ...
func (tm *TipoManifestacao) UpdateTipoManifestacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE TIPOMANIFESTACAO SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(tm.DescricaoTipoManifestacao, tm.CodigoTipoManifestacao)

	return err
}

//GetTipoManifestacao ...
func (tm *TipoManifestacao) GetTipoManifestacao(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM TIPOMANIFESTACAO
					WHERE CODIGO =  ?`, tm.CodigoTipoManifestacao).Scan(&tm.CodigoTipoManifestacao, &tm.DescricaoTipoManifestacao, &tm.StatusTipoManifestacao)
	if err != nil {
		return err
	}

	return err
}

//GetTipoManifestacoes ...
func (tm *TipoManifestacao) GetTipoManifestacoes(db *sql.DB) ([]TipoManifestacao, error) {
	var values []interface{}
	var where []string

	if tm.CodigoTipoManifestacao != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, tm.CodigoTipoManifestacao)
	}

	if tm.DescricaoTipoManifestacao != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, tm.DescricaoTipoManifestacao)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM TIPOMANIFESTACAO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	manifestacoes := []TipoManifestacao{}
	defer rows.Close()
	for rows.Next() {
		var tmo TipoManifestacao
		if err = rows.Scan(&tmo.CodigoTipoManifestacao, &tmo.DescricaoTipoManifestacao, &tmo.StatusTipoManifestacao); err != nil {
			return nil, err
		}
		manifestacoes = append(manifestacoes, tmo)
	}
	return manifestacoes, nil
}

//DeleteTipoManifestacao ...
func (tm *TipoManifestacao) DeleteTipoManifestacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE TIPOMANIFESTACAO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, tm.CodigoTipoManifestacao)

	return err
}
