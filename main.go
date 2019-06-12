package main

import (
	"Go-MagicMirror/api"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	tpl *template.Template
)

func numbers(w http.ResponseWriter, req *http.Request) {
	data, _ := http.Get("http://numbersapi.com/random/trivia")
	defer data.Body.Close()

	bodyBytes, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "data: %v\n\n", string(bodyBytes))

}

//---------------------------------------testing javascript EventSource------------------
type Notification struct {
	Message []string
}

var notifier chan *Notification

func notification(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %v", req.RemoteAddr)
	notifier = make(chan *Notification)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	go l()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(<-notifier)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
	fmt.Printf(buf.String())
	fmt.Println("done")

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

}
func l() {
	for {
		s := &Notification{
			Message: []string{strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())},
		}

		notifier <- s
	}
}

//---------------------------------------------------------

func hvv(w http.ResponseWriter, req *http.Request) {
	//Dummy Data
	dummy := `{"realtimeSchedules":[{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":41,"footpathTime":4,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"13:54"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"13:58"}},"line":{"name":"340","direction":"Ehestorfer Heuweg","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:01"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"14:28"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"14:30"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"14:35"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]},{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":50,"footpathTime":13,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"14:05"}},"to":{"name":"Rehrstieg","arrTime":{"date":"12.06.2019","time":"14:14"}},"line":{"name":"Fußweg","direction":"","origin":"","type":{"simpleType":"FOOTPATH","shortInfo":""}}},{"from":{"name":"Rehrstieg","depTime":{"date":"12.06.2019","time":"14:14"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"14:16"}},"line":{"name":"251","direction":"Heykenaukamp (Kehre)","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:21"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"14:48"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"14:50"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"14:55"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]},{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":47,"footpathTime":5,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"14:18"}},"to":{"name":"S Neugraben","arrTime":{"date":"12.06.2019","time":"14:25"}},"line":{"name":"340","direction":"S Neugraben","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neugraben","depTime":{"date":"12.06.2019","time":"14:29"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"14:58"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"15:00"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"15:05"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]},{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":50,"footpathTime":13,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"14:25"}},"to":{"name":"Rehrstieg","arrTime":{"date":"12.06.2019","time":"14:34"}},"line":{"name":"Fußweg","direction":"","origin":"","type":{"simpleType":"FOOTPATH","shortInfo":""}}},{"from":{"name":"Rehrstieg","depTime":{"date":"12.06.2019","time":"14:34"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"14:36"}},"line":{"name":"251","direction":"Finkenwerder (Fähre)","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:41"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"15:08"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"15:10"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"15:15"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]},{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":41,"footpathTime":4,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"14:44"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"14:48"}},"line":{"name":"340","direction":"Ehestorfer Heuweg","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:51"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"15:18"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"15:20"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"15:25"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]}]}`
	time.Sleep(4 * time.Second)
	fmt.Fprintf(w, dummy)

	//var request = &api.HVVGetRouteRequest{
	//	Origin:       api.Station{Name: "Moorburger Ring"},
	//	Destinations: api.Station{Name: "Schlump"},
	//	DateTime:     api.DateTime{Date: "12.06.2019", Time: "14:00"},
	//	Language:     api.GERMAN,
	//	Amount:       4,
	//	Apikey:       api.APIKEY_HVV,
	//	Username:     api.APIKEY_HVV_USER,
	//}
	//
	//resp, err := c.HVVGetRoute(request)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//json, err := json.Marshal(resp)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//fmt.Fprintf(w, string(json))
}

func googleCalender(w http.ResponseWriter, req *http.Request) {
	//Dummy Data
	dummy := `{"items":[{"start":{"date":"2019-06-10"},"summary":"Pfingstferien"},{"start":{"date":"2019-06-17"},"summary":"MPI, OpenMP"},{"start":{"date":"2019-06-24"},"summary":"SiW - Machine Learning"},{"start":{"date":"2019-06-24"},"summary":"Mathe Bonusklausur"},{"start":{"date":"2019-06-27"},"summary":"Sommerferien"},{"start":{"date":"2019-07-11"},"summary":"AOK Betrag"},{"start":{"date":"2019-07-17"},"summary":"Mathe Klausur"},{"start":{"date":"2019-07-18"},"summary":"FGI Klausur"}]}`
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, dummy)

	//var request api.GoogleCalenderRequest
	//resp, err := c.GoogleCalender(&request)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//json, err := json.Marshal(resp)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//fmt.Println("GOOGLECALENDER Done")
	//fmt.Fprintf(w, string(json))
}

