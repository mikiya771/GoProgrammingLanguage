package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	encType := flag.String("enc", "jpg", "valid variables: jpg, png, gif")
	var err error
	var kind string
	flag.Parse()
	switch *encType {
	case "jpg", "jpeg":
		kind, err = toJPG(os.Stdin, os.Stdout)
	case "png":
		kind, err = toPNG(os.Stdin, os.Stdout)
	case "gif":
		kind, err = toGIF(os.Stdin, os.Stdout)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "covert failed: error message is %s, evaluated image type %s /n", err, kind)
	}
}
func toJPG(in io.Reader, out io.Writer) (string, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return "unknown", err
	}

	return kind, jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
func toPNG(in io.Reader, out io.Writer) (string, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return "unknown", err
	}

	return kind, png.Encode(out, img)
}
func toGIF(in io.Reader, out io.Writer) (string, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return "unknown", err
	}

	return kind, gif.Encode(out, img, &gif.Options{Quantizer: nil, Drawer: nil})
}
