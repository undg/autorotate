# autorotate

`autorotate` is Linux/X11 utility for 2in1 laptops and other devices with touchscreen.

It's a wrapper around xrandr and xinput to synchronize rotation of screen and digitizer. In fact it's rotate all input devices including stylus, touchpad, touchscreen, mouse...

Laptops with axis sensor can use `auto` command to auto detect orientation. Flag `--daemon` is designed to run in background as a service.


## Service (systemd)

To install service, copy `_config/systemd/user/autorotate.service` to `.config/systemd/user/autorotate.service` (@todo, create installation script)
In line starting with `ExecStart` edit `--display` part if necessary. (@todo, create config that will read display value)

Enable service if you want to run in after PC reboot.
`systemctl --user enable rotate.service`

Start it manually with
`systemctl --user start rotate.service`
