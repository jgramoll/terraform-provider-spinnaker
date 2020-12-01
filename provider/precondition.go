package provider

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

type precondition struct {
	FailPipeline bool                    `mapstructure:"fail_pipeline"`
	Type         client.PreconditionType `mapstructure:"type"`
	Context      map[string]interface{}  `mapstructure:"context"`
}

func newPrecondition(t client.PreconditionType) *precondition {
	return &precondition{
		FailPipeline: true,
		Type:         t,
		Context:      map[string]interface{}{},
	}
}

func toClientPreconditions(p *[]precondition) (*[]client.Precondition, error) {
	clientPreconditions := []client.Precondition{}

	for _, precondition := range *p {
		var preconditionMap map[string]interface{}
		if err := mapstructure.Decode(precondition, &preconditionMap); err != nil {
			return nil, err
		}

		// TODO is there a better way to data clean up?
		context, ok := preconditionMap["context"].(map[string]interface{})
		if ok {
			newContext := map[string]interface{}{}

			switch precondition.Type {
			case client.PreconditionClusterSizeType:
				regions, ok := context["regions"].(string)
				if ok {
					regionsArray := strings.Split(regions, ",")
					for i, r := range regionsArray {
						regionsArray[i] = strings.TrimSpace(r)
					}
					context["regions"] = regionsArray
				}

				expectedStr, ok := context["expected"].(string)
				if ok {
					expected, err := strconv.Atoi(expectedStr)
					if err == nil {
						context["expected"] = expected
					}
				}
			}

			for k, v := range context {
				newContext[strcase.ToCamel(k)] = v
			}
			preconditionMap["context"] = newContext
		}

		clientPrecondition, err := client.ParsePrecondition(preconditionMap, precondition.Type)
		if err != nil {
			return nil, err
		}
		clientPreconditions = append(clientPreconditions, clientPrecondition)
	}

	return &clientPreconditions, nil
}

func fromClientPreconditions(clientPreconditions *[]client.Precondition) (*[]precondition, error) {
	p := []precondition{}

	for _, clientPrecondition := range *clientPreconditions {
		precondition := newPrecondition(clientPrecondition.GetType())

		if err := mapstructure.Decode(clientPrecondition, precondition); err != nil {
			log.Printf("[ERROR] parsing check precondition stage %s\n", err)
			return nil, err
		}

		// TODO is there a better way to data clean up?
		newContext := map[string]interface{}{}
		for k, v := range precondition.Context {
			var val string
			switch reflect.TypeOf(v).Kind() {
			default:
				val = fmt.Sprint(v)
			case reflect.Slice:
				vArr, ok := v.([]string)
				if ok {
					val = strings.Join(vArr, ",")
				} else {
					val = fmt.Sprint(v)
				}
			}
			newContext[strcase.ToSnake(k)] = val
		}
		precondition.Context = newContext

		p = append(p, *precondition)
	}

	return &p, nil
}
