package main
import "testing"

type rectest struct{
	recs Rectangle
	result float64
}
type cirtest struct{
	cirs Circle
	result float64
}

var rectests = []rectest{{Rectangle{0,0,2,2},8},{Rectangle{0,0,4,4},16}}
var cirtests = []cirtest{{Circle{0,0,4},25.132741228718345}}

func TestPerimeter(t *testing.T){
	for _,cases:= range rectests{
		v1:=cases.recs.Peri()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
	for _,cases:= range cirtests{
		v1:=cases.cirs.Peri()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
}
