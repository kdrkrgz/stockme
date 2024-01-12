package client

import (
	"fmt"
	"net/http"
)

func getStockUrl() map[string]string {
	return map[string]string{
		"ALTINS1": "altins1-darphane-altin-sertifikasi",
		"GOLD":    "gram-altin",
		"USD":     "amerikan-dolari",
		"EUR":     "euro",
		"GBP":     "sterlin",
		"XU100":   "xu100-bist-100",
		"BTC":     "bitcoin",
		"BRENT":   "brent-petrol",
		"THYAO":   "thyao-turk-hava-yollari",
		"EREGL":   "eregl-eregli-demir-celik",
		"YKBNK":   "ykbnk-yapi-ve-kredi-bank",
		"ISCTR":   "isctr-is-bankasi-c",
		"AKBNK":   "akbnk-akbank",
		"TUPRS":   "tuprs-tupras",
		"ASELS":   "asels-aselsan",
		"SAHOL":   "sahol-sabanci-holding",
		"GARAN":   "garan-garanti-bankasi",
		"YEOTK":   "yeotk-yeo-teknoloji-enerji",
		"KCHOL":   "kchol-koc-holding",
		"HALKB":   "halkb-t-halk-bankasi",
		"KBORU":   "kboru-kuzey-boru",
		"EKGYO":   "ekgyo-emlak-konut-gmyo",
		"KRDMD":   "krdmd-kardemir-d",
		"TCELL":   "tcell-turkcell",
		"GUBRF":   "gubrf-gubre-fabrik",
		"PETKM":   "petkm-petkim",
		"KONTR":   "kontr-kontrolmatik-teknoloji",
		"SISE":    "sise-sise-cam",
		"ASTOR":   "astor-astor-enerji",
		"SASA":    "sasa-sasa-polyester",
		"KOZAL":   "kozal-koza-altin",
		"BIMAS":   "bimas-bim-magazalar",
		"TTKOM":   "ttkom-turk-telekom",
		"MIATK":   "miatk-mia-teknoloji",
		"PGSUS":   "pgsus-pegasus",
		"REEDR":   "reedr-reeder-teknoloji",
		"HEKTS":   "hekts-hektas",
		"GESAN":   "gesan-girisim-elektrik-sanayi",
		"FROTO":   "froto-ford-otosan",
		"VAKBN":   "vakbn-vakiflar-bankasi",
		"TOASO":   "toaso-tofas-oto-fab",
		"BRSAN":   "brsan-borusan-boru-sanayi",
		"MGROS":   "mgros-migros-ticaret",
		"AEFES":   "aefes-anadolu-efes",
		"OYAKC":   "oyakc-oyak-cimento",
		"ARCLK":   "arclk-arcelik",
		"FENER":   "fener-fenerbahce-futbol",
		"DOHOL":   "dohol-dogan-holding",
		"KOZAA":   "kozaa-koza-madencilik",
		"ENJSA":   "enjsa-enerjisa-enerji",
		"ALARK":   "alark-alarko-holding",
		"ZOREN":   "zoren-zorlu-enerji",
		"CIMSA":   "cimsa-cimsa",
		"AKSEN":   "aksen-aksa-enerji",
		"OTKAR":   "otkar-otokar",
	}

}

type IStockClient interface {
	GetLastPrice() (*http.Response, error)
}

type StockClient struct {
	StockSymbol string
	Url         string
}

func NewStockClient(stockSymbol, url string) StockClient {
	return StockClient{
		StockSymbol: stockSymbol,
		Url:         url,
	}

}

func checkSymbol(stockSymbolUrlMapping map[string]string, stockSymbol string) bool {
	_, ok := stockSymbolUrlMapping[stockSymbol]
	return ok
}

func (sc *StockClient) GetLastPrice() (*http.Response, error) {
	stockSymbol := sc.StockSymbol
	stockSymbolUrlMapping := getStockUrl()
	if !checkSymbol(stockSymbolUrlMapping, stockSymbol) {
		return nil, fmt.Errorf("invalid symbol, please use supported symbols")
	}
	endpoint := stockSymbolUrlMapping[stockSymbol]
	baseUrl := sc.Url
	uri := fmt.Sprintf("%s/%s", baseUrl, endpoint)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting stock price")
	}
	// defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server returned unsuccessful status code: %v", resp.StatusCode)
	}

	return resp, nil

}
