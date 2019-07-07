package main

import (
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

type GenParams struct {
	randType    string
	intRange    [2]int
	extAlphabet string
	strLen      int
}

const lat string = "abcdefghijklmnopqrstuvwxyz"
const latNum string = "abcdefghijklmnopqrstuvwxyz1234567890"

func Generate(params *GenParams) (string, int) {
	var result string
	var err = 0
	switch params.randType {
	case "num":
		result = genNum(params.intRange[0], params.intRange[1])
	case "str":
		result = genString(params.strLen, lat)
	case "strnum":
		result = genString(params.strLen, latNum)
	case "uuid":
		result = genUUID(&err)
	case "extstr":
		result = genString(params.strLen, params.extAlphabet)
	default:
		err = 1
	}
	return result, err

}

func genNum(a int, b int) string {
	var num = rand.Intn(a+b) + a
	return strconv.Itoa(num)

}

func genString(length int, alphabet string) string {
	var alphabetLen = len(alphabet)
	var result string
	for i := 0; i < length; i++ {
		var alphabetIndex = rand.Intn(alphabetLen)
		result += string(alphabet[alphabetIndex])

	}
	return result
}

func genUUID(error *int) string {
	id, err := uuid.NewUUID()
	if err != nil {
		*error = 1
		return ""
	}
	return id.String()

}
