package config

type AppTeamMakerService struct {
	Port int `json:"port"`
}

type ApplicationServer struct {
	AppTeamMakerService AppTeamMakerService `json:"app_team_maker_api_service"`
}
