# metagopenoffice

[![Build Status](https://travis-ci.org/kevinborras/metagopenoffice.svg?branch=master)](https://travis-ci.org/kevinborras/metagopenoffice)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinborras/metagopenoffice)](https://goreportcard.com/badge/github.com/kevinborras/metagopenoffice)

Open Office metadata extractor written in Go

The main features of metagopenoffice are:

* Read metadata of odt files.
* Read metadata of odp files.
* Read metadata of ods files.

## How to use ?
---

```golang
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinborras/metagopenoffice"
)

func main() {

	file, err := os.Open("ods_sample.ods")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	content, err := gopenoffice.GetMetada(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generator: ", content.Meta.Generator)
	fmt.Println("InitialCreator: ", content.Meta.InitialCreator)
	fmt.Println("CreationDate: ", content.Meta.CreationDate)
	fmt.Println("Creator: ", content.Meta.Creator)
	fmt.Println("Date: ", content.Meta.Date)
	fmt.Println("Language: ", content.Meta.Language)
	fmt.Println("EditingCycles: ", content.Meta.EditingCycles)
	fmt.Println("EditingDuration: ", content.Meta.EditingDuration)
	fmt.Println("Title: ", content.Meta.Title)
	fmt.Println("Description: ", content.Meta.Description)
	fmt.Println("Subject: ", content.Meta.Subject)
	fmt.Println("PrintDate: ", content.Meta.PrintDate)
	fmt.Println("Keyword: ", content.Meta.Keyword)
	fmt.Println()
	fmt.Println("PageCount: ", content.Meta.Stats.PageCount)
	fmt.Println("ImageCount: ", content.Meta.Stats.ImageCount)
	fmt.Println("ObjectCount: ", content.Meta.Stats.ObjectCount)
	fmt.Println("ParagraphCount: ", content.Meta.Stats.ParagraphCount)
	fmt.Println("WordCount: ", content.Meta.Stats.WordCount)
	fmt.Println("CharCount: ", content.Meta.Stats.CharCount)

}
```

Output

```bash
Generator:  MicrosoftOffice/14.0 MicrosoftExcel/CalculationVersion-9303
InitialCreator:
CreationDate:  2008-01-04T03:08:19Z
Creator:
Date:  2012-07-04T03:25:20Z
Language:
EditingCycles:
EditingDuration:
Title:  Welcome to File Extension FYI Center!
Description:
Subject:
PrintDate:
Keyword:

PageCount:  0
ImageCount:  0
ObjectCount:  0
ParagraphCount:  0
WordCount:  0
CharCount:  0
```