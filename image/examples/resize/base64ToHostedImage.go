package example

import (
	"fmt"
	"github.com/micro/micro-go/image"
)

// Resize an input base64 encoded image and store the resulting resized image on our CDN.
func Base64toHostedImage() {
	imageService := image.NewImageService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := imageService.Resize(image.ResizeRequest{
		Base64: "data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
		Height: 100,
		Name:   "cat.png",
		Width:  100,
	})
	fmt.Println(rsp)
}
