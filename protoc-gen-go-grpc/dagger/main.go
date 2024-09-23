// A generated module for ProtocGenGoGRPC functions

package main

import (
	"context"
	"fmt"
	"path"

	"dagger/protoc-gen-go-grpc/internal/dagger"
)

type ProtocGenGoGRPC struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// Custom image to use to run protoc.
	// +optional
	// renovate image: datasource=docker depName=golang versioning=docker
	// +default="golang:1.23.1-bookworm"
	goImage string,
	// +optional
	// renovate: datasource=github-tags depName=protocolbuffers/protobuf versioning="regex:^v?(?<major>\\d+)\\.(?<minor>\\d+)$"
	// +default="28.2"
	protobufVersion string,
	// +optional
	// renovate: datasource=go depName=google.golang.org/protobuf/cmd/protoc-gen-go versioning=semver
	// +default="v1.33.0"
	protocGenGoVersion string,
	// +optional
	// renovate: datasource=go depName=google.golang.org/grpc/cmd/protoc-gen-go-grpc versioning=semver
	// +default="v1.3.0"
	protocGenGoGRPCVersion string,
) *ProtocGenGoGRPC {
	protobufRelURL := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%v/protoc-%v-linux-x86_64.zip",
		protobufVersion, protobufVersion)

	protobuf := dag.Container().
		From(goImage).
		WithExec([]string{"apt", "update"}).
		WithExec([]string{"apt", "install", "-y", "unzip"}).
		WithExec([]string{"curl", "-LO", protobufRelURL}).
		WithExec([]string{"unzip", "protoc-*.zip", "-d", "/usr/local"}).
		WithExec([]string{"rm", "-rf", "protoc-*.zip"}).
		WithExec([]string{"apt", "purge", "-y", "unzip"}).
		WithExec([]string{"rm", "-rf", "/var/lib/apt/lists/*"}).
		WithExec([]string{"go", "install",
			fmt.Sprintf("google.golang.org/protobuf/cmd/protoc-gen-go@%v", protocGenGoVersion)}).
		WithExec([]string{"go", "install",
			fmt.Sprintf("google.golang.org/grpc/cmd/protoc-gen-go-grpc@%v", protocGenGoGRPCVersion)})

	return &ProtocGenGoGRPC{
		Ctr: protobuf,
	}
}

// Container get the current container
func (m *ProtocGenGoGRPC) Container() *dagger.Container {
	return m.Ctr
}

// Run runs protoc on proto files, returning the generated go files as a directory.
//
//	Example: dagger call run --source . \
//	    --go-opt module=github.com/cloudnative-pg/cnpg-i \
//	    --go-grpcopt module=github.com/cloudnative-pg/cnpg-i \
//	    --proto-path proto -o .
func (m *ProtocGenGoGRPC) Run(
	ctx context.Context,
	// The source directory.
	source *dagger.Directory,
	// The path to the proto files, relative to the source directory.
	protoPath string,
	// go_opt flag to pass to protoc.
	goOpt string,
	// go-grpc_opt flag to pass to protoc.
	goGRPCOpt string,
) (*dagger.Directory, error) {
	args := []string{"/usr/local/bin/protoc"}
	args = append(args, "--go_out=/out/")
	args = append(args, fmt.Sprintf("--go_opt=%v", goOpt))
	args = append(args, "--go-grpc_out=/out/")
	args = append(args, fmt.Sprintf("--go-grpc_opt=%v", goGRPCOpt))
	protos, err := source.Directory(protoPath).Entries(ctx)
	if err != nil {
		return nil, err
	}
	for i := range protos {
		protos[i] = path.Join(protoPath, protos[i])
	}
	args = append(args, protos...)

	buildDir := m.Ctr.
		WithMountedDirectory("/src", source).
		WithExec([]string{"mkdir", "-p", "/out"}).
		WithWorkdir("/src").
		WithExec(args).
		Directory("/out")
	return buildDir, nil
}
