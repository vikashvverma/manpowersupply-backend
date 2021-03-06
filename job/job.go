package job

import (
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

const unknownID = -1

type Fetcher interface {
	Fetch(string, int64, int64, string) ([]Job, error)
	SaveAndUpdate(*Job) (int64, error)
	Delete(int64) error
	Industry() ([]Industry, error)
	JobType(string) ([]Type, error)
}

type job struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &job{Execer: e}
}

func (j *job) Fetch(jobID string, page, limit int64, jobType string) ([]Job, error) {
	return findAll(j.Execer, jobID, page, limit, jobType)
}

func (j *job) SaveAndUpdate(job *Job) (int64, error) {
	rowsAffected, err := upsert(j.Execer, job)
	if err != nil {
		return unknownID, fmt.Errorf("saveAndUpdate: could save/update job: %s", err)
	}
	return rowsAffected, err
}

func (j *job) Delete(jobID int64) error {
	return delete(j.Execer, jobID)
}

func (j *job) Industry() ([]Industry, error) {
	return industry(j.Execer)
}

func (j *job) JobType(typeID string) ([]Type, error) {
	return jobTypes(j.Execer, typeID)
}
