//    - NOTE: 2^9 = 28^8 * 2
//    - polynomial long division
//    - XOR (^) if decimal is >= 256 then XOR with 285 -> recursive
//    - When adding exponents, if the exponent becomes greater than or equal to 256, simply apply modulo 255
//       - 2^170 * 2^164 = 2^(170+164) = 2^334 -> 2^334%255 = 2^79
//    - https://www.thonky.com/qr-code-tutorial/log-antilog-table
//
//
// Generator Polynomial for 2 Error Correction Codewords!
//
//    - TODO:
//       - create message polynomial: <NUMBER>^len-(index+1) + <NUMBER>^len-(index+1), if ^n == 0 -> return just number
//       - generator polynomial: x^10 + α^251x^9 + α^67x^8 + α^46x^7 + α^61x^6 + α^118x^5 + α^70x^4 + α^64x^3 + α^94x^2 + α^32x + α^45
//
//       - division:
//          1) multiply message x^10 -> <NUMBER>x^(exp + 10)
//          2) multiply generator by x^15, to have same exponentional number as message, if one "component" do not have a^n, add a^0
//          3) now you can divide
//       - Multiply generator, repeat 16x
//          1) multiply generator by index[0] from message, like its 32x^25, 32 in table == 5 so it's a^5
//          2) if exponent >255, do % 255
//          3) now convert it to numbers, where a == 2, so output will be 32x^25 + 2x^24 + 101x^23 + 10x^22 + 97x^21 + 197x^20 + 15x^19 + 47x^18 + 134x^17 + 74x^16 + 5x^15
//          4) do XOR with message a generator, (NUM_MESSAGE XOR NUM_GENERATOR)x^GENRATOR_EXP
//          4) NOTE: if something after it is 0x^n, remove it

//          5) If it's 16x done (16 = len message), remove exponents and X == result

package encoder

const (
   CW = 10
   ALPHA = 2
)

type Poly struct {
   x          int
   power     int
}
type GeneratorPoly struct {
   alpha      int
   alphaPower int
   x          int
   xPower     int
}

