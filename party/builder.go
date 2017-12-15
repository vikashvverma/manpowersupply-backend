package party

import (
	"database/sql"
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

const (
	schema       = "manpower"
	partyTable   = "party"
	query        = "query"
	partyPerPage = 10
)

func saveQuery(e repository.Execer, q *Query) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s.%s(queryer_id, query) VALUES($1, $2)", schema, query)
	return e.Exec(query, q.QueryerID, q.Query)
}

func saveParty(e repository.Execer, p *Party) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s.%s(name, address, phone, mobile, email) VALUES($1, $2, $3, $4, $5) RETURNING id", schema, partyTable)
	return e.QueryRow(query, p.Name, p.Address, p.Phone, p.Mobile, p.Email)
}

func findAll(e repository.Execer, id string, page int64) ([]Party, error) {
	query := fmt.Sprintf(`SELECT a.id, a.name, a.address, a.phone, a.mobile, a.email,
				b.queryer_id, b.query, b.query_date FROM %s.%s a INNER JOIN %s.%s b
				ON a.id = b.queryer_id WHERE a.id::TEXT LIKE '%s' ORDER BY id DESC
				OFFSET %d LIMIT %d`,
		schema, partyTable, schema, query, id+"%", page*partyPerPage, partyPerPage)
	res, err := e.Query(query, partyScanner)
	if err != nil {
		return nil, fmt.Errorf("findAll: error querying database: %s", err)
	}
	parties := res.([]Party)
	return parties, nil
}

func partyScanner(rows *sql.Rows) (interface{}, error) {
	var results []Party
	defer rows.Close()

	for rows.Next() {
		var result Party

		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Address,
			&result.Phone,
			&result.Mobile,
			&result.Email,
			&result.Query.QueryerID,
			&result.Query.Query,
			&result.Query.QueryDate,
		)
		if err != nil {
			return nil, fmt.Errorf("partyScanner: error scanning row: %s", err)
		}
		results = append(results, result)
	}

	return results, nil
}
