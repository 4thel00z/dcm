package surveys

import (
	"fmt"
	"github.com/4thel00z/dcm/pkg/libdcm"
	"github.com/4thel00z/dcm/pkg/libdcm/questions"
	"github.com/AlecAivazis/survey/v2"
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
			Created:                  nil,
			Intents:                  []libdcm.Intent{},
		}
	)

	err := survey.Ask(questions.Domain, &domain)
	if err != nil {
		return domain, fmt.Errorf("survey.Ask(questions.Domain, &domain): %e", err)
	}
	for {
		confirm, err := libdcm.ConfirmWithHelp(
			"Do you want to add another intent?",
			"Intents of the skill",
		)

		if err != nil {
			return domain, fmt.Errorf("libdcm.ConfirmWithHelp: %e", err)
		}
		if !confirm {
			break
		}
		intent, err := Intent()

		if err != nil {
			return domain, err
		}

		domain.Intents = append(domain.Intents, intent)
	}

	return domain, nil
}
