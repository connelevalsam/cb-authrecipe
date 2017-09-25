package grifts

import (
	"github.com/connelevalsam/authrecipe/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
