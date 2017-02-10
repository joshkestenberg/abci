package dummy

import (
	"github.com/tendermint/abci/types"
	"github.com/tendermint/merkleeyes/app"
)

type DummyApplication struct {
	// just a simple wrapper around merkleeyes
	types.Application
}

func (a *DummyApplication) assertApplication() types.Application {
	return a
}

func NewDummyApplication() *DummyApplication {
	return &DummyApplication{
		Application: app.NewMerkleEyesApp("", 100),
	}
}
