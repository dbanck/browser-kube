package loader

import (
	"context"
	"fmt"
	"os"
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

	// Test side-effects
	expectManifest := fmt.Sprintf("%s/save/%s/manifest.json", basePath, imageName)
	_, err = appFS.Stat(expectManifest)

	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", expectManifest)
	}

	expectedContent := fmt.Sprintf("%s/content/%s/wasm_bg.wasm", basePath, imageName)
	_, err = appFS.Stat(expectedContent)

	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", expectedContent)
	}

	// Test end-result
	if len(bytes) == 0 {
		t.Error("Bytes array is empty")
	}
}
