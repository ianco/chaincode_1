/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package integration

import (
	"fmt"
	//"strconv"
	"testing"

	//fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
)

func TestChainQueries(t *testing.T) {

	testSetup := &BaseSetupImpl{
		ConfigFile:      "../config_test.yaml",
		ChainID:         "mychannel",
		ChannelConfig:   "../mychannel.tx",
		ConnectEventHub: true,
	}

	if err := testSetup.Initialize(); err != nil {
		t.Fatalf(err.Error())
	}

	chain := testSetup.Chain
	//client := testSetup.Client

	Logger.Infof("InstallAndInstantiateMyCC()!!!")
	if err := testSetup.InstallAndInstantiateMyCC(); err != nil {
		t.Fatalf("InstallAndInstantiateMyCC return error: %v", err)
	}

	// Test Query Info - retrieve values before transaction
	Logger.Infof("QueryInfo()!!!")
	info, err := chain.QueryInfo()
	if err != nil {
		t.Fatalf("QueryInfo return error: %v", err)
	}
	Logger.Infof("QueryInfo [%s]", info)

	// run some queries on the "my-channel" chaincode
	Logger.Infof("!!! now testing custom chaincode")
	value, err := testMyChannelQueries(testSetup)
	if err != nil {
		Logger.Infof("testMyChannelQueries() return error: %v", err)
	} else {
		Logger.Infof("testMyChannelQueries() = %s", value)
	}

}

func testMyChannelQueries(testSetup *BaseSetupImpl) (string, error) {
	Logger.Infof("Query()!!!")

	value, err := testSetup.QueryMyAsset()
	if err != nil {
		return "", fmt.Errorf("QueryMyAsset return error: %v", err)
	}
	Logger.Infof("QueryMyAsset() = %s", value)

	value, err = testSetup.UpdateMyAsset()
	if err != nil {
		return "", fmt.Errorf("UpdateMyAsset return error: %v", err)
	}
	Logger.Infof("UpdateMyAsset() = %s", value)

	value, err = testSetup.QueryMyAsset()
	if err != nil {
		return "", fmt.Errorf("QueryMyAsset return error: %v", err)
	}
	Logger.Infof("QueryMyAsset() = %s", value)

	return value, err
}

