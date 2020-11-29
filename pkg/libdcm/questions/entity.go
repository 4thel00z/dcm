package questions

import "github.com/AlecAivazis/survey/v2"

var (
	Entity = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Name of the entity?",
			},
			Validate: survey.Required,
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Type of the entity?",
				Options: []string{
					"ZIP_CODE",
					"CITY",
					"ROOM",
					"TIMEZONE",
					"FREETEXT",
					"DEVICE_NAME",
					"OTHER_DEVICE_NAMES",
					"PHONE_NAMES",
					"DYNAMIC_DEVICE_DATA",
					"STT_TEXT",
					"STT_RAW_RESPONSE",
					"USER_ID_HASH",
					"VERIFICATION",
				},
				Help: `ZIP_CODE: ZIP code, e.g. 81549
CITY: City name, e.g. "Munich"
ROOM: Room of a house, e.g. "Living Room", "Bedroom"
TIMEZONE: Timezone, e.g. Europe/Berlin", notes = "For a list, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
FREETEXT: Free text
DEVICE_NAME: The name of the device sending the request
OTHER_DEVICE_NAMES: A list with the user devices of the user. Does not contain the device that sent the request
PHONE_NAMES: A list with the phones that are registered to the user
DYNAMIC_DEVICE_DATA: Device metadata from the dynamic value section
STT_TEXT: Text from the STT. Only available if skill has STT_RESULT permission
STT_RAW_RESPONSE: Raw response from the STT provider. Only available if skill has STT_RAW_RESPONSE permission.
USER_ID_HASH: A unique hash for identifying the user. Only available if skill has USER_ID_HASH permission
VERIFICATION: Result of the voice verification`,
			},
			Validate: survey.Required,
		},
		{
			Name: "fillPolicy",
			Prompt: &survey.MultiSelect{
				Message: "Fill policy of the entity?",
				Options: []string{
					"environment",
					"device",
					"context",
					"verification",
				},
				Help: `environment:A fill policy which fills entities from the environment, for example the current date/time, STT text, etc.
device: Fills entities with data from the device metadata.
context: Fills entities with values from the entities in the context.
verification: Fill policy which which verifies the voice of the user and puts the result the result as attribute.`,
			},
			Validate: survey.Required,
		},
		{
			Name: "useFromSession",
			Prompt: &survey.Confirm{
				Message: "Should entity be used from session?",
				Help:    `Flag if entity is used from session`,
				Default: false,
			},
			Validate: survey.Required,
		},
		{
			Name: "mappable",
			Prompt: &survey.Confirm{
				Message: "Should entity be be canonized?",
				Help:    `Flag if entity should use a mapping service to discover its canonical value`,
				Default: true,
			},
			Validate: survey.Required,
		},
	}
)
