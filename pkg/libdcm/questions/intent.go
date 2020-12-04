package questions

import "github.com/AlecAivazis/survey/v2"

var (
	FirstIntent = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of your Intent?",
			},
			Validate: survey.Required,
		},
		{
			Name: "entityFiller",
			Prompt: &survey.Input{
				Message: "What is the name of the entityFiller reprompt? (skip this if no reprompt is needed)",
				Default: "",
			},
		},
		{
			Name: "fallbackToNoSession",
			Prompt: &survey.Confirm{
				Message: "Do you want to fallback to the session for entity retrieval?",
				Default: true,
			},
		},
	}
)
