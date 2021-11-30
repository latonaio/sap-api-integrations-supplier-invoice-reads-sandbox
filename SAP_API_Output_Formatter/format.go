package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-supplier-invoice-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToSupplierInvoice(raw []byte, l *logger.Logger) *SupplierInvoice {
	pm := &responses.SupplierInvoice{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &SupplierInvoice{
		SupplierInvoice                data.SupplierInvoice,
		FiscalYear                     data.FiscalYear,
		CompanyCode                    data.CompanyCode,
		DocumentDate                   data.DocumentDate,
		PostingDate                    data.PostingDate,
		SupplierInvoiceIDByInvcgParty  data.SupplierInvoiceIDByInvcgParty,
		InvoicingParty                 data.InvoicingParty,
		DocumentCurrency               data.DocumentCurrency,
		InvoiceGrossAmount             data.InvoiceGrossAmount,
		DocumentHeaderText             data.DocumentHeaderText,
		PaymentTerms                   data.PaymentTerms,
		DueCalculationBaseDate         data.DueCalculationBaseDate,
		NetPaymentDays                 data.NetPaymentDays,
		PaymentBlockingReason          data.PaymentBlockingReason,
		AccountingDocumentType         data.AccountingDocumentType,
		BPBankAccountInternalID        data.BPBankAccountInternalID,
		SupplierInvoiceStatus          data.SupplierInvoiceStatus,
		DirectQuotedExchangeRate       data.DirectQuotedExchangeRate,
		SupplyingCountry               data.SupplyingCountry,
		PaymentMethod                  data.PaymentMethod,
		InvoiceReference               data.InvoiceReference,
		SupplierPostingLineItemText    data.SupplierPostingLineItemText,
		TaxIsCalculatedAutomatically   data.TaxIsCalculatedAutomatically,
		BusinessArea                   data.BusinessArea,
		SupplierInvoiceIsCreditMemo    data.SupplierInvoiceIsCreditMemo,
		ReverseDocument                data.ReverseDocument,
		ReverseDocumentFiscalYear      data.ReverseDocumentFiscalYear,
		TaxCode                        data.TaxCode,
		DocumentCurrency               data.DocumentCurrency,
		TaxAmount                      data.TaxAmount,
		TaxBaseAmountInTransCrcy       data.TaxBaseAmountInTransCrcy,
		SupplierInvoiceItem            data.SupplierInvoiceItem,
		PurchaseOrder                  data.PurchaseOrder,
		PurchaseOrderItem              data.PurchaseOrderItem,
		Plant                          data.Plant,
		TaxCode                        data.TaxCode,
		DocumentCurrency               data.DocumentCurrency,
		SupplierInvoiceItemAmount      data.SupplierInvoiceItemAmount,
		PurchaseOrderQuantityUnit      data.PurchaseOrderQuantityUnit,
		QuantityInPurchaseOrderUnit    data.QuantityInPurchaseOrderUnit,
		PurchaseOrderPriceUnit         data.PurchaseOrderPriceUnit,
		QtyInPurchaseOrderPriceUnit    data.QtyInPurchaseOrderPriceUnit,
		SupplierInvoiceItemText        data.SupplierInvoiceItemText,
		PurchasingDocumentItemCategory data.PurchasingDocumentItemCategory,
		SupplierInvoiceItem            data.SupplierInvoiceItem,
		CompanyCode                    data.CompanyCode,
		CostCenter                     data.CostCenter,
		ControllingArea                data.ControllingArea,
		BusinessArea                   data.BusinessArea,
		ProfitCenter                   data.ProfitCenter,
		FunctionalArea                 data.FunctionalArea,
		GLAccount                      data.GLAccount,
		SalesOrder                     data.SalesOrder,
		SalesOrderItem                 data.SalesOrderItem,
		CostObject                     data.CostObject,
		WBSElement                     data.WBSElement,
		DocumentCurrency               data.DocumentCurrency,
		SuplrInvcAcctAssignmentAmount  data.SuplrInvcAcctAssignmentAmount,
		TaxCode                        data.TaxCode,
		WorkItem                       data.WorkItem,
		MasterFixedAsset               data.MasterFixedAsset,
		FixedAsset                     data.FixedAsset,
		DebitCreditCode                data.DebitCreditCode,
		InternalOrder                  data.InternalOrder,
		ProjectNetwork                 data.ProjectNetwork,
		ProfitabilitySegment           data.ProfitabilitySegment,
	}
}
