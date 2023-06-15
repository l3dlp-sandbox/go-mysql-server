// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encodings

// Latin7_estonian_cs_RuneWeight returns the weight of a given rune based on its relational sort order from
// the `latin7_estonian_cs` collation.
func Latin7_estonian_cs_RuneWeight(r rune) int32 {
	weight, ok := latin7_estonian_cs_Weights[r]
	if ok {
		return weight
	} else {
		return 2147483647
	}
}

// latin7_estonian_cs_Weights contain a map from rune to weight for the `latin7_estonian_cs` collation. The
// map primarily contains mappings that have a random order. Mappings that fit into a sequential range (and are long
// enough) are defined in the calling function to save space.
var latin7_estonian_cs_Weights = map[rune]int32{
	0:    0,
	128:  1,
	1:    2,
	2:    3,
	3:    4,
	4:    5,
	5:    6,
	6:    7,
	7:    8,
	8:    9,
	14:   10,
	15:   11,
	16:   12,
	17:   13,
	18:   14,
	19:   15,
	20:   16,
	21:   17,
	22:   18,
	23:   19,
	24:   20,
	25:   21,
	26:   22,
	27:   23,
	28:   24,
	29:   25,
	30:   26,
	31:   27,
	127:  28,
	129:  29,
	131:  30,
	136:  31,
	138:  32,
	140:  33,
	144:  34,
	152:  35,
	154:  36,
	156:  37,
	159:  38,
	39:   39,
	45:   40,
	173:  41,
	150:  42,
	151:  43,
	32:   44,
	160:  45,
	9:    46,
	10:   47,
	11:   48,
	12:   49,
	13:   50,
	33:   51,
	34:   52,
	35:   53,
	36:   54,
	37:   55,
	38:   56,
	40:   57,
	41:   58,
	42:   59,
	44:   60,
	46:   61,
	47:   62,
	58:   63,
	59:   64,
	63:   65,
	64:   66,
	91:   67,
	92:   68,
	93:   69,
	94:   70,
	95:   71,
	96:   72,
	123:  73,
	124:  74,
	125:  75,
	126:  76,
	166:  77,
	141:  78,
	157:  79,
	8220: 80,
	143:  81,
	142:  82,
	8217: 83,
	158:  84,
	145:  85,
	146:  86,
	130:  87,
	147:  88,
	148:  89,
	132:  90,
	139:  91,
	155:  92,
	43:   93,
	60:   94,
	61:   95,
	62:   96,
	177:  97,
	171:  98,
	187:  99,
	215:  100,
	247:  101,
	162:  102,
	163:  103,
	164:  104,
	167:  105,
	169:  106,
	172:  107,
	174:  108,
	176:  109,
	181:  110,
	182:  111,
	183:  112,
	134:  113,
	135:  114,
	149:  115,
	133:  116,
	137:  117,
	48:   118,
	188:  119,
	189:  120,
	190:  121,
	49:   122,
	185:  123,
	50:   124,
	178:  125,
	51:   126,
	179:  127,
	52:   128,
	53:   129,
	54:   130,
	55:   131,
	56:   132,
	57:   133,
	65:   134,
	97:   135,
	256:  136,
	257:  137,
	197:  138,
	229:  139,
	260:  140,
	261:  141,
	198:  142,
	230:  143,
	66:   144,
	98:   145,
	67:   146,
	99:   147,
	262:  148,
	263:  149,
	268:  150,
	269:  151,
	68:   152,
	100:  153,
	69:   154,
	101:  155,
	201:  156,
	233:  157,
	278:  158,
	279:  159,
	274:  160,
	275:  161,
	280:  162,
	281:  163,
	70:   164,
	102:  165,
	71:   166,
	103:  167,
	290:  168,
	291:  169,
	72:   170,
	104:  171,
	73:   172,
	105:  173,
	298:  174,
	299:  175,
	302:  176,
	303:  177,
	74:   178,
	106:  179,
	75:   180,
	107:  181,
	310:  182,
	311:  183,
	76:   184,
	108:  185,
	315:  186,
	316:  187,
	321:  188,
	322:  189,
	77:   190,
	109:  191,
	78:   192,
	110:  193,
	323:  194,
	324:  195,
	325:  196,
	326:  197,
	79:   198,
	111:  199,
	211:  200,
	243:  201,
	332:  202,
	333:  203,
	216:  204,
	248:  205,
	80:   206,
	112:  207,
	81:   208,
	113:  209,
	82:   210,
	114:  211,
	342:  212,
	343:  213,
	83:   214,
	115:  215,
	346:  216,
	347:  217,
	223:  218,
	352:  219,
	353:  220,
	90:   221,
	122:  222,
	377:  223,
	378:  224,
	379:  225,
	380:  226,
	381:  227,
	382:  228,
	84:   229,
	116:  230,
	153:  231,
	85:   232,
	117:  233,
	362:  234,
	363:  235,
	370:  236,
	371:  237,
	86:   238,
	118:  239,
	87:   240,
	119:  241,
	213:  242,
	245:  243,
	196:  244,
	228:  245,
	214:  246,
	246:  247,
	220:  248,
	252:  249,
	88:   250,
	120:  251,
	89:   252,
	121:  253,
	8221: 254,
	8222: 255,
}