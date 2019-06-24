package main

import (
	"context"
	"github.com/Shodske/calc-grpc/pkg/calculator"
	"google.golang.org/grpc"
	"math"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	// Start the server so it will listen to connection.
	go main()

	// To test the server we will use the generated client to make sure the server accepts the requests and returns
	// correct responses.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	c := calculator.NewCalculatorClient(conn)

	// Test calling the add endpoint
	type addArgs struct {
		ctx    context.Context
		values *calculator.Values
	}
	addTests := []struct {
		name    string
		args    addArgs
		want    *calculator.Result
		wantErr bool
	}{
		{name: "base-case", args: addArgs{ctx: context.Background(), values: &calculator.Values{X: 35.0, Y: 12.5}}, want: &calculator.Result{Value: 47.5}, wantErr: false},
		{name: "large-values", args: addArgs{ctx: context.Background(), values: &calculator.Values{X: math.MaxFloat64, Y: math.MaxFloat64}}, want: &calculator.Result{Value: math.Inf(1)}, wantErr: false},
		{name: "negative-values", args: addArgs{ctx: context.Background(), values: &calculator.Values{X: -1242.25, Y: -666}}, want: &calculator.Result{Value: -1908.25}, wantErr: false},
		{name: "large-negative-values", args: addArgs{ctx: context.Background(), values: &calculator.Values{X: -math.MaxFloat64, Y: -math.MaxFloat64}}, want: &calculator.Result{Value: math.Inf(-1)}, wantErr: false},
	}
	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Add(tt.args.ctx, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test calling the sum endpoint

	// This is a large collection that should sum up to 199990000
	largeCollection := &calculator.Collection{Values: make([]float64, 20000)}
	for i := 0; i < 20000; i++ {
		largeCollection.Values[i] = float64(i)
	}

	type sumArgs struct {
		ctx context.Context
		col *calculator.Collection
	}
	sumTests := []struct {
		name    string
		args    sumArgs
		want    *calculator.Result
		wantErr bool
	}{
		{
			name:    "base-case",
			args:    sumArgs{ctx: context.Background(), col: &calculator.Collection{[]float64{1.0, 2.0, 3.0}}},
			want:    &calculator.Result{Value: 6.0},
			wantErr: false,
		},
		{
			name:    "empty-collection",
			args:    sumArgs{ctx: context.Background(), col: &calculator.Collection{[]float64{}}},
			want:    &calculator.Result{Value: 0.0},
			wantErr: false,
		},
		{
			name:    "large-collection",
			args:    sumArgs{ctx: context.Background(), col: largeCollection},
			want:    &calculator.Result{Value: 199990000},
			wantErr: false,
		},
		{
			name:    "large-collection-values",
			args:    sumArgs{ctx: context.Background(), col: &calculator.Collection{[]float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}}},
			want:    &calculator.Result{Value: math.Inf(1)},
			wantErr: false,
		},
	}
	for _, tt := range sumTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Sum(tt.args.ctx, tt.args.col)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test evaluate endpoint

	equalT := &calculator.Comparison{
		Values:   &calculator.Values{X: 12.0, Y: 12.0},
		Operator: calculator.Comparison_EQUAL,
	}
	notEqualF := &calculator.Comparison{
		Values:   &calculator.Values{X: 4.4, Y: 4.4},
		Operator: calculator.Comparison_NOT_EQUAL,
	}
	greaterF := &calculator.Comparison{
		Values:   &calculator.Values{X: -51.2, Y: 123.321},
		Operator: calculator.Comparison_GREATER,
	}
	lessT := &calculator.Comparison{
		Values:   &calculator.Values{X: 34.5, Y: 987.6},
		Operator: calculator.Comparison_LESS,
	}

	trueResult := &calculator.BooleanResult{true}
	falseResult := &calculator.BooleanResult{false}

	type evalArgs struct {
		ctx context.Context
		cmp *calculator.Comparison
	}
	tests := []struct {
		name    string
		args    evalArgs
		want    *calculator.BooleanResult
		wantErr bool
	}{
		{name: "equal-true", args: evalArgs{ctx: context.Background(), cmp: equalT}, want: trueResult, wantErr: false},
		{name: "not-equal-false", args: evalArgs{ctx: context.Background(), cmp: notEqualF}, want: falseResult, wantErr: false},
		{name: "greater-false", args: evalArgs{ctx: context.Background(), cmp: greaterF}, want: falseResult, wantErr: false},
		{name: "less-true", args: evalArgs{ctx: context.Background(), cmp: lessT}, want: trueResult, wantErr: false},
		{name: "invalid-operator", args: evalArgs{ctx: context.Background(), cmp: &calculator.Comparison{Operator: 99999}}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Evaluate(tt.args.ctx, tt.args.cmp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
