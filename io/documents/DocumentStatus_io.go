package documents

import (
	"errors"
	"obas/api"
	domain3 "obas/domain/documents"
)

const DocumentStatus = api.BASE_URL + "Document/status"

func CreateDocumentStatus(obj domain3.DocumentStatus) (domain3.DocumentStatus, error) {
	entity := domain3.DocumentStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(DocumentStatus + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDocumentStatusFor(documentId, statusId string) (domain3.DocumentStatus, error) {
	entity := domain3.DocumentStatus{}
	resp, _ := api.Rest().Get(DocumentStatus + "/getforstatus/" + documentId + "/" + statusId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDocumentStatus(documentId string) (domain3.DocumentStatus, error) {
	entity := domain3.DocumentStatus{}
	resp, _ := api.Rest().Get(DocumentStatus + "/getfordocument/" + documentId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetdocumentStatues() (domain3.DocumentStatus, error) {
	entity := domain3.DocumentStatus{}
	resp, _ := api.Rest().Get(DocumentStatus + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
