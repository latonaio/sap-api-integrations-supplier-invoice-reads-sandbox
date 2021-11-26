package file_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	SupplierInvoice   struct {
		SupplierInvoice                string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		QuantityInPurchaseOrderUnit    float64     `json:"quantity"`
		PickedQuantity                 float64     `json:"picked_quantity"`
		SupplierInvoiceItemAmount      float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	MaterialCode            string      `json:"material_code"`
	InvoicingParty          string      `json:"plant/supplier"`
	Stock                   float64     `json:"stock"`
	AccountingDocumentType  string      `json:"document_type"`
	SupplierInvoice         string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	PostingDate             string      `json:"validated_date"`
	Deleted                 string      `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	SupplierInvoice struct {
		SupplierInvoice               string      `json:"SupplierInvoice"`
		FiscalYear                    int         `json:"FiscalYear"`
		CompanyCode                   string      `json:"CompanyCode"`
		DocumentDate                  string      `json:"DocumentDate"`
		PostingDate                   string      `json:"PostingDate"`
		SupplierInvoiceIDByInvcgParty string      `json:"SupplierInvoiceIDByInvcgParty"`
		InvoicingParty                string      `json:"InvoicingParty"`
		DocumentCurrency              string      `json:"DocumentCurrency"`
		InvoiceGrossAmount            float64    `json:"InvoiceGrossAmount"`
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
		Tax                           struct {
			TaxCode                  string `json:"TaxCode"`
			DocumentCurrency         string `json:"DocumentCurrency"`
			TaxAmount                float64 `json:"TaxAmount"`
			TaxBaseAmountInTransCrcy float64 `json:"TaxBaseAmountInTransCrcy"`
		} `json:"Tax"`
		PurchaseOrderReference struct {
			SupplierInvoiceItem            int    `json:"SupplierInvoiceItem"`
			PurchaseOrder                  string `json:"PurchaseOrder"`
			PurchaseOrderItem              int    `json:"PurchaseOrderItem"`
			Plant                          string `json:"Plant"`
			TaxCode                        string `json:"TaxCode"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			SupplierInvoiceItemAmount      float64 `json:"SupplierInvoiceItemAmount"`
			PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
			QuantityInPurchaseOrderUnit    float64 `json:"QuantityInPurchaseOrderUnit"`
			PurchaseOrderPriceUnit         int    `json:"PurchaseOrderPriceUnit"`
			QtyInPurchaseOrderPriceUnit    float64 `json:"QtyInPurchaseOrderPriceUnit"`
			SupplierInvoiceItemText        string `json:"SupplierInvoiceItemText"`
			PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
		} `json:"PurchaseOrderReference"`
		AccountAssignment struct {
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
		} `json:"AccountAssignment"`
	} `json:"SupplierInvoice"`
	APISchema       string `json:"api_schema"`
	SupplierInvoice string `json:"supplier_invoice"`
	Deleted         string `json:"deleted"`
}
