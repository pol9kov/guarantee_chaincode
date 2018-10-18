package com

import "encoding/xml"

type Request struct {
	XMLName xml.Name `xml:"request"`

	Channels []string `xml:"channels>channel"`

	EntityName string `xml:"entity_name"`
	Type       string `xml:"type"`
	TypeAttr   string `xml:"type_attr"`

	Args []string `xml:"args>arg"`

	// Edit Fields //////////////////////////////
	FieldPaths  []FieldPath `xml:"field_paths>field_path"`
	FieldValues []string    `xml:"field_values>field_value"`
	/////////////////////////////////////////////
}

type FieldPath struct {
	XMLName xml.Name `xml:"field_path"`

	FieldPath []string `xml:"field"`
}
