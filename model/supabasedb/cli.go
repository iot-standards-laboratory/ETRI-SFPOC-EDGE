package supabasedb

import (
	"etri-sfpoc-edge/model"

	"github.com/supabase-community/supabase-go"
)

var DefaultDB model.I_DBHandler = nil
var client *supabase.Client

func init() {
	var err error
	client, err = supabase.NewClient(
		"https://supabase.godopu.com",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyAgCiAgICAicm9sZSI6ICJhbm9uIiwKICAgICJpc3MiOiAic3VwYWJhc2UtZGVtbyIsCiAgICAiaWF0IjogMTY0MTc2OTIwMCwKICAgICJleHAiOiAxNzk5NTM1NjAwCn0.dc_X5iR_VP_qT0zsiyj_I_OZ2T9FtRU2BBNWN8Bu4GE",
		nil,
	)
	if err != nil {
		panic(err)
	}

	DefaultDB = &_supabaseDB{}
}
