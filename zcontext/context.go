package zcontext

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

type contextKey string

var (
	_serviceName string
)

const (
	AuthorizationHeaderName            = "authorization"
	SchemaOperationName                = "__schema"
	reqIdCtxKey             contextKey = "traceId"
	salesOrgCtxKey          contextKey = "salesOrg"
	customerCodeCtxKey      contextKey = "customerCode"
	shipToCodeCtxKey        contextKey = "shipToCode"
	languageCtxKey          contextKey = "language"
	loggerCtxKey            contextKey = "logger"
	operationNameCtxKey     contextKey = "operationName"
	spanNameCtxKey          contextKey = "spanName"
	sourceCtxKey            contextKey = "source"
)

func SetServiceName(serviceName string) {
	if serviceName == "" {
		panic("Service name can not be empty")
	}
	_serviceName = serviceName
}

func ServiceName() string {
	return _serviceName
}

// WithDBTimeout sets db timeout in context from DB_TIMEOUT_SEC
func WithDBTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Second*time.Duration(getDBTimeout()))
}

func getDBTimeout() int {
	dbTimeoutInt, err := strconv.ParseInt(os.Getenv("DB_TIMEOUT_SEC"), 10, 32)
	if err != nil {
		return 300
	}
	return int(dbTimeoutInt)
}

// WithSalesOrg Bind salesOrg to context variable
func WithSalesOrg(ctx context.Context, salesOrg string) context.Context {
	return context.WithValue(ctx, salesOrgCtxKey, salesOrg)
}

// GetSalesOrg get salesOrg from context
func SalesOrg(ctx context.Context) string {
	return readCtx(ctx, salesOrgCtxKey)
}

// WithLanguage binds language to context variable
func WithLanguage(ctx context.Context, language string) context.Context {
	return context.WithValue(ctx, languageCtxKey, language)
}

// Language get language from context
func Language(ctx context.Context) string {
	return readCtx(ctx, languageCtxKey)
}

// WithTraceID Bind correlation id to context variable
func WithTraceID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, reqIdCtxKey, reqID)
}

func WithCustomerCode(ctx context.Context, customerCode string) context.Context {
	return context.WithValue(ctx, customerCodeCtxKey, customerCode)
}

func WithShipToCode(ctx context.Context, shipToCode string) context.Context {
	return context.WithValue(ctx, shipToCodeCtxKey, shipToCode)
}

func CustomerCode(ctx context.Context) string {
	return readCtx(ctx, customerCodeCtxKey)
}

func ShipToCode(ctx context.Context) string {
	return readCtx(ctx, shipToCodeCtxKey)
}

func WithSource(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, sourceCtxKey, val)
}

func Source(ctx context.Context) string {
	return readCtx(ctx, sourceCtxKey)
}

func readCtx(ctx context.Context, key contextKey) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}

func WithSpanName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, spanNameCtxKey, name)
}

func SpanName(ctx context.Context) string {
	return readCtx(ctx, spanNameCtxKey)
}

// BackgroundContext creates a background context with decorated fields
func BackgroundContext() context.Context {
	return WithTraceID(context.Background(), GenerateTraceIDString())
}

func GenerateTraceIDString() string {
	return GenerateTraceID().String()
}

func GenerateTraceID() trace.TraceID {
	id := uuid.New()
	return trace.TraceID(id)
}
