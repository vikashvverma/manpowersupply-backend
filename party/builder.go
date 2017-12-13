package party

import (
	"database/sql"
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

const(
	schema = "manpower"
	table = "party"
	partyPerPage = 10
)

func save(e repository.Execer ,p *Party) (int64, error){
	query:=fmt.Sprintf("INSERT INTO %s.%s(name, address, phone, mobile, email) VALUES($1, $2, $3, $4, $5)", schema, table)
	return e.Exec(query, p.Name, p.Address, p.Phone, p.Mobile, p.Email)
}

func findAll(e repository.Execer, page int64) ([]Party, error) {
	query := fmt.Sprintf("SELECT id, name, address, phone, mobile, email FROM %s.%s ORDER BY id DESC OFFSET %d LIMIT %d", schema, table, page * partyPerPage, partyPerPage)

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
		)
		if err != nil {
			return nil, fmt.Errorf("partyScanner: error scanning row: %s", err)
		}
		results = append(results, result)
	}

	return results, nil
}
