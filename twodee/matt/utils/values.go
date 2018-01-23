package utils

import "fmt"

type valueHitMap int

const (
	emptyMapping valueHitMap = 0
	oneMapping   valueHitMap = 1
	twoMapping   valueHitMap = 1 << 1
	threeMapping valueHitMap = 1 << 2
	fourMapping  valueHitMap = 1 << 3
	fiveMapping  valueHitMap = 1 << 4
	sixMapping   valueHitMap = 1 << 5
	sevenMapping valueHitMap = 1 << 6
	eightMapping valueHitMap = 1 << 7
	nineMapping  valueHitMap = 1 << 8

	fullMapping = oneMapping | twoMapping | threeMapping | fourMapping | fiveMapping | sixMapping | sevenMapping | eightMapping | nineMapping
)

func valueToMapping(value int) valueHitMap {
	if value < 1 || value > 9 {
		return 0
	}
	return 1 << uint(value-1)
}

func (v valueHitMap) valueIsSet(value int) bool {
	return (v & valueToMapping(value)) > 0
}

func (v valueHitMap) setValue(value int) valueHitMap {
	var orValue valueHitMap
	switch value {
	case 1:
		orValue = oneMapping
	case 2:
		orValue = twoMapping
	case 3:
		orValue = threeMapping
	case 4:
		orValue = fourMapping
	case 5:
		orValue = fiveMapping
	case 6:
		orValue = sixMapping
	case 7:
		orValue = sevenMapping
	case 8:
		orValue = eightMapping
	case 9:
		orValue = nineMapping
	}
	return v | orValue
}

func (v valueHitMap) numberSet() int {
	num := 0
	for i := 1; v > 0; i++ {
		if v&oneMapping > 0 {
			num++
		}
		v = v >> 1
	}
	return num
}

func (v valueHitMap) getValues() []int {
	values := []int{}
	for i := 1; v > 0; i++ {
		if v&oneMapping > 0 {
			values = append(values, i)
		}
		v = v >> 1
	}
	return values
}

func (v valueHitMap) inverse() valueHitMap {
	return v ^ fullMapping
}

func (v valueHitMap) String() string {
	return fmt.Sprintf("%v", v.getValues())
}
