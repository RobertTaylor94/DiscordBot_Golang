package utility

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func GetDiceImage(id string, rolls []Roll) (*discordgo.File, error) {
	paths := make([]string, 0)
	files := make([]*os.File, 0)
	images := make([]image.Image, 0)

	for _, roll := range rolls {
		path := getFilePath(roll)
		paths = append(paths, path)
	}

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		files = append(files, file)
	}

	for _, file := range files {
		image, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}
		images = append(images, image)
	}

	totalWidth := 0
	totalHeight := 0

	for _, image := range images {
		totalWidth += image.Bounds().Dx()
		totalHeight = image.Bounds().Dy()
	}

	rgba := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	currentX := 0
	for _, img := range images {
		r := image.Rectangle{image.Point{currentX, 0}, image.Point{currentX + img.Bounds().Dx(), totalHeight}}
		draw.Draw(rgba, r, img, image.Point{0, 0}, draw.Src)
		currentX += img.Bounds().Dx()
	}

	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Printf("error creating out.png: %v", err)
	}

	png.Encode(out, rgba)

	imgPath := filepath.Join("output.png")
	file, _ := os.Open(imgPath)

	dsFile := &discordgo.File{
		Name:        "image.png",
		ContentType: "image/png",
		Reader:      file,
	}

	return dsFile, nil
}

// paths, _ := utility.GetDiceImage("", rolls)

func getFilePath(roll Roll) string {
	imgName := fmt.Sprintf("d%vs%v.png", roll.dice, roll.roll)
	imgFolder := fmt.Sprintf("d%v", roll.dice)
	imgPath := filepath.Join("assets", imgFolder, imgName)
	return imgPath
}
