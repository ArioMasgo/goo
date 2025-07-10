package health

import "github.com/asaskevich/govalidator"

type JWTToken struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestLogin struct {
	Username  string `json:"username" db:"username" valid:"required~Username is blank"`
	Password  string `json:"password" db:"password" valid:"required~Password is blank"`
	IPAddress string `json:"ip_address,omitempty" db:"ip_address" valid:"required~IP Address no detected,ipv4~IP no valid"`
}

func (m *RequestLogin) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
