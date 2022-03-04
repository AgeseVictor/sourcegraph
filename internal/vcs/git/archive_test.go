package git

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/authz"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	"github.com/sourcegraph/sourcegraph/internal/types"
)

func TestArchiveReaderForRepoWithSubRepoPermissionsFiltersFile(t *testing.T) {
	ctx := actor.WithActor(context.Background(), &actor.Actor{
		UID: 1,
	})
	repoName := MakeGitRepository(t,
		"echo abcd > file1",
		"git add file1",
		"GIT_COMMITTER_NAME=a GIT_COMMITTER_EMAIL=a@a.com GIT_COMMITTER_DATE=2006-01-02T15:04:05Z git commit -m commit1 --author='a <a@a.com>' --date 2006-01-02T15:04:05Z",
		"echo foo > file2",
		"git add file2",
		"GIT_COMMITTER_NAME=a GIT_COMMITTER_EMAIL=a@a.com GIT_COMMITTER_DATE=2006-01-02T15:04:06Z git commit -m commit1 --author='a <a@a.com>' --date 2006-01-02T15:04:06Z",
	)
	const commitID = "3d689662de70f9e252d4f6f1d75284e23587d670"

	checker := authz.NewMockSubRepoPermissionChecker()
	checker.EnabledFunc.SetDefaultHook(func() bool {
		return true
	})
	checker.PermissionsFunc.SetDefaultHook(func(ctx context.Context, i int32, content authz.RepoContent) (authz.Perms, error) {
		if content.Path == "file2" {
			return authz.None, nil
		}
		return authz.Read, nil
	})

	repo := &types.Repo{Name: repoName, ID: 1}

	opts := gitserver.ArchiveOptions{
		Format:  ArchiveFormatZip,
		Treeish: commitID,
		Paths:   []string{"."},
	}
	readCloser, err := ArchiveReader(ctx, checker, repo.Name, opts)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer readCloser.Close()
	validateFilesInZipArchive(readCloser, t, []string{"file1"})
}

func validateFilesInZipArchive(rc io.ReadCloser, t *testing.T, expectedFiles []string) {
	t.Helper()
	buff := bytes.NewBuffer([]byte{})
	size, err := io.Copy(buff, rc)
	if err != nil {
		t.Fatalf("error copying file contents: %s", err)
	}
	reader := bytes.NewReader(buff.Bytes())
	zipReader, err := zip.NewReader(reader, size)
	if err != nil {
		t.Fatalf("error creating zip reader: %s", err)
	}
	if len(zipReader.File) != len(expectedFiles) {
		t.Errorf("expected zip archive to have %d files, got %d files instead", len(expectedFiles), len(zipReader.File))
	}
	for _, zf := range zipReader.File {
		match := false
		for _, ef := range expectedFiles {
			if zf.Name == ef {
				match = true
			}
		}
		if match == false {
			t.Errorf("zip archive missing file: %s", zf.Name)
		}
	}
}

func TestArchiveReaderForRepoWithoutSubRepoPermissions(t *testing.T) {
	ctx := actor.WithActor(context.Background(), &actor.Actor{
		UID: 1,
	})
	repoName := MakeGitRepository(t,
		"echo abcd > file1",
		"git add file1",
		"GIT_COMMITTER_NAME=a GIT_COMMITTER_EMAIL=a@a.com GIT_COMMITTER_DATE=2006-01-02T15:04:05Z git commit -m commit1 --author='a <a@a.com>' --date 2006-01-02T15:04:05Z",
	)
	const commitID = "3d689662de70f9e252d4f6f1d75284e23587d670"

	checker := authz.NewMockSubRepoPermissionChecker()
	checker.EnabledFunc.SetDefaultHook(func() bool {
		return true
	})

	repo := &types.Repo{Name: repoName, ID: 1}

	opts := gitserver.ArchiveOptions{
		Format:  ArchiveFormatZip,
		Treeish: commitID,
		Paths:   []string{"."},
	}
	readCloser, err := ArchiveReader(ctx, checker, repo.Name, opts)
	if err != nil {
		t.Fatalf("Error should not be thrown because ArchiveReader is invoked for a repo without sub-repo permissions, got error: %s", err)
	}
	err = readCloser.Close()
	if err != nil {
		t.Error("Error during closing a reader")
	}
}
