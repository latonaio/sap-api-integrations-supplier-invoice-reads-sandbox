package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-supplier-invoice-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetSupplierInvoice(supplierInvoice, fiscalYear, purchaseOrder, purchaseOrderItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "Tax":
			func() {
				c.Tax(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "Account":
			func() {
				c.Account(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "PurchaseOrder":
			func() {
				c.PurchaseOrder(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementHeader("A_SupplierInvoice", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementHeader(api, supplierInvoice, fiscalYear string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, supplierInvoice, fiscalYear)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Tax(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementTax("A_SupplierInvoiceTax", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementTax(api, supplierInvoice, fiscalYear string) (*sap_api_output_formatter.Tax, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithTax(req, supplierInvoice, fiscalYear)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToTax(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Account(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementAccount("A_SuplrInvcItemAcctAssgmt", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementAccount(api, supplierInvoice, fiscalYear string) (*sap_api_output_formatter.Account, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithAccount(req, supplierInvoice, fiscalYear)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchaseOrder(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementPurchaseOrder("A_SuplrInvcItemPurOrdRef", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementPurchaseOrder(api, purchaseOrder, purchaseOrderItem string) (*sap_api_output_formatter.PurchaseOrder, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPurchaseOrder(req, purchaseOrder, purchaseOrderItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchaseOrder(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, supplierInvoice, fiscalYear string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithTax(req *http.Request, supplierInvoice, fiscalYear string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithAccount(req *http.Request, supplierInvoice, fiscalYear string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPurchaseOrder(req *http.Request, purchaseOrder, purchaseOrderItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem))
	req.URL.RawQuery = params.Encode()
}
