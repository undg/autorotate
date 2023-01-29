# autorotate

`autorotate` is Linux utility for 2in1 devices with touchscreen that's run with Xorg.

It's a wrapper around xrandr and xinput to synchronize rotation of screen and digitizer. In fact it's rotate all input devices including stylus, touchpad, touchscreen, mouse...

Laptops with axis sensor can use `auto` command to auto detect orientation. Flag `--deamon` is designed to run as a service.
