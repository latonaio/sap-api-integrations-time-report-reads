package responses

type TimeReportCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                       string `json:"ObjectID"`
			Description                    string `json:"Description"`
			LanguageCode                   string `json:"languageCode"`
			EmployeeUUID                   string `json:"EmployeeUUID"`
			EndDate                        string `json:"EndDate"`
			ID                             string `json:"ID"`
			InformationLifeCycleStatusCode string `json:"InformationLifeCycleStatusCode"`
			RejectionReason                string `json:"RejectionReason"`
			LanguageCode1                  string `json:"languageCode1"`
			ReportName                     string `json:"ReportName"`
			LanguageCode2                  string `json:"languageCode2"`
			StartDate                      string `json:"StartDate"`
			Status                         string `json:"Status"`
			CreationDateTime               string `json:"CreationDateTime"`
			CreationIdentityUUID           string `json:"CreationIdentityUUID"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			LastChangeIdentityUUID         string `json:"LastChangeIdentityUUID"`
			TotalDuration                  string `json:"TotalDuration"`
			EntityLastChangedOn            string `json:"EntityLastChangedOn"`
			ETag                           string `json:"ETag"`
			CodTimeEntry                   struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CodTimeEntry"`
			CodTimeReportAttachmentFolder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CodTimeReportAttachmentFolder"`
			CodTimeReportParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CodTimeReportParty"`
		} `json:"results"`
	} `json:"d"`
}
