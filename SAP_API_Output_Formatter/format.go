package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-return-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		CustomerReturn:                data.CustomerReturn,
		CustomerReturnType:            data.CustomerReturnType,
		SalesOrganization:             data.SalesOrganization,
		DistributionChannel:           data.DistributionChannel,
		OrganizationDivision:          data.OrganizationDivision,
		SalesGroup:                    data.SalesGroup,
		SalesOffice:                   data.SalesOffice,
		SalesDistrict:                 data.SalesDistrict,
		SoldToParty:                   data.SoldToParty,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		SenderBusinessSystemName:      data.SenderBusinessSystemName,
		LastChangeDateTime:            data.LastChangeDateTime,
		PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderType:     data.CustomerPurchaseOrderType,
		CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
		CustomerReturnDate:            data.CustomerReturnDate,
		TotalNetAmount:                data.TotalNetAmount,
		TransactionCurrency:           data.TransactionCurrency,
		SDDocumentReason:              data.SDDocumentReason,
		PricingDate:                   data.PricingDate,
		RequestedDeliveryDate:         data.RequestedDeliveryDate,
		ShippingType:                  data.ShippingType,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		DeliveryBlockReason:           data.DeliveryBlockReason,
		IncotermsClassification:       data.IncotermsClassification,
		CustomerPaymentTerms:          data.CustomerPaymentTerms,
		PaymentMethod:                 data.PaymentMethod,
		CustomerTaxClassification1:    data.CustomerTaxClassification1,
		RetsMgmtProcess:               data.RetsMgmtProcess,
		ReferenceSDDocument:           data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:   data.ReferenceSDDocumentCategory,
		CustomerReturnApprovalReason:  data.CustomerReturnApprovalReason,
		SalesDocApprovalStatus:        data.SalesDocApprovalStatus,
		RetsMgmtLogProcgStatus:        data.RetsMgmtLogProcgStatus,
		RetsMgmtCompnProcgStatus:      data.RetsMgmtCompnProcgStatus,
		RetsMgmtProcessingStatus:      data.RetsMgmtProcessingStatus,
		OverallSDProcessStatus:        data.OverallSDProcessStatus,
		TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
		OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		CustomerReturn:                 data.CustomerReturn,
		CustomerReturnItem:             data.CustomerReturnItem,
		HigherLevelItem:                data.HigherLevelItem,
		CustomerReturnItemCategory:     data.CustomerReturnItemCategory,
		CustomerReturnItemText:         data.CustomerReturnItemText,
		PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
		Material:                       data.Material,
		MaterialByCustomer:             data.MaterialByCustomer,
		RequestedQuantity:              data.RequestedQuantity,
		RequestedQuantityUnit:          data.RequestedQuantityUnit,
		ItemGrossWeight:                data.ItemGrossWeight,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		ItemVolume:                     data.ItemVolume,
		ItemVolumeUnit:                 data.ItemVolumeUnit,
		TransactionCurrency:            data.TransactionCurrency,
		NetAmount:                      data.NetAmount,
		MaterialGroup:                  data.MaterialGroup,
		Batch:                          data.Batch,
		ProductionPlant:                data.ProductionPlant,
		StorageLocation:                data.StorageLocation,
		ShippingPoint:                  data.ShippingPoint,
		ShippingType:                   data.ShippingType,
		DeliveryPriority:               data.DeliveryPriority,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		ProductTaxClassification1:      data.ProductTaxClassification1,
		SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
		ItemBillingBlockReason:         data.ItemBillingBlockReason,
		ProfitCenter:                   data.ProfitCenter,
		RetsMgmtProcess:                data.RetsMgmtProcess,
		RetsMgmtProcessItem:            data.RetsMgmtProcessItem,
		ReturnReason:                   data.ReturnReason,
		RetsMgmtProcessingBlock:        data.RetsMgmtProcessingBlock,
		CustRetItmFollowUpActivity:     data.CustRetItmFollowUpActivity,
		ReturnsMaterialHasBeenReceived: data.ReturnsMaterialHasBeenReceived,
		ReturnsRefundType:              data.ReturnsRefundType,
		ReturnsRefundProcgMode:         data.ReturnsRefundProcgMode,
		ReturnsRefundExtent:            data.ReturnsRefundExtent,
		PrelimRefundIsDetermined:       data.PrelimRefundIsDetermined,
		ReturnsRefundRjcnReason:        data.ReturnsRefundRjcnReason,
		ReplacementMaterial:            data.ReplacementMaterial,
		ReplacementMaterialQuantity:    data.ReplacementMaterialQuantity,
		ReplacementMaterialQtyUnit:     data.ReplacementMaterialQtyUnit,
		ReplacementMaterialIsRequested: data.ReplacementMaterialIsRequested,
		ReplacementMatlSupplyingPlant:  data.ReplacementMatlSupplyingPlant,
		NextPlantForFollowUpActivity:   data.NextPlantForFollowUpActivity,
		ReturnsTransshipmentPlant:      data.ReturnsTransshipmentPlant,
		Supplier:                       data.Supplier,
		SupplierRetMatlAuthzn:          data.SupplierRetMatlAuthzn,
		SuplrRetMatlAuthznIsRequired:   data.SuplrRetMatlAuthznIsRequired,
		CustomerRetMatlInspResultCode:  data.CustomerRetMatlInspResultCode,
		NextFllwUpActivityForMatlInsp:  data.NextFllwUpActivityForMatlInsp,
		RetMatlInspResultCode:          data.RetMatlInspResultCode,
		ProductIsInspectedAtCustSite:   data.ProductIsInspectedAtCustSite,
		CustRetMatlAuthzn:              data.CustRetMatlAuthzn,
		CRMLogicalSystem:               data.CRMLogicalSystem,
		CRMObjectUUID:                  data.CRMObjectUUID,
		CRMObjectID:                    data.CRMObjectID,
		CRMObjectType:                  data.CRMObjectType,
		RetsMgmtItmLogProcgStatus:      data.RetsMgmtItmLogProcgStatus,
		RetsMgmtItmCompnProcgStatus:    data.RetsMgmtItmCompnProcgStatus,
		RetsMgmtItmProcgStatus:         data.RetsMgmtItmProcgStatus,
		ReturnsDocumentStatus:          data.ReturnsDocumentStatus,
		ReturnsDocumentApprovalStatus:  data.ReturnsDocumentApprovalStatus,
		SDProcessStatus:                data.SDProcessStatus,
		ReferenceSDDocument:            data.ReferenceSDDocument,
		ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
		ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
		SDDocumentRejectionStatus:      data.SDDocumentRejectionStatus,
	}

	return item, nil
}
