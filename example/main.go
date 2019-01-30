/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinborras/metagopenoffice"
)

func main() {

	file, err := os.Open("../test/ods_sample.ods")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	content, err := metagopenoffice.GetMetada(file)
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
