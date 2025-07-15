package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/devlup-labs/Libr/core/db/internal/node"
)

func main() {
	// config.InitConnection()

	url := "https://raw.githubusercontent.com/cherry-aggarwal/LIBR/refs/heads/main/docs/database_ips.csv"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch CSV: %v", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	var ip string
	var port string

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("skipping bad row: %v", err)
			continue
		}

		if len(row) < 2 {
			log.Printf("skipping row with too few columns: %v", row)
			continue
		}

		ip := row[0]
		port := row[1]

		fmt.Printf("Using IP=%q, Port=%q\n", ip, port)
		break
		// call something like:
		// if node is not working skip to next
	}

	address := ip + ":" + port

	localNode := &node.Node{
		NodeId: node.GenerateNodeID(address),
		IP:     ip,
		Port:   port,
	}

	fmt.Println(ip)
	fmt.Println(port)
	fmt.Println(localNode)

	// id := node.GenerateNodeID(address)
	// fmt.Println(hex.EncodeToString(id[:]))

	// fmt.Println(localNode)

	// rt := routing.GetOrCreateRoutingTable(localNode)
	// fmt.Println("Routing table created with port:", rt.SelfPort)
	// // Optional: Bootstrap to known node
	// bootstrapAddr := os.Getenv("BOOTSTRAP")
	// if bootstrapAddr != "" {
	// 	fmt.Println("Bootstrapping with", bootstrapAddr)
	// 	network.Bootstrap(bootstrapAddr, localNode, rt)
	// }

	// server.SetupRoutes(localNode, rt)
	// data, err := json.MarshalIndent(rt, "", "  ")
	// if err != nil {
	// 	log.Println("Error marshalling routing table:", err)
	// 	return
	// }
	// fmt.Println(string(data))
	// fmt.Println("Kademlia node running at http://" + address)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	fmt.Println("Failed to start server:", err)
	// }
	// fmt.Println(rt)
}
