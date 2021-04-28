package plumbing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Common struct {
	Env  string
	Keys map[string]string
}

func Config() *Common {
	cmn := Common{}
	if len(os.Args) > 1 {
		cmn.Env = os.Args[1]
	}
	cmn.Keys = keysFromEnv()

	return &cmn
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
