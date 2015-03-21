package main

import "fmt"
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
}
