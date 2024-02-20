package models

type ResponseBodyOtherRecords struct {
	Identification id     `json:"identificacao"`
	Name           string `json:"nome,omitempty"`
	MP             []mp   `json:"mp,omitempty"`
	BNMP           []bnmp `json:"bnmp"`
}
type mp struct {
	Sigill        string      `json:"sigla"`
	Confidence    string      `json:"confianca_associacao"`
	Subject       string      `json:"assunto"`
	MPUnity       string      `json:"unidade_mp"`
	NumberMP      string      `json:"numero_mp"`
	LawsuitNumber string      `json:"numero_processo"`
	ProcedureType string      `json:"tipo_procedimento"`
	Situation     string      `json:"situacao"`
	YearLawsuit   string      `json:"ano_do_inquerito"`
	UF            string      `json:"uf"`
	Autos         string      `json:"autos"`
	Laws          []law       `json:"tipo_de_ocorrencia"`
	CrimeFound    bool        `json:"tipificacao_identificada"`
	Movements     []movements `json:"movimentacoes"`
}

type movements struct {
	URL      string `json:"url"`
	Date     string `json:"data"`
	Detail   string `json:"detalhe"`
	Movement string `json:"movimento"`
}

type bnmp struct {
	Sigill             string `json:"sigla"`
	Confidence         string `json:"confianca_associacao"`
	UF                 string `json:"uf"`
	Situation          string `json:"situacao"`
	Organ              string `json:"orgao"`
	PrisonType         string `json:"especie_prisao"`
	DocumentNumber     string `json:"numero_documento"`
	DocumentExpedition string `json:"data_expedicao"`
	OtherName          string `json:"alcunha_acusado"`
	MotherName         string `json:"filiacao_materna_acusado"`
	FatherName         string `json:"filiacao_paterna_acusado"`
	BirthDate          string `json:"data_nascimento_acusado"`
	Nationality        string `json:"nacionalidade_acusado"`
	PlaceOfBirth       string `json:"naturalidade_acusado"`
	Profession         string `json:"profissao_acusado"`
	LawsuitLocation    string `json:"processo_local"`
	Judge              string `json:"magistrado"`
	ValidityDate       string `json:"data_validade"`
	CreationDate       string `json:"data_criacao"`
	Recapture          bool   `json:"recaptura"`
	Decision           string `json:"sintese_decisao"`
	Execution          string `json:"cumprimento"`
	Observation        string `json:"observacao"`
	PlaceOfCrime       string `json:"local_ocorrencia"`
	PenaltyTime        string `json:"tempo_pena"`
	Regime             string `json:"regime_prisional"`
	CrimeFound         bool   `json:"tipificacao_identificada"`
	Laws               []law  `json:"tipo_de_ocorrencia"`
}
