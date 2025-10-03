package main

import (
	"context"
	"main/internal/dagger"
)

type Protolint struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// Protolint image to use.
	// +optional
	// renovate image: datasource=docker depName=yoheimuta/protolint versioning=docker
	// +default="yoheimuta/protolint:0.56.4"
	Image string,
) *Protolint {
	return &Protolint{
		Ctr: dag.Container().From(Image),
	}
}

// Lint runs protolint on proto files.
//
// Example usage: dagger call lint --source /path/ --args "-config_path=.protolint.yaml" --args .
func (m *Protolint) Lint(
	ctx context.Context,
	// The directory of the repository.
	source *dagger.Directory,
	// A list of arguments to pass to commitlint.
	// +optional
	args []string,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"lint"}, args...), dagger.ContainerWithExecOpts{UseEntrypoint: true})
}
