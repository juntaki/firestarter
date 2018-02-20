// Code generated by protoc-gen-twirp v5.2.0, DO NOT EDIT.
// source: proto/config.proto

/*
Package proto is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.2.0.

It is generated from these files:
	proto/config.proto
*/
package proto

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import log "log"
import http "net/http"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import strconv "strconv"
import json "encoding/json"
import url "net/url"

// =======================
// ConfigService Interface
// =======================

type ConfigService interface {
	GetConfig(context.Context, *GetConfigListRequest) (*ConfigList, error)

	SetConfig(context.Context, *Config) (*SetConfigResponse, error)

	DeleteConfig(context.Context, *Config) (*DeleteConfigResponse, error)
}

// =============================
// ConfigService Protobuf Client
// =============================

type configServiceProtobufClient struct {
	client HTTPClient
	urls   [3]string
}

// NewConfigServiceProtobufClient creates a Protobuf client that implements the ConfigService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewConfigServiceProtobufClient(addr string, client HTTPClient) ConfigService {
	prefix := urlBase(addr) + ConfigServicePathPrefix
	urls := [3]string{
		prefix + "GetConfig",
		prefix + "SetConfig",
		prefix + "DeleteConfig",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &configServiceProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &configServiceProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *configServiceProtobufClient) GetConfig(ctx context.Context, in *GetConfigListRequest) (*ConfigList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "GetConfig")
	out := new(ConfigList)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	return out, err
}

func (c *configServiceProtobufClient) SetConfig(ctx context.Context, in *Config) (*SetConfigResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "SetConfig")
	out := new(SetConfigResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

func (c *configServiceProtobufClient) DeleteConfig(ctx context.Context, in *Config) (*DeleteConfigResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteConfig")
	out := new(DeleteConfigResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[2], in, out)
	return out, err
}

// =========================
// ConfigService JSON Client
// =========================

type configServiceJSONClient struct {
	client HTTPClient
	urls   [3]string
}

// NewConfigServiceJSONClient creates a JSON client that implements the ConfigService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewConfigServiceJSONClient(addr string, client HTTPClient) ConfigService {
	prefix := urlBase(addr) + ConfigServicePathPrefix
	urls := [3]string{
		prefix + "GetConfig",
		prefix + "SetConfig",
		prefix + "DeleteConfig",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &configServiceJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &configServiceJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *configServiceJSONClient) GetConfig(ctx context.Context, in *GetConfigListRequest) (*ConfigList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "GetConfig")
	out := new(ConfigList)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	return out, err
}

func (c *configServiceJSONClient) SetConfig(ctx context.Context, in *Config) (*SetConfigResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "SetConfig")
	out := new(SetConfigResponse)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

func (c *configServiceJSONClient) DeleteConfig(ctx context.Context, in *Config) (*DeleteConfigResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteConfig")
	out := new(DeleteConfigResponse)
	err := doJSONRequest(ctx, c.client, c.urls[2], in, out)
	return out, err
}

// ============================
// ConfigService Server Handler
// ============================

type configServiceServer struct {
	ConfigService
	hooks *twirp.ServerHooks
}

func NewConfigServiceServer(svc ConfigService, hooks *twirp.ServerHooks) TwirpServer {
	return &configServiceServer{
		ConfigService: svc,
		hooks:         hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *configServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// ConfigServicePathPrefix is used for all URL paths on a twirp ConfigService server.
// Requests are always: POST ConfigServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const ConfigServicePathPrefix = "/twirp/proto.ConfigService/"

func (s *configServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "proto")
	ctx = ctxsetters.WithServiceName(ctx, "ConfigService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/proto.ConfigService/GetConfig":
		s.serveGetConfig(ctx, resp, req)
		return
	case "/twirp/proto.ConfigService/SetConfig":
		s.serveSetConfig(ctx, resp, req)
		return
	case "/twirp/proto.ConfigService/DeleteConfig":
		s.serveDeleteConfig(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *configServiceServer) serveGetConfig(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetConfigJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetConfigProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *configServiceServer) serveGetConfigJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(GetConfigListRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *ConfigList
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.GetConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ConfigList and nil error while calling GetConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) serveGetConfigProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(GetConfigListRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *ConfigList
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.GetConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ConfigList and nil error while calling GetConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) serveSetConfig(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSetConfigJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSetConfigProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *configServiceServer) serveSetConfigJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SetConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(Config)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *SetConfigResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.SetConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SetConfigResponse and nil error while calling SetConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) serveSetConfigProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SetConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Config)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *SetConfigResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.SetConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SetConfigResponse and nil error while calling SetConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) serveDeleteConfig(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDeleteConfigJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDeleteConfigProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *configServiceServer) serveDeleteConfigJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(Config)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteConfigResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.DeleteConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteConfigResponse and nil error while calling DeleteConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) serveDeleteConfigProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteConfig")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Config)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteConfigResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.DeleteConfig(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteConfigResponse and nil error while calling DeleteConfig. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *configServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *configServiceServer) ProtocGenTwirpVersion() string {
	return "v5.2.0"
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
}

