package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseRequisition            string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
			PurchasingDocument             string `json:"PurchasingDocument"`
			PurchasingDocumentItem         string `json:"PurchasingDocumentItem"`
			PurReqnReleaseStatus           string `json:"PurReqnReleaseStatus"`
			PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
			PurchaseRequisitionItemText    string `json:"PurchaseRequisitionItemText"`
			AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
			Material                       string `json:"Material"`
			MaterialGroup                  string `json:"MaterialGroup"`
			PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
			RequestedQuantity              string `json:"RequestedQuantity"`
			BaseUnit                       string `json:"BaseUnit"`
			PurchaseRequisitionPrice       string `json:"PurchaseRequisitionPrice"`
			PurReqnPriceQuantity           string `json:"PurReqnPriceQuantity"`
			MaterialGoodsReceiptDuration   string `json:"MaterialGoodsReceiptDuration"`
			ReleaseCode                    string `json:"ReleaseCode"`
			PurchaseRequisitionReleaseDate string `json:"PurchaseRequisitionReleaseDate"`
			PurchasingOrganization         string `json:"PurchasingOrganization"`
			PurchasingGroup                string `json:"PurchasingGroup"`
			Plant                          string `json:"Plant"`
			CompanyCode                    string `json:"CompanyCode"`
			SourceOfSupplyIsAssigned       bool   `json:"SourceOfSupplyIsAssigned"`
			SupplyingPlant                 string `json:"SupplyingPlant"`
			OrderedQuantity                string `json:"OrderedQuantity"`
			DeliveryDate                   string `json:"DeliveryDate"`
			CreationDate                   string `json:"CreationDate"`
			ProcessingStatus               string `json:"ProcessingStatus"`
			ExternalApprovalStatus         string `json:"ExternalApprovalStatus"`
			PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
			Supplier                       string `json:"Supplier"`
			FixedSupplier                  string `json:"FixedSupplier"`
			RequisitionerName              string `json:"RequisitionerName"`
			PurReqnItemCurrency            string `json:"PurReqnItemCurrency"`
			MaterialPlannedDeliveryDurn    string `json:"MaterialPlannedDeliveryDurn"`
			StorageLocation                string `json:"StorageLocation"`
			PurReqnSourceOfSupplyType      string `json:"PurReqnSourceOfSupplyType"`
			ConsumptionPosting             string `json:"ConsumptionPosting"`
			PurReqnOrigin                  string `json:"PurReqnOrigin"`
			IsPurReqnBlocked               string `json:"IsPurReqnBlocked"`
			PurchaseRequisitionStatus      string `json:"PurchaseRequisitionStatus"`
			Batch                          string `json:"Batch"`
			GoodsReceiptIsExpected         bool   `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool   `json:"GoodsReceiptIsNonValuated"`
			RequirementTracking            string `json:"RequirementTracking"`
			MRPController                  string `json:"MRPController"`
			Reservation                    string `json:"Reservation"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			IsDeleted                      string `json:"IsDeleted"`
			ToItemDeliveryAddress struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurchaseReqnDeliveryAddress"`
			ToItemAccount struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurchaseReqnAcctAssgmt"`
		} `json:"results"`
	} `json:"d"`
}
