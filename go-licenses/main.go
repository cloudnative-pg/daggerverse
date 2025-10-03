// A generated module for GoLicenses functions
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

	"dagger/go-licenses/internal/dagger"
)

type GoLicenses struct {
	// +private
	Ctr *dagger.Container
}

func New(
// GoLicenses to use.
// +optional
// renovate: datasource=github-tags depName=google/go-licenses versioning=semver
// +default="2.0.1"
	goLicensesVersion string,
// +optional
// renovate image: datasource=docker depName=golang versioning=docker
// +default="golang:1.25.1-alpine"
	goImage string,
// +optional
	ctr *dagger.Container,
) *GoLicenses {
	if ctr != nil {
		return &GoLicenses{
			Ctr: ctr,
		}
	}
	return &GoLicenses{
		Ctr: dag.Container().From(goImage).
			WithExec([]string{"go", "install",
				fmt.Sprintf("github.com/google/go-licenses@v%v", goLicensesVersion)}),
	}
}

func (m *GoLicenses) GoLicenses(
	ctx context.Context,
// The directory of the repository.
	source *dagger.Directory,
// +optional
// A list of arguments to pass to go-licenses.
	args []string,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"go-licenses"}, args...))
}
