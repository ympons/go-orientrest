package orientrest

type OClass struct {
	Name             string            `json:"name"`
	ShortName        string            `json:"shortname,omitempty"`
	Properties       []OProperty       `json:"properties"`
	DefaultClusterId int               `json:"defaultCluster"`
	ClusterIds       []int             `json:"clusters"`
	SuperClass       string            `json:"superClass"`
	OverSize         float32           `json:"oversize,omitempty"`
	StrictMode       bool              `json:strictMode,omitempty`
	AbstractClass    bool              `json:"abstractClass,omitempty"`
	ClusterSelection string            `json:"clusterSelection"`
	CustomFields     map[string]string `json:"customFields,omitempty"`
	Records          int               `json:"records"`
}

type OProperty struct {
	Id           int
	Name         string            `json:"name"`
	Fullname     string            `json:"fullname,omitempty"`
	Type         string            `json:"type"`
	NotNull      bool              `json:"notNull"`
	Collate      string            `json:"collate,omitempty"`
	Mandatory    bool              `json:"mandatory"`
	Min          string            `json:"min"`
	Max          string            `json:"max"`
	CustomFields map[string]string `json:"customFields"`
	Readonly     bool              `json:"readonly"`
	Indexed      bool              `json:"indexed"`
}

type ODocument struct {
	Rid       string             `json:"@rid"`
	Version   int                `json:"@version"`
	Fields    map[string]*OField `json:"fields,omitempty"`
	Type      string             `json:"@type"`
	ClassName string             `json:"@class"`
}

type OField struct {
	Id    int
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
