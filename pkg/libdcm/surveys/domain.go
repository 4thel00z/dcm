package surveys

import (
	"github.com/4thel00z/dcm/pkg/libdcm"
	"github.com/4thel00z/dcm/pkg/libdcm/questions"
	"github.com/AlecAivazis/survey/v2"
	"time"
)

func Domain() (libdcm.DomainContextMetadata, error) {
	var (
		domain = libdcm.DomainContextMetadata{
			Version:                  1,
			ExportVersion:            "1",
			Model:                    nil,
			ModelMinConfidence:       0.5,
			AnswerModel:              nil,
			AnswerModelMinConfidence: 0.5,
			Type:                     "nuance",
			Created:                  time.Now(),
			Intents:                  nil,
		}
	)

	err := survey.Ask(questions.Domain, &domain)

	for {
		confirm, err := libdcm.ConfirmWithHelp(
			"Do you want to add another intent?",
			"Intents of the skill",
		)

		if err != nil {
			return libdcm.DomainContextMetadata{}, nil
		}
		if !confirm {
			break
		}
		intent, err := Intent()

		if err != nil {
			return libdcm.DomainContextMetadata{}, nil
		}

		domain.Intents = append(domain.Intents, intent)
	}

	return domain, err
}
