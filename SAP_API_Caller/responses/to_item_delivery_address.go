package responses

type ToItemDeliveryAddress struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
			PurchaseRequisition     string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem string `json:"PurchaseRequisitionItem"`
			AddressID               string `json:"AddressID"`
			Country                 string `json:"Country"`
			Region                  string `json:"Region"`
			StreetName              string `json:"StreetName"`
			CityName                string `json:"CityName"`
			CorrespondenceLanguage  string `json:"CorrespondenceLanguage"`
			FaxNumber               string `json:"FaxNumber"`
			PhoneNumber             string `json:"PhoneNumber"`
	} `json:"d"`
}
