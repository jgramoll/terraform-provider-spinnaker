package provider

import "github.com/jgramoll/terraform-provider-spinnaker/client"

type locked struct {
	UI            bool   `mapstructure:"ui"`
	Description   string `mapstructure:"description"`
	AllowUnlockUI bool   `mapstructure:"allow_unlock_ui"`
}

type lockedArray []locked

func (array lockedArray) toClientLocked() *client.Locked {
	if len(array) > 0 {
		l := array[0]
		return &client.Locked{
			UI:            l.UI,
			Description:   l.Description,
			AllowUnlockUI: l.AllowUnlockUI,
		}
	}

	return nil
}

func fromClientLocked(clientLocked *client.Locked) lockedArray {
	if clientLocked == nil {
		return lockedArray{}
	}
	return []locked{
		{
			UI:            clientLocked.UI,
			Description:   clientLocked.Description,
			AllowUnlockUI: clientLocked.AllowUnlockUI,
		},
	}
}
