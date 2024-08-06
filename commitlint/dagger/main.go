// This module runs commitlint to validate conventional commits.
package main

import (
	"context"
	"main/internal/dagger"
)

type Commitlint struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// Commitlint image to use.
	// +optional
	// renovate image: datasource=docker depName=commitlint/commitlint versioning=docker
	// +default="commitlint/commitlint:19.3.1"
	Image string,
) *Commitlint {
	return &Commitlint{
		Ctr: dag.Container().From(Image),
	}
}

// Lint runs commitlint to lint commit messages.
//
// Example usage: dagger call lint --source /path/to/your/repo --args arg1 --args arg2
func (m *Commitlint) Lint(
	ctx context.Context,
	// The directory of the repository.
	source *dagger.Directory,
	// +optional
	// A list of arguments to pass to commitlint.
	args []string,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args, dagger.ContainerWithExecOpts{UseEntrypoint: true})
}
