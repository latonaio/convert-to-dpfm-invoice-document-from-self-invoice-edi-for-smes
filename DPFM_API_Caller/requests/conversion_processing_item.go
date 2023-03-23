package requests

type ConversionProcessingItem struct {
	ConvertingInvoiceDocumentItem          *string `json:"ConvertingInvoiceDocumentItem"`
	ConvertedInvoiceDocumentItem           *int    `json:"ConvertedInvoiceDocumentItem"`
	ConvertingProduct                      *string `json:"ConvertingProduct"`
	ConvertedProduct                       *string `json:"ConvertedProduct"`
	ConvertingBuyer                        *string `json:"ConvertingBuyer"`
	ConvertedBuyer                         *int    `json:"ConvertedBuyer"`
	ConvertingSeller                       *string `json:"ConvertingSeller"`
	ConvertedSeller                        *int    `json:"ConvertedSeller"`
	ConvertingDeliverToParty               *string `json:"ConvertingDeliverToParty"`
	ConvertedDeliverToParty                *int    `json:"ConvertedDeliverToParty"`
	ConvertingDeliverFromParty             *string `json:"ConvertingDeliverFromParty"`
	ConvertedDeliverFromParty              *int    `json:"ConvertedDeliverFromParty"`
	ConvertingTransactionTaxClassification *string `json:"ConvertingTransactionTaxClassification"`
	ConvertedTransactionTaxClassification  *string `json:"ConvertedTransactionTaxClassification"`
	ConvertingProject                      *string `json:"ConvertingProject"`
	ConvertedProject                       *string `json:"ConvertedProject"`
	ConvertingOrderID                      *string `json:"ConvertingOrderID"`
	ConvertedOrderID                       *int    `json:"ConvertedOrderID"`
	ConvertingOrderItem                    *string `json:"ConvertingOrderItem"`
	ConvertedOrderItem                     *int    `json:"ConvertedOrderItem"`
	ConvertingOriginDocument               *string `json:"ConvertingOriginDocument"`
	ConvertedOriginDocument                *int    `json:"ConvertedOriginDocument"`
	ConvertingOriginDocumentItem           *string `json:"ConvertingOriginDocumentItem"`
	ConvertedOriginDocumentItem            *int    `json:"ConvertedOriginDocumentItem"`
	ConvertingReferenceDocument            *string `json:"ConvertingReferenceDocument"`
	ConvertedReferenceDocument             *int    `json:"ConvertedReferenceDocument"`
	ConvertingReferenceDocumentItem        *string `json:"ConvertingReferenceDocumentItem"`
	ConvertedReferenceDocumentItem         *int    `json:"ConvertedReferenceDocumentItem"`
}
