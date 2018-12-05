package model

import (
	"database/sql"
	"strings"
)

//Demanda struct
type Demanda struct {
	CodigoDemanda    int64  `json:"codigo"`
	DescricaoDemanda string `json:"descricao"`
	StatusDemanda    int64  `json:"status"`
}

//InsertDemanda ...
func (d *Demanda) InsertDemanda(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO DEMANDA (DESCRICAO, STATUS)
								VALUES
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(d.DescricaoDemanda, d.StatusDemanda)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	d.CodigoDemanda = id
	return nil
}

//UpdateDemanda ...
func (d *Demanda) UpdateDemanda(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE DEMANDA SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(d.DescricaoDemanda, d.CodigoDemanda)

	return err
}

//GetDemanda ...
func (d *Demanda) GetDemanda(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM DEMANDA
					WHERE CODIGO =  ?`, d.CodigoDemanda).Scan(&d.CodigoDemanda, &d.DescricaoDemanda, &d.StatusDemanda)
	if err != nil {
		return err
	}

	return err
}

//GetDemandas ...
func (d *Demanda) GetDemandas(db *sql.DB) ([]Demanda, error) {
	var values []interface{}
	var where []string

	if d.CodigoDemanda != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, d.CodigoDemanda)
	}

	if d.DescricaoDemanda != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, d.DescricaoDemanda)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM DEMANDA
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

	if err != nil {
		return nil, err
	}

	demandas := []Demanda{}
	defer rows.Close()
	for rows.Next() {
		var dm Demanda
		if err = rows.Scan(&dm.CodigoDemanda, &dm.DescricaoDemanda, &dm.StatusDemanda); err != nil {
			return nil, err
		}
		demandas = append(demandas, dm)
	}
	return demandas, nil
}

//DeleteDemanda ...
func (d *Demanda) DeleteDemanda(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE DEMANDA SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, d.CodigoDemanda)

	return err
}
