package job

import (
	"database/sql"
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

const (
	schema     = "manpower"
	jobTable   = "job"
)

func findAll(e repository.Execer, id string, page, limit int64) ([]Job, error) {
	query := fmt.Sprintf(`SELECT a.job_id, a.type_id, a.title, a.industry, a.location,
				a.date_created, a.date_updated, a.available FROM %s.%s a WHERE
				a.job_id::TEXT LIKE '%s' ORDER BY id DESC OFFSET %d LIMIT %d`,
				schema, jobTable, id+"%", page*limit, limit)
	println()
	println(query)
	println()
	res, err := e.Query(query, jobScanner)
	if err != nil {
		return nil, fmt.Errorf("findAll: error querying database: %s", err)
	}
	jobs := res.([]Job)
	return jobs, nil
}

func jobScanner(rows *sql.Rows) (interface{}, error) {
	var results []Job
	defer rows.Close()

	for rows.Next() {
		var result Job

		err := rows.Scan(
			&result.JobID,
			&result.TypeID,
			&result.Title,
			&result.Industry,
			&result.Location,
			&result.DateCreated,
			&result.DateUpdated,
			&result.Available,
		)
		if err != nil {
			return nil, fmt.Errorf("partyScanner: error scanning row: %s", err)
		}
		results = append(results, result)
	}

	return results, nil
}
