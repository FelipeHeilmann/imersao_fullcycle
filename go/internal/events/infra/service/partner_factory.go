package service

import (
	"fmt"
)

type PartnerFactory interface {
	Create(parterId int) (Partner, error)
}

type DefaultParterFactory struct {
	partnerBaseUrls map[int]string
}

func newPartnerFactory(parterBaseURls map[int]string) PartnerFactory {
	return &DefaultParterFactory{partnerBaseUrls: parterBaseURls}
}

func (factory *DefaultParterFactory) Create(parterId int) (Partner, error) {
	baseUrl, ok := factory.partnerBaseUrls[parterId]
	if !ok {
		return nil, fmt.Errorf("partner with id %d not found", parterId)
	}

	switch parterId {
	case 1:
		return &Partner1{BaseUrl: baseUrl}, nil
	default:
		return nil, fmt.Errorf("partner with id %d not found", parterId)
	}
}
