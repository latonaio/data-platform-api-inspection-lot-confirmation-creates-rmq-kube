package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-lot-confirmation-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionLotConfirmation:           *&data.InspectionLotConfirmation,
		InspectionLotConfirmationHeaderText: data.InspectionLotConfirmationHeaderText,
	}
}