func powerToInt(n int) int {
   switch n {
   case 0:
      return 1
   case 1:
      return 2
   case 2:
      return 4
   case 3:
      return 8
   case 4:
      return 16
   case 5:
      return 32
   case 6:
      return 64
   case 7:
      return 128
   case 8:
      return 29
   case 9:
      return 58
   case 10:
      return 116
   case 11:
      return 232
   case 12:
      return 205
   case 13:
      return 135
   case 14:
      return 19
   case 15:
      return 38
   case 16:
      return 76
   case 17:
      return 152
   case 18:
      return 45
   case 19:
      return 90
   case 20:
      return 180
   case 21:
      return 117
   case 22:
      return 234
   case 23:
      return 201
   case 24:
      return 143
   case 25:
      return 3
   case 26:
      return 6
   case 27:
      return 12
   case 28:
      return 24
   case 29:
      return 48
   case 30:
      return 96
   case 31:
      return 192
   case 32:
      return 157
   case 33:
      return 39
   case 34:
      return 78
   case 35:
      return 156
   case 36:
      return 37
   case 37:
      return 74
   case 38:
      return 148
   case 39:
      return 53
   case 40:
      return 106
   case 41:
      return 212
   case 42:
      return 181
   case 43:
      return 119
   case 44:
      return 238
   case 45:
      return 193
   case 46:
      return 159
   case 47:
      return 35
   case 48:
      return 70
   case 49:
      return 140
   case 50:
      return 5
   case 51:
      return 10
   case 52:
      return 20
   case 53:
      return 40
   case 54:
      return 80
   case 55:
      return 160
   case 56:
      return 93
   case 57:
      return 186
   case 58:
      return 105
   case 59:
      return 210
   case 60:
      return 185
   case 61:
      return 111
   case 62:
      return 222
   case 63:
      return 161
   case 64:
      return 95
   case 65:
      return 190
   case 66:
      return 97
   case 67:
      return 194
   case 68:
      return 153
   case 69:
      return 47
   case 70:
      return 94
   case 71:
      return 188
   case 72:
      return 101
   case 73:
      return 202
   case 74:
      return 137
   case 75:
      return 15
   case 76:
      return 30
   case 77:
      return 60
   case 78:
      return 120
   case 79:
      return 240
   case 80:
      return 253
   case 81:
      return 231
   case 82:
      return 211
   case 83:
      return 187
   case 84:
      return 107
   case 85:
      return 214
   case 86:
      return 177
   case 87:
      return 127
   case 88:
      return 254
   case 89:
      return 225
   case 90:
      return 223
   case 91:
      return 163
   case 92:
      return 91
   case 93:
      return 182
   case 94:
      return 113
   case 95:
      return 226
   case 96:
      return 217
   case 97:
      return 175
   case 98:
      return 67
   case 99:
      return 134
   case 100:
      return 17
   case 101:
      return 34
   case 102:
      return 68
   case 103:
      return 136
   case 104:
      return 13
   case 105:
      return 26
   case 106:
      return 52
   case 107:
      return 104
   case 108:
      return 208
   case 109:
      return 189
   case 110:
      return 103
   case 111:
      return 206
   case 112:
      return 129
   case 113:
      return 31
   case 114:
      return 62
   case 115:
      return 124
   case 116:
      return 248
   case 117:
      return 237
   case 118:
      return 199
   case 119:
      return 147
   case 120:
      return 59
   case 121:
      return 118
   case 122:
      return 236
   case 123:
      return 197
   case 124:
      return 151
   case 125:
      return 51
   case 126:
      return 102
   case 127:
      return 204
   case 128:
      return 133
   case 129:
      return 23
   case 130:
      return 46
   case 131:
      return 92
   case 132:
      return 184
   case 133:
      return 109
   case 134:
      return 218
   case 135:
      return 169
   case 136:
      return 79
   case 137:
      return 158
   case 138:
      return 33
   case 139:
      return 66
   case 140:
      return 132
   case 141:
      return 21
   case 142:
      return 42
   case 143:
      return 84
   case 144:
      return 168
   case 145:
      return 77
   case 146:
      return 154
   case 147:
      return 41
   case 148:
      return 82
   case 149:
      return 164
   case 150:
      return 85
   case 151:
      return 170
   case 152:
      return 73
   case 153:
      return 146
   case 154:
      return 57
   case 155:
      return 114
   case 156:
      return 228
   case 157:
      return 213
   case 158:
      return 183
   case 159:
      return 115
   case 160:
      return 230
   case 161:
      return 209
   case 162:
      return 191
   case 163:
      return 99
   case 164:
      return 198
   case 165:
      return 145
   case 166:
      return 63
   case 167:
      return 126
   case 168:
      return 252
   case 169:
      return 229
   case 170:
      return 215
   case 171:
      return 179
   case 172:
      return 123
   case 173:
      return 246
   case 174:
      return 241
   case 175:
      return 255
   case 176:
      return 227
   case 177:
      return 219
   case 178:
      return 171
   case 179:
      return 75
   case 180:
      return 150
   case 181:
      return 49
   case 182:
      return 98
   case 183:
      return 196
   case 184:
      return 149
   case 185:
      return 55
   case 186:
      return 110
   case 187:
      return 220
   case 188:
      return 165
   case 189:
      return 87
   case 190:
      return 174
   case 191:
      return 65
   case 192:
      return 130
   case 193:
      return 25
   case 194:
      return 50
   case 195:
      return 100
   case 196:
      return 200
   case 197:
      return 141
   case 198:
      return 7
   case 199:
      return 14
   case 200:
      return 28
   case 201:
      return 56
   case 202:
      return 112
   case 203:
      return 224
   case 204:
      return 221
   case 205:
      return 167
   case 206:
      return 83
   case 207:
      return 166
   case 208:
      return 81
   case 209:
      return 162
   case 210:
      return 89
   case 211:
      return 178
   case 212:
      return 121
   case 213:
      return 242
   case 214:
      return 249
   case 215:
      return 239
   case 216:
      return 195
   case 217:
      return 155
   case 218:
      return 43
   case 219:
      return 86
   case 220:
      return 172
   case 221:
      return 69
   case 222:
      return 138
   case 223:
      return 9
   case 224:
      return 18
   case 225:
      return 36
   case 226:
      return 72
   case 227:
      return 144
   case 228:
      return 61
   case 229:
      return 122
   case 230:
      return 244
   case 231:
      return 245
   case 232:
      return 247
   case 233:
      return 243
   case 234:
      return 251
   case 235:
      return 235
   case 236:
      return 203
   case 237:
      return 139
   case 238:
      return 11
   case 239:
      return 22
   case 240:
      return 44
   case 241:
      return 88
   case 242:
      return 176
   case 243:
      return 125
   case 244:
      return 250
   case 245:
      return 233
   case 246:
      return 207
   case 247:
      return 131
   case 248:
      return 27
   case 249:
      return 54
   case 250:
      return 108
   case 251:
      return 216
   case 252:
      return 173
   case 253:
      return 71
   case 254:
      return 142
   case 255:
      return 1
   default:
      return -1
   }
}

