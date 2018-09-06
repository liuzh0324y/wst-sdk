package omigad

type httpconfig struct {
	Server struct {
		Appname    string `yaml:"appname"`
		Httpport   string `yaml:"httpport"`
		Runmodel   string `yaml:"runmode"`
		Copy       bool   `yaml:"copyrequestbody"`
		Endpoint   string `yaml:"endpoint"`
		AccessKey  string `yaml:"accesskey"`
		SecretKey  string `yaml:"secretkey"`
		Bucket     string `yaml:"bucket"`
		FilePath   string `yaml:"filepath"`
		PutHost    string `yaml:"puthost"`
		PutPort    string `yaml:"putport"`
		PutPath    string `yaml:"putpath"`
		ManagerURL string `yaml:"managerurl"`
	} `yaml:"server"`
}
