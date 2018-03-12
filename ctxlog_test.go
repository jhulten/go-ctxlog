package ctxlog

import (
	"bytes"
	"context"
	"log"
	"reflect"
	"testing"
)

func TestNewContextFromWriter(t *testing.T) {
	var out bytes.Buffer
	ctx := WithWriter(context.Background(), &out)
	testContextCreation(t, ctx)
}

func TestNewContextFromLogger(t *testing.T) {
	var out bytes.Buffer
	ctx := WithLogger(context.Background(), log.New(&out, "TestNewContext", log.Lshortfile))
	testContextCreation(t, ctx)
}

func TestNewContext(t *testing.T) {
	var out bytes.Buffer
	ctx := WithLogContext(context.Background(), &LogContext{log.New(&out, "TestNewContext", log.Lshortfile)})
	testContextCreation(t, ctx)
}

func testContextCreation(t *testing.T, ctx context.Context) {
	v := ctx.Value(contextKey)
	if reflect.TypeOf(v) != reflect.PtrTo(reflect.TypeOf(LogContext{})) {
		t.Fatalf("type mismatch: %v, expected %v", reflect.TypeOf(v), reflect.PtrTo(reflect.TypeOf(LogContext{})))
	}
}

func TestFromContext(t *testing.T) {
	var out bytes.Buffer
	ctx := context.WithValue(context.Background(), contextKey, &LogContext{log.New(&out, "TestNewContext", log.Lshortfile)})
	v, ok := FromContext(ctx)
	if !ok {
		t.Fatalf("context did not find logContext: %v", ctx)
	}
	if reflect.TypeOf(v) != reflect.PtrTo(reflect.TypeOf(LogContext{})) {
		t.Fatalf("type mismatch: %v, expected %v", reflect.TypeOf(v), reflect.PtrTo(reflect.TypeOf(LogContext{})))
	}
}

func TestLog(t *testing.T) {

}

func TestPLog(t *testing.T) {

}
