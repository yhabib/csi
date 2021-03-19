package vm

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
//  - Can I add byte to int? 0x08 + 3?
//  - Overflow is down automatically because of the size of the variable
// -  I can't have negative numbers, same overflow 2-3=255

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
		pos := registers[0]
		op := memory[pos]

		switch op {
		case Load:
			registers[memory[pos+1]] = memory[memory[pos+2]]

		case Store:
			memory[memory[pos+2]] = registers[memory[pos+1]]

		case Halt:
			break loop

		case Add:
			registers[memory[pos+1]] += registers[memory[pos+2]]

		case Sub:
			registers[memory[pos+1]] -= registers[memory[pos+2]]

		case Jump:
			registers[0] = memory[pos+1]
			continue loop

		case Beqz:
			reg, offset := memory[pos+1], memory[pos+2]
			if registers[reg] == 0 {
				registers[0] += offset
			}

		case Addi:
			registers[memory[pos+1]] += memory[pos+2]

		case Subi:
			registers[memory[pos+1]] -= memory[pos+2]

		}
		registers[0] += 3
	}
}
