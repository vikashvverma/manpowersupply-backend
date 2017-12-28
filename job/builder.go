package job

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vikashvverma/manpowersupply-backend/repository"
	"strings"
)

const (
	schema        = "manpower"
	jobTable      = "job"
	industryTable = "industry"
	jobType       = "job_type"
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
	rowsAffected, err := e.Exec(query, jobID)
	if err != nil {
		return fmt.Errorf("delete: could not delete job having id: %d, err: %s", jobID, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("delete: no job found with jobID: %d", jobID)
	}

	return nil
}
func findAll(e repository.Execer, id string, page, limit int64, jobType string) ([]Job, error) {
	query := fmt.Sprintf(`SELECT a.job_id, a.type_id, a.title, a.industry, a.location,
				a.date_created, a.date_updated, a.available FROM %s.%s a WHERE
				a.job_id::TEXT LIKE '%s' AND LOWER(a.industry) LIKE '%s' ORDER BY id DESC OFFSET %d LIMIT %d`,
		schema, jobTable, id+"%", strings.ToLower(jobType)+"%", page*limit, limit)
	res, err := e.Query(query, jobScanner)
	if err != nil {
		return nil, fmt.Errorf("findAll: error querying database: %s", err)
	}
	jobs := res.([]Job)
	return jobs, nil
}

func jobTypes(e repository.Execer, typeID string) ([]Type, error) {
	query := fmt.Sprintf("SELECT type_id, title FROM %s.%s WHERE type_id::TEXT  LIKE '%s'", schema, jobType, string(typeID)+"%")
	fmt.Println()
	fmt.Println(typeID)
	fmt.Println(query)
	fmt.Println()
	types, err := e.Query(query, jobTypeScanner)
	if err != nil {
		return nil, fmt.Errorf("jobTypes: could not query job types: %s", err)
	}
	return types.([]Type), nil
	return nil, nil
}

func industry(e repository.Execer) ([]Industry, error) {
	query := fmt.Sprintf("SELECT type_id, industry FROM %s.%s", schema, industryTable)

	types, err := e.Query(query, industryScanner)
	if err != nil {
		return nil, fmt.Errorf("industry: could not query job types: %s", err)
	}
	return types.([]Industry), nil
}

func industryScanner(rows *sql.Rows) (interface{}, error) {
	var results []Industry
	defer rows.Close()

	for rows.Next() {
		var result Industry

		err := rows.Scan(&result.TypeID, &result.Industry)
		if err != nil {
			return nil, fmt.Errorf("industryScanner: error scanning row: %s", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func jobTypeScanner(rows *sql.Rows) (interface{}, error) {
	var results []Type
	defer rows.Close()

	for rows.Next() {
		var result Type

		err := rows.Scan(&result.TypeID, &result.Title)
		if err != nil {
			return nil, fmt.Errorf("industryScanner: error scanning row: %s", err)
		}
		results = append(results, result)
	}
	return results, nil
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
