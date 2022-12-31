package docker

import "testing"

func TestDockerBuild(t *testing.T) {
	t.Log("start docker build test")
	DockerBuild("/Users/bron.kang/Projects/complex/sample-codes/java/17/springboot/3.0/plus-two-number/", "plus-two-number:0.1")
}
