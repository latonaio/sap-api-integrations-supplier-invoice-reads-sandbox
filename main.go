package main

import (
	sap_api_caller "sap-api-integrations-supplier-invoice-reads/SAP_API_Caller"
	"sap-api-integrations-supplier-invoice-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Supplier_Invoice_sample3.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetSupplierInvoice(
		inoutSDC.SupplierInvoice.SupplierInvoice,
		inoutSDC.SupplierInvoice.FiscalYear,
		inoutSDC.SupplierInvoice.PurchaseOrder.PurchaseOrder,
		inoutSDC.SupplierInvoice.PurchaseOrder.PurchaseOrderItem,
	)
}
