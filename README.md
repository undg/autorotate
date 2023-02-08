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

Laptops with axis sensor can be run without any command to detect orientation. Flag `--daemon` is designed to run in background as a service.

Check display name with `list` command.

`autorotate list`

single run and exit:

`autorotate --display eDP`

`autorotate left --display eDP`

`autorotate help`

optionally to keep app running

`autorotate --daemon --display eDP`

![image](https://user-images.githubusercontent.com/5306983/217210748-93221f5f-8dab-4645-84b2-10505d149206.png)

## Service (systemd)

To install service, copy file

`cp _config/systemd/user/rotate.service ~/.config/systemd/user/rotate.service `

In line starting with `ExecStart` edit `--display` part if necessary.

If you want to run service after PC reboot run command:

`systemctl --user enable rotate.service`

Start it manually with

`systemctl --user start rotate.service`

![image](https://user-images.githubusercontent.com/5306983/217212712-5b81a5ab-6ab3-4abf-b628-b151339a9d0e.png)
