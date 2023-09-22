package svc

import "strconv"

type Server struct {
	Schema  string
	Host    string
	Port    uint16
	Context string
}

func (s Server) url() string {

	return s.Schema + "://" + s.Host + ":" + strconv.FormatUint(uint64(s.Port), 10) + s.Context

}
