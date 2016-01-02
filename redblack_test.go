package montecarlo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_RedBlack(t *testing.T) {

	var inputs = []int{10,123,2455,36452,43,5562,624521,723,831,9,102132}

	Convey("RedBlack, balanced",t,func() {
		
		r := RedBlack(inputs,100,50)
		So(r,ShouldEqual,1.0)

		r1 := RedBlack(inputs,100,50)
		So(r1,ShouldEqual,r)
	})

	Convey("RedBlack, unbalanced to black",t,func() {

		r := RedBlack(inputs,100,80)
		So(r,ShouldEqual,1.0)

		bs,ws,rs := RedWhiteBlackBox(inputs,100,80)
		So(bs + ws + rs,ShouldEqual,len(inputs))
		So(bs,ShouldEqual,11)
		So(ws,ShouldEqual,0)
		So(rs,ShouldEqual,0)
	})

	Convey("RedBlack, unbalanced to red",t,func() {

		r := RedBlack(inputs,100,10)
		So(r,ShouldEqual,0.0)

		bs,ws,rs := RedWhiteBlackBox(inputs,100,10)
		So(bs + ws + rs,ShouldEqual,len(inputs))
		So(bs,ShouldEqual,1)
		So(ws,ShouldEqual,1)
		So(rs,ShouldEqual,9)
	})

}

func Test_WhiteBlack(t *testing.T) {

	var inputs = []int{12345,6789,101234,234567,202492,29484,19389389,193848,1932934,1938458,596904}

	Convey("WhiteBlack, balanced",t,func() {

		r := WhiteBlack(inputs,100,50)
		So(r,ShouldEqual,0.0)

		r1 := WhiteBlack(inputs,100,50)
		So(r1,ShouldEqual,r)
	})


}
		

