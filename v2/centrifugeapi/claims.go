package centrifugeapi

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/centrifugal/centrifuge"
// 	"github.com/golang-jwt/jwt/v4"
// )

// type AuthTokenClaims struct {
// 	ID                   string   `json:"id"`   // 유저 ID
// 	Name                 string   `json:"name"` // 유저 이름
// 	Role                 []string `json:"role"` // 유저 역할
// 	Base                 string   `json:"base"` // base url
// 	Sub                  string   `json:"sub"`
// 	Channels             []string `json:"channels"`
// 	jwt.RegisteredClaims          // 표준 토큰 Claims
// }

// const TokenHmacSecret = "46b38493-147e-4e3f-86e0-dc5ec54f5133"

// func IssueJWT(id, name, role, base string, channels []string) (string, error) {
// 	claims := AuthTokenClaims{
// 		ID:       id,
// 		Name:     name,
// 		Role:     []string{role},
// 		Base:     base,
// 		Sub:      id,
// 		Channels: channels, // client side subscription시 nil 전달
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			// A usual scenario is to set the expiration time relative to the current time
// 			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString([]byte(TokenHmacSecret))

// 	return ss, err
// }

// func ParseJWTwithClaims(ss string) (*AuthTokenClaims, error) {
// 	claims := AuthTokenClaims{}

// 	_, err := jwt.ParseWithClaims(ss, &claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(TokenHmacSecret), nil
// 	})

// 	if err != nil {
// 		return &claims, err
// 	}

// 	return &claims, nil
// }

// func NewClient(id, name, role string) error {
// 	ss, _ := IssueJWT(id, name, role, fmt.Sprintf("/%s/", role), []string{"public/status"})
// 	node, err := centrifuge.New(centrifuge.Config{
// 		LogLevel:   centrifuge.LogLevelDebug,
// 		LogHandler: handleLog,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	broker, err := centrifuge.NewMemoryBroker(node, centrifuge.MemoryBrokerConfig{
// 		HistoryMetaTTL: 120 * time.Second,
// 	})

// 	node.SetBroker(broker)

// 	node.OnConnecting(func(ctx context.Context, e centrifuge.ConnectEvent) (centrifuge.ConnectReply, error) {
// 		log.Printf("client connecting with token: %s", e.Token)
// 		token, err := ParseJWTwithClaims(e.Token)
// 		if err != nil {
// 			if err == jwt.ErrTokenExpired {
// 				return centrifuge.ConnectReply{}, centrifuge.ErrorTokenExpired
// 			}
// 			return centrifuge.ConnectReply{}, centrifuge.DisconnectInvalidToken
// 		}
// 		subs := make(map[string]centrifuge.SubscribeOptions, len(token.Channels))
// 		for _, ch := range token.Channels {
// 			subs[ch] = centrifuge.SubscribeOptions{}
// 		}
// 		return centrifuge.ConnectReply{
// 			Credentials: &centrifuge.Credentials{
// 				UserID:   token.ID,
// 				ExpireAt: token.ExpiresAt.Unix(),
// 			},
// 			Subscriptions: subs,
// 		}, nil
// 	})

// 	node.OnConnect(func(client *centrifuge.Client) {
// 		transport := client.Transport()
// 		log.Printf("user %s connected via %s with protocol: %s", client.UserID(), transport.Name(), transport.Protocol())
// 	})

// 	if err := node.Run(); err != nil {
// 		log.Fatal(err)
// 	}

// }

// func handleLog(e centrifuge.LogEntry) {
// 	log.Printf("%s: %v", e.Message, e.Fields)
// }
