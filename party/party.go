package party

import (
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

type Fetcher interface {
	Fetch(int64) ([]Party, error)
	Save(*Party) error
}

type party struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &party{Execer: e}
}

func (p *party) Fetch(page int64) ([]Party, error) {
	return findAll(p.Execer, page)
}

func (p *party) Save(party *Party) error {
	_, err := save(p.Execer, party)
	if err !=nil{
		return fmt.Errorf("save: could not save party: %s", err)
	}
	return nil
}
