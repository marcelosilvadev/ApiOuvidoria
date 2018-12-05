package model

import (
	"database/sql"
	"strings"
)

//Classificacao struct
type Classificacao struct {
	CodigoClassificacao    int64  `json:"codigo"`
	DescricaoClassificacao string `json:"descricao"`
	StatusClassificacao    int64  `json:"status"`
}

//InsertClassificacao ...
func (c *Classificacao) InsertClassificacao(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO CLASSIFICACAO (DESCRICAO, STATUS)
								VALUES
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(c.DescricaoClassificacao, c.StatusClassificacao)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.CodigoClassificacao = id
	return nil
}

//UpdateClassificacao ...
func (c *Classificacao) UpdateClassificacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CLASSIFICACAO SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(c.DescricaoClassificacao, c.CodigoClassificacao)

	return err
}

//GetClassificacao ...
func (c *Classificacao) GetClassificacao(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM SEGMENTO
					WHERE CODIGO =  ?`, c.CodigoClassificacao).Scan(&c.CodigoClassificacao, &c.DescricaoClassificacao, &c.StatusClassificacao)
	if err != nil {
		return err
	}

	return err
}

//GetClassificacoes ...
func (c *Classificacao) GetClassificacoes(db *sql.DB) ([]Classificacao, error) {
	var values []interface{}
	var where []string

	if c.CodigoClassificacao != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, c.CodigoClassificacao)
	}

	if c.DescricaoClassificacao != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, c.DescricaoClassificacao)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM CLASSIFICACAO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	classificacoes := []Classificacao{}
	defer rows.Close()
	for rows.Next() {
		var cl Classificacao
		if err = rows.Scan(&cl.CodigoClassificacao, &cl.DescricaoClassificacao, &cl.StatusClassificacao); err != nil {
			return nil, err
		}
		classificacoes = append(classificacoes, cl)
	}
	return classificacoes, nil
}

//DeleteClassificacao ...
func (c *Classificacao) DeleteClassificacao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CLASSIFICACAO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, c.CodigoClassificacao)

	return err
}
