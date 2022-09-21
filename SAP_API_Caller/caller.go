package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchase-requisition-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseRequisition(purchaseRequisition, purchaseRequisitionItem, purchasingDocument, purchasingDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseRequisition)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "ItemDeliveryAddress":
			func() {
				c.ItemDeliveryAddress(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "ItemAccount":
			func() {
				c.ItemAccount(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "PurchasingDocument":
			func() {
				c.PurchasingDocument(purchasingDocument, purchasingDocumentItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseRequisition string) {
	headerData, err := c.callPurchaseRequisitionSrvAPIRequirementHeader("A_PurchaseRequisitionHeader", purchaseRequisition)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(headerData)
	}

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(itemData)
	}

	itemDeliveryAddressData, err := c.callToItemDeliveryAddress(itemData[0].ToItemDeliveryAddress)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(itemDeliveryAddressData)
	}

	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(itemAccountData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementHeader(api, purchaseRequisition string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, purchaseRequisition)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDeliveryAddress(url string) (*sap_api_output_formatter.ToItemDeliveryAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDeliveryAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAccount(url string) ([]sap_api_output_formatter.ToItemAccount, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseRequisition, purchaseRequisitionItem string) {
	itemData, err := c.callPurchaseRequisitionSrvAPIRequirementItem("A_PurchaseRequisitionItem", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(itemData)
    }
	
	itemDeliveryAddressData, err := c.callToItemDeliveryAddress(itemData[0].ToItemDeliveryAddress)
	if err != nil {
		c.log.Error(err)
	} else {
	    c.log.Info(itemDeliveryAddressData)
    }
	
	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
    } else {
	    c.log.Info(itemAccountData)
    }
	return
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementItem(api, purchaseRequisition, purchaseRequisitionItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, purchaseRequisition, purchaseRequisitionItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemDeliveryAddress(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementItemDeliveryAddress("A_PurReqAddDelivery", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementItemDeliveryAddress(api, purchaseRequisition, purchaseRequisitionItem string) ([]sap_api_output_formatter.ItemDeliveryAddress, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItemDeliveryAddress(map[string]string{}, purchaseRequisition, purchaseRequisitionItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemDeliveryAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemAccount(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementItemAccount("A_PurReqnAcctAssgmt", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementItemAccount(api, purchaseRequisition, purchaseRequisitionItem string) ([]sap_api_output_formatter.ItemAccount, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItemAccount(map[string]string{}, purchaseRequisition, purchaseRequisitionItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchasingDocument(purchasingDocument, purchasingDocumentItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementPurchasingDocument("A_PurchaseRequisitionItem", purchasingDocument, purchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementPurchasingDocument(api, purchasingDocument, purchasingDocumentItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")

	param := c.getQueryWithPurchasingDocument(map[string]string{}, purchasingDocument, purchasingDocumentItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, purchaseRequisition string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseRequisition eq '%s'", purchaseRequisition)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, purchaseRequisition, purchaseRequisitionItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemAccount(params map[string]string, purchaseRequisition, purchaseRequisitionItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemDeliveryAddress(params map[string]string, purchaseRequisition, purchaseRequisitionItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem)
	return params
}

func (c *SAPAPICaller) getQueryWithPurchasingDocument(params map[string]string, purchasingDocument, purchasingDocumentItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", purchasingDocument, purchasingDocumentItem)
	return params
}
