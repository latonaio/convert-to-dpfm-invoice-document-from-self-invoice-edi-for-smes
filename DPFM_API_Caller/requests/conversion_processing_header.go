package requests

type ConversionProcessingHeader struct {
	ConvertingInvoiceDocument *string `json:"ConvertingInvoiceDocument"`
	ConvertedInvoiceDocument  *int    `json:"ConvertedInvoiceDocument"`
	ConvertingBillToParty     *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty      *int    `json:"ConvertedBillToParty"`
	ConvertingBillFromParty   *string `json:"ConvertingBillFromParty"`
	ConvertedBillFromParty    *int    `json:"ConvertedBillFromParty"`
	ConvertingPayer           *string `json:"ConvertingPayer"`
	ConvertedPayer            *int    `json:"ConvertedPayer"`
	ConvertingPayee           *string `json:"ConvertingPayee"`
	ConvertedPayee            *int    `json:"ConvertedPayee"`
	ConvertingPaymentMethod   *string `json:"ConvertingPaymentMethod"`
	ConvertedPaymentMethod    *string `json:"ConvertedPaymentMethod"`
}
