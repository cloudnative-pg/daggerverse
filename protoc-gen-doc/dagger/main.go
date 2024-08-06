package main

import (
	"fmt"
	"main/internal/dagger"
)

type ProtocGenDoc struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// ProtocGenDoc image to use.
	// +optional
	// renovate image: datasource=docker depName=pseudomuto/protoc-gen-doc versioning=docker
	// +default="pseudomuto/protoc-gen-doc:1.5"
	Image string,
) *ProtocGenDoc {
	return &ProtocGenDoc{
		Ctr: dag.Container().From(Image),
	}
}

// Generate runs protoc-gen-doc on proto files, returning the generated documentation as a directory.
//
// Example usage: dagger call generate --proto-dir /path/ --doc-opt "markdown,docs.md"
func (m *ProtocGenDoc) Generate(
	// The directory of the proto files.
	protoDir *dagger.Directory,
	// +optional
	// +default="markdown,docs.md"
	// The doc_opt flag to pass to protoc-gen-doc.
	docOpt string,
) *dagger.Directory {
	const outDir = "/out"

	return m.Ctr.
		WithMountedDirectory("/protos", protoDir).
		WithExec([]string{"mkdir", outDir}).
		WithExec([]string{fmt.Sprintf("--doc_opt=%v", docOpt)}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Directory(outDir)
}
