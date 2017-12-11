package party

import (
	"database/sql"
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

func findAll(e repository.Execer) ([]Party, error) {
	query := "SELECT id, name, address, phone, mobile, email FROM manpower.party ORDER BY id"

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
