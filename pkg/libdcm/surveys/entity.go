package surveys

import (
	"github.com/4thel00z/dcm/pkg/libdcm"
	"github.com/4thel00z/dcm/pkg/libdcm/questions"
	"github.com/AlecAivazis/survey/v2"
)

func Entity() (libdcm.Entity, error) {
	entity := libdcm.Entity{}

	err := survey.Ask(questions.Entity, &entity)
	if err != nil {
		return libdcm.Entity{}, err
	}

	err = survey.Ask(questions.Ttl, &entity.TTL)
	if err != nil {
		return libdcm.Entity{}, err
	}

	return entity, err

}

func Entities() (entities libdcm.Entities, err error) {
	entities = libdcm.Entities{}

	for {
		confirmation, err := libdcm.Confirm("Do you want to define another entity?")
		if err != nil {
			return nil, err
		}
		if !confirmation {
			break
		}
		entity, err := Entity()
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return
}
