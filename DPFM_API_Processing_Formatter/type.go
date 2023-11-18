package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLot           int     `json:"InspectionLot"`
	Operations              int     `json:"Operations"`
	OperationsItem          int     `json:"OperationsItem"`
	OperationID             int     `json:"OperationID"`
	ConfirmationCountingID  int     `json:"ConfirmationCountingID"`
	ConfirmationText        *string `json:"ConfirmationText"`
	OperationVarianceReason *string `json:"OperationVarianceReason"`
}
