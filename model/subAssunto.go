package model

import (
	"database/sql"
	"strings"
)

//SubAssunto struct
type SubAssunto struct {
	Codigo              int64  `json:"codigo"`
	CodigoAssunto       int64  `json:"codigoAssunto"`
	DescricaoSubAssunto string `json:"descricao"`
	StatusSubAssunto    int64  `json:"status"`
}

//InsertSubAssunto ...
func (a *SubAssunto) InsertSubAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO SUBASSUNTO (CODIGOASSUNTO, DESCRICAO, STATUS)
								VALUES
								(?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(a.CodigoAssunto, a.DescricaoSubAssunto, a.StatusSubAssunto)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	a.Codigo = id
	return nil
}

//UpdateSubAssunto ...
func (a *SubAssunto) UpdateSubAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SUBASSUNTO 
								  SET CODIGOASSUNTO = ?,
								  DESCRICAO = ?
								  WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(a.CodigoAssunto, a.DescricaoSubAssunto, a.Codigo)

	return err
}

//GetSubAssunto ...
func (a *SubAssunto) GetSubAssunto(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, CODIGOASSUNTO, DESCRICAO, STATUS
					FROM SUBASSUNTO
					WHERE CODIGO =  ?`, a.Codigo).Scan(&a.CodigoAssunto, &a.DescricaoSubAssunto, &a.StatusSubAssunto)
	if err != nil {
		return err
	}

	return err
}

//GetSubAssuntos ...
func (a *SubAssunto) GetSubAssuntos(db *sql.DB) ([]SubAssunto, error) {
	var values []interface{}
	var where []string

	if a.Codigo != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, a.Codigo)
	}

	if a.DescricaoSubAssunto != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, a.DescricaoSubAssunto)
	}

	rows, err := db.Query(`SELECT CODIGO, CODIGOASSUNTO, DESCRICAO, STATUS
					FROM SUBASSUNTO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	SubAssuntos := []SubAssunto{}
	defer rows.Close()
	for rows.Next() {
		var sa SubAssunto
		if err = rows.Scan(&sa.Codigo, &sa.CodigoAssunto, &sa.DescricaoSubAssunto, &sa.StatusSubAssunto); err != nil {
			return nil, err
		}
		SubAssuntos = append(SubAssuntos, sa)
	}
	return SubAssuntos, nil
}

//DeleteSubAssunto ...
func (a *SubAssunto) DeleteSubAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SUBASSUNTO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, a.Codigo)

	return err
}
