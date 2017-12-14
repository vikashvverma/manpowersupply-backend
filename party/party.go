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
	lastInsertID, err := saveParty(p.Execer, party)
	if err != nil {
		return fmt.Errorf("saveParty: could not party: %s", err)
	}
	fmt.Println("lastInsertID: ", lastInsertID)
	party.Query.QueryerID = lastInsertID
	_, err = saveQuery(p.Execer, &party.Query)
	if err != nil {
		return fmt.Errorf("saveParty: could not query: %s", err)
	}
	return nil
}
