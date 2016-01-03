package montecarlo

import (
	"sync"
	"fmt"
)

var (
	ErrCatalogueExhasted error = fmt.Errorf("catalogue exhasted")
	ErrInvalidCandiates error = fmt.Errorf("invalid candiate count")
	ErrInvalidCatalogue error = fmt.Errorf("invalid catalogue")
	ErrMinGreaterThanMax error = fmt.Errorf("min >= max")
	ErrMinLessThanZero error = fmt.Errorf("min < 0")
)

type Cataloguer interface {
	Seek(int) error
	Read([]byte) (int,error)
	Pos() int
	Len() int
}

const DefaultRollForwards int = 10000

type data struct {
	sync.Mutex
	z,w,counter,position int
	catalogue Cataloguer
}

func (d *data) next(roll bool) error {

	in := make([]byte,4)
	n,err := d.catalogue.Read(in)
	if err != nil {
		return err
	}
	if n != 4 {
		return ErrCatalogueExhasted
	}
	
	y := int(in[0])
	if y == 0 {
		y = 1
	}
	z := int(in[1])
	if z == 0 {
		z = 1
	}

	d.z = y * z

	x := int(in[2])
	if x == 0 {
		x = 1
	}

	w := int(in[3])
	if w == 0 {
		w = 1
	}

	d.w = x * w

	d.position += 4

	if d.position >= d.catalogue.Len() {
	
		if roll {
			d.catalogue.Seek(0)
		} else {
			return ErrCatalogueExhasted
		}
	}
	d.counter = 0	
	return nil
}

func (d *data) read(list []int) int {

	for i := 0; i < len(list); i++ {
		d.z = 36969 * (d.z & 65535) + (d.z >> 16)
		d.w = 18000 * (d.w & 65535) + (d.w >> 16)
		list[i] = (d.z << 16) + d.w
	}
	return len(list)
}

type Configuration struct {

	Roll bool /* roll catalogue around when hit the end ? */
	RollForwards int
}

type Montecarlo struct {
	internal *data
	
	Config *Configuration
}

func (m *Montecarlo) Linear(min,max int) (int,error) {
		
	if min >= max {
		return 0,ErrMinGreaterThanMax
	}

	if min < 0 {
		return 0,ErrMinLessThanZero
	}


	r := make([]int,100) /* FIXME */
	
	m.internal.Lock()
	defer m.internal.Unlock()

	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards {
		if err := m.internal.next(m.Config.Roll); err != nil {
			return 0,err
		}
	}

	return Linear(r,min,max),nil
}

func (m *Montecarlo) RedBlack() (float64,error) {
	
	r := make([]int,100) /* FIXME, add variable count */
	
	m.internal.Lock()
	defer m.internal.Unlock()
	
	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards { 
		if err := m.internal.next(m.Config.Roll); err != nil {
			return 0.0,err
		}
	}

	return RedBlack(r,len(r),len(r) / 2),nil
}

func (m *Montecarlo) WhiteBlack() (float64,error) {

	r := make([]int,100) /* FIXME */

	m.internal.Lock()
	defer m.internal.Unlock()

	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards {
			if err := m.internal.next(m.Config.Roll); err != nil {
				return 0.0,err
			}
	}

	return WhiteBlack(r,len(r),len(r) / 2),nil
}


func (m *Montecarlo) Election(candiates int) ([]int,error) {

	if candiates <= 0 {
		return nil,ErrInvalidCandiates
	}
	
	m.internal.Lock()
	defer m.internal.Unlock()

	r := make([]int,(candiates * 100)) /* FIXME, add variable count */
	
	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards {
		if err := m.internal.next(m.Config.Roll); err != nil {
			return nil,err
		}
	}

	return Election(candiates,r),nil
}

func (m *Montecarlo) ElectionList(list []int) ([]int,error) {

	if len(list) == 0 {
		return nil,ErrInvalidCandiates
	}

	m.internal.Lock()
	defer m.internal.Unlock()

	r := make([]int,(len(list) * 100))

	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards {
		if err := m.internal.next(m.Config.Roll); err != nil {
			return nil,err
		}
	}

	results := Election(len(list),r)
	out := make([]int,0)

	for i := 0; i < len(results); i++ {		
			out = append(out,list[results[i] - 1])
	}

	return out,nil
}
			

func (m *Montecarlo) Elimination(candiates,rounds int)([]int,error) {

	if candiates <= 0 {
		return nil,ErrInvalidCandiates
	}

	m.internal.Lock()
	defer m.internal.Unlock()

	r := make([]int,(candiates * 100)) /* FIXME */
	
	n := m.internal.read(r)
	m.internal.counter += n

	if m.internal.counter >= m.Config.RollForwards {
		if err := m.internal.next(m.Config.Roll); err != nil {
			return nil,err
		}
	}

	return Elimination(candiates,rounds,r),nil
}

func New(config *Configuration,cat Cataloguer) (*Montecarlo,error) {

	if cat == nil {
		return nil,ErrInvalidCatalogue
	}

	if config == nil {
		config = &Configuration{Roll:false,RollForwards:DefaultRollForwards}
	}

	d := &data{catalogue:cat}	
	if err := d.next(false); err != nil {
		return nil,err
	}

	m := &Montecarlo{}
	m.internal = d

	m.Config = config
	
	return m,nil
}

	


