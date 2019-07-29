package io

import (
	"encoding/json"
	"errors"
	"obas/src/api"
	domain "obas/src/domain/mail"
)

const mailconfig = api.BASE_URL + "/mail"

type LogEvent domain.MailConfig

func GetMailConfigs() ([]domain.MailConfig, error) {
	entities := []domain.MailConfig{}
	resp, _ := api.Rest().Get(mailconfig + "/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func GetMailConfig(id string) (domain.MailConfig, error) {
	entity := domain.MailConfig{}
	resp, _ := api.Rest().Get(mailconfig + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func CreateMailConfig(entity domain.MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailconfig + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func UpdateMailConfig(entity domain.MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailconfig + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func DeleteMailConfig(entity domain.MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailconfig + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
