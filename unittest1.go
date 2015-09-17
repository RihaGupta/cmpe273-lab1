package main
import "testing"

type test1 struct{
	recs Rectangle
	result float64
}
type test2 struct{
	cirs Circle
	result float64
}

var rec = []test1{{Rectangle{2,2},8},{Rectangle{4,4},16}}
var cir = []test2{{Circle{4},25.132741228718345}}

func TestPerimeter(t *testing.T){
	for _,cases:= range rec{
		v1:=cases.recs.Peri()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
	for _,cases:= range cir{
		v1:=cases.cirs.Peri()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
}
