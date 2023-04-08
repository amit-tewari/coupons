package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/signintech/gopdf"
)

func main() {
	// denomination of the coupons
	denomination, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Print(err.Error())
		return
	}

	// how many sheets to print
	sheets, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Print(err.Error())
		return
	}

	// Add some prefix about the event
	eventPrefix := os.Args[3]

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		Unit: gopdf.Unit_PT,
		PageSize: gopdf.Rect{ //595.28, 841.89 = A4
			W: 595.28,
			H: 841.89,
		},
	})

	// Add fonts
	err = pdf.AddTTFFont(
		"TakaoPGothic",
		"/usr/share/fonts/truetype/freefont/FreeSans.ttf", //need to be adjusted
		//"../go/ttf/TakaoPGothic.ttf",
	)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// Generate coupons in a 3,9 grid
	for sheet := 1; sheet <= sheets; sheet++ {
		pdf.AddPage()
		pdf.SetLineWidth(1)
		pdf.SetLineType("dotted")

		// horizontal lines
		for i := 0.0; i <= 9; i++ {
			pdf.Line(0.28, 93.3*i+0.89,
				595, 93.3*i+0.89)
		}

		// vertical lines
		for i := 0.0; i <= 3; i++ {
			pdf.Line(i*198+0.28, 0.89,
				i*198+0.28, 841.0)
		}

		err = pdf.SetFont("TakaoPGothic", "", 12)
		if err != nil {
			log.Print(err.Error())
			return
		}

		// Print coupons serial
		// Serial includes: Event Prefix, Coupon denomination, Sheet number, row & column number
		for i := 0.0; i <= 3; i++ {
			for j := 0.0; j <= 9; j++ {
				pdf.SetXY(i*198+0.28+2,
					93.3*j+0.89+2)
				pdf.Cell(nil, fmt.Sprintf("%s-%d-%d-%d%d",
					eventPrefix,
					denomination,
					sheet,
					int(j),
					int(i)))
			}
		}

		// Print denomination in bigger font size
		err = pdf.SetFont("TakaoPGothic", "", 48)
		if err != nil {
			log.Print(err.Error())
			return
		}
		for i := 0.0; i <= 3; i++ {
			for j := 0.0; j <= 9; j++ {
				pdf.SetXY(i*198+0.28+2+50,
					93.3*j+0.89+2+25)
				pdf.Cell(nil, fmt.Sprintf("%d",
					denomination))
			}
		}
	}

	// Save to file
	pdf.WritePdf(fmt.Sprintf("%s-denom-%d-sheets-%d.pdf",
		eventPrefix,
		denomination,
		sheets))
}
