package types

type UserData struct {
	Bio               string `json:"bio"`
	Company           string `json:"company"`
	Email             string `json:"email"`
	FullName          string `json:"name"`
	Locale            string `json:"locale"`
	Location          string `json:"location"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Provider          string `json:"provider"`
	ProviderID        int    `json:"providerId"`
	Username          string `json:"login"`
	VerifiedEmail     bool   `json:"verified_email"`
}
