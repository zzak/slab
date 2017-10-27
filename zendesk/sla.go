package zendesk

import (
	c "github.com/tylerconlee/slab/config"
)

// CheckSLA will grab the tickets from GetAllTickets, parse the SLA fields and // compare them to the current time
func CheckSLA() {
	config := c.LoadConfig()
	tickets := GetAllTickets(config.Zendesk.User, config.Zendesk.APIKey, config.Zendesk.URL)
	Log.Debug(tickets)

}

// getPriorityLevel takes an individual ticket row from the Zendesk output and // returns a string of what priority level the ticket is tagged with
func getPriorityLevel() {

}
