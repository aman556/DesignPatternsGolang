package main

import "fmt"

type mobile interface {
	chargeAppleMobile()
}

type APPLEMobile struct{}

func (a *APPLEMobile) chargeAppleMobile() {
	fmt.Println("Apple mobile charging.")
}

type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}

type adapterAndroid struct {
	ap *adapteeAndroid
}

func (a *adapterAndroid) chargeAppleMobile() {
	a.ap.chargeAndroidMobile()
}

type adapteeAndroid struct{}

func (a *adapteeAndroid) chargeAndroidMobile() {
	fmt.Println("Android mobile charging.")
}

func main() {
	apple := &APPLEMobile{}
	client := &client{}
	client.chargeMobile(apple)

	adapteeAndroid := &adapteeAndroid{}
	android := &adapterAndroid{
		ap: adapteeAndroid,
	}
	client.chargeMobile(android)

}
