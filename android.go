package android

import (
	"errors"
	"github.com/google/gousb"
	"log"
)

func Devices() ([]*Device, error) {
	context := gousb.NewContext()
	defer func() {
		err := context.Close()
		if err != nil {
			log.Println("Android devices detector: failed to close context", err)
		}
	}()

	devices, err := context.OpenDevices(isAndroidDevice)

	defer func() {
		for _, device := range devices {
			_ := device.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []*Device
	for _, device := range devices {
		result = append(
			result,
			mapLibUsbDevicesToInternalModel(device.Desc),
		)
	}

	return nil, errors.New("error")
}

// TODO detect android device
func isAndroidDevice(description *gousb.DeviceDesc) bool {
	return true
}

// TODO: fetch description
func mapLibUsbDevicesToInternalModel(description *gousb.DeviceDesc) *Device {
	return &Device{
		Description: "",

		Bus:     description.Bus,
		Address: description.Address,

		Vendor:  int(description.Vendor),
		Product: int(description.Product),
	}
}
