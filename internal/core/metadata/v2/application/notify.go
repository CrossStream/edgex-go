package application

import (
	"context"

	"net/http"

	"github.com/edgexfoundry/go-mod-bootstrap/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/di"
	v2HttpClient "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// addDeviceCallback invoke device service's callback function fot adding new device
func addDeviceCallback(ctx context.Context, dic *di.Container, device dtos.Device) {
	lc := container.LoggingClientFrom(dic.Get)
	ds, err := DeviceServiceByName(device.ServiceName, ctx, dic)
	if err != nil {
		lc.Errorf("fail to query device service by serviceName %s, err: %v", device.ServiceName, err)
		return
	}
	deviceServiceCallbackClient := v2HttpClient.NewDeviceServiceCallbackClient(ds.BaseAddress)
	response, err := deviceServiceCallbackClient.AddDeviceCallback(ctx, requests.AddDeviceRequest{Device: device})
	if err != nil {
		lc.Infof("fail to invoke device service callback for adding device %s, err: %v", device.Name, err)
	}
	if response.StatusCode != http.StatusOK {
		lc.Infof("fail to invoke device service callback for adding device %s, err: %s", device.Name, response.Message)
	}
}

// updateDeviceCallback invoke device service's callback function fot updating device
func updateDeviceCallback(ctx context.Context, dic *di.Container, device dtos.UpdateDevice) {
	lc := container.LoggingClientFrom(dic.Get)
	ds, err := DeviceServiceByName(*device.ServiceName, ctx, dic)
	if err != nil {
		lc.Errorf("fail to query device service by serviceName %s, err: %v", *device.ServiceName, err)
		return
	}
	deviceServiceCallbackClient := v2HttpClient.NewDeviceServiceCallbackClient(ds.BaseAddress)
	response, err := deviceServiceCallbackClient.UpdateDeviceCallback(ctx, requests.UpdateDeviceRequest{Device: device})
	if err != nil {
		lc.Infof("fail to invoke device service callback for updating device %s, err: %v", *device.Name, err)
		return
	}
	if response.StatusCode != http.StatusOK {
		lc.Infof("fail to invoke device service callback for updating device %s, err: %s", *device.Name, response.Message)
	}
}

// deleteDeviceCallback invoke device service's callback function fot deleting device
func deleteDeviceCallback(ctx context.Context, dic *di.Container, device models.Device) {
	lc := container.LoggingClientFrom(dic.Get)
	ds, err := DeviceServiceByName(device.ServiceName, ctx, dic)
	if err != nil {
		lc.Errorf("fail to query device service by serviceName %s, err: %v", device.ServiceName, err)
		return
	}
	deviceServiceCallbackClient := v2HttpClient.NewDeviceServiceCallbackClient(ds.BaseAddress)
	response, err := deviceServiceCallbackClient.DeleteDeviceCallback(ctx, device.Id)
	if err != nil {
		lc.Infof("fail to invoke device service callback for deleting device %s, err: %v", device.Name, err)
		return
	}
	if response.StatusCode != http.StatusOK {
		lc.Infof("fail to invoke device service callback for deleting device %s, err: %s", device.Name, response.Message)
	}
}
