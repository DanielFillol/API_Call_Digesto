package models

type ResponseBody struct {
	Identification id        `json:"identificacao"`
	Name           string    `json:"nome,omitempty"`
	Pagination     page      `json:"pagination"`
	Lawsuits       []Lawsuit `json:"processos"`
}

type id struct {
	IdType string `json:"tipo,omitempty"`
	Value  string `json:"valor,omitempty"`
}

type page struct {
	EndCursor   string `json:"endCursor,omitempty"`
	HasNextPage bool   `json:"hasNextPage,omitempty"`
}

type Lawsuit struct {
	UF               string           `json:"UF,omitempty"`
	DistYear         string           `json:"ano_distribuicao,omitempty"`
	LawsuitYear      string           `json:"ano_do_processo,omitempty"`
	Subject          string           `json:"assunto,omitempty"`
	CourtCity        string           `json:"comarca,omitempty"`
	Confidence       string           `json:"confianca_associacao,omitempty"`
	MostRecentMove   string           `json:"data_andamento_mais_recente,omitempty"`
	MostRecentUpdate string           `json:"data_ultima_atualizacao,omitempty"`
	Forum            string           `json:"forum,omitempty"`
	Link             string           `json:"link,omitempty"`
	Nature           string           `json:"natureza,omitempty"`
	CoverName        string           `json:"nome_na_capa,omitempty"`
	LawsuitNumber    string           `json:"numero_processo,omitempty"`
	Role             string           `json:"papel,omitempty"`
	PassivePole      bool             `json:"polo_passivo,omitempty"`
	MainLawsuit      bool             `json:"processo_principal,omitempty"`
	RelatedLawsuits  []relatedLawsuit `json:"processos_relacionados,omitempty"`
	Laws             []law            `json:"tipificacao"`
	CrimeFound       bool             `json:"tipificacao_identificada,omitempty"`
	JusticeType      string           `json:"tipo_processo,omitempty"`
	Court            string           `json:"tribunal,omitempty"`
}

type law struct {
	Law   []string `json:"lei,omitempty"`
	Crime string   `json:"tipo_de_ocorrencia,omitempty"`
}

type relatedLawsuit struct {
	LawSuitNumber string `json:"numero_processo,omitempty"`
	Link          string `json:"link,omitempty"`
}
