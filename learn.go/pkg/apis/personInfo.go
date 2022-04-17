package apis

//
//type PersonalInformation struct {
//	Nickname string  `json:"nickname"`
//	Password float64 `json:"Password"`
//	Account  int     `json:"Account"`
//}

func (*PersonalInformation) TableName() string {
	return "user_info"
}
