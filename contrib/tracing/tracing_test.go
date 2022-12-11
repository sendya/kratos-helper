package tracing_test

import (
	"testing"

	"github.com/sendya/kratos-helper/contrib/tracing"
)

type testTracing struct {
	endpoint string
	customName string
}

func (r testTracing) GetEndpoint() string {
	return r.endpoint
}

func (r testTracing) GetCustomName() string {
	return r.customName
}

func TestInitTrace(t *testing.T) {
	trConf := &testTracing{
		endpoint: "http://localhost:14268/api/traces",
		customName: "myapp",
	}

	tracing.InitTracer(trConf)
}