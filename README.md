GROB: Forked from talkiepi, removed all the GPIO stuff to make it a little more simple


# -------------------------------------------- Talkiepi install process  
# --install required libs
sudo apt update
sudo apt -y install libopenal-dev libopus-dev libasound2-dev libopenal1 libopusfile-dev


# --download asound config to set default audio capture, note: this does not take affect till next reboot
sudo wget https://raw.githubusercontent.com/GR0B/talkiepi/master/conf/asoundrc/.asoundrc -O ~/.asoundrc
sudo chmod a+r ~/.asoundrc


# --download precompiled talkiepi client bin (else follow the Go compile instruction instead of this step)
mkdir ~/bin
cd ~/bin
sudo wget https://raw.githubusercontent.com/GR0B/talkiepi/master/bin-rpi/talkiepi -O /home/cs/bin/talkiepi
sudo chmod a+rx ~/bin/talkiepi 



# --download talkiepi mumble service and set to auto start
sudo wget https://raw.githubusercontent.com/GR0B/talkiepi/master/conf/systemd/mumble.service -O /etc/systemd/system/mumble.service
sudo chmod a+r /etc/systemd/system/mumble.service

# update user, client name and channel values
sudo nano /etc/systemd/system/mumble.service


# --enable service
sudo systemctl enable mumble.service
sudo systemctl start mumble.service


# --reboot to test, 
sudo reboot now







# talkiepi
![assembled1](doc/talkiepi_assembled_1.jpg "Assembled talkiepi 1")
![assembled2](doc/talkiepi_assembled_2.jpg "Assembled talkiepi 2")

talkiepi is a headless capable Mumble client written in Go, written for walkie talkie style use on the Pi using GPIO pins for push to talk and LED status.  It is a fork of [barnard](https://github.com/layeh/barnard), which was a great jump off point for me to learn golang and have something that worked relatively quickly.


## 3D printable enclosure

In the stl directory are the stl files for the enclosure I have designed specifically for the Raspberry Pi B+ board layout (I am using a Raspberry Pi 3 Model B) and the PCB and components from the [US Robotics USB Speakerphone](https://www.amazon.com/USRobotics-USB-Internet-Speakerphone-USR9610/dp/B000E6IL10/ref=sr_1_1?ie=UTF8&qid=1472691020&sr=8-1&keywords=us+robotics+speakerphone).
I will be posting a blog post shortly with a full component list and build guide.  For more information regarding building a full talkiepi device, go check out my blog at [projectable.me](http://projectable.me).


## Installing talkiepi

I have put together an install guide [here](doc/README.md).


## GPIO

You can edit your pin assignments in `talkiepi.go`
```go
const (
	OnlineLEDPin       uint = 18
	ParticipantsLEDPin uint = 23
	TransmitLEDPin     uint = 24
	ButtonPin          uint = 25
)
```

Here is a basic schematic of how I am currently controlling the LEDs and pushbutton:

![schematic](doc/gpio_diagram.png "GPIO Diagram")


## Pi Zero Fixes
I have compiled libopenal without ARM NEON support so that it works on the Pi Zero. The packages can be found in the [workarounds](/workarounds/). directory of this repo, install the libopenal1 package over your existing libopenal install.


## License

MPL 2.0

## Author

- talkiepi - [Daniel Chote](https://github.com/dchote)
- Barnard,Gumble Author - Tim Cooper (<tim.cooper@layeh.com>)

