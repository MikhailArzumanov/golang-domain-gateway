package router

import (
	"fmt"
	"proxy/config"
)

func hostEntryToHostAddr(entry config.HostEntry) string {
	address := fmt.Sprintf("%s://%s:%s",
		entry.DestScheme, entry.DestHost, entry.DestPort,
	)
	return address
}
