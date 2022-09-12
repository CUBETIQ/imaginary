package main

// Version stores the current package semantic version
var Version = "dev"

var CustomBuild = "cubetiqs.com"

// Versions represents the used versions for several significant dependencies
type Versions struct {
	ImaginaryVersion string `json:"imaginary"`
	BimgVersion      string `json:"bimg"`
	VipsVersion      string `json:"libvips"`
	CustomBuild      string `json:"custom_build"`
}
