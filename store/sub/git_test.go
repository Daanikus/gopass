package sub

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGit(t *testing.T) {
	ctx := context.Background()

	tempdir, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Failed to create tempdir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()

	s, err := createSubStore(tempdir)
	assert.NoError(t, err)

	t.Skip("flaky")
	assert.NoError(t, s.GitInit(ctx, "", "", ""))
	assert.NoError(t, s.Git(ctx, "status"))
}