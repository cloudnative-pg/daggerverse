// A generated module for ControllerGen functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"

	"dagger/controller-gen/internal/dagger"
)

type ControllerGen struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// ControllerGen to use.
	// +optional
	// renovate: datasource=github-tags depName=kubernetes-sigs/controller-tools versioning=semver
	// +default="0.16.2"
	controllerGenVersion string,
	// +optional
	// renovate image: datasource=docker depName=golang versioning=docker
	// +default="golang:1.24.1-bookworm"
	goImage string,
	// +optional
	ctr *dagger.Container,
) *ControllerGen {
	if ctr != nil {
		return &ControllerGen{
			Ctr: ctr,
		}
	}
	return &ControllerGen{
		Ctr: dag.Container().From(goImage).
			WithExec([]string{"go", "install",
				fmt.Sprintf("sigs.k8s.io/controller-tools/cmd/controller-gen@v%v", controllerGenVersion)}),
	}
}

func (m *ControllerGen) ControllerGen(
	ctx context.Context,
	// The directory of the repository.
	source *dagger.Directory,
	// +optional
	// A list of arguments to pass to controller-gen.
	args []string,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"controller-gen"}, args...))
}
