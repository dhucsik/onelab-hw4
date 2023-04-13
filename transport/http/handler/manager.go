package handler

import (
	"github.com/dhucsik/onelab-hw4/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