func intToPower(n int) int {
   switch n {
   case 1:
      return 0
   case 2:
      return 1
   case 3:
      return 25
   case 4:
      return 2
   case 5:
      return 50
   case 6:
      return 26
   case 7:
      return 198
   case 8:
      return 3
   case 9:
      return 223
   case 10:
      return 51
   case 11:
      return 238
   case 12:
      return 27
   case 13:
      return 104
   case 14:
      return 199
   case 15:
      return 75
   case 16:
      return 4
   case 17:
      return 100
   case 18:
      return 224
   case 19:
      return 14
   case 20:
      return 52
   case 21:
      return 141
   case 22:
      return 239
   case 23:
      return 129
   case 24:
      return 28
   case 25:
      return 193
   case 26:
      return 105
   case 27:
      return 248
   case 28:
      return 200
   case 29:
      return 8
   case 30:
      return 76
   case 31:
      return 113
   case 32:
      return 5
   case 33:
      return 138
   case 34:
      return 101
   case 35:
      return 47
   case 36:
      return 225
   case 37:
      return 36
   case 38:
      return 15
   case 39:
      return 33
   case 40:
      return 53
   case 41:
      return 147
   case 42:
      return 142
   case 43:
      return 218
   case 44:
      return 240
   case 45:
      return 18
   case 46:
      return 130
   case 47:
      return 69
   case 48:
      return 29
   case 49:
      return 181
   case 50:
      return 194
   case 51:
      return 125
   case 52:
      return 106
   case 53:
      return 39
   case 54:
      return 249
   case 55:
      return 185
   case 56:
      return 201
   case 57:
      return 154
   case 58:
      return 9
   case 59:
      return 120
   case 60:
      return 77
   case 61:
      return 228
   case 62:
      return 114
   case 63:
      return 166
   case 64:
      return 6
   case 65:
      return 191
   case 66:
      return 139
   case 67:
      return 98
   case 68:
      return 102
   case 69:
      return 221
   case 70:
      return 48
   case 71:
      return 253
   case 72:
      return 226
   case 73:
      return 152
   case 74:
      return 37
   case 75:
      return 179
   case 76:
      return 16
   case 77:
      return 145
   case 78:
      return 34
   case 79:
      return 136
   case 80:
      return 54
   case 81:
      return 208
   case 82:
      return 148
   case 83:
      return 206
   case 84:
      return 143
   case 85:
      return 150
   case 86:
      return 219
   case 87:
      return 189
   case 88:
      return 241
   case 89:
      return 210
   case 90:
      return 19
   case 91:
      return 92
   case 92:
      return 131
   case 93:
      return 56
   case 94:
      return 70
   case 95:
      return 64
   case 96:
      return 30
   case 97:
      return 66
   case 98:
      return 182
   case 99:
      return 163
   case 100:
      return 195
   case 101:
      return 72
   case 102:
      return 126
   case 103:
      return 110
   case 104:
      return 107
   case 105:
      return 58
   case 106:
      return 40
   case 107:
      return 84
   case 108:
      return 250
   case 109:
      return 133
   case 110:
      return 186
   case 111:
      return 61
   case 112:
      return 202
   case 113:
      return 94
   case 114:
      return 155
   case 115:
      return 159
   case 116:
      return 10
   case 117:
      return 21
   case 118:
      return 121
   case 119:
      return 43
   case 120:
      return 78
   case 121:
      return 212
   case 122:
      return 229
   case 123:
      return 172
   case 124:
      return 115
   case 125:
      return 243
   case 126:
      return 167
   case 127:
      return 87
   case 128:
      return 7
   case 129:
      return 112
   case 130:
      return 192
   case 131:
      return 247
   case 132:
      return 140
   case 133:
      return 128
   case 134:
      return 99
   case 135:
      return 13
   case 136:
      return 103
   case 137:
      return 74
   case 138:
      return 222
   case 139:
      return 237
   case 140:
      return 49
   case 141:
      return 197
   case 142:
      return 254
   case 143:
      return 24
   case 144:
      return 227
   case 145:
      return 165
   case 146:
      return 153
   case 147:
      return 119
   case 148:
      return 38
   case 149:
      return 184
   case 150:
      return 180
   case 151:
      return 124
   case 152:
      return 17
   case 153:
      return 68
   case 154:
      return 146
   case 155:
      return 217
   case 156:
      return 35
   case 157:
      return 32
   case 158:
      return 137
   case 159:
      return 46
   case 160:
      return 55
   case 161:
      return 63
   case 162:
      return 209
   case 163:
      return 91
   case 164:
      return 149
   case 165:
      return 188
   case 166:
      return 207
   case 167:
      return 205
   case 168:
      return 144
   case 169:
      return 135
   case 170:
      return 151
   case 171:
      return 178
   case 172:
      return 220
   case 173:
      return 252
   case 174:
      return 190
   case 175:
      return 97
   case 176:
      return 242
   case 177:
      return 86
   case 178:
      return 211
   case 179:
      return 171
   case 180:
      return 20
   case 181:
      return 42
   case 182:
      return 93
   case 183:
      return 158
   case 184:
      return 132
   case 185:
      return 60
   case 186:
      return 57
   case 187:
      return 83
   case 188:
      return 71
   case 189:
      return 109
   case 190:
      return 65
   case 191:
      return 162
   case 192:
      return 31
   case 193:
      return 45
   case 194:
      return 67
   case 195:
      return 216
   case 196:
      return 183
   case 197:
      return 123
   case 198:
      return 164
   case 199:
      return 118
   case 200:
      return 196
   case 201:
      return 23
   case 202:
      return 73
   case 203:
      return 236
   case 204:
      return 127
   case 205:
      return 12
   case 206:
      return 111
   case 207:
      return 246
   case 208:
      return 108
   case 209:
      return 161
   case 210:
      return 59
   case 211:
      return 82
   case 212:
      return 41
   case 213:
      return 157
   case 214:
      return 85
   case 215:
      return 170
   case 216:
      return 251
   case 217:
      return 96
   case 218:
      return 134
   case 219:
      return 177
   case 220:
      return 187
   case 221:
      return 204
   case 222:
      return 62
   case 223:
      return 90
   case 224:
      return 203
   case 225:
      return 89
   case 226:
      return 95
   case 227:
      return 176
   case 228:
      return 156
   case 229:
      return 169
   case 230:
      return 160
   case 231:
      return 81
   case 232:
      return 11
   case 233:
      return 245
   case 234:
      return 22
   case 235:
      return 235
   case 236:
      return 122
   case 237:
      return 117
   case 238:
      return 44
   case 239:
      return 215
   case 240:
      return 79
   case 241:
      return 174
   case 242:
      return 213
   case 243:
      return 233
   case 244:
      return 230
   case 245:
      return 231
   case 246:
      return 173
   case 247:
      return 232
   case 248:
      return 116
   case 249:
      return 214
   case 250:
      return 244
   case 251:
      return 234
   case 252:
      return 168
   case 253:
      return 80
   case 254:
      return 88
   case 255:
      return 175
   default:
      return -1
   }
}

