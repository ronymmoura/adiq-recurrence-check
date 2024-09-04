package sql

import (
	"database/sql"
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
}

func (conn *DbConn) GetAssinaturas() ([]Assinatura, error) {
	assinatura := Assinatura{}
	assinaturas := []Assinatura{}

	rows, err := conn.Query("SELECT * FROM WEB_ASSINAT_CRED")
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

		assinaturas = append(assinaturas, assinatura)
	}

	return assinaturas, nil
}
