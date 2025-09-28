package driver

import (
	"time"
)

type Controller interface {
	RegisterRouter()
}

type RegistryContract interface {
	RunApplication()
}

func Run(rv RegistryContract) {
	if rv != nil {
		rv.RunApplication()
	}
}

type ApplicationData struct {
	AppName       string `json:"appName"`
	AppInstanceID string `json:"appInstanceID"`
	StartTime     string `json:"startTime"`
}

func NewApplicationData(appName, appInstanceID string) ApplicationData {
	return ApplicationData{
		AppName:       appName,
		AppInstanceID: appInstanceID,
		StartTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
}

// func (a ApplicationData) RunReporting(cfg *config.Config, publisher service.PublishMessageService) {

//if publisher == nil {
//	return
//}
//
//type Data struct {
//	Now time.Time `json:"now"`
//}
//
//data := &Data{
//	Now: time.Now(),
//}
//
//ticker := time.NewTicker(time.Duration(cfg.Reporting.IntervalPing) * time.Second)
//go func() {
//	for range ticker.C {
//		err := publisher.PublishMessage(context.Background(), 0, "event.health.iamstillalive", data)
//		if err != nil {
//			return
//		}
//	}
//}()

//ticker.Stop() //put it somewhere

// }
