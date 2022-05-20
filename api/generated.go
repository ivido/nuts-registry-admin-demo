// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /web/auth)
	CreateSession(ctx echo.Context) error

	// (GET /web/private)
	CheckSession(ctx echo.Context) error

	// (PUT /web/private/credential/{type}/issuer/{did})
	UpdateCredentialIssuer(ctx echo.Context, pType string, did string) error

	// (GET /web/private/credentials/issuers)
	GetCredentialIssuers(ctx echo.Context) error

	// (GET /web/private/customers)
	GetCustomers(ctx echo.Context) error

	// (POST /web/private/customers)
	ConnectCustomer(ctx echo.Context) error

	// (GET /web/private/customers/{id})
	GetCustomer(ctx echo.Context, id int) error

	// (PUT /web/private/customers/{id})
	UpdateCustomer(ctx echo.Context, id int) error

	// (GET /web/private/customers/{id}/services)
	GetServicesForCustomer(ctx echo.Context, id int) error

	// (POST /web/private/customers/{id}/services)
	EnableCustomerService(ctx echo.Context, id int) error

	// (DELETE /web/private/customers/{id}/services/{type})
	DisableCustomerService(ctx echo.Context, id int, pType string) error

	// (POST /web/private/organizations)
	SearchOrganizations(ctx echo.Context) error

	// (GET /web/private/service-provider)
	GetServiceProvider(ctx echo.Context) error

	// (PUT /web/private/service-provider)
	UpdateServiceProvider(ctx echo.Context) error

	// (GET /web/private/service-provider/endpoints)
	GetEndpoints(ctx echo.Context) error

	// (POST /web/private/service-provider/endpoints)
	RegisterEndpoint(ctx echo.Context) error

	// (DELETE /web/private/service-provider/endpoints/{id})
	DeleteEndpoint(ctx echo.Context, id string) error

	// (GET /web/private/service-provider/services)
	GetServices(ctx echo.Context) error

	// (POST /web/private/service-provider/services)
	AddService(ctx echo.Context) error

	// (POST /web/private/vc)
	IssueVC(ctx echo.Context) error

	// (GET /web/private/vc/templates)
	GetVCTemplates(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateSession converts echo context to params.
func (w *ServerInterfaceWrapper) CreateSession(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateSession(ctx)
	return err
}

// CheckSession converts echo context to params.
func (w *ServerInterfaceWrapper) CheckSession(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CheckSession(ctx)
	return err
}

// UpdateCredentialIssuer converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCredentialIssuer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "type" -------------
	var pType string

	err = runtime.BindStyledParameterWithLocation("simple", false, "type", runtime.ParamLocationPath, ctx.Param("type"), &pType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// ------------- Path parameter "did" -------------
	var did string

	err = runtime.BindStyledParameterWithLocation("simple", false, "did", runtime.ParamLocationPath, ctx.Param("did"), &did)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter did: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCredentialIssuer(ctx, pType, did)
	return err
}

// GetCredentialIssuers converts echo context to params.
func (w *ServerInterfaceWrapper) GetCredentialIssuers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCredentialIssuers(ctx)
	return err
}

// GetCustomers converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomers(ctx)
	return err
}

// ConnectCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) ConnectCustomer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConnectCustomer(ctx)
	return err
}

// GetCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomer(ctx, id)
	return err
}

// UpdateCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCustomer(ctx, id)
	return err
}

// GetServicesForCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) GetServicesForCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetServicesForCustomer(ctx, id)
	return err
}

// EnableCustomerService converts echo context to params.
func (w *ServerInterfaceWrapper) EnableCustomerService(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EnableCustomerService(ctx, id)
	return err
}

// DisableCustomerService converts echo context to params.
func (w *ServerInterfaceWrapper) DisableCustomerService(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "type" -------------
	var pType string

	err = runtime.BindStyledParameterWithLocation("simple", false, "type", runtime.ParamLocationPath, ctx.Param("type"), &pType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DisableCustomerService(ctx, id, pType)
	return err
}

// SearchOrganizations converts echo context to params.
func (w *ServerInterfaceWrapper) SearchOrganizations(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SearchOrganizations(ctx)
	return err
}

// GetServiceProvider converts echo context to params.
func (w *ServerInterfaceWrapper) GetServiceProvider(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetServiceProvider(ctx)
	return err
}

// UpdateServiceProvider converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateServiceProvider(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateServiceProvider(ctx)
	return err
}

// GetEndpoints converts echo context to params.
func (w *ServerInterfaceWrapper) GetEndpoints(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetEndpoints(ctx)
	return err
}

// RegisterEndpoint converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterEndpoint(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterEndpoint(ctx)
	return err
}

// DeleteEndpoint converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteEndpoint(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteEndpoint(ctx, id)
	return err
}

// GetServices converts echo context to params.
func (w *ServerInterfaceWrapper) GetServices(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetServices(ctx)
	return err
}

// AddService converts echo context to params.
func (w *ServerInterfaceWrapper) AddService(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddService(ctx)
	return err
}

// IssueVC converts echo context to params.
func (w *ServerInterfaceWrapper) IssueVC(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.IssueVC(ctx)
	return err
}

// GetVCTemplates converts echo context to params.
func (w *ServerInterfaceWrapper) GetVCTemplates(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVCTemplates(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/web/auth", wrapper.CreateSession)
	router.GET(baseURL+"/web/private", wrapper.CheckSession)
	router.PUT(baseURL+"/web/private/credential/:type/issuer/:did", wrapper.UpdateCredentialIssuer)
	router.GET(baseURL+"/web/private/credentials/issuers", wrapper.GetCredentialIssuers)
	router.GET(baseURL+"/web/private/customers", wrapper.GetCustomers)
	router.POST(baseURL+"/web/private/customers", wrapper.ConnectCustomer)
	router.GET(baseURL+"/web/private/customers/:id", wrapper.GetCustomer)
	router.PUT(baseURL+"/web/private/customers/:id", wrapper.UpdateCustomer)
	router.GET(baseURL+"/web/private/customers/:id/services", wrapper.GetServicesForCustomer)
	router.POST(baseURL+"/web/private/customers/:id/services", wrapper.EnableCustomerService)
	router.DELETE(baseURL+"/web/private/customers/:id/services/:type", wrapper.DisableCustomerService)
	router.POST(baseURL+"/web/private/organizations", wrapper.SearchOrganizations)
	router.GET(baseURL+"/web/private/service-provider", wrapper.GetServiceProvider)
	router.PUT(baseURL+"/web/private/service-provider", wrapper.UpdateServiceProvider)
	router.GET(baseURL+"/web/private/service-provider/endpoints", wrapper.GetEndpoints)
	router.POST(baseURL+"/web/private/service-provider/endpoints", wrapper.RegisterEndpoint)
	router.DELETE(baseURL+"/web/private/service-provider/endpoints/:id", wrapper.DeleteEndpoint)
	router.GET(baseURL+"/web/private/service-provider/services", wrapper.GetServices)
	router.POST(baseURL+"/web/private/service-provider/services", wrapper.AddService)
	router.POST(baseURL+"/web/private/vc", wrapper.IssueVC)
	router.GET(baseURL+"/web/private/vc/templates", wrapper.GetVCTemplates)

}
