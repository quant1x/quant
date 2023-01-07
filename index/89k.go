package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
)

type KNode struct {
	K float64 `name:"价格" json:"k"`
	N int     `name:"周期" json:"n"`
}

func (this *KNode) Reset() {
	this.K = 0.00
	this.N = 0
}

func (this *KNode) Set(k float64, n int) {
	this.K = k
	this.N = n
}

func (this *KNode) Comp(k float64) int {
	if this.K <= 0.00 || this.N == 0 {
		return 0
	}
	if this.K == k {
		return 0
	}
	if this.K < k {
		return -1
	}
	return 1
}

func (this *KNode) Lt(k float64) bool {
	return this.Comp(k) <= 0
}

func (this *KNode) Gt(k float64) bool {
	return this.Comp(k) >= 0
}

type K89 struct {
	data   []MaLine
	P0     KNode
	P5     KNode
	P7     KNode
	P8     KNode
	P9     KNode
	P10    KNode
	P11    KNode
	Signal MaLine
}

func (this *K89) Reset() {
	this.P0.Reset()
	this.P5.Reset()
	this.P7.Reset()
	this.P8.Reset()
	this.P9.Reset()
	this.P10.Reset()
	this.P11.Reset()
}

func (self *K89) Len() int {
	return len(self.data)
}

func (self *K89) Data() interface{} {
	return self.Signal
}

// 89日均线

