package comet

import (
	"errors"
	"maoim/api/protocal"
)

type Ring struct {
	rp uint64
	wp uint64
	num uint64
	mask uint64
	data []protocal.Proto
}

func New(num uint64) *Ring {
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num-1
		}
		num <<= 1
	}

	r := &Ring{
		data: make([]protocal.Proto, num),
		num: num,
		mask: num - 1,
	}
	return r
}

func (r *Ring) Get() (*protocal.Proto, error) {
	if r.rp == r.wp {
		return nil, errors.New("ring is empty")
	}
	return &r.data[r.rp&r.mask], nil
}

func (r *Ring) GetIncr()  {
	r.rp++
}

func (r *Ring) Set() (*protocal.Proto, error) {
	if r.wp - r.rp >= r.num {
		return nil, errors.New("ring is full")
	}
	return &r.data[r.wp&r.mask], nil
}

func (r *Ring) SetIncr() {
	r.wp++
}

func (r *Ring) Reset() {
	r.wp = 0
	r.rp = 0
}