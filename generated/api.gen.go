// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	Password    *string `json:"password,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// PutProfileJSONBody defines parameters for PutProfile.
type PutProfileJSONBody struct {
	FullName    *string `json:"full_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// PostRegisterJSONBody defines parameters for PostRegister.
type PostRegisterJSONBody struct {
	FullName    *string `json:"full_name,omitempty"`
	Password    *string `json:"password,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody PostLoginJSONBody

// PutProfileJSONRequestBody defines body for PutProfile for application/json ContentType.
type PutProfileJSONRequestBody PutProfileJSONBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody PostRegisterJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// User login
	// (POST /login)
	PostLogin(ctx echo.Context) error
	// Get user profile
	// (GET /profile)
	GetProfile(ctx echo.Context) error
	// Update user profile
	// (PUT /profile)
	PutProfile(ctx echo.Context) error
	// Register a new user
	// (POST /register)
	PostRegister(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostLogin(ctx)
	return err
}

// GetProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProfile(ctx)
	return err
}

// PutProfile converts echo context to params.
func (w *ServerInterfaceWrapper) PutProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutProfile(ctx)
	return err
}

// PostRegister converts echo context to params.
func (w *ServerInterfaceWrapper) PostRegister(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRegister(ctx)
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

	router.POST(baseURL+"/login", wrapper.PostLogin)
	router.GET(baseURL+"/profile", wrapper.GetProfile)
	router.PUT(baseURL+"/profile", wrapper.PutProfile)
	router.POST(baseURL+"/register", wrapper.PostRegister)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RVwW7UMBD9lWjgGDUpW5DIrUVQFfVQFSoOVVW5ySRxldhmPG5Zqvw7sr3Z7m6DAGm1",
	"Wk5JPC/jmffe2E9Q6t5ohYotFE9gyxZ7EV4/Emm6RGu0sugXDGmDxBJDuEdrRRMCPDcIBVgmqRoYhhQI",
	"vztJWEFxvQTepCNQ391jyQFosXQkef7F7xsT36EgpGPH7fPXJ029YCjg87evkMYqfaYYhWXmltnA4BNL",
	"VetQm+TOR64sUnKJjbRMgqVWiVBVcq4bqZLjizNI4QHJSq2ggMOD/CCHIQVtUAkjoYBZWErBCG5DlVnn",
	"fw20aMv+6ckJmc8qKOBCWw7ZIbKBlk90NffAUitGFf4RxnSyDH9l91arZwleEm6EtY+aKv9ej3wsF9NN",
	"FVIwrVZ4q1x/hxQUEz/OUTWe18NZCr1Uy888nRBxU601XZkchoXoj1Dhmzz3jwptSdJwJDNSbF1ZorW1",
	"63xlR1PAE1ElC6Yi5vAl5koJx60m+RMrD3o7lehMMZISXWKRHpAS9EaObnN9L2g+2iFK6AOZIV3LLri5",
	"wQk5T5EvFpC/aXqBTQiZJD5gtUJAN4/tzbbf3mKYoLheH6Prm+FmtftT5MR5Bsa2vVvclIvdWtvbsLFn",
	"4FaJHjcs+S5fs+Rsbw09autMJXhS2fyf6HlNWEMBr7LnkzhbHMPZ+hkcKv7D0Mx2t/emW4/y97vb/INW",
	"dSdLXhmT3Wy8hfG7Cs7ZmEB/ClG4n6K5f3+tXI6ofRjJ/+BSmrpHPPUj23s2w/vi5qVdR78lIlH4GGzr",
	"Uw2/AgAA//8NYb+bPAoAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

