package questions

import "github.com/AlecAivazis/survey/v2"

var (
	Domain = []*survey.Question{
		{
			Name: "modelMinConfidence",
			Prompt: &survey.Input{
				Message: "What is the minimal confidence at which you want your model to be triggered?",
				Default: "0.5",
			},
			Validate: survey.Required,
		},
		{
			Name: "answerModelMinConfidence",
			Prompt: &survey.Input{
				Message: "What is the minimal confidence at which you want your reprompt model to be triggered?",
				Default: "0.5",
			},
			Validate: survey.Required,
		},
	}
)
