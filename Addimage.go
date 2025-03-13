package main

import (
    // "github.com/gofiber/fiber/v2"
    "github.com/pdftron/pdftron-go/v2"
    "fmt"
    "log"
)

func Addimgoperation() {
    var inputPath = "./inputfiles/"
    var outputpath = "./output/output.pdf"

    doc := pdftron.NewPDFDoc()
    f := pdftron.NewElementBuilder()
    writer := pdftron.NewElementWriter()
    page := doc.PageCreate()
    writer.Begin(page)

    img := pdftron.ImageCreate(doc.GetSDFDoc(), inputPath + "sky.jpeg")
    element := f.CreateImage(img, 50.0, 500.0, float64(img.GetImageWidth()/2), float64(img.GetImageHeight()/2))
    writer.WritePlacedElement(element)


    writer.End()
    doc.PagePushBack(page)

    err := doc.Save(outputpath, uint(pdftron.SDFDocE_linearized))
    if err != nil {
        fmt.Println("Error while saving the file")
    }

    log.Printf("Image successfully added to the PDF file.")
}