/*
{89K图形}

MA89:=MA(CLOSE,N1);
{计算N1日内最低价}
L89:=LLV(LOW,N1);
{计算N1日内最高价}
H89:=HHV(HIGH,N1);

{确定⑤}
K5:L89,NODRAW,COLORGREEN;
{确定⑦}
N7:BARSLAST(H89=HIGH),NODRAW,COLORLIGRAY;
T7:=REF(HIGH,N7);
K7:T7,NODRAW,COLORRED;
{确定⑧}
T8:=LLV(LOW,N7),NODRAW,COLORGREEN;
N8:BARSLAST(T8=LOW AND N1>N7),NODRAW,COLORLIGRAY;
K8:T8,NODRAW,COLORGREEN;
{确定⑨}
T9:=HHV(HIGH,N8);
N9X:=BARSLAST(T9=HIGH AND N7>N8);
N9:IFF(N9X=0,N9X+1,N9X),NODRAW,COLORLIGRAY;
K9:REF(HIGH,N9X),NODRAW,COLORRED;
{确定⑩}
K10:LLV(LOW,N9),NODRAW,COLORGREEN;
N10:BARSLAST(K10=LOW AND N8>N9),NODRAW,COLORLIGRAY;

{比对周期长度}
C_N:=5;
{量价关系最低校对比率}
C_S:=0.191;
{涂画与股价的纵向比率}
C_PX:=0.002;
{真阳线}
{C_ISMALE:=(CLOSE > REF(CLOSE,1)) AND (CLOSE > OPEN);}
C_ISMALE:=CLOSE > REF(CLOSE,1);
{成交量较上一个周期放大}
C_VOL:= VOL>REF(VOL,1);
{成交量均线周期}
VOL_PERIOD:=5;
{成交量比例}
VOLSCALE:=1+C_S;
{高股价或指数的计算方法, 比MAVOL5高出C_S/10且比前一日方法}
X_INDEX:=VOL>=MA(VOL, VOL_PERIOD)*(1 + C_S/10);
{一般股价的计算方法}
X_GENERAL:=VOL>=MA(VOL, VOL_PERIOD)*VOLSCALE;
{指数类或者高股价类的成交量不太可能像个股那样成倍放量, 这里做一个降级处理}
X:=IFF(CLOSE>=100, X_INDEX, X_GENERAL) AND C_ISMALE AND C_VOL;
{确定X上一次确立成立是在哪一天}
DN:=BARSLAST(X);
{放量上攻作为一个小阶段的起点, 该起点K线的最低价作为止盈止损线}
ZS_VOL:=REF(VOL,DN);
ZS_LOW:=REF(LOW,DN);
ZS:ZS_LOW,NODRAW,COLORLIGRAY;
N11:BARSLAST(CLOSE<ZS_LOW AND N9>N10),NODRAW,COLORLIGRAY;
{股价突破止损线}
{C0:BARSSINCEN(CLOSE<ZS_LOW,DN),NODRAW;}
{C0:FINDHIGHBARS(CLOSE,1,DN-1,1),NODRAW;}
C00:=BARSLASTCOUNT(CLOSE<ZS_LOW),NODRAW;
C01:=BARSLASTCOUNT(CLOSE>=ZS_LOW),NODRAW;
C02:=IFF(C01>0,REF(C00,1),C01),NODRAW;
C03:=REF(CLOSE,C02-1),NODRAW;
C1:BARSLAST(CLOSE>ZS_LOW AND N9>N10),NODRAW,COLORLIGRAY;
{股价跌破止损线}
C2:BARSLAST(CLOSE<ZS_LOW AND N9>N10),NODRAW,COLORLIGRAY;
C3:N10>0,NODRAW;
B0:=C1=0 AND C2=1 AND C3;
B:50*B0,COLORYELLOW;

{筹码锁定因子}
CM0:=50-LFS;
CM:CM0,COLORSTICK;
CM1:EMA(CM0,3);

{斜率}
XL:(K7-K9)/(N7-N9),NODRAW;
P0:CLOSE,COLORWHITE;
P1:-1*XL*N9+K9,COLORCYAN;
*/
func (self *K89) Load(code string) error {
	kls := cache.LoadKLine(code)
	if kls == nil {
		return ErrCode
	} else if len(kls) < 1 {
		return ErrData
	}

	count := len(kls)
	const w = 89
	var (
		t        K89
		ma89             = 0.00
		zhisun   float64 = 0
		huicai   bool    = false
		bConsole         = true
	)
	for i, v := range kls {
		var (
			price  float64
			volume int64
		)

		tmp := MaLine{DayKLine: v}

		hds := kls[:i+1]
		// 第一步, 计算MA89
		if i+1 >= w {
			price = SUM(hds, Close, w) / float64(w)
			volume = int64(SUM(hds, Volume, w)) / int64(w)
			tmp.MA30 = price
			tmp.MA30Volume = volume
			// 重置MA89均线
			ma89 = price
		} else {
			continue
		}
		// 第二步, 低于收盘价低于MA89
		if v.Close < ma89 {
			// 重置⑤
			if t.P5.Gt(v.Low) {
				t.Reset()
				huicai = false
				zhisun = 0
				t.P5.Set(v.Low, i)
				if bConsole {
					fmt.Printf("%s, ⑤: %.2f\n", v.Date, v.Low)
				}
				continue
			}
		}
		// 第三步, 股价不再创新低之后
		if t.P5.Lt(v.Low) {
			// 找从P5开始的最高价
			n5 := i - t.P5.N + 1
			hp := HHV(hds, High, n5)
			if t.P7.Lt(hp) {
				t.P7.Set(hp, i)
				if bConsole {
					fmt.Printf("%s, ⑦: %.2f\n", v.Date, hp)
				}
			}

			use89k := true

			if !use89k {
				hv := HHV(hds, Volume, n5)
				if int64(hv) == v.Volume {
					zhisun = v.Low
					if bConsole {
						fmt.Printf("%s, 画止损线: %.2f\n", v.Date, zhisun)
					}
					continue
				}
			} else /*if zhisun == 0.00*/ {
				n := BARSSINCEN(hds, Volume, n5, func(a, b float64) bool {
					if a == 0.00 {
						return false
					}
					return b/a >= 2
				})
				if n > 0 {
					v := hds[i-n]
					zhisun = v.Low
					if bConsole {
						fmt.Printf("%s, 画止损线: %.2f\n", v.Date, zhisun)
					}
					//continue
				}
			}
		}
		// 第四步, 股价不再创新高之后
		if t.P7.Gt(v.High) {
			// 开始回落
			n7 := i - t.P7.N + 1
			lc := LLV(hds, Close, n7)
			if lc < zhisun {
				huicai = true
				if bConsole {
					fmt.Printf("%s, 跌破止损: %.2f, 收盘: %.2f, 卖出\n", v.Date, zhisun, v.Close)
				}
				continue
			}
			// 重置⑧
			if t.P8.Gt(v.Low) {
				t.P8.Set(v.Low, i)
				if bConsole {
					fmt.Printf("%s, ⑧: %.2f\n", v.Date, v.Low)
				}
				continue
			}
		}
		// 第五步, 股价第2次不再新低
		if huicai {
			if zhisun < v.Close {
				if bConsole {
					fmt.Printf("%s, 回踩止损: %.2f, 突破, 收盘: %.2f, 买入\n", v.Date, zhisun, v.Close)
				}
				huicai = false
				zhisun = 0
				t.Reset()
				self.Signal = tmp
			}
		}
		self.data = append(self.data, tmp)
		// 输出最后2组数据
		if category.DEBUG && count < i+3 {
			//fmt.Printf("day: %s, MA5: %.2f, MA10: %.2f, MA20: %.2f, MA30: %.2f\n", self.data[i].Date, self.data[i].MA5, self.data[i].MA10, self.data[i].MA20, self.data[i].MA30)
		}
	}
	return nil
}
