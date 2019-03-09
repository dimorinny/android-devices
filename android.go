package android

import (
	"github.com/google/gousb"
	"log"
)

const (
	usbInterfaceAdbClass    gousb.Class    = 0xFF
	usbInterfaceAdbSubclass gousb.Class    = 0x42
	usbInterfaceAdbPortocol gousb.Protocol = 0x1
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
		// avoid device opening
		return false
	})

	return devices, err
}

func isAndroidDevice(description *gousb.DeviceDesc) bool {
	return description.Class == usbInterfaceAdbClass &&
		description.SubClass == usbInterfaceAdbSubclass &&
		description.Protocol == usbInterfaceAdbPortocol
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
