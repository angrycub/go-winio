package fs

import (
	"os"
	"testing"
)

func TestGetFSTypeOfKnownDrive(t *testing.T) {
	fsType, err := GetFileSystemType("C:\\")
	if err != nil {
		t.Fatal(err)
	}

	if fsType == "" {
		t.Fatal("No filesystem type name returned")
	}
}

func TestGetFSTypeOfInvalidPath(t *testing.T) {
	// [filepath.VolumeName] doesn't mandate that the drive letters matches [a-zA-Z].
	// Instead, use non-character drive.
	_, err := GetFileSystemType(`No:\`)
	if err != ErrInvalidPath {
		t.Fatalf("Expected `ErrInvalidPath`, got %v", err)
	}
}

func TestGetFSTypeOfValidButAbsentDrive(t *testing.T) {
	drive := ""
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		possibleDrive := string(letter) + ":\\"
		if _, err := os.Stat(possibleDrive); os.IsNotExist(err) {
			drive = possibleDrive
			break
		}
	}
	if drive == "" {
		t.Skip("Every possible drive exists")
	}

	_, err := GetFileSystemType(drive)
	if err == nil {
		t.Fatalf("GetFileSystemType %s unexpectedly succeeded", drive)
	}
	if !os.IsNotExist(err) {
		t.Fatalf("GetFileSystemType %s failed with %v, expected 'ErrNotExist' or similar", drive, err)
	}
}
