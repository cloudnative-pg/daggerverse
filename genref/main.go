// A generated module for Genref functions
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

	"dagger/genref/internal/dagger"
)

type Genref struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// genref version to use.
	// +optional
	// renovate: datasource=go depName=github.com/kubernetes-sigs/reference-docs/genref versioning=semver
	// +default="015aaac611407c4fe591bc8700d2c67b7521efca"
	genrefVersion string,
	// +optional
	// renovate image: datasource=docker depName=golang versioning=docker
	// +default="golang:1.24.1-bookworm"
	goImage string,
	// +optional
	ctr *dagger.Container,
) *Genref {
	if ctr != nil {
		return &Genref{
			Ctr: ctr,
		}
	}
	return &Genref{
		Ctr: dag.Container().From(goImage).
			WithExec([]string{"go", "install",
				fmt.Sprintf("github.com/kubernetes-sigs/reference-docs/genref@%v", genrefVersion)}),
	}
}

func (m *Genref) Genref(
	ctx context.Context,
	// The directory of the repository.
	source *dagger.Directory,
	// +optional
	// +default="docs"
	// The directory from where the command should run
	workDir string,
	// +optional
	// A list of arguments to pass to genref.
	args []string,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir(fmt.Sprintf("/src/%s", workDir)).
		WithExec(append([]string{"genref"}, args...))
}
