package system_services

import (
	"github.com/oswaldom-code/api-template-gin/src/adapters/repository"
	"github.com/oswaldom-code/api-template-gin/src/aplication/system_services/ports"
)

type Health interface {
	TestDb() error
}
type healthImp struct {
	r ports.Store
}

func HealthService() Health {
	return &healthImp{r: repository.NewRepository()}
}

func (p *healthImp) TestDb() error {
	return p.r.TestDb()
}
