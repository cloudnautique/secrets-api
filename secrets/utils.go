package secrets

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (us *UnencryptedSecret) HandleVaultConfig() error {
	valid, err := us.VaultPolicyConfig.Valid()
	if err != nil {
		return err
	}

	if us.ClearText != "" && valid {
		return errors.New("ClearText Must be empty when passing Vault config")
	}

	if valid {
		vpc, err := json.Marshal(us.VaultPolicyConfig)
		if err != nil {
			return err
		}
		us.ClearText = string(vpc)
	}

	return nil
}

func (vp VaultPolicyConfig) Valid() (bool, error) {
	var err error
	valid := true
	msg := "vault config: %s can not be blank"

	if vp.Policy == "" && vp.IssuingToken == "" && vp.VaultURL == "" {
		return false, nil
	}

	if vp.Policy == "" {
		valid = false
		err = fmt.Errorf(msg, "policy")
	}

	if vp.IssuingToken == "" {
		valid = false
		err = fmt.Errorf(msg, "Issuing Token")
	}

	if vp.VaultURL == "" {
		valid = false
		err = fmt.Errorf(msg, "Vault URL")
	}

	return valid, err
}
