package model

import "errors"

type Account struct {
	ProviderID        string `json:"provider_id"`
	ProviderType      string `json:"provider_type"`
	ProviderAccountID string `json:"provider_account_id"`
	UserID            string `json:"user_id"`
}

func NewAccount(providerID, providerType, providerAccountID, userID string) (*Account, error) {
	account := Account{}
	if err := account.SetProviderID(providerID); err != nil {
		return nil, err
	}
	if err := account.SetProviderType(providerType); err != nil {
		return nil, err
	}
	if err := account.SetProviderAccountID(providerAccountID); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *Account) SetProviderID(providerID string) error {
	if providerID == "" {
		return errors.New("ProviderID is blank")
	}
	a.ProviderID = providerID
	return nil
}

func (a *Account) SetProviderType(providerType string) error {
	if providerType == "" {
		return errors.New("ProviderType is blank")
	}
	a.ProviderType = providerType
	return nil
}

func (a *Account) SetProviderAccountID(providerAccountID string) error {
	if providerAccountID == "" {
		return errors.New("ProviderAccountID is blank")
	}
	a.ProviderAccountID = providerAccountID
	return nil
}
