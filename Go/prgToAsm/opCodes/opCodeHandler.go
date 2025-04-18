package opCodeHandler

var opCodesMap = map[byte]OpCodeInfo{

	0x00: {'brk', 1, '', ''},
	0x01: {'ora', 2, '($', ',x)'},
	0x02: {'kil', 1, '', ''},
	0x03: {'slo', 2, '($', ',x)'},
	0x04: {'nop', 2, '$', ''},
	0x05: {'ora', 2, '$', ''},
	0x06: {'asl', 2, '$', ''},
	0x07: {'slo', 2, '$', ''},
	0x08: {'php', 1, '', ''},
	0x09: {'ora', 2, '#$', ''},
	0x0a: {'asl', 1, '', ''},
	0x0b: {'anc', 2, '#$', ''},
	0x0c: {'nop', 3, '', ''},
	0x0d: {'ora', 3, '', ''},
	0x0e: {'asl', 3, '', ''},
	0x0f: {'slo', 3, '', ''},
	0x10: {'bpl', 2, '*', ''},
	0x11: {'ora', 2, '($', '),y'},
	0x12: {'kil', 1, '', ''},
	0x13: {'slo', 2, '($', '),y'},
	0x14: {'nop', 2, '$', ',x'},
	0x15: {'ora', 2, '$', ',x'},
	0x16: {'asl', 2, '$', ',x'},
	0x17: {'slo', 2, '$', ',x'},
	0x18: {'clc', 1, '', ''},
	0x19: {'ora', 3, '', ',y'},
	0x1a: {'nop', 1, '', ''},
	0x1b: {'slo', 3, '', ',y'},
	0x1c: {'nop', 3, '', ',x'},
	0x1d: {'ora', 3, '', ',x'},
	0x1e: {'asl', 3, '', ',x'},
	0x1f: {'slo', 3, '', ',x'},
	0x20: {'jsr', 3, '', ''},
	0x21: {'and', 2, '($', ',x)'},
	0x22: {'kil', 1, '', ''},
	0x23: {'rla', 2, '($', ',x)'},
	0x24: {'bit', 2, '$', ''},
	0x25: {'and', 2, '$', ''},
	0x26: {'rol', 2, '$', ''},
	0x27: {'rla', 2, '$', ''},
	0x28: {'plp', 1, '', ''},
	0x29: {'and', 2, '#$', ''},
	0x2a: {'rol', 1, '', ''},
	0x2b: {'anc', 2, '#$', ''},
	0x2c: {'bit', 3, '', ''},
	0x2d: {'and', 3, '', ''},
	0x2e: {'rol', 3, '', ''},
	0x2f: {'rla', 3, '', ''},
	0x30: {'bmi', 2, '*', ''},
	0x31: {'and', 2, '($', '),y'},
	0x32: {'kil', 1, '', ''},
	0x33: {'rla', 2, '($', '),y'},
	0x34: {'nop', 2, '$', ',x'},
	0x35: {'and', 2, '$', ',x'},
	0x36: {'rol', 2, '$', ',x'},
	0x37: {'rla', 2, '$', ',x'},
	0x38: {'sec', 1, '', ''},
	0x39: {'and', 3, '', ',y'},
	0x3a: {'nop', 1, '', ''},
	0x3b: {'rla', 3, '', ',y'},
	0x3c: {'nop', 3, '', ',x'},
	0x3d: {'and', 3, '', ',x'},
	0x3e: {'rol', 3, '', ',x'},
	0x3f: {'rla', 3, '', ',x'},
	0x40: {'rti', 1, '', ''},
	0x41: {'eor', 2, '($', ',x)'},
	0x42: {'kil', 1, '', ''},
	0x43: {'sre', 2, '($', ',x)'},
	0x44: {'nop', 2, '$', ''},
	0x45: {'eor', 2, '$', ''},
	0x46: {'lsr', 2, '$', ''},
	0x47: {'sre', 2, '$', ''},
	0x48: {'pha', 1, '', ''},
	0x49: {'eor', 2, '#$', ''},
	0x4a: {'lsr', 1, '', ''},
	0x4b: {'alr', 2, '#$', ''},
	0x4c: {'jmp', 3, '', ''},
	0x4d: {'eor', 3, '', ''},
	0x4e: {'lsr', 3, '', ''},
	0x4f: {'sre', 3, '', ''},
	0x50: {'bvc', 2, '*', ''},
	0x51: {'eor', 2, '($', '),y'},
	0x52: {'kil', 1, '', ''},
	0x53: {'sre', 2, '($', '),y'},
	0x54: {'nop', 2, '$', ',x'},
	0x55: {'eor', 2, '$', ',x'},
	0x56: {'lsr', 2, '$', ',x'},
	0x57: {'sre', 2, '$', ',x'},
	0x58: {'cli', 1, '', ''},
	0x59: {'eor', 3, '', ',y'},
	0x5a: {'nop', 1, '', ''},
	0x5b: {'sre', 3, '', ',y'},
	0x5c: {'nop', 3, '', ',x'},
	0x5d: {'eor', 3, '', ',x'},
	0x5e: {'lsr', 3, '', ',x'},
	0x5f: {'sre', 3, '', ',x'},
	0x60: {'rts', 1, '', ''},
	0x61: {'adc', 2, '($', ',x)'},
	0x62: {'kil', 1, '', ''},
	0x63: {'rra', 2, '($', ',x)'},
	0x64: {'nop', 2, '$', ''},
	0x65: {'adc', 2, '$', ''},
	0x66: {'ror', 2, '$', ''},
	0x67: {'rra', 2, '$', ''},
	0x68: {'pla', 1, '', ''},
	0x69: {'adc', 2, '#$', ''},
	0x6a: {'ror', 1, '', ''},
	0x6b: {'arr', 2, '#$', ''},
	0x6c: {'jmp', 3, '($', ')'},
	0x6d: {'adc', 3, '', ''},
	0x6e: {'ror', 3, '', ''},
	0x6f: {'rra', 3, '', ''},
	0x70: {'bvs', 2, '*', ''},
	0x71: {'adc', 2, '($', '),y'},
	0x72: {'kil', 1, '', ''},
	0x73: {'rra', 2, '($', '),y'},
	0x74: {'nop', 2, '$', ',x'},
	0x75: {'adc', 2, '$', ',x'},
	0x76: {'ror', 2, '$', ',x'},
	0x77: {'rra', 2, '$', ',x'},
	0x78: {'sei', 1, '', ''},
	0x79: {'adc', 3, '', ',y'},
	0x7a: {'nop', 1, '', ''},
	0x7b: {'rra', 3, '', ',y'},
	0x7c: {'nop', 3, '', ',x'},
	0x7d: {'adc', 3, '', ',x'},
	0x7e: {'ror', 3, '', ',x'},
	0x7f: {'rra', 3, '', ',x'},
	0x80: {'nop', 2, '#$', ''},
	0x81: {'sta', 2, '($', ',x)'},
	0x82: {'nop', 2, '#$', ''},
	0x83: {'sax', 2, '($', ',x)'},
	0x84: {'sty', 2, '$', ''},
	0x85: {'sta', 2, '$', ''},
	0x86: {'stx', 2, '$', ''},
	0x87: {'sax', 2, '$', ''},
	0x88: {'dey', 1, '', ''},
	0x89: {'nop', 2, '#$', ''},
	0x8a: {'txa', 1, '', ''},
	0x8b: {'xaa', 2, '#$', ''},
	0x8c: {'sty', 3, '', ''},
	0x8d: {'sta', 3, '', ''},
	0x8e: {'stx', 3, '', ''},
	0x8f: {'sax', 3, '', ''},
	0x90: {'bcc', 2, '*', ''},
	0x91: {'sta', 2, '($', '),y'},
	0x92: {'kil', 1, '', ''},
	0x93: {'ahx', 2, '($', '),y'},
	0x94: {'sty', 2, '$', ',x'},
	0x95: {'sta', 2, '$', ',x'},
	0x96: {'stx', 2, '$', ',y'},
	0x97: {'sax', 2, '$', ',y'},
	0x98: {'tya', 1, '', ''},
	0x99: {'sta', 3, '', ',y'},
	0x9a: {'txs', 1, '', ''},
	0x9b: {'tas', 3, '', ',y'},
	0x9c: {'shy', 3, '', ',x'},
	0x9d: {'sta', 3, '', ',x'},
	0x9e: {'shx', 3, '', ',y'},
	0x9f: {'sha', 3, '', ',y'},
	0xa0: {'ldy', 2, '#$', ''},
	0xa1: {'lda', 2, '($', ',x)'},
	0xa2: {'ldx', 2, '#$', ''},
	0xa3: {'lax', 2, '($', ',x)'},
	0xa4: {'ldy', 2, '$', ''},
	0xa5: {'lda', 2, '$', ''},
	0xa6: {'ldx', 2, '$', ''},
	0xa7: {'lax', 2, '$', ''},
	0xa8: {'tay', 1, '', ''},
	0xa9: {'lda', 2, '#$', ''},
	0xaa: {'tax', 1, '', ''},
	0xab: {'lax', 2, '#$', ''},
	0xac: {'ldy', 3, '', ''},
	0xad: {'lda', 3, '', ''},
	0xae: {'ldx', 3, '', ''},
	0xaf: {'lax', 3, '', ''},
	0xb0: {'bcs', 2, '*', ''},
	0xb1: {'lda', 2, '($', '),y'},
	0xb2: {'kil', 1, '', ''},
	0xb3: {'lax', 2, '($', '),y'},
	0xb4: {'ldy', 2, '$', ',x'},
	0xb5: {'lda', 2, '$', ',x'},
	0xb6: {'ldx', 2, '$', ',y'},
	0xb7: {'lax', 2, '$', ',y'},
	0xb8: {'clv', 1, '', ''},
	0xb9: {'lda', 3, '', ',y'},
	0xba: {'tsx', 1, '', ''},
	0xbb: {'las', 3, '', ',y'},
	0xbc: {'ldy', 3, '', ',x'},
	0xbd: {'lda', 3, '', ',x'},
	0xbe: {'ldx', 3, '', ',y'},
	0xbf: {'lax', 3, '', ',y'},
	0xc0: {'cpy', 2, '#$', ''},
	0xc1: {'cmp', 2, '($', ',x)'},
	0xc2: {'nop', 2, '#$', ''},
	0xc3: {'dcp', 2, '$', ''},
	0xc4: {'cpy', 2, '$', ''},
	0xc5: {'cmp', 2, '$', ''},
	0xc6: {'dec', 2, '$', ''},
	0xc7: {'dcp', 2, '$', ''},
	0xc8: {'iny', 1, '', ''},
	0xc9: {'cmp', 2, '#$', ''},
	0xca: {'dex', 1, '', ''},
	0xcb: {'axs', 2, '#$', ''},
	0xcc: {'cpy', 3, '', ''},
	0xcd: {'cmp', 3, '', ''},
	0xce: {'dec', 3, '', ''},
	0xcf: {'dcp', 3, '', ''},
	0xd0: {'bne', 2, '*', ''},
	0xd1: {'cmp', 2, '($', '),y'},
	0xd2: {'kil', 1, '', ''},
	0xd3: {'dcp', 2, '($', '),y'},
	0xd4: {'nop', 2, '$', ',x'},
	0xd5: {'cmp', 2, '$', ',x'},
	0xd6: {'dec', 2, '$', ',x'},
	0xd7: {'dcp', 2, '$', ',x'},
	0xd8: {'cld', 1, '', ''},
	0xd9: {'cmp', 3, '', ',y'},
	0xda: {'nop', 1, '', ''},
	0xdb: {'dcp', 3, '', ',x'},
	0xdc: {'nop', 3, '', ',x'},
	0xdd: {'cmp', 3, '', ',x'},
	0xde: {'dec', 3, '', ',x'},
	0xdf: {'dcp', 3, '', ',x'},
	0xe0: {'cpx', 2, '#$', ''},
	0xe1: {'sbc', 2, '($', ',x)'},
	0xe2: {'nop', 2, '#$', ''},
	0xe3: {'isc', 2, '($', ',x)'},
	0xe4: {'cpx', 2, '$', ''},
	0xe5: {'sbc', 2, '$', ''},
	0xe6: {'inc', 2, '$', ''},
	0xe7: {'isc', 2, '$', ''},
	0xe8: {'inx', 1, '', ''},
	0xe9: {'sbc', 2, '#$', ''},
	0xea: {'nop', 1, '', ''},
	0xeb: {'sbc', 2, '#$', ''},
	0xec: {'cpx', 3, '', ''},
	0xed: {'sbc', 3, '', ''},
	0xee: {'inc', 3, '', ''},
	0xef: {'isc', 3, '', ''},
	0xf0: {'beq', 2, '*', ''},
	0xf1: {'sbc', 2, '($', '),y'},
	0xf2: {'kil', 1, '', ''},
	0xf3: {'isc', 2, '($', '),y'},
	0xf4: {'nop', 2, '$', ',x'},
	0xf5: {'sbc', 2, '$', ',x'},
	0xf6: {'inc', 2, '$', ',x'},
	0xf7: {'isc', 2, '$', ',x'},
	0xf8: {'sed', 1, '', ''},
	0xf9: {'sbc', 3, '', ',y'},
	0xfa: {'nop', 1, '', ''},
	0xfb: {'isc', 3, '', ',y'},
	0xfc: {'nop', 3, '', ',x'},
	0xfd: {'sbc', 3, '', ',x'},
	0xfe: {'inc', 3, '', ',x'},
	0xff: {'isc', 3, '', ',x'}
}

