package gb

const opCodeExt uint16 = 0xCB

var instruction = map[uint16]Instruction{
	0x01:   {"LD BC,nn", LD_BC_nn},
	0x02:   {"LD (BC),A", LD_pBC_A},
	0x03:   {"INC BC", INC_BC},
	0x04:   {"INC B", INC_B},
	0x05:   {"DEC B", DEC_B},
	0x06:   {"LD B,n", LD_B_n},
	0x07:   {"RLCA", RLCA},
	0x0C:   {"INC C", INC_C},
	0x0D:   {"DEC C", DEC_C},
	0x0A:   {"LD A,(BC)", LD_A_pBC},
	0x0B:   {"DEC BC", DEC_BC},
	0x0E:   {"LD C,n", LD_C_n},
	0x0F:   {"RRCA", RRCA},
	0x11:   {"LD DE,nn", LD_DE_nn},
	0x12:   {"LD (DE),A", LD_pDE_A},
	0x13:   {"INC DE", INC_DE},
	0x14:   {"INC D", INC_D},
	0x15:   {"DEC D", DEC_D},
	0x16:   {"LD D,n", LD_D_n},
	0x17:   {"RLA", RLA},
	0x1A:   {"LD A,(DE)", LD_A_pDE},
	0x1B:   {"DEC DE", DEC_DE},
	0x1C:   {"INC E", INC_E},
	0x1D:   {"DEC E", DEC_E},
	0x1E:   {"LD E,n", LD_E_n},
	0x1F:   {"RRA", RRA},
	0x20:   {"JR NZ,n", JR_NZ_n},
	0x21:   {"LD HL,nn", LD_HL_nn},
	0x22:   {"LD (HL+),A", LDI_pHL_A},
	0x23:   {"INC HL", INC_HL},
	0x24:   {"INC H", INC_H},
	0x25:   {"DEC H", DEC_H},
	0x26:   {"LD H,n", LD_H_n},
	0x28:   {"JR Z,n", JR_Z_n},
	0x2A:   {"LD A,(HL+)", LDI_A_pHL},
	0x2B:   {"DEC HL", DEC_HL},
	0x2C:   {"INC L", INC_L},
	0x2D:   {"DEC L", DEC_L},
	0x2E:   {"LD L,n", LD_L_n},
	0x30:   {"JR NC,n", JR_NC_n},
	0x31:   {"LD SP,nn", LD_SP_nn},
	0x32:   {"LD (HL-),A", LDD_pHL_A},
	0x33:   {"INC SP", INC_SP},
	0x34:   {"INC (HL)", INC_pHL},
	0x35:   {"DEC (HL)", DEC_pHL},
	0x38:   {"JR C,n", JR_C_n},
	0x3A:   {"LD A,(HL-)", LDD_A_pHL},
	0x3B:   {"DEC SP", DEC_SP},
	0x3C:   {"INC A", INC_A},
	0x3D:   {"DEC A", DEC_A},
	0x3E:   {"LD A,n", LD_A_n},
	0x47:   {"LD B,A", LD_B_A},
	0x4F:   {"LD C,A", LD_C_A},
	0x57:   {"LD D,A", LD_D_A},
	0x5F:   {"LD E,A", LD_E_A},
	0x67:   {"LD H,A", LD_H_A},
	0x6F:   {"LD L,A", LD_L_A},
	0x77:   {"LD (HL),A", LD_pHL_A},
	0x78:   {"LD A,B", LD_A_B},
	0x79:   {"LD A,C", LD_A_C},
	0x7A:   {"LD A,D", LD_A_D},
	0x7B:   {"LD A,E", LD_A_E},
	0x7C:   {"LD A,H", LD_A_H},
	0x7D:   {"LD A,L", LD_A_L},
	0x7E:   {"LD A,(HL)", LD_A_pHL},
	0x7F:   {"LD A,A", LD_A_A},
	0xAF:   {"XOR A", XOR_A_A},
	0xA8:   {"XOR B", XOR_A_B},
	0xA9:   {"XOR C", XOR_A_C},
	0xAA:   {"XOR D", XOR_A_D},
	0xAB:   {"XOR E", XOR_A_E},
	0xAC:   {"XOR H", XOR_A_H},
	0xAD:   {"XOR L", XOR_A_L},
	0xAE:   {"XOR (HL)", XOR_A_pHL},
	0xB8:   {"CP B", CP_B},
	0xB9:   {"CP C", CP_C},
	0xBA:   {"CP D", CP_D},
	0xBB:   {"CP E", CP_E},
	0xBC:   {"CP H", CP_H},
	0xBD:   {"CP L", CP_L},
	0xBE:   {"CP (HL)", CP_pHL},
	0xFE:   {"CP n", CP_n},
	0xBF:   {"CP A", CP_A},
	0xC0:   {"RET NZ", RET_NZ},
	0xC1:   {"POP BC", POP_BC},
	0xC4:   {"CALL NZ,nn", CALL_NZ_nn},
	0xC5:   {"PUSH BC", PUSH_BC},
	0xC8:   {"RET Z", RET_Z},
	0xC9:   {"RET", RET},
	0xCC:   {"CALL Z,nn", CALL_Z_nn},
	0xCD:   {"CALL nn", CALL_nn},
	0xD0:   {"RET NC", RET_NC},
	0xD1:   {"POP DE", POP_DE},
	0xD4:   {"CALL NC,nn", CALL_NC_nn},
	0xD5:   {"PUSH DE", PUSH_DE},
	0xD8:   {"RET C", RET_C},
	0xDC:   {"CALL C,nn", CALL_C_nn},
	0xEA:   {"LD (nn),A", LD_pnn_A},
	0xE0:   {"LD ($FF00+n),A", LD_IO_n_A},
	0xE1:   {"POP HL", POP_HL},
	0xE2:   {"LD ($FF00+C),A", LD_IO_C_A},
	0xE5:   {"PUSH HL", PUSH_HL},
	0xEE:   {"XOR n", XOR_A_n},
	0xF0:   {"LD A,($FF00+n)", LD_A_IO_n},
	0xF1:   {"POP AF", POP_AF},
	0xFA:   {"LD A,(nn)", LD_A_pnn},
	0xF2:   {"LD A,($FF00+C)", LD_A_IO_C},
	0xF5:   {"PUSH AF", PUSH_AF},
	0xCB00: {"RLC B", RLC_B},
	0xCB01: {"RLC C", RLC_C},
	0xCB02: {"RLC D", RLC_D},
	0xCB03: {"RLC E", RLC_E},
	0xCB04: {"RLC H", RLC_H},
	0xCB05: {"RLC L", RLC_L},
	0xCB06: {"RLC (HL)", RLC_pHL},
	0xCB07: {"RLC A", RLC_A},
	0xCB08: {"RRC B", RRC_B},
	0xCB09: {"RRC C", RRC_C},
	0xCB0A: {"RRC D", RRC_D},
	0xCB0B: {"RRC E", RRC_E},
	0xCB0C: {"RRC H", RRC_H},
	0xCB0D: {"RRC L", RRC_L},
	0xCB0E: {"RRC (HL)", RRC_pHL},
	0xCB0F: {"RRC A", RRC_A},
	0xCB10: {"RL B", RL_B},
	0xCB11: {"RL C", RL_C},
	0xCB12: {"RL D", RL_D},
	0xCB13: {"RL E", RL_E},
	0xCB14: {"RL H", RL_H},
	0xCB15: {"RL L", RL_L},
	0xCB16: {"RL (HL)", RL_pHL},
	0xCB17: {"RL A", RL_A},
	0xCB18: {"RR B", RR_B},
	0xCB19: {"RR C", RR_C},
	0xCB1A: {"RR D", RR_D},
	0xCB1B: {"RR E", RR_E},
	0xCB1C: {"RR H", RR_H},
	0xCB1D: {"RR L", RR_L},
	0xCB1E: {"RR (HL)", RR_pHL},
	0xCB1F: {"RR A", RR_A},
	0xCB40: {"BIT 0,B", BIT_0_B},
	0xCB41: {"BIT 0,C", BIT_0_C},
	0xCB42: {"BIT 0,D", BIT_0_D},
	0xCB43: {"BIT 0,E", BIT_0_E},
	0xCB44: {"BIT 0,H", BIT_0_H},
	0xCB45: {"BIT 0,L", BIT_0_L},
	0xCB46: {"BIT 0,(HL)", BIT_0_pHL},
	0xCB47: {"BIT_0,A", BIT_0_A},
	0xCB48: {"BIT_1,B", BIT_1_B},
	0xCB49: {"BIT_1,C", BIT_1_C},
	0xCB4A: {"BIT_1,D", BIT_1_D},
	0xCB4B: {"BIT_1,E", BIT_1_E},
	0xCB4C: {"BIT_1,H", BIT_1_H},
	0xCB4D: {"BIT_1,L", BIT_1_L},
	0xCB4E: {"BIT_1,(HL)", BIT_1_pHL},
	0xCB4F: {"BIT_1,A", BIT_1_A},
	0xCB50: {"BIT_2,B", BIT_2_B},
	0xCB51: {"BIT_2,C", BIT_2_C},
	0xCB52: {"BIT_2,D", BIT_2_D},
	0xCB53: {"BIT_2,E", BIT_2_E},
	0xCB54: {"BIT_2,H", BIT_2_H},
	0xCB55: {"BIT_2,L", BIT_2_L},
	0xCB56: {"BIT_2,(HL)", BIT_2_pHL},
	0xCB57: {"BIT_2,A", BIT_2_A},
	0xCB58: {"BIT_3,B", BIT_3_B},
	0xCB59: {"BIT_3,C", BIT_3_C},
	0xCB5A: {"BIT_3,D", BIT_3_D},
	0xCB5B: {"BIT_3,E", BIT_3_E},
	0xCB5C: {"BIT_3,H", BIT_3_H},
	0xCB5D: {"BIT_3,L", BIT_3_L},
	0xCB5E: {"BIT_3,(HL)", BIT_3_pHL},
	0xCB5F: {"BIT_3,A", BIT_3_A},
	0xCB60: {"BIT_4,B", BIT_4_B},
	0xCB61: {"BIT_4,C", BIT_4_C},
	0xCB62: {"BIT_4,D", BIT_4_D},
	0xCB63: {"BIT_4,E", BIT_4_E},
	0xCB64: {"BIT_4,H", BIT_4_H},
	0xCB65: {"BIT_4,L", BIT_4_L},
	0xCB66: {"BIT_4,(HL)", BIT_4_pHL},
	0xCB67: {"BIT_4,A", BIT_4_A},
	0xCB68: {"BIT_5,B", BIT_5_B},
	0xCB69: {"BIT_5,C", BIT_5_C},
	0xCB6A: {"BIT_5,D", BIT_5_D},
	0xCB6B: {"BIT_5,E", BIT_5_E},
	0xCB6C: {"BIT_5,H", BIT_5_H},
	0xCB6D: {"BIT_5,L", BIT_5_L},
	0xCB6E: {"BIT_5,(HL)", BIT_5_pHL},
	0xCB6F: {"BIT_5,A", BIT_5_A},
	0xCB70: {"BIT_6,B", BIT_6_B},
	0xCB71: {"BIT_6,C", BIT_6_C},
	0xCB72: {"BIT_6,D", BIT_6_D},
	0xCB73: {"BIT_6,E", BIT_6_E},
	0xCB74: {"BIT_6,H", BIT_6_H},
	0xCB75: {"BIT_6,L", BIT_6_L},
	0xCB76: {"BIT_6,(HL)", BIT_6_pHL},
	0xCB77: {"BIT_6,A", BIT_6_A},
	0xCB78: {"BIT_7,B", BIT_7_B},
	0xCB79: {"BIT_7,C", BIT_7_C},
	0xCB7A: {"BIT_7,D", BIT_7_D},
	0xCB7B: {"BIT_7,E", BIT_7_E},
	0xCB7C: {"BIT_7,H", BIT_7_H},
	0xCB7D: {"BIT_7,L", BIT_7_L},
	0xCB7E: {"BIT_7,(HL)", BIT_7_pHL},
	0xCB7F: {"BIT_7,A", BIT_7_A},
}
