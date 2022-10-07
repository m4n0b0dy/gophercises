package main

import (
"testing"
"os"
)


func TestMatchAndReplacePath(t *testing.T) {
	p := MyFilePath{Path:"testing.txt"}
	p.MatchAndReplacePath("test", "vest")
	if p.NewPath != "vesting.txt" {
		t.Fatalf("didnt work")
	}
}


func TestWalkDirectory(t *testing.T) {
	d := "testing/"
	os.RemoveAll(d)
	CopyDir("bckp", d)
	err := WalkDirectory(d, "n", "m")
	if err != nil {
		t.Fatalf(`%v`, err)
	}
	os.RemoveAll(d)
}