package video12

import (
	"context"
	"github.com/evt/video12/scheduler"
	"log"
	"net/http"

	"github.com/evt/video12/config"
	"github.com/evt/video12/db"
	"github.com/evt/video12/db/migrations"
	"github.com/evt/video12/server"
)

var s *server.Server

func init() {
	ctx := context.Background()

	// config
	cfg := config.Get()

	pgDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run Postgres migrations
	if err := migrations.Run(pgDB); err != nil {
		log.Fatal(err)
	}

	// create google cloud scheduler client
	sch, err := scheduler.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create new server instance
	s = server.Init(ctx, cfg, pgDB, sch)
}

// ScheduleCall
func ScheduleCall(w http.ResponseWriter, r *http.Request) {
	s.ScheduleCall(w, r)
}


// CallRoom
func CallRoom(w http.ResponseWriter, r *http.Request) {
	s.CallRoom(w, r)
}
