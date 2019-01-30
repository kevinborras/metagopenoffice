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
Generator:  StarOffice/8_Beta$Linux OpenOffice.org_project/680m66$Build-8852$CWS-sdksample
InitialCreator:  Jürgen Schmidt
CreationDate:  2002-12-18T12:28:35
Creator:  Jürgen Schmidt
Date:  2002-12-18T12:31:15
Language:  en-US
EditingCycles:  3
EditingDuration:  PT2M40S

PageCount:  1
ImageCount:  0
ObjectCount:  0
ParagraphCount:  7
WordCount:  77
CharCount:  511
```