func fuel(w http.ResponseWriter, req *http.Request) {

	//Dummy Data
	dummy := `{"stations":[{"brand":"Elan","street":"Neuwiedenthaler Str.","houseNumber":"122","diesel":1.159,"e5":1.439,"e10":1.419,"isOpen":true},{"brand":"Star","street":"Neuwiedenthaler Straße","houseNumber":"131","diesel":1.159,"e5":1.439,"e10":1.419,"isOpen":true},{"brand":"Oil!","street":"Cuxhavener Strasse 123","houseNumber":"","diesel":1.159,"e5":1.439,"e10":1.419,"isOpen":true},{"brand":"Shell","street":"Cuxhavener Str. 361","houseNumber":"","diesel":1.169,"e5":1.449,"e10":1.429,"isOpen":true},{"brand":"Total","street":"Cuxhavener Str.","houseNumber":"380","diesel":1.169,"e5":1.449,"e10":1.429,"isOpen":true}]}`
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, dummy)

	//var tk = &api.TankerkoenigRequest{
	//	Lon:    api.Longitude,
	//	Lat:    api.Latitude,
	//	Radius: api.Radius,
	//	GasTyp: api.GasTypAll,
	//	ApiKey: api.APIKEY_TANKERKOENIG,
	//}
	//
	//resp, err := c.Tankerkoenig(tk)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//json, err := json.Marshal(resp)
	//if err != nil {
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//
	//fmt.Println(string(json))
	//fmt.Println("Fuelprice Done")
	//fmt.Fprintf(w, string(json))
}

func weather(w http.ResponseWriter, req *http.Request) {
	//Dummy data
	dummy := `{"name":"Heimfeld","dt":1560373329,"weather":[{"id":741,"main":"Fog","description":"mostly Sunny","icon":"50n"}],"main":{"temp":15.45,"temp_min":13.33,"temp_max":17.22,"humidity":87},"clouds":{"all":72},"sys":{"sunrise":1560307904,"sunset":1560368903}}`
	fmt.Fprintf(w, dummy)
}

func forecast(w http.ResponseWriter, req *http.Request) {
	dummy := `{"list":[{"dt":1560384000,"main":{"temp":13.26,"temp_min":11.85,"temp_max":13.26,"humidity":95},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":37}},{"dt":1560394800,"main":{"temp":11.51,"temp_min":10.45,"temp_max":11.51,"humidity":94},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":1}},{"dt":1560405600,"main":{"temp":15.15,"temp_min":14.45,"temp_max":15.15,"humidity":83},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":30}},{"dt":1560416400,"main":{"temp":19.2,"temp_min":18.85,"temp_max":19.2,"humidity":58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":4}},{"dt":1560427200,"main":{"temp":21.31,"temp_min":21.31,"temp_max":21.31,"humidity":50},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":2}},{"dt":1560438000,"main":{"temp":21.96,"temp_min":21.96,"temp_max":21.96,"humidity":49},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0}},{"dt":1560448800,"main":{"temp":20.17,"temp_min":20.17,"temp_max":20.17,"humidity":64},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":2}},{"dt":1560459600,"main":{"temp":15.01,"temp_min":15.01,"temp_max":15.01,"humidity":77},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":2}},{"dt":1560470400,"main":{"temp":12.79,"temp_min":12.79,"temp_max":12.79,"humidity":85},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":1}},{"dt":1560481200,"main":{"temp":12,"temp_min":12,"temp_max":12,"humidity":81},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0}},{"dt":1560492000,"main":{"temp":17.35,"temp_min":17.35,"temp_max":17.35,"humidity":68},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":16}},{"dt":1560502800,"main":{"temp":24.05,"temp_min":24.05,"temp_max":24.05,"humidity":53},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"clouds":{"all":76}},{"dt":1560513600,"main":{"temp":22.35,"temp_min":22.35,"temp_max":22.35,"humidity":69},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":82}},{"dt":1560524400,"main":{"temp":26.96,"temp_min":26.96,"temp_max":26.96,"humidity":54},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":45}},{"dt":1560535200,"main":{"temp":21.86,"temp_min":21.86,"temp_max":21.86,"humidity":82},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":36}},{"dt":1560546000,"main":{"temp":18.17,"temp_min":18.17,"temp_max":18.17,"humidity":91},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":81}},{"dt":1560556800,"main":{"temp":17.85,"temp_min":17.85,"temp_max":17.85,"humidity":94},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":81}},{"dt":1560567600,"main":{"temp":17.51,"temp_min":17.51,"temp_max":17.51,"humidity":96},"weather":[{"id":501,"main":"Rain","description":"moderate rain","icon":"10d"}],"clouds":{"all":100}},{"dt":1560578400,"main":{"temp":18.2,"temp_min":18.2,"temp_max":18.2,"humidity":90},"weather":[{"id":501,"main":"Rain","description":"moderate rain","icon":"10d"}],"clouds":{"all":100}},{"dt":1560589200,"main":{"temp":23.11,"temp_min":23.11,"temp_max":23.11,"humidity":72},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":{"all":100}},{"dt":1560600000,"main":{"temp":22.27,"temp_min":22.27,"temp_max":22.27,"humidity":84},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":100}},{"dt":1560610800,"main":{"temp":22.45,"temp_min":22.45,"temp_max":22.45,"humidity":71},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"clouds":{"all":64}},{"dt":1560621600,"main":{"temp":18.13,"temp_min":18.13,"temp_max":18.13,"humidity":77},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":47}},{"dt":1560632400,"main":{"temp":13.04,"temp_min":13.04,"temp_max":13.04,"humidity":94},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"clouds":{"all":55}},{"dt":1560643200,"main":{"temp":14.47,"temp_min":14.47,"temp_max":14.47,"humidity":75},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"clouds":{"all":68}},{"dt":1560654000,"main":{"temp":11.19,"temp_min":11.19,"temp_max":11.19,"humidity":96},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":17}},{"dt":1560664800,"main":{"temp":14.69,"temp_min":14.69,"temp_max":14.69,"humidity":78},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":9}},{"dt":1560675600,"main":{"temp":19.08,"temp_min":19.08,"temp_max":19.08,"humidity":65},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0}},{"dt":1560686400,"main":{"temp":21.37,"temp_min":21.37,"temp_max":21.37,"humidity":58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":10}},{"dt":1560697200,"main":{"temp":21.85,"temp_min":21.85,"temp_max":21.85,"humidity":62},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":20}},{"dt":1560708000,"main":{"temp":19.59,"temp_min":19.59,"temp_max":19.59,"humidity":83},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":32}},{"dt":1560718800,"main":{"temp":15.44,"temp_min":15.44,"temp_max":15.44,"humidity":88},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":1}},{"dt":1560729600,"main":{"temp":14.27,"temp_min":14.27,"temp_max":14.27,"humidity":92},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0}},{"dt":1560740400,"main":{"temp":13.03,"temp_min":13.03,"temp_max":13.03,"humidity":94},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":5}},{"dt":1560751200,"main":{"temp":15.85,"temp_min":15.85,"temp_max":15.85,"humidity":90},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"clouds":{"all":51}},{"dt":1560762000,"main":{"temp":20.53,"temp_min":20.53,"temp_max":20.53,"humidity":61},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"clouds":{"all":51}},{"dt":1560772800,"main":{"temp":22.49,"temp_min":22.49,"temp_max":22.49,"humidity":57},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":39}},{"dt":1560783600,"main":{"temp":22.55,"temp_min":22.55,"temp_max":22.55,"humidity":58},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":42}},{"dt":1560794400,"main":{"temp":21.07,"temp_min":21.07,"temp_max":21.07,"humidity":70},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":23}},{"dt":1560805200,"main":{"temp":15.04,"temp_min":15.04,"temp_max":15.04,"humidity":87},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":3}}]}`
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, dummy)
}

