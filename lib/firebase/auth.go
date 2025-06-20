package firebase

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/auth"
)

func VerifyIDToken(idToken string) (*auth.Token, error) {
	client, err := App.Auth(context.Background())
	if err != nil {
		fmt.Println("❌ Firebase Auth クライアントの取得に失敗:", err)
		log.Printf("Failed to get Auth client: %v", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		fmt.Println("❌ IDトークンの検証に失敗:", err)
		log.Printf("Failed to verify ID token: %v", err)
		return nil, err
	}

	fmt.Println("✅ IDトークンの検証に成功: UID =", token.UID)
	return token, nil
}
