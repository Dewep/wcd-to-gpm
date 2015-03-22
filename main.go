package main

import "fmt"
import "flag"
import "strings"
import "./lib/config"
import "./lib/whatcd"

func search(cfg config.Config, search string) {
	wcd := whatcd.Init(cfg)

	fmt.Println()
	wcd.Ratio()
	fmt.Println()
	wcd.SearchTorrents(strings.Join(flag.Args(), " "))
	fmt.Println()
}

func detail(cfg config.Config, id int) {
	wcd := whatcd.Init(cfg)

	fmt.Println()
	wcd.Ratio()
	fmt.Println()
	wcd.DetailTorrent(id)
	fmt.Println()
}

func download(cfg config.Config, id int) {
	wcd := whatcd.Init(cfg)

	fmt.Println()
	wcd.Ratio()
	fmt.Println()
	fmt.Println(wcd.DownloadURL(id))
	fmt.Println()
}

func upload(cfg config.Config, files string[]) {
	gpm := gpmusic.Init(cfg)
	for i := range files {
		gpm.Upload(files[i])
	}
}

func main() {
	cfg := config.Get("config.ini")

	isSearch := flag.Bool("search", false, "Search a torrent: ./wcd-to-gmp -search django unchained")
	detailPtr := flag.Int("detail", 0, "Detail a torrent: ./wcd-to-gmp -detail 30647830")
	downloadPtr := flag.Int("download", 0, "Download a torrent: ./wcd-to-gmp -download 30647830")
	isUpload := flag.Bool("upload", false, "Upload a music: ./wcd-to-gmp -upload file.mp3")
	flag.Parse()

	if *isSearch {
		search(cfg, strings.Join(flag.Args(), " "))
	} else if *detailPtr > 0 {
		detail(cfg, *detailPtr)
	} else if *downloadPtr > 0 {
		download(cfg, *downloadPtr)
	} else if *isUpload {
		upload(cfg, flag.Args())
	}
}
