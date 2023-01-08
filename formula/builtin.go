package formula

//type Float float64

const (
	FloatDefault float64 = 0.00
)

// REF 引用
func REF(slice []float64, n int) float64 {
	count := len(slice)
	if count <= n {
		return FloatDefault
	}
	return slice[count-1-n]
}

// 指标计算接口
type algorithmHandler = func(a, b float64) float64

// 切片通用遍历方法
func slice_universal[V int64 | float64](slice []V, n int, iterator func(a, b V) V) V {
	count := len(slice)
	// TODO 小于等于 改为 小于, 允许最大尺寸查询, 当天也算
	if count < n {
		return V(0)
	}
	var (
		ret    V    = 0
		inited bool = false
	)
	pos := count - n
	for i := 0; i < n; i++ {
		cur := slice[pos+i]
		if !inited {
			ret = cur
			inited = true
			continue
		}
		ret = iterator(ret, cur)
	}
	return ret
}

// SUM 计算n周期内的flag的总和
func SUM[V int64 | float64](slice []V, n int) V {
	return slice_universal(slice, n, func(a, b V) V {
		return a + b
	})
}

// 计算n周期的算术平均值
func MA[V int64 | float64](slice []V, n int) V {
	count := len(slice)
	// TODO 注意这个点可能存在问题
	if count < n {
		return V(0)
	}
	v := SUM(slice, n)
	return v / V(n)
}

// 这个sma公式有个计算起点，当周期数小于等于n时，是按ma计算的。超过n后，就可以按sma计算了。
// 示例：Y = SMA（收盘价，30，2）= [（当天的收盘价* 2 + Y *前一个交易日的（30-1））/ 30，即今天收盘价的扩展乘以2并计算30天平均值，今天的价值对平均值有很大的影响，即，平均价格中的权重很大。
// 在大多数情况下，M取1而不取2，Y是循环参考或先前的值，类似于a = a +1。严格来说，Y不是平均值，包括股票上市首日的价值，但权重越远，价值就越小，其变化比平均值ma慢。
//
// SMA SMA(X,N,M), 求X的N日移动平均，M为权重。算法：若Y=SMA(X,N,M) 则 Y=(M*X+(N-M)*Y')/N，其中Y'表示上一周期Y值，N必须大于M。
func SMA(slice []float64, n, m int) float64 {
	if n <= m {
		return FloatDefault
	}
	//count := len(slice)
	_n := float64(n)
	_m := float64(m)
	pervious := FloatDefault
	for i, x := range slice {
		if i < n-1 {
			//pervious = MA(slice[:i+1], n)
			pervious = (pervious*_n + x) / _n
		} else {
			pervious = (_m*x + (_n-_m)*pervious) / _n
		}
	}

	return pervious
}

// EMA（Exponential Moving Average）是指数移动平均值。
// 也叫EXPMA指标，它也是一种趋向类指标，指数移动平均值是以指数式递减加权的移动平均。
// EXPMA=（当日或当期收盘价－上一日或上期EXPMA）/N+上一日或上期EXPMA，其中，首次上期EXPMA值为上一期收盘价，N为天数。

// 当天EMA=昨天的EMA+加权因子*（当天的收盘价-昨天的EMA）
// = 加权因子*当天的收盘价+（1-加权因子）*昨天的EMA
// 加权因子=2/(N+1);
// N就是上面所说的周期 ，比如周期12 则加权的因子就是 2/13；
// 当天EMA=2/13*当天的收盘价+11/13*昨天的EMA
// 计算过程：（每日你看到的EMA计算结果是从上市第一天就开始累积了）
// 股票上市第一天：当天EMA1 = 当天收盘价
// 第二天：EMA2 = 2/13 * 当天收盘价 + 11/13 * EMA1
// 第三天：EMA3 = 2/13 * 当天收盘价 + 11/13 * EMA2

// 若求X的N日指数平滑移动平均, 则表达式为: EMA(X, N)
// 算法是: 若Y=EMA(X, N), 则Y=[2*X+(N-1)*Y’]/(N+1), 其中Y’表示上一周期的Y值。
func EXMPA(hds []float64, n int, x int) float64 {
	count := len(hds)
	if n > x+1 || count < n {
		return 0.000
	}

	// Weighting factor, 计算加权因子
	factor := float64(2.0000 / (n + 1.0000))
	// 第一天ema等于当天收盘价
	ema := hds[0]
	for i := 1; i < n; i++ {
		hd := hds[i]
		// 第二天以后，当天收盘 收盘价乘以系数再加上昨天EMA乘以1-系数
		ema = hd*factor + ema*(1.0000-factor)
	}
	return ema
}
