package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-requisition-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionType:        data.PurchaseRequisitionType,
		SourceDetermination:            data.SourceDetermination,
        ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		ReleaseCode:                    data.ReleaseCode,
		PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		Plant:                          data.Plant,
		CompanyCode:                    data.CompanyCode,
		SourceOfSupplyIsAssigned:       data.SourceOfSupplyIsAssigned,
		SupplyingPlant:                 data.SupplyingPlant,
		OrderedQuantity:                data.OrderedQuantity,
		DeliveryDate:                   data.DeliveryDate,
		CreationDate:                   data.CreationDate,
		ProcessingStatus:               data.ProcessingStatus,
		ExternalApprovalStatus:         data.ExternalApprovalStatus,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		Supplier:                       data.Supplier,
		FixedSupplier:                  data.FixedSupplier,
		RequisitionerName:              data.RequisitionerName,
		PurReqnItemCurrency:            data.PurReqnItemCurrency,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		StorageLocation:                data.StorageLocation,
		PurReqnSourceOfSupplyType:      data.PurReqnSourceOfSupplyType,
		ConsumptionPosting:             data.ConsumptionPosting,
		PurReqnOrigin:                  data.PurReqnOrigin,
		IsPurReqnBlocked:               data.IsPurReqnBlocked,
		PurchaseRequisitionStatus:      data.PurchaseRequisitionStatus,
		Batch:                          data.Batch,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		RequirementTracking:            data.RequirementTracking,
		MRPController:                  data.MRPController,
		Reservation:                    data.Reservation,
		LastChangeDateTime:             data.LastChangeDateTime,
		IsDeleted:                      data.IsDeleted,
        ToItemDeliveryAddress:          data.ToItemDeliveryAddress.Deferred.URI,
        ToItemAccount:                  data.ToItemAccount.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToItemDeliveryAddress(raw []byte, l *logger.Logger) ([]ItemDeliveryAddress, error) {
	pm := &responses.ItemDeliveryAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemDeliveryAddress. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	itemDeliveryAddress := make([]ItemDeliveryAddress, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemDeliveryAddress = append(itemDeliveryAddress, ItemDeliveryAddress{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		AddressID:                      data.AddressID,
		Country:                        data.Country,
		Region:                         data.Region,
		StreetName:                     data.StreetName,
		CityName:                       data.CityName,
		CorrespondenceLanguage:         data.CorrespondenceLanguage,
		FaxNumber:                      data.FaxNumber,
		PhoneNumber:                    data.PhoneNumber,
		})
	}

	return itemDeliveryAddress, nil
}

func ConvertToItemAccount(raw []byte, l *logger.Logger) ([]ItemAccount, error) {
	pm := &responses.ItemAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	itemAccount := make([]ItemAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemAccount = append(itemAccount, ItemAccount{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
        PurchaseReqnAcctAssgmtNumber:   data.PurchaseReqnAcctAssgmtNumber,
		CostCenter:                     data.CostCenter,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		ProjectNetwork:                 data.ProjectNetwork,
		CostElement:                    data.CostElement,
		CostObject:                     data.CostObject,
		GLAccount:                      data.GLAccount,
		BusinessArea:                   data.BusinessArea,
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
		OrderID:                        data.OrderID,
		UnloadingPointName:             data.UnloadingPointName,
		ControllingArea:                data.ControllingArea,
		ProfitabilitySegment:           data.ProfitabilitySegment,
		ProfitCenter:                   data.ProfitCenter,
		FunctionalArea:                 data.FunctionalArea,
		GoodsRecipientName:             data.GoodsRecipientName,
		CostCtrActivityType:            data.CostCtrActivityType,
		WBSElement:                     data.WBSElement,
		IsDeleted:                      data.IsDeleted,
		})
	}

	return itemAccount, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		ReleaseCode:                    data.ReleaseCode,
		PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		Plant:                          data.Plant,
		CompanyCode:                    data.CompanyCode,
		SourceOfSupplyIsAssigned:       data.SourceOfSupplyIsAssigned,
		SupplyingPlant:                 data.SupplyingPlant,
		OrderedQuantity:                data.OrderedQuantity,
		DeliveryDate:                   data.DeliveryDate,
		CreationDate:                   data.CreationDate,
		ProcessingStatus:               data.ProcessingStatus,
		ExternalApprovalStatus:         data.ExternalApprovalStatus,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		Supplier:                       data.Supplier,
		FixedSupplier:                  data.FixedSupplier,
		RequisitionerName:              data.RequisitionerName,
		PurReqnItemCurrency:            data.PurReqnItemCurrency,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		StorageLocation:                data.StorageLocation,
		PurReqnSourceOfSupplyType:      data.PurReqnSourceOfSupplyType,
		ConsumptionPosting:             data.ConsumptionPosting,
		PurReqnOrigin:                  data.PurReqnOrigin,
		IsPurReqnBlocked:               data.IsPurReqnBlocked,
		PurchaseRequisitionStatus:      data.PurchaseRequisitionStatus,
		Batch:                          data.Batch,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		RequirementTracking:            data.RequirementTracking,
		MRPController:                  data.MRPController,
		Reservation:                    data.Reservation,
		LastChangeDateTime:             data.LastChangeDateTime,
		IsDeleted:                      data.IsDeleted,
        ToItemDeliveryAddress:          data.ToItemDeliveryAddress.Deferred.URI,
        ToItemAccount:                  data.ToItemAccount.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemDeliveryAddress(raw []byte, l *logger.Logger) (*ToItemDeliveryAddress, error) {
	pm := &responses.ToItemDeliveryAddress{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemDeliveryAddress. unmarshal error: %w", err)
	}
	
	return &ToItemDeliveryAddress{
		PurchaseRequisition:            pm.D.PurchaseRequisition,
		PurchaseRequisitionItem:        pm.D.PurchaseRequisitionItem,
		AddressID:                      pm.D.AddressID,
		Country:                        pm.D.Country,
		Region:                         pm.D.Region,
		StreetName:                     pm.D.StreetName,
		CityName:                       pm.D.CityName,
		CorrespondenceLanguage:         pm.D.CorrespondenceLanguage,
		FaxNumber:                      pm.D.FaxNumber,
		PhoneNumber:                    pm.D.PhoneNumber,
	}, nil
}

func ConvertToToItemAccount(raw []byte, l *logger.Logger) ([]ToItemAccount, error) {
	pm := &responses.ToItemAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemAccount := make([]ToItemAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemAccount = append(toItemAccount, ToItemAccount{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
        PurchaseReqnAcctAssgmtNumber:   data.PurchaseReqnAcctAssgmtNumber,
		CostCenter:                     data.CostCenter,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		ProjectNetwork:                 data.ProjectNetwork,
		CostElement:                    data.CostElement,
		CostObject:                     data.CostObject,
		GLAccount:                      data.GLAccount,
		BusinessArea:                   data.BusinessArea,
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
		OrderID:                        data.OrderID,
		UnloadingPointName:             data.UnloadingPointName,
		ControllingArea:                data.ControllingArea,
		ProfitabilitySegment:           data.ProfitabilitySegment,
		ProfitCenter:                   data.ProfitCenter,
		FunctionalArea:                 data.FunctionalArea,
		GoodsRecipientName:             data.GoodsRecipientName,
		CostCtrActivityType:            data.CostCtrActivityType,
		WBSElement:                     data.WBSElement,
		IsDeleted:                      data.IsDeleted,
		})
	}

	return toItemAccount, nil
}
