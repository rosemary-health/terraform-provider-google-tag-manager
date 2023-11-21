package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func nullableStringValue(s string) types.String {
	if s != "" {
		return types.StringValue(s)
	} else {
		return types.StringNull()
	}
}

func toResourceStringArray(list []string) []types.String {
	var rv []types.String

	for _, v := range list {
		rv = append(rv, types.StringValue(v))
	}

	return rv
}

func unwrapStringArray(list []types.String) []string {
	var rv []string

	for _, v := range list {
		rv = append(rv, v.ValueString())
	}

	return rv
}
