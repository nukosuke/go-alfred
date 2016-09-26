package alfred

import (
	"strings"
	"fmt"
	"encoding/xml"
	"encoding/json"
)

type AlfredResponse struct {
	Items []AlfredResponseItem `json:"items"`
	XMLName struct{} `json:"-" xml:"items"`
}

type AlfredResponseItem struct {
	Valid bool `json:"valid" xml:"valid,attr"`
	Arg string `json:"arg,omitempty" xml:"arg,attr,omitempty"`
	Uid string `json:"uid,omitempty" xml:"uid,attr,omitempty"`
	Title string `json:"title" xml:"title"`
	Subtitle string `json:"subtitle" xml:"subtitle"`
	Icon string `json:"icon" xml:"icon"`

	XMLName struct{} `json:"-" xml:"item"`
}

const xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

func NewResponse() *AlfredResponse {
	return new(AlfredResponse).Init()
}

func (response *AlfredResponse) Init() *AlfredResponse {
	response.Items = []AlfredResponseItem{}
	return response
}

func (response *AlfredResponse) AddItem(item *AlfredResponseItem) {
	response.Items = append(response.Items, *item)
}

func (response *AlfredResponse) PrintJSON() {
	var jsonOutput, _ = json.Marshal(response)
	fmt.Print(string(jsonOutput))
}

func (response *AlfredResponse) PrintXML() {
	var xmlOutput, _ = xml.Marshal(response)
	fmt.Print(xmlHeader, string(xmlOutput))
}

// for backward compatibility (Alfred 2)
func (response *AlfredResponse) Print() {
	response.PrintXML()
}

func InitTerms(params []string) {
	for index, term := range params {
		params[index] = strings.ToLower(term)
	}
}

func MatchesTerms(queryTerms []string, itemName string) bool {
	nameLower := strings.ToLower(itemName)

	for _, term := range queryTerms {
		if ! strings.Contains(nameLower, term) { return false }
	}

	return true
}
