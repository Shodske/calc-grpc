package server

import (
	"reflect"
	"testing"

	"github.com/Shodske/calc-grpc/pkg/calculator"
	"golang.org/x/net/context"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want *server
	}{
		{name: "base-case", want: &server{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Add(t *testing.T) {
	type args struct {
		ctx    context.Context
		values *calculator.Values
	}
	tests := []struct {
		name    string
		s       *server
		args    args
		want    *calculator.Result
		wantErr bool
	}{
		{name: "base-case", s: &server{}, args: args{ctx: nil, values: &calculator.Values{X: 4.25, Y: 2.5}}, want: &calculator.Result{Value: 6.75}, wantErr: false},
		{name: "nil-values", s: &server{}, args: args{ctx: nil, values: nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.Add(tt.args.ctx, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Sum(t *testing.T) {
	type args struct {
		ctx context.Context
		col *calculator.Collection
	}
	tests := []struct {
		name    string
		s       *server
		args    args
		want    *calculator.Result
		wantErr bool
	}{
		{name: "base-case", s: &server{}, args: args{ctx: nil, col: &calculator.Collection{Values: []float64{2.0, -3.125, 4.5}}}, want: &calculator.Result{Value: 3.375}, wantErr: false},
		{name: "empty-collection", s: &server{}, args: args{ctx: nil, col: &calculator.Collection{Values: []float64{}}}, want: &calculator.Result{Value: 0.0}, wantErr: false},
		{name: "one-value", s: &server{}, args: args{ctx: nil, col: &calculator.Collection{Values: []float64{-4.25}}}, want: &calculator.Result{Value: -4.25}, wantErr: false},
		{name: "nil-collection", s: &server{}, args: args{ctx: nil, col: nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.Sum(tt.args.ctx, tt.args.col)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Evaluate(t *testing.T) {
	equalT := &calculator.Comparison{
		Values:   &calculator.Values{X: 12.0, Y: 12.0},
		Operator: calculator.Comparison_EQUAL,
	}
	equalF := &calculator.Comparison{
		Values:   &calculator.Values{X: 7.4, Y: -22.2},
		Operator: calculator.Comparison_EQUAL,
	}
	notEqualT := &calculator.Comparison{
		Values:   &calculator.Values{X: 13.37, Y: 90.01},
		Operator: calculator.Comparison_NOT_EQUAL,
	}
	notEqualF := &calculator.Comparison{
		Values:   &calculator.Values{X: 4.4, Y: 4.4},
		Operator: calculator.Comparison_NOT_EQUAL,
	}
	greaterT := &calculator.Comparison{
		Values:   &calculator.Values{X: 100000000.0, Y: 99999999.9},
		Operator: calculator.Comparison_GREATER,
	}
	greaterF := &calculator.Comparison{
		Values:   &calculator.Values{X: -51.2, Y: 123.321},
		Operator: calculator.Comparison_GREATER,
	}
	lessT := &calculator.Comparison{
		Values:   &calculator.Values{X: 34.5, Y: 987.6},
		Operator: calculator.Comparison_LESS,
	}
	lessF := &calculator.Comparison{
		Values:   &calculator.Values{X: 123.456, Y: 1.7},
		Operator: calculator.Comparison_LESS,
	}

	trueResult := &calculator.BooleanResult{Value: true}
	falseResult := &calculator.BooleanResult{Value: false}

	type args struct {
		ctx context.Context
		cmp *calculator.Comparison
	}
	tests := []struct {
		name    string
		s       *server
		args    args
		want    *calculator.BooleanResult
		wantErr bool
	}{
		{name: "equal-true", s: &server{}, args: args{ctx: nil, cmp: equalT}, want: trueResult, wantErr: false},
		{name: "equal-false", s: &server{}, args: args{ctx: nil, cmp: equalF}, want: falseResult, wantErr: false},
		{name: "not-equal-true", s: &server{}, args: args{ctx: nil, cmp: notEqualT}, want: trueResult, wantErr: false},
		{name: "not-equal-false", s: &server{}, args: args{ctx: nil, cmp: notEqualF}, want: falseResult, wantErr: false},
		{name: "greater-true", s: &server{}, args: args{ctx: nil, cmp: greaterT}, want: trueResult, wantErr: false},
		{name: "greater-false", s: &server{}, args: args{ctx: nil, cmp: greaterF}, want: falseResult, wantErr: false},
		{name: "less-true", s: &server{}, args: args{ctx: nil, cmp: lessT}, want: trueResult, wantErr: false},
		{name: "less-false", s: &server{}, args: args{ctx: nil, cmp: lessF}, want: falseResult, wantErr: false},
		{name: "invalid-operator", s: &server{}, args: args{ctx: nil, cmp: &calculator.Comparison{Operator: 99999}}, want: nil, wantErr: true},
		{name: "nil-comparison", s: &server{}, args: args{ctx: nil, cmp: nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.Evaluate(tt.args.ctx, tt.args.cmp)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
