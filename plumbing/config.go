package plumbing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Common struct {
	Env      string
	Keys     map[string]string
	Commands map[string]string
}

const (
	IdxCommand    = 2
	IdxCollection = 3
	IdxModifier   = 4
	IdxParams     = 5
)

func Config() *Common {
	cmn := Common{}
	if len(os.Args) > 1 && len(os.Args) < 6 {
		cmn.Env = os.Args[1]
		cmn.Commands = setCommandConfig(os.Args)
	}
	cmn.Keys = keysFromEnv()

	return &cmn
}

func setCommandConfig(a []string) map[string]string {
	m := make(map[string]string)
	m["Command"] = a[IdxCommand]
	m["Collection"] = a[IdxCollection]
	m["Modifier"] = a[IdxModifier]
	m["Params"] = a[IdxParams]
	return m
}

func keysFromEnv() map[string]string {
	f, err := os.Open("./.env")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		kv := strings.Split(ln, "=")
		key := kv[0]
		val := kv[1]
		m[key] = val
	}

	return m
}
