package conf

import (
	"encoding/json"
	io "io/ioutil"
	re "regexp"

	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Plug struct {
		Bl []struct {
			Cmd   string `json:"cmd"`
			Allow bool   `json:"allow"`
			Dmsg  string `json:"dmsg"`
		} `json:"commands"`
		Desc bool `json:"description"`
	} `json:"plugin"`
}

type Db struct {
	All []struct {
		Cmd  []string `yaml:"commands"`
		Meth string   `yaml:"method"`
		Path []string `yaml:"path"`
	} `yaml:"methods"`
}

type nallowed struct {
	path    []string
	method  string
	command string
	dmsg    string
	tmpstat bool
}

var (
	jd     Config
	mapd   Db
	na     []nallowed
	config string = "/etc/docksec/main.json"
	db     string = "/usr/share/docksec/api.yml"
)

func init() {
	var (
		con []byte
		err error
	)
	if con, err = io.ReadFile(config); err != nil {
		logrus.Fatal(err)
	}
	if err = json.Unmarshal(con, &jd); err != nil {
		logrus.Fatal(err)
	}
	if con, err = io.ReadFile(db); err != nil {
		logrus.Fatal(err)
	}
	if err = yaml.Unmarshal(con, &mapd); err != nil {
		logrus.Fatal(err)
	}
	for _, jdata := range jd.Plug.Bl {
		for _, ydata := range mapd.All {
			for _, cmds := range ydata.Cmd {
				if jdata.Cmd == cmds {
					na = append(na, nallowed{ydata.Path, ydata.Meth, jdata.Cmd, jdata.Dmsg, jdata.Allow})
				}
			}
		}
	}
}

func GetDescStat() bool {
	return jd.Plug.Desc
}

func GetStatus(reqm, requ string) (bool, string) {
	for _, v := range na {
		stat, _ := re.MatchString(v.method, reqm)
		if stat {
			for _, _v := range v.path {
				stat, _ = re.MatchString(_v, requ)
				if !stat {
					continue
				}
				if !v.tmpstat {
					if jd.Plug.Desc {
						return true, v.dmsg
					}
					return true, ""
				}
				return false, ""
			}
		}
	}
	return false, ""
}
