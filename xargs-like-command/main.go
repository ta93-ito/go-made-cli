package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	iFlag := flag.String("I", "", "option like xargs's \"-I\".")

	flag.Parse()
	ph := *iFlag

	sc.Split(bufio.ScanLines)

	var cmd, subArgs string

	if ph != "" {
		cmd = os.Args[3]
		subArgs = strings.Join(os.Args[4:], " ")
	} else {
		cmd = os.Args[1]
	}

	args := getArgs(ph, subArgs)

	for _, arg := range args {
		sub := strings.Split(arg, " ")
		out, err := exec.Command(cmd, sub...).CombinedOutput()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	}
}

// コマンドに渡す引数一覧を取得
func getArgs(ph, subArgs string) []string {
	inputs := make([]string, 0)
	for {
		sc.Scan()
		arg := sc.Text()
		if arg == "" {
			return inputs
		}
		if subArgs != "" {
			arg = strings.ReplaceAll(subArgs, ph, arg)
		}
		inputs = append(inputs, arg)
	}
}
