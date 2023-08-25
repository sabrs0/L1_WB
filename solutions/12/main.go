package main

import "fmt"

var (
	str = []string{"cat", "cat", "dog", "cat", "tree"}
)

type Set struct {
	elems map[string]int
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func NewSet(elems []string) Set {
	set := Set{
		elems: make(map[string]int),
	}
	for _, v := range elems {
		set.elems[v]++
	}
	return set
}
func (s Set) InterSection(set2 Set) []string {

	ans := []string{}
	for key := range set2.elems {
		if _, isOK := s.elems[key]; isOK {

			for i := 0; i < minInt(s.elems[key], set2.elems[key]); i++ {
				ans = append(ans, key)
			}
		}
	}
	return ans
}
func (s Set) Union(set2 Set) []string {

	ans := []string{}
	for key, val := range set2.elems {
		for i := 0; i < val; i++ {
			ans = append(ans, key)
		}
	}
	for key, val := range s.elems {
		for i := 0; i < val; i++ {
			ans = append(ans, key)
		}
	}
	return ans
}
func main() {
	s1 := NewSet(str)
	s2 := NewSet(str)
	fmt.Println(s1.InterSection(s2))
	fmt.Println(s1.Union(s2))

}
