package loader

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// Manifest resembles a docker Manifest file in an abbreviated form
type Manifest struct {
	Layers []string
}

// UnpackTar takes a afero filesystem and a reader; a tar reader loops over the tarfile
// creating the file structure in the fs along the way, and writing any files
// From https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func UnpackTar(fs *afero.Afero, dst string, r io.Reader) error {
	tr := tar.NewReader(r)
	err := fs.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := fs.Stat(target); err != nil {
				if err := fs.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			buf := new(bytes.Buffer)
			buf.ReadFrom(tr)
			fileContent := buf.Bytes()
			err := fs.WriteFile(target, fileContent, 0644)
			if err != nil {
				return err
			}
		}
	}
}

// Loader exposes functions to load WASM binaries from
type Loader interface {
	// Loads a resource to the file system and returns content
	Load(ctx context.Context, id string) ([]byte, error)
}

// DockerImageLoader implements the Loader interface for docker images
type DockerImageLoader struct {
	fs                     *afero.Afero
	basePath               string
	pathsInsideDockerImage []string
}

// NewDockerImageLoader creates a new DockerImageLoader
func NewDockerImageLoader(fs afero.Fs, basePath string, pathsInsideDockerImage []string) *DockerImageLoader {
	return &DockerImageLoader{fs: &afero.Afero{Fs: fs}, pathsInsideDockerImage: pathsInsideDockerImage, basePath: basePath}
}

// Load loads an image resource to the file system
// Inspired by https://www.madebymikal.com/quick-hack-extracting-the-contents-of-a-docker-image-to-disk/
func (l *DockerImageLoader) Load(ctx context.Context, imageName string) ([][]byte, error) {
	// Ensure base path exists
	l.fs.MkdirAll(l.basePath, 0755)
	l.fs.MkdirAll(filepath.Join(l.basePath, "save"), 0755)
	l.fs.MkdirAll(filepath.Join(l.basePath, "content"), 0755)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, errors.Wrap(err, "Could initialize docker cli")
	}

	// First we need to pull the image from the registry
	responseBody, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{All: true})
	if err != nil {
		return nil, errors.Wrap(err, "Could not load image")
	}
	// We need to consume the reader for the reference to be there
	buf := new(strings.Builder)
	io.Copy(buf, responseBody)
	fmt.Println("Loading docker image")
	fmt.Println(buf.String())

	responseBody.Close()

	// Now that we have the image we need to use `docker save` to get it as a tar
	reader, err := cli.ImageSave(ctx, []string{imageName})
	if err != nil {
		return nil, errors.Wrap(err, "Could not save image")
	}

	// we extract the tar.gz file as seen in
	imagePath := filepath.Join(l.basePath, "save", imageName)
	err = UnpackTar(l.fs, imagePath, reader)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unpack image")
	}

	// Get the manifest.json
	manifestFile, err := l.fs.ReadFile(filepath.Join(imagePath, "manifest.json"))
	if err != nil {
		return nil, errors.Wrap(err, "Could not find manifest.json in unpacked image")
	}

	var manifests []Manifest
	err = json.Unmarshal(manifestFile, &manifests)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unmarshal manifest.json in unpacked image")
	}

	if len(manifests) < 1 {
		return nil, errors.Wrap(err, "No manifests found in manifest.json")
	}

	contentPath := filepath.Join(l.basePath, "content", imageName)
	for _, layer := range manifests[0].Layers {
		layerContent, err := l.fs.ReadFile(filepath.Join(imagePath, layer))

		if err != nil {
			return nil, errors.Wrap(err, "Could not find layer from manifest.json")
		}
		err = UnpackTar(l.fs, contentPath, bytes.NewReader(layerContent))

		if err != nil {
			return nil, errors.Wrap(err, "Could not unpack layer from manifest.json")
		}
	}

	// we read the requested files
	files := [][]byte{}
	for _, filePath := range l.pathsInsideDockerImage {
		fileContent, err := l.fs.ReadFile(filepath.Join(contentPath, filePath))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Could not find requested file: %q", filePath))
		}
		files = append(files, fileContent)
	}

	return files, nil
}

// removeTmpFiles removes tmp files created by LOAD
func (l *DockerImageLoader) removeTmpFiles(ctx context.Context, imageName string) error {
	println("TODO: implement loader cleanup")
	return nil
}
