package linear

import (
	"fmt"
	"log"
	"os"

	"gitee.com/quant1x/pandas"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	ColNames = []string{"feature", "document", "machine", "load_time",
		"search_time", "reduce_and_save"}

	ResColNames = []string{"feature", "document", "machine", "total"}
)

// dataPrepare 数据预处理函数
func dataPrepare(clsDF *pandas.DataFrame) {
	// 获取total列
	*clsDF = clsDF.Select(ColNames)
	totalSeries := clsDF.Rapply(getTotal)
	totalSeries.SetNames("total")
	*clsDF = clsDF.CBind(totalSeries)

	// document列 *2/1000
	*clsDF = clsDF.Select(ResColNames)
	newDocSeries := clsDF.Rapply(getDoc)
	newDocSeries.SetNames("new_doc")
	*clsDF = clsDF.CBind(newDocSeries)
	*clsDF = clsDF.Drop([]string{"document"})
	*clsDF = clsDF.Rename("document", "new_doc")
	*clsDF = clsDF.Select(ResColNames)
}

func dataPlot(actPoints, expPoints plotter.XYs) {
	plt := plot.New()
	//if err != nil {
	//	panic(err)
	//}
	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	if err := plotutil.AddLinePoints(plt,
		"expPoints", expPoints,
		"actPoints", actPoints,
	); err != nil {
		panic(err)
	}

	if err := plt.Save(5*vg.Inch, 5*vg.Inch, "classification-fit.png"); err != nil {
		panic(err)
	}
}

// FitClassification 分类曲线拟合函数
func FitClassification() {
	clsData, err := os.Open("classification_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer clsData.Close()
	clsDF := pandas.ReadCSV(clsData)
	// 数据预处理
	dataPrepare(&clsDF)
	// 数据预处理完成
	fmt.Println("数据预处理完成...")
	fmt.Println(clsDF)

	// 数据拟合
	actPoints, expPoints, fa, fb := dataOptimize(&clsDF)
	// 拟合完成，输出fa,fb
	fmt.Println("Fa", fa, "Fb", fb)

	// 数据绘图
	dataPlot(actPoints, expPoints)
	fmt.Println("绘制完成，图形地址: classification-fit.png")
}
