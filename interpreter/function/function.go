package function

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

type Function struct {
	Scope            context.Scope
	Call             func(ctx *context.Context, args ...value.Value) (value.Value, error)
	CanStatementCall bool
	IsIdentArgument  func(i int) bool
}

func Exists(scope context.Scope, name string) (*Function, error) {
	fn, ok := builtinFunctions[name]
	if !ok {
		return nil, errors.WithStack(
			fmt.Errorf("Function %s is not defined", name),
		)
	} else if (fn.Scope & scope) == 0 {
		return nil, errors.WithStack(
			fmt.Errorf("Function %s could not call on %s scope", name, scope.String()),
		)
	}
	return fn, nil
}

func Inject(fns map[string]*Function) error {
	for key, fn := range fns {
		if _, ok := builtinFunctions[key]; ok {
			return errors.WithStack(
				fmt.Errorf("Function %s already defiend and could not override", key),
			)
		}
		builtinFunctions[key] = fn
	}
	return nil
}
