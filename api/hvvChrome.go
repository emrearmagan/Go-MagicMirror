package api

import (
	"context"
	"errors"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

//---------------------------------------------OLD AND NOT USED---------------------------------------------------------
type Traffic struct {
	Start       string
	Arrive      string
	Traffic     []string
	Duration    string
	Destination string
}

var Urls = map[string]string{
	"Dammtor":      "https://geofox.hvv.de/jsf/home.seam?clear=true&onefield=true&language=de&start=Moorburger%20Ring&startCity=Hamburg&startType=STATION&destination=Dammtor%20(Messe%2FCCH)&destinationCity=Hamburg&destType=STATION&execute=true",
	"Informatikum": "https://geofox.hvv.de/jsf/home.seam?execute=true&date=29.03.2019&time=23:18&language=de&start=Moorburger%20Ring&startCity=Hamburg&startType=STATION&destination=Informatikum&destinationCity=Hamburg&destinationType=STATION&timeIsDeparture=1&wayBy=train",
	"Gym-S":        "https://geofox.hvv.de/jsf/home.seam?execute=true&date=29.03.2019&time=23:18&language=de&start=Moorburger%20Ring&startCity=Hamburg&startType=STATION&destination=Francoper%20Straße&destinationCity=Hamburg&destinationType=STATION&timeIsDeparture=1&wayBy=train",
}

var key string

func APICall(k string) ([]Traffic, error) {
	if _, ok := Urls[k]; !ok {
		return nil, errors.New("error: no match found")
	}

	key = k

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var res string
	err = c.Run(ctxt, text("c-schedule-table", &res))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("overview: %s", res)

	/*
	//writing result in a text file, only for maint purposes

	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(emptyFile)
	defer emptyFile.Close()

	err = ioutil.WriteFile("empty.txt", []byte(res), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
*/
	r := transform(res)

	return r, nil
}

func text(s string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(Urls[key]),
		chromedp.WaitVisible(s, chromedp.BySearch),
		chromedp.WaitReady(s, chromedp.BySearch),
		chromedp.Text(s, res),
	}
}

//------------------------------------------------------------------------------------------------------------

func transform(rcvd string) []Traffic {
	var ts []Traffic
	var e Traffic

	//delete traffic information and split in individual strings[]
	rcvd = replaceString(rcvd, "", -1, "Für diese Fahrt liegen Hinweise vor.")
	split := strings.Split(rcvd, "€")
	for _, v := range split {
		sd := strings.Split(v, " ")
		sd, err := formatString(sd)
		if err != nil {
			continue
		}
		//fmt.Println(sd)
		e = Traffic{
			sd[0],
			sd[1],
			sd[2 : len(sd)-1],
			sd[len(sd)-1],
			key,
		}
		ts = append(ts, e)
	}

	//fmt.Println(ts)
	return ts
}

//formatting given string
func formatString(s []string) ([]string, error) {
	var rs []string

	for _, str := range s {
		str = replaceString(str, "", -1, " ", "Bus", "-bis")

		//fmt.Println(len(s))
		//fmt.Println(len(s))

		//some names are being repeated, getting rid of them
		n := len(str) / 2
		if str[:n] == str[n:] {
			str = str[:n]
		}

		//only take valid elements that are not empty
		if !(len(strings.TrimSpace(str)) == 0) {
			rs = append(rs, str)
		}
	}

	//delete last two elements, no needed (price & stops)
	if len(rs) < 1 {
		return rs, errors.New("slice bounds out of range ")
	}
	n := len(rs)
	rs = rs[:n-2]

	//last index is sometimes messy, traffic and needed time are mixed in some cases (S3S3S21S210:46)
	//splitting them apart in (S3 S3 S21 S21 0:46) and them removing doubles
	var temp []string
	for _, str := range rs {
		if len(str) > 6 {
			//traffic
			str1 := str[:len(str)-4]
			//duration
			str2 := str[len(str)-4:]

			var h string
			for _, v := range str1[:] {
				if strings.Contains(string(v), "S") {
					if !(len(strings.TrimSpace(string(v))) == 0) {
						temp = append(temp, h)
						h = string(v)
					}
				} else {
					h += string(v)
				}
			}
			temp = append(temp, str2)
		} else {
			temp = append(temp, str)
		}
	}

	var r []string
	for _, str := range temp {
		if !(len(strings.TrimSpace(str)) == 0) {
			r = append(r, str)
		}
	}

	r = sliceUniqMaps(r)
	return r, nil
}

//replaces the given words of a string
func replaceString(s string, with string, i int, element ... string) string {
	for _, v := range element {
		s = strings.Replace(s, v, with, i)
	}

	return s
}

//returns a slice with no duplicates
func sliceUniqMaps(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0

	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
