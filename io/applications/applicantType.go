package applications

import (
	"errors"
	"obas/api"
	domain "obas/domain/application"
)

const applicanttypeURL = api.BASE_URL + "/application/applicanttype"

func CreateApplicantType(obj domain.ApplicantType) (domain.ApplicantType, error) {
	entity := domain.ApplicantType{}
	resp, _ := api.Rest().SetBody(obj).Post(applicanttypeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetApplicantType(applicantTypeId string) (domain.ApplicantType, error) {
	entity := domain.ApplicantType{}
	resp, _ := api.Rest().Get(applicanttypeURL + "/get/" + applicantTypeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetApplicantTypes() ([]domain.ApplicantType, error) {
	entity := []domain.ApplicantType{}
	//entity = append(entity, domain.ApplicantType{"1", "Matric Applicant", ""})
	//entity = append(entity, domain.ApplicantType{"2", "University Applicant", ""})
	resp, _ := api.Rest().Get(applicanttypeURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteApplicantType(obj domain.ApplicantType) (bool, error) {
	resp, _ := api.Rest().SetBody(obj).Post(applicanttypeURL + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicantType(obj domain.ApplicantType) (domain.ApplicantType, error) {
	entity := domain.ApplicantType{}
	resp, _ := api.Rest().SetBody(obj).Post(applicanttypeURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
