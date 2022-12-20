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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MaxItemsAllowed - defines the max Consent items that can be returned for GetConsentHistory
const MaxItemsAllowed = 100

// MinKeyLength - defines the minimum key length
const MinKeyLength = 4

// ConsentContract contract for managing CRUD for Consent
type ConsentContract struct {
	contractapi.Contract
}

func (c *ConsentContract) Ping(pCtx contractapi.TransactionContextInterface) (string, error) {

	log.Print("Ping: enter")
	defer log.Print("Ping: exit")

	dt := time.Now()
	log.Println("Current date and time is: ", dt.String())

	return "returned from BC Successfully", nil

} // end of Ping

// ConsentExists returns true when Consent with given ID exists in world state
func (c *ConsentContract) ConsentExists(ctx contractapi.TransactionContextInterface, ConsentID string) (bool, error) {
	log.Print("ConsentExists: enter")
	defer log.Print("ConsentExists: exit")

	data, err := ctx.GetStub().GetState(ConsentID)
	if err != nil {
		return false, err
	}

	return data != nil, nil

} // end of ConsentExists

// CreateConsent creates a new instance of Consent
func (c *ConsentContract) CreateConsent(
	pCtx contractapi.TransactionContextInterface,
	pConsentId string,
	pPatientId string,
	pServiceId string,
	pTenantId string,
	pDataTypeIds []string,
	pConsentOption []string,
	pCreationTime int64,
	pExpirationTime int64,
	pFHIRResourceID string,
	pFHIRResourceVersion string,
	pFHIRPolicy string,
	pFHIRStatus string,
	pFHIRProvisionType string,
	pFHIRProvisionAction string,
	pFHIRPerformerIDSystem string,
	pFHIRPerformerIDValue string,
	pFHIRPerformerDisplay string,
	pFHIRRecipientIDSystem string,
	pFHIRRecipientIDValue string,
	pFHIRRecipientDisplay string,
) error {

	log.Print("CreateConsent: enter")
	defer log.Print("CreateConsent: exit")

	if len(pConsentId) < MinKeyLength {
		errorString := "CreateConsent: ERROR key minimum length is = " + strconv.Itoa(MinKeyLength)
		log.Println(errorString)
		return errors.New(errorString)
	}

	exists, err := c.ConsentExists(pCtx, pConsentId)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("the Consent %s already exists", pConsentId)
	}

	consent := new(Consent)
	consent.ConsentID = pConsentId
	consent.PatientID = pPatientId
	consent.ServiceID = pServiceId
	consent.TenantID = pTenantId
	consent.DatatypeIDs = pDataTypeIds
	consent.ConsentOption = pConsentOption
	consent.Creation = pCreationTime
	consent.Expiration = pExpirationTime
	consent.FHIRResourceID = pFHIRResourceID
	consent.FHIRResourceVersion = pFHIRResourceVersion
	consent.FHIRPolicy = pFHIRPolicy
	consent.FHIRStatus = pFHIRStatus
	consent.FHIRProvisionType = pFHIRProvisionType
	consent.FHIRProvisionAction = pFHIRProvisionAction
	consent.FHIRPerformerIDSystem = pFHIRPerformerIDSystem
	consent.FHIRPerformerIDValue = pFHIRPerformerIDValue
	consent.FHIRPerformerDisplay = pFHIRPerformerDisplay
	consent.FHIRRecipientIDSystem = pFHIRRecipientIDSystem
	consent.FHIRRecipientIDValue = pFHIRRecipientIDValue
	consent.FHIRRecipientDisplay = pFHIRRecipientDisplay

	bytes, _ := json.Marshal(consent)
	return pCtx.GetStub().PutState(pConsentId, bytes)

} // end of CreateConsent

// ReadConsent retrieves an instance of Consent from the world state
func (c *ConsentContract) ReadConsent(pCtx contractapi.TransactionContextInterface, pConsentId string) (*Consent, error) {
	log.Print("ReadConsent: enter")
	defer log.Print("ReadConsent: exit")

	exists, err := c.ConsentExists(pCtx, pConsentId)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the Consent %s does not exist", pConsentId)
	}

	bytes, _ := pCtx.GetStub().GetState(pConsentId)

	consent := new(Consent)

	err = json.Unmarshal(bytes, consent)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Consent")
	}

	return consent, nil

} // end of ReadConsent

