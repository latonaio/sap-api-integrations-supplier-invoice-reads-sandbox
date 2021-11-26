package sap_api_caller

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
     SupplierInvoice               string      `json:"SupplierInvoice"`
     FiscalYear                    int         `json:"FiscalYear"`
     CompanyCode                   string      `json:"CompanyCode"`
     DocumentDate                  string      `json:"DocumentDate"`
     PostingDate                   string      `json:"PostingDate"`
     SupplierInvoiceIDByInvcgParty string      `json:"SupplierInvoiceIDByInvcgParty"`
     InvoicingParty                string      `json:"InvoicingParty"`
     DocumentCurrency              string      `json:"DocumentCurrency"`
     InvoiceGrossAmount            float64     `json:"InvoiceGrossAmount"`
     DocumentHeaderText            string      `json:"DocumentHeaderText"`
     PaymentTerms                  string      `json:"PaymentTerms"`
     DueCalculationBaseDate        string      `json:"DueCalculationBaseDate"`
     NetPaymentDays                int         `json:"NetPaymentDays"`
     PaymentBlockingReason         string      `json:"PaymentBlockingReason"`
     AccountingDocumentType        string      `json:"AccountingDocumentType"`
     BPBankAccountInternalID       string      `json:"BPBankAccountInternalID"`
     SupplierInvoiceStatus         string      `json:"SupplierInvoiceStatus"`
     DirectQuotedExchangeRate      string      `json:"DirectQuotedExchangeRate"`
     SupplyingCountry              string      `json:"SupplyingCountry"`
     PaymentMethod                 string      `json:"PaymentMethod"`
     InvoiceReference              string      `json:"InvoiceReference"`
     SupplierPostingLineItemText   string      `json:"SupplierPostingLineItemText"`
     TaxIsCalculatedAutomatically  string      `json:"TaxIsCalculatedAutomatically"`
     BusinessArea                  string      `json:"BusinessArea"`
     SupplierInvoiceIsCreditMemo   string      `json:"SupplierInvoiceIsCreditMemo"`
     ReverseDocument               string      `json:"ReverseDocument"`
     ReverseDocumentFiscalYear     int         `json:"ReverseDocumentFiscalYear"`
}

type Tax struct {
     SupplierInvoice               string      `json:"SupplierInvoice"`
     FiscalYear                    int         `json:"FiscalYear"`
     TaxCode                       string      `json:"TaxCode"`
     DocumentCurrency              string      `json:"DocumentCurrency"`
     TaxAmount                     float64     `json:"TaxAmount"`
     TaxBaseAmountInTransCrcy      float64     `json:"TaxBaseAmountInTransCrcy"`
}

type PurchaseOrderReference struct {
     SupplierInvoice                string     `json:"SupplierInvoice"`
     FiscalYear                     int        `json:"FiscalYear"`
     SupplierInvoiceItem            int        `json:"SupplierInvoiceItem"`
     PurchaseOrder                  string     `json:"PurchaseOrder"`
     PurchaseOrderItem              int        `json:"PurchaseOrderItem"`
     Plant                          string     `json:"Plant"`
     TaxCode                        string     `json:"TaxCode"`
     DocumentCurrency               string     `json:"DocumentCurrency"`
     SupplierInvoiceItemAmount      float64    `json:"SupplierInvoiceItemAmount"`
     PurchaseOrderQuantityUnit      string     `json:"PurchaseOrderQuantityUnit"`
     QuantityInPurchaseOrderUnit    float64    `json:"QuantityInPurchaseOrderUnit"`
     PurchaseOrderPriceUnit         int        `json:"PurchaseOrderPriceUnit"`
     QtyInPurchaseOrderPriceUnit    float64    `json:"QtyInPurchaseOrderPriceUnit"`
     SupplierInvoiceItemText        string     `json:"SupplierInvoiceItemText"`
     PurchasingDocumentItemCategory string     `json:"PurchasingDocumentItemCategory"`
}

type AccountAssignment struct {
     SupplierInvoice               string      `json:"SupplierInvoice"`
     FiscalYear                    int         `json:"FiscalYear"`
     SupplierInvoiceItem           int         `json:"SupplierInvoiceItem"`
     CompanyCode                   string      `json:"CompanyCode"`
     CostCenter                    string      `json:"CostCenter"`
     ControllingArea               string      `json:"ControllingArea"`
     BusinessArea                  string      `json:"BusinessArea"`
     ProfitCenter                  string      `json:"ProfitCenter"`
     FunctionalArea                string      `json:"FunctionalArea"`
     GLAccount                     string      `json:"GLAccount"`
     SalesOrder                    string      `json:"SalesOrder"`
     SalesOrderItem                string      `json:"SalesOrderItem"`
     CostObject                    string      `json:"CostObject"`
     WBSElement                    string      `json:"WBSElement"`
     DocumentCurrency              string      `json:"DocumentCurrency"`
     SuplrInvcAcctAssignmentAmount float64     `json:"SuplrInvcAcctAssignmentAmount"`
     TaxCode                       string      `json:"TaxCode"`
     WorkItem                      string      `json:"WorkItem"`
     MasterFixedAsset              string      `json:"MasterFixedAsset"`
     FixedAsset                    string      `json:"FixedAsset"`
     DebitCreditCode               string      `json:"DebitCreditCode"`
     InternalOrder                 string      `json:"InternalOrder"`
     ProjectNetwork                string      `json:"ProjectNetwork"`
     ProfitabilitySegment          int         `json:"ProfitabilitySegment"`
}
