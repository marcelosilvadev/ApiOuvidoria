package model

import (
	"database/sql"
	"strings"
)

type Assunto struct {
	CodigoAssunto    int64  `json:"codigo"`
	DescricaoAssunto string `json:"descricao"`
	CodigoSegmento   int64  `json:"codSegmento"`
	StatusAssunto    int64  `json:"status"`
}

//InsertAssunto ...
func (a *Assunto) InsertAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO ASSUNTO (DESCRICAO, CODSEGMENTO, STATUS)
								VALUES
								(?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(a.DescricaoAssunto, a.CodigoSegmento, a.StatusAssunto)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	a.CodigoAssunto = id
	return nil
}

//UpdateAssunto ...
func (a *Assunto) UpdateAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE ASSUNTO SET DESCRICAO = ?, CODSEGMENTO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(a.DescricaoAssunto, a.CodigoSegmento, a.CodigoAssunto)

	return err
}

//GetAssunto ...
func (a *Assunto) GetAssunto(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, CODSEGMENTO, STATUS
					FROM ASSUNTO
					WHERE CODIGO =  ?`, a.CodigoAssunto).Scan(&a.CodigoAssunto, &a.DescricaoAssunto, &a.CodigoSegmento, &a.StatusAssunto)
	if err != nil {
		return err
	}

	return err
}

//GetAssuntos ...
func (a *Assunto) GetAssuntos(db *sql.DB) ([]Assunto, error) {
	var values []interface{}
	var where []string

	if a.CodigoAssunto != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, a.CodigoAssunto)
	}

	if a.DescricaoAssunto != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, a.DescricaoAssunto)
	}

	if a.CodigoSegmento != 0 {
		where = append(where, "CODSEGMENTO = ?")
		values = append(values, a.CodigoSegmento)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, CODSEGMENTO, STATUS
					FROM ASSUNTO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

	if err != nil {
		return nil, err
	}

	assuntos := []Assunto{}
	defer rows.Close()
	for rows.Next() {
		var as Assunto
		if err = rows.Scan(&as.CodigoAssunto, &as.DescricaoAssunto, &as.CodigoSegmento, &as.StatusAssunto); err != nil {
			return nil, err
		}
		assuntos = append(assuntos, as)
	}
	return assuntos, nil
}

//DeleteAssunto ...
func (a *Assunto) DeleteAssunto(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE ASSUNTO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, a.CodigoAssunto)

	return err
}
