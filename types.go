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

type Parties []Party

type Party struct {
	CompanyID                    string   `json:"companyId"`
	Name                         string   `json:"name"`
	AltName                      string   `json:"altName"`
	StreetName                   string   `json:"streetName"`
	PostBox                      string   `json:"postBox"`
	City                         string   `json:"city"`
	PostalZone                   string   `json:"postalZone"`
	CountryCode                  string   `json:"countryCode"`
	PhoneNo                      string   `json:"phoneNo"`
	OrgNo                        string   `json:"orgNo"`
	VatNo                        string   `json:"vatNo"`
	GLN                          string   `json:"gln"`
	PeppolParticipantIdentifiers []string `json:"peppolParticipantIdentifiers"`
	Address                      string   `json:"address"`
	Address2                     string   `json:"address2"`
	PostalCode                   string   `json:"postalCode"`
	Connected                    bool     `json:"connected"`
	// Which capabilities a company may have to receive e-invoices
	// * SendsElectronicInvoices - Company can send e-invoices
	// * CannotSendElectronicInvoices - Company cannot send e-invoices
	// * ContactRequiredBeforeSendingElectronicInvoices - Company sends e-invoices. You have to contact the company first to setup e-invoicing
	// * ErpSystemSupportsSendingElectronicInvoices - Company uses an ERP-system that support sending e-invoices, but it’s not activated
	ReceiveElectronicInvoiceCapability string `json:"receiveElectronicInvoiceCapability"`
	// Which capabilities a company may have to receive e-invoices
	// * ReceivesElectronicOrders - Company can receive e-order
	// * CannotReceiveElectronicOrders - Company cannot receive e-order
	// * ContactRequiredBeforeReceivingElectronicOrders - Company receives e-order. You can send orders but expect that they will be stopped unless you have contacted the company first and agreed on which information that is required
	// * ErpSystemSupportsReceivingElectronicOrders - Company uses an ERP-system that support receiving e-orders, but it’s not activated
	SendsElectronicOrderCapability string `json:"sendsElectronicOrderCapability"`
}
