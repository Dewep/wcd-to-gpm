package gpmusic

import "fmt"
import "../config"
import "net/url"
import "strings"
import "strconv"
import "github.com/kdvh/whatapi"

type Gpmusic struct {
	Email string
}

func Init(config config.Config) Gpmusic {
	var gpm Gpmusic

	gmp.Email = config.Googleplaymusic.Username

	return gpm
}

func (gpm *Gpmusic) Upload(file string) {
	fmt.Println("Upload", file)
}
