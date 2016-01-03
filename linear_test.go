package montecarlo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func Test_Linear(t *testing.T) {

	var inputs = []int{10,123,2455,36452,43,5562,624521,723,831,9,102132}

	Convey("Linear [0,100]",t,func() {
	
		for i := 0; i < 10; i++ {
			r := Linear(inputs,0,100)
			fmt.Printf("Linear [0,100] %d] %d\n",i,r)
			So(r,ShouldEqual,23)
		}
	})
	
}
