package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET")) // 環境変数で管理推奨

func GenerateToken(userID uint) (string, error) {
	fmt.Println("🔐 Generating JWT for userID:", userID)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("❌ Failed to sign token:", err)
		return "", err
	}

	fmt.Println("✅ JWT generated successfully:", signedToken)
	return signedToken, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	fmt.Println("🔍 Verifying JWT token:", tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("❌ Invalid signing method")
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("❌ Token verification failed:", err)
		return nil, err
	}

	fmt.Println("✅ Token verification succeeded")
	return token, nil
}

func VerifyTokenAndExtractUserID(tokenStr string) (uint, error) {
	fmt.Println("🔍 Extracting userID from token...")

	token, err := VerifyToken(tokenStr)
	if err != nil || !token.Valid {
		fmt.Println("❌ Invalid or failed-to-verify token:", err)
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("❌ Failed to assert JWT claims")
		return 0, jwt.ErrInvalidKey
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		fmt.Println("❌ user_id not found or invalid type in token claims")
		return 0, jwt.ErrInvalidKey
	}

	userID := uint(userIDFloat)
	fmt.Println("✅ Extracted userID from token:", userID)
	return userID, nil
}
