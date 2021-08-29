// +build wireinject
package di

import "github.com/google/wire"

func BuildInjector() (*Event, error) {
	panic(wire.Build(EventSet))
}
