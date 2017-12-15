package job

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

const (
	schema   = "manpower"
	jobTable = "job"
)

func upsert(e repository.Execer, j *Job) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s.%s(job_id, title, industry, location,
			date_updated, available, type_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT(job_id) DO UPDATE SET
			(title, industry, location, available, type_id) = ($2, $3, $4, $6, $7)`,
		schema, jobTable)
	return e.Exec(query, j.JobID, j.Title, j.Industry, j.Location, time.Now(), j.Available, j.TypeID)
}

func delete(e repository.Execer, jobID int64) error {
	query := fmt.Sprintf("DELETE FROM %s.%s WHERE job_id = $1", schema, jobTable)
	_, err := e.Exec(query, jobID)
	if err != nil {
		return fmt.Errorf("delete: could not delete job having id: %d, err: ", jobID, err)
	}
	return nil
}
func findAll(e repository.Execer, id string, page, limit int64) ([]Job, error) {
	query := fmt.Sprintf(`SELECT a.job_id, a.type_id, a.title, a.industry, a.location,
				a.date_created, a.date_updated, a.available FROM %s.%s a WHERE
				a.job_id::TEXT LIKE '%s' ORDER BY id DESC OFFSET %d LIMIT %d`,
		schema, jobTable, id+"%", page*limit, limit)
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
