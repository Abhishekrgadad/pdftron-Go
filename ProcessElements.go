package main

import (
    "fmt"
    . "github.com/pdftron/pdftron-go/v2"
    
)

var inputPath = "../../TestFiles/"

func ProcessElements(reader ElementReader) {
    element := reader.Next()
    for element.GetMp_elem().Swigcptr() != 0 {
        if element.GetType() == ElementE_path {
        } else if element.GetType() == ElementE_text {
            data := element.GetTextString()
            fmt.Println(data)
        } else if element.GetType() == ElementE_form {
            reader.FormBegin()
            ProcessElements(reader)
            reader.End()
        }
        element = reader.Next()
    }
}

func main() {
    PDFNetInitialize("Add your Key")

    fmt.Println("-------------------------------------------------")
    fmt.Println("Sample 1 - Extract text data from all pages in the document.")
    fmt.Println("Opening the input PDF...")

    doc := NewPDFDoc(inputPath + "newsletter.pdf")
    doc.InitSecurityHandler()

    pageReader := NewElementReader()
    itr := doc.GetPageIterator()

    for itr.HasNext() {
        pageReader.Begin(itr.Current())
        ProcessElements(pageReader)
        pageReader.End()
        itr.Next()
    }

    doc.Close()
    PDFNetTerminate()

    fmt.Println("Done.")
}
