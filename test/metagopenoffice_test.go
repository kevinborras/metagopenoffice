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

package test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/kevinborras/metagopenoffice"
)

var odtExpectedResult = metagopenoffice.OpenOfficeXML{
	Meta: metagopenoffice.Metadata{
		Generator:       "StarOffice/8_Beta$Linux OpenOffice.org_project/680m66$Build-8852$CWS-sdksample",
		InitialCreator:  "Jürgen Schmidt",
		CreationDate:    "2002-12-18T12:28:35",
		Creator:         "Jürgen Schmidt",
		Date:            "2002-12-18T12:31:15",
		Language:        "en-US",
		EditingCycles:   "3",
		EditingDuration: "PT2M40S",
		Stats: metagopenoffice.Statistics{
			PageCount:      1,
			ImageCount:     0,
			ObjectCount:    0,
			ParagraphCount: 7,
			WordCount:      77,
			CharCount:      511,
		},
	},
}

var odsExpectedResult = metagopenoffice.OpenOfficeXML{
	Meta: metagopenoffice.Metadata{
		Generator:       "MicrosoftOffice/14.0 MicrosoftExcel/CalculationVersion-9303",
		InitialCreator:  "",
		CreationDate:    "2008-01-04T03:08:19Z",
		Creator:         "",
		Date:            "2012-07-04T03:25:20Z",
		Language:        "",
		EditingCycles:   "",
		EditingDuration: "",
		Stats: metagopenoffice.Statistics{
			PageCount:      0,
			ImageCount:     0,
			ObjectCount:    0,
			ParagraphCount: 0,
			WordCount:      0,
			CharCount:      0,
		},
	},
}

var odpExpectedResult = metagopenoffice.OpenOfficeXML{
	Meta: metagopenoffice.Metadata{
		Generator:       "OpenOffice.org/2.0$Linux OpenOffice.org_project/680m7$Build-9044",
		InitialCreator:  "",
		CreationDate:    "2006-09-02T14:46:49",
		Creator:         "Laurent Godard",
		Date:            "2006-09-19T09:55:18",
		Language:        "fr-FR",
		EditingCycles:   "155",
		EditingDuration: "P1DT4H4M7S",
		Stats: metagopenoffice.Statistics{
			PageCount:      0,
			ImageCount:     0,
			ObjectCount:    240,
			ParagraphCount: 0,
			WordCount:      0,
			CharCount:      0,
		},
	},
}

func TestODTDocumment(t *testing.T) {
	file, err := os.Open("test.odt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagopenoffice.GetMetada(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(odtExpectedResult, actualResult) {
		t.Fatal("Not passing")
	}
}

func TestODSDocumment(t *testing.T) {
	file, err := os.Open("ods_sample.ods")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagopenoffice.GetMetada(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(odsExpectedResult, actualResult) {
		t.Fatalf("Not passing")
	}
}

func TestODPDocumment(t *testing.T) {
	file, err := os.Open("tuesday_d6.odp")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagopenoffice.GetMetada(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(odpExpectedResult, actualResult) {
		t.Fatal("Not passing")
	}
}
