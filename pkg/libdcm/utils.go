package libdcm

import (
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

func Confirm(m string) (confirm bool, err error) {
	err = survey.AskOne(&survey.Confirm{
		Message: m,
	}, &confirm)
	return
}
func ConfirmWithHelp(m, help string) (confirm bool, err error) {
	err = survey.AskOne(&survey.Confirm{
		Message: m,
		Help:    help,
	}, &confirm)
	return
}

func SelectMultipleStringArrays(continuePrompt, multiSelectPrompt, help string, choices ...string) (target [][]string, err error) {
	target = [][]string{}
	p := &survey.MultiSelect{
		Message: multiSelectPrompt,
		Options: choices,
		Help:    help,
	}

	for {
		var current []string
		confirmation, err := Confirm(continuePrompt)
		if err != nil {
			return nil, err
		}
		if !confirmation {
			break
		}

		err = survey.AskOne(p, &current)
		if err != nil {
			return nil, err
		}
		target = append(target, current)
	}

	return
}

func SelectMultipleStrings(continuePrompt, multiSelectPrompt, help string, choices ...string) ([]string, error) {
	target := []string{}
	p := &survey.MultiSelect{
		Message: multiSelectPrompt,
		Options: choices,
		Help:    help,
	}

	if continuePrompt != "" {
		confirmation, err := Confirm(continuePrompt)
		if err != nil {
			return target, err
		}
		if !confirmation {
			return target, err
		}
	}

	err := survey.AskOne(p, &target)
	if err != nil {
		return target, err
	}

	return target, err
}

func Multilines(prompt string) []string {
	content := ""
	err := survey.AskOne(&survey.Multiline{
		Message: prompt,
	}, &content)
	if err != nil {
		return []string{}
	}
	return strings.Split(content, "\n")
}
