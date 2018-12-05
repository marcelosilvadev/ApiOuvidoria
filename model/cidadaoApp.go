package model

import (
	"database/sql"
	"strings"
)

//Cidadao struct
type Cidadao struct {
	CodigoCidadao  int64  `json:"codigo"`
	Nome           string `json:"nome"`
	Cpf            string `json:"cpf"`
	DataNascimento string `json:"dataNascimento"`
	Telefone       string `json:"telefone"`
	Senha          string `json:"senha"`
	StatusCidadao  int64  `json:"status"`
}

//InsertCidadao ...
func (c *Cidadao) InsertCidadao(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO CIDADAO (NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS)
								VALUES
								(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(c.Nome, c.Cpf, c.DataNascimento, c.Telefone, c.StatusCidadao)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.CodigoCidadao = id
	return nil
}

//UpdateCidadao ...
func (c *Cidadao) UpdateCidadao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CIDADAO
									SET NOME = ?,
										CPF = ?,
										DATANASCIMENTO = ?,
										TELEFONE = ?, 
									WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(c.Nome, c.Cpf, c.DataNascimento, c.Telefone, c.StatusCidadao, c.CodigoCidadao)

	return err
}

//GetCidadao ...
func (c *Cidadao) GetCidadao(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
					FROM CIDADAO
					WHERE CODIGO =  ?`, c.CodigoCidadao).Scan(&c.CodigoCidadao, &c.Nome, &c.Cpf, &c.DataNascimento, &c.Telefone, &c.StatusCidadao)
	if err != nil {
		return err
	}

	return err
}

//GetCidadaos ...
func (c *Cidadao) GetCidadaos(db *sql.DB) ([]Cidadao, error) {
	var values []interface{}
	var where []string

	if c.CodigoCidadao != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, c.CodigoCidadao)
	}

	if c.Nome != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, c.Nome)
	}

	rows, err := db.Query(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
					FROM CIDADAO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

	if err != nil {
		return nil, err
	}

	cidadaos := []Cidadao{}
	defer rows.Close()
	for rows.Next() {
		var cid Cidadao
		if err = rows.Scan(&cid.CodigoCidadao, &cid.Nome, &cid.Cpf, &cid.DataNascimento, &cid.Telefone, &cid.StatusCidadao); err != nil {
			return nil, err
		}
		cidadaos = append(cidadaos, cid)
	}
	return cidadaos, nil
}

//DeleteCidadao ...
func (c *Cidadao) DeleteCidadao(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CIDADAO
									SET STATUS = 1
									WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, c.CodigoCidadao)

	return err
}
