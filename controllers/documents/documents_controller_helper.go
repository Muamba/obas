package controllers

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	domain "obas/domain/documents"
	documentIO "obas/io/documents"
)

/**
Get document types
 */
func GetDocumentTypes(app *config.Env) ([]domain.DocumentType, genericHelper.PageToast) {
	var documentTypes []domain.DocumentType
	var alert genericHelper.PageToast
	documentTypes, err := documentIO.GetDocumentTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve document types!"}
	}
	return documentTypes, alert
}

/**
Get document given document id
 */
func GetDocument(app *config.Env, documentId string) (domain.Document, genericHelper.PageToast) {
	var document domain.Document
	var alert genericHelper.PageToast
	document, err := documentIO.GetDocument(documentId)
	if err != nil {
		errorDesc := " ~ Could not retrieve document with id: " + documentId
		app.ErrorLog.Println(err.Error() + errorDesc)
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve document!"}
	}
	return document, alert
}

/**
Get document type given document type id
 */
func GetDocumentType(app *config.Env, documentTypeId string) (domain.DocumentType, genericHelper.PageToast) {
	var documentType domain.DocumentType
	var alert genericHelper.PageToast
	documentType, err := documentIO.GetDocumentType(documentTypeId)
	if err != nil {
		errorDesc := " ~ Could not retrieve document type for id: " + documentTypeId
		app.ErrorLog.Println(err.Error() + errorDesc)
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve document type!"}
	}
	return documentType, alert
}
