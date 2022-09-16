package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseRequisition     string `json:"PurchaseRequisition"`
			PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
			SourceDetermination     bool   `json:"SourceDetermination"`
			ToItem                  struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurchaseReqnItem"`
		} `json:"results"`
	} `json:"d"`
}
