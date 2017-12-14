package party

import (
	"fmt"

	"github.com/vikashvverma/manpowersupply-backend/repository"
)

type Fetcher interface {
	Fetch(string, int64) ([]Party, error)
	Save(*Party) error
}

type party struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &party{Execer: e}
}

func (p *party) Fetch(id string, page int64) ([]Party, error) {
	return findAll(p.Execer, id, page)
}

func (p *party) Save(party *Party) error {
	lastInsertID, err := saveParty(p.Execer, party)
	if err != nil {
		return fmt.Errorf("saveParty: could not party: %s", err)
	}

	party.Query.QueryerID = lastInsertID
	_, err = saveQuery(p.Execer, &party.Query)
	if err != nil {
		return fmt.Errorf("saveParty: could not query: %s", err)
	}
	return nil
}
