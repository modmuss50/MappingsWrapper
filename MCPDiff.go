package mcpwrapper

type MCPDiff struct {
	NewFields  []SRGField  `json:"newFields"`
	NewMethods []SRGMethod `json:"newMethods"`
	NewParams  []SRGParm   `json:"newParams"`

	LostFields  []SRGField  `json:"lostFields"`
	LostMethods []SRGMethod `json:"lostMethods"`
	LostParams  []SRGParm   `json:"lostParams"`

	ChangedMethods []ChangedMethod `json:"changedMethods"`
	ChangedFields  []ChangedField  `json:"changedFields"`
	ChangedParms   []ChangedParm   `json:"changedParms"`
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
	OldParms SRGParm `json:"oldParms"`
	NewParms SRGParm `json:"newParms"`
}

func GenerateDiff(oldMappings SRGNames, newMappings SRGNames) MCPDiff {
	//TODO find all new things, and if not found add to lost

	//TODO find all changed, add the changes
}

func DiffToString(diff MCPData) {
	//Make it a bunch of strings so it can be shown to users
}
