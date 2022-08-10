package responses

type TimeReportPartyCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                string `json:"ObjectID"`
			ParentObjectID          string `json:"ParentObjectID"`
			DeterminationMethodCode string `json:"DeterminationMethodCode"`
			AddressHostTypeCode     string `json:"AddressHostTypeCode"`
			MainIndicator           bool   `json:"MainIndicator"`
			PartyID                 string `json:"PartyID"`
			PartyTypeCode           string `json:"PartyTypeCode"`
			RoleCategoryCode        string `json:"RoleCategoryCode"`
			RoleCode                string `json:"RoleCode"`
			ETag                    string `json:"ETag"`
			CodTimeReport           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CodTimeReport"`
		} `json:"results"`
	} `json:"d"`
}
