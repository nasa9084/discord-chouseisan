package p

import (
	"crypto/ed25519"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

var applicationPublicKey ed25519.PublicKey

func init() {
	appPubKeyEnv := os.Getenv("APPLICATION_PUBLICKEY")
	if appPubKeyEnv == "" {
		panic("environment variable APPLICATION_PUBLICKEY is not set")
	}

	applicationPublicKey = ed25519.PublicKey([]byte(appPubKeyEnv))
}

// Handler is a http handler for Cloud Functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("incoming request")

	if !discordgo.VerifyInteraction(r, applicationPublicKey) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid request signature"))
		return
	}
}
