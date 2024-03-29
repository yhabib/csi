package vm

import (
	"log"
)

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Notes:
//  - Can I add byte to int? 0x08 + 3? Yes, byte is just an int capped at 255
//  - Overflow is down automatically because of the size of the variable
// -  I can't have negative numbers, overflow takes care 2-3=255

// From Oz:
// You could theoretically write a program which produces instructions at run time, places them
// in memory, and then points the program counter at them (this is what JIT compilation is).
// For this trick you need a segment that’s both writable and executable…. if you were using the data
// segment here you’d have to leave it executable; if using the text segment you’d need to make it writable

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {
	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
loop:
	for {
		pc := registers[0]
		op := memory[pc]

		switch op {
		case Load:
			registers[memory[pc+1]] = memory[memory[pc+2]]

		// Memory protection
		case Store:
			mp := memory[pc+2]
			if mp > 7 {
				log.Fatalf("Terminating program due an attempt to overwritte instructions sections")
			}
			memory[mp] = registers[memory[pc+1]]

		case Halt:
			break loop

		case Add:
			registers[memory[pc+1]] += registers[memory[pc+2]]

		case Sub:
			registers[memory[pc+1]] -= registers[memory[pc+2]]

		// Memory protection
		case Jump:
			newpc := memory[pc+1]
			if newpc < 8 {
				log.Fatalf("Terminating program due an attempt to jump out of the instructions section")
			}

			// This is not so straightforward, some instrucions betweeen new pcition and current one could avoid potential infinite loop
			// if newpc < registers[0] {
			// 	log.Fatalf("Terminating program due an attempt to jump back in the intructions section, pcsible infinite loop")
			// }
			registers[0] = newpc
			continue loop

		// Branch Equal to Zero
		case Beqz:
			reg, offset := memory[pc+1], memory[pc+2]
			if registers[reg] == 0 {
				registers[0] += offset
			}

		case Addi:
			registers[memory[pc+1]] += memory[pc+2]

		case Subi:
			registers[memory[pc+1]] -= memory[pc+2]

		default:
			// panic(fmt.Sprintf("unknown operation %d"))
			log.Fatalf("Terminting due unknow operation %d", op)
		}
		registers[0] += 3
	}
}
