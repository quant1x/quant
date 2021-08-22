package security

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/quant1x/quant/data/constant"
	"github.com/quant1x/quant/utils"
	"github.com/mymmsc/gox/errors"
	"github.com/mymmsc/gox/logger"
)

const (
	MARKET_SH string = "sh" // 上海
	MARKET_SZ string = "sz" // 深圳
	MARKET_HK string = "hk" // 香港
)

var (
	// ResourcesPath 资源路径
	ResourcesPath = "resources"
	// MarketName 市场名称
	MarketName = map[string]string{
		MARKET_SH: "上海",
		MARKET_SZ: "深圳",
		MARKET_HK: "香港",
	}
	MarketSecurity = map[string]int32{
		MARKET_SH: int32(constant.QotMarket_QotMarket_CNSH_Security),
		MARKET_SZ: int32(constant.QotMarket_QotMarket_CNSZ_Security),
		MARKET_HK: int32(constant.QotMarket_QotMarket_HK_Security),
	}
	// 证券市场缓存
	cacheSecurity = map[string]StaticBasic{}
)

var (
	// ErrCacheNotExist 没有缓存
	ErrCacheNotExist = errors.New("Cache not exist")
	// ErrCodeNotExist 证券代码不存在
	ErrCodeNotExist = errors.New("Securities code does not exist")
)

func GetStockCode(market string, code string) string {
	return fmt.Sprintf("%s%s", market, code)
}

//go:embed resources/*
var resources embed.FS

// Security 两个字段确定一支股票
type Security struct {
	Market int32  `json:"market,omitempty"` //QotMarket,股票市场
	Code   string `json:"code,omitempty"`   //股票代码
}

type StaticBasic struct {
	Security      Security `json:"security,omitempty"`      //股票
	Id            int64    `json:"id,omitempty"`            //股票ID
	LotSize       int32    `json:"lotSize,omitempty"`       //每手数量,期权以及期货类型表示合约乘数
	SecType       int32    `json:"secType,omitempty"`       //Qot_Common.SecurityType,股票类型
	Name          string   `json:"name,omitempty"`          //股票名字
	ListTime      string   `json:"listTime,omitempty"`      //上市时间字符串
	Delisting     bool     `json:"delisting,omitempty"`     //是否退市
	ListTimestamp float64  `json:"listTimestamp,omitempty"` //上市时间戳, 秒数
}

// 生成指数静态信息
func genIndexStaticInfo(market, code, name, listTime string, id int64, lotSize int32) (*StaticBasic, error) {
	marketSecurity, ok := MarketSecurity[market]
	if !ok {
		return nil, ErrCodeNotExist
	}
	fullCode := GetStockCode(market, code)
	listTimestamp, err := utils.ParseTime(listTime)
	if err != nil {
		return nil, err
	}
	return &StaticBasic{
		Security: Security{
			Market: marketSecurity,
			Code:   fullCode,
		},
		Id:            id,
		LotSize:       lotSize,
		SecType:       int32(constant.SecurityType_SecurityType_Index),
		Name:          name,
		ListTime:      listTime,
		Delisting:     false,
		ListTimestamp: float64(listTimestamp.Unix()),
	}, nil
}

