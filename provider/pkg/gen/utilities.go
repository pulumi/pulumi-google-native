// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"strings"
)

var s = codegen.NewStringSet()

// apiNameToSdkName returns an SDK name for a given API parameter or property name.
// In particular, it converts pluralized ID names to singular ID names.
func apiNameToSdkName(name string) string {
	if strings.HasSuffix(name, "Id") {
		name = inflector.Singularize(name[:len(name)-2])
		switch name {
		case "project", "location", "managedZone":
			return name
		case "datum":
			return "dataId"
		default:
			return name + "Id"
		}
	}
	return name
}

// ToLowerCamel converts a string to lowerCamelCase.
// The code is adopted from https://github.com/iancoleman/strcase but changed in several ways to handle
// all the cases that are found in Azure in a most user-friendly way.
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	if r := rune(s[0]); r == '_' {
		s = s[1:]
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

// ToUpperCamel converts a string to UpperCamelCase.
func ToUpperCamel(s string) string {
	return toCamelInitCase(s, true)
}

func toCamelInitCase(s string, initCase bool) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

var uppercaseAcronym = map[string]bool{
	"ID":  true,
	"TTL": true,
}
