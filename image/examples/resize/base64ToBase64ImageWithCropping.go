package example

import (
	"fmt"
	"github.com/micro/micro-go/image"
)

// Resize a base64 encoded image on the fly without storing it anywhere
func Base64toBase64imageWithCropping() {
	imageService := image.NewImageService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := imageService.Resize(image.ResizeRequest{
		Base64: "data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
		CropOptions: image.CropOptions{
			Height: 50,
			Width:  50,
		},
		Height: 100,
		Width:  100,
	})
	fmt.Println(rsp)
}
