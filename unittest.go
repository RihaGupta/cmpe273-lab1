package main
import "testing"

type testpair struct{
	value int
	result int
}

var tests =[]testpair {{3,1},{4,2},{5,3},{6,5}}

func TestFibonacci(t *testing.T){
	for _, cases := range tests{
		v := Fibonacci(cases.value)
		if v != cases.result{
			t.Error("For",cases.value,
			"expected",cases.result,
			"got",v,
			)
		}
	}
}
