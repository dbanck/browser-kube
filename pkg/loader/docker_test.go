package loader

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestLoad(t *testing.T) {
	// appFS := afero.NewMemMapFs()
	// basePath := "/images"
	appFS := afero.NewOsFs()
	basePath := "/Users/danielschmidt/Downloads/afero"
	imageName := "danielmschmidt/hello-wasm:latest"

	loader := NewDockerImageLoader(appFS, basePath, []string{"/main.wasm"})
	bytes, err := loader.Load(context.Background(), imageName)
	if err != nil {
		t.Error(err.Error())
	}

	// Test side-effects
	expectManifest := fmt.Sprintf("%s/save/%s/manifest.json", basePath, imageName)
	_, err = appFS.Stat(expectManifest)

	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", expectManifest)
	}

	expectedContent := fmt.Sprintf("%s/content/%s/main.wasm", basePath, imageName)
	_, err = appFS.Stat(expectedContent)

	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", expectedContent)
	}

	// Test end-result
	if len(bytes) == 0 {
		t.Error("Bytes array is empty")
	}
}
