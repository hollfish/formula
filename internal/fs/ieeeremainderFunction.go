package fs

import "github.com/yidane/formula/opt"

type IEEERemainderFunction struct {

}

func (*IEEERemainderFunction) Name() string {
	panic("implement me")
}

func (*IEEERemainderFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewIEEERemainderFunction() *IEEERemainderFunction{
	return &IEEERemainderFunction{}
}

