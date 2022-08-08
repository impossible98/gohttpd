package ip

import (
	// import third-party packages
	"github.com/andreaskoch/myip"
)

func LocalIP() string {
	remoteIPProvider, ipProviderError := myip.NewLocalIPProvider()

	if ipProviderError != nil {
		panic(ipProviderError)
	}

	remoteIPv4Addresses, remoteIPError := remoteIPProvider.GetIPv4Addresses()

	if remoteIPError != nil {
		panic(remoteIPError)
	}
	// return
	return remoteIPv4Addresses[0].String()
}
