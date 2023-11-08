package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLotConfirmation           int     `json:"InspectionLotConfirmation"`
	InspectionLotConfirmationHeaderText *string `json:"InspectionLotConfirmationHeaderText"`
}

type InspectionUpdates struct {
	InspectionLotConfirmation               int     `json:"InspectionLotConfirmation"`
	Inspection                              int     `json:"Inspection"`
	InspectionLotConfirmationInspectionText *string `json:"InspectionLotConfirmationInspectionText"`
}
