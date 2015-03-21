package main

import "fmt"
import "net/url"
import "github.com/kdvh/whatapi"
import "./lib/config"

func main() {
	cfg := config.Get("config.ini")

	fmt.Printf("Whatcd:\t\t\t%#v\n", cfg.Whatcd)
	fmt.Printf("Transmission:\t\t%#v\n", cfg.Transmission)
	fmt.Printf("Googleplaymusic:\t%#v\n", cfg.Googleplaymusic)

	wcd := whatapi.NewSite("https://what.cd/")
	wcd.Login(cfg.Whatcd.Username, cfg.Whatcd.Password)

	account := wcd.GetAccount()
	fmt.Printf("\nRatio What.cd: %f\n", account.UserStats.Ratio)

	parameters := url.Values{}
	parameters.Add("format", "MP3")
	response := wcd.SearchTorrents("django unchained", parameters)

	for i := range response.Results {
		group := response.Results[i]
		fmt.Println("************************************")
		fmt.Printf("%s  ||  %s  ||  %d\n", group.Artist, group.GroupName, group.GroupYear)
		fmt.Println("------------------------------------")
		fmt.Println("ID  ||  Format  ||  Size  ||  Seeders  ||  Leechers  ||  Freeleech")
		fmt.Println("------------------------------------")
		for j := range group.Torrents {
			torrent := group.Torrents[j]
			fmt.Println("------------------------------------")
			fmt.Printf("%d", torrent.TorrentID)
			fmt.Printf("  ||  %s %s", torrent.Format, torrent.Encoding)
			fmt.Printf("  ||  %d Mo", torrent.Size / 1000000)
			fmt.Printf("  ||  %d", torrent.Seeders)
			fmt.Printf("  ||  %d", torrent.Leechers)
			if torrent.IsFreeleech {
				fmt.Println("  ||  YES")
			} else {
				fmt.Println("  ||  NO")
			}
		}
		fmt.Println("------------------------------------")
	}
	fmt.Println("************************************")
}
