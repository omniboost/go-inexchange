package inexchange

import (
	"time"

	"github.com/cydev/zero"
)

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

type DocumentsOutboundSendDocumentAs struct {
	// - possible values: Paper, Electronic, Pdf, BusinessToConsumer
	Type               string                              `json:"Type"`
	Paper              DocumentsOutboundPaper              `json:"Paper,omitzero"`
	Electronic         DocumentsOutboundElectronic         `json:"Electronic,omitzero"`
	Pdf                DocumentsOutboundPDF                `json:"Pdf,omitzero"`
	BusinessToConsumer DocumentsOutboundBusinessToConsumer `json:"BusinessToConsumer,omitzero"`
}

type DocumentsOutboundPaper struct {
	RecipientAddress struct {
		Name        string `json:"Name"`
		Department  string `json:"Department"`
		StreetName  string `json:"StreetName"`
		PostBox     string `json:"PostBox"`
		PostalZone  string `json:"PostalZone"`
		City        string `json:"City"`
		CountryCode string `json:"CountryCode"`
	} `json:"RecipientAddress"`
	ReturnAddress struct {
		Name        string `json:"Name"`
		Department  string `json:"Department"`
		StreetName  string `json:"StreetName"`
		PostBox     string `json:"PostBox"`
		PostalZone  string `json:"PostalZone"`
		City        string `json:"City"`
		CountryCode string `json:"CountryCode"`
	} `json:"ReturnAddress"`
}

func (do DocumentsOutboundPaper) IsZero() bool {
	return zero.IsZero(do)
}

type DocumentsOutboundElectronic struct {
	RecipientID string `json:"RecipientId"`
}

func (do DocumentsOutboundElectronic) IsZero() bool {
	return zero.IsZero(do)
}

type DocumentsOutboundPDF struct {
	RecipientEmail string `json:"RecipientEmail"`
	RecipientName  string `json:"RecipientName"`
	SenderEmail    string `json:"SenderEmail"`
	SenderName     string `json:"SenderName"`
}

func (do DocumentsOutboundPDF) IsZero() bool {
	return zero.IsZero(do)
}

type DocumentsOutboundBusinessToConsumer struct {
	FMI      string `json:"FMI"`
	SSN      string `json:"SSN"`
	Provider string `json:"Provider"`
}

func (do DocumentsOutboundBusinessToConsumer) IsZero() bool {
	return zero.IsZero(do)
}

type DocumentsOutboundRecipientInformation struct {
	GLN         string `json:"GLN"`
	OrgNo       string `json:"OrgNo"`
	VatNo       string `json:"VatNo"`
	Name        string `json:"Name"`
	RecipientNo string `json:"RecipientNo"`
	CountryCode string `json:"CountryCode"`
}

func (do DocumentsOutboundRecipientInformation) IsZero() bool {
	return zero.IsZero(do)
}

type DocumentsOutboundDocument struct {
	DocumentFormat         string   `json:"DocumentFormat"`
	ErpDocumentID          string   `json:"ErpDocumentId"`
	DocumentURI            string   `json:"DocumentUri"`
	RenderedDocumentFormat string   `json:"RenderedDocumentFormat"`
	RenderedDocumentURI    string   `json:"RenderedDocumentUri"`
	Language               string   `json:"Language"`
	Culture                string   `json:"Culture"`
	Attachments            []string `json:"Attachments"`
}

func (do DocumentsOutboundDocument) IsZero() bool {
	return zero.IsZero(do)
}
