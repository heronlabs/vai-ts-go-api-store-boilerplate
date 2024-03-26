package presenters

import models "go-api-store-boilerplate/src/application/api/presenters/models"

type JsonPresenter struct{}

func (jp JsonPresenter) Envelope(payload any) models.JsonModel {

	return models.JsonModel{
		Payload: payload,
	}

}
