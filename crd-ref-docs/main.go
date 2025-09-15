// A Dagger module for generating CRD reference documentation using the crd-ref-docs tool

package main

import (
	"context"
	"fmt"

	"dagger/crd-ref-docs/internal/dagger"
)

type CrdRefDocs struct {
	// +private
	Ctr *dagger.Container
}

type Renderer string

const (
	// Asciidoc is the Asciidoc renderer.
	Asciidoc Renderer = "asciidoctor"
	// Markdown is the Markdown renderer.
	Markdown Renderer = "markdown"
)

type OutputMode string

const (
	// Group is the group output mode.
	Group OutputMode = "group"
	// Single is the single output mode.
	Single OutputMode = "single"
)

func New(
	ctx context.Context,
// Go image to use.
// +optional
// renovate image: datasource=docker depName=golang versioning=docker
// +default="golang:1.24.5-alpine"
	Image string,
// CrdRefDocs version to use.
// +optional
// +default="master"
	Version string,
) *CrdRefDocs {
	ctr := dag.Container().From(Image).
		WithExec([]string{"go", "install",
			fmt.Sprintf("github.com/elastic/crd-ref-docs@%v", Version)})
	return &CrdRefDocs{
		Ctr: ctr,
	}
}

func (m *CrdRefDocs) Generate(
	ctx context.Context,
// The directory of the sources
	src *dagger.Directory,
// The path of the CRD files, relative to the source directory.
	sourcePath string,
// +optional
// The path of the config file, relative to the source directory.
	configFile string,
// +optional
// The path of the template director, relative to the source directory.
	templatesDir string,
// +optional
// +default="asciidoctor"
// The renderer for the generated documentation.
	renderer Renderer,
// +optional
// +default="single"
// The output mode for the generated documentation.
	outputMode OutputMode,
// +optional
// +default="INFO"
// Log level.
	logLevel string,
// +optional
// +default=10
// Maximum recursion level for type discovery.
	maxDepth int,
// +optional
// Output path for the generated documentation.
	outputPath string,
) *dagger.Container {
	command := []string{"crd-ref-docs",
		"--log-level", logLevel,
		"--max-depth", fmt.Sprintf("%d", maxDepth),
		"--output-mode", string(outputMode),
		"--renderer", string(renderer),
		"--source-path", sourcePath,
	}
	ctr := m.Ctr.WithMountedDirectory("/src", src).WithWorkdir("/src")
	if configFile != "" {
		command = append(command, "--config", configFile)
	}
	if templatesDir != "" {
		command = append(command, "--templates-dir", templatesDir)
	}
	if outputPath != "" {
		command = append(command, "--output-path", outputPath)
	}
	return ctr.WithExec(command)
}
