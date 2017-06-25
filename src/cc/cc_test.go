/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"testing"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	//"github.com/ianco/chaincode_1/model"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query_config")})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func TestConfigCC_Init(t *testing.T) {
	scc := new(ConfigCC)
	stub := shim.NewMockStub("ex02", scc)

	s1 := "{\"difficulty_rating\": 4, \"startup\": {\"ai_nation_count\": 3, \"start_up_cash\": 10, \"ai_start_up_cash\": 20, \"ai_aggressiveness\": 3, \"start_up_independent_town\": 15, \"start_up_raw_site\": 55, \"difficulty_level\": 4 }}"

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte(s1), []byte("tbd"), []byte("tbd")})

	checkState(t, stub, "Configuration", strings.Replace(s1, " ", "", -1))
}

func TestConfigCC_Query(t *testing.T) {
	scc := new(ConfigCC)
	stub := shim.NewMockStub("ex02", scc)

	s1 := "{\"difficulty_rating\": 4, \"startup\": {\"ai_nation_count\": 3, \"start_up_cash\": 10, \"ai_start_up_cash\": 20, \"ai_aggressiveness\": 3, \"start_up_independent_town\": 15, \"start_up_raw_site\": 55, \"difficulty_level\": 4 }}"

	// Init A=345 B=456
	checkInit(t, stub, [][]byte{[]byte("init"), []byte(s1), []byte("tbd"), []byte("tbd")})

	// Query A
	checkQuery(t, stub, "Configuration", strings.Replace(s1, " ", "", -1))
}

func TestConfigCC_Invoke(t *testing.T) {
	scc := new(ConfigCC)
	stub := shim.NewMockStub("ex02", scc)

	s1 := "{\"difficulty_rating\": 4, \"startup\": {\"ai_nation_count\": 3, \"start_up_cash\": 10, \"ai_start_up_cash\": 20, \"ai_aggressiveness\": 3, \"start_up_independent_town\": 15, \"start_up_raw_site\": 55, \"difficulty_level\": 4 }}"

	s2 := "{\"difficulty_rating\": 1, \"startup\": {\"ai_nation_count\": 3, \"start_up_cash\": 10, \"ai_start_up_cash\": 20, \"ai_aggressiveness\": 3, \"start_up_independent_town\": 15, \"start_up_raw_site\": 55, \"difficulty_level\": 4 }}"

	// Init A=567 B=678
	checkInit(t, stub, [][]byte{[]byte("init"), []byte(s1), []byte("tbd"), []byte("tbd")})

	// Invoke A->B for 123
	checkInvoke(t, stub, [][]byte{[]byte("update_config"), []byte(s2)})
	checkQuery(t, stub, "Configuration", strings.Replace(s2, " ", "", -1))

	// Invoke B->A for 234
	checkInvoke(t, stub, [][]byte{[]byte("update_config"), []byte(s1)})
	checkQuery(t, stub, "Configuration", strings.Replace(s1, " ", "", -1))
}
