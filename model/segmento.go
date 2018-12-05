package model

import (
	"database/sql"
	"strings"
)

//Segmento struct
type Segmento struct {
	CodigoSegmento    int64  `json:"codigo"`
	DescricaoSegmento string `json:"descricao"`
	StatusSegmento    int64  `json:"status"`
}

//InsertSegmento ...
func (s *Segmento) InsertSegmento(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO SEGMENTO (DESCRICAO, STATUS)
								VALUES
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(s.DescricaoSegmento, s.StatusSegmento)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	s.CodigoSegmento = id
	return nil
}

//UpdateSegmento ...
func (s *Segmento) UpdateSegmento(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SEGMENTO SET DESCRICAO = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(s.DescricaoSegmento, s.CodigoSegmento)

	return err
}

//GetSegmento ...
func (s *Segmento) GetSegmento(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, DESCRICAO, STATUS
					FROM SEGMENTO
					WHERE CODIGO =  ?`, s.CodigoSegmento).Scan(&s.CodigoSegmento, &s.DescricaoSegmento, &s.StatusSegmento)
	if err != nil {
		return err
	}

	return err
}

//GetSegmentos ...
func (s *Segmento) GetSegmentos(db *sql.DB) ([]Segmento, error) {
	var values []interface{}
	var where []string

	if s.CodigoSegmento != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, s.CodigoSegmento)
	}

	if s.DescricaoSegmento != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, s.DescricaoSegmento)
	}

	rows, err := db.Query(`SELECT CODIGO, DESCRICAO, STATUS
					FROM SEGMENTO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	segmentos := []Segmento{}
	defer rows.Close()
	for rows.Next() {
		var sg Segmento
		if err = rows.Scan(&sg.CodigoSegmento, &sg.DescricaoSegmento, &sg.StatusSegmento); err != nil {
			return nil, err
		}
		segmentos = append(segmentos, sg)
	}
	return segmentos, nil
}

//DeleteSegmento ...
func (s *Segmento) DeleteSegmento(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE SEGMENTO SET STATUS = ? WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, s.CodigoSegmento)

	return err
}