// UpdateConsent retrieves an instance of Consent from the world state and updates its value
func (c *ConsentContract) UpdateConsent(
	pCtx contractapi.TransactionContextInterface,
	pConsentId string,
	pPatientId string,
	pServiceId string,
	pTenantId string,
	pDataTypeIds []string,
	pConsentOption []string,
	pCreationTime int64,
	pExpirationTime int64,
	pFHIRResourceID string,
	pFHIRResourceVersion string,
	pFHIRPolicy string,
	pFHIRStatus string,
	pFHIRProvisionType string,
	pFHIRProvisionAction string,
	pFHIRPerformerIDSystem string,
	pFHIRPerformerIDValue string,
	pFHIRPerformerDisplay string,
	pFHIRRecipientIDSystem string,
	pFHIRRecipientIDValue string,
	pFHIRRecipientDisplay string,
) error {

	log.Print("UpdateConsent: enter")
	defer log.Print("UpdateConsent: exit")

	exists, err := c.ConsentExists(pCtx, pConsentId)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the Consent %s does not exist", pConsentId)
	}

	consent := new(Consent)
	consent.ConsentID = pConsentId
	consent.PatientID = pPatientId
	consent.ServiceID = pServiceId
	consent.TenantID = pTenantId
	consent.DatatypeIDs = pDataTypeIds
	consent.ConsentOption = pConsentOption
	consent.Creation = pCreationTime
	consent.Expiration = pExpirationTime
	consent.FHIRResourceID = pFHIRResourceID
	consent.FHIRResourceVersion = pFHIRResourceVersion
	consent.FHIRPolicy = pFHIRPolicy
	consent.FHIRStatus = pFHIRStatus
	consent.FHIRProvisionType = pFHIRProvisionType
	consent.FHIRProvisionAction = pFHIRProvisionAction
	consent.FHIRPerformerIDSystem = pFHIRPerformerIDSystem
	consent.FHIRPerformerIDValue = pFHIRPerformerIDValue
	consent.FHIRPerformerDisplay = pFHIRPerformerDisplay
	consent.FHIRRecipientIDSystem = pFHIRRecipientIDSystem
	consent.FHIRRecipientIDValue = pFHIRRecipientIDValue
	consent.FHIRRecipientDisplay = pFHIRRecipientDisplay

	bytes, _ := json.Marshal(consent)

	return pCtx.GetStub().PutState(pConsentId, bytes)
}

// DeleteConsent deletes an instance of Consent from the world state
func (c *ConsentContract) DeleteConsent(pCtx contractapi.TransactionContextInterface, pConsentId string) error {

	log.Print("DeleteConsent: enter")
	defer log.Print("DeleteConsent: exit")

	exists, err := c.ConsentExists(pCtx, pConsentId)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the Consent %s does not exist", pConsentId)
	}

	return pCtx.GetStub().DelState(pConsentId)

} // end of DeleteConsent

