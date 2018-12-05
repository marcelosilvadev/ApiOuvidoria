CREATE DATABASE	OUVIDORIA;

USE OUVIDORIA;

CREATE TABLE Ouvidoria (
    Ano INT,
    Sequencia INT AUTO_INCREMENT,
    DataInclusao DATETIME,
    CodUsuario INT,
    CodClassificacao INT, -- Tipo manifestação na tela
    CodSegmento INT,
    CodOrigem INT,
    CodAssunto INT,
    CodSituacao VARCHAR(25),
    TextoOuvidoria BLOB,
    CartaResposta BLOB,
    CodPrioridade VARCHAR(20),
    CodDemanda INT,
    CodTipoManifestacao INT,
    Matricula INT,
    CONSTRAINT PK_Ouvidoria_Ano_Sequencia PRIMARY KEY (Ano , Sequencia),
    CONSTRAINT fk_Ouvidoria_CodOrigem FOREIGN KEY (CodOrigem)
        REFERENCES Origem (Codigo),
    CONSTRAINT fk_OuvidoriaAssunto_CodAssunto FOREIGN KEY (CodAssunto)
        REFERENCES Assunto (Codigo),
    CONSTRAINT fk_OuvidoriaDemanda_CodDemanda FOREIGN KEY (CodDemanda)
        REFERENCES Demanda (Codigo),
    CONSTRAINT fk_OuvidoriaTipoManifestacao_CodTipoManifestacao FOREIGN KEY (CodTipoManifestacao)
        REFERENCES TipoManifestacao (Codigo),
    CONSTRAINT fk_OuvidoriaClassificacao_CodClassificacao FOREIGN KEY (CodClassificacao)
        REFERENCES Classificacao (Codigo),
    CONSTRAINT fk_OuvidoriaPrioridade_CodPrioridade FOREIGN KEY (CodPrioridade)
        REFERENCES Prioridade (Codigo),
    CONSTRAINT fk_OuvidoriaSituacao_CodSituacao FOREIGN KEY (CodSituacao)
        REFERENCES Situacao (Codigo),
    CONSTRAINT fk_OuvidoriaUsuario_CodUsuario FOREIGN KEY (CodUsuario)
        REFERENCES Cidadao (Codigo),
    CONSTRAINT fk_OuvidoriaSegmento_CodSegmento FOREIGN KEY (CodSegmento)
        REFERENCES Segmento (Codigo)
);

CREATE TABLE ImagemOuvidoria(
    Codigo INT AUTO_INCREMENT,
    AnoOuvidoria INT,
    SequenciaOuvidoria INT,
    Imagem LONGBLOB,
    CONSTRAINT PK_ImagemOuvidoria_Codigo PRIMARY KEY (codigo),
    CONSTRAINT FK_ImagemOuvidoria FOREIGN KEY (AnoOuvidoria, SequenciaOuvidoria) 
        REFERENCES Ouvidoria (Ano, Sequencia)
);

CREATE TABLE Segmento(
    Codigo  INT NOT NULL AUTO_INCREMENT,
    Descricao VARCHAR (30),
    Status INT,
    CONSTRAINT PK_Segmento_Codigo PRIMARY KEY (Codigo)
);

-- Assunto da Manifestação
CREATE TABLE Assunto (
    Codigo INT NOT NULL AUTO_INCREMENT,
    Descricao VARCHAR(50),
    CodSegmento INT NULL,
    Status INT,
    CONSTRAINT PK_Assunto_Codigo PRIMARY KEY (Codigo),
    CONSTRAINT FK_AssuntoSegmento_CodSegmento FOREIGN KEY (CodSegmento)
        REFERENCES Segmento (Codigo)
);
 
-- Texto escrito sobre a sua manifestação
CREATE TABLE SubAssunto(
    Codigo INT AUTO_INCREMENT,
    CodigoAssunto INT,
    Descricao BLOB,
    Status INT,
    CONSTRAINT PK_SubAssunto_Codigo PRIMARY KEY (Codigo),
    CONSTRAINT FK_SubAssunto_CodigoAssunto FOREIGN KEY (CodigoAssunto)
	    REFERENCES Assunto (Codigo)
);

#Ambiente em que foi realizado a manifestação (APP,SIGS)
CREATE TABLE Demanda (
    Codigo INT NOT NULL AUTO_INCREMENT,
    Descricao VARCHAR(50) NULL,
    Status INT,
    CONSTRAINT PK_OuvidoriaCanalDemanda_Codigo PRIMARY KEY (Codigo)
);
 
 #Tipo da manifestação (Denuncia, Elogio, Sugestão, Reclamação, Solicitação)
 CREATE TABLE Classificacao (
    Codigo INT NOT NULL AUTO_INCREMENT,
    Descricao VARCHAR(20),
    Status INT,
    CONSTRAINT PK_Classificacao_Codigo PRIMARY KEY (Codigo)
);

#Local da manifestação
CREATE TABLE Origem(
    Codigo INT NOT NULL auto_increment,
    Segmento INT,
    Descricao VARCHAR(80),
    Cep varchar(10),
    Endereco VARCHAR (90),
    Numero VARCHAR(4),
    Bairro VARCHAR(60),
    Cidade VARCHAR(60),
    Latitude VARCHAR(20),
    Longitude VARCHAR(20),
    Status INT,
    CONSTRAINT pk_Origem_Codigo PRIMARY KEY (Codigo),
    CONSTRAINT fk_Origem_Segmento FOREIGN KEY (Segmento)
        REFERENCES Segmento (Codigo)
);
        
#Define quando será tratada essa manifestação
CREATE TABLE Prioridade(
    Codigo int NOT NULL auto_increment,
    Descricao varchar(20) NULL,
    Status INT,
    constraint pk_Prioridade_Codigo PRIMARY KEY (Codigo)
);

#Pedido já foi concluído, encaminhado, fechado, em andamento	
CREATE TABLE Situacao(
    Codigo int NOT NULL auto_increment,
    Descricao varchar(25) NULL,
    Status INT,
    constraint pk_Situacao_Codigo PRIMARY KEY (Codigo)
);

#Se é de sigilo, anônimo ou informado o cidadão
CREATE TABLE TipoManifestacao(
    Codigo int NOT NULL auto_increment,
    Descricao varchar(10) NULL,
    Status INT,
    constraint pk_TipoManifestacao_Codigo PRIMARY KEY (Codigo)); 
 
#Tabela de cadastro das pessoas pelo APP
CREATE TABLE Cidadao (
    Codigo INT NOT NULL AUTO_INCREMENT,
    Nome VARCHAR(60),
    Cpf VARCHAR(15),
    DataNascimento DATE,
    Telefone VARCHAR(12),
    Senha VARCHAR(30),
    Status INT,
    CONSTRAINT pk_Cidadao_Codigo PRIMARY KEY (Codigo)
);