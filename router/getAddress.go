package router

import "proxy/config"

func getAddress(
	table []config.HostEntry,
	domainHost string,
) string {
	address := ""
	for _, entry := range table {
		if entry.DomainKey == domainHost {
			address = hostEntryToHostAddr(entry)
		}
	}
	return address
}
