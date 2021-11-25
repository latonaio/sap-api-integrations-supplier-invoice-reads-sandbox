package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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

func (c *SAPAPICaller) AsyncGetSupplierInvoice(SupplierInvoice, FiscalYear, PurchaseOrder, PurchaseOrderItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		c.Header(SupplierInvoice, FiscalYear)
		wg.Done()
	}()

	go func() {
		c.Tax(SupplierInvoice, FiscalYear)
		wg.Done()
	}()
	
	go func() {
		c.PurchaseOrder(PurchaseOrder, PurchaseOrderItem)
		wg.Done()
	}()
	wg.Wait()
	go func() {
		c.Account(SupplierInvoice, FiscalYear)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(SupplierInvoice, FiscalYear string) {
	res, err := c.callSupplierInvoiceSrvAPIRequirementHeader("A_SupplierInvoice", SupplierInvoice, FiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Tax(SupplierInvoice, FiscalYear string) {
	res, err := c.callSupplierInvoiceSrvAPIRequirementTax("A_SupplierInvoice(SupplierInvoice='{SupplierInvoice}',FiscalYear='{FiscalYear}')/to_SupplierInvoiceItemGLAcct", SupplierInvoice, FiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

func (c *SAPAPICaller) PurchaseOrder(PurchaseOrder, PurchaseOrderItem string) {
	res, err := c.callSupplierInvoiceSrvAPIRequirementPurchaseOrder("A_SuplrInvcItemPurOrdRef", PurchaseOrder, PurchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)
	
func (c *SAPAPICaller) Account(SupplierInvoice, FiscalYear string) {
	res, err := c.callSupplierInvoiceSrvAPIRequirementAccount("A_SuplrInvcItemAcctAssgmt", PurchaseOrder, PurchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)


func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementHeader(api, SupplierInvoice, FiscalYear string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "SupplierInvoice, FiscalYear")
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", SupplierInvoice, FiscalYear))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementTax(api, SupplierInvoice, FiscalYear string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "SupplierInvoice, FiscalYear")
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq 'RE'", SupplierInvoice, FiscalYear))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}


func (c *SAPAPICaller) callSupplierInvoiceSrvAPIPurchaseOrder(api, PurchaseOrder, PurchaseOrderItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseOrder, PurchaseOrderItem")
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", PurchaseOrder, PurchaseOrderItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementAccount(api, SupplierInvoice, FiscalYear string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "SupplierInvoice, FiscalYear")
	params.Add("$filter", fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq 'RE'", SupplierInvoice, FiscalYear))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

