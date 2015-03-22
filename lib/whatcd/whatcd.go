package whatcd

import "fmt"
import "../config"
import "net/url"
import "strings"
import "strconv"
import "github.com/kdvh/whatapi"

type Whatcd struct {
	site *whatapi.Site
}

func Init(config config.Config) Whatcd {
	var wcd Whatcd

	wcd.site = whatapi.NewSite("https://what.cd/")
	wcd.site.Login(config.Whatcd.Username, config.Whatcd.Password)

	return wcd
}

func (wcd *Whatcd) Ratio() {
	/*account := wcd.site.GetAccount()
	fmt.Printf("Ratio What.cd: %f\n", account.UserStats.Ratio)*/
	fmt.Printf("Ratio What.cd: -\n")
}

func (wcd *Whatcd) SearchTorrents(search string) {
	parameters := url.Values{}
	parameters.Add("format", "MP3")
	response := wcd.site.SearchTorrents(search, parameters)

	for i := range response.Results {
		group := response.Results[i]
		fmt.Println("\n")
		fmt.Printf("Artist:  %s\n", group.Artist)
		fmt.Printf("Group:   %s\n", group.GroupName)
		fmt.Printf("Year:    %d\n", group.GroupYear)
		fmt.Println("---------------------------------------------------------------------------------")
		fmt.Println("||     #ID     ||      FORMAT     ||   SIZE   ||  NB  ||   S   ||   L   ||  F  ||")
		fmt.Println("---------------------------------------------------------------------------------")
		for j := range group.Torrents {
			torrent := group.Torrents[j]
			fmt.Printf("||  %9d", torrent.TorrentID)
			fmt.Printf("  ||  %13s", torrent.Format + " " + torrent.Encoding)
			fmt.Printf("  ||  %3d Mo", torrent.Size / 1000000)
			fmt.Printf("  ||  %2d", torrent.FileCount)
			fmt.Printf("  ||  %3d", torrent.Seeders)
			fmt.Printf("  ||  %3d", torrent.Leechers)
			if torrent.IsFreeleech {
				fmt.Println("  ||  X  ||")
			} else {
				fmt.Println("  ||     ||")
			}
			fmt.Println("---------------------------------------------------------------------------------")
		}
	}
}

func (wcd *Whatcd) DetailTorrent(id int) {
	response := wcd.site.GetTorrent(id, url.Values{})
	group := response.Group
	torrent := response.Torrent

	fmt.Println("\n************ TRACK ************")
	fmt.Printf("Category:\t%s\n", group.CategoryName)
	fmt.Printf("Name:\t\t%s\n", group.Name)
	fmt.Printf("Year:\t\t%d\n", group.Year)
	fmt.Printf("Artist:\t\t")
	for i := range group.MusicInfo.Artists {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", group.MusicInfo.Artists[i].Name)
	}
	if len(group.MusicInfo.Artists) == 0 {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	fmt.Printf("   With:\t")
	for i := range group.MusicInfo.With {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", group.MusicInfo.With[i].Name)
	}
	if len(group.MusicInfo.With) == 0 {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	fmt.Println("\n************ FORMAT ************")
	fmt.Printf("Format:\t\t%s %s\n", torrent.Format, torrent.Encoding)
	fmt.Printf("Size:\t\t%d Mo\n", torrent.Size / 1000000)
	fmt.Printf("Files (%d):\n", torrent.FileCount)
	files := strings.Split(torrent.FileList, "|||")
	for i := range files {
		file := strings.Split(files[i], "{{{")
		size, _ := strconv.Atoi(strings.Split(file[1], "}}}")[0])
		fmt.Printf("   -> %6d Ko:  %s\n", size / 1000, file[0])
	}

	fmt.Println("\n************ TORRENT ************")
	fmt.Printf("Seeders:\t%d\n", torrent.Seeders)
	fmt.Printf("Leechers:\t%d\n", torrent.Leechers)
	fmt.Printf("Snatched:\t%d\n", torrent.Snatched)
	if torrent.FreeTorrent {
		fmt.Printf("FreeTorrent:\tYes\n")
	} else {
		fmt.Printf("FreeTorrent:\tNo\n")
	}
}

func (wcd *Whatcd) DownloadURL(id int) string {
	return wcd.site.CreateDownloadURL(id)
}
