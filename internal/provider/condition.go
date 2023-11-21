package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/api/tagmanager/v2"
)

var conditionSchema = schema.ListNestedAttribute{
	Optional: true,
	NestedObject: schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Description: "Condition type.",
				Required:    true},
			"parameter": parameterSchema,
		},
	},
}

type ResourceConditionModel struct {
	Type      types.String             `tfsdk:"type"`
	Parameter []ResourceParameterModel `tfsdk:"parameter"`
}

// Equal compares two resource conditions.
func (m ResourceConditionModel) Equal(o ResourceConditionModel) bool {
	if !m.Type.Equal(o.Type) ||
		len(m.Parameter) != len(o.Parameter) {
		return false
	}

	for i, p := range m.Parameter {
		if !p.Equal(o.Parameter[i]) {
			return false
		}
	}

	return true
}

func toApiCondition(resourceCondition []ResourceConditionModel) []*tagmanager.Condition {
	condition := make([]*tagmanager.Condition, len(resourceCondition))

	for i, rc := range resourceCondition {
		var parameter []*tagmanager.Parameter
		if rc.Parameter != nil {
			parameter = toApiParameter(rc.Parameter)
		}

		condition[i] = &tagmanager.Condition{
			Type:      rc.Type.ValueString(),
			Parameter: parameter,
		}
	}
	return condition
}

func toResourceCondition(condition []*tagmanager.Condition) []ResourceConditionModel {
	resourceCondition := make([]ResourceConditionModel, len(condition))

	for i, c := range condition {
		var resourceParameter []ResourceParameterModel
		if c.Parameter != nil {
			resourceParameter = toResourceParameter(c.Parameter)
		}

		resourceCondition[i] = ResourceConditionModel{
			Type:      nullableStringValue(c.Type),
			Parameter: resourceParameter,
		}
	}

	return resourceCondition
}
