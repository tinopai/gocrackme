package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/denisbrodbeck/machineid"
)

func MessageBox(hh uintptr, ca, tt string, ff uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hh),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ca))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tt))),
		uintptr(ff))

	return int(ret)
}

func main() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}

	hashed := sha256.Sum256([]byte(id[:6]))
	fmt.Printf("Your password: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	if sha256.Sum256([]byte(input.Text())) == hashed {
		MessageBox(0, "Congratulations! You solved it", "MachineID CrackMe", 0)
	} else {
		MessageBox(0, "Wrong password!", "MachineID CrackMe", 0)
	}
}
