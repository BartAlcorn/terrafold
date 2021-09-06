package scaffold

type Profile struct {
	AwsProfile string   `json:"aws_profile"`
	AwsAccount string   `json:"aws_account"`
	Stages     []string `json:"stages"`
	Triggers   []string `json:"triggers"`
}

type Lambda struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Stages      []string   `json:"stages"`
	Triggers    []string   `json:"triggers"`
	Triggering  Triggering `json:"triggering"`
	Overwrite   bool       `json:"overwrite"`
	Trigger     string     `json:"trigger,omitempty"`
	Stage       string     `json:"stage,omitempty"`
}

type SNS struct {
	Name string `json:"name"`
}

type SQS struct {
	Name string `json:"name"`
}

type API struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type Triggering struct {
	API API    `json:"api"`
	SNS string `json:"sns"`
	SQS string `json:"sqs"`
}

type Folder struct {
	Lambda Lambda `json:"lambda"`
	LPath  string `json:"lpath,omitempty"`
}
