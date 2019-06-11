package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tpl *template.Template
)

// This is just dummy data to fill the Webpage
type calender struct {
	Event string
	Date  string
}

type forecast struct {
	Date      string
	Icon      string
	Humidity  string
	TopDegree string
	LowDegree string
	Rain      string
}

type hvv struct {
	TotalTime string
	Line      []Hv
	Time      string
}

type Hv struct {
	From string
	To   string
	Line string
	Typ  string
}

type traffic struct {
	Status  int
	Goal    string
	Time    string
	Address string
}
type Stations struct {
	Brand  string
	Diesel float32
	E10    float32
	Street string
}

type ff struct {
	Stations []Stations
}

func init() {
	tpl = template.Must(template.ParseGlob("client/templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	a := calender{
		"Feiertag",
		"Heute um 14:00 Uhr",
	}

	b := calender{
		"Halloween",
		"Morgen um 07:00 Uhr",
	}

	c := calender{
		"Termin 1",
		"02. März um 11 Uhr",
	}

	d := calender{
		"Termin 2",
		"14. April um 23 Uhr",
	}

	e := calender{
		"Hochzeit von John",
		"24. Mai um 11 Uhr",
	}

	f := calender{
		"Geburtstagsfeier von Nick",
		"15. Juni um 12 Uhr",
	}

	g := calender{
		"Abschlussfeier",
		"30. August",
	}

	h := calender{
		"Semesterferien",
		"25. Dezemeber",
	}
	//-----------------------------------------------
	day1 := forecast{
		"9",
		"01d.png",
		"10%",
		"10°",
		"10°",
		"30%",
	}

	day2 := forecast{
		"12",
		"10d.png",
		"1%",
		"15°",
		"13°",
		"20%",
	}

	day3 := forecast{
		"17",
		"02n.png",
		"20%",
		"9°",
		"1°",
		"0%",
	}

	day4 := forecast{
		"Di",
		"09d.png",
		"90%",
		"15°",
		"10°",
		"30%",
	}

	day5 := forecast{
		"Mi",
		"01n.png",
		"5%",
		"28°",
		"10°",
		"50%",
	}

	day6 := forecast{
		"Do",
		"50d.png",
		"9%",
		"10°",
		"10°",
		"30%",
	}

	//-------------------------------

	h1 := hvv{
		"18:24-19:10",
		[]Hv{
			{
				"Moorburger Ring",
				"Neddduwiedenthal",
				"340",
				"Bus",
			},
			{
				"Neuwiedenthal",
				"Harburg",
				"S3",
				"Train",
			},
		},
		"0:46h",
	}

	h2 := hvv{
		"18:24-19:10",
		[]Hv{
			{
				"Moorburger Ring",
				"Neuwiedenthal",
				"340",
				"Bus",
			},
			{
				"Neuwiedenthal",
				"Harburg",
				"S3",
				"Train",
			},
			{
				"Harburg",
				"Pinneberg",
				"U3",
				"U",
			},
		},
		"0:20h",
	}
	//
	//h3 := hvv{
	//	"18:24-19:10",
	//	"S31-Altona",
	//	"1:00h",
	//	"S-Neuwiedenthal",
	//}
	//
	//h4 := hvv{
	//	"18:24-19:10",
	//	"S-Neuwiedenthal",
	//	"0:30h",
	//	"Moorburger Ring",
	//}
	//
	//h5 := hvv{
	//	"18:24-19:10",
	//	"Uni-Hamburg",
	//	"0:09h",
	//	"Home",
	//}
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

	f1 := Stations{
		"Aral",
		1.059,
		1.059,
		"Neuwiedenthaler Straße 11",
	}

	f2 := Stations{
		"Shell",
		1059,
		1059,
		"Neugrabener Markt 24",
	}

	f3 := Stations{
		"Elan",
		1059,
		1059,
		"Cuxhavener Straße 99",
	}

	f4 := Stations{
		"Total",
		1059,
		1059,
		"Neuwiedenthaler Straße 20",
	}

	forecasts := []forecast{day1, day2, day3, day4, day5, day6}
	calenders := []calender{a, b, c, d, e, f, g, h}
	hvvs := []hvv{h1, h2}
	traffics := []traffic{t1, t2, t3, t4}

	var tmp []Stations
	tmp = append(tmp, f1, f2, f3, f4)
	fuel := ff{tmp}
	//fp:= fuelprice.New(lat,long,radius,typ,sort)

	data := struct {
		Calender []calender
		Forecast []forecast
		HVV      []hvv
		Traffic  []traffic
		//FuelPrice fuelprice.FuelPrice
		FuelPrice ff
	}{
		calenders,
		forecasts,
		hvvs,
		traffics,
		fuel, //fp
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		panic(err)
	}

}

func main() {

	fmt.Println("Starting Server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	//serving static files
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/client/css/", http.StripPrefix("/client/css", http.FileServer(http.Dir("client/css"))))
	mux.Handle("/client/img/", http.StripPrefix("/client/img", http.FileServer(http.Dir("client/img"))))
	mux.Handle("/client/js/", http.StripPrefix("/client/js", http.FileServer(http.Dir("client/js"))))

	log.Fatal(http.ListenAndServe(":8080", mux))

}
