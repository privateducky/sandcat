package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"./api"
	"./execute"
	"./util"
)

type executorFlags []string

var iteration = 60
var executors executorFlags

func askForInstructions(profile map[string]string) {
	commands := api.Instructions(profile)
	if commands != nil && len(commands.([]interface{})) > 0 {
		cmds := reflect.ValueOf(commands)
		for i := 0; i < cmds.Len(); i++ {
			cmd := cmds.Index(i).Elem().String()
			fmt.Println("[*] Running instruction")
			command := util.Unpack([]byte(cmd))
			api.Drop(profile["server"], command["payload"].(string))
			api.Execute(profile, command)
		}
	} else {
		time.Sleep(time.Duration(iteration) * time.Second)
	}
}

func buildProfile(server string, group string, executors []string) map[string]interface{} {
	host, _ := os.Hostname()
	user, _ := user.Current()
	paw := fmt.Sprintf("%s$%s", host, user.Username)
	arch := runtime.GOARCH
	profile := make(map[string]interface{})
	profile["paw"] = paw
	profile["server"] = server
	profile["group"] = group
	profile["architecture"] = arch
	profile["platform"] = runtime.GOOS
	profile["location"] = os.Args[0]
	profile["pid"] = strconv.Itoa(os.Getpid())
	profile["ppid"] = strconv.Itoa(os.Getppid())
	profile["executors"] = executors
	return profile
}

func (i *executorFlags) String() string {
	return fmt.Sprint((*i))
}

func (i *executorFlags) Set(value string) error {
	for _, exec := range strings.Split(value, ",") {
		*i = append(*i, exec)
	}
	return nil
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	executors = []string{execute.DetermineExecutor(runtime.GOOS)}
	server := flag.String("server", "http://localhost:8888", "The FQDN of the server")
	group := flag.String("group", "my_group", "Attach a group to this agent")
	flag.Var(&executors, "executors", "Comma separated list of executors")
	flag.Parse()

	profile := buildProfile(*server, *group, executors)
	for {
		askForInstructions(profile)
	}
}

var key = "3TEU4UD15V29OBJB7U9HNCR2JPWL1U"
