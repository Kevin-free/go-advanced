// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	pb "biohouse/api"
	"biohouse/internal/dao"
	"biohouse/internal/server/http"
	"biohouse/internal/server/tcp"
	"biohouse/internal/service"

	"github.com/google/wire"
)

var daoProvider = wire.NewSet(dao.New)
var serviceProvider = wire.NewSet(service.New, wire.Bind(new(pb.BiohouseCometServer), new(*service.Service)))

func InitApp() (*App, func(), error) {
	panic(wire.Build(daoProvider, serviceProvider, http.New, tcp.New, NewApp))
}
