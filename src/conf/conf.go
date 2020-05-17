package conf

import (
    "encoding/json"
    io "io/ioutil"
    re "regexp"
    "github.com/sirupsen/logrus"
    "github.com/go-yaml/yaml"
    "fmt"
)

type Config struct {
    Plug  struct {
        Sock bool `json:"socket"`
        Tcp uint32 `json:"port"`
        Bl []struct {
            Cmd string `json:"cmd"`
            Allow bool `json:"allow"`
            Dmsg string `json:"dmsg"`
            Amsg string `json:"admin-msg"`
        } `json:"commands"`
        Desc bool `json:"description"`
    } `json:"plugin"`
}

type Db struct {
    All []struct {
        Cmd []string `yaml:"commands"`
        Meth string `yaml:"method"`
        Path []string `yaml:"path"`
    } `yaml:"methods"`
}

type nallowed struct {
    path []string
    method string
    command string
    dmsg string
    amsg string
    tmpstat bool
}

var jd Config
var mapd Db
var na []nallowed
var config string = "/etc/docksec/main.json"
var db string = "/usr/share/docksec/api.yml"

func init(){
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
    if err := yaml.Unmarshal(con, &mapd){
        logrus.Fatal(err)
    }
    for _, jdata := range jd.Plug.Bl {
        for _,  ydata := range mapd.All {
            for _, cmds := range ydata.Cmd {
                if jdata.Cmd == cmds {
                    na = append(na, nallowed{ydata.Path, ydata.Meth, jdata.Cmd, jdata.Dmsg, jdata.Amsg, jdata.Allow})
                }
            }
        }
    }
}


func GetPort() uint32 {
    return jd.Plug.Tcp
}

func GetSockStat() bool {
    return jd.Plug.Sock
}

func GetDescStat() bool {
    return jd.Plug.Desc
}

func GetStatus(reqm, requ string) (bool, string, string){
    for _, v := range na {
        _status, _ := re.MatchString(v.method, reqm)
        if _status {
            for _, _v := range v.path {
                _status, _ = re.MatchString(_v, requ)
                if ! _status {
                    continue
                }
                if ! v.tmpstat {
                    if jd.Plug.Desc {
                        return true, v.dmsg, v.amsg
                    }
                    return true, "", ""
                }
                return false, "", ""
            }
        }
    }
    return false, "", ""
}
