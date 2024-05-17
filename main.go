package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	// Membuat konteks gambar baru dengan ukuran 800x600
	dc := gg.NewContext(800, 600)

	// Mengatur warna teks menjadi putih
	white := color.RGBA{255, 255, 255, 255}
	dc.SetColor(white)

	// Mengatur ukuran font
	fontPath := "arial.ttf" // Ganti dengan path font yang Anda miliki
	if err := dc.LoadFontFace(fontPath, 48); err != nil {
		panic(err)
	}

	// Teks yang akan diulang-ulang dalam watermark
	text := "satria"
	repetitions := 20

	// Menggambar teks secara berulang-ulang
	for i := 0; i < repetitions; i++ {
		for j := 0; j < repetitions; j++ {
			// Mengatur posisi dan sudut teks
			x := float64(j * 100)
			y := float64(i * 100)
			angle := 45.0

			// Menggambar teks dengan posisi dan sudut yang ditentukan
			dc.RotateAbout(gg.Radians(angle), x, y)
			dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
			dc.RotateAbout(-gg.Radians(angle), x, y)
		}
	}

	// Simpan gambar dengan watermark
	if err := dc.SavePNG("watermarked_image.png"); err != nil {
		panic(err)
	}

	fmt.Println("Watermark berhasil dibuat!")
}
