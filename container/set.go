package container

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(map[interface{}]struct{})
}

func (s Set) Add(val interface{}) bool {
	_, ok := s[val]
	if ok {
		return false
	}
	s[val] = struct{}{}
	return true
}
func (s Set) AddAll(all Set) {
	for k := range all {
		s.Add(k)
	}
}

func (s Set) RemoveAll(r Set) {
	for k := range r {
		delete(s, k)
	}
}

func (s Set) Size() int {
	return len(s.List())
}
func (s Set) List() []interface{} {
	var ret []interface{}
	for k := range s {
		ret = append(ret, k)
	}
	return ret
}
