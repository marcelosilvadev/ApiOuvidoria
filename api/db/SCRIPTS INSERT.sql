USE OUVIDORIA;

insert into Segmento(descricao, status)
values
("Saúde",1),
("Educação",1),
("Trânsito",1),
("Meio Ambiente",1);

INSERT INTO Origem (Codigo,Segmento,Descricao,Cep,Endereco,Numero,Bairro,Cidade,Latitude,Longitude,Status)
VALUES
(default, 1,	'AMBULATORIO DE SAUDE MENTAL ADULTO', '',	'RUA ALUIZIO PACHECO FERREIRA', '4810',  'JARDIM SANTA EUGENIA',  'FRANCA',	'-20.5191942',	'-47.3954618',	1),
(default, 1,	'NAIA - AMBULATORIO DE SAUDE MENTAL INFANTIL', '',	'RUA OUVIDOR FREIRE', '2109', 'CENTRO',  'FRANCA', '-20.5344865',	'-47.4022458', 1),
(default, 1, 'CASA DO DIABETICO', '',	'RUA DIONIZIO FACIOLLI',  '1148',  'CENTRO', 'FRANCA', '-20.5441972',	'-47.3964312', 1),
(default, 1, 	'CCI-CENTRO DE CONVIVENCIA DO IDOSO', '',	'RUA OUVIDOR FREIRE', '2109', 'CENTRO', 'FRANCA',	'-20.5344865',	'-47.4022458', 1),
(default, 1,	'CENTRO DE DIAGNOSTICO POR IMAGEM', '', 	'AV. FLAVIO ROCHA - DR.', '4780', 'JARDIM REDENTOR', 'FRANCA',	'-20.5106013',	'-47.4016564', 1),
(default, 1,	'CENTRO OFTALMOLOGICO', '',	'RUA ANTONIO DE PADUA FARIA - DR.', '2301', 'PROL. VILA INDUSTRIAL',  'FRANCA',	'-20.5468785',	'-47.3968249', 1),
(default, 1,	'CENTRO SAUDE I FRANCA', '',	'RUA OUVIDOR FREIRE', '2109', 'CENTRO', 'FRANCA', '-20.5344865',	'-47.4022458', 1),
(default, 1,	'CEREST', '', 'RUA MARECHAL DEODORO - FRANCA', '2097', 'BAIRRO SAO JOSE',  'FRANCA',	'-20.5391074',	'-47.3958258',  1),
(default, 1,	'ESF VERA CRUZ UBS LUIZA', '',	'AV. NELSON JAPAULO', '1731', 'JARDIM LUIZA', 'FRANCA',	'-20.4907152',	'-47.4302067', 1),
(default, 1, 	'LABORATORIO DE FRANCA', '',	'RUA OUVIDOR FREIRE', '2109',  'CENTRO', 'FRANCA',	'-20.5344865',	'-47.4022458', 1),
(default, 1,	'NGA - SAUDE AUDITIVA','',	'RUA ALUIZIO PACHECO FERREIRA', '4010', 'VILA TOTOLI', 'FRANCA',	'-20.5191942',	'-47.3954618', 1),
(default, 1,	'NGA FRANCA', '',	'RUA ALUIZIO PACHECO FERREIRA', '4010', 'VILA TOTOLI', 'FRANCA',	'-20.5191942',	'-47.3954618', 1),
(default, 1, 	'NUCLEO DE SAUDE DA FAMILIA CITY PETROPOL', '',	'RUA ALBERT SABIN', '1471', 'CITY PETROPOLIS',  'FRANCA',	'-20.4688501',	'-47.3963086', 1),
(default, 1,	'NUCLEO DE SAUDE DA FAMILIA ESMERALDA', '',	'RUA ANTONIO ALBINO DA SILVA', '160',  'PARQUE DAS ESMERALDAS', 'FRANCA',	'-20.5456236',	'-47.4511838',	1),
(default, 1,	'NUCLEO DE SAUDE DA FAMILIA PAINEIRAS',	'', 'RUA ANTONIO MARIANO', '230', 'JARDIM PAINEIRA', 'FRANCA',	'-20.4726087',	'-47.4157090', 1),
(default, 1,	'PRONTO SOCORRO DE REFERENCIA', '',	'AV. CHICO JULIO', '5125', 'VILA IMPERADOR', 'FRANCA',	'-20.5088663',	'-47.3996674', 1),
(default, 1,	'PRONTO SOCORRO INFANTIL', '',	'RUA ALUIZIO PACHECO FERREIRA', '3912', 'PARQUE SANTA ADELIA',  'FRANCA', '-20.5191942', '-47.3954618', 1),
(default, 1,	'SECRETARIA DE SAUDE','',	'AV. FLAVIO ROCHA - DR.', '4780', 'JARDIM REDENTOR', 'FRANCA',	'-20.5106013','-47.4016564',1),
(default, 1,	'UBS AEROPORTO I','',	'RUA CYRO EDUARDO ROSA FALEIROS - DR.', '601',  'JARDIM AEROPORTO',  'FRANCA',	'-20.5744292',	'-47.3726741', 1),
(default, 1,	'UBS AEROPORTO III','',	'RUA DENIZAR TREVIZANI', '1971', 'JARDIM AEROPORTO', 'FRANCA',	'-20.5961795',	'-47.3658093', 1),
(default, 1,	'UBS ANGELA ROSA', '',	'RUA ANGELA ROSA SCARABUCCI', '691', 'JARDIM ANGELA ROSA', 'FRANCA',	'-20.5454125',	'-47.3906371',1),
(default, 1,	'UBS BRASILANDIA','',	'RUA PORTO VELHO FRANCA', '1871', 'JARDIM BRASILANDIA', 'FRANCA',	'-20.5313664','-47.3656421', 1),
(default, 1,	'UBS ESTACAO','',	'AV. SANTOS DUMONT', '288', 'BAIRRO DA ESTAÇAO', 'FRANCA',	'-20.5365864',	'-47.4140001', 1),
(default, 1,	'UBS GUANABARA CS II', '',	'RUA JAMIL ABDALLA', '351', 'JARDIM GUANABARA', 'FRANCA',	'-20.5196240',	'-47.4094720',1),
(default, 1,	'UBS PAULISTA','',	'RUA LUIZ GAMA', '1880',  'JARDIM PAULISTA',  'FRANCA',	'-20.5206273',	'-47.3740249', 1),
(default, 1,	'UBS PAULISTANO','',	'RUA GILBERTO DE AGUILAR', '790',  'JARDIM PAULISTANO', 'FRANCA',	'-20.5227166',	'-47.3588368',1),
(default, 1,	'UBS PLANALTO',	'','RUA OFELIA SOARES RUSSO', '1140', 'JARDIM PLANALTO', 'FRANCA', '-20.5172411',	'-47.3863577', 1),
(default, 1,	'UBS PROGRESSO','',	'RUA HERCILIO BATISTA DE AVELAR', '1211', 'PARQUE PROGRESSO', 'FRANCA',	'-20.5611399', '-47.3976464', 1),
(default, 1,	'UBS SANTA TEREZINHA','',	'RUA FREI AGOSTINHO DA PIEDADE', '500', 'VILA SANTA TEREZINHA',  'FRANCA',	'-20.5009115',	'-47.4024206', 1),
(default, 1,	'UBS SAO SEBASTIAO','',	'RUA AMELIO BORGES CAMPOS', '603', 'VILA SAO SEBASTIAO', 'FRANCA',	'-20.5238382',	'-47.4266136', 1),
(default, 1,	'UBS VICENTE LEPORACE','',	'RUA NORBERTO BASSALO', '820', 'PARQUE VICENTE LEPORACE I', 'FRANCA',	'-20.4984030',	'-47.4162399', 1),
(default, 1,	'UPA ANITA','',	'RUA ALCINA DE LIMA SILVEIRA', '1805', 'JARDIM DAS PALMEIRAS', 'FRANCA',	'-20.5216236',	'-47.4423927', 1),
(default, 1,	'UBS DISTRITAL HORTOMIRAMONTES','',	'RUA LUIZ BELCHIOR', '1040', 'PARQUE DO HORTO', 'FRANCA',	'-20.491506',	'-47.401296', 1),
(default, 1,	'UPA AEROPORTO','',	'RUA CYRO EDUARDO ROSA FALEIROS', '222', 'PROLONGAMENTO JARDIM AEROPORTO I', 'FRANCA',	'-20.5781551',	'-47.3824339', 1),
(default, 1,	'NUCLEO DE SAUDE DA FAMILIA PALMA','', 'AV. LEILA SCARABUCCI GUIMARAES', '2842','JARDIM PALMA', 'FRANCA',	'-20.5157959',	'-47.3691389', 1),
(default, 1,	'AMB DE CRIANCA DE ALTO RISCO','',	'RUA OUVIDOR FREIRE', '2109', 'CENTRO', 'FRANCA',	'-20.5344865',	'-47.4022458', 1),
(default, 1,	'SERVICO DE ASSISTENCIA ESPECIALIZADA','',	'RUA GENERAL OSORIO', '1425', 'ESTACAO', 'FRANCA', 	'-20.5347616',	'-47.4030904',1),
(default, 1,	'NUCLEO DE SAUDE DA FAMILIA PALMA','',	'AV. LEILA SCARABUCCI GUIMARAES', '2842',  'JARDIM PALMA', 'FRANCA',	'-20.5157959',	'-47.3691389', 1),
(default, 1,	'SERVICO DE ASSISTENCIA ESPECIALIZADA','',	'RUA GENERAL OSORIO', '1425',  'ESTACAO', 'FRANCA',	'-20.5347616',	'-47.4030904', 1);

INSERT INTO Assunto (Codigo,Descricao,CodSegmento,Status)
values
(default,'CIRURGIA',1,1),
(default,'CONSULTA / ATENDIMENTO / TRATAMENTO',1,1),
(default,'DIAGNOSTICO',1,1),
(default,'MEDICAMENTO',1,1),
(default,'INSUMOS E CORRELATOS',1,1),
(default,'RECURSOS HUMANOS',1,1),
(default,'DOCUMENTOS',1 ,1),
(default,'CONSULTA / ATENDIMENTO / TRATAMENTO',1,1),
(default,'ESTABELECIMENTO DE SAÚDE',1,1),
(default,'RECURSOS MATERIAIS',1,1),
(default,'VIGILÂNCIA SANITÁRIA',1,1),
(default,'TRANSPORTE',1,1),
(default,'NÃO PERTINENTE A SAÚDE',1,1),
(default,'INFORMAÇÃO',1,1);

INSERT INTO SubAssunto (Codigo, CodigoAssunto,Descricao,Status)
values
(default,8,'DIFICULDADE DE CONTATO TELEFÔNICO',1),
(default,8,'ESPAÇO FÍSICO',1),
(default,8,'HORÁRIO DE FUNCIONAMENTO',1),
(default,8,'SUPERLOTAÇÃO NA UNIDADE',1),
(default,8,'FALTA DE VAGAS',1),
(default,8,'RECUSA NO ATENDIMENTO',1),
(default,8,'ATRASO DE PROFISSIONAL',1),
(default,8,'NÃO CUMPRIMENTO DE CARGA HORÁRIA',1),
(default,8,'ROTINAS E PROTOCOLOS',1),
(default,8,'PRESTADORES DE SERVIÇOS',1),
(default,8,'ROTINAS E PROTOCOLOS',1),
(default,8,'FALTA DE PROFISSIONAL',1);


INSERT INTO Demanda (Codigo,Descricao,Status)
values
(default,'APP',1),
(default,'SIGS',1);

INSERT INTO Classificacao (Codigo,Descricao,Status)
values
(default,'Denuncia',1),
(default,'Elogio',1),
(default,'Sugestao',1),
(default,'Reclamação',1),
(default,'Solicitação',1);

#concluído, encaminhado, fechado, em andamento	
INSERT INTO Situacao (Codigo,Descricao,Status)
values
(default,'Concluído',1),
(default,'Fechado',1),
(default,'Em andamento',1),
(default,'Encaminhado',1);