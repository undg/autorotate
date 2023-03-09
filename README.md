# autorotate

`autorotate` is Linux/X11 utility for 2in1 laptops and other devices with touchscreen. No ROOT privileges necessary!

## Install
Download and unzip [autorotate.zip](https://github.com/undg/autorotate/releases/latest). Give execution permission with command:
```
unzip autorotate.zip
cd bin/
chmod +x autorotate
```
## Compile from source

To compile binary file in `bin` directory, run the command:
`
go build -o bin/autorotate
cd bin/
chmod +x autorotate
`

## Usage 
Command `left` `right` `portrait` and `lanscape` are wrappers around xrandr and xinput to synchronize rotation of screen and digitizer. In fact they will rotate all xinput devices including stylus, touchpad, touchscreen, mouse and ~~webcam~~.

Laptops with axis sensor can be run without any command to detect orientation. Flag `--daemon` is designed to run in background as a service.

Check display name with `list` command.

`autorotate list`

In following examples `eDP` should be replaced with your display name.
 
Single run and exit:

`autorotate left --display eDP`

`autorotate help`

You keep app running with command that should be added to autostart in your window manager.

`autorotate --display eDP`


![image](https://user-images.githubusercontent.com/5306983/217210748-93221f5f-8dab-4645-84b2-10505d149206.png)

