package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/tarm/serial"
	"log"
)

func main() {
	info := accessory.Info{
		Name:         "csi0n",
		SerialNumber: "051AC-23AAM1",
		Manufacturer: "Apple",
		Model:        "AB",
	}
	acc := accessory.NewSwitch(info)

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("Client changed switch to on")
			sendStrCom("COM3",38400,[]byte{0x3A,0x00,0x01,0x0A,0x00,0x31,0x23},acc,false)
		} else {
			log.Println("Client changed switch to off")
			sendStrCom("COM3",38400,[]byte{0x3A,0x00,0x01,0x0A,0x01,0x30,0x23},acc,true)
		}
	})

	config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Println(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}

func sendStrCom(com string,baud int,b[] byte,acc *accessory.Switch,value bool)  {
	c := &serial.Config{Name: com, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
	}
	n, err := s.Write(b)
	if err != nil {
		log.Println(err)
	}
	if n==0 {
		log.Println("byte[] length is 0")
	}
	s.Close();
	acc.Switch.On.SetValue(!value)
	//buf := make([]byte, 5)
	//n, err = s.Read(buf)
	//log.Printf("start 4")
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("start 5")
	//log.Printf("%q", buf[:n])
}