func init() {
	logger.Infof("开始初始化市场静态数据...")
	// 1.上海指数
	// 上证综合指数 000001.sh
	/*code := GetStockCode(MARKET_SH, "000001")
	listTime := "1990-12-19"
	listTimestamp, _ := utils.ParseTime(listTime)
	index := StaticBasic{
		Security: Security{
			Market: int32(constant.QotMarket_QotMarket_CNSH_Security),
			Code: code,
		},
		Id:            1000001,
		LotSize:       100,
		SecType:       int32(constant.SecurityType_SecurityType_Index),
		Name:          "上证指数",
		ListTime:      listTime,
		Delisting:     false,
		ListTimestamp: float64(listTimestamp.Unix()),
	}*/

	index, err := genIndexStaticInfo(MARKET_SH, "000001", "上证指数", "1990-12-19", 1000001, 100)
	if err == nil && index != nil {
		cacheSecurity[index.Security.Code] = *index
	}

	// 中证500 sh000905 2007-01-15
	/*code = GetStockCode(MARKET_SH, "000905")
	listTime = "2007-01-15"
	listTimestamp, _ = utils.ParseTime(listTime)
	index = StaticBasic{
		Security: Security{
			Market: int32(constant.QotMarket_QotMarket_CNSH_Security),
			Code: code,
		},
		Id:            1000905,
		LotSize:       100,
		SecType:       int32(constant.SecurityType_SecurityType_Index),
		Name:          "中证500",
		ListTime:      listTime,
		Delisting:     false,
		ListTimestamp: float64(listTimestamp.Unix()),
	}
	cacheSecurity[code] = index*/
	index, err = genIndexStaticInfo(MARKET_SH, "000905", "中证500", "2007-01-15", 1000905, 100)
	if err == nil && index != nil {
		cacheSecurity[index.Security.Code] = *index
	}

	// 2. 深圳指数
	// 深证成指, sz399001, 1993-01-03
	/*code = GetStockCode(MARKET_SZ, "399001")
	listTime = "1993-01-03"
	listTimestamp, _ = utils.ParseTime(listTime)
	index = StaticBasic{
		Security: Security{
			Market: int32(constant.QotMarket_QotMarket_CNSZ_Security),
			Code: code,
		},
		Id:            2399001,
		LotSize:       100,
		SecType:       int32(constant.SecurityType_SecurityType_Index),
		Name:          "深证成指",
		ListTime:      listTime,
		Delisting:     false,
		ListTimestamp: float64(listTimestamp.Unix()),
	}
	cacheSecurity[code] = index*/
	index, err = genIndexStaticInfo(MARKET_SZ, "399001", "深证成指", "1993-01-03", 2399001, 100)
	if err == nil && index != nil {
		cacheSecurity[index.Security.Code] = *index
	}
	// 创业板指, sz399006, 2010-06-02
	/*code = GetStockCode(MARKET_SZ, "399006")
	listTime = "2010-06-02"
	listTimestamp, _ = utils.ParseTime(listTime)
	index = StaticBasic{
		Security: Security{
			Market: int32(constant.QotMarket_QotMarket_CNSZ_Security),
			Code: code,
		},
		Id:            2399001,
		LotSize:       100,
		SecType:       int32(constant.SecurityType_SecurityType_Index),
		Name:          "创业板指",
		ListTime:      listTime,
		Delisting:     false,
		ListTimestamp: float64(listTimestamp.Unix()),
	}
	cacheSecurity[code] = index*/
	index, err = genIndexStaticInfo(MARKET_SZ, "399006", "创业板指", "2010-06-02", 2399006, 100)
	if err == nil && index != nil {
		cacheSecurity[index.Security.Code] = *index
	}
	// 3. 香港指数 恒生指数时间不准确
	// 4. 加载 上海的个股信息
	for market, name := range MarketName {
		logger.Infof("开始加载 %s 个股静态信息...", name)
		list, err := getStaticBasic(market)
		if err != nil {
			logger.Errorf("上海个股信息加载失败")
		} else {
			for _, item := range list {
				if item.Delisting || item.ListTimestamp == 0.00 {
					// 跳过退市和时间戳为0的个股
					continue
				}
				code := GetStockCode(market, item.Security.Code)
				cacheSecurity[code] = item
			}
		}
		logger.Infof("开始加载 %s 个股静态信息...OK", name)
	}
	logger.Infof("开始初始化市场静态数据...OK")
}

func getStaticBasic(market string) (list []StaticBasic, err error) {
	filename := fmt.Sprintf("%s/%s.json", ResourcesPath, market)
	fileBuf, err := resources.ReadFile(filename)
	if err != nil {
		logger.Debugf("market[%s]: K线数据文件不存在", market)
		return nil, ErrCacheNotExist
	}
	err = json.Unmarshal(fileBuf, &list)
	return
}

func GetBasicInfo(code string) (*StaticBasic, error) {
	info, ok := cacheSecurity[code]
	if !ok || info.Delisting {
		return nil, ErrCodeNotExist
	}
	return &info, nil
}
