package vpn

import (
	// TODO: Which Wireguard lib?
	"golang.zx2c4.com/wireguard/wgctrl" // has docs: https://pkg.go.dev/golang.zx2c4.com/wireguard/wgctrl
	// "golang.zx2c4.com/wireguard" has no docs: https://pkg.go.dev/golang.zx2c4.com/wireguard
	// idk yet if wgctrl can do everything we need it to
	// https://www.wireguard.com/xplatform/
)

func Create() (*wgctrl.Client, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	return client, nil
}
