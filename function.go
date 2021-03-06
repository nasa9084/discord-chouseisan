package p

import (
	"crypto/ed25519"
	"encoding/hex"
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

	decoded, err := hex.DecodeString(appPubKeyEnv)
	if err != nil {
		panic("decoding APPLICATION_PUBLICKEY")
	}

	applicationPublicKey = ed25519.PublicKey(decoded)
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
