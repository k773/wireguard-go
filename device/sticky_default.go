//go:build !linux

package device

import (
	"github.com/k773/wireguard-go/conn"
	"github.com/k773/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
