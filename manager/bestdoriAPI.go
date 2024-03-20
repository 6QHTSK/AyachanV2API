package manager

import (
	"encoding/json"
	"fmt"
	"github.com/6QHTSK/ayachan/config"
	"github.com/6QHTSK/ayachan/model/bestdoriChart"
	"io"
	"net/http"
	"net/url"
	"time"
)

type chartDataResponse struct {
	Diff  int                           `json:"diff"`
	Chart bestdoriChart.BestdoriV2Chart `json:"chart"`
}

type errorResponse struct {
	ErrorCode int    `json:"err_code"`
	ErrorMsg  string `json:"err_msg"`
}

func httpGet(url string, object *chartDataResponse) (errorCode int, err error) {
	var Client = http.Client{
		Timeout: time.Second * 20, // 20秒超时
	}

	res, err := Client.Get(url)
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer func(Body io.ReadCloser) {
		BErr := Body.Close()
		if BErr != nil {
			err = BErr
		}
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		var ErrMessage errorResponse
		err = json.Unmarshal(body, &ErrMessage)
		if err != nil {
			return http.StatusBadGateway, err
		}
		return res.StatusCode, fmt.Errorf("[%d]%s", ErrMessage.ErrorCode, ErrMessage.ErrorMsg)
	} else {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = json.Unmarshal(body, object)
		if err != nil {
			return http.StatusBadGateway, err
		}
		return http.StatusOK, nil
	}
}

func bestdoriProxyManager(server string, chartID int, diff int) (Chart bestdoriChart.BestdoriV2Chart, serverDiff int, err error) {
	mapDataParam, err := url.Parse(fmt.Sprintf("/post/%s/%d/full?diff=%d", server, chartID, diff))
	mapDataUrl := config.BestdoriAPIUrl.ResolveReference(mapDataParam)
	var request chartDataResponse
	_, err = httpGet(mapDataUrl.String(), &request)
	if err != nil {
		return nil, 0, err
	}
	return request.Chart, request.Diff, nil
}

func GetBandoriChart(chartID int, diff int) (Chart bestdoriChart.BestdoriV2Chart, err error) {
	Chart, _, err = bestdoriProxyManager("bandori", chartID, diff)
	return Chart, err
}

func GetBestdoriChart(chartID int) (Chart bestdoriChart.BestdoriV2Chart, serverDiff int, err error) {
	return bestdoriProxyManager("bestdori", chartID, 0) // Bestdori Server Ignore Diff
}
