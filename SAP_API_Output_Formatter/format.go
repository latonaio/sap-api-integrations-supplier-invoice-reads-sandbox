package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-supplier-invoice-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Header{
			SupplierInvoice:               data.SupplierInvoice,
			FiscalYear:                    data.FiscalYear,
			CompanyCode:                   data.CompanyCode,
			DocumentDate:                  data.DocumentDate,
			PostingDate:                   data.PostingDate,
			SupplierInvoiceIDByInvcgParty: data.SupplierInvoiceIDByInvcgParty,
			InvoicingParty:                data.InvoicingParty,
			DocumentCurrency:              data.DocumentCurrency,
			InvoiceGrossAmount:            data.InvoiceGrossAmount,
			DocumentHeaderText:            data.DocumentHeaderText,
			PaymentTerms:                  data.PaymentTerms,
			DueCalculationBaseDate:        data.DueCalculationBaseDate,
			NetPaymentDays:                data.NetPaymentDays,
			PaymentBlockingReason:         data.PaymentBlockingReason,
			AccountingDocumentType:        data.AccountingDocumentType,
			BPBankAccountInternalID:       data.BPBankAccountInternalID,
			SupplierInvoiceStatus:         data.SupplierInvoiceStatus,
			DirectQuotedExchangeRate:      data.DirectQuotedExchangeRate,
			SupplyingCountry:              data.SupplyingCountry,
			PaymentMethod:                 data.PaymentMethod,
			InvoiceReference:              data.InvoiceReference,
			SupplierPostingLineItemText:   data.SupplierPostingLineItemText,
			TaxIsCalculatedAutomatically:  data.TaxIsCalculatedAutomatically,
			BusinessArea:                  data.BusinessArea,
			SupplierInvoiceIsCreditMemo:   data.SupplierInvoiceIsCreditMemo,
			ReverseDocument:               data.ReverseDocument,
			ReverseDocumentFiscalYear:     data.ReverseDocumentFiscalYear,
	}, nil
}

func ConvertToTax(raw []byte, l *logger.Logger) (*Tax, error) {
	pm := &responses.Tax{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Tax. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Tax{
			SupplierInvoice:          data.SupplierInvoice,
			FiscalYear:               data.FiscalYear,
			TaxCode:                  data.TaxCode,
			DocumentCurrency:         data.DocumentCurrency,
			TaxAmount:                data.TaxAmount,
			TaxBaseAmountInTransCrcy: data.TaxBaseAmountInTransCrcy,
	}, nil
}

func ConvertToAccount(raw []byte, l *logger.Logger) (*Account, error) {
	pm := &responses.Account{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Account. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Account{
			SupplierInvoice:               data.SupplierInvoice,
			FiscalYear:                    data.FiscalYear,
			SupplierInvoiceItem:           data.SupplierInvoiceItem,
			CompanyCode:                   data.CompanyCode,
			CostCenter:                    data.CostCenter,
			ControllingArea:               data.ControllingArea,
			BusinessArea:                  data.BusinessArea,
			ProfitCenter:                  data.ProfitCenter,
			FunctionalArea:                data.FunctionalArea,
			GLAccount:                     data.GLAccount,
			SalesOrder:                    data.SalesOrder,
			SalesOrderItem:                data.SalesOrderItem,
			CostObject:                    data.CostObject,
			WBSElement:                    data.WBSElement,
			DocumentCurrency:              data.DocumentCurrency,
			SuplrInvcAcctAssignmentAmount: data.SuplrInvcAcctAssignmentAmount,
			TaxCode:                       data.TaxCode,
			WorkItem:                      data.WorkItem,
			MasterFixedAsset:              data.MasterFixedAsset,
			FixedAsset:                    data.FixedAsset,
			DebitCreditCode:               data.DebitCreditCode,
			InternalOrder:                 data.InternalOrder,
			ProjectNetwork:                data.ProjectNetwork,
			ProfitabilitySegment:          data.ProfitabilitySegment,
	}, nil
}

func ConvertToPurchaseOrder(raw []byte, l *logger.Logger) (*PurchaseOrder, error) {
	pm := &responses.PurchaseOrder{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchaseOrder. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchaseOrder{
			SupplierInvoice:                data.SupplierInvoice,
			FiscalYear:                     data.FiscalYear,
			SupplierInvoiceItem:            data.SupplierInvoiceItem,
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			Plant:                          data.Plant,
			TaxCode:                        data.TaxCode,
			DocumentCurrency:               data.DocumentCurrency,
			SupplierInvoiceItemAmount:      data.SupplierInvoiceItemAmount,
			PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
			QuantityInPurchaseOrderUnit:    data.QuantityInPurchaseOrderUnit,
			PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
			QtyInPurchaseOrderPriceUnit:    data.QtyInPurchaseOrderPriceUnit,
			SupplierInvoiceItemText:        data.SupplierInvoiceItemText,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
	}, nil
}
