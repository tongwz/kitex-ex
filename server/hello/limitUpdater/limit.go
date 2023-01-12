package limitUpdater

import "github.com/cloudwego/kitex/pkg/limit"

// MyLimiterUpdater define your limiter updater to update limit threshold
type MyLimiterUpdater struct {
	updater limit.Updater
}

func (lu *MyLimiterUpdater) YourChange() {
	// your logic: set new option as needed
	newOpt := &limit.Option{
		MaxConnections: 10000,
		MaxQPS:         2000,
	}
	// update limit config
	isUpdated := lu.updater.UpdateLimit(newOpt)
	// your logic
	if isUpdated {
		//lu.UpdateControl()
	}
}

func (lu *MyLimiterUpdater) UpdateControl(u limit.Updater) {
	lu.updater = u
}
