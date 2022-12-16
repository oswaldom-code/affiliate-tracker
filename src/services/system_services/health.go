package system_services

import (
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/repository"
	"github.com/oswaldom-code/affiliate-tracker/src/services/ports"
)

type Health interface {
	TestDb() error
}
type healthImp struct {
	r ports.Repository
}

func HealthService() Health {
	return &healthImp{r: repository.NewRepository()}
}

func (p *healthImp) TestDb() error {
	return p.r.TestDb()
}
