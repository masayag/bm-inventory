// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/filanov/bm-inventory/restapi/operations/inventory"
)

// NewBMInventoryAPI creates a new BMInventory instance
func NewBMInventoryAPI(spec *loads.Document) *BMInventoryAPI {
	return &BMInventoryAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		InventoryCreateImageHandler: inventory.CreateImageHandlerFunc(func(params inventory.CreateImageParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.CreateImage has not yet been implemented")
		}),
		InventoryDeregisterClusterHandler: inventory.DeregisterClusterHandlerFunc(func(params inventory.DeregisterClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.DeregisterCluster has not yet been implemented")
		}),
		InventoryDeregisterNodeHandler: inventory.DeregisterNodeHandlerFunc(func(params inventory.DeregisterNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.DeregisterNode has not yet been implemented")
		}),
		InventoryGetClusterHandler: inventory.GetClusterHandlerFunc(func(params inventory.GetClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.GetCluster has not yet been implemented")
		}),
		InventoryGetImageHandler: inventory.GetImageHandlerFunc(func(params inventory.GetImageParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.GetImage has not yet been implemented")
		}),
		InventoryGetNextStepsHandler: inventory.GetNextStepsHandlerFunc(func(params inventory.GetNextStepsParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.GetNextSteps has not yet been implemented")
		}),
		InventoryGetNodeHandler: inventory.GetNodeHandlerFunc(func(params inventory.GetNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.GetNode has not yet been implemented")
		}),
		InventoryListClustersHandler: inventory.ListClustersHandlerFunc(func(params inventory.ListClustersParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.ListClusters has not yet been implemented")
		}),
		InventoryListImagesHandler: inventory.ListImagesHandlerFunc(func(params inventory.ListImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.ListImages has not yet been implemented")
		}),
		InventoryListNodesHandler: inventory.ListNodesHandlerFunc(func(params inventory.ListNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.ListNodes has not yet been implemented")
		}),
		InventoryPostStepReplyHandler: inventory.PostStepReplyHandlerFunc(func(params inventory.PostStepReplyParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.PostStepReply has not yet been implemented")
		}),
		InventoryRegisterClusterHandler: inventory.RegisterClusterHandlerFunc(func(params inventory.RegisterClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.RegisterCluster has not yet been implemented")
		}),
		InventoryRegisterNodeHandler: inventory.RegisterNodeHandlerFunc(func(params inventory.RegisterNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation inventory.RegisterNode has not yet been implemented")
		}),
	}
}

/*BMInventoryAPI Bare metal inventory */
type BMInventoryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// InventoryCreateImageHandler sets the operation handler for the create image operation
	InventoryCreateImageHandler inventory.CreateImageHandler
	// InventoryDeregisterClusterHandler sets the operation handler for the deregister cluster operation
	InventoryDeregisterClusterHandler inventory.DeregisterClusterHandler
	// InventoryDeregisterNodeHandler sets the operation handler for the deregister node operation
	InventoryDeregisterNodeHandler inventory.DeregisterNodeHandler
	// InventoryGetClusterHandler sets the operation handler for the get cluster operation
	InventoryGetClusterHandler inventory.GetClusterHandler
	// InventoryGetImageHandler sets the operation handler for the get image operation
	InventoryGetImageHandler inventory.GetImageHandler
	// InventoryGetNextStepsHandler sets the operation handler for the get next steps operation
	InventoryGetNextStepsHandler inventory.GetNextStepsHandler
	// InventoryGetNodeHandler sets the operation handler for the get node operation
	InventoryGetNodeHandler inventory.GetNodeHandler
	// InventoryListClustersHandler sets the operation handler for the list clusters operation
	InventoryListClustersHandler inventory.ListClustersHandler
	// InventoryListImagesHandler sets the operation handler for the list images operation
	InventoryListImagesHandler inventory.ListImagesHandler
	// InventoryListNodesHandler sets the operation handler for the list nodes operation
	InventoryListNodesHandler inventory.ListNodesHandler
	// InventoryPostStepReplyHandler sets the operation handler for the post step reply operation
	InventoryPostStepReplyHandler inventory.PostStepReplyHandler
	// InventoryRegisterClusterHandler sets the operation handler for the register cluster operation
	InventoryRegisterClusterHandler inventory.RegisterClusterHandler
	// InventoryRegisterNodeHandler sets the operation handler for the register node operation
	InventoryRegisterNodeHandler inventory.RegisterNodeHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *BMInventoryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BMInventoryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BMInventoryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BMInventoryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BMInventoryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BMInventoryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BMInventoryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BMInventoryAPI
