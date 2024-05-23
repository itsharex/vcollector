package fixed

import (
	"github.com/gvcgo/vcollector/pkgs/crawlers/crawler"
	"github.com/gvcgo/vcollector/pkgs/version"
)

func init() {
	crawler.RegisterCrawler(NewCygwin())
}

type Cygwin struct {
	SDKName string
	Version version.VersionList
}

func NewCygwin() *Cygwin {
	return &Cygwin{
		SDKName: "cygwin",
		Version: make(version.VersionList),
	}
}

func (c *Cygwin) GetSDKName() string {
	return c.SDKName
}

func (c *Cygwin) Start() {
	c.Version["latest"] = version.Version{
		version.Item{
			Url:       "https://cygwin.com/setup-x86_64.exe",
			Arch:      "amd64",
			Os:        "windows",
			Installer: version.Executable,
		},
	}
}

func (c *Cygwin) GetVersions() []byte {
	result, _ := c.Version.Marshal()
	return result
}

func (c *Cygwin) HomePage() string {
	return "https://cygwin.com/"
}
