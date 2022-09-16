package responses

type ToItemAccount struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseRequisition          string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem      string `json:"PurchaseRequisitionItem"`
			PurchaseReqnAcctAssgmtNumber string `json:"PurchaseReqnAcctAssgmtNumber"`
			CostCenter                   string `json:"CostCenter"`
			MasterFixedAsset             string `json:"MasterFixedAsset"`
			FixedAsset                   string `json:"FixedAsset"`
			ProjectNetwork               string `json:"ProjectNetwork"`
			CostElement                  string `json:"CostElement"`
			CostObject                   string `json:"CostObject"`
			GLAccount                    string `json:"GLAccount"`
			BusinessArea                 string `json:"BusinessArea"`
			SalesOrder                   string `json:"SalesOrder"`
			SalesOrderItem               string `json:"SalesOrderItem"`
			SalesOrderScheduleLine       string `json:"SalesOrderScheduleLine"`
			OrderID                      string `json:"OrderID"`
			UnloadingPointName           string `json:"UnloadingPointName"`
			ControllingArea              string `json:"ControllingArea"`
			ProfitabilitySegment         string `json:"ProfitabilitySegment"`
			ProfitCenter                 string `json:"ProfitCenter"`
			FunctionalArea               string `json:"FunctionalArea"`
			GoodsRecipientName           string `json:"GoodsRecipientName"`
			CostCtrActivityType          string `json:"CostCtrActivityType"`
			WBSElement                   string `json:"WBSElement"`
			IsDeleted                    string `json:"IsDeleted"`
		} `json:"results"`
	} `json:"d"`
}
