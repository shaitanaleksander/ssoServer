package repository

import (
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"log"
	"net"
)


func CreateSession(urlDB string) *mgo.Session {
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInf, err := mgo.ParseURL(urlDB)
	if err != nil {
		log.Fatal(err)
	}
	dialInf.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session ,err :=  mgo.DialWithInfo(dialInf)
	if err != nil {
		log.Fatal(err)
	}
	return session
}