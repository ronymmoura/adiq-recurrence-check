package sql

import "time"

type FichaCobranca struct {
	NoPessoa            string    `json:"no_pessoa"`
	CPF                 string    `json:"nr_cpf"`
	SqCobranca          int       `json:"sq_cobranca"`
	CdPessoa            int       `json:"cd_pessoa"`
	SqTipoCobranca      int       `json:"sq_tipo_cobranca"`
	SqLocalCobranca     int       `json:"sq_local_cobranca"`
	SqPlanoPrevidencial int       `json:"sq_plano_previdencial"`
	DtReferencia        time.Time `json:"dt_referencia"`
	DtCompetencia       time.Time `json:"dt_competencia"`
	DtRegistro          time.Time `json:"dt_registro"`
	VlCobranca          float64   `json:"vl_cobranca"`
	IrStatusCobranca    string    `json:"ir_status_cobranca"`
}

type FichaFinanceira struct {
	NoPessoa            string    `json:"no_pessoa"`
	CPF                 string    `json:"nr_cpf"`
	SqFicha             int       `json:"sq_ficha"`
	SqPlanoPrevidencial int       `json:"sq_plano_previdencial"`
	SqContratoTrabalho  int       `json:"sq_contrato_trabalho"`
	SqTipoCobranca      int       `json:"sq_tipo_cobranca"`
	DtReferencia        time.Time `json:"dt_referencia"`
	DtCompetencia       time.Time `json:"dt_competencia"`
	DtAporte            time.Time `json:"dt_aporte"`
	VlContribuicao      float64   `json:"vl_contribuicao"`
}

func (fc FichaCobranca) Exists(list []FichaCobranca) bool {
	for _, x := range list {
		if x.CPF == fc.CPF &&
			x.DtCompetencia == fc.DtCompetencia &&
			x.VlCobranca == fc.VlCobranca &&
			x.SqCobranca != fc.SqCobranca {
			return true
		}
	}

	return false
}

func (fc FichaFinanceira) Exists(list []FichaFinanceira) bool {
	for _, x := range list {
		if x.CPF == fc.CPF &&
			x.DtCompetencia == fc.DtCompetencia &&
			x.VlContribuicao == fc.VlContribuicao &&
			x.SqTipoCobranca == fc.SqTipoCobranca &&
			x.SqFicha != fc.SqFicha {
			return true
		}
	}

	return false
}

func (conn *DbConn) GetFichaCobranca() ([]FichaCobranca, error) {
	query := `
SELECT 
	pe.no_pessoa,
	pf.nr_cpf,
	fc.sq_cobranca,
	fc.cd_pessoa,
	fc.sq_tipo_cobranca,
	fc.sq_local_cobranca,
	fc.sq_plano_previdencial,
	fc.dt_referencia,
	fc.dt_competencia,
	fc.dt_registro,
	fc.vl_cobranca,
	fc.ir_status_cobranca
FROM fi_ficha_financ_cobranca fc
LEFT JOIN fi_pessoa_fisica pf on pf.cd_pessoa = fc.cd_pessoa
LEFT JOIN fi_pessoa pe on pe.cd_pessoa = pf.cd_pessoa
WHERE SQ_TIPO_COBRANCA IN (3, 1)
  AND SQ_LOCAL_COBRANCA = 6
ORDER BY nr_cpf, DT_REGISTRO DESC`

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fichaCobranca := FichaCobranca{}
	ficha := []FichaCobranca{}
	for rows.Next() {
		err := rows.Scan(
			&fichaCobranca.NoPessoa,
			&fichaCobranca.CPF,
			&fichaCobranca.SqCobranca,
			&fichaCobranca.CdPessoa,
			&fichaCobranca.SqTipoCobranca,
			&fichaCobranca.SqLocalCobranca,
			&fichaCobranca.SqPlanoPrevidencial,
			&fichaCobranca.DtReferencia,
			&fichaCobranca.DtCompetencia,
			&fichaCobranca.DtRegistro,
			&fichaCobranca.VlCobranca,
			&fichaCobranca.IrStatusCobranca,
		)

		if err != nil {
			return nil, err
		}

		ficha = append(ficha, fichaCobranca)
	}

	return ficha, nil
}

func (conn *DbConn) GetFichaFinanc() ([]FichaFinanceira, error) {
	query := `
SELECT
	pe.no_pessoa,
	pf.nr_cpf,
	cp.sq_ficha,
	cp.sq_plano_previdencial,
	cp.sq_contrato_trabalho,
	cp.sq_tipo_cobranca,
	cp.dt_referencia,
	cp.dt_competencia,
	cp.dt_aporte,
	cp.vl_contribuicao
FROM fi_ficha_contrib_previdencial cp
LEFT JOIN fi_contrato_trabalho ct on ct.sq_contrato_trabalho = cp.sq_contrato_trabalho
LEFT JOIN fi_pessoa_fisica pf on pf.cd_pessoa = ct.cd_pessoa
LEFT JOIN fi_pessoa pe on pe.cd_pessoa = pf.cd_pessoa
WHERE SQ_TIPO_COBRANCA IN (3, 1)
  AND NR_CPF IS NOT NULL
ORDER BY nr_cpf, DT_APORTE DESC`

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fichaFinanc := FichaFinanceira{}
	ficha := []FichaFinanceira{}

	for rows.Next() {
		err := rows.Scan(
			&fichaFinanc.NoPessoa,
			&fichaFinanc.CPF,
			&fichaFinanc.SqFicha,
			&fichaFinanc.SqPlanoPrevidencial,
			&fichaFinanc.SqContratoTrabalho,
			&fichaFinanc.SqTipoCobranca,
			&fichaFinanc.DtReferencia,
			&fichaFinanc.DtCompetencia,
			&fichaFinanc.DtAporte,
			&fichaFinanc.VlContribuicao,
		)
		if err != nil {
			return nil, err
		}

		ficha = append(ficha, fichaFinanc)
	}

	return ficha, nil
}
