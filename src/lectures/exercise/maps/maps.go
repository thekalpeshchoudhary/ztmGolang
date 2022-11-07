//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
func dispayServerInfo(serverStatus map[string]int) {
	//  - number of servers
	fmt.Println("")
	fmt.Println("Total number of Servers:", len(serverStatus))
	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
	var (
		onlineServers, offlineServers, miantServers, retiredServers int
	)
	for _, serverStatus := range serverStatus {
		switch serverStatus {
		case Online:
			onlineServers += 1
		case Offline:
			offlineServers += 1
		case Maintenance:
			miantServers += 1
		case Retired:
			retiredServers += 1
		}
	}
	fmt.Println("---Server Status---")
	fmt.Println(" Online :", onlineServers, "\n Offline :", offlineServers, "\n Maintainence :", miantServers, "\n Retired", retiredServers)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	serverStatus := make(map[string]int)
	for _, eachServer := range servers {
		//* Set all of the server statuses to `Online` when creating the map
		serverStatus[eachServer] = Online
	}
	//* After creating the map, perform the following actions:
	//  - call display server info function
	dispayServerInfo(serverStatus)

	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired

	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline

	//  - call display server info function
	dispayServerInfo(serverStatus)

	//  - change server status of all servers to `Maintenance`
	for serverName := range serverStatus {
		serverStatus[serverName] = Maintenance
	}

	//  - call display server info function
	dispayServerInfo(serverStatus)

}
