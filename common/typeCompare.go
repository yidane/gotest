package main

type Stringer interface {
	String() string
}
type String struct {
	data string
}

func (s *String) String() string {
	return s.data
}

func GetString() *String {
	return nil
}
func CheckString(s Stringer) bool {
	return s == nil
}
