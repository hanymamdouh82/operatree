package project

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func writeSubjectMeta(t *testing.T, dir, body string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "METADATA.yml"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestHydrateFromDisk_AddsNewSubjectByFolderName(t *testing.T) {
	root := t.TempDir()
	mod := filepath.Join(root, "01_EVENTS")
	writeSubjectMeta(t, filepath.Join(mod, "alpha"), "uuid: u1\ntype: EVENT\nname: stored\n")
	if err := os.MkdirAll(filepath.Join(mod, "not-a-subject"), 0o755); err != nil {
		t.Fatal(err)
	}

	p := Project{Modules: []module.Module{{Name: "01_EVENTS", AbsPath: mod}}}
	hydrateFromDisk(&p)

	got := p.Modules[0].Subjects
	if len(got) != 1 {
		t.Fatalf("got %d subjects, want 1", len(got))
	}
	if got[0].UUID != "u1" || got[0].Name != "alpha" || got[0].DirName != filepath.Join(mod, "alpha") {
		t.Errorf("unexpected subject %+v", got[0])
	}
}

func TestHydrateFromDisk_RefreshesExistingByUUID(t *testing.T) {
	root := t.TempDir()
	mod := filepath.Join(root, "01_EVENTS")
	dir := filepath.Join(mod, "alpha")
	writeSubjectMeta(t, dir, "uuid: u1\ntype: EVENT\nname: stored\nstatus: done\n")

	p := Project{Modules: []module.Module{{
		Name: "01_EVENTS", AbsPath: mod,
		Subjects: []subject.Subject{{UUID: "u1", Type: subject.SubjectEvent, Name: "alpha", Status: "todo", DirName: dir}},
	}}}
	hydrateFromDisk(&p)

	got := p.Modules[0].Subjects
	if len(got) != 1 {
		t.Fatalf("got %d subjects, want 1 (refreshed, not duplicated)", len(got))
	}
	if got[0].Status != "done" {
		t.Errorf("Status = %q, want done (refreshed from disk)", got[0].Status)
	}
}

func TestHydrateFromDisk_KeepsSubjectWhoseFolderIsGone(t *testing.T) {
	root := t.TempDir()
	mod := filepath.Join(root, "01_EVENTS")
	if err := os.MkdirAll(mod, 0o755); err != nil {
		t.Fatal(err)
	}

	p := Project{Modules: []module.Module{{
		Name: "01_EVENTS", AbsPath: mod,
		Subjects: []subject.Subject{{UUID: "u1", Type: subject.SubjectEvent, Name: "gone", DirName: filepath.Join(mod, "gone")}},
	}}}
	hydrateFromDisk(&p)

	got := p.Modules[0].Subjects
	if len(got) != 1 || got[0].UUID != "u1" {
		t.Fatalf("vanished-folder subject must be kept, got %+v", got)
	}
}
