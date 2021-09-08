package deploy

type Package struct {
	Project string `json:"project"`
	Trigger string `json:"trigger"`
	Version string `json:"version"`
}
