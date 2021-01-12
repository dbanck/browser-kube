package loader

import (
	"context"
	"testing"

	"github.com/spf13/afero"
)

func TestLoad(t *testing.T) {
	appFS := afero.NewMemMapFs()
	basePath := "/images"
	imageName := "danielmschmidt/hello-wasm:latest"

	loader := NewDockerImageLoader(appFS, basePath, []string{"/wasm_bg.wasm"})
	bytes, err := loader.Load(context.Background(), imageName)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Test end-result
	if len(bytes) == 0 {
		t.Error("Bytes array is empty")
	}
}
