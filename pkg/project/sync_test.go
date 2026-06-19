package project

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func TestSyncModule_RefreshesFromDiskWithoutLosingDirName(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "a-subject")
	if err := os.MkdirAll(dir, 0775); err != nil {
		t.Fatal(err)
	}

	onDisk := subject.Subject{UUID: "uuid-a", Type: subject.SubjectTask, Name: "edited-on-disk", DirName: dir}
	if err := onDisk.WriteMetadata(); err != nil {
		t.Fatal(err)
	}

	m := module.Module{
		Subjects: []subject.Subject{
			{UUID: "uuid-a", Type: subject.SubjectTask, Name: "stale-name", DirName: dir},
		},
	}

	var result SyncResult
	if err := syncModule(&m, true, &result); err != nil {
		t.Fatalf("syncModule() error = %v", err)
	}

	if len(m.Subjects) != 1 {
		t.Fatalf("got %d subjects, want 1", len(m.Subjects))
	}

	got := m.Subjects[0]
	if got.Name != "edited-on-disk" {
		t.Errorf("Name = %q, want %q (disk edit should be picked up)", got.Name, "edited-on-disk")
	}
	if got.DirName != dir {
		t.Errorf("DirName = %q, want %q (must not be wiped — it's excluded from YAML)", got.DirName, dir)
	}
	if len(result.Updated) != 1 {
		t.Errorf("result.Updated = %+v, want 1 entry", result.Updated)
	}
}

func TestSyncModule_MissingOrMalformedSubjectIsSkippedNotLost(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "corrupt-subject")
	if err := os.MkdirAll(dir, 0775); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, subject.METADATA_FILE), []byte("not: valid: yaml: : :"), 0664); err != nil {
		t.Fatal(err)
	}

	m := module.Module{
		Subjects: []subject.Subject{
			{UUID: "uuid-corrupt", Type: subject.SubjectTask, Name: "corrupt-subject", DirName: dir},
		},
	}

	var result SyncResult
	if err := syncModule(&m, true, &result); err != nil {
		t.Fatalf("syncModule() error = %v", err)
	}

	if len(m.Subjects) != 1 || m.Subjects[0].UUID != "uuid-corrupt" {
		t.Errorf("malformed metadata should leave the existing entry untouched, got: %+v", m.Subjects)
	}
	if len(result.Skipped) != 1 {
		t.Errorf("result.Skipped = %+v, want 1 entry", result.Skipped)
	}
}

func TestDiscoverModule_FindsNewSubjectAndAssignsUUID(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "new-subject")
	if err := os.MkdirAll(dir, 0775); err != nil {
		t.Fatal(err)
	}
	newSubj := subject.Subject{Type: subject.SubjectTask, Name: "new-subject", DirName: dir}
	if err := newSubj.WriteMetadata(); err != nil {
		t.Fatal(err)
	}

	// a flat, non-subject directory must be ignored, not erroring discovery
	if err := os.MkdirAll(filepath.Join(root, "flat-dir"), 0775); err != nil {
		t.Fatal(err)
	}
	// a directory with metadata but no valid type must be ignored
	noTypeDir := filepath.Join(root, "no-type")
	if err := os.MkdirAll(noTypeDir, 0775); err != nil {
		t.Fatal(err)
	}
	noType := subject.Subject{Name: "no-type", DirName: noTypeDir}
	if err := noType.WriteMetadata(); err != nil {
		t.Fatal(err)
	}

	m := module.Module{AbsPath: root}

	var result SyncResult
	if err := discoverModule(&m, true, &result); err != nil {
		t.Fatalf("discoverModule() error = %v", err)
	}

	if len(result.New) != 1 {
		t.Fatalf("result.New = %+v, want exactly 1 (flat-dir and no-type must be excluded)", result.New)
	}
	if len(m.Subjects) != 1 {
		t.Fatalf("got %d subjects indexed, want 1", len(m.Subjects))
	}
	if m.Subjects[0].UUID == "" {
		t.Error("newly discovered subject should have been assigned a UUID")
	}
}

func TestDiscoverModule_SkipsAlreadyIndexedByUUID(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "tracked-subject")
	if err := os.MkdirAll(dir, 0775); err != nil {
		t.Fatal(err)
	}
	tracked := subject.Subject{UUID: "uuid-tracked", Type: subject.SubjectTask, Name: "tracked-subject", DirName: dir}
	if err := tracked.WriteMetadata(); err != nil {
		t.Fatal(err)
	}

	m := module.Module{
		AbsPath:  root,
		Subjects: []subject.Subject{tracked},
	}

	var result SyncResult
	if err := discoverModule(&m, true, &result); err != nil {
		t.Fatalf("discoverModule() error = %v", err)
	}

	if len(result.New) != 0 {
		t.Errorf("result.New = %+v, want none — subject is already indexed by UUID", result.New)
	}
	if len(m.Subjects) != 1 {
		t.Errorf("got %d subjects, want 1 (no duplicate)", len(m.Subjects))
	}
}

func TestSync_DryRunDoesNotWriteChanges(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "new-subject")
	if err := os.MkdirAll(dir, 0775); err != nil {
		t.Fatal(err)
	}
	newSubj := subject.Subject{Type: subject.SubjectTask, Name: "new-subject", DirName: dir}
	if err := newSubj.WriteMetadata(); err != nil {
		t.Fatal(err)
	}

	p := &Project{
		absDir:  root,
		Modules: []module.Module{{AbsPath: root}},
	}

	result, err := Sync(p, false)
	if err != nil {
		t.Fatalf("Sync() error = %v", err)
	}

	if len(result.New) != 1 {
		t.Fatalf("result.New = %+v, want 1 (dry-run should still report what it would do)", result.New)
	}
	if len(p.Modules[0].Subjects) != 0 {
		t.Errorf("dry-run must not mutate the in-memory project: got %+v", p.Modules[0].Subjects)
	}
	if _, err := os.Stat(filepath.Join(root, METADATA_FILE)); err == nil {
		t.Error("dry-run must not write project METADATA.yml")
	}
}
