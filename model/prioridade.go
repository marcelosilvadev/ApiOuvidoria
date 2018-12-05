package model

import (
	"database/sql"
	"strings"
)

//Prioridade struct
type Prioridade struct {
	CodigoPrioridade    int64  `json:"codigo"`
	DescricaoPrioridade string `json:"descricao"`
	StatusPrioridade    int64  `json:"status"`
}

//InsertPrioridade ...
func (p *Prioridade) InsertPrioridade(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO PRIORIDADE (DESCRICAO, STATUS)
									VALUES (?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(p.DescricaoPrioridade, p.StatusPrioridade)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	p.CodigoPrioridade = id
	return nil
}

//UpdatePrioridade ...
func (p *Prioridade) UpdatePrioridade(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE PRIORIDADE SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(p.DescricaoPrioridade, p.CodigoPrioridade)

	return err
}

//GetPrioridade ...
func (p *Prioridade) GetPrioridade(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM PRIORIDADE
					WHERE CODIGO =  ?`, p.CodigoPrioridade).Scan(&p.CodigoPrioridade, &p.DescricaoPrioridade, &p.StatusPrioridade)
	if err != nil {
		return err
	}

	return err
}

//GetPrioridades ...
func (p *Prioridade) GetPrioridades(db *sql.DB) ([]Prioridade, error) {
	var values []interface{}
	var where []string

	if p.CodigoPrioridade != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, p.CodigoPrioridade)
	}

	if p.DescricaoPrioridade != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, p.DescricaoPrioridade)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM PRIORIDADE
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

	if err != nil {
		return nil, err
	}

	prioridade := []Prioridade{}
	defer rows.Close()
	for rows.Next() {
		var pri Prioridade
		if err = rows.Scan(&pri.CodigoPrioridade, &pri.DescricaoPrioridade, &pri.StatusPrioridade); err != nil {
			return nil, err
		}
		prioridade = append(prioridade, pri)
	}
	return prioridade, nil
}

//DeletePrioridade ...
func (p *Prioridade) DeletePrioridade(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE PRIORIDADE SET STATUS = 1 WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, p.CodigoPrioridade)

	return err
}
