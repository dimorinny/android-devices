package android

import (
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

	devices, err := getDevicesDescriptions(context)
	if err != nil {
		return nil, err
	}

	var result []*Device
	for _, device := range devices {
		if isAndroidDevice(device) {
			result = append(
				result,
				mapLibUsbDevicesToInternalModel(device),
			)
		}
	}

	return result, nil
}

func getDevicesDescriptions(context *gousb.Context) ([]*gousb.DeviceDesc, error) {
	var devices []*gousb.DeviceDesc

	_, err := context.OpenDevices(func(description *gousb.DeviceDesc) bool {
		devices = append(devices, description)
		return false
	})

	return devices, err
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
