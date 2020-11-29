package libdcm

import "time"

type DomainContextMetadata struct {
	Version                  int       `json:"version"`
	ExportVersion            string    `json:"exportVersion" survey:"exportVersion"`
	Model                    *string   `json:"model" survey:"model"`
	ModelMinConfidence       float64   `json:"modelMinConfidence" survey:"modelMinConfidence"`
	AnswerModel              *string   `json:"answerModel" survey:"answerModel"`
	AnswerModelMinConfidence float64   `json:"answerModelMinConfidence" survey:"answerModelMinConfidence"`
	Type                     string    `json:"type" survey:"type"`
	Created                  time.Time `json:"created"`
	Intents                  []Intent  `json:"intents" survey:"intents"`
}

type TTL struct {
	Interactions int `json:"interactions" survey:"interactions"`
	Minutes      int `json:"minutes" survey:"minutes"`
}

type Entity struct {
	Name           string   `json:"name" survey:"name"`
	Type           string   `json:"type" survey:"type"`
	FillPolicy     []string `json:"fillPolicy" survey:"fillPolicy"`
	TTL            TTL      `json:"ttl" survey:"ttl"`
	UseFromSession bool     `json:"useFromSession" survey:"useFromSession"`
	Mappable       bool     `json:"mappable" survey:"mappable"`
}

type Entities []Entity

func (es Entities) Names() []string {
	names := []string{}
	for _, e := range es {
		names = append(names, e.Name)
	}
	return names
}

func (es Entities) Promptable() Entities {
	promptables := Entities{}
	for _, e := range es {
		for _, policy := range e.FillPolicy {
			if policy == "prompt" {
				promptables = append(promptables, e)
			}
		}
	}
	return promptables
}

type Prompt struct {
	Entities []string `json:"entities" survey:"entities"`
	Messages []string `json:"messages" survey:"messages"`
}

type Intent struct {
	Name                string     `json:"name" survey:"name"`
	TTL                 TTL        `json:"ttl" survey:"ttl"`
	RequiredEntities    [][]string `json:"requiredEntities" survey:"requiredEntities"`
	EntityFiller        string     `json:"entityFiller" survey:"entityFiller"`
	FallbackToNoSession bool       `json:"fallbackToNoSession" survey:"fallbackToNoSession"`
	Entities            []Entity   `json:"entities" survey:"entities"`
	Prompts             []Prompt   `json:"prompts" survey:"prompts"`
	RequiredTokens      []string   `json:"requiredTokens" survey:"requiredTokens"`
	Errors              []Error    `json:"errors" survey:"errors"`
}

type Error struct {
	Code            int        `json:"errorCode" survey:"errorCode"`
	MatchesEntities []string   `json:"matchesEntities" survey:"matchesEntities"`
	Messages        [][]string `json:"messages" survey:"messages"`
}
