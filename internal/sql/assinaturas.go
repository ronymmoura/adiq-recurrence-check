package sql

import (
	"database/sql"
	"fmt"
	"time"
)

type Assinatura struct {
	Oid                 int             `json:"OID_ASSINAT_CRED"`
	CPF                 string          `json:"NR_CPF"`
	SqPlanoPrevidencial string          `json:"SQ_PLANO_PREVIDENCIAL"`
	DataCriacao         time.Time       `json:"DTA_CRIACAO"`
	VaultId             string          `json:"COD_VAULT"`
	IdPlano             string          `json:"COD_PLANO_ASSINAT"`
	IdAssinat           *sql.NullString `json:"COD_ID_ASSINAT"`
	Status              string          `json:"IND_STATUS"`
	NumPedido           int             `json:"NUM_PEDIDO"`
	Pagamentos          []Pagamento
}

type Pagamento struct {
	Oid                 int       `json:"OID_PAG_CRED"`
	DataPagamento       time.Time `json:"DTA_PAGAMENTO"`
	Valor               float32   `json:"VAL_PAG"`
	IdPagamento         string    `json:"COD_ID_PAGAMENTO"`
	CodigoAutorizacao   string    `json:"COD_AUTORIZACAO"`
	Infos               string    `json:"TXT_INFOS"`
	Lancado             string    `json:"IND_LANCADO"`
	OidAssinatCred      int       `json:"OID_ASSINAT_CRED"`
	SqPlanoPrevidencial string    `json:"SQ_PLANO_PREVIDENCIAL"`
}

func (conn *DbConn) GetAssinaturas() ([]Assinatura, error) {
	assinatura := Assinatura{}
	assinaturas := []Assinatura{}

	rows, err := conn.Query("SELECT * FROM WEB_ASSINAT_CRED ORDER BY NR_CPF, COD_PLANO_ASSINAT, COD_ID_ASSINAT")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&assinatura.Oid,
			&assinatura.CPF,
			&assinatura.SqPlanoPrevidencial,
			&assinatura.DataCriacao,
			&assinatura.VaultId,
			&assinatura.IdPlano,
			&assinatura.IdAssinat,
			&assinatura.Status,
			&assinatura.NumPedido,
		)

		if err != nil {
			return nil, err
		}

		pagamento := Pagamento{}
		assinatura.Pagamentos = []Pagamento{}

		tsql := fmt.Sprintf("SELECT * FROM WEB_PAG_CRED WHERE OID_ASSINAT_CRED=@p1")
		rowsPag, err := conn.Query(tsql, assinatura.Oid)

		if err != nil {
			return nil, err
		}

		defer rowsPag.Close()

		for rowsPag.Next() {
			err := rowsPag.Scan(
				&pagamento.Oid,
				&pagamento.SqPlanoPrevidencial,
				&pagamento.DataPagamento,
				&pagamento.Valor,
				&pagamento.IdPagamento,
				&pagamento.CodigoAutorizacao,
				&pagamento.Infos,
				&pagamento.OidAssinatCred,
				&pagamento.Lancado,
			)

			if err != nil {
				return nil, err
			}

			assinatura.Pagamentos = append(assinatura.Pagamentos, pagamento)
		}

		assinaturas = append(assinaturas, assinatura)
	}

	return assinaturas, nil
}
