package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectPrompt(message string, options []string) (int, error) {
	questions := []*survey.Question{
		{
			Name: "choice",
			Prompt: &survey.Select{
				Message: message,
				Options: options,
			},
			Validate: survey.Required,
		},
	}

	choices := struct {
		Choice int
	}{}

	err := survey.Ask(questions, &choices)
	if err != nil {
		return -1, err
	}

	return choices.Choice, nil
}

func MultiSelectPrompt(message string, options []string) ([]string, error) {

	choices := []string{}
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}

	err := survey.AskOne(prompt, &choices)
	if err != nil {
		return nil, err
	}

	return choices, nil
}
