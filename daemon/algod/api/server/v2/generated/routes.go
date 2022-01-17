// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get a Merkle proof for a transaction in a block.
	// (GET /v2/blocks/{round}/transactions/{txid}/proof)
	GetProof(ctx echo.Context, round uint64, txid string, params GetProofParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetProof converts echo context to params.
func (w *ServerInterfaceWrapper) GetProof(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProofParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProof(ctx, round, txid, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/blocks/:round/transactions/:txid/proof", wrapper.GetProof, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
