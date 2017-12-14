package job

import "github.com/vikashvverma/manpowersupply-backend/repository"

type Fetcher interface {
	Fetch(string, int64, int64) ([]Job, error)
	Update(Job) error
	Delete(Job) error
}

type job struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &job{Execer: e}
}

func (j *job) Fetch(jobID string, page, limit int64) ([]Job, error) {
	return findAll(j.Execer, jobID, page, limit)
}

func (j *job) Update(job Job) error {
	return nil
}

func (j *job) Delete(job Job) error {
	return nil
}
