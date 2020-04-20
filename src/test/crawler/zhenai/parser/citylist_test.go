package parser

import (
	"io/ioutil"
	"testing"
)

const resultSize = 470

func TestParseCityList(t *testing.T) {
	bytes, e := ioutil.ReadFile("citylist_test_data.html")
	if e != nil {
		panic(e)
	}
	result := ParseCityList(bytes)
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d,but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d,but had %d", resultSize, len(result.Items))
	}
}
