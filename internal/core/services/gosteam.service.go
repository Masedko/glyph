package services

import (
	"context"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam"
	"github.com/sirupsen/logrus"
	"go-glyph-v2/internal/core/dtos"
	"log"
	"time"
)

type GoSteamService struct {
	dotaClient *dota2.Dota2
}

func NewGoSteamService(username, password, twoFactorCode, authCode string) *GoSteamService {
	steamLoginInfo := new(steam.LogOnDetails)
	steamLoginInfo.Username = username
	steamLoginInfo.Password = password
	steamLoginInfo.TwoFactorCode = twoFactorCode
	steamLoginInfo.AuthCode = authCode
	sc := steam.NewClient()
	err := steam.InitializeSteamDirectory()
	if err != nil {
		log.Fatal(err)
	}
	dc := dota2.New(sc, logrus.New())
	go func(dc *dota2.Dota2) {
		for event := range sc.Events() {
			switch e := event.(type) {
			case *steam.ConnectedEvent:
				sc.Auth.LogOn(steamLoginInfo)
			case *steam.LoggedOnEvent:
				log.Println("Logged on to Steam")
				dc.SetPlaying(true)
				time.Sleep(5 * time.Second)
				dc.SayHello()
			case *steam.LogOnFailedEvent:
				log.Printf("LogOn failed. Reason: %v\n", e.Result)
			case *events.GCConnectionStatusChanged:
				log.Println("GCConnectionStatusChanged")
				isReady := e.NewState == protocol.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
				if !isReady {
					log.Println("GCConnectionStatusChanged: Not ready")
					dc.SayHello()
				}
			case steam.FatalErrorEvent:
				log.Print(e)
			case error:
				log.Print(e)
			}

		}
	}(dc)
	server := sc.Connect()
	log.Printf("Steam sc connected %s\n", server.String())

	return &GoSteamService{dotaClient: dc}
}

func (s GoSteamService) GetMatchFromGoSteamService(matchID int) (dtos.Match, error) {
	ctx := context.Background()
	matchDetails, err := s.dotaClient.RequestMatchDetails(ctx, uint64(matchID))
	if err != nil {
		return dtos.Match{}, err
	}

	return dtos.Match{
		ID: matchID,
		// TODO: add some validation to this, probably not xD
		Cluster:    int(matchDetails.Match.GetCluster()),
		ReplaySalt: int(matchDetails.Match.GetReplaySalt()),
	}, nil
}
