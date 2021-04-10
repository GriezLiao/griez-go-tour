package convert

import "strconv"

type StrTo string

func (st StrTo) String() string {
	return string(st)
}

func (st StrTo) Int() (int, error) {
	return strconv.Atoi(st.String())
}

func (st StrTo) MustInt() int {
	i, _ := strconv.Atoi(st.String())
	return i
}

func (st StrTo) UInt32() (uint32, error) {
	i, err := strconv.Atoi(st.String())
	return uint32(i), err
}

func (st StrTo) MustUInt32() uint32 {
	i, _ := strconv.Atoi(st.String())
	return uint32(i)
}
