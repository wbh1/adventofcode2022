package util

type Set []interface{}

func (s *Set) Add(value interface{}) {
	for _, val := range *s {
		if val == value {
			return
		}
	}

	*s = append(*s, value)
}