// GetConsentHistory returns the chain of custody for an asset since issuance.
func (t *ConsentContract) GetConsentHistory(ctx contractapi.TransactionContextInterface, pConsentId string) ([]*HistoryQueryResult, error) {

	log.Print("GetConsentHistory: enter")
	defer log.Print("GetConsentHistory: exit")

	log.Printf("GetConsentHistory: ID %v", pConsentId)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(pConsentId)
	if err != nil {
		log.Println("GetConsentHistory() : ERROR getting history for key resultsIterator")
		return nil, errors.New("GetConsentHistory() : ERROR calling GetHistoryForKey")
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()

		if err != nil {
			log.Println("GetConsentHistory() : ERROR getting next results iterator")
			return nil, errors.New("ERROR getting next results iterator")
		}

		var idxConsent Consent

		if response.IsDelete || len(response.Value) < 1 {
			// IBP has in issue mapping objects if they are nil
			// so set any empty objects,  this could happen in the case of a delete

			// if nil copy empty key
			idxConsent.ConsentID = pConsentId
			// if null copy in empty array
			dataTypeIds := []string{"", ""}
			idxConsent.DatatypeIDs = dataTypeIds
		} else {
			err := json.Unmarshal(response.Value, &idxConsent)
			if err != nil {
				log.Println("GetConsentHistory() : consent record with error = ", idxConsent)
				log.Println("GetConsentHistory() : ERROR unmarshalling iterations")
				return nil, errors.New("GetConsentHistory() : ERROR unmarshalling iterations")
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			log.Println("GetConsentHistory() : ERROR getting TimeStanp")
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &idxConsent,
			IsDelete:  response.IsDelete,
		}

		records = append(records, &record)
	}

	return records, nil

} // end of GetConsentHistory

// constructQueryResponseFromIterator constructs a slice of Consents from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*Consent, error) {

	log.Println("constructQueryResponseFromIterator: enter")
	defer log.Println("constructQueryResponseFromIterator: exit")

	var Consents = make([]*Consent, 0)

	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var Consent Consent
		err = json.Unmarshal(queryResult.Value, &Consent)
		if err != nil {
			return nil, err
		}
		Consents = append(Consents, &Consent)
	}

	return Consents, nil

} // end of constructQueryResponseFromIterator

// takes a single argument that is JSON of the KeyHash to write to the ledger
func (cc *ConsentContract) ReadAllConsents(pCtx contractapi.TransactionContextInterface) ([]*Consent, error) {
	log.Println("ReadAllConsents: enter")
	defer log.Println("ReadAllConsents: exit")

	stub := pCtx.GetStub()
	resultsIterator, readErr := stub.GetStateByRange("", "")
	if readErr != nil {
		log.Println("read: Error call GetStateByRange:", readErr)
		return nil, readErr
	}

	return constructQueryResponseFromIterator(resultsIterator)

} // end of ReadAllConsents

// GetConsentByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *ConsentContract) GetConsentByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Consent, error) {

	log.Println("GetConsentByRange: enter")
	defer log.Println("GetConsentByRange: exit")

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)

} // end of GetConsentByRange

// QueryConsentByPatientID queries for Consent based on PatientID and TenantID.
// Only available on state databases that support rich query (e.g. CouchDB)
// Sorts by Creation in desc
// Example: Parameterized rich query
func (t *ConsentContract) QueryConsentByPatientID(ctx contractapi.TransactionContextInterface, patientID string, tenantID string) ([]*Consent, error) {
	log.Print("QueryConsentByPatientID: enter")
	defer log.Print("QueryConsentByPatientID: exit")

	queryString := fmt.Sprintf(`{"selector":{"PatientID":"%s", "TenantID":"%s"}, "sort":[{"Creation": "desc"}], "use_index":["_design/indexPatientIDDoc", "indexPatientID"]}`, patientID, tenantID)

	return getQueryResultForQueryString(ctx, queryString)
}

// QueryConsent uses a query string to perform a query for Consent.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryConsentForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *ConsentContract) QueryConsent(ctx contractapi.TransactionContextInterface, queryString string) ([]*Consent, error) {
	return getQueryResultForQueryString(ctx, queryString)
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Consent, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// GetConsentByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *ConsentContract) GetConsentByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) ([]*Consent, error) {

	resultsIterator, _, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// QueryConsentWithPagination uses a query string, page size and a bookmark to perform a query
// for Consent. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryConsentForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *ConsentContract) QueryConsentWithPagination(ctx contractapi.TransactionContextInterface, patientID string, tenantID string, pageSize int, bookmark string) (*PaginatedQueryResult, error) {
	queryString := fmt.Sprintf(`{"selector":{"PatientID":"%s", "TenantID":"%s"}, "sort":[{"Creation": "desc"}], "use_index":["_design/indexPatientIDDoc", "indexPatientID"]}`, patientID, tenantID)
	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	Consent, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             Consent,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil

} // end of getQueryResultForQueryStringWithPagination
