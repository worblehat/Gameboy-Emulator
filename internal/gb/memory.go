package gb

import (
	"fmt"
)

const BootROMSize = 256
const CartROM0Size = 0x4000
const vramSize = 0x2000
const hramSize = 0x7F
const ioMemSize = 0x80

type Memory struct {
	bootROM  [BootROMSize]byte
	cartROM0 [CartROM0Size]byte
	vram     [vramSize]byte
	hram     [hramSize]byte
	ioMem    [ioMemSize]byte
}

func NewMemory(bootROM [BootROMSize]byte, cardROM0 [CartROM0Size]byte) *Memory {
	return &Memory{
		bootROM:  bootROM,
		cartROM0: cardROM0,
	}
}

func (m *Memory) Read8(addr uint16) uint8 {
	if addr >= 0x0000 && addr < 0x0100 {
		return m.bootROM[addr]
	} else if addr >= 0x0100 && addr < 0x4000 {
		return m.cartROM0[addr-0x0100]
	} else if addr >= 0x8000 && addr < 0xA000 {
		return m.vram[addr-0x8000]
	} else if addr >= 0xFF00 && addr < 0xFF80 {
		value := m.ioMem[addr-0xFF00]
		fmt.Printf("Reading from I/O Memory at 0x%04X: 0x%02X.\n", addr, value)
		return value
	} else if addr >= 0xFF80 && addr < 0xFFFF {
		return m.hram[addr-0xFF80]
	}
	panic(fmt.Sprintf("Read from unknown memory address 0x%X", addr))
}

func (m *Memory) Read16(addr uint16) uint16 {
	// Little endian
	loByte := uint16(m.Read8(addr))
	hiByte := uint16(m.Read8(addr + 1))
	return (hiByte << 8) | loByte
}

func (m *Memory) Write8(addr uint16, val uint8) {
	if addr >= 0x8000 && addr < 0xA000 {
		m.vram[addr-0x8000] = val
	} else if addr >= 0xFF00 && addr < 0xFF80 {
		fmt.Printf("Writing to I/O Memory at 0x%04X: 0x%02X.\n", addr, val)
		m.ioMem[addr-0xFF00] = val
	} else if addr >= 0xFF80 && addr < 0xFFFF {
		m.hram[addr-0xFF80] = val
	} else {
		panic(fmt.Sprintf("Write to non-writable memory address 0x%X", addr))
	}
}

func (m *Memory) Write16(addr uint16, val uint16) {
	// Little endian
	loByte := uint8(val)
	hiByte := uint8(val >> 8)
	m.Write8(addr, loByte)
	m.Write8(addr+1, hiByte)
}
