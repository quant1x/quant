package main

/*
N:=15;
M:=9;
VAR1:=C-REF(C,1);
VAR2:=O-REF(O,1);
VAR3:=H-REF(H,1);
VAR4:=L-REF(L,1);
VAR5:=(VAR1+VAR2+VAR3+VAR4)/4;
VJ:=VAR5;
VAR9:=IF(CAPITAL<=0,1,100/CAPITAL);
{VL:=VAR9*VOL;}
VL:=VOL;
LJ:=VJ*VL;

K:SUM(LJ,N),LINETHICK1,COLORWHITE;
D:EMA(K,M),COLORRED,LINETHICK2;
IF(D<REF(D,1), D,DRAWNULL),COLORGREEN,LINETHICK2;

0,COLORGRAY,DOTLINE;
DRAWTEXTABS(0,0,'QQ交流群:142713'), COLORYELLOW;

B1:=REF(K,2)<REF(K,1) AND REF(K,1)>K;
B2:=D>REF(D,1) AND CROSS(K,D);
B3:=BARSSINCEN(B2>0,N);
B4:=COUNT(B1,B3);
S:IF(D>REF(D,1) AND B4<7,B4,0),COLORGREEN,NODRAW;
S1:=REF(K,2)>REF(K,1) AND REF(K,1)<K;
S2:=D<REF(D,1) AND CROSS(D,K);
S3:=BARSSINCEN(S2>0,N);
S4:=COUNT(S1,S3);
B:IF(D<REF(D,1) AND S4<7,S4,0),COLORRED,NODRAW;
*/

/*
1号策略 5日均线10日均线金叉后的第一个收盘价等于或者低于10日线，MACD红柱。
金叉后的第一次最低价低于10日线
2号策略 5日均线小于10日均线，当日收盘小于10日均线
*/

/*
短线策略

ZLW0:=IF((CLOSE > 200),(CLOSE * 1.01),(CLOSE * 1.07));
ZLW1:=IF((CLOSE < 10),(CLOSE * 1.05),ZLW0);
ZSW0:=IF((CLOSE > 200),(CLOSE * 0.99),(CLOSE * 0.93));
ZSW1:=IF((CLOSE < 10),(CLOSE * 0.95),ZSW0);

ZSW:ZSW1;
ZLW:ZLW1;

PT:=REF(HIGH,1)-REF(LOW,1);
ZX:(HIGH + LOW + CLOSE)/3;
YL1:2*ZX-LOW;
YL2:ZX + PT;
ZC1:2*ZX-HIGH;
ZC2:ZX-PT;
*/