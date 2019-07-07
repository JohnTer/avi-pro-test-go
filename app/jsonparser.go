package main

import (
	"encoding/json"
)

func getRange(params *GenParams, dat map[string]interface{}) int {
	var temp float64

	RangeVar, getErr := dat["intRange"].([]interface{})
	if !getErr || len(RangeVar) != 2 {
		return 1
	}

	temp, getErr = RangeVar[0].(float64)
	if !getErr {
		return 1
	} else {
		params.intRange[0] = int(temp)
	}
	temp, getErr = RangeVar[1].(float64)
	if !getErr {
		return 1
	} else {
		params.intRange[1] = int(temp)
	}
	return 0

}

func getStrLen(params *GenParams, dat map[string]interface{}) int {
	var gerr int
	var getErr bool
	var temp float64
	temp, getErr = dat["strLen"].(float64)
	if !getErr {
		gerr = 1
		return gerr
	} else {
		params.strLen = int(temp)
	}
	return gerr
}

func getExtAlphabet(params *GenParams, dat map[string]interface{}) int {
	var gerr int
	var getErr bool
	params.extAlphabet, getErr = dat["extAlphabet"].(string)
	if !getErr {
		gerr = 1
		return gerr
	}
	return gerr
}

func getRandType(params *GenParams, dat map[string]interface{}) int {
	var gerr int
	var getErr bool
	params.randType, getErr = dat["randType"].(string)
	if !getErr {
		gerr = 1
		return gerr
	}
	return gerr
}

func getDat(js string, params *GenParams) (map[string]interface{}, int) {
	gerr := 0
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(js), &dat)
	if err != nil {
		gerr = 1
		return dat, gerr
	}
	return dat, gerr
}

func ParseJSON(js string) (GenParams, int) {
	var gerr int
	var params = GenParams{}
	dat, gerr := getDat(js, &params)
	if gerr != 0 {
		return params, gerr
	}

	gerr = getRandType(&params, dat)
	if gerr != 0 {
		return params, gerr
	}

	switch params.randType {
	case "num":
		gerr = getRange(&params, dat)
	case "extstr":
		gerr = getExtAlphabet(&params, dat)
		if gerr != 0 {
			return params, gerr
		}
		gerr = getStrLen(&params, dat)
	case "str":
		gerr = getStrLen(&params, dat)
	case "strnum":
		gerr = getStrLen(&params, dat)
	case "uuid":
		break
	default:
		gerr = 1
	}
	return params, gerr

}
