package cmd

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"slices"

	"github.com/Long-Software/Bex/apps/cmd/image/internal/logger"
	"github.com/Long-Software/Bex/packages/log"
	"github.com/spf13/cobra"

	_ "image/jpeg" // required for decoding JPEG
	_ "image/png"
)

var formats = []string{
	"jpeg",
	"jpg",
	"png",
}
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress the image",
	Long:  "Compress the image file provided",
	Run: func(cmd *cobra.Command, args []string) {
		input, err := cmd.Flags().GetString("input")
		quality, err := cmd.Flags().GetInt("quality")

		file, err := os.Open(input)
		if err != nil {
			logger.NewLog(log.ERROR, err.Error())
			return
		}
		defer file.Close()
		img, format, err := image.Decode(file)
		if err != nil {

			logger.NewLog(log.INFO, format)
			logger.NewLog(log.ERROR, err.Error())
			return
		}
		if !slices.Contains(formats, format) {
			logger.NewLog(log.ERROR, "Unsupported format: "+format)
			return
		}
		outFile, err := os.Create("output_" + input)
		if err != nil {
			logger.NewLog(log.ERROR, err.Error())
			return
		}
		defer outFile.Close()

		switch format {
		case "jpg", "jpeg":
			options := jpeg.Options{Quality: quality} // adjust 1-100, lower = more compression
			err = jpeg.Encode(outFile, img, &options)
		case "png":
			encoder := png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(outFile, img)
		default:
			logger.NewLog(log.ERROR, "Unsupported output format. Use .jpg, .jpeg, or .png")
			return
		}
		if err != nil {
			logger.NewLog(log.ERROR, err.Error())
			return
		}


		logger.NewLog(log.INFO, "Image compressed and saved to: "+outFile.Name())
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	compressCmd.Flags().StringP("input", "i", "", "input image file")
	compressCmd.Flags().IntP("quality", "q", 60, "quality of the image")
}
