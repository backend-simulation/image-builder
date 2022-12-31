package docker

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
)

const dockerTar string = "dockerfile.tar"

func DockerBuild(dockerfilePath string, tag string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	tar := new(archivex.TarFile)
	tar.Create(dockerTar)
	tar.AddAll(dockerfilePath, false)
	tar.Close()

	dockerBuildContext, err := os.Open(dockerTar)
	defer os.RemoveAll(dockerTar)

	opt := types.ImageBuildOptions{
		Tags:    []string{tag},
		Context: dockerBuildContext,
	}

	imageBuildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, opt)
	if err != nil {
		panic(err)
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}

	containersPruneReport, err := cli.ContainersPrune(ctx, filters.Args{})
	if err != nil {
		log.Fatal(err, " :unable to prune container")
	}
	log.Printf("prune containers: %+v\n", containersPruneReport.ContainersDeleted)

	imagesPruneReport, err := cli.ImagesPrune(ctx, filters.Args{})
	if err != nil {
		log.Fatal(err, " :unable to prune image")
	}
	log.Printf("prune images: %+v\n", imagesPruneReport.ImagesDeleted)
}
