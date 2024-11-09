package assets

import (
    resource "github.com/quasilyte/ebitengine-resource"
    _ "image/png"
)

const (
    ImageNone resource.ImageID = iota
    ImageWitch
)

func registerImageResources(loader *resource.Loader) {
    imageResources := map[resource.ImageID]resource.ImageInfo{
        ImageWitch: {Path: "images/witch.png"},
    }

    for id, res := range imageResources {
        loader.ImageRegistry.Set(id, res)
    }
}