package main

import (
	"bytes"
	"fmt"
)

// IntSet, Ch 6.5, gopl.io
type IntSet struct {
	words []uint64
}

// Has checks membership.
func (s *IntSet) Sum(x int) bool {
	return true
}

func (hh *IntSet) Sum2(u string) bool{
	return true
}


func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add sets and element.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	//log.Printf("Add(%d), %d, %d", x, word, bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Unionwith merges t into s.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the number of elements.
func (s *IntSet) Len() uint64 {
	var length uint64
	for _, word := range s.words {
		var i uint
		for i = 0; i < 64; i++ {
			if word&(1<<i) != 0 {
				length++
			}
		}
	}
	return length
}

// Remove removes an element.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements.
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	var t = IntSet{words: make([]uint64, len(s.words))}
	for i := range s.words {
		t.words[i] = s.words[i]
	}
	return &t
}

// String prints the set.
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('<')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("<") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('>')
	return buf.String()
}

func main(){
	gg:=IntSet{[]uint64{1,2,3}}
	gg.Add(4)
	gg_point:=&IntSet{}
	gg_point.Add(100)
	(*gg_point).Add(101)

	fmt.Println(gg)
	fmt.Println(gg_point) //指针
	fmt.Println(*gg_point) //取值

	type struct1 struct {
		i1  int
		f1  float32
		str string
	}

	ms := new(struct1)
	(*ms).i1 = 10
	ms.f1 = 15.5
	ms.str= "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)


var x IntSet
x.Add(111)  //ok 因为x是变量
	(&x).Add(123123)
//IntSet{[]uint64{1,2}}.Add(1111)  //error add方法接收的是指针

}
