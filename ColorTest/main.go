package main

import (
	"github.com/michelle083/img_mod/Colors"
	"github.com/michelle083/img_mod/GetPic"
	"github.com/michelle083/img_mod/Grayscale"
	"github.com/michelle083/img_mod/Text"
)

func main() {
	GetPic.GetPic()
	Colors.Colors("downloaded_image.jpg")
	Grayscale.Grayscale("downloaded_image.jpg")
	Text.Text()
}