func multiplyPow(msg []GeneratorPoly, n int) []GeneratorPoly {
   res := msg

   for i := range res {
      res[i].xPower += n
   }

   return res
}

func getMessagePolynomial(msg []int) []Poly {
   result := []Poly{}

   for index, value := range msg {
      nPower := len(msg) - (index + 1) + CW

      result = append(result, Poly{ value, nPower })
   }

   return result 
}

func polyToInt(poly []Poly) []int {
   result := []int{}

   for _, value := range poly {
      x := value.x

      if x != 0 {
         result = append(result, x)
      }

   }

   return result
}

func multiplyGenMsg(msg []Poly, generator []GeneratorPoly) []Poly {
   alphaPower := intToPower(msg[0].x)
   genMultiplyed := generator

   // multiply power
   for i, value := range genMultiplyed {
      x := 0
      xPower := 0

      value.alphaPower += alphaPower

      if value.alphaPower > 255 {
         value.alphaPower = value.alphaPower % 255
      }

      // calculate
      value.x = powerToInt(value.alphaPower)
      value.alphaPower = 0
      value.alpha = 0


      // XOR result
      x = msg[i].x ^ value.x
      xPower = value.xPower

      msg[i].x = x
      msg[i].power = xPower
   }

   diff := len(msg) - len(genMultiplyed)
   for i:=0; i < diff; i++ {
      index := len(genMultiplyed) + i

      msg[index].x = msg[index].x ^ 0
   }

   if msg[0].x == 0 {
      msg = msg[1:]
      msg = append(msg, Poly{
         0, 0,
      })
   }



   return msg
}

func Encoder(msg []int) []int {
   msgPoly := getMessagePolynomial(msg)

   generator := []GeneratorPoly{ {ALPHA, 0, 1, 10}, {ALPHA, 251, 1, 9}, {ALPHA, 67, 1, 8}, {ALPHA, 46, 1, 7}, {ALPHA, 61, 1, 6}, {ALPHA, 118, 1, 5}, {ALPHA, 70, 1, 4}, {ALPHA, 64, 1, 3}, {ALPHA, 94, 1, 2}, {ALPHA, 32, 1, 1}, {ALPHA, 45, 1, 0} }

   generatorDiff := msgPoly[0].power - generator[0].xPower
   generatorPoly := multiplyPow(generator, generatorDiff)

   for i:=0; i < len(msg); i++ {
      msgPoly = multiplyGenMsg(msgPoly, generatorPoly)
   }

   return polyToInt(msgPoly)
}
