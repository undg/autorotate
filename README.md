# autorotate

`autorotate` is Linux/X11 utility for 2in1 laptops and other devices with touchscreen. No ROOT privileges necessary!

## Install
Download and unzip [autorotate.zip](https://github.com/undg/autorotate/releases/latest). Give execution permission with command:
```
cd bin/
chmod +x autocommand
```

## Usage 
Command `left` `right` `portrait` and `lanscape` are wrappers around xrandr and xinput to synchronize rotation of screen and digitizer. In fact they will rotate all xinput devices including stylus, touchpad, touchscreen, mouse and ~~webcam~~.

Laptops with axis sensor can use `auto` command to detect orientation. Flag `--daemon` is designed to run in background as a service.


## Service (systemd)

To install service, copy file `_config/systemd/user/autorotate.service` to `.config/systemd/user/autorotate.service`

In line starting with `ExecStart` edit `--display` part if necessary.

If you want to run service after PC reboot run command:

`systemctl --user enable rotate.service`

Start it manually with

`systemctl --user start rotate.service`

## Roadmap
- [x] Synchronize display and digitizer
- [x] Run as daemon in background
- [ ] Config to store chosen screen name and input devices
- [ ] Script to install systemd service
- [ ] Publish in AUR
