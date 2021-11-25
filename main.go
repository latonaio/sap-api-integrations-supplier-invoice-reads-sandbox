package main

import (
    sap_api_caller "ap-api-integrations-supplier-invoice-reads/SAP_API_Caller"
    "ap-api-integrations-supplier-invoice-reads/file_reader"

    "github.com/latonaio/golang-logging-library/logger"
)

func main() {
    l := logger.NewLogger()
    fr := file_reader.NewFileReader()
    inoutSDC := fr.ReadSDC("./Inputs//SDC_Supplier_Invoice_sample.json")
    caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
    )

    caller.AsyncGetSupplierInvoice(
 		inoutSDC.SupplierInvoice.FiscalYear,
        inoutSDC.PurchaseOrder.PurchaseOrderItem,
    )
}