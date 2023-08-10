package handler

import (
	"sawitpro/generated"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBody(t *testing.T) {
	tests := []struct {
		name           string
		paramName      string
		paramPhone     string
		paramPass      string
		expectedErrStr string
	}{
		{
			name:           "Valid Body",
			paramPhone:     "+62345678990",
			paramName:      "John Doe",
			paramPass:      "Pass123!",
			expectedErrStr: "",
		},
		{
			name:           "inValid Phone legnth",
			paramPhone:     "345678990",
			paramName:      "John Doe",
			paramPass:      "Pass123!",
			expectedErrStr: "Invalid phone number:phone number must be between 10 and 13 characters",
		},
		{
			name:           "inValid Phone area",
			paramPhone:     "34567899000",
			paramName:      "John Doe",
			paramPass:      "Pass123!",
			expectedErrStr: "Invalid phone number:phone number must start with +62",
		},
		{
			name:           "inValid Name",
			paramPhone:     "+62345678990",
			paramName:      "oe",
			paramPass:      "Pass123!",
			expectedErrStr: "Invalid full name:full name must be between 3 and 60 characters",
		},
		{
			name:           "inValid Password",
			paramPhone:     "+62345688990",
			paramName:      "oase",
			paramPass:      "Pass!",
			expectedErrStr: "Invalid password:password must be between 6 and 64 characters",
		},
		// Add more test cases for different scenarios
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			param := generated.PostRegisterJSONBody{}
			param.FullName = &test.paramName
			param.PhoneNumber = &test.paramPhone
			param.Password = &test.paramPass

			err := validateBody(param)
			if test.expectedErrStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedErrStr)
			}
		})
	}
}
