package services

import (
	"context"
	"errors"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam"
	"github.com/sirupsen/logrus"
	"go-glyph-v2/internal/core/dtos"
	"log"
	"strings"
	"sync"
	"time"
)

type GoSteamService struct {
	dotaClient      *dota2.Dota2
	steamLoginInfos []*steam.LogOnDetails
	counter         uint
	lock            sync.Mutex
}

func NewGoSteamService(usernames, passwords string) *GoSteamService {
	var steamLoginInfos []*steam.LogOnDetails
	u := strings.Split(usernames, " ")
	p := strings.Split(passwords, " ")
	for i := 0; i < len(u); i++ {
		steamLoginInfos = append(steamLoginInfos, &steam.LogOnDetails{
			Username: u[i],
			Password: p[i],
		})
	}
	steamLoginInfo := steamLoginInfos[0]

	dc, err := initDotaClient(steamLoginInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &GoSteamService{dotaClient: dc, steamLoginInfos: steamLoginInfos, counter: 1, lock: sync.Mutex{}}
}

func (s *GoSteamService) GetMatchFromGoSteamService(matchID int) (dtos.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // It's important to call cancel to release resources if operation completes before timeout.

	matchDetails, err := s.dotaClient.RequestMatchDetails(ctx, uint64(matchID))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = s.changeClient()
			if err != nil {
				return dtos.Match{}, err
			}
			return s.GetMatchFromGoSteamService(matchID)
		}
		return dtos.Match{}, err
	}

	return dtos.Match{
		ID: matchID,
		// TODO: add some validation to this, probably not xD
		Cluster:    int(matchDetails.Match.GetCluster()),
		ReplaySalt: int(matchDetails.Match.GetReplaySalt()),
	}, nil
}

func (s *GoSteamService) changeClient() error {
	s.lock.Lock()
	if s.counter >= uint(len(s.steamLoginInfos)) {
		s.counter = 0
	}
	dc, err := initDotaClient(s.steamLoginInfos[s.counter])
	if err != nil {
		return err
	}
	s.dotaClient = dc
	s.counter++
	s.lock.Unlock()
	return nil
}

func initDotaClient(steamLoginInfo *steam.LogOnDetails) (*dota2.Dota2, error) {
	sc := steam.NewClient()
	err := steam.InitializeSteamDirectory()
	if err != nil {
		return nil, err
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
	return dc, nil
}
