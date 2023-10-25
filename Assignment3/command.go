package main

import "fmt"

type Command interface {
    execute()
}

type Device interface {
    on()
    off()
}

type Button struct {
    command Command
}

func (b *Button) press() {
    b.command.execute()
}

type OnCommand struct {
    device Device
}

func (c *OnCommand) execute() {
    c.device.on()
    fmt.Println("Turning on TV")
}

type OffCommand struct {
    device Device
}

func (c *OffCommand) execute() {
    c.device.off()
    fmt.Println("Turning off TV")
}

type Tv struct {
    isRunning bool
}

func (t *Tv) on() {
    t.isRunning = true
    fmt.Println("TV is turned on")
}

func (t *Tv) off() {
    t.isRunning = false
    fmt.Println("TV is turned off")
}
func main() {
    tv := &Tv{}
    onCommand := &OnCommand{
        device: tv,
    }
    offCommand := &OffCommand{
        device: tv,
    }
    onButton := &Button{
        command: onCommand,
    }
    offButton := &Button{
        command: offCommand,
    }
    fmt.Println("Pressing the on button...")
    onButton.press()
    fmt.Println("Pressing the off button...")
    offButton.press()
}