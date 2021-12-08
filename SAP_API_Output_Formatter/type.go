package sap_api_output_formatter

type SupplierInvoice struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	SupplierInvoice string `json:"supplier_invoice"`
	Deleted         string `json:"deleted"`
}

type Header struct {
	SupplierInvoice               string `json:"SupplierInvoice"`
	FiscalYear                    string `json:"FiscalYear"`
	CompanyCode                   string `json:"CompanyCode"`
	DocumentDate                  string `json:"DocumentDate"`
	PostingDate                   string `json:"PostingDate"`
	SupplierInvoiceIDByInvcgParty string `json:"SupplierInvoiceIDByInvcgParty"`
	InvoicingParty                string `json:"InvoicingParty"`
	DocumentCurrency              string `json:"DocumentCurrency"`
	InvoiceGrossAmount            string `json:"InvoiceGrossAmount"`
	DocumentHeaderText            string `json:"DocumentHeaderText"`
	PaymentTerms                  string `json:"PaymentTerms"`
	DueCalculationBaseDate        string `json:"DueCalculationBaseDate"`
	NetPaymentDays                string `json:"NetPaymentDays"`
	PaymentBlockingReason         string `json:"PaymentBlockingReason"`
	AccountingDocumentType        string `json:"AccountingDocumentType"`
	BPBankAccountInternalID       string `json:"BPBankAccountInternalID"`
	SupplierInvoiceStatus         string `json:"SupplierInvoiceStatus"`
	DirectQuotedExchangeRate      string `json:"DirectQuotedExchangeRate"`
	SupplyingCountry              string `json:"SupplyingCountry"`
	PaymentMethod                 string `json:"PaymentMethod"`
	InvoiceReference              string `json:"InvoiceReference"`
	SupplierPostingLineItemText   string `json:"SupplierPostingLineItemText"`
	TaxIsCalculatedAutomatically  bool   `json:"TaxIsCalculatedAutomatically"`
	BusinessArea                  string `json:"BusinessArea"`
	SupplierInvoiceIsCreditMemo   string `json:"SupplierInvoiceIsCreditMemo"`
	ReverseDocument               string `json:"ReverseDocument"`
	ReverseDocumentFiscalYear     string `json:"ReverseDocumentFiscalYear"`
}

type Tax struct {
	SupplierInvoice          string `json:"SupplierInvoice"`
	FiscalYear               string `json:"FiscalYear"`
	TaxCode                  string `json:"TaxCode"`
	DocumentCurrency         string `json:"DocumentCurrency"`
	TaxAmount                string `json:"TaxAmount"`
	TaxBaseAmountInTransCrcy string `json:"TaxBaseAmountInTransCrcy"`
}

type Account struct {
	SupplierInvoice               string `json:"SupplierInvoice"`
	FiscalYear                    string `json:"FiscalYear"`
	SupplierInvoiceItem           string `json:"SupplierInvoiceItem"`
	CompanyCode                   string `json:"CompanyCode"`
	CostCenter                    string `json:"CostCenter"`
	ControllingArea               string `json:"ControllingArea"`
	BusinessArea                  string `json:"BusinessArea"`
	ProfitCenter                  string `json:"ProfitCenter"`
	FunctionalArea                string `json:"FunctionalArea"`
	GLAccount                     string `json:"GLAccount"`
	SalesOrder                    string `json:"SalesOrder"`
	SalesOrderItem                string `json:"SalesOrderItem"`
	CostObject                    string `json:"CostObject"`
	WBSElement                    string `json:"WBSElement"`
	DocumentCurrency              string `json:"DocumentCurrency"`
	SuplrInvcAcctAssignmentAmount string `json:"SuplrInvcAcctAssignmentAmount"`
	TaxCode                       string `json:"TaxCode"`
	WorkItem                      string `json:"WorkItem"`
	MasterFixedAsset              string `json:"MasterFixedAsset"`
	FixedAsset                    string `json:"FixedAsset"`
	DebitCreditCode               string `json:"DebitCreditCode"`
	InternalOrder                 string `json:"InternalOrder"`
	ProjectNetwork                string `json:"ProjectNetwork"`
	ProfitabilitySegment          string `json:"ProfitabilitySegment"`
}

type PurchaseOrder           struct {
	SupplierInvoice                string `json:"SupplierInvoice"`
	FiscalYear                     string `json:"FiscalYear"`
	SupplierInvoiceItem            string `json:"SupplierInvoiceItem"`
	PurchaseOrder                  string `json:"PurchaseOrder"`
	PurchaseOrderItem              string `json:"PurchaseOrderItem"`
	Plant                          string `json:"Plant"`
	TaxCode                        string `json:"TaxCode"`
	DocumentCurrency               string `json:"DocumentCurrency"`
	SupplierInvoiceItemAmount      string `json:"SupplierInvoiceItemAmount"`
	PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
	QuantityInPurchaseOrderUnit    string `json:"QuantityInPurchaseOrderUnit"`
	PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
	QtyInPurchaseOrderPriceUnit    string `json:"QtyInPurchaseOrderPriceUnit"`
	SupplierInvoiceItemText        string `json:"SupplierInvoiceItemText"`
	PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
}
