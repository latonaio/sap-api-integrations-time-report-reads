package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-time-report-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToTimeReportCollection(raw []byte, l *logger.Logger) ([]TimeReportCollection, error) {
	pm := &responses.TimeReportCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to TimeReportCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	timeReportCollection := make([]TimeReportCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		timeReportCollection = append(timeReportCollection, TimeReportCollection{
			ObjectID:                       data.ObjectID,
			Description:                    data.Description,
			LanguageCode:                   data.LanguageCode,
			EmployeeUUID:                   data.EmployeeUUID,
			EndDate:                        data.EndDate,
			ID:                             data.ID,
			InformationLifeCycleStatusCode: data.InformationLifeCycleStatusCode,
			RejectionReason:                data.RejectionReason,
			LanguageCode1:                  data.LanguageCode1,
			ReportName:                     data.ReportName,
			LanguageCode2:                  data.LanguageCode2,
			StartDate:                      data.StartDate,
			Status:                         data.Status,
			CreationDateTime:               data.CreationDateTime,
			CreationIdentityUUID:           data.CreationIdentityUUID,
			LastChangeDateTime:             data.LastChangeDateTime,
			LastChangeIdentityUUID:         data.LastChangeIdentityUUID,
			TotalDuration:                  data.TotalDuration,
			EntityLastChangedOn:            data.EntityLastChangedOn,
			ETag:                           data.ETag,
		})
	}

	return timeReportCollection, nil
}

func ConvertToTimeReportPartyCollection(raw []byte, l *logger.Logger) ([]TimeReportPartyCollection, error) {
	pm := &responses.TimeReportPartyCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to TimeReportPartyCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	timeReportPartyCollection := make([]TimeReportPartyCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		timeReportPartyCollection = append(timeReportPartyCollection, TimeReportPartyCollection{
			ObjectID:                  data.ObjectID,
			ParentObjectID:            data.ParentObjectID,
			DeterminationMethodCode:   data.DeterminationMethodCode,
			AddressHostTypeCode:       data.AddressHostTypeCode,
			MainIndicator:             data.MainIndicator,
			PartyID:                   data.PartyID,
			PartyTypeCode:             data.PartyTypeCode,
			RoleCategoryCode:          data.RoleCategoryCode,
			RoleCode:                  data.RoleCode,
			ETag:                      data.ETag,
		})
	}

	return timeReportPartyCollection, nil
}
