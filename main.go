package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/brutella/dnssd"
	_ "github.com/lib/pq"

	"github.com/go-ocf/go-coap"
	"github.com/thomas-fossati/gordo/model"
)

var g_m *model.Model

func main() {
	pgCfg := []string{
		"host=localhost",
		"port=15432",
		"user=postgres",
		"password=123",
		"dbname=gordo",
		"sslmode=disable",
	}

	db, err := sql.Open("postgres", strings.Join(pgCfg, " "))
	if err != nil {
		log.Fatal("DB init failed", err)
	}

	g_m = model.NewModel(db)

	mux := coap.NewServeMux()
	mux.Handle("/rd-lookup/res", coap.HandlerFunc(LookupRes))

	// start DNS-SD
	go dnsSDStart()

	log.Fatal(coap.ListenAndServe(":5683", "udp", mux))
}

func dnsSDStart() error {
	r, err := dnssd.NewResponder()
	if err != nil {
		return err
	}

	// read from DB and configure and add services to the responder
	for _, sv := range dnsSDInit() {
		_, err := r.Add(sv)
		if err != nil {
			return err
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.Respond(ctx)

	return nil
}

func dnsSDInit() []dnssd.Service {
	x, _ := dnssd.NewService(
		dnssd.Config{
			Domain: "local",
			Name:   "Spot",
			Type:   "_oic-d-light._udp",
			Port:   12345,
		},
	)

	// dummy
	return []dnssd.Service{x}
}

func LookupRes(w coap.ResponseWriter, req *coap.Request) {
	// query the data model
	rs, err := g_m.ResourceLookup(req.Msg.Query())
	if err != nil {
		fmt.Println("TODO send 5.xx: ", err)
	}

	w.SetContentFormat(coap.AppLinkFormat)

	// format the result into link-format
	payload := []byte(rs.LinkFormat())

	if _, err := w.Write(payload); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
