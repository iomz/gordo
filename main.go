package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"

	_ "github.com/lib/pq"

	"github.com/thomas-fossati/go-coap"
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
	mux.Handle("/rd-lookup/res", coap.FuncHandler(LookupRes))

	log.Fatal(coap.ListenAndServe("udp", ":5683", mux))
}

func LookupRes(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	// query the data model
	rs, err := g_m.ResourceLookup(m.Query())
	if err != nil {
		fmt.Println("TODO send 5.xx: ", err)
	}

	// format the result into link-format
	payload := []byte(rs.LinkFormat())

	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   payload,
		}
		res.SetOption(coap.ContentFormat, coap.AppLinkFormat)

		//	log.Printf("Transmitting from A %#v", res)
		return res
	}
	return nil
}
