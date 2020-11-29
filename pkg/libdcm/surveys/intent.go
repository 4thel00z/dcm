package surveys

import (
	"github.com/4thel00z/dcm/pkg/libdcm"
	"github.com/4thel00z/dcm/pkg/libdcm/questions"
	"github.com/AlecAivazis/survey/v2"
)

func Intent() (libdcm.Intent, error) {
	intent := libdcm.Intent{}

	err := survey.Ask(questions.FirstIntent, &intent)
	if err != nil {
		return libdcm.Intent{}, err
	}
	entities, err := Entities()
	if err != nil {
		return libdcm.Intent{}, err
	}
	intent.Entities = entities
	intent.RequiredEntities = [][]string{}
	if len(entities) > 0 {

		intent.RequiredEntities, err = libdcm.SelectMultipleStringArrays(
			"Do you want to select more required entity pairs?",
			"Select required entities which are OR evaluated, all together are AND evaluated",
			"Entities in the outer list are AND evaluated, lower lists are OR evaluated",
			entities.Names()...,
		)
	}

	if err != nil {
		return libdcm.Intent{}, err
	}

	err = survey.Ask(questions.Ttl, &intent.TTL)
	if err != nil {
		return libdcm.Intent{}, err
	}

	requiredTokens, err := libdcm.SelectMultipleStrings("Do you want to select required tokens?",
		"Select required tokens",
		"Select required tokens, if you don't know what this is skip it.",
		"CVI",
	)

	// just ignore error
	if err != nil {
		requiredTokens = []string{}
	}

	intent.RequiredTokens = requiredTokens
	promptableEntities := entities.Promptable().Names()

	for {
		currentPrompt := libdcm.Prompt{}

		confirm, err := libdcm.ConfirmWithHelp(
			"Do you want to add another entity prompt?",
			"Available prompts if required entity is missing",
		)
		if err != nil {
			return libdcm.Intent{}, nil

		}
		if !confirm {
			break
		}

		pick, err := libdcm.SelectMultipleStrings("",
			"Pick entities for this prompt",
			"Entity combinations for which will be prompted. Need to correspond to reprompt model.",
			promptableEntities...,
		)
		if err != nil {
			return libdcm.Intent{}, nil
		}

		currentPrompt.Entities = pick
		currentPrompt.Messages = libdcm.Multilines("Write entity prompts separated by a newline character")

		intent.Prompts = append(intent.Prompts, currentPrompt)
	}

	return intent, err

}
