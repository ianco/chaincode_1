package main

import (
	"strings"
	"testing"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("ConfigTest")

func TestConfigConverstions(t *testing.T) {
	var c ConfigData
	var s1, s2 string
	s1 = "{\"difficulty_rating\": 4, \"startup\": {\"ai_nation_count\": 3, \"start_up_cash\": 10, \"ai_start_up_cash\": 20, \"ai_aggressiveness\": 3, \"start_up_independent_town\": 15, \"start_up_raw_site\": 55, \"difficulty_level\": 4 }}"
	log.Infof("json [%s]", s1)
	c, err1 := Json2Config(s1)
	if err1 != nil {
		t.Fatalf("Json2Config return error: %v", err1)
	}
	log.Infof("ConfigData [%+v]", c)
	s2, err2 := Config2Json(c)
	if err2 != nil {
		t.Fatalf("Config2Json return error: %v", err2)
	}
	log.Infof("returned [%s]", s2)
	if strings.Replace(s1, " ", "", -1) != strings.Replace(s2, " ", "", -1) {
		t.Fatalf("Error strings are not equal")
	} else {
		log.Infof("Strings are the same")
	}
}
