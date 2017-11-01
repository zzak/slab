package main

import (
	"time"

	"github.com/tylerconlee/slab/sla"

	Zen "github.com/tylerconlee/slab/zendesk"
)

func RunTimer(interval time.Duration) {
	t := time.NewTicker(interval * time.Minute)
	for {
		active := Zen.CheckSLA()
		log.Debug(active)
		sla.InitSLA()
		for _, ticket := range active {
			if ticket.Priority != nil {
				//timer := sla.GetTimer(ticket)
				sla.GetBreach(ticket)
			}
		}
		<-t.C
	}
}
