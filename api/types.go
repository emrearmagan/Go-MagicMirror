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
)

const (
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
)

const (
	GERMAN  = Language("de")
	ENGLISH = Language("en")
	SPANISH = Language("sp")
)

const (
	Latitude  = Coords("53.48114369667324")
	Longitude = Coords("9.872239785725588")
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

type ApiConfig struct {
	host string
	path string
}

//---------------------------------------------------------------------------------------------------------------------

// commonResponse contains the common response fields to most API. This is used for all responses
//type CommonResponse struct {
//	//Status Tankerkoenig
//	Status Status `json:"status,omitempty"`
//	//Status Openweather
//	Status1 Status `json:"cod,omitempty"` //

//	// ErrorMessage is the explanatory field added when Status is an error.
//	//ErrorMessage for Tankerkoening and Openweather
//	ErrorMessage ErrorMessage `json:"message,omitempty"`

//	//ErrorMessage1 for Google DistanceMatrix
//	ErrorMessage1 ErrorMessage `json:"error_message,omitempty"`
//}

type Status interface{}
type ErrorMessage interface{}

// Since Status and Errormessage are called differently on most API calls, we sort them out here
//func (c *CommonResponse) StatusCheck() {
//	if c.Status == nil {
//		c.Status = c.Status1
//	}
//	if c.ErrorMessage == nil {
//		c.ErrorMessage = c.ErrorMessage1
//	}
//}
// defines an interface for all API Requests
// StatusError returns an error if this object has a Status different
//func (c *CommonResponse) StatusError() error {
//	switch status := c.Status.(type) {
//	case int:
//		if status != 200 {
//			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
//		}
//	case float64:
//		if status != 200 {
//			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
//		}
//	case string:
//		if status != "OK" && c.Status != "ok" && status != "ZERO_RESULTS" && c.Status != "200" {
//			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
//		}
//	}
//	return nil
//}
