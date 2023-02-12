package db

import (
	"context"
	"database/sql"
)

type ControlResetData struct {
	Ip       string
	Port     int
	HostName string

	ClientIp  string
	ClientKey string
}

func SelectControlIp(conn *sql.Conn, ctx context.Context) (string, error) {
}

func SelectControlPort(conn *sql.Conn, ctx context.Context) (int, error) {
}

func SelectControlAddr(conn *sql.Conn, ctx context.Context) (ip string, port int, err error) {
	port, err = SelectControlPort(conn, ctx)
	if err != nil {
		return
	}
	ip, err = SelectControlIp(conn, ctx)
	return
}

func ResetControl() error {
}
