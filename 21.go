package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

//---------------------------------------------------------
type someService interface {
	convertToJson(string) string
}

type XMLData struct{}

func (s *XMLData) convertToJson(data string) string {
	newStr := fmt.Sprintf("Converted to json: %s", data)
	return newStr
}

//---------------------------------------------------------
type otherService interface {
	handleJSONData() string
}

//---------------------------------------------------------
type adapter struct {
	service someService
	data    string
}

func (a *adapter) handleJSONData() string {
	if a.service != nil {
		newStr := fmt.Sprintf("Adapter: %s", a.data)
		newStr = a.service.convertToJson(newStr)
		return newStr
	}
	return a.data
}