func (o *BMInventoryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.InventoryCreateImageHandler == nil {
		unregistered = append(unregistered, "inventory.CreateImageHandler")
	}
	if o.InventoryDeregisterClusterHandler == nil {
		unregistered = append(unregistered, "inventory.DeregisterClusterHandler")
	}
	if o.InventoryDeregisterNodeHandler == nil {
		unregistered = append(unregistered, "inventory.DeregisterNodeHandler")
	}
	if o.InventoryGetClusterHandler == nil {
		unregistered = append(unregistered, "inventory.GetClusterHandler")
	}
	if o.InventoryGetImageHandler == nil {
		unregistered = append(unregistered, "inventory.GetImageHandler")
	}
	if o.InventoryGetNextStepsHandler == nil {
		unregistered = append(unregistered, "inventory.GetNextStepsHandler")
	}
	if o.InventoryGetNodeHandler == nil {
		unregistered = append(unregistered, "inventory.GetNodeHandler")
	}
	if o.InventoryListClustersHandler == nil {
		unregistered = append(unregistered, "inventory.ListClustersHandler")
	}
	if o.InventoryListImagesHandler == nil {
		unregistered = append(unregistered, "inventory.ListImagesHandler")
	}
	if o.InventoryListNodesHandler == nil {
		unregistered = append(unregistered, "inventory.ListNodesHandler")
	}
	if o.InventoryPostStepReplyHandler == nil {
		unregistered = append(unregistered, "inventory.PostStepReplyHandler")
	}
	if o.InventoryRegisterClusterHandler == nil {
		unregistered = append(unregistered, "inventory.RegisterClusterHandler")
	}
	if o.InventoryRegisterNodeHandler == nil {
		unregistered = append(unregistered, "inventory.RegisterNodeHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BMInventoryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BMInventoryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BMInventoryAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BMInventoryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BMInventoryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BMInventoryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the b m inventory API
func (o *BMInventoryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BMInventoryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/images"] = inventory.NewCreateImage(o.context, o.InventoryCreateImageHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{cluster_id}"] = inventory.NewDeregisterCluster(o.context, o.InventoryDeregisterClusterHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/nodes/{node_id}"] = inventory.NewDeregisterNode(o.context, o.InventoryDeregisterNodeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}"] = inventory.NewGetCluster(o.context, o.InventoryGetClusterHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/images/{image_id}"] = inventory.NewGetImage(o.context, o.InventoryGetImageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/nodes/{node_id}/next-steps"] = inventory.NewGetNextSteps(o.context, o.InventoryGetNextStepsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/nodes/{node_id}"] = inventory.NewGetNode(o.context, o.InventoryGetNodeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters"] = inventory.NewListClusters(o.context, o.InventoryListClustersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/images"] = inventory.NewListImages(o.context, o.InventoryListImagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/nodes"] = inventory.NewListNodes(o.context, o.InventoryListNodesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/nodes/{node_id}/next-steps/reply"] = inventory.NewPostStepReply(o.context, o.InventoryPostStepReplyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = inventory.NewRegisterCluster(o.context, o.InventoryRegisterClusterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/nodes"] = inventory.NewRegisterNode(o.context, o.InventoryRegisterNodeHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BMInventoryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BMInventoryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BMInventoryAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BMInventoryAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BMInventoryAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
