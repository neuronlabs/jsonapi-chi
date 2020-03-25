package jsonapichi

import (
	"fmt"

	"github.com/go-chi/chi"
	handler "github.com/neuronlabs/jsonapi-handler"
	"github.com/neuronlabs/jsonapi-handler/middlewares"
	"github.com/neuronlabs/neuron-core/controller"
)

// GroupModel creates chi router group for given 'model' with provided 'endpoints' for the Default neuron Controller.
// By default all endpoints are set.
func GroupModel(creator *handler.Creator, model interface{}, endpoints ...EndpointType) chi.Router {
	r := chi.NewRouter()
	routeModelC(controller.Default(), creator, r, model, endpoints...)
	return r
}

// GroupModelC creates chi router group for given 'model' with provided 'endpoints' for given 'c' controller.
// By default all endpoints are set.
func GroupModelC(c *controller.Controller, creator *handler.Creator, model interface{}, endpoints ...EndpointType) chi.Router {
	r := chi.NewRouter()
	routeModelC(c, creator, r, model, endpoints...)
	return r
}

// RouteModel routes model endpoints for given 'r' chi Router.
// By default all endpoints are set.
func RouteModel(r chi.Router, creator *handler.Creator, model interface{}, endpoints ...EndpointType) {
	routeModelC(controller.Default(), creator, r, model, endpoints...)
}

// RouteModelC routes model endpoints for given 'r' chi Router for given 'c' neuron controller.
// By default all endpoints are set.
func RouteModelC(c *controller.Controller, r chi.Router, creator *handler.Creator, model interface{}, endpoints ...EndpointType) {
	routeModelC(c, creator, r, model, endpoints...)
}

func routeModelC(c *controller.Controller, creator *handler.Creator, r chi.Router, model interface{}, endpoints ...EndpointType) {
	// by default if no endpoints are provided the function takes all endpoints.
	if len(endpoints) == 0 {
		endpoints = allEndpoints()
	}
	mappedModel := c.MustGetModelStruct(model)

	// iterate over provided endpoints and store related handler functions.
	for _, endpoint := range endpoints {
		switch endpoint {
		case Create:
			r.With(middlewares.CheckJSONAPIContentType).Post("/"+mappedModel.Collection(), creator.Create(model))
		case Get:
			r.With(middlewares.AcceptJSONAPIMediaType, GetID).Get(fmt.Sprintf("/%s/{id}", mappedModel.Collection()), creator.Get(model))
		case GetRelatedFields:
			handlers := creator.GetRelatedHandlers(model)
			for field, handlerFunc := range handlers {
				r.With(middlewares.AcceptJSONAPIMediaType, GetID).Get(fmt.Sprintf("/%s/{id}/%s", mappedModel.Collection(), field.NeuronName()), handlerFunc)
			}
		case GetRelationships:
			handlers := creator.GetRelationShipHandlers(model)
			for field, handlerFunc := range handlers {
				r.With(middlewares.AcceptJSONAPIMediaType, GetID).Get(fmt.Sprintf("/%s/{id}/relationships/%s", mappedModel.Collection(), field.NeuronName()), handlerFunc)
			}
		case List:
			r.With(middlewares.AcceptJSONAPIMediaType).Get(fmt.Sprintf("/%s", mappedModel.Collection()), creator.List(model))
		case Patch:
			r.With(middlewares.CheckJSONAPIContentType, GetID).Patch(fmt.Sprintf("/%s", mappedModel.Collection()), creator.Patch(model))
		case PatchRelationships:
			handlers := creator.PatchRelationshipHandlers(model)
			for field, handlerFunc := range handlers {
				r.With(middlewares.CheckJSONAPIContentType, GetID).Patch(fmt.Sprintf("/%s/{id}/relationships/%s", mappedModel.Collection(), field.NeuronName()), handlerFunc)
			}
		case Delete:
			r.With(GetID).Delete(fmt.Sprintf("/%s", mappedModel.Collection()), creator.Delete(model))
		}
	}
}
