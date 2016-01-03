package montecarlo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"fmt"
)

type sequenced struct {
	Count int
	Total int
}

func (seq *sequenced) Seek(n int) error {

	if n < 0 {
		return fmt.Errorf("invalid seek position %d < 0",n)
	}

	if n >= seq.Total {
		return fmt.Errorf("invalid seek position %d >= %d",n,seq.Total)
	}

	seq.Count = n

	return nil
}

func (seq *sequenced) Read(n []byte) (int,error) {

	for i := 0; i < len(n); i++ {
		n[i] = byte(seq.Count)
		seq.Count++		
	}
	return len(n),nil
}

func (seq *sequenced) Pos() int {

	return seq.Count
}

func (seq *sequenced) Len() int {

	return seq.Total
}


func Test_Basic(t *testing.T) {

	var mon *Montecarlo

	Convey("New",t,func() {

		m,err := New(nil,&sequenced{Count:0,Total:1000})
		So(err,ShouldBeNil)
		So(m,ShouldNotBeNil)
		So(m.internal,ShouldNotBeNil)

		mon = m
	})

	Convey("Linear",t,func() {

		for i := 0; i < 10; i++ {
			r,err := mon.Linear(30,100)
			So(err,ShouldBeNil)
			So(r,ShouldBeGreaterThanOrEqualTo,30)
			So(r,ShouldBeLessThanOrEqualTo,100)
			fmt.Printf("Linear[30,100] %02d] %d\n",i,r)
		}
	})

	Convey("Linearv",t,func() {

		r,err := mon.Linearv(10,3,122)
		So(err,ShouldBeNil)
		So(len(r),ShouldEqual,10)
		for i := 0; i < len(r); i++ {
			So(r[i],ShouldBeGreaterThanOrEqualTo,3)
			So(r[i],ShouldBeLessThanOrEqualTo,122)
			fmt.Printf("Linearv[3,122] %02d] %d\n",i,r[i])
		}
	})

	Convey("RedBlack",t,func() {

		for i := 0; i < 10; i++ {
			r,err := mon.RedBlack()
			So(err,ShouldBeNil)
			fmt.Printf("RedBlack %02d] %.2f\n",i,r)
		}
	})

	Convey("WhiteBlack",t,func() {

		for i := 0; i < 10; i++ {
			r,err := mon.WhiteBlack()
			So(err,ShouldBeNil)
			fmt.Printf("WhiteBlack %02d] %.2f\n",i,r)
		}
	})

	Convey("Election",t,func() {

		for i := 0; i < 5; i++ {
			r,err := mon.Election(5)
			So(err,ShouldBeNil)
			fmt.Printf("Election %d] %v\n",i,r)
		}
	})

	Convey("Election List",t,func() {

		list := make([]int,0)
		list = append(list,4)
		list = append(list,3)
		list = append(list,1)
		list = append(list,4)
		list = append(list,2)
	
		for i := 0; i < 5; i++ {
			r,err := mon.ElectionList(list)
			So(err,ShouldBeNil)
			fmt.Printf("ElectionList %d] %v\n",i,r)
		}
	})

	Convey("Elimination",t,func() {
	
		for i := 0; i < 5; i++ {
			r,err := mon.Elimination(5,i)
			So(err,ShouldBeNil)
			So(len(r),ShouldEqual,5 - i)
			fmt.Printf("Elimination %d] %v\n",i,r)
		}
	})
}
