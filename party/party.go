package party

import (
	"github.com/vikashvverma/manpowersupply-backend/repository"
)

type Fetcher interface {
	Fetch() ([]Party, error)
}

type party struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &party{Execer: e}
}

func (p *party) Fetch()([]Party, error){
	return findAll(p.Execer)
}
