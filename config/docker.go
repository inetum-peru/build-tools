package config

import (
	"fmt"
)

const (
	targetImage string = "inetum-peru/build-tools"
)

// Docker struct field.
type Docker struct {
	TargetImage string
	Tag         string
	TagLatest   string
}

// ImageTag return image tag.
func (d *Docker) ImageTag() string {
	return fmt.Sprintf("%s:%s", d.TargetImage, d.Tag)
}

// ImageTagLatest return image tag latest.
func (d *Docker) ImageTagLatest() string {
	return fmt.Sprintf("%s:%s", d.TargetImage, d.TagLatest)
}
