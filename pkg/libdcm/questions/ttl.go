package questions

import "github.com/AlecAivazis/survey/v2"

var (
	Ttl = []*survey.Question{
		{
			Name: "interactions",
			Prompt: &survey.Input{
				Message: "How many interactions do you you want to allow for this intent?",
				Default: "0",
			},
			Validate: survey.Required,
		},
		{
			Name: "minutes",
			Prompt: &survey.Input{
				Message: "How many minutes do you you want to allow this intent to be active?",
				Default: "0",
			},
			Validate: survey.Required,
		},
	}
)
