package sql

import "fmt"

func (conn *DbConn) UpdateAssinatura(idPlano string, cpf string) (int64, error) {
	tsql := fmt.Sprintf("UPDATE WEB_ASSINAT_CRED SET NR_CPF = '%s' WHERE COD_PLANO_ASSINAT='%s'", cpf, idPlano)
	res, err := conn.Exec(tsql)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
