package cli

import (
	gcli "github.com/urfave/cli"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
)

func emulatorCheckMessageSignatureCmd() gcli.Command {
	name := "emulatorCheckMessageSignature"
	return gcli.Command{
		Name:        name,
		Usage:       "Check a message signature matches the given address.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "message",
				Usage: "The message that the signature claims to be signing.",
			},
			gcli.StringFlag{
				Name:  "signature",
				Usage: "Signature of the message.",
			},
			gcli.StringFlag{
				Name:  "address",
				Usage: "Address that issued the signature.",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			message := c.String("message")
			signature := c.String("signature")
			address := c.String("address")
			deviceWallet.DeviceCheckMessageSignature(deviceWallet.DeviceTypeEmulator, message, signature, address)
		},
	}
}
