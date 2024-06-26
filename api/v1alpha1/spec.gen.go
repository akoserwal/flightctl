// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"Qk4h3wkBAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
