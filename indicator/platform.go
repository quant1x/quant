package indicator

import (
	"gitee.com/quant1x/pandas"
	. "gitee.com/quant1x/pandas/formula"
	"gitee.com/quant1x/pandas/stat"
)

// Platform 平台整理
func Platform(df pandas.DataFrame) pandas.DataFrame {
	var (
		OPEN  = df.ColAsNDArray("open")  // 开盘价
		CLOSE = df.ColAsNDArray("close") // 收盘价
		//HIGH    = df.ColAsNDArray("high")   // 最高价
		//LOW     = df.ColAsNDArray("low")    // 最低价
		VOL = df.ColAsNDArray("volume") // 成交量
		//DATALEN = df.Nrow()                 // 数据长度
	)
	//{T02: 平台, V1.0.0 2023-03-08}
	//BL1:=VOL/REF(VOL,1);
	LB := VOL.Div(REF(VOL, 1))
	//STH:=MAX(OPEN,CLOSE);
	STH := MAX(OPEN, CLOSE)
	//STL:=MIN(OPEN,CLOSE);
	STL := MIN(OPEN, CLOSE)
	//BLN1:=BARSLAST(BL1>=2.00),NODRAW;
	BLN := BARSLAST(LB.Gte(2.00))
	//倍量周期:BLN1,NODRAW;
	//{为HHV修复BLN1的值,需要+1}
	//BLN:=IFF(BLN1>=0,BLN1+1,BLN1);
	bln1 := IFF(BLN.Gte(0.00), BLN.Add(1), BLN)
	//倍量H:REF(STH,BLN);
	BLH := REF(STH, bln1)
	//倍量L:REF(STL,BLN);
	BLL := REF(STL, bln1)
	//BLVH:=HHV(VOL,BLN),NODRAW;
	BLVH := HHV(VOL, bln1)
	//BLVHN:=BARSLAST(VOL=BLVH),NODRAW;
	//BLVHN := BARSLAST(VOL.Eq(BLVH))
	//BLVL:=LLV(VOL,BLVHN),NODRAW;
	//BLVL := LLV(VOL, BLVHN)
	//BLVLN:=BARSLAST(VOL=BLVL),NODRAW;
	//BLVLN := BARSLAST(VOL.Eq(BLVL))
	//
	//SL:VOL/BLVH,NODRAW;
	SL := VOL.Div(BLVH)
	//SL21:=BARSLASTCOUNT(SL<=0.50),NODRAW;
	//SL21 := BARSLASTCOUNT(SL.Lte(0.50))
	SL21 := BARSSINCEN(SL.Lte(0.50), BLN)
	//{为REF修复SL21,需要-1}
	//SL2:=IFF(SL21>0,SL21-1,DRAWNULL),NODRAW;
	//SL2 := IFF(SL21.Gt(0), SL21.Sub(1), SL21)
	//SL2 := SL21.Sub(1)
	//SL2 = SL21
	//SL3:=REF(VOL,SL2);
	//SL3 := REF(VOL, SL2)
	//SLN:=BARSLAST(SL3=VOL);
	//SLN := BARSLAST(SL3.Eq(VOL))
	SLN := SL21
	//
	//缩量周期:SLN,NODRAW,COLORYELLOW;
	//缩量H:REF(STH,SLN+1),DOTLINE,COLORLIRED;
	SLH := REF(STH, SLN.Add(1))
	//缩量L:REF(STL,SLN),DOTLINE,COLORLIGREEN;
	SLL := REF(STL, SLN)

	oBlN := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLN", BLN)
	oBlH := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLH", BLH)
	oBlL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLL", BLL)
	oSL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SL", SL)
	oBLVH := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLVH", BLVH)
	oSlN := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLN", SLN)
	oSlH := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLH", SLH)
	oSlL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLL", SLL)

	df = df.Join(oBLVH).Join(oBlN).Join(oBlH).Join(oBlL).Join(oSlN).Join(oSlH).Join(oSlL).Join(oSL)
	return df
}

func v1Platform(df pandas.DataFrame) pandas.DataFrame {
	var (
		OPEN  = df.ColAsNDArray("open")  // 开盘价
		CLOSE = df.ColAsNDArray("close") // 收盘价
		//HIGH    = df.ColAsNDArray("high")   // 最高价
		//LOW     = df.ColAsNDArray("low")    // 最低价
		VOL = df.ColAsNDArray("volume") // 成交量
		//DATALEN = df.Nrow()                 // 数据长度
	)
	//{T02: 平台, V1.0.0 2023-03-08}
	//BL1:=VOL/REF(VOL,1);
	LB := VOL.Div(REF(VOL, 1))
	//STH:=MAX(OPEN,CLOSE);
	STH := MAX(OPEN, CLOSE)
	//STL:=MIN(OPEN,CLOSE);
	STL := MIN(OPEN, CLOSE)
	//BLN1:=BARSLAST(BL1>=2.00),NODRAW;
	BLN := BARSLAST(LB.Gte(2.00))
	//倍量周期:BLN1,NODRAW;
	//{为HHV修复BLN1的值,需要+1}
	//BLN:=IFF(BLN1>=0,BLN1+1,BLN1);
	bln1 := IFF(BLN.Gte(0.00), BLN.Add(1), BLN)
	//倍量H:REF(STH,BLN);
	BLH := REF(STH, bln1)
	//倍量L:REF(STL,BLN);
	BLL := REF(STL, bln1)
	//BLVH:=HHV(VOL,BLN),NODRAW;
	BLVH := HHV(VOL, bln1)
	//BLVHN:=BARSLAST(VOL=BLVH),NODRAW;
	//BLVHN := BARSLAST(VOL.Eq(BLVH))
	//BLVL:=LLV(VOL,BLVHN),NODRAW;
	//BLVL := LLV(VOL, BLVHN)
	//BLVLN:=BARSLAST(VOL=BLVL),NODRAW;
	//BLVLN := BARSLAST(VOL.Eq(BLVL))
	//
	//SL:VOL/BLVH,NODRAW;
	SL := VOL.Div(BLVH)
	//SL21:=BARSLASTCOUNT(SL<=0.50),NODRAW;
	SL21 := BARSLASTCOUNT(SL.Lte(0.50))
	//SL21 := formula.BARSSINCEN(SL.Lte(0.50), BLN)
	//{为REF修复SL21,需要-1}
	//SL2:=IFF(SL21>0,SL21-1,DRAWNULL),NODRAW;
	//SL2 := IFF(SL21.Gt(0), SL21.Sub(1), SL21)
	SL2 := SL21.Sub(1)
	//SL2 = SL21
	//SL3:=REF(VOL,SL2);
	SL3 := REF(VOL, SL2)
	//SLN:=BARSLAST(SL3=VOL);
	SLN := BARSLAST(SL3.Eq(VOL))
	//
	//缩量周期:SLN,NODRAW,COLORYELLOW;
	//缩量H:REF(STH,SLN+1),DOTLINE,COLORLIRED;
	SLH := REF(STH, SLN.Add(1))
	//缩量L:REF(STL,SLN),DOTLINE,COLORLIGREEN;
	SLL := REF(STL, SLN)

	oBlN := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLN", BLN)
	oBlH := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLH", BLH)
	oBlL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "BLL", BLL)
	oSL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SL", SL)
	oSlN := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLN", SLN)
	oSlH := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLH", SLH)
	oSlL := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "SLL", SLL)

	df = df.Join(oBlN).Join(oBlH).Join(oBlL).Join(oSlN).Join(oSlH).Join(oSlL).Join(oSL)
	return df
}
