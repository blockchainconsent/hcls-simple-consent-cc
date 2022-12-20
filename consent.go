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

import "time"

type Consent struct {
	ConsentID             string   `json:"ConsentID"`
	PatientID             string   `json:"PatientID"`
	ServiceID             string   `json:"ServiceID"`
	TenantID              string   `json:"TenantID"`
	DatatypeIDs           []string `json:"DatatypeIDs"`
	ConsentOption         []string `json:"ConsentOption"`
	Creation              int64    `json:"Creation"`
	Expiration            int64    `json:"Expiration"`
	FHIRResourceID        string   `json:"FHIRResourceID"`
	FHIRResourceVersion   string   `json:"FHIRResourceVersion"`
	FHIRPolicy            string   `json:"FHIRPolicy"`
	FHIRStatus            string   `json:"FHIRStatus"`
	FHIRProvisionType     string   `json:"FHIRProvisionType"`
	FHIRProvisionAction   string   `json:"FHIRProvisionAction"`
	FHIRPerformerIDSystem string   `json:"FHIRPerformerIDSystem"`
	FHIRPerformerIDValue  string   `json:"FHIRPerformerIDValue"`
	FHIRPerformerDisplay  string   `json:"FHIRPerformerDisplay"`
	FHIRRecipientIDSystem string   `json:"FHIRRecipientIDSystem"`
	FHIRRecipientIDValue  string   `json:"FHIRRecipientIDValue"`
	FHIRRecipientDisplay  string   `json:"FHIRRecipientDisplay"`
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *Consent  `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*Consent `json:"records"`
	FetchedRecordsCount int32      `json:"fetchedRecordsCount"`
	Bookmark            string     `json:"bookmark"`
}
