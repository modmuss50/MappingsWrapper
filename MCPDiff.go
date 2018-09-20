package mcpwrapper

import "fmt"

type MCPDiff struct {
	NewFields  []SRGField  `json:"newFields"`
	NewMethods []SRGMethod `json:"newMethods"`
	NewParams  []SRGParam  `json:"newParams"`

	LostFields  []SRGField  `json:"lostFields"`
	LostMethods []SRGMethod `json:"lostMethods"`
	LostParams  []SRGParam  `json:"lostParams"`

	ChangedMethods []ChangedMethod `json:"changedMethods"`
	ChangedFields  []ChangedField  `json:"changedFields"`
	ChangedParams  []ChangedParm   `json:"changedParams"`
}

type ChangedMethod struct {
	OldMethod SRGMethod `json:"oldMethod"`
	NewMethod SRGMethod `json:"newMethod"`
}

type ChangedField struct {
	OldField SRGField `json:"oldField"`
	NewField SRGField `json:"newField"`
}

type ChangedParm struct {
	OldParam SRGParam `json:"oldParam"`
	NewParam SRGParam `json:"newParam"`
}

func GenerateDiff(oldMappings SRGNames, newMappings SRGNames) MCPDiff {

	diff := MCPDiff{}

	//Find new and changed methods
	for _, method := range newMappings.Methods {
		oldMethod, err := FindMethodSRGName(method.Searge, oldMappings)
		if err != nil { //We have a new method as it was not found in the old mappings
			diff.NewMethods = append(diff.NewMethods, method)
		} else if method.Name != oldMethod.Name {
			diff.ChangedMethods = append(diff.ChangedMethods, ChangedMethod{OldMethod: oldMethod, NewMethod: method})
		}
	}

	//Find lost methods
	for _, method := range oldMappings.Methods {
		_, err := FindMethodSRGName(method.Searge, newMappings)
		if err != nil { //We have a lost method as it was not found in the new mappings
			diff.LostMethods = append(diff.NewMethods, method)
		}
	}

	//Find new and changed fields
	for _, field := range newMappings.Fields {
		oldField, err := FindFieldSRGName(field.Searge, oldMappings)
		if err != nil { //We have a new field as it was not found in the old mappings
			diff.NewFields = append(diff.NewFields, field)
		} else if field.Name != oldField.Name {
			diff.ChangedFields = append(diff.ChangedFields, ChangedField{OldField: oldField, NewField: field})
		}
	}

	//Find lost fields
	for _, field := range oldMappings.Fields {
		_, err := FindFieldSRGName(field.Searge, newMappings)
		if err != nil { //We have a lost field as it was not found in the new mappings
			diff.LostFields = append(diff.NewFields, field)
		}
	}

	//Find new params
	for _, param := range newMappings.Params {
		oldParam, err := FindParamSRGName(param.Searge, oldMappings)
		if err != nil { //We have a new param as it was not found in the old mappings
			diff.NewParams = append(diff.NewParams, param)
		} else if param.Name != oldParam.Name {
			diff.ChangedParams = append(diff.ChangedParams, ChangedParm{OldParam: oldParam, NewParam: param})
		}
	}

	//Find lost params
	for _, param := range oldMappings.Params {
		_, err := FindParamSRGName(param.Searge, newMappings)
		if err != nil { //We have a lost param as it was not found in the new mappings
			diff.LostParams = append(diff.NewParams, param)
		}
	}

	//TODO find all changed, add the changes

	return diff
}

func DiffToString(diff MCPDiff) string {
	response := ""

	response += fmt.Sprintf("%d New Fields\n", len(diff.NewFields))
	response += fmt.Sprintf("%d New Methods\n", len(diff.NewMethods))
	response += fmt.Sprintf("%d New Params\n", len(diff.NewParams))

	response += fmt.Sprintf("%d Lost Fields\n", len(diff.LostFields))
	response += fmt.Sprintf("%d Lost Methods\n", len(diff.LostMethods))
	response += fmt.Sprintf("%d Lost Params\n", len(diff.LostParams))

	response += fmt.Sprintf("%d Changed Fields\n", len(diff.ChangedFields))
	response += fmt.Sprintf("%d Changed Methods\n", len(diff.ChangedMethods))
	response += fmt.Sprintf("%d Changed Params\n", len(diff.ChangedParams))

	for _, method := range diff.NewMethods {
		response += fmt.Sprintf("New Method: `%s` `%s`\n", method.Searge, method.Name)
	}
	for _, method := range diff.LostMethods {
		response += fmt.Sprintf("Lost Method: `%s` `%s`\n", method.Searge, method.Name)
	}
	for _, change := range diff.ChangedMethods {
		response += fmt.Sprintf("Changed Method: `%s` `%s` -> `%s`\n", change.NewMethod.Searge, change.OldMethod.Name, change.NewMethod.Name)
	}

	for _, field := range diff.NewFields {
		response += fmt.Sprintf("New Field: `%s` `%s`\n", field.Searge, field.Name)
	}
	for _, field := range diff.LostFields {
		response += fmt.Sprintf("Lost Field: `%s` `%s`\n", field.Searge, field.Name)
	}
	for _, change := range diff.ChangedFields {
		response += fmt.Sprintf("Changed Field: `%s` `%s` -> `%s`\n", change.NewField.Searge, change.OldField.Name, change.NewField.Name)
	}

	for _, param := range diff.NewParams {
		response += fmt.Sprintf("New Param: `%s` `%s`\n", param.Searge, param.Name)
	}
	for _, param := range diff.LostParams {
		response += fmt.Sprintf("Lost Param: `%s` `%s`\n", param.Searge, param.Name)
	}
	for _, change := range diff.ChangedParams {
		response += fmt.Sprintf("Changed Param: `%s` `%s` -> `%s`\n", change.NewParam.Searge, change.OldParam.Name, change.NewParam.Name)
	}

	return response
}
