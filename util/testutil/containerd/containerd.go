package containerd

import (
	"context"
	"testing"
	"time"

	containerd "github.com/containerd/containerd/v2/client"
)

func GetVersion(t *testing.T, cdAddress string) string {
	t.Helper()

	cdClient, err := containerd.New(cdAddress, containerd.WithTimeout(60*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	defer cdClient.Close()
	ctx := context.TODO()
	cdVersion, err := cdClient.Version(ctx)
	if err != nil {
		t.Fatal(err)
	}
	return cdVersion.Version
}
