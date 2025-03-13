package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/pdftron/pdftron-go/v2"
)

func main() {
	PDFNetInitialize("YOUR_LICENSE_KEY")
	defer PDFNetTerminate()

	filename := "sample.pdf"
	doc, err := NewPDFDoc(filename)
	if err != nil {
		log.Fatalf("Error opening PDF: %v", err)
	}
	defer doc.Close()

	pageNum := doc.GetPageCount()
	fmt.Printf("Total pages: %d\n", pageNum)

	for i := 1; i <= pageNum; i++ {
		page := doc.GetPage(i)
		if page == nil {
			fmt.Printf("Skipping invalid page: %d\n", i)
			continue
		}

		pageReader := NewElementReader()
		pageReader.Begin(page)
		defer pageReader.Destroy()

		element := pageReader.Next()
		for element.GetMp_elem().Swigcptr() != 0 {
			switch element.GetType() {
			case ElementE_text:
				fmt.Println("Text element found")
			case ElementE_image:
				fmt.Println("Image element found")
			default:
				fmt.Println("Other element found")
			}
			element = pageReader.Next()
		}
		pageReader.End()
	}

	fmt.Println("PDF processing completed.")
}
