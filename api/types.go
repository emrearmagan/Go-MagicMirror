package api

type Units string
type Mode string
type Language string
type Coords string
type GasTyp string

//--------------------------------------------------Google API----------------------------------------------------------
// Travel mode preferences.
const (
	TravelModeDriving   = Mode("driving")
	TravelModeWalking   = Mode("walking")
	TravelModeBicycling = Mode("bicycling")
	TravelModeTransit   = Mode("transit")

	AvoidTolls    = "tolls"
	AvoidHighways = "highways"
	AvoidFerries  = "ferries"
	AvoidIndoor   = "indoor"
)

const (
	// See, edit, share, and permanently delete all the calendars you can
	// access using Google Calendar
	CalendarScope = "https://www.googleapis.com/auth/calendar"

	// View and edit events on all your calendars
	CalendarEventsScope = "https://www.googleapis.com/auth/calendar.events"

	// View events on all your calendars
	CalendarEventsReadonlyScope = "https://www.googleapis.com/auth/calendar.events.readonly"

	// View your calendars
	CalendarReadonlyScope = "https://www.googleapis.com/auth/calendar.readonly"

	// View your Calendar settings
	CalendarSettingsReadonlyScope = "https://www.googleapis.com/auth/calendar.settings.readonly"
)

//---------------------------------------------------------------------------------------------------------------------
// Units to use on request.
const (
	UnitsMetric   = Units("metric")
	UnitsImperial = Units("imperial")

	GERMAN  = Language("de")
	ENGLISH = Language("en")
	SPANISH = Language("sp")

	Latitude  = Coords("53.460210")
	Longitude = Coords("9.951299")
)

//--------------------------------------------Tankerkoenig Api----------------------------------------------------------
const (
	Radius         = 2.0
	GasTypAll      = GasTyp("all")
	GasTypE5       = GasTyp("e5")
	GasTypE10      = GasTyp("e10")
	GasTypDiesel   = GasTyp("diesel")
	SortbyDistance = "dist"
	SortbyPrice    = "price"
)

type apiConfig struct {
	host string
	path string
}

type apiKey string

//---------------------------------------------------------------------------------------------------------------------

type Status interface{}
type ErrorMessage interface{}
