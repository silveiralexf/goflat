package middleware

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

// SessionData holds the user session data which is used by any page that has <nav> bar
type SessionData struct {
	Username              string
	AvatarUrl             string
	IsGuest               bool
	IsAdmin               bool
	EnabledOauthProviders []string
	Pb                    *pocketbase.PocketBase
	ApisRequestInfo       *models.RequestInfo
	ApisAuthRecord        *models.Record
}

func NewSession(pb *pocketbase.PocketBase, ctx echo.Context) *SessionData {
	s := SessionData{}

	s.Pb = pb

	err := s.setOauthProviderNames()
	if err != nil {
		s.Pb.Logger().Warn(err.Error())
	}

	s.ApisRequestInfo = apis.RequestInfo(ctx)
	s.ApisAuthRecord = s.ApisRequestInfo.AuthRecord // nil if not authenticated as regular auth record

	if s.ApisRequestInfo.Admin == nil && s.ApisAuthRecord == nil {
		s.IsGuest = true
		return &s
	}

	if s.ApisRequestInfo.Admin != nil {
		s.IsAdmin = true
		s.Username = s.ApisRequestInfo.Admin.Email
		return &s
	}

	s.Username = s.ApisRequestInfo.AuthRecord.Username()
	err = s.setAvatarUrl()
	if err != nil {
		// TODO: later add guest avatar instead of failing
		s.Pb.Logger().Warn("no avatar was found, later set some pretty picture")
	}

	return &s
}

func (s *SessionData) setOauthProviderNames() error {
	oauthProviders := s.Pb.Settings().NamedAuthProviderConfigs()
	s.EnabledOauthProviders = make([]string, 0, len(oauthProviders))
	for name, config := range oauthProviders {
		if config.Enabled {
			s.EnabledOauthProviders = append(s.EnabledOauthProviders, name)
		}
	}

	if len(s.EnabledOauthProviders) == 0 {
		return fmt.Errorf("no valid oauth providers are configured")
	}
	return nil
}

func (s *SessionData) setAvatarUrl() error {
	user, err := s.Pb.Dao().FindAuthRecordByUsername("users", s.ApisAuthRecord.Username())
	if err != nil {
		// TODO: later set avatar URL for guest user
		return fmt.Errorf("failed to retrieve user record by username: %v", err)
	}

	avatarFileName := s.ApisAuthRecord.SchemaData()["avatar"]
	userBaseDir := user.BaseFilesPath()
	s.AvatarUrl = s.Pb.Settings().Meta.AppUrl + "/api/files/" + userBaseDir + "/" + fmt.Sprintf("%v", avatarFileName)

	return nil
}
