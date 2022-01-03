package main

import (
  	"math/rand"
  	"time"
  	"strings"
  	"net/http"
  	"bytes"
  	"io/ioutil"
  	log "github.com/sirupsen/logrus"
)

func z(s string) string {
	variants := []string{"", "", "", "", "", "", "", "", "", "z", "s"}
	s += variants[rand.Intn(len(variants))]
	return s
}

func xor(s string) string {
	variants := []string{"xor", "zor"}
	s += variants[rand.Intn(len(variants))]
	return z(s)
}

func er(s string) string {
	return s+"er"
}

func k(s string) string {
	s += "k"
	var variants []func(string) string
	variants = append(variants, xor)
	variants = append(variants, er)
	return variants[rand.Intn(len(variants))](s)
}

func xxor(s string) string {
	variants := []string{"xxor", "xzor", "zxor", "zzor"}
	s += variants[rand.Intn(len(variants))]
	return z(s)
}

func c(s string) string {
	s += "c"
	var variants []func(string) string
	variants = append(variants, xxor)
	variants = append(variants, k)
	return variants[rand.Intn(len(variants))](s)
}

var latest string = ""

func ha() string{
	h := "ha"
	var variants []func(string) string
	variants = append(variants, xxor)
	variants = append(variants, c)
	h = variants[rand.Intn(len(variants))](h)
	return h
}

func leet(s string) string{
	subs := map[string][]string{
		"h": []string{"h", "#", "h", "#", "h", "#", "][", "}{", ")("},
		"a": []string{"a", "4", "@", "^"},
		"c": []string{"c", "(", "<", "$"},
		"e": []string{"e", "3", "&"},
		"r": []string{"r", "9", "7", "2", "?"},
		"x": []string{"x", "ecks", "][", "}{", ")(", "*"},
		"o": []string{"o", "0", "o", "0", "()", "[]"},
		"z": []string{"z", "z", "2", "2", "s"},
	}
	ret := ""
	for _, char := range s {
		variants, ok := subs[string(char)]
		if ok {
			ret += variants[rand.Intn(len(variants))]
		}else {
			ret += string(char)
		}
	}
	return ret
}

func hacker() string{
	hc := ""
	for {
		hc = leet(ha())
		if len(hc) > 9 { continue }
		symbols := 0
		symbols += strings.Count(hc, "(")
		symbols += strings.Count(hc, ")")
		symbols += strings.Count(hc, "[")
		symbols += strings.Count(hc, "]")
		symbols += strings.Count(hc, "{")
		symbols += strings.Count(hc, "}")
		if symbols > 3 { continue}
		if hc != latest {
			latest = hc
			break
		}
	}
	return hc
}

func set(bio string, token string) {
	url := "https://api.github.com/user"
	bio = "{\"bio\":\""+bio+"\"}"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer([]byte(bio)))
    if err != nil {
    	log.Fatal(err)
    }
    req.Header.Set("Accept", "application/vnd.github.v3+json")
    req.Header.Set("Authorization", "token "+token)
    resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Wrong answer code ", resp.StatusCode)
	}
}

func main() {
	b, err := ioutil.ReadFile("/etc/dynbio.txt")
	if err != nil {
		log.Fatal(err)
	}
	token := strings.TrimSuffix(string(b), "\n")
	rand.Seed(time.Now().Unix())
	for {
		bio := "Coder, "+hacker()+", sometimes artist"
    	set(bio, token)
    	log.Info(bio)
    	time.Sleep(30 * time.Second)
    }
}
