package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTClaims_Validate(t *testing.T) {
	tests := []struct {
		name         string
		token        string
		expectedCode int
		expectedErr  error
	}{
		{
			name:         "Valid Token",
			token:        generateTestToken(123),
			expectedCode: 200,
			expectedErr:  nil,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jwt := NewJWT()
			payload, statusCode, err := jwt.Validate(test.token)

			assert.Equal(t, test.expectedCode, statusCode)
			assert.Equal(t, test.expectedErr, err)

			if test.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, int64(123), payload.UserID)
			}
		})
	}
}

func generateTestToken(userID int64) string {
	jwt := NewJWT()
	token, _ := jwt.GenerateJWTToken(userID)
	return token
}
