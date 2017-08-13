package common

type Address struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	Address3  string `json:"address3,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
}
