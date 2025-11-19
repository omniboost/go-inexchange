package inexchange

import "time"

type Documents []Document

type Document struct {
	ID            string `json:"id"`
	DocumentType  string `json:"documentType"`
	CurrentStatus struct {
		Status string    `json:"status"`
		Time   time.Time `json:"time"`
	} `json:"currentStatus"`
	ErpDocumentID string `json:"erpDocumentId"`
}
