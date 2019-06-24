package server

import (
	"errors"
	"fmt"
	"github.com/Shodske/calc-grpc/pkg/calculator"
	"golang.org/x/net/context"
)

// Server struct that implements CalculatorServer.
type server struct {
}

// Constructor to initialize a new server.
func NewServer() *server {
	return &server{}
}

// Add function required by the CalculatorServer interface, adds 2 values and returns the result.
func (*server) Add(ctx context.Context, values *calculator.Values) (*calculator.Result, error) {
	if values == nil {
		return nil, errors.New("could not add values, nil values received")
	}

	res := &calculator.Result{
		Value: values.X + values.Y,
	}

	return res, nil
}

// Sum function required by the CalculatorServer interface, sums all values in the collection and returns the result.
func (*server) Sum(ctx context.Context, col *calculator.Collection) (*calculator.Result, error) {
	if col == nil {
		return nil, errors.New("could not sum collection, nil collection received")
	}

	res := &calculator.Result{}
	for _, v := range col.Values {
		res.Value += v
	}

	return res, nil
}

// Evaluate function evaluates given comparison and returns a boolean result.
func (*server) Evaluate(ctx context.Context, cmp *calculator.Comparison) (*calculator.BooleanResult, error) {
	if cmp == nil {
		return nil, errors.New("could not evaluate comparison, nil comparison received")
	}

	res := &calculator.BooleanResult{}
	switch cmp.Operator {
	case calculator.Comparison_EQUAL:
		res.Value = cmp.Values.X == cmp.Values.Y
	case calculator.Comparison_NOT_EQUAL:
		res.Value = cmp.Values.X != cmp.Values.Y
	case calculator.Comparison_GREATER:
		res.Value = cmp.Values.X > cmp.Values.Y
	case calculator.Comparison_LESS:
		res.Value = cmp.Values.X < cmp.Values.Y
	default:
		return nil, errors.New(fmt.Sprintf("could not evaluator comparison, invalid operator %d given", cmp.Operator))
	}

	return res, nil
}