var c *api.Client

func init() {
	tpl = template.Must(template.ParseGlob("client/templates/*"))

	//Creates a new Client
	c = api.NewClient()
}

func main() {
	router := newRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/notification/", notification)
	router.HandleFunc("/calender/", googleCalender)
	router.HandleFunc("/fuel/", fuel)
	router.HandleFunc("/hvv/", hvv)
	router.HandleFunc("/weather/", weather)
	router.HandleFunc("/forecast/", forecast)

	fmt.Println("Starting Server on port:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	//log.Fatal(http.ListenAndServe(":8080", limiter(router)))
}

//@todo mux router
//return a new http.ServerMux
func newRouter() *http.ServeMux {
	r := http.NewServeMux()

	//serving static files
	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.Handle("/client/css/", http.StripPrefix("/client/css", http.FileServer(http.Dir("client/css"))))
	r.Handle("/client/img/", http.StripPrefix("/client/img", http.FileServer(http.Dir("client/img"))))
	r.Handle("/client/js/", http.StripPrefix("/client/js", http.FileServer(http.Dir("client/js"))))

	return r
}

//@todo limiter gives error for static files
var limiter = rate.NewLimiter(1, 2)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
-------------------------------------This is just dummy data to fill the Webpage
*/

type f struct {
	Date      string
	Icon      string
	Humidity  string
	TopDegree string
	LowDegree string
	Rain      string
}

type traffic struct {
	Status  int
	Goal    string
	Time    string
	Address string
}

func index(w http.ResponseWriter, req *http.Request) {
	//----------------------

	t1 := traffic{
		0,
		"Arbeit",
		"10 min",
		"Neuwiedenthaler Straße 1",
	}

	t2 := traffic{
		1,
		"Geomatikum",
		"30 min",
		"Bundesstraße 55",
	}

	t3 := traffic{
		1,
		"Uni Hamburg",
		"35 min",
		"Grindelalle 3",
	}

	t4 := traffic{
		2,
		"Informatikum",
		"45 min",
		"Vogt-Informatikum 11",
	}

	//----------------------

	traffics := []traffic{t1, t2, t3, t4}

	data := struct {
		Traffic []traffic
	}{
		traffics,
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		panic(err)
	}

}
