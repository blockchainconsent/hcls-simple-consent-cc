/*******************************************************************************
 * IBM Confidential
 *
 * OCO Source Materials
 *
 * (c) Copyright IBM Corporation 2022 All Rights Reserved.
 *
 * The source code for this program is not published or otherwise
 * divested of its trade secrets, irrespective of what has been
 * deposited with the U.S. Copyright Office.
 *******************************************************************************/

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	consentContract := new(ConsentContract)
	consentContract.Info.Version = "0.0.1"
	consentContract.Info.Description = "Simple Consent Smart Contract"
	consentContract.Info.License = new(metadata.LicenseMetadata)
	consentContract.Info.License.Name = "IBM"
	consentContract.Info.Contact = new(metadata.ContactMetadata)
	consentContract.Info.Contact.Name = "Richard M. Scott"

	chaincode, err := contractapi.NewChaincode(consentContract)
	chaincode.Info.Title = "Simple Consent Chaincode"
	chaincode.Info.Version = "0.0.1"
	if err != nil {
		panic("Failed to create chaincode from ConsentContract: " + err.Error())
	}

	err = chaincode.Start()
	if err != nil {
		panic("Failed to start chaincode: " + err.Error())
	}
}
