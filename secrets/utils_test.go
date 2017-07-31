package secrets

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

type validCase struct {
	Config     VaultPolicyConfig
	Result     bool
	ErrorState bool
}

type handleVaultConfigCase struct {
	Secret     UnencryptedSecret
	ErrorState bool
}

var (
	validCases = []validCase{
		validCase{
			Config: VaultPolicyConfig{
				VaultURL:     "",
				Policy:       "",
				IssuingToken: "",
			},
			Result:     false,
			ErrorState: false,
		},
		validCase{
			Config: VaultPolicyConfig{
				VaultURL:     "http://localhost",
				Policy:       "",
				IssuingToken: "",
			},
			Result:     false,
			ErrorState: true,
		},
		validCase{
			Config: VaultPolicyConfig{
				VaultURL:     "http://localhost",
				Policy:       "realPolicy",
				IssuingToken: "",
			},
			Result:     false,
			ErrorState: true,
		},
		validCase{
			Config: VaultPolicyConfig{
				VaultURL:     "http://localhost",
				Policy:       "blah",
				IssuingToken: "aaa-aaa-bbb",
			},
			Result:     true,
			ErrorState: false,
		},
	}

	handleVaultConfigCases = []handleVaultConfigCase{
		handleVaultConfigCase{
			Secret: UnencryptedSecret{
				ClearText: "Hello",
				VaultPolicyConfig: VaultPolicyConfig{
					VaultURL:     "http://localhost",
					Policy:       "blah",
					IssuingToken: "aaa-aaa-bbb",
				},
			},
			ErrorState: true,
		},
		handleVaultConfigCase{
			Secret: UnencryptedSecret{
				ClearText: "",
				VaultPolicyConfig: VaultPolicyConfig{
					VaultURL:     "http://localhost",
					Policy:       "blah",
					IssuingToken: "aaa-aaa-bbb",
				},
			},
			ErrorState: false,
		},
		handleVaultConfigCase{
			Secret: UnencryptedSecret{
				ClearText: "badkitty",
				VaultPolicyConfig: VaultPolicyConfig{
					VaultURL:     "",
					Policy:       "",
					IssuingToken: "",
				},
			},
			ErrorState: false,
		},
	}
)

func TestValidFunction(t *testing.T) {
	for _, c := range validCases {
		val, err := c.Config.Valid()
		if val != c.Result {
			logrus.Errorf("%#v", c)
			t.Errorf("Got: %v expected: %v", val, c.Result)
		}

		if c.ErrorState && err == nil {
			logrus.Errorf("%#v", c)
			t.Error("Expected Error")
		}
	}
}

func TestHandleConfigFunction(t *testing.T) {
	for _, c := range handleVaultConfigCases {
		err := c.Secret.HandleVaultConfig()
		if c.ErrorState && err == nil {
			t.Error("Expected Error")
		}
	}
}
