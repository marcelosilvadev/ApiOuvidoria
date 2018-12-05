package model

import (
	"database/sql"
	"strings"
)

//Origem struct
type Origem struct {
	CodigoOrigem    int64  `json:"codigo"`
	CodigoSegmento  int64  `json:"segmento"`
	DescricaoOrigem string `json:"descricao"`
	Cep             string `json:"cep"`
	Endereco        string `json:"endereco"`
	Numero          string `json:"numero"`
	Bairro          string `json:"bairro"`
	Cidade          string `json:"cidade"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	StatusOrigem    int64  `json:"status"`
}

//InsertOrigem ...
func (o *Origem) InsertOrigem(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO ORIGEM (SEGMENTO, DESCRICAO, CEP , ENDERECO, NUMERO,
								  BAIRRO, CIDADE, LATITUDE, LONGITUDE, STATUS)
								VALUES
								(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(o.CodigoSegmento, o.DescricaoOrigem, o.Cep, o.Endereco, o.Numero,
		o.Bairro, o.Cidade, o.Latitude, o.Longitude, o.StatusOrigem)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	o.CodigoOrigem = id
	return nil
}

//UpdateOrigem ...
func (o *Origem) UpdateOrigem(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE ORIGEM 
								  SET SEGMENTO = ?,
								  DESCRICAO = ?,
								  CEP = ?,
								  ENDERECO = ?,
								  NUMERO = ?,
								  BAIRRO = ?,
								  CIDADE = ?,
								  LATITUDE = ?,
								  LONGITUDE = ?
								  WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(o.CodigoSegmento, o.DescricaoOrigem, o.Cep, o.Endereco, o.Numero,
		o.Bairro, o.Cidade, o.Latitude, o.Longitude, o.CodigoOrigem)

	return err
}

//GetOrigem ...
func (o *Origem) GetOrigem(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, SEGMENTO, DESCRICAO, CEP, 
						ENDERECO, NUMERO, BAIRRO, CIDADE, LATITUDE, LONGITUDE, STATUS
					FROM ORIGEM
					WHERE CODIGO =  ?`, o.CodigoOrigem).Scan(&o.CodigoOrigem, &o.CodigoSegmento, &o.DescricaoOrigem, &o.Cep,
		&o.Endereco, &o.Numero, &o.Bairro, &o.Cidade, &o.Latitude,
		&o.Longitude, &o.CodigoOrigem)
	if err != nil {
		return err
	}

	return err
}

//GetOrigens ...
func (o *Origem) GetOrigens(db *sql.DB) ([]Origem, error) {
	var values []interface{}
	var where []string

	if o.CodigoOrigem != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, o.CodigoOrigem)
	}

	if o.DescricaoOrigem != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, o.DescricaoOrigem)
	}

	rows, err := db.Query(`SELECT CODIGO, SEGMENTO, DESCRICAO, CEP, 
						   ENDERECO, NUMERO, BAIRRO, CIDADE, LATITUDE, LONGITUDE, STATUS
					FROM ORIGEM
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	Origens := []Origem{}
	defer rows.Close()
	for rows.Next() {
		var or Origem
		if err = rows.Scan(&or.CodigoOrigem, &or.CodigoSegmento, &or.DescricaoOrigem, &or.Cep, &or.Endereco,
			&or.Numero, &or.Bairro, &or.Cidade, &or.Latitude, &or.Longitude, &or.StatusOrigem); err != nil {
			return nil, err
		}
		Origens = append(Origens, or)
	}
	return Origens, nil
}

//DeleteOrigem ...
func (o *Origem) DeleteOrigem(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE ORIGEM SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, o.CodigoOrigem)

	return err
}
