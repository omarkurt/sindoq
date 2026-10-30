//go:build microsandbox && cgo

package sindoq

import (
	"testing"

	"github.com/happyhackingspace/sindoq/internal/factory"
)

func TestMicrosandboxConfigReachesProvider(t *testing.T) {
	constructor, ok := factory.DefaultRegistry.GetConstructor("microsandbox")
	if !ok {
		t.Fatal("microsandbox provider should be registered")
	}
	if _, err := constructor(MicrosandboxConfig{DefaultImage: "alpine:3.19"}); err != nil {
		t.Errorf("public MicrosandboxConfig should be accepted, got %v", err)
	}
	if _, err := constructor(nil); err != nil {
		t.Errorf("nil config should use defaults, got %v", err)
	}
	if _, err := constructor(DockerConfig{}); err == nil {
		t.Error("unrelated config type should be rejected")
	}
}
