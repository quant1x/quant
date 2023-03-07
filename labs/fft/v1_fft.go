package fft

import "math"

// 输入初始数据切片pr，切片长度n，在pr上操作，返回fft后的pr
func v1_fft(pr []float64, n int) []float64 {
	k := int(math.Floor(math.Log2(float64(n))))
	var pi = make([]float64, n)
	var fi = make([]float64, n)
	var fr = make([]float64, n)
	var it, m, is, i, j, nv, l0 int
	var p, q, s, vr, vi, poddr, poddi float64
	for it = 0; it <= n-1; it++ {
		m = it
		is = 0
		for i = 0; i <= k-1; i++ {
			j = m / 2
			is = 2*is + (m - 2*j)
			m = j
		}
		fr[it] = pr[is]
		fi[it] = pi[is]
	}
	pr[0] = 1.0
	pi[0] = 0.0
	p = 2 * math.Pi / float64(n)
	pr[1] = math.Cos(p) //将w=e^-j2pi/n用欧拉公式表示
	pi[1] = -math.Sin(p)

	for i = 2; i <= n-1; i++ {
		p = pr[i-1] * pr[1]
		q = pi[i-1] * pi[1]
		s = (pr[i-1] + pi[i-1]) * (pr[1] + pi[1])
		pr[i] = p - q
		pi[i] = s - p - q
	}
	for it = 0; it <= n-2; it = it + 2 {
		vr = fr[it]
		vi = fi[it]
		fr[it] = vr + fr[it+1]
		fi[it] = vi + fi[it+1]
		fr[it+1] = vr - fr[it+1]
		fi[it+1] = vi - fi[it+1]
	}
	m = n / 2
	nv = 2
	for l0 = k - 2; l0 >= 0; l0-- { //蝴蝶操作
		m = m / 2
		nv = 2 * nv
		for it = 0; it <= (m-1)*nv; it = it + nv {
			for j = 0; j <= (nv/2)-1; j++ {

				p = pr[m*j] * fr[it+j+nv/2]
				q = pi[m*j] * fi[it+j+nv/2]
				s = pr[m*j] + pi[m*j]
				s = s * (fr[it+j+nv/2] + fi[it+j+nv/2])
				poddr = p - q
				poddi = s - p - q
				fr[it+j+nv/2] = fr[it+j] - poddr
				fi[it+j+nv/2] = fi[it+j] - poddi
				fr[it+j] = fr[it+j] + poddr
				fi[it+j] = fi[it+j] + poddi
			}
		}
	}

	for i = 0; i <= n-1; i++ {
		pr[i] = math.Sqrt(fr[i]*fr[i] + fi[i]*fi[i]) //幅值计算
	}
	return pr
}
