package main

import "fmt"

type ProtocGenDoc struct {
	// +private
	Ctr *Container
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
// Example usage: dagger call run --proto-dir /path/ --doc-opt "markdown,docs.md"
func (m *ProtocGenDoc) Generate(
	// The directory of the proto files.
	protoDir *Directory,
	// +optional
	// +default="markdown,docs.md"
	// The doc_opt flag to pass to protoc-gen-doc.
	docOpt string,
) *Directory {
	const outDir = "/out"

	return m.Ctr.
		WithMountedDirectory("/protos", protoDir).
		WithExec([]string{"mkdir", outDir}, ContainerWithExecOpts{SkipEntrypoint: true}).
		WithExec([]string{fmt.Sprintf("--doc_opt=%v", docOpt)}).
		Directory(outDir)
}
