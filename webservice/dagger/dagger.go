// A generated module for Go functions
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

package dagger

import (
	"context"
	"dagger/go/internal/dagger"
	"fmt"
)

type Go struct{}

const postgresPassword = "password"
const postgresSvc = "postgres"

func (m *Go) Postgres(ctx context.Context, src *dagger.Directory) *dagger.Service {
	pg := dag.Container().
		From("postgres:17.2").
		WithEnvVariable("POSTGRES_PASSWORD", postgresPassword).
		WithExposedPort(5432).
		AsService()

	s, _ := dag.Container().
		From("flyway/flyway:11.2.0-alpine").
		WithMountedDirectory("/flyway/project", src).
		WithServiceBinding("postgres", pg).
		WithExec(
			[]string{`-workingDirectory="project"`,
				"-url=jdbc:postgresql://postgres:5432/postgres?user=postgres&password=password",
				"migrate",
			},
			dagger.ContainerWithExecOpts{UseEntrypoint: true},
		).
		Stdout(ctx)
	fmt.Println("flyway migrations result", s)

	return pg
}

func (m *Go) Redis() *dagger.Service {
	return dag.Container().
		From("redis:7.4.2").
		WithExposedPort(6379).
		AsService()
}

func (m *Go) Rebrow(svc *dagger.Service) *dagger.Service {
	return dag.Container().
		From("marian/rebrow:latest").
		WithExposedPort(5001).
		WithServiceBinding("redis", svc).
		AsService()
}

func (m *Go) Psql(svc *dagger.Service) *dagger.Container {
	return dag.Container().
		From("postgres:17.2").
		WithEnvVariable("PGPASSWORD", postgresPassword).
		WithServiceBinding("postgres", svc).
		Terminal(dagger.ContainerTerminalOpts{Cmd: []string{"psql", "-U", "postgres", "-h", "postgres"}})
}

func (m *Go) RedisCLI(svc *dagger.Service) *dagger.Container {
	return dag.Container().
		From("redis").
		WithServiceBinding("redis", svc).
		WithEntrypoint([]string{"redis-cli", "-h", "redis"}).
		Terminal()
}

func (m *Go) Build(src *dagger.Directory) *dagger.Container {
	build := dag.Container().
		From("golang:1.23.5").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod")).
		WithMountedCache("/root/.cache/go-build", dag.CacheVolume("go-build")).
		WithWorkdir("/work/src").
		WithFiles(".", []*dagger.File{src.File("go.mod"), src.File("go.sum")}).
		WithExec([]string{"go", "mod", "download"}).
		WithDirectory(".", src).
		WithExec([]string{"go", "build", "-o", "main", "./cmd"})

	return dag.Container().
		From("alpine:3.21.2").
		WithWorkdir("/app").
		WithFile("./main", build.File("./main")).
		WithEntrypoint([]string{"./main"})
}

func (m *Go) Run(src *dagger.Directory) *dagger.Service {
	return m.Build(src).
		WithExposedPort(8080).
		AsService()
}
