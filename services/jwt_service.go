package services

import (
	"fmt"
	"time"

	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "pcams",
		issure:    "pcams",
	}
}

type Claim struct {
	Sum uint `json:"sum`
	jwt.StandardClaims
}

type ClaimAccessCam struct {
	Cams []uint `json:"cams`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil

}

func (s *jwtService) GenerateTokenCamAccess(camAccessPermission models.CamAccessPermission) (string, error) {

	var expiresAt time.Time

	if camAccessPermission.DurationPermitions != 0 {
		camAccessPermission.DurationPermitions = camAccessPermission.DurationPermitions * 60000000000
		expiresAt = camAccessPermission.StartPermissionDate.Add(time.Minute * time.Duration(time.Duration(camAccessPermission.DurationPermitions).Minutes()))
	} else {
		expiresAt = camAccessPermission.StopPermissionDate
		camAccessPermission.DurationPermitions = uint64(camAccessPermission.StopPermissionDate.Sub(camAccessPermission.StartPermissionDate))
	}
	durationpermitionsMinutes := time.Duration(camAccessPermission.DurationPermitions).Minutes()
	claim := &ClaimAccessCam{
		camAccessPermission.CamId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    s.issure,
			IssuedAt:  camAccessPermission.StartPermissionDate.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}
	for _, c := range camAccessPermission.CamId {
		db := database.GetDataBase()
		err = db.QueryRow(
			`INSERT INTO cam_access_permission (token, alias, camid, startpermissiondate, stoppermissiondate, durationpermitions, userid) VALUES($1, $2, $3, $4, $5, $6, $7)
			RETURNING id`, t,
			camAccessPermission.Alias,
			c,
			camAccessPermission.StartPermissionDate,
			expiresAt,
			int(durationpermitionsMinutes),
			camAccessPermission.UserId,
		).Scan(&camAccessPermission.ID)
		database.CheckError(err)
	}

	return t, nil

}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil

}
