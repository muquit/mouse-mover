package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32            = syscall.MustLoadDLL("user32.dll")
	getLastInputInfo  = user32.MustFindProc("GetLastInputInfo")
	getCursorPos      = user32.MustFindProc("GetCursorPos")
	setCursorPos      = user32.MustFindProc("SetCursorPos")
	mouse_event       = user32.MustFindProc("mouse_event")
	kernel32          = syscall.MustLoadDLL("kernel32.dll")
	getTickCount      = kernel32.MustFindProc("GetTickCount")
	getSystemMetrics  = user32.MustFindProc("GetSystemMetrics")
)

type LASTINPUTINFO struct {
	CbSize uint32
	DwTime uint32
}

const (
	VERSION = "1.0.1"
	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004
	SM_CXSCREEN         = 0x0000
	SM_CYSCREEN         = 0x0001
)

type Corner struct {
	x, y int
	name string
}

func getIdleTime() uint32 {
	var lastInputInfo LASTINPUTINFO
	lastInputInfo.CbSize = uint32(unsafe.Sizeof(lastInputInfo))
	
	r, _, _ := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r == 0 {
		return 0
	}
	
	currentTickCount, _, _ := getTickCount.Call()
	return uint32(currentTickCount) - lastInputInfo.DwTime
}

func getScreenDimensions() (width, height int) {
	w, _, _ := getSystemMetrics.Call(SM_CXSCREEN)
	h, _, _ := getSystemMetrics.Call(SM_CYSCREEN)
	return int(w), int(h)
}

func getCornerPosition(corner string) Corner {
	width, height := getScreenDimensions()
	corners := map[string]Corner{
		"ulc": {0, 0, "Upper Left"},                   // Upper Left Corner
		"urc": {width, 0, "Upper Right"},              // Upper Right Corner
		"blc": {0, height, "Bottom Left"},             // Bottom Left Corner
		"brc": {width, height, "Bottom Right"},        // Bottom Right Corner
	}
	
	if pos, exists := corners[strings.ToLower(corner)]; exists {
		return pos
	}
	return corners["ulc"] // Default to upper left if invalid
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Show version info")
	idleTimePtr := flag.Int("idle", 60, "Idle time in seconds before moving mouse")
	cornerPtr := flag.String("corner", "ulc", "Corner to move mouse to (ulc, urc, blc, brc)")
	tickPtr := flag.Int("tick", 0, "Print idle time every X seconds (0 to disable)")
	flag.Parse()

	if (showVersion) {
		fmt.Printf("mouse-mover.exe v%s\n",VERSION)
		os.Exit(0)
	}
	
	corner := getCornerPosition(*cornerPtr)
	fmt.Printf("Mouse will move to %s corner (%d,%d) after %d seconds of idle time\n", 
		corner.name, corner.x, corner.y, *idleTimePtr)
	if *tickPtr > 0 {
		fmt.Printf("Will print idle time every %d seconds\n", *tickPtr)
	}
	
	lastTick := time.Now()
	
	for {
		idleTime := getIdleTime()
		currentTime := time.Now()
		
		// Print idle time if tick is enabled
		if *tickPtr > 0 && currentTime.Sub(lastTick) >= time.Duration(*tickPtr)*time.Second {
			fmt.Printf("Current idle time: %.2f seconds\n", float64(idleTime)/1000.0)
			lastTick = currentTime
		}
		
		if idleTime >= uint32(*idleTimePtr * 1000) { // Convert seconds to milliseconds
			// Move mouse to specified corner
			setCursorPos.Call(uintptr(corner.x), uintptr(corner.y))
			
			// Perform click
			mouse_event.Call(
				MOUSEEVENTF_LEFTDOWN,
				uintptr(corner.x),
				uintptr(corner.y),
				0,
				0,
			)
			time.Sleep(time.Millisecond * 50)
			mouse_event.Call(
				MOUSEEVENTF_LEFTUP,
				uintptr(corner.x),
				uintptr(corner.y),
				0,
				0,
			)
			
			fmt.Printf("Moved mouse to %s corner and clicked at %v\n", 
				corner.name, time.Now().Format("15:04:05"))
			
			// Wait a bit before checking again to avoid continuous clicking
			time.Sleep(time.Second * 2)
		}
		
		time.Sleep(time.Second) // Check every second
	}
}
