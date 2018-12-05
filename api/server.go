package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/ApiOuvidoria/api/handler"
	"github.com/gorilla/mux"
)

// App struct ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//StartServer ...
func (a *App) StartServer() {
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix("/api/v1").Subrouter()
	//Health para teste de funcionamento da API
	s.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)

	//EndPoints para Segmentos da ouvidoria
	s.HandleFunc("/segmento", handler.InsertSegmento).Methods(http.MethodPost)
	s.HandleFunc("/segmento/{id:[0-9]+}", handler.UpdateSegmento).Methods(http.MethodPut)
	s.HandleFunc("/segmento/{id:[0-9]+}", handler.DeleteSegmento).Methods(http.MethodDelete)
	s.HandleFunc("/segmento/{id:[0-9]+}", handler.GetSegmento).Methods(http.MethodGet)
	s.HandleFunc("/segmento", handler.GetSegmentos).Methods(http.MethodGet)

	//EndPoints para Situação da ouvidoria
	s.HandleFunc("/situacao", handler.InsertSituacao).Methods(http.MethodPost)
	s.HandleFunc("/situacao/{id:[0-9]+}", handler.UpdateSituacao).Methods(http.MethodPut)
	s.HandleFunc("/situacao/{id:[0-9]+}", handler.DeleteSituacao).Methods(http.MethodDelete)
	s.HandleFunc("/situacao/{id:[0-9]+}", handler.GetSituacao).Methods(http.MethodGet)
	s.HandleFunc("/situacao", handler.GetSituacoes).Methods(http.MethodGet)

	//EndPoints para Tipo Manifestacao da ouvidoria
	s.HandleFunc("/tipomanifestacao", handler.InsertTipoManifestacao).Methods(http.MethodPost)
	s.HandleFunc("/tipomanifestacao/{id:[0-9]+}", handler.UpdateTipoManifestacaoo).Methods(http.MethodPut)
	s.HandleFunc("/tipomanifestacao/{id:[0-9]+}", handler.DeleteTipoManifestacao).Methods(http.MethodDelete)
	s.HandleFunc("/tipomanifestacao/{id:[0-9]+}", handler.GetTipoManifestacao).Methods(http.MethodGet)
	s.HandleFunc("/tipomanifestacao", handler.GetTipoManifestacoes).Methods(http.MethodGet)

	//EndPoints para Demanda da ouvidoria
	s.HandleFunc("/demanda", handler.InsertDemanda).Methods(http.MethodPost)
	s.HandleFunc("/demanda/{id:[0-9]+}", handler.UpdateDemanda).Methods(http.MethodPut)
	s.HandleFunc("/demanda/{id:[0-9]+}", handler.DeleteDemanda).Methods(http.MethodDelete)
	s.HandleFunc("/demanda/{id:[0-9]+}", handler.GetDemanda).Methods(http.MethodGet)
	s.HandleFunc("/demanda", handler.GetDemandas).Methods(http.MethodGet)

	//EndPoints para Prioridade da ouvidoria
	s.HandleFunc("/prioridade", handler.InsertPrioridade).Methods(http.MethodPost)
	s.HandleFunc("/prioridade/{id:[0-9]+}", handler.UpdatePrioridade).Methods(http.MethodPut)
	s.HandleFunc("/prioridade/{id:[0-9]+}", handler.DeletePrioridade).Methods(http.MethodDelete)
	s.HandleFunc("/prioridade/{id:[0-9]+}", handler.GetPrioridade).Methods(http.MethodGet)
	s.HandleFunc("/prioridade", handler.GetPrioridades).Methods(http.MethodGet)

	//EndPoints para Cidadao da ouvidoria
	s.HandleFunc("/cidadao", handler.InsertCidadao).Methods(http.MethodPost)
	s.HandleFunc("/cidadao/{id:[0-9]+}", handler.UpdateCidadao).Methods(http.MethodPut)
	s.HandleFunc("/cidadao/{id:[0-9]+}", handler.DeleteCidadao).Methods(http.MethodDelete)
	s.HandleFunc("/cidadao/{id:[0-9]+}", handler.GetCidadao).Methods(http.MethodGet)
	s.HandleFunc("/cidadao", handler.GetCidadaos).Methods(http.MethodGet)

	//EndPoints para classificacao da ouvidoria
	s.HandleFunc("/classificacao", handler.InsertClassificacao).Methods(http.MethodPost)
	s.HandleFunc("/classificacao/{id:[0-9]+}", handler.UpdateClassificacao).Methods(http.MethodPut)
	s.HandleFunc("/classificacao/{id:[0-9]+}", handler.DeleteClassificacao).Methods(http.MethodDelete)
	s.HandleFunc("/classificacao/{id:[0-9]+}", handler.GetClassificacao).Methods(http.MethodGet)
	s.HandleFunc("/classificacao", handler.GetClassificacoes).Methods(http.MethodGet)

	//EndPoints para assunto da ouvidoria
	s.HandleFunc("/assunto", handler.InsertAssunto).Methods(http.MethodPost)
	s.HandleFunc("/assunto/{id:[0-9]+}", handler.UpdateAssunto).Methods(http.MethodPut)
	s.HandleFunc("/assunto/{id:[0-9]+}", handler.DeleteAssunto).Methods(http.MethodDelete)
	s.HandleFunc("/assunto/{id:[0-9]+}", handler.GetAssunto).Methods(http.MethodGet)
	s.HandleFunc("/assunto", handler.GetAssunto).Methods(http.MethodGet)

	//EndPoints para origem da ouvidoria
	s.HandleFunc("/origem", handler.InsertOrigem).Methods(http.MethodPost)
	s.HandleFunc("/origem/{id:[0-9]+}", handler.UpdateOrigem).Methods(http.MethodPut)
	s.HandleFunc("/origem/{id:[0-9]+}", handler.DeleteOrigem).Methods(http.MethodDelete)
	s.HandleFunc("/origem/{id:[0-9]+}", handler.GetOrigem).Methods(http.MethodGet)
	s.HandleFunc("/origem", handler.GetOrigens).Methods(http.MethodGet)

	//EndPoints para origem da ouvidoria
	s.HandleFunc("/SubAssunto", handler.InsertOrigem).Methods(http.MethodPost)
	s.HandleFunc("/SubAssunto/{id:[0-9]+}", handler.UpdateSubAssunto).Methods(http.MethodPut)
	s.HandleFunc("/SubAssunto/{id:[0-9]+}", handler.DeleteSubAssunto).Methods(http.MethodDelete)
	s.HandleFunc("/SubAssunto/{id:[0-9]+}", handler.GetSubAssunto).Methods(http.MethodGet)
	s.HandleFunc("/SubAssunto", handler.GetSubAssuntos).Methods(http.MethodGet)

	//EndPoints para ouvidoria
	s.HandleFunc("/ouvidoria", handler.InsertOuvidoria).Methods(http.MethodPost)
	// s.HandleFunc("/ouvidoria/{id:[0-9]+}", handler.UpdateClassificacao).Methods(http.MethodPut)
	// s.HandleFunc("/ouvidoria/{id:[0-9]+}", handler.DeleteClassificacao).Methods(http.MethodDelete)
	// s.HandleFunc("/ouvidoria/{id:[0-9]+}", handler.GetClassificacao).Methods(http.MethodGet)
	// s.HandleFunc("/ouvidoria", handler.GetClassificacoes).Methods(http.MethodGet)

	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 8081
	log.Printf("Starting Server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}
