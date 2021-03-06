package goma

import (
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
)

type ComposeImageRequest struct {
	img image.Image
	x   int
	y   int
}

func ComposeRGBA(width int, height int, requests []*ComposeImageRequest) image.Image {
	canvas := image.NewNRGBA(image.Rect(0, 0, width, height))

	for _, req := range requests {
		canvas = imaging.Overlay(canvas, req.img, image.Pt(req.x, req.y), 1.0)
	}

	return canvas
}

func ComposeCMYK(width int, height int, requests []*ComposeImageRequest) image.Image {
	canvas := image.NewCMYK(image.Rect(0, 0, width, height))

	for _, req := range requests {
		dstRect := image.Rectangle{image.Pt(req.x, req.y), image.Pt(req.x+req.img.Bounds().Max.X, req.y+req.img.Bounds().Max.Y)}
		draw.Draw(canvas, dstRect, req.img, image.Pt(0, 0), draw.Over)
	}

	return canvas
}
