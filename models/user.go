package models

type AssignableUser struct {
	Self         string `json:"self"`
	AccountId    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls   struct {
		X16 string `json:"16x16"`
		X24 string `json:"24x24"`
		X32 string `json:"32x32"`
		X48 string `json:"48x48"`
	} `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`
}
