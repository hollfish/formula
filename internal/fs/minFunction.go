package fs

import "github.com/yidane/formula/opt"

type MinFunction struct {
}

func (*MinFunction) Name() string {
	panic("implement me")
}

func (*MinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewMinFunction() *MinFunction{
	return &MinFunction{}
}
