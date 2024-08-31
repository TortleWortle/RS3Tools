package main

import (
	"image/png"
	"os"
	"testing"
)

func TestCompareSucceeds(t *testing.T) {
	active, err := os.Open("input/wise-active.png")
	if err != nil {
		t.Fatalf("opening active: %v", err)
	}

	buff, err := os.Open("input/wise-buff.png")
	if err != nil {
		t.Fatalf("opening buff: %v", err)
	}

	activeImg, err := png.Decode(active)
	if err != nil {
		t.Fatalf("decoding active: %v", err)
	}
	buffImg, err := png.Decode(buff)
	if err != nil {
		t.Fatalf("decoding buff: %v", err)
	}

	percent, err := overlapPercent(activeImg, buffImg)
	if err != nil {
		t.Fatalf("calc overlap: %v", err)
	}
	if percent < 0.99 {
		t.Errorf("overlap percent is less than 0.99: %f", percent)
	}
}

func TestCompareFails(t *testing.T) {
	active, err := os.Open("input/bone-shield.png")
	if err != nil {
		t.Fatalf("opening active: %v", err)
	}

	buff, err := os.Open("input/wise-buff.png")
	if err != nil {
		t.Fatalf("opening buff: %v", err)
	}

	activeImg, err := png.Decode(active)
	if err != nil {
		t.Fatalf("decoding active: %v", err)
	}
	buffImg, err := png.Decode(buff)
	if err != nil {
		t.Fatalf("decoding buff: %v", err)
	}

	percent, err := overlapPercent(activeImg, buffImg)
	if err != nil {
		t.Fatalf("calc overlap: %v", err)
	}

	if percent > 0.9 {
		t.Errorf("overlap percent is over 0.9: %f", percent)
	}
}
