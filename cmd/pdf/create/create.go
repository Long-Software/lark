package create

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Long-Software/Sonality/internal/utils"
 
	"github.com/jung-kurt/gofpdf"
)

var (
	dir = flag.String("dir", "", "Directory path containing the images to create the pdf")
)

type CreateCommand struct{}

func (m *CreateCommand) Run() {
	flag.Usage = m.Help
	flag.CommandLine.Parse(os.Args[2:])
	if *dir == "" {
		log.Fatal("Please provide a directory path using the -dir flag. Use -h for help.")
	}
	_, filename := filepath.Split(*dir)
	pdf := gofpdf.New("P", "mm", "A4", "")
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && utils.IsImageFile(info.Name()) {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			img, format, err := utils.DecodeImage(file)
			if err != nil {
				return err
			}

			width, height := utils.Dimension(img.Bounds())

			file.Seek(0, 0)
			pdf.AddPageFormat("P", gofpdf.SizeType{Wd: width, Ht: height})
			pdf.RegisterImageOptionsReader(info.Name(), gofpdf.ImageOptions{ImageType: format}, file)
			pdf.ImageOptions(info.Name(), 0, 0, width, height, false, gofpdf.ImageOptions{ImageType: format}, 0, "")
			fmt.Printf("Added %s: %s \n", format, info.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error reading folder: %v\n", err)
		return
	}
	err = pdf.OutputFileAndClose(filename + ".pdf")
	if err != nil {
		fmt.Printf("Error saving PDF: %v\n", err)
		return
	}

	fmt.Printf("PDF generated successfully: %s.pdf\n", filename)
}

func (m *CreateCommand) Help() {
	fmt.Println("usage: pdf create [options]")
	flag.PrintDefaults()
}
