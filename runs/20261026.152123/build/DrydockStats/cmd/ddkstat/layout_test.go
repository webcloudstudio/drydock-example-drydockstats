package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// moduleRoot is the application root, two directories above this package.
const moduleRoot = "../.."

func TestGoModDeclaresModuleAndGoVersion(t *testing.T) {
	data, err := os.ReadFile(moduleRoot + "/go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	var module, goVersion string
	var requires int
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "module":
			module = fields[1]
		case "go":
			goVersion = fields[1]
		case "require":
			requires++
		}
	}
	if module != "ddkstat" {
		t.Errorf("module = %q, want %q", module, "ddkstat")
	}
	var major, minor int
	if _, err := fmt.Sscanf(goVersion, "%d.%d", &major, &minor); err != nil {
		t.Fatalf("go directive %q is not major.minor: %v", goVersion, err)
	}
	if major < 1 || (major == 1 && minor < 22) {
		t.Errorf("go directive = %q, want 1.22 or newer", goVersion)
	}
	if requires != 0 {
		t.Errorf("found %d require directives, want none", requires)
	}
}

func TestLayoutDirectoriesExist(t *testing.T) {
	for _, dir := range []string{"cmd/ddkstat", "internal", "bin"} {
		info, err := os.Stat(moduleRoot + "/" + dir)
		if err != nil {
			t.Errorf("stat %s: %v", dir, err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("%s is not a directory", dir)
		}
	}
}

func TestNoVendorOrDependencySums(t *testing.T) {
	if _, err := os.Stat(moduleRoot + "/vendor"); err == nil {
		t.Error("vendor directory exists, want none")
	}
	sum, err := os.ReadFile(moduleRoot + "/go.sum")
	if err == nil && strings.TrimSpace(string(sum)) != "" {
		t.Errorf("go.sum is not empty: %q", sum)
	}
}
