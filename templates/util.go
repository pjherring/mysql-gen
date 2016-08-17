package templates

import "strings"

var TemplateReplacer *strings.Replacer = strings.NewReplacer("\t", "", "\n", "", " ", "")
