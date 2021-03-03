package collections

type Set struct {
	m map[int]bool
	name string
}

func (set Set)add(value int)  {
	set.m[value] = true
}

func (set Set)delete(value int)  {
	delete(set.m,value)
}

func (set Set)contains(value int) bool {
	return set.m[value]
}

func (set *Set)setName(name string)  {
	set.name = name
}