// WriteError writes an HTTP response with a valid Twirp error format.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON (instead of protobuf)
	resp.WriteHeader(statusCode)                          // HTTP response status code

	respBody := marshalErrorToJSON(twerr)
	_, err2 := resp.Write(respBody)
	if err2 != nil {
		log.Printf("unable to send error message %q: %s", twerr, err2)
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// closebody closes a response or request body and just logs
// any error encountered while closing, since errors are
// considered very unusual.
func closebody(body io.Closer) {
	if err := body.Close(); err != nil {
		log.Printf("error closing body: %q", err)
	}
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.2.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}
	var tj twerrJSON
	if err := json.Unmarshal(respBodyBytes, &tj); err != nil {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}
func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrappedError implements the github.com/pkg/errors.Causer interface, allowing errors to be
// examined for their root cause.
type wrappedError struct {
	msg   string
	cause error
}

func wrapErr(err error, msg string) error { return &wrappedError{msg: msg, cause: err} }
func (e *wrappedError) Cause() error      { return e.cause }
func (e *wrappedError) Error() string     { return e.msg + ": " + e.cause.Error() }

// clientError adds consistency to errors generated in the client
func clientError(desc string, err error) twirp.Error {
	return twirp.InternalErrorWith(wrapErr(err, desc))
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest is common code to make a request to the remote twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read response body", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return clientError("failed to unmarshal proto response", err)
	}
	return nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return clientError("failed to unmarshal json response", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xa9, 0x40, 0x7f, 0x2e, 0xb0, 0xe0, 0x6a, 0xc8, 0x04, 0x13, 0xd3, 0x34, 0x31, 0xe9,
	0x0a, 0x13, 0x48, 0x5c, 0xe9, 0x42, 0x21, 0x31, 0x26, 0xac, 0x06, 0x7c, 0x80, 0x32, 0x5e, 0xb1,
	0xb1, 0xb4, 0xb5, 0x1d, 0x0c, 0xbe, 0x9f, 0x4f, 0xe3, 0x53, 0x98, 0xce, 0x94, 0xf2, 0x23, 0xab,
	0xe6, 0x7c, 0xe7, 0xde, 0xce, 0xcd, 0x39, 0x80, 0x69, 0x96, 0xc8, 0xe4, 0x46, 0x24, 0xf1, 0x5b,
	0xb8, 0x1c, 0x28, 0x81, 0x4d, 0xf5, 0xf1, 0x7a, 0x70, 0xf1, 0x44, 0x72, 0xac, 0x9c, 0x69, 0x98,
	0x4b, 0x4e, 0x9f, 0x6b, 0xca, 0xa5, 0x77, 0x0e, 0xdd, 0xd9, 0x96, 0x73, 0xca, 0xd3, 0x24, 0xce,
	0xa9, 0x18, 0x9e, 0x50, 0x44, 0x92, 0x8e, 0xf8, 0xaf, 0x01, 0xa6, 0x46, 0x78, 0x05, 0x30, 0x0e,
	0xa2, 0x68, 0x11, 0x88, 0x8f, 0xe7, 0x09, 0x33, 0x5c, 0xc3, 0x77, 0x38, 0x88, 0x8a, 0x60, 0x1f,
	0xec, 0xf1, 0x7b, 0x10, 0xc7, 0x14, 0xe5, 0xec, 0xcc, 0xad, 0xfb, 0x0e, 0xb7, 0x45, 0xa9, 0x11,
	0xa1, 0x31, 0xa7, 0x8d, 0x64, 0x75, 0xb5, 0xd5, 0x90, 0xb4, 0x91, 0xd8, 0x03, 0x93, 0xd3, 0x92,
	0x36, 0x29, 0x6b, 0x2a, 0x6a, 0x66, 0x4a, 0xa1, 0x0b, 0xad, 0x17, 0x3e, 0x9d, 0xd3, 0x2a, 0x8d,
	0x02, 0x49, 0xcc, 0x54, 0x66, 0x6b, 0xbd, 0x43, 0xe8, 0x41, 0xfb, 0x31, 0x79, 0xfd, 0xae, 0x46,
	0x2c, 0x35, 0xd2, 0x5e, 0xec, 0x31, 0x64, 0x60, 0xa9, 0xbb, 0xb3, 0x15, 0xb3, 0x5d, 0xc3, 0xb7,
	0xb9, 0x25, 0xb4, 0x2c, 0x9c, 0x07, 0x21, 0xc3, 0x24, 0xce, 0x99, 0xa3, 0xce, 0xb4, 0x02, 0x2d,
	0xbd, 0x11, 0xc0, 0x2e, 0x2e, 0xbc, 0x06, 0x53, 0xc7, 0xca, 0x0c, 0xb7, 0xee, 0xb7, 0x86, 0x1d,
	0x1d, 0xef, 0xa0, 0x4c, 0xa8, 0x34, 0x87, 0x3f, 0x06, 0x74, 0x34, 0x9a, 0x51, 0xf6, 0x15, 0x0a,
	0xc2, 0x7b, 0x70, 0xaa, 0xe0, 0xf1, 0xb2, 0xdc, 0x3a, 0x55, 0x45, 0xbf, 0x7b, 0xf0, 0xcb, 0xc2,
	0xf1, 0x6a, 0x78, 0x0b, 0x4e, 0xd5, 0x0f, 0x1e, 0x3e, 0xda, 0x67, 0xa5, 0xfc, 0x5f, 0x60, 0x0d,
	0xef, 0xa0, 0xbd, 0x5f, 0xe1, 0xf1, 0xea, 0xf6, 0x90, 0x93, 0x35, 0xd7, 0x16, 0xa6, 0x72, 0x47,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa1, 0x3c, 0xcd, 0x51, 0x02, 0x00, 0x00,
}
