package sap_api_output_formatter

type TimeReport struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	APISchema      string `json:"api_schema"`
	TimeReportCode string `json:"time_report_code"`
	Deleted        bool   `json:"deleted"`
}

type TimeReportCollection struct {
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
}

type TimeReportPartyCollection struct {
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
}
