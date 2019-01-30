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

package metagopenoffice

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"errors"
	"os"
	"strings"
)

//Statistics is a struct that contains the document-statistics
type Statistics struct {
	PageCount      int `xml:"page-count,attr"`
	ImageCount     int `xml:"image-count,attr"`
	ObjectCount    int `xml:"object-count,attr"`
	ParagraphCount int `xml:"paragraph-count,attr"`
	WordCount      int `xml:"word-count,attr"`
	CharCount      int `xml:"character-count,attr"`
}

//Metadata is a struct that contains metadata fields
type Metadata struct {
	Generator       string     `xml:"generator"`
	InitialCreator  string     `xml:"initial-creator"`
	CreationDate    string     `xml:"creation-date"`
	Creator         string     `xml:"creator"`
	Date            string     `xml:"date"`
	Language        string     `xml:"language"`
	EditingCycles   string     `xml:"editing-cycles"`
	EditingDuration string     `xml:"editing-duration"`
	Stats           Statistics `xml:"document-statistic"`
}

//OpenOfficeXML contains the fields of te file meta.xml
type OpenOfficeXML struct {
	Meta Metadata `xml:"meta"`
}

//GetMetada function
func GetMetada(document *os.File) (metadata OpenOfficeXML, err error) {

	fileName := document.Name()

	dot := strings.LastIndex(fileName, ".")
	after := fileName[:dot] + ".zip"
	err = os.Rename(fileName, after)
	if err != nil {
		return metadata, errors.New("Failed to rename as .zip")
	}

	read, err := zip.OpenReader(after)
	if err != nil {
		os.Rename(after, fileName)
		return metadata, errors.New("Failed to open the file")
	}
	var xmlFile string
	for _, file := range read.File {
		if file.Name == "meta.xml" {
			rc, err := file.Open()
			scanner := bufio.NewScanner(rc)
			for scanner.Scan() {
				xmlFile += scanner.Text()
			}
			if err != nil {
				os.Rename(after, fileName)
				return metadata, errors.New("Failed to open meta.xml")
			}
			defer rc.Close()
		}
	}
	if err := xml.Unmarshal([]byte(xmlFile), &metadata); err != nil {
		os.Rename(after, fileName)
		return metadata, errors.New("Failed to Unmarshal")
	}
	read.Close()
	os.Rename(after, fileName)
	return metadata, err
}
