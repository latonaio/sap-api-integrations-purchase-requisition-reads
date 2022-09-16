package sap_api_output_formatter

type PurchaseRequisition struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	APISchema           string `json:"api_schema"`
	PurchaseRequisition string `json:"purchase_requisition"`
	Deleted             bool   `json:"deleted"`
}

type Header struct {
	PurchaseRequisition     string `json:"PurchaseRequisition"`
	PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
	SourceDetermination     bool   `json:"SourceDetermination"`
	ToItem                  string `json:"to_PurchaseReqnItem"`
}

type Item struct {
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
	ToItemDeliveryAddress          string `json:"to_PurchaseReqnDeliveryAddress"`
	ToItemAccount                  string `json:"to_PurchaseReqnAcctAssgmt"`
}

type ItemDeliveryAddress struct {
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
}

type ItemAccount struct {
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
}

type ToItem struct {
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
	ToItemDeliveryAddress          string `json:"to_PurchaseReqnDeliveryAddress"`
	ToItemAccount                  string `json:"to_PurchaseReqnAcctAssgmt"`
}

type ToItemDeliveryAddress struct {
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
}

type ToItemAccount struct {
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
}

