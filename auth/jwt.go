import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("SECRET_KEY")

type Claims struct {
  Username string `json:"username"`
  Role string `json:"role"`
  jwt.StandardClaims
}

func GenerateToken(user *User) (string, error) {
  claims := Claims{
    user.Username,
    user.Role,
    jwt.StandardClaims{
      // set exp time
    }
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
  token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
      return jwtKey, nil
  })

  if claims, ok := token.Claims.(*Claims); ok && token.Valid {
    return claims, nil
  } else {
    return nil, err 
  }
}
