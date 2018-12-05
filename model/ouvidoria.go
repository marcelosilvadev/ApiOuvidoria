package model

import (
	"database/sql"
	"time"
)

//Imagem struct
type Imagem struct {
	Codigo int64  `json:"codigo"`
	Imagem string `json:"imagem"`
}

//Ouvidoria struct
type Ouvidoria struct {
	Ano                 int64    `json:"ano"`
	Sequencia           int64    `json:"sequencia"`
	DataInclusao        string   `json:"dataInclusao"`
	CodUsuario          int64    `json:"codUsuario"`
	CodClassificacao    int64    `json:"codClassificacao"`
	CodSegmento         int64    `json:"codSegmento"`
	CodOrigem           int64    `json:"codOrigem"`
	CodAssunto          int64    `json:"codigoAssunto"`
	CodSituacao         string   `json:"codSituacao"`
	Imagens             []Imagem `json:"imagens"`
	TextoOuvidoria      string   `json:"textoOuvidoria"`
	CartaResposta       string   `json:"cartaResposta"`
	CodDemanda          int64    `json:"codDemanda"`
	Matricula           int64    `json:"matricula"`
	CodTipoManifestacao int64    `json:"codTipoManifestacao"`
	CodPrioridade       int64   `json:"codPrioridade"`
}

//InsertOuvidoria ...
func (o *Ouvidoria) InsertOuvidoria(db *sql.DB) error {
	date := time.Now()

	db.QueryRow(`SELECT Codigo FROM Situacao WHERE Descricao = 'Encaminhado'`).Scan(&o.CodSituacao)
	db.QueryRow(`SELECT Codigo FROM Demanda WHERE Descricao = 'APP'`).Scan(&o.CodDemanda)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`INSERT INTO OUVIDORIA (Ano, 
													DataInclusao,
													CodUsuario, 
													CodClassificacao,
													CodSegmento,
													CodOrigem,
													CodAssunto,
													CodSituacao,
													TextoOuvidoria,
													CodDemanda)
								VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(date.Year(), date, o.CodUsuario, o.CodClassificacao, o.CodSegmento, o.CodOrigem, o.CodAssunto, o.CodSituacao, o.TextoOuvidoria, o.CodDemanda)
	if err != nil {
		return err
	}
	sequencia, _ := res.LastInsertId()

	o.Sequencia = sequencia
	o.DataInclusao = date.Format("2006-01-02T15:04:05-0700")
	o.Ano = int64(date.Year())

	stmt, err = tx.Prepare(`INSERT INTO ImagemOuvidoria (AnoOuvidoria, SequenciaOuvidoria, Imagem) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	for i, imagem := range o.Imagens {
		res, err = stmt.Exec(date.Year(), sequencia, imagem.Imagem)
		if err != nil {
			return err
		}
		codigo, _ := res.LastInsertId()
		o.Imagens[i].Codigo = codigo
	}

	tx.Commit()

	return nil
}
