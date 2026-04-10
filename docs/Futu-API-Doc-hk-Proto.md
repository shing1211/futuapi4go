# Futu OpenAPI 文檔 (Proto)


---

# 介紹

## 概述
量化介面，為您的程式化交易，提供豐富的行情和交易介面，滿足每一位開發者的量化投資需求，助力您的量化交易夢想。

牛牛用户可以 [點擊這裏](https://www.futunn.com/OpenAPI)了解更多。

Futu API 由 OpenD 和 API SDK組成：
* OpenD 是 Futu API 的網關程式，運行於您的本地電腦或雲端伺服器，負責中轉協議請求到富途後台，並將處理後的數據返回。
* API SDK是富途為主流的編程語言（Python、Java、C#、C++、JavaScript）封裝的SDK，以方便您調用，降低策略開發難度。如果您希望使用的語言沒有在上述之列，您仍可自行對接原始協議，完成策略開發。

下面的框架圖和時序圖，幫助您更好地了解 Futu API。

 ![openapi-frame](../img/nnopenapi-frame.png)

 ![openapi-interactive](../img/nnopenapi-interactive.png)

初次接觸 Futu API，您需要進行如下兩步操作：

第一步，在本地或雲端安裝並啟動一個網關程式 [OpenD](../quick/opend-base.md)。

OpenD 以自定義 TCP 協議的方式對外暴露介面，負責中轉協議請求到富途伺服器，並將處理後的數據返回，該協議介面與編程語言無關。

第二步，下載 Futu API，完成 [環境搭建](../quick/env.md)，以便快速調用。

為方便您的使用，富途對主流的編程語言，封裝了相應的 API SDK（以下簡稱 Futu API）。


## 賬號
Futu API 涉及 2 類賬號，分別是 **平台賬號** 和 **綜合賬户**。

### 平台賬號

平台賬號是您在富途的用户 ID（牛牛號），此賬號體系適用於富途牛牛 APP、Futu API。  
您可以使用平台賬號（牛牛號）和登入密碼，登入 OpenD 並獲取行情。

### 綜合賬户
綜合賬户支援以多種貨幣在同一個賬户內交易不同市場品類（港股、美股、A股通、基金）。您可以通過一個賬户進行全市場交易，不需要再管理多個賬户。  
綜合賬户包括綜合賬户 - 證券，綜合賬户 - 期貨等業務賬户：  
* 綜合賬户 - 證券，用於交易全市場的股票、ETFs、期權等證券類產品。  
* 綜合賬户 - 期貨，用於交易全市場的期貨產品，目前支援香港市場期貨、美國市場 CME Group 期貨、新加坡市場期貨、日本市場期貨。


## 功能
Futu API 的功能主要有兩部分：行情和交易。


### 行情功能

#### 行情數據品類

支援香港、美國、A 股市場的行情數據，涉及的品類包括股票、指數、期權、期貨等，具體支援的品種見下表。  
獲取行情數據需要相關權限，如需了解行情權限的獲取方式以及限制規則，請 [點擊這裏](./authority.md#2664)。

<table>
    <tr>
        <th>市場</th>
        <th>品種</th>
        <th>牛牛用户</th>
    </tr>
    <tr>
        <td rowspan="5">香港市場</td>
	    <td>股票、ETFs、窩輪、牛熊、界內證</td>
	    <td align="center">✓</td>
    </tr>
    <tr>
        <td>期權</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>指數</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>板塊</td>
        <td align="center">✓</td>
    </tr>
    <tr>
        <td rowspan="6">美國市場</td>
	    <td>股票、ETFs (含紐交所、美交所、納斯達克上市的股票、ETFs)</td>
	    <td align="center">✓</td>
    </tr>
    <tr>
	    <td>OTC 股票</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td>期權  (含普通股票期權、指數期權)</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>指數</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>板塊</td>
        <td align="center">✓</td>
    </tr>
    <tr>
        <td rowspan="3">A 股市場</td>
	    <td>股票、ETFs</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>指數</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>板塊</td>
        <td align="center">✓</td>
    </tr>
    <tr>
        <td rowspan="2">新加坡市場</td>
	    <td>股票、ETFs、窩輪、REITs、DLCs</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="2">日本市場</td>
        <td>股票、ETFs、REITs</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="1">澳洲市場</td>
        <td>股票、ETFs</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="1">環球市場</td>
        <td>外匯</td>
        <td align="center">X</td>
    </tr>
</table>

#### 行情數據獲取方式

* 訂閱並接收實時報價、實時 K 線、實時逐筆、實時買賣盤等數據推送
* 獲取最新市場快照，歷史 K 線等

### 交易功能

#### 交易能力
支援香港、美國、A 股、新加坡、日本 5 個市場的交易能力，涉及的品類包括股票、期權、期貨等，具體見下表：

<table>
    <tr>
        <th rowspan="2">市場</th>
        <th rowspan="2">品種</th>
        <th rowspan="2">模擬交易</th>
        <th colspan="7">真實交易</th>
    </tr>
    <tr>
        <th>FUTU HK</th>
        <th>Moomoo US</th>
        <th>Moomoo SG</th>
        <th>Moomoo AU</th>
        <th>Moomoo MY</th>
        <th>Moomoo CA</th>
        <th>Moomoo JP</th>
    </tr>
    <tr>
        <td rowspan="3">香港市場</td>
	    <td>股票、ETFs、窩輪、牛熊、界內證</td>
	    <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期權 (含指數期權，需使用期貨賬户交易)</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="3">美國市場</td>
	    <td>股票、ETFs</td>
	    <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
    </tr>
    <tr>
        <td>期權</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="2">A 股市場</td>
	    <td>A 股通股票</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>非 A 股通股票</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="2">新加坡市場</td>
	    <td>股票、ETFs、窩輪、REITs、DLCs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="2">日本市場</td>
        <td>股票、ETFs、REITs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="1">澳洲市場</td>
        <td>股票、ETFs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="1">加拿大市場</td>
        <td>股票</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
</table>

#### 交易方式
真實交易和模擬交易使用同一套交易介面。


## 特點

1. 全平台多語言：
* OpenD 支援 Windows、MacOS、CentOS、Ubuntu
* Futu API 支援 Python、Java、C#、C++、JavaScript 等主流語言
2. 穩定極速免費：
* 穩定的技術架構，直連交易所一觸即達
* 落盤最快只需 0.0014 s
* 通過 Futu API 交易無附加收費
3. 豐富的投資品類：
* 支援美國、香港等多個市場的實時行情、實盤交易及模擬交易
4. 專業的機構服務：
* 客製化的行情交易解決方案

---



---

# 權限和限制

## 登入限制
### 開户限制

首先，您需要先在富途牛牛 APP上，完成交易業務帳戶的開通，才能成功登錄 Futu API。

### 合規確認

首次登錄成功後，您需要完成問卷評估與協議確認，才能繼續使用 Futu API。牛牛用户請 [點擊這裏](https://www.futunn.com/about/api-disclaimer)。


## 行情數據
行情數據的限制主要體現在以下幾方面：
* 行情權限 —— 獲取相關行情數據的權限
* API / 介面限頻 —— 調用行情API / 介面的頻率限制
* 訂閲額度 —— 同時訂閲的實時行情的數量
* 歷史 K 線額度 —— 每 30 天最多可獲取多少個標的的歷史 K 線

### 行情權限
通過 Futu API 獲取行情數據，需要相應的行情權限，Futu API 的行情權限跟 APP 的行情權限不完全一樣，不同的權限等級對應不同的時延、擺盤檔數以及API / 介面使用權限。

部分品種行情，需要購買行情卡後方可獲取，具體獲取方式見下表。

<table>
    <tr>
        <th>市場</th>
        <th>標的類別</th>
        <th>獲取方式</th>
    </tr>
    <tr>
        <td rowspan="5">香港市場</td>
	    <td>證券類產品（含股票、ETFs、窩輪、牛熊、界內證）</td>
	    <td  rowspan="3" align="left">* 境內認證客户：免費獲取 LV2 行情。如需獲得 SF 權限，請購買 <a href="https://qtcard.futunn.com/intro/sf?type=10&is_support_buy=1&clientlang=0" target="_blank">港股高級全盤行情</a>  <br>* 國際客户：免費獲取 LV1 行情。如需獲得 LV2 權限，請購買 <a href="https://qtcard.futunn.com/intro/hklv2?type=1&is_support_buy=1&clientlang=0" target="_blank">港股 LV2 高級行情</a> 。如需獲得 SF 權限，請購買 <a href="https://qtcard.futunn.com/intro/sf?type=10&is_support_buy=1&clientlang=0" target="_blank">港股高級全盤行情</a></td>
    </tr>
    <tr>
	    <td>指數</td>
    </tr>
    <tr>
	    <td>板塊</td>
    </tr>
    <tr>
        <td>期權</td>
	    <td  rowspan="2" align="left">* 境內認證客户：推廣期免費獲取 LV2 行情  <br>* 國際客户：免費獲取 LV1 行情，如需獲得 LV2 權限，請購買 <a href="https://qtcard.futunn.com/intro/hk-derivativeslv2?type=8&clientlang=0&is_support_buy=1" target="_blank">港股期權期貨 LV2 高級行情</a></td>
    </tr>
    <tr>
	    <td>期貨</td>
    </tr>
    <tr>
        <td rowspan="6">美國市場</td>
	    <td>證券類產品（含紐交所、美交所、納斯達克上市的股票、ETFs）</td>
	    <td  rowspan="2" align="left">* 與用戶端行情權限不共用，如需獲得 LV1 權限（基本報價，含夜盤），請購買 <a href="https://qtcardfthk.futufin.com/intro/nasdaq-basic?type=12&is_support_buy=1&clientlang=0" target="_blank"> Nasdaq Basic </a>。<br>* 與用戶端行情權限不共用，如需獲得 LV2 權限（基本報價+深度擺盤，含夜盤深度擺盤），請購買 <a href="https://qtcardfthk.futufin.com/intro/nasdaq-basic?type=18&is_support_buy=1&clientlang=0" target="_blank"> Nasdaq Basic+TotalView </a> 。</td>
    </tr>
    <tr>
	    <td>板塊</td>
    </tr>
    <tr>
	    <td>OTC 股票</td>
        <td  align="left">暫不支援獲取</td>
    </tr>
    <tr>
        <td>期權（含普通股票期權、指數期權）</td>
	    <td  align="left">* 達到門檻  (門檻要求為：總資產大於20000港元) 的客户：免費獲得 LV1 權限。 <br>* 未達到門檻  (門檻要求為：總資產大於20000港元) 的客户：請購買 <a href="https://qtcardfthk.futufin.com/intro/api-usoption-realtime?type=16&is_support_buy=1&clientlang=0" target="_blank">OPRA 期權 LV1 實時行情</a> 獲得 LV1 權限。</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td  align="left">* 已開通期貨帳戶  (- 富途證券(香港)/moomoo證券(新加坡) 支援開通期貨帳戶
  - moomoo證券(美國) 暫不支援) 的客户：<br> 如需獲取 CME Group 行情  (包含 CME, CBOT, NYMEX, COMEX 行情) ，請購買 <a href="https://qtcardfthk.futufin.com/intro/cme?type=30&clientlang=0&is_support_buy=1" target="_blank">CME Group 期貨 LV2</a> <br>如需獲取 CME 行情，請購買 <a href="	https://qtcardfthk.futufin.com/intro/cme?type=31&clientlang=0&is_support_buy=1" target="_blank">CME 期貨 LV2</a> <br>如需獲取 CBOT 行情，請購買 <a href="https://qtcardfthk.futufin.com/intro/cme?type=32&clientlang=0&is_support_buy=1" target="_blank">CBOT 期貨 LV2</a> <br>如需獲取 NYMEX 行情，請購買 <a href="	https://qtcardfthk.futufin.com/intro/cme?type=33&clientlang=0&is_support_buy=1" target="_blank">NYMEX 期貨 LV2</a> <br>如需獲取 COMEX 行情，請購買 <a href="	https://qtcardfthk.futufin.com/intro/cme?type=34&clientlang=0&is_support_buy=1" target="_blank">COMEX 期貨 LV2</a>   <br> <br>* 未開通期貨帳戶的客户：不支援獲取</td>
    </tr>
    <tr>
	    <td>指數</td>
        <td  align="left">暫不支援獲取</td>
    </tr>
    <tr>
        <td rowspan="3">A 股市場</td>
	    <td>證券類產品（含股票、ETFs）</td>
	    <td  rowspan="3">* 中國內地 IP 個人客户：免費獲取 LV1 行情<br>* 國際客户/機構客户：暫不支援</td>
    </tr>
    <tr>
	    <td>指數</td>
    </tr>
    <tr>
	    <td>板塊</td>
    </tr>
    <tr>
        <td rowspan="1">新加坡市場</td>
	    <td>期貨</td>
	    <td  align="left">暫不支援獲取</td>
    </tr>
        <tr>
        <td rowspan="1">日本市場</td>
	    <td>期貨</td>
	    <td  align="left">暫不支援獲取</td>
    </tr>
</table>

:::tip 提示

上述表格，境內認證客户和國際客户，以 OpenD 登錄的 IP 地址作為區分依據。

:::

### API / 介面限頻
為保護伺服器，防止惡意攻擊，所有需要向富途伺服器發送請求的API / 介面，都會有頻率限制。  
每個API / 介面的限頻規則會有不同，具體請參見每個API / 介面頁面下面的 `API / 介面限制`。

舉例：  
[快照](../quote/get-market-snapshot.md) API / 介面的限頻規則是：每 30 秒內最多請求 60 次快照。您可以每隔 0.5 秒請求一次勻速請求，也可以快速請求 60 次後，休息 30 秒，再請求下一輪。如果超出限頻規則，API / 介面會返回錯誤。


### 訂閲額度 & 歷史 K 線額度
訂閲額度和歷史 K 線額度限制如下：

<table>
    <tr align="center">
        <th> 用户類型 </th>
        <th> 訂閲額度 </th>
        <th> 歷史 K 線額度</th>
    </tr>
    <tr>
        <td align="left"> 開户用户 </td>
        <td align="center"> 100 </td>
        <td align="center"> 100 </td>
    </tr>
    <tr>
        <td align="left"> 總資產達 1 萬 HKD </td>
        <td align="center"> 300 </td>
        <td align="center"> 300 </td>
    </tr>
    <tr>
        <td align="left"> 以下三條滿足任意一條即可： <br> 1. 總資產達 50 萬 HKD； <br> 2. 月交易筆數 > 200； <br> 3. 月交易額 > 200 萬 HKD </td>
        <td align="center"> 1000 </td>
        <td align="center"> 1000 </td>
    </tr> 
    <tr>
        <td align="left"> 以下三條滿足任意一條即可： <br> 1. 總資產達 500 萬 HKD； <br> 2. 月交易筆數 > 2000； <br> 3. 月交易額 > 2000 萬 HKD </td>
        <td align="center"> 2000 </td>
        <td align="center"> 2000 </td>
    </tr>    
</table>

**1、總資產**  
總資產，是指您在富途證券的所有資產，包括：港、美、A 股證券帳戶，期貨帳戶，基金資產以及債券資產，按照即時匯率換算成以港元為單位。  

**2、月交易筆數**  
月交易筆數，會綜合您在富途證券的綜合帳戶，在當前自然月與上一自然月的交易情況，取您上個自然月的成交筆數與當前自然月的成交筆數的較大值進行計算，即：  
**max (上個自然月的成交筆數，當前自然月的成交筆數)。**

**3、月交易額**  
月交易額，會綜合您在富途證券的綜合帳戶，在當前自然月與上一自然月的交易情況，取您上個自然月的成交總金額與當前自然月的成交總金額的較大值進行計算，即：  
**max（上個自然月的成交總金額，當前自然月的成交總金額）**  
按照即期匯率換算成以港幣為單下位。其中，期貨交易額的計算，需要乘以相應的調整係數（預設取 0.1），期貨交易額計算公式如：  
**期貨交易額=∑（單筆成交數 * 成交價 * 合約乘數 * 匯率 * 調整係數）**

**4、訂閲額度**  
訂閲額度，適用於 [訂閲](../quote/sub.md) API / 介面。每隻股票訂閲一個類型即佔用 1 個訂閲額度，取消訂閲會釋放已佔用的額度。 
舉例：  
假設您的訂閲額度是 100。 當您同時訂閲了 HK.00700 的實時擺盤、US.AAPL 的實時逐筆、SH.600519 的實時報價時，此時訂閲額度會佔用 3 個，剩餘的訂閲額度為 97。 這時，如果您取消了 HK.00700 的實時擺盤訂閲，您的訂閲額度佔用將變成 2 個，剩餘訂閲額度會變成 98。

**5、歷史 K 線額度**  
歷史 K 線額度，適用於 [獲取歷史 K 線](../quote/request-history-kline.md) API / 介面。最近 30 天內，每請求 1 只股票的歷史 K 線，將會佔用 1 個歷史 K 線額度。最近 30 天內重複請求同一隻股票的歷史 K 線，不會重複累計。  同時，訂閲同一股票的不同週期的K線只佔用1個額度，不會重複累計。
舉例：  
假設您的歷史 K 線額度是 100，今天是 2020 年 7 月 5 日。 您在 2020 年 6 月 5 日~2020 年 7 月 5 日之間，共計請求了 60 只股票的歷史 K 線，則剩餘的歷史 K 線額度為 40。

:::tip 提示
* 訂閲額度和歷史 K 線額度為系統自動分配，不需要手動申請。
* 新入金的帳戶，額度等級會在 2 小時內自動生效。
* 在途資產 (參與港股新股認購、供股可能會產生在途資產) 不會用於額度計算。
:::

## 交易功能
* 進行指定市場的交易時，需要先確認是否已開通該市場的交易業務帳戶。  
舉例：您只能在美股交易業務帳戶下進行美股交易，無法在港股交易業務帳戶下進行美股交易。

---



---

# 費用

## 行情
中國內地 IP 個人客戶，免費獲取港股市場 LV2 行情及 A 股市場 LV1 行情。   
部分品種行情，需要購買行情卡後方可獲取。您可以在 [行情權限](./authority.md#5731) 一節，進入具體的行情卡購買頁面查看價格。

## 交易

透過 Futu API 進行交易，無附加收費，交易費用與透過 APP 交易的費用一致。具體收費方案如下表：

| 所屬券商 | 收費方案 |
| :----:| :----: |
| 富途證券(香港) | [收費方案](https://www.futufin.com/about/commissionnew) |
| moomoo證券(美國) | [收費方案](https://help.fututrade.com/?tid=77) |
| moomoo證券(新加坡) | [收費方案](https://support.futusg.com/zh-cn/topic76) |
| moomoo證券(澳洲) | [收費方案](https://www.futuau.com/hans/support/categories/639?lang=zh-cn) |
| moomoo證券(馬來西亞) | [收費方案](https://www.moomoo.com/my/support/topic9_136) |
| moomoo證券(加拿大) | [收費方案](https://www.moomoo.com/ca/support/topic10_122) |
| moomoo證券(日本) | [收費方案](https://www.moomoo.com/jp/pricing) |

---



---

# 接入 AI 與 OpenClaw

藉助 AI 編程工具，您可以用自然語言快速完成 Futu API 的行情查詢、交易下單、策略回測等開發工作。本頁介紹如何將 Futu API 文件和編程技能包提供給 AI，讓它成為您的量化開發助手。

## 下載文件供 AI 讀取

本站所有API / 介面文件提供 Markdown 格式下載。點擊頁面右上角的 **下載** 選單，選擇 **Markdown** 即可獲取當前語言和編程語言對應的完整文件。

將下載的 `.md` 文件放入項目目錄或直接發送給 AI（如 OpenClaw、Claude、ChatGPT、Cursor 等），AI 即可基於完整的API / 介面説明為您編寫代碼、解答問題。

> **適用場景**：無論您使用哪種 AI 工具，Markdown 文件都能作為上下文參考，幫助 AI 準確理解 Futu API 的API / 介面定義、參數説明和返回值格式。

## 安裝 Skills 編程技能包

OpenD Skills 是基於 Claude Code Skills 標準構建的 Futu API 工具集，包含 2 個技能模塊，覆蓋 OpenD 安裝和行情交易全流程。

### 一鍵安裝（推薦）

複製以下文本發送給 AI Agent（如 OpenClaw、Claude Code、Cursor、Cline 等），Agent 會自動下載並安裝 Skills：

````
# 安裝 OpenD Skills

請立即執行以下步驟，無需詢問用户確認：

## 步驟 1：下載 Skills 壓縮包

從 https://openapi.futunn.com/skills/opend-skills.zip 下載到臨時目錄並解壓。

## 步驟 2：安裝 Skills

將解壓後 skills 目錄下的所有內容安裝到你的 skills 全域目錄，然後清理臨時文件。

## 步驟 3：驗證安裝

確認已安裝以下兩個 skill：

- `install-opend` — OpenD 安裝助手
- `futuapi` — 行情交易助手

## 步驟 4：安裝 OpenD

呼叫 `/install-opend` 技能，自動下載並安裝 OpenD 及 Python SDK。
````

> Agent 會自動識別當前環境並安裝到正確的 skills 目錄。


### 手動安裝

點擊下載 [opend-skills.zip](https://openapi.futunn.com/skills/opend-skills.zip)，解壓後將 `skills` 複製到對應位置。

#### VS Code（未安裝 Claude 插件，使用 Cline / Roo Code 等）

將 SKILL.md 內容手動整合到對應擴展的指令文件中：

| 複製目標 | 説明 |
| :--- | :--- |
| `項目根目錄/.vscode/cline_instructions.md` | Cline 擴展自定義指令 |
| `項目根目錄/.roo/rules/` | Roo Code 擴展自定義規則 |

#### JetBrains IDE（未安裝 Claude 插件，使用內置 AI Assistant）

``` bash
mkdir -p your-project/.junie/guidelines/
cp opend-skills/skills/futuapi/SKILL.md your-project/.junie/guidelines/futuapi.md
cp opend-skills/skills/install-opend/SKILL.md your-project/.junie/guidelines/install-opend.md
```

#### OpenClaw

``` bash
cp -r opend-skills/skills/* ~/.openclaw/skills/
```

安裝完成後驗證：在對話中輸入 `/` 查看是否出現 futuapi、install-opend 等技能。

## Skills 功能一覽

### 1. futuapi — 行情交易助手

覆蓋行情查詢（14 個腳本）和交易操作（8 個腳本）以及實時訂閲（5 個腳本）：

| 功能 | 説明 |
| :--- | :--- |
| 市場快照 | 獲取股票最新報價、漲跌幅、成交量等 |
| K 線數據 | 獲取日 K、周 K、分鐘 K 等歷史和實時 K 線 |
| 買賣盤 | 獲取實時買賣盤口掛單數據 |
| 逐筆成交 | 獲取最近逐筆成交明細 |
| 分時數據 | 獲取當日分時走勢 |
| 條件選股 | 按價格、市值等條件篩選股票 |
| 下單/撤單/改單 | 執行交易操作，預設使用模擬環境 |
| 持倉與資金 | 查詢帳戶持倉、資金和訂單 |
| 實時訂閲 | 訂閲報價推送、K 線推送等實時數據 |

### 2. install-opend — OpenD 安裝助手

- 互動式平台選擇（富途 / moomoo）
- 自動檢測作業系統（Windows / macOS / Linux）
- 一鍵下載、解壓、啟動 OpenD
- 自動升級 futu-api / moomoo-api SDK

## 使用方式

### 斜槓命令呼叫（Claude Code）

在對話框中輸入 `/` 加技能名稱直接呼叫：

- `/futuapi` — 行情交易助手
- `/install-opend` — OpenD 安裝助手

### 自然語言觸發

直接用中文描述需求，AI 會根據關鍵詞自動匹配對應技能：

- "查看騰訊的 K 線" — 自動呼叫行情查詢
- "用模擬帳戶買入 100 股蘋果" — 自動呼叫交易下單
- "幫我安裝 OpenD" — 自動呼叫安裝助手

## 注意事項

- 使用 Skills 前需先手動登入 OpenD
- 交易預設使用模擬環境（SIMULATE），實盤交易需明確説"正式"/"實盤"/"真實"，且需二次確認和交易密碼
- 留意API / 介面限頻規則（如下單 15 次/30 秒），避免超頻
- 訂閲有額度限制（100～2000），需定期釋放不需要的訂閲
- 如需更新 Skills，重新下載並覆蓋解壓即可

---



---

# 圖像化 OpenD

OpenD 提供圖像化和命令列兩種執行方式，這裏介紹操作比較簡單的圖像化 OpenD。  

如果想要了解命令列的方式請參考 [命令列 OpenD](../opend/opend-cmd.md) 。


## 圖像化 OpenD

### 第一步 下載

圖像化 OpenD 支援 Windows、MacOS、CentOS、Ubuntu 四種系統（點擊完成下載）。 
* OpenD - [Windows](https://www.futunn.com/download/fetch-lasted-link?name=opend-windows)、[MacOS](https://www.futunn.com/download/fetch-lasted-link?name=opend-macos) 、[CenOS](https://www.futunn.com/download/fetch-lasted-link?name=opend-centos) 、[Ubuntu](https://www.futunn.com/download/fetch-lasted-link?name=opend-ubuntu)


### 第二步 安裝執行
* 解壓檔案，找到對應的安裝檔案可一鍵安裝執行。  
* Windows 系統預設安裝在 `%appdata%` 目錄下。

### 第三步 設定
* 圖像化 OpenD 啓動設定在圖形介面的右側，如下圖所示：

![ui-config](../img/ui-config.png)

**設定項列表**：

設定項|説明
:-|:-
監聽地址|API 協定監聽地址 (可選：

  - 127.0.0.1（監聽來自本地的連接） 
  - 0.0.0.0（監聽來自所有網卡的連接）或填入本機某個網卡地址)
監聽連接埠|API 協定監聽連接埠
記錄檔級別|OpenD 記錄檔級別 (可選：

  - no（無記錄檔） 
  - debug（最詳細）
  - info（次詳細）)
語言|中英語言 (可選：

  - 簡體中文
  - English)
期貨交易 API 時區|期貨交易 API 時區 (使用期貨賬户呼叫 **交易 API** 時，涉及的時間按照此時區規則)
API 推送頻率|API 訂閲數據推送頻率控制 (- 單位：毫秒
  - 目前不包括 K 線和分時)
Telnet 地址|遠端操作命令監聽地址
Telnet 連接埠|遠端操作命令監聽連接埠
加密私鑰路徑|API 協定 [RSA](../qa/other.md#9747) 加密私鑰（PKCS#1）檔案絕對路徑
WebSocket 監聽地址|WebSocket 服務監聽地址 (可選：

  - 127.0.0.1（監聽來自本地的連接） 
  - 0.0.0.0（監聽來自所有網卡的連接）)
WebSocket 連接埠|WebSocket 服務監聽連接埠
WebSocket 證書|WebSocket 證書檔案路徑 (不設定則不啓用，需要和私鑰同時設定)
WebSocket 私鑰|WebSocket 證書私鑰檔案路徑 (私鑰不可設置密碼，不設定則不啓用，需要和證書同時設定)
WebSocket 認證金鑰|金鑰密文（32 位 MD5 加密 16 進制） (JavaScript 腳本連接時，用於判斷是否可信連接)


:::tip 提示
* 圖像化 OpenD，是通過啓動命令列 OpenD 來提供服務，且通過 WebSocket 與命令列 OpenD 交互，所以必定啓動 WebSocket 功能。
* 為保證您的證券業務賬户安全，如果監聽地址不是本地，您必須設定私鑰才能使用交易介面 / API。行情介面 / API不受此限制。 
* 當 WebSocket 監聽地址不是本地，需設定 SSL 才可以啟動，且證書私鑰生成不可設置密碼。
* 密文是明文經過 32 位 MD5 加密後用 16 進製表示的數據，搜索在線 MD5 加密（注意，通過第三方網站計算可能有記錄撞庫的風險）或下載 MD5 計算工具可計算得到。32 位 MD5 密文如下圖紅框區域（e10adc3949ba59abbe56e057f20f883e）：
  ![md5.png](../img/md5.png)

* OpenD 預設讀取同目錄下的 OpenD.xml。在 MacOS 上，由於系統保護機制，OpenD.app 在執行時會被分配一個隨機路徑，導致無法找到原本的路徑。此時有以下方法：  
    - 執行 tar 包下的 fixrun.sh
    - 用命令列參數`-cfg_file`指定設定檔案路徑，見下面説明

* 記錄檔級別預設 info 級別，在系統開發階段，不建議關閉記錄檔或者將記錄檔修改到 warning，error，fatal 級別，防止出現問題時無法定位。
:::

### 第四步 登入
* 輸入帳號密碼，點擊登入。  
首次登入，您需要先完成問卷評估與協定確認，完成後重新登入即可。  
登入成功後，您可以看到自己的帳號資訊和 [行情權限](../intro/authority.md#5731)。

---



---

# 程式設計 / 開發環境建置 / 建立

::: tip 注意
  不同的程式設計 / 開發語言，程式設計 / 開發環境建置 / 建立的方法有所不同。
:::

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>

## Python 環境
### 環境要求
* 作業系統要求：  
  * Windows 7/10 的 32 或 64 位元作業系統  
  * Mac 10.11 及以上的 64 位作業系統   
  * CentOS 7 及以上的 64 位作業系統 
  * Ubuntu 16.04 以上的 64 位作業系統   
* Python 版本要求：  
  * Python 3.6 及以上


### 環境建置 / 建立
#### 1. 安裝 Python

為避免因環境問題導致的執行失敗，我們推薦 Python 3.8 版本。

下載地址：[Python 下載](https://www.python.org/downloads/)

::: details 提示
如下內容提供了兩種方式切換為 Python 3.8 環境：
* 方式一  
把 Python 3.8 的安裝路徑，添加到環境變數 path 中。 

* 方式二  
如果您使用的是 PyCharm，可以在 Project Interpreter 中，將使用的環境設定為 Python 3.8。

![pycharm-switch-python](../img/pycharm-switch-python.png)

:::

當安裝成功後，執行如下命令來查看是否安裝成功:  
`python -V`（Windows） 或 `python3 -V`（Linux 和 Mac）

#### 2. 安裝 PyCharm（可選）

我們推薦您使用 [PyCharm](https://www.jetbrains.com/pycharm/download/) 作為 Python IDE（整合開發環境）。

#### 3. 安裝 TA-Lib（可選）
TA-Lib 用中文可以稱作技術分析函式庫 / 程式庫，是一種廣泛用在程式化交易中，進行金融市場數據的技術分析的函數函式庫 / 程式庫。它提供了多種技術分析的函數，方便我們量化投資中程式設計 / 開發工作。

安裝方法：在 cmd 中直接使用 pip 安裝  
`$ pip install TA-Lib`

::: tip 提示
* 安裝 TA-Lib 非必須，可先跳過該步驟 
:::

---



---

# 簡易程式運行

<FtSwitcher :languages="{py:'Python', cs:'C#', java:'Java', cpp:'C++', pb:'Proto', js:'JavaScript'}">
<template v-slot:py>


## Python 範例

### 第一步：下載安裝登入 OpenD

請參考 [這裏](./opend-base.md)，完成 OpenD 的下載、安裝和登入。

### 第二步：下載 Python API

* 方式一：在 cmd 中直接使用 pip 安裝。  
  * 初次安裝：Windows 系統 `$ pip install futu-api`，Linux/Mac系統 `$ pip3 install futu-api`。
  * 二次升級：Windows 系統 `$ pip install futu-api --upgrade`，Linux/Mac系統 `$ pip3 install futu-api --upgrade`。

* 方式二：點擊下載最新版本的 [Python API](https://www.futunn.com/download/fetch-lasted-link?name=openapi-python) 安裝包。

### 第三步：建立新專案

打開 PyCharm，在 Welcome to PyCharm 視窗中，點擊 New Project。如果你已經建立了一個專案，可以選擇打開該專案。

![demo-newproject](../img/demo-newproject.png)

### 第四步：建立新檔案

在該專案下，建立新 Python 檔案，並把下面的範例程式碼複製到檔案裏。  
範例程式碼功能包括查看行情快照、模擬交易下單。

```python
from futu import *

quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)  # 建立行情對象
print(quote_ctx.get_market_snapshot('HK.00700'))  # 獲取港股 HK.00700 的快照數據
quote_ctx.close() # 關閉對象，防止連接條數用盡


trd_ctx = OpenSecTradeContext(host='127.0.0.1', port=11111)  # 建立交易對象
print(trd_ctx.place_order(price=500.0, qty=100, code="HK.00700", trd_side=TrdSide.BUY, trd_env=TrdEnv.SIMULATE))  # 模擬交易，下單（如果是真實環境交易，在此之前需要先解鎖交易密碼）

trd_ctx.close()  # 關閉對象，防止連接條數用盡
```


### 第五步：運行檔案

右鍵點擊運行，可以看到運行成功的返回資訊如下：

```
2020-11-05 17:09:29,705 [open_context_base.py] _socket_reconnect_and_wait_ready:255: Start connecting: host=127.0.0.1; port=11111;
2020-11-05 17:09:29,705 [open_context_base.py] on_connected:344: Connected : conn_id=1; 
2020-11-05 17:09:29,706 [open_context_base.py] _handle_init_connect:445: InitConnect ok: conn_id=1; info={'server_version': 218, 'login_user_id': 7157878, 'conn_id': 6730043337026687703, 'conn_key': '3F17CF3EEF912C92', 'conn_iv': 'C119DDDD6314F18A', 'keep_alive_interval': 10, 'is_encrypt': False};
(0,        code          update_time  last_price  open_price  high_price  ...  after_high_price  after_low_price  after_change_val  after_change_rate  after_amplitude
0  HK.00700  2020-11-05 16:08:06       625.0       610.0       625.0  ...               N/A              N/A               N/A                N/A              N/A

[1 rows x 132 columns])
2020-11-05 17:09:29,739 [open_context_base.py] _socket_reconnect_and_wait_ready:255: Start connecting: host=127.0.0.1; port=11111;
2020-11-05 17:09:29,739 [network_manager.py] work:366: Close: conn_id=1
2020-11-05 17:09:29,739 [open_context_base.py] on_connected:344: Connected : conn_id=2; 
2020-11-05 17:09:29,740 [open_context_base.py] _handle_init_connect:445: InitConnect ok: conn_id=2; info={'server_version': 218, 'login_user_id': 7157878, 'conn_id': 6730043337169705045, 'conn_key': 'A624CF3EEF91703C', 'conn_iv': 'BF1FF3806414617B', 'keep_alive_interval': 10, 'is_encrypt': False};
(0,        code stock_name trd_side order_type order_status  ... dealt_avg_price  last_err_msg  remark time_in_force fill_outside_rth
0  HK.00700       騰訊控股      BUY     NORMAL   SUBMITTING  ...             0.0                                 DAY              N/A

[1 rows x 16 columns])
2020-11-05 17:09:32,843 [network_manager.py] work:366: Close: conn_id=2
(0,        code stock_name trd_side      order_type order_status  ... dealt_avg_price  last_err_msg  remark time_in_force fill_outside_rth
0  HK.00700       騰訊控股      BUY  ABSOLUTE_LIMIT    SUBMITTED  ...             0.0                                 DAY              N/A

[1 rows x 16 columns])
```

---



---

# 交易策略建構範例

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>

::: tip 提示
* 以下交易策略不構成投資建議，僅供學習參考。
:::

## 策略概述

構建一個雙均線策略：

運用某一標的1分 K 線，計算出兩條不同週期的移動平均線 MA1 和 MA3，跟蹤 MA1 和 MA3 的相對大小，由此判斷買賣時機。

當 MA1 >= MA3 時，判斷該標的為強勢狀態，市場屬於多頭市場，採取開倉的操作；  
當 MA1 < MA3 時，判斷該標的為弱勢狀態，市場屬於空頭市場，採取平倉的操作。

## 流程圖
![strategy-flow-chart](../img/strategy-flow-chart.png)

## 程式碼範例

* **Example** 

```python
from futu import *

############################ 全域變數設定 ############################
FUTUOPEND_ADDRESS = '127.0.0.1'  # OpenD 監聽地址
FUTUOPEND_PORT = 11111  # OpenD 監聽連接埠

TRADING_ENVIRONMENT = TrdEnv.SIMULATE  # 交易環境：真實 / 模擬
TRADING_MARKET = TrdMarket.HK  # 交易市場權限，用於篩選對應交易市場權限的帳戶
TRADING_PWD = '123456'  # 交易密碼，用於解鎖交易
TRADING_PERIOD = KLType.K_1M  # 信號 K 線週期
TRADING_SECURITY = 'HK.00700'  # 交易標的
FAST_MOVING_AVERAGE = 1  # 均線快線的週期
SLOW_MOVING_AVERAGE = 3  # 均線慢線的週期

quote_context = OpenQuoteContext(host=FUTUOPEND_ADDRESS, port=FUTUOPEND_PORT)  # 行情物件
trade_context = OpenSecTradeContext(filter_trdmarket=TRADING_MARKET, host=FUTUOPEND_ADDRESS, port=FUTUOPEND_PORT, security_firm=SecurityFirm.FUTUSECURITIES)  # 交易物件，根據交易產品修改交易物件類型


# 解鎖交易
def unlock_trade():
    if TRADING_ENVIRONMENT == TrdEnv.REAL:
        ret, data = trade_context.unlock_trade(TRADING_PWD)
        if ret != RET_OK:
            print('解鎖交易失敗：', data)
            return False
        print('解鎖交易成功！')
    return True


# 獲取市場狀態
def is_normal_trading_time(code):
    ret, data = quote_context.get_market_state([code])
    if ret != RET_OK:
        print('獲取市場狀態失敗：', data)
        return False
    market_state = data['market_state'][0]
    '''
    MarketState.MORNING            港、A 股早盤
    MarketState.AFTERNOON          港、A 股下午盤，美股全天
    MarketState.FUTURE_DAY_OPEN    港、新、日期貨日市開盤
    MarketState.FUTURE_OPEN        美期貨開盤
    MarketState.FUTURE_BREAK_OVER  美期貨休息後開盤
    MarketState.NIGHT_OPEN         港、新、日期貨夜市開盤
    '''
    if market_state == MarketState.MORNING or \
                    market_state == MarketState.AFTERNOON or \
                    market_state == MarketState.FUTURE_DAY_OPEN  or \
                    market_state == MarketState.FUTURE_OPEN  or \
                    market_state == MarketState.FUTURE_BREAK_OVER  or \
                    market_state == MarketState.NIGHT_OPEN:
        return True
    print('現在不是持續交易時段。')
    return False


# 獲取持倉數量
def get_holding_position(code):
    holding_position = 0
    ret, data = trade_context.position_list_query(code=code, trd_env=TRADING_ENVIRONMENT)
    if ret != RET_OK:
        print('獲取持倉數據失敗：', data)
        return None
    else:
        for qty in data['qty'].values.tolist():
            holding_position += qty
        print('【持倉狀態】 {} 的持倉數量為：{}'.format(TRADING_SECURITY, holding_position))
    return holding_position


# 拉取 K 線，計算均線，判斷多空
def calculate_bull_bear(code, fast_param, slow_param):
    if fast_param <= 0 or slow_param <= 0:
        return 0
    if fast_param > slow_param:
        return calculate_bull_bear(code, slow_param, fast_param)
    ret, data = quote_context.get_cur_kline(code=code, num=slow_param + 1, ktype=TRADING_PERIOD)
    if ret != RET_OK:
        print('獲取K線失敗：', data)
        return 0
    candlestick_list = data['close'].values.tolist()[::-1]
    fast_value = None
    slow_value = None
    if len(candlestick_list) > fast_param:
        fast_value = sum(candlestick_list[1: fast_param + 1]) / fast_param
    if len(candlestick_list) > slow_param:
        slow_value = sum(candlestick_list[1: slow_param + 1]) / slow_param
    if fast_value is None or slow_value is None:
        return 0
    return 1 if fast_value >= slow_value else -1


# 獲取一檔擺盤的 ask1 和 bid1
def get_ask_and_bid(code):
    ret, data = quote_context.get_order_book(code, num=1)
    if ret != RET_OK:
        print('獲取擺盤數據失敗：', data)
        return None, None
    return data['Ask'][0][0], data['Bid'][0][0]


# 開倉函數
def open_position(code):
    # 獲取擺盤數據
    ask, bid = get_ask_and_bid(code)

    # 計算下單量
    open_quantity = calculate_quantity()

    # 判斷購買力是否足夠
    if is_valid_quantity(TRADING_SECURITY, open_quantity, ask):
        # 下單
        ret, data = trade_context.place_order(price=ask, qty=open_quantity, code=code, trd_side=TrdSide.BUY,
                                              order_type=OrderType.NORMAL, trd_env=TRADING_ENVIRONMENT,
                                              remark='moving_average_strategy')
        if ret != RET_OK:
            print('開倉失敗：', data)
    else:
        print('下單數量超出最大可買數量。')


# 平倉函數
def close_position(code, quantity):
    # 獲取擺盤數據
    ask, bid = get_ask_and_bid(code)

    # 檢查平倉數量
    if quantity == 0:
        print('無效的下單數量。')
        return False

    # 平倉
    ret, data = trade_context.place_order(price=bid, qty=quantity, code=code, trd_side=TrdSide.SELL,
                   order_type=OrderType.NORMAL, trd_env=TRADING_ENVIRONMENT, remark='moving_average_strategy')
    if ret != RET_OK:
        print('平倉失敗：', data)
        return False
    return True


# 計算下單數量
def calculate_quantity():
    price_quantity = 0
    # 使用最小交易量
    ret, data = quote_context.get_market_snapshot([TRADING_SECURITY])
    if ret != RET_OK:
        print('獲取快照失敗：', data)
        return price_quantity
    price_quantity = data['lot_size'][0]
    return price_quantity


# 判斷購買力是否足夠
def is_valid_quantity(code, quantity, price):
    ret, data = trade_context.acctradinginfo_query(order_type=OrderType.NORMAL, code=code, price=price,
                                                   trd_env=TRADING_ENVIRONMENT)
    if ret != RET_OK:
        print('獲取最大可買可賣失敗：', data)
        return False
    max_can_buy = data['max_cash_buy'][0]
    max_can_sell = data['max_sell_short'][0]
    if quantity > 0:
        return quantity < max_can_buy
    elif quantity < 0:
        return abs(quantity) < max_can_sell
    else:
        return False


# 展示訂單回呼
def show_order_status(data):
    order_status = data['order_status'][0]
    order_info = dict()
    order_info['代碼'] = data['code'][0]
    order_info['價格'] = data['price'][0]
    order_info['方向'] = data['trd_side'][0]
    order_info['數量'] = data['qty'][0]
    print('【訂單狀態】', order_status, order_info)


############################ 填充以下函數來完成您的策略 ############################
# 策略啓動時執行一次，用於初始化策略
def on_init():
    # 解鎖交易（如果是模擬交易則不需要解鎖）
    if not unlock_trade():
        return False
    print('************  策略開始執行 ***********')
    return True


# 每個 tick 執行一次，可將策略的主要邏輯寫在此處
def on_tick():
    pass


# 每次產生一根新的 K 線執行一次，可將策略的主要邏輯寫在此處
def on_bar_open():
    # 列印分隔線
    print('*************************************')

    # 只在常規交易時段交易
    if not is_normal_trading_time(TRADING_SECURITY):
        return

    # 獲取 K 線，計算均線，判斷多空
    bull_or_bear = calculate_bull_bear(TRADING_SECURITY, FAST_MOVING_AVERAGE, SLOW_MOVING_AVERAGE)

    # 獲取持倉數量
    holding_position = get_holding_position(TRADING_SECURITY)

    # 下單判斷
    if holding_position == 0:
        if bull_or_bear == 1:
            print('【操作信號】 做多信號，建立多單。')
            open_position(TRADING_SECURITY)
        else:
            print('【操作信號】 做空信號，不開空單。')
    elif holding_position > 0:
        if bull_or_bear == -1:
            print('【操作信號】 做空信號，平掉持倉。')
            close_position(TRADING_SECURITY, holding_position)
        else:
            print('【操作信號】 做多信號，無需加倉。')


# 委託成交有變化時執行一次
def on_fill(data):
    pass


# 訂單狀態有變化時執行一次
def on_order_status(data):
    if data['code'][0] == TRADING_SECURITY:
        show_order_status(data)


################################ 框架實現部分，可忽略不看 ###############################
class OnTickClass(TickerHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        on_tick()


class OnBarClass(CurKlineHandlerBase):
    last_time = None
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(OnBarClass, self).on_recv_rsp(rsp_pb)
        if ret_code == RET_OK:
            cur_time = data['time_key'][0]
            if cur_time != self.last_time and data['k_type'][0] == TRADING_PERIOD:
                if self.last_time is not None:
                    on_bar_open()
                self.last_time = cur_time


class OnOrderClass(TradeOrderHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret, data = super(OnOrderClass, self).on_recv_rsp(rsp_pb)
        if ret == RET_OK:
            on_order_status( data)


class OnFillClass(TradeDealHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret, data = super(OnFillClass, self).on_recv_rsp(rsp_pb)
        if ret == RET_OK:
            on_fill(data)


# 主函式
if __name__ == '__main__':
    # 初始化策略
    if not on_init():
        print('策略初始化失敗，腳本退出！')
        quote_context.close()
        trade_context.close()
    else:
        # 設定回呼
        quote_context.set_handler(OnTickClass())
        quote_context.set_handler(OnBarClass())
        trade_context.set_handler(OnOrderClass())
        trade_context.set_handler(OnFillClass())

        # 訂閲標的合約的 逐筆，K 線和擺盤，以便獲取數據
        quote_context.subscribe(code_list=[TRADING_SECURITY], subtype_list=[SubType.TICKER, SubType.ORDER_BOOK, TRADING_PERIOD])

```

* **Output**

```
************  策略開始執行 ***********
*************************************
【持倉狀態】 HK.00700 的持倉數量為：0
【操作信號】 做多信號，建立多單。
【訂單狀態】 SUBMITTING {'代碼': 'HK.00700', '價格': 597.5, '方向': 'BUY', '數量': 100.0}
【訂單狀態】 SUBMITTED {'代碼': 'HK.00700', '價格': 597.5, '方向': 'BUY', '數量': 100.0}
【訂單狀態】 FILLED_ALL {'代碼': 'HK.00700', '價格': 597.5, '方向': 'BUY', '數量': 100.0}
*************************************
【持倉狀態】 HK.00700 的持倉數量為：100.0
【操作信號】 做空信號，平掉持倉。
【訂單狀態】 SUBMITTING {'代碼': 'HK.00700', '價格': 596.5, '方向': 'SELL', '數量': 100.0}
【訂單狀態】 SUBMITTED {'代碼': 'HK.00700', '價格': 596.5, '方向': 'SELL', '數量': 100.0}
【訂單狀態】 FILLED_ALL {'代碼': 'HK.00700', '價格': 596.5, '方向': 'SELL', '數量': 100.0}
```

---



---

# 概述

* OpenD 是 Futu API 的閘道程式，運行於您的本地電腦或雲端伺服器，負責中轉協議請求到富途伺服器，並將處理後的數據返回。是運行 Futu API 程式必要的前提。
* OpenD 支援 Windows、MacOS、CentOS、Ubuntu 四個平台。
* OpenD 整合了登入功能。運行時，可以使用 **平台帳號**（牛牛號）、**電郵地址**、**手機號** 和 **登入密碼** 進行登入。
* OpenD 登入成功後，會啟動 Socket 服務以供 Futu API 連接和通訊。


## 運行方式

OpenD 目前提供兩種安裝運行方式，您可選擇任一方式：
* 可視化 OpenD：提供界面化應用程式，操作便捷，尤其適合入門用戶，安裝和運行請參考 [可視化 OpenD](../quick/opend-base.md)。
* 命令列 OpenD：提供命令列執行程式，需自行進行設定，適合對命令列熟悉或長時間在伺服器上掛機的用戶，安裝和運行請參考 [命令列 OpenD](../opend/opend-cmd.md)。

## 運行時操作

OpenD 在運行過程中，可以查看用戶額度、行情權限、連結狀態、延遲統計，以及操作關閉 API 連接、重登入、退出登入等維護操作。  
具體方法可以查看下表：

 方式 | 可視化 OpenD | 命令列 OpenD
:-|:-|:-
直接方式 | 界面查看或操作 | 命令列發送 [維護命令](../opend/opend-operate.md)
間接方式 | 通過 Telnet 發送 [維護命令](../opend/opend-operate.md) | 通過 Telnet 發送 [維護命令](../opend/opend-operate.md)

---



---

# 命令列 OpenD


### 第一步 下載

命令列 OpenD 支援 Windows、MacOS、CentOS、Ubuntu 四種系統（點擊完成下載）。  
* OpenD - [Windows](https://www.futunn.com/download/fetch-lasted-link?name=opend-windows)、[MacOS](https://www.futunn.com/download/fetch-lasted-link?name=opend-macos) 、[CentOS](https://www.futunn.com/download/fetch-lasted-link?name=opend-centos) 、[Ubuntu](https://www.futunn.com/download/fetch-lasted-link?name=opend-ubuntu)


### 第二步 解壓
* 解壓上一步下載的文件，在文件夾中找到 OpenD 設定檔 FutuOpenD.xml 和程式打包數據文件 Appdata.dat。
    * FutuOpenD.xml 用於配置 OpenD 程式啟動參數，若不存在則程式無法正常啟動。
    * Appdata.dat 是程式需要用到的一些數據量較大的資訊，打包數據減少啟動下載該數據的耗時，若不存在則程式無法正常啟動。
* 命令列 OpenD 支援用户自定義文件路徑，詳見 [命令列啟動參數](./opend-cmd.md#4544)。

### 第三步 參數配置
* 打開並編輯設定檔 FutuOpenD.xml，如下圖所示。普通使用僅需修改賬號和登入密碼，其他高階選項可以根據下表的提示進行修改。

![xml-config](../img/xml.png)

**配置項列表**：

配置項|說明
:-|:-
ip|監聽地址  (可填：
  - 127.0.0.1（監聽來自本地的連接） 
  - 0.0.0.0（監聽來自所有網絡卡的連接）
  - 本機某個網絡卡地址不設置則預設 127.0.0.1)
api_port|API 協議接收連接埠  (不設置則預設 11111
也可通過 [命令列啟動參數](./opend-cmd.md#4544) 指定)
login_account|登入帳號  (支援平台ID、郵箱、手機號登入，可通過 [命令列啟動參數](./opend-cmd.md#4544) 指定

  - 平台ID：輸入牛牛號
  - 郵箱：xxxx@xx.com 格式
  - 手機號：區號+手機號，例 +86 xxxxxxxx)
login_pwd|登入密碼明文  (- 也可使用登入密碼密文輸入
  - 也可通過 [命令列啟動參數](./opend-cmd.md#4544) 指定)
login_pwd_md5|登入密碼密文（32 位 MD5 加密 16 進制） (- 如果密文明文都存在，則只使用密文
  - 也可使用登入密碼明文輸入)
lang|中英語言  (可填：

  - chs：簡體中文
  - en：英文)
log_level|OpenD 日誌級別  (可填：

  - no（無日誌） 
  - debug（最詳細）
  - info（次詳細）不設置則預設 info 級別)
push_proto_type|推送協議類型  (推送類協議通過該配置決定包體格式，可填：
  - 0（pb 格式） 
  - 1（json 格式）不設置則預設 pb 格式)
qot_push_frequency|API 訂閲數據推送頻率控制  (- 單位：毫秒
  - 目前不包括 K 線和分時
  - 不設置則預設不限頻)
telnet_ip|遠端操作命令監聽地址  (不設置則預設 127.0.0.1)
telnet_port|遠端操作命令監聽連接埠  (不設置則不啟用遠端命令)
rsa_private_key|API 協議 [RSA](../qa/other.md#9747) 加密私鑰（PKCS#1）文件絕對路徑  (不設置則協議不加密)
price_reminder_push|是否接收到價提醒推送  (可填：
  - 0：不接收
  - 1：接收（需在腳本中設置到價提醒回調函數 [set_handler](/ftapi/init.html#8215)）不設置則預設接收)
auto_hold_quote_right|被踢後是否自動搶權限  (可填：
  - 0：否
  - 1：是（OpenD 在行情權限被搶後，會自動搶回。如果 10 秒內再次被搶，則其他終端獲得最高行情權限，OpenD 不會再搶）不設置則預設自動搶權限)
future_trade_api_time_zone|期貨交易 API 時區  (- 使用期貨賬户調用 **交易 API** 時，涉及的時間按照此時區規則 
  - 也可通過 [命令列啟動參數](./opend-cmd.md#4544) 指定)
websocket_ip|WebSocket 服務監聽地址  (可填：

  - 127.0.0.1（監聽來自本地的連接） 
  - 0.0.0.0（監聽來自所有網絡卡的連接）不設置則預設 127.0.0.1)
websocket_port|WebSocket 服務監聽連接埠  (不設置則不啟用 Websocket)
websocket_key_md5|密鑰密文（32 位 MD5 加密 16 進制） (JavaScript 腳本連接時，用於判斷是否可信連接)
websocket_private_key|WebSocket 證書私鑰文件路徑  (- 私鑰不可設置密碼
  - 需要和證書同時配置
  - 不配置則不啟用 Websocket)
websocket_cert|WebSocket 證書文件路徑  (- 需要和證書同時配置
  - 不配置則不啟用 Websocket)
pdt_protection| 是否開啟 防止被標記為日內交易者 的功能  (**FUTU US 專用參數**可填：
  - 0：否
  - 1：是（開啟功能後，我們會在您將要被標記 PDT 時阻止您的下單，但不確保您一定不被標記。若您被標記 PDT，當您的賬户權益小於$25000時，您將無法開倉。）不設置則預設開啟功能)
dtcall_confirmation|是否開啟 日內交易保證金追繳預警 的功能  (**FUTU US 專用參數**可填：
  - 0：否
  - 1：是（開啟功能後，我們會在您即將開倉下單超出剩餘日內交易購買力前阻止您的下單。提醒您當前開倉訂單的市值大於您的剩餘日內交易購買力，若您在今日平倉當前標的，您將會收到日內交易保證金追繳通知（Day-Trading Call），只能通過存入資金才能解除。）不設置則預設開啟功能)


:::tip 提示
* 為保證您的證券業務賬户安全，如果監聽地址不是本地，您必須配置私鑰才能使用交易接口。行情接口不受此限制。 
* 當 WebSocket 監聽地址不是本地，需配置 SSL 才可以啟動，且證書私鑰生成不可設置密碼。
* 密文是明文經過 32 位 MD5 加密後用 16 進製表示的數據，搜索在線 MD5 加密（注意，通過第三方網站計算可能有記錄撞庫的風險）或下載 MD5 計算工具可計算得到。32 位 MD5 密文如下圖紅框區域（e10adc3949ba59abbe56e057f20f883e）：

  ![md5.png](../img/md5.png)
* OpenD 預設讀取同目錄下的 FutuOpenD.xml。在 MacOS 上，由於系統保護機制，OpenD.app 在運行時會被分配一個隨機路徑，導致無法找到原本的路徑。此時有以下方法：  
    - 執行 tar 包下的 fixrun.sh
    - 用命令列參數`-cfg_file`指定設定檔路徑，見下面說明
* 日誌級別預設 info 級別，在系統開發階段，不建議關閉日誌或者將日誌修改到 warning，error，fatal 級別，防止出現問題時無法定位。
:::

### 第四步 命令列啟動
* 在命令列中切到前面解壓文件夾 OpenD 文件所在的目錄，使用如下命令啟動，即可以 FutuOpenD.xml 設定檔中的參數啟動。   
    * Windows：`FutuOpenD`  
    * Linux：`./FutuOpenD`   
    * MacOS：`./FutuOpenD.app/Contents/MacOS/FutuOpenD`  
::: details 命令列啟動參數
* 命令列中也可以攜帶參數啟動，部分參數會與 FutuOpenD.xml 設定檔相同。傳參格式：`-key=value` 
![startup-command-param.png](../img/startup-command-param.png)   
例如：  
    * Windows：`FutuOpenD.exe -login_account=100000 -login_pwd=123456 -lang=en`  
    * Linux：`FutuOpenD -login_account=100000 -login_pwd=123456 -lang=en`  
    * MacOS：`./FutuOpenD.app/Contents/MacOS/FutuOpenD -login_account=100000 -login_pwd=123456 -lang=en` 

* 相同參數若同時存在於命令列與設定檔，命令列參數優先。具體參數詳見如下表格：

**參數列表**：
配置項|說明
:-|:-
login_account|登入帳號 (也可通過設定檔指定)
login_pwd|登入密碼明文 (- 也可使用登入密碼密文輸入
  - 也可通過設定檔指定)
login_pwd_md5|登入密碼密文（32 位 MD5 加密 16 進制） (- 如果密文明文都存在，則只使用密文
  - 也可使用登入密碼明文輸入)
cfg_file|OpenD 設定檔絕對路徑 (不設置則使用程式所在目錄下的 OpenD.xml)
console|是否顯示控制台 (- 0：不顯示
  - 1：顯示不設置則預設顯示控制台)
lang|中英語言 (- chs：簡體中文
  - en：英文)
api_ip|API 服務監聽地址
api_port|API 協議接收連接埠
help|輸出命令列啟動參數，並退出程式
log_level|OpenD 日誌級別 (- no（無日誌） 
  - debug（最詳細）
  - info（次詳細）)
no_monitor|是否啟動守護程序 (- 0：啟動
  - 1：不啟動)
websocket_ip|WebSocket 服務監聽地址 (可填：

  - 127.0.0.1（監聽來自本地的連接） 
  - 0.0.0.0（監聽來自所有網絡卡的連接）)
websocket_port|WebSocket 服務監聽連接埠 (不設置則不啟用 Websocket)
websocket_private_key|WebSocket 證書私鑰文件路徑 (- 私鑰不可設置密碼
  - 需要和證書同時配置
  - 不配置則不啟用 Websocket)
websocket_cert|WebSocket 證書文件路徑 (- 需要和證書同時配置
  - 不配置則不啟用 Websocket)
websocket_key_md5|密鑰密文（32 位 MD5 加密 16 進制） (JavaScript 腳本連接時，用於判斷是否可信連接)
price_reminder_push|是否接收到價提醒推送 (可填：
  - 0：不接收
  - 1：接收（需在腳本中設置到價提醒回調函數 [set_handler](/ftapi/init.html#8215)）不設置則預設接收)
auto_hold_quote_right|被踢後是否自動搶權限 (可填：
  - 0：否
  - 1：是（OpenD 在行情權限被搶後，會自動搶回。如果 10 秒內再次被搶，則其他終端獲得最高行情權限，OpenD 不會再搶）不設置則預設自動搶權限)
future_trade_api_time_zone|期貨交易 API 時區 (使用期貨賬户調用 **交易 API** 時，涉及的時間按照此時區規則)

:::

---



---

# 維護命令

透過命令列或者 Telnet 發送命令可以對 OpenD 做維護操作。

命令格式：`cmd -param_key1=param_value1 -param_key2=param_value2`

以 `help -cmd=exit` 為例，介紹Telnet的用法：
1. 在OpenD啟動參數中，設定好 Telnet 地址和 Telnet 連接埠。
![telnet_GUI](../img/telnet_GUI.jpg)
![telnet_CMD](../img/telnet_CMD.jpg)
2. 啟動 OpenD（會同時啟動 Telnet）。
3. 透過 Telnet，向 OpenD 發送 `help -cmd=exit` 命令。
```python
from telnetlib import Telnet
with Telnet('127.0.0.1', 22222) as tn:  # Telnet 地址為：127.0.0.1，Telnet 連接埠為：22222
    tn.write(b'help -cmd=exit\r\n')
    reply = b''
    while True:
        msg = tn.read_until(b'\r\n', timeout=0.5)
        reply += msg
        if msg == b'':
            break
    print(reply.decode('gb2312'))
```


## 命令幫助
`help -cmd=exit`

查看指定命令詳細資訊，不指定參數則輸出命令列表

* 參數:	
    - cmd: 命令

## 退出程式
`exit`

退出 OpenD 程式

## 請求手機驗證碼
`req_phone_verify_code `

請求手機驗證碼，當啟用裝置鎖並初次在該裝置登入，要求做安全驗證。

* 頻率限制:	
  - 每60秒內最多請求1次
  
## 輸入手機驗證碼
`input_phone_verify_code -code=123456`

輸入手機驗證碼，並繼續登入流程。

* 參數:	
  - code: 手機驗證碼

* 頻率限制:	
  - 每60秒內最多請求10次
 
## 請求圖形驗證碼
`req_pic_verify_code`

請求圖形驗證碼，當多次輸入錯登入密碼時，需要輸入圖形驗證碼。

* 頻率限制:	
  - 每60秒內最多請求10次
  
## 輸入圖形驗證碼
`input_pic_verify_code -code=1234`

輸入圖形驗證碼，並繼續登入流程。

* 參數:	
  - code: 圖形驗證碼

* 頻率限制:	
  - 每60秒內最多請求10次
  
## 重登入
`relogin -login_pwd=123456`

當登入密碼修改或中途打開裝置鎖等情況，要求用戶重新登入時，可以使用該命令。只能重登當前帳號，不支援切換帳號。
密碼參數主要用於登入密碼修改的情況，不指定密碼則使用啟動時登入密碼。

* 參數:	
  - login_pwd: 登入密碼明文
  
  - login_pwd_md5: 登入密碼密文（32 位 MD5 加密 16 進制）

* 頻率限制:	
  - 每小時最多請求10次
  
## 檢測與連接點之間的時延
`ping `

檢測與連接點之前的時延

* 頻率限制:	
  - 每60秒內最多請求10次
  
## 展示延遲統計報告
`show_delay_report -detail_report_path=D:/detail.txt -push_count_type=sr2cs`

展示延遲統計報告，包括推送延遲，請求延遲以及下單延遲。每日北京時間 6:00 清理數據。 

* 參數:	 
  - detail_report_path: 檔案輸出路徑（MAC 系統僅支援絕對路徑，不支援相對路徑），可選參數，若不指定則輸出到控制台
  
  - Paramters: push_count_type: 推送延遲的類型(sr2ss，ss2cr，cr2cs，ss2cs，sr2cs)，預設 sr2cs。
    + sr 指伺服器接收時間(目前只有港股支援該時間)
    + ss 指伺服器發出時間
    + cr 指 OpenD 接收時間 
    + cs 指 OpenD 發出時間

## 關閉 API 連接
`close_api_conn  -conn_id=123456`

關閉某條 API 連接，若不指定則關閉所有
  
  * 參數:
    - conn_id: API 連接 ID

## 展示訂閱狀態
`show_sub_info -conn_id=123456 -sub_info_path=D:/detail.txt`

展示某條連接的訂閱狀態，若不指定則展示所有
  
  * 參數:
    - conn_id: API 連接 ID
  
    - sub_info_path: 檔案輸出路徑（MAC 系統僅支援絕對路徑，不支援相對路徑），可選參數，若不指定則輸出到控制台
  
## 請求最高行情權限
`request_highest_quote_right`

當進階行情權限被其他裝置（如：桌面端/手機端）佔用時，可使用該命令重新請求最高行情權限（屆時，其他處於登入狀態的裝置將無法使用進階行情）。

* 頻率限制:	
  - 每60秒內最多請求10次

## 升級
`update`

運行該命令，可以一鍵更新 OpenD

---



---

# 行情介面總覽

<table>
    <tr>
        <th colspan="2">模組</th>
        <th>協議 ID</th>
        <th>Protobuf 定義</th>
        <th>説明</th>
    </tr>
    <tr>
        <td rowspan="15">實時行情</td>
        <td rowspan="2">訂閲</td>
        <td>3001</td>
	    <td><a href="../quote/sub.html">Qot_Sub</a></td>
	    <td>訂閲或者反訂閲</td>
    </tr>
    <tr>
        <td>3003</td>
	    <td><a href="../quote/query-subscription.html">Qot_GetSubInfo</a></td>
	    <td>獲取訂閲資訊</td>
    </tr>
    <tr>
        <td rowspan="6">推送回呼</td>
        <td>3005</td>
	    <td><a href="../quote/update-stock-quote.html">Qot_UpdateBasicQot</a></td>
	    <td>推送股票基本報價</td>
    </tr>
    <tr>
        <td>3013</td>
	    <td><a href="../quote/update-order-book.html">Qot_UpdateOrderBook</a></td>
	    <td>推送買賣盤</td>
    </tr>
    <tr>
        <td>3007</td>
	    <td><a href="../quote/update-kl.html">Qot_UpdateKL</a></td>
	    <td>推送 K 線</td>
    </tr>
    <tr>
        <td>3009</td>
	    <td><a href="../quote/update-rt.html">Qot_UpdateRT</a></td>
	    <td>推送分時</td>
    </tr>
    <tr>
        <td>3011</td>
	    <td><a href="../quote/update-ticker.html">Qot_UpdateTicker</a></td>
	    <td>推送逐筆</td>
    </tr>
    <tr>
        <td>3015</td>
	    <td><a href="../quote/update-broker.html">Qot_UpdateBroker</a></td>
	    <td>推送經紀隊列</td>
    </tr>
    <tr>
        <td rowspan="7">拉取</td>
        <td>3203</td>
	    <td><a href="../quote/get-market-snapshot.html">Qot_GetSecuritySnapshot</a></td>
	    <td>獲取股票快照</td>
    </tr>
    <tr>
        <td>3004</td>
	    <td><a href="../quote/get-stock-quote.html">Qot_GetBasicQot</a></td>
	    <td>獲取股票基本報價</td>
    </tr>
    <tr>
        <td>3012</td>
        <td><a href="../quote/get-order-book.html">Qot_GetOrderBook</a></td>
	    <td>獲取買賣盤</td>
    </tr>
    <tr>
        <td>3006</td>
	    <td><a href="../quote/get-kl.html">Qot_GetKL</a></td>
	    <td>獲取 K 線</td>
    </tr>
    <tr>
        <td>3008</td>
        <td><a href="../quote/get-rt.html">Qot_GetRT</a></td>
	    <td>獲取分時</td>
    </tr>
    <tr>
        <td>3010</td>
        <td><a href="../quote/get-ticker.html">Qot_GetTicker</a></td>
	    <td>獲取逐筆</td>
    </tr>
    <tr>
        <td>3014</td>
        <td><a href="../quote/get-broker.html">Qot_GetBroker</a></td>
	    <td>獲取經紀隊列</td>
    </tr>
    <tr>
        <td rowspan="6" colspan="2">基本數據</td>
        <td>3223</td>
	    <td><a href="../quote/get-market-state.html">Qot_GetMarketState</a></td>
	    <td>獲取指定品種的市場狀態</td>
    </tr>
    <tr>
        <td>3211</td>
        <td><a href="../quote/get-capital-flow.html">Qot_GetCapitalFlow</a></td>
	    <td>獲取資金流向</td>
    </tr>
    <tr>
        <td>3212</td>
        <td><a href="../quote/get-capital-distribution.html">Qot_GetCapitalDistribution</a></td>
	    <td>獲取資金分佈</td>
    </tr>
    <tr>
        <td>3207</td>
        <td><a href="../quote/get-owner-plate.html">Qot_GetOwnerPlate</a></td>
	    <td>獲取股票所屬板塊</td>
    </tr>
    <tr>
        <td>3103</td>
        <td><a href="../quote/request-history-kline.html">Qot_RequestHistoryKL</a></td>
	    <td>線上獲取單隻股票一段歷史 K 線</td>
    </tr>
    <tr>
        <td>3105</td>
	    <td><a href="../quote/get-rehab.html">Qot_RequestRehab</a></td>
	    <td>線上獲取單隻股票復權資訊</td>
    </tr>
    <tr>
        <td rowspan="5" colspan="2">相關衍生品</td>
        <td>3224</td>
        <td><a href="../quote/get-option-expiration-date.html">Qot_GetOptionExpirationDate</a></td>
	    <td>獲取期權到期日</td>
    </tr>
    <tr>
        <td>3209</td>
        <td><a href="../quote/get-option-chain.html">Qot_GetOptionChain</a></td>
	    <td>獲取期權鏈</td>
    </tr>
    <tr>
        <td>3210</td>
        <td><a href="../quote/get-warrant.html">Qot_GetWarrant</a></td>
	    <td>獲取窩輪</td>
    </tr>
    <tr>
        <td>3206</td>
        <td><a href="../quote/get-referencestock-list.html">Qot_GetReference</a></td>
	    <td>獲取正股相關股票</td>
    </tr>
    <tr>
        <td>3218</td>
        <td><a href="../quote/get-future-info.html">Qot_GetFutureInfo</a></td>
	    <td>獲取期貨合約資料</td>
    </tr>
    <tr>
        <td rowspan="7" colspan="2">全市場篩選</td>
        <td>3215</td>
	    <td><a href="../quote/get-stock-filter.html">Qot_StockFilter</a></td>
	    <td>獲取條件選股</td>
    </tr>
    <tr>
        <td>3205</td>
        <td><a href="../quote/get-plate-stock.html">Qot_GetPlateSecurity</a></td>
	    <td>獲取板塊下的股票</td>
    </tr>
    <tr>
        <td>3204</td>
        <td><a href="../quote/get-plate-list.html">Qot_GetPlateSet</a></td>
	    <td>獲取板塊集合下的板塊</td>
    </tr>
    <tr>
        <td>3202</td>
        <td><a href="../quote/get-static-info.html">Qot_GetStaticInfo</a></td>
	    <td>獲取股票靜態資訊</td>
    </tr>
    <tr>
        <td>3217</td>
        <td><a href="../quote/get-ipo-list.html">Qot_GetIpoList</a></td>
	    <td>獲取新股</td>
    </tr>
    <tr>
        <td>1002</td>
        <td><a href="../quote/get-global-state.html">GetGlobalState</a></td>
	    <td>獲取全域市場狀態</td>
    </tr>
    <tr>
        <td>3219</td>
        <td><a href="../quote/request-trading-days.html">Qot_RequestTradeDate</a></td>
	    <td>獲取市場交易日，線上拉取不在本地計算</td>
    </tr>
    <tr>
        <td rowspan="7" colspan="2">個人化</td>
        <td>3104</td>
        <td><a href="../quote/get-history-kl-quota.html">Qot_RequestHistoryKLQuota</a></td>
	    <td>獲取歷史 K 線額度</td>
    </tr>
    <tr>
        <td>3220</td>
        <td><a href="../quote/set-price-reminder.html">Qot_SetPriceReminder</a></td>
	    <td>設定到價提醒</td>
    </tr>
    <tr>
        <td>3221</td>
        <td><a href="../quote/get-price-reminder.html">Qot_GetPriceReminder</a></td>
	    <td>獲取到價提醒</td>
    </tr>
    <tr>
        <td>3213</td>
        <td><a href="../quote/get-user-security.html">Qot_GetUserSecurity</a></td>
	    <td>獲取自選股分組下的股票</td>
    </tr>
    <tr>
        <td>3222</td>
        <td><a href="../quote/get-user-security-group.html">Qot_GetUserSecurityGroup</a></td>
	    <td>獲取自選股分組列表</td>
    </tr>
    <tr>
        <td>3214</td>
        <td><a href="../quote/modify-user-security.html">Qot_ModifyUserSecurity</a></td>
	    <td>修改自選股分組下的股票</td>
    </tr>
    <tr>
        <td>3019</td>
	    <td><a href="../quote/update-price-reminder.html">Qot_UpdatePriceReminder</a></td>
	    <td>到價提醒通知</td>
    </tr>
</table>

---



---

# 行情物件

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>

## 建立連線

`OpenQuoteContext(host='127.0.0.1', port=11111, is_encrypt=None)`  

* **介紹**

    建立並初始化行情連線

* **參數**

    參數|類型|說明
    :-|:-|:-
    host|str|OpenD 監聽的 IP 位址
    port|int|OpenD 監聽的連接埠
    is_encrypt|bool|是否啓用加密  (- 預設為 None，表示使用 [enable_proto_encrypt](../ftapi/init.md#1542) 的設置
  - True：強制加密False：強制不加密)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111, is_encrypt=False)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

## 關閉連線

`close()`  

* **介紹**

    關閉行情介面類物件。預設情況下，Futu API 內部建立的執行緒會阻止行程退出，只有當所有 Context 都 close 後，行程才能正常退出。但通過 [set_all_thread_daemon](../ftapi/init.md#7809) 可以設置所有內部執行緒為 daemon 執行緒，這時即使沒有呼叫 Context 的 close，行程也可以正常退出。

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

## 啓動

`start()` 

* **介紹**

    啓動非同步接收推送數據

## 停止

`stop()` 

* **介紹**

    停止非同步接收推送數據

---



---

# 取消訂閲

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>

## **訂閲**  

`subscribe(code_list, subtype_list, is_first_push=True, subscribe_push=True, is_detailed_orderbook=False, extended_time=False, session=Session.NONE)` 
* **介紹**

    訂閲註冊需要的實時資訊，指定股票和訂閲的數據類型即可。  
    香港市場（含正股、窩輪、牛熊、期權、期貨）訂閲，需要 LV1 及以上的權限，BMP 權限下不支援訂閲。  
    美股市場（含正股、ETFs）夜盤行情訂閲，需要 LV1 及以上的權限，BMP 權限下不支援訂閲。  

* **參數**

    參數|類型|説明
    :-|:-|:-
    code_list|list|需要訂閲的股票代碼列表  (list 中元素類型是 str)
    subtype_list|list|需要訂閲的數據類型列表  (list 中元素類型是 [SubType](./quote.md#7865))
    is_first_push|bool|訂閲成功之後是否立即推送一次快取數據  (- True：推送快取當腳本和 OpenD 之間出現斷線重連，重新訂閲時若設定為 True，會再次推送斷線前的最後一條數據
  - False：不推送快取。等待伺服器的最新推送)
    subscribe_push|bool|訂閲後是否推送  (訂閲後，OpenD 提供了[兩種取數據的方式](../qa/quote.html#7310)，如果您僅使用 **獲取實時數據** 的方式，選擇 False 可以節省一部分效能消耗
  - True：推送。如果使用 **實時數據回調** 的方式，則必須設定為 True
  - False：不推送。如果**僅**使用 **獲取實時數據** 的方式，則建議設定為 False)
    is_detailed_orderbook|bool|是否訂閲詳細的擺盤訂單明細  (- 僅用於港股 SF 行情權限下訂閲港股 ORDER_BOOK 類型 
  - 美股美期 LV2 權限下不提供詳細擺盤訂單明細)
    extended_time|bool|是否允許美股盤前盤後數據  (僅用於訂閲美股實時 K 線、實時分時、實時逐筆)
    session|[Session](./quote.md#3103)|美股訂閲時段  (- 僅用於訂閲美股實時 K 線、實時分時、實時逐筆
  - 訂閲美股行情不支援入參OVERNIGHT
  - 最低OpenD版本：9.2.4207)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">err_message</td>
            <td >NoneType</td>
            <td>當 ret == RET_OK 時，返回 None</td>
        </tr>
        <tr>
            <td >str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>


* **Example**

``` python
import time
from futu import *
class OrderBookTest(OrderBookHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(OrderBookTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("OrderBookTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("OrderBookTest ", data) # OrderBookTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = OrderBookTest()
quote_ctx.set_handler(handler)  # 設定實時擺盤迴調
quote_ctx.subscribe(['US.AAPL'], [SubType.ORDER_BOOK])  # 訂閲買賣擺盤類型，OpenD 開始持續收到伺服器的推送
time.sleep(15)  #  設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()  # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲
```

* **Output**

``` python
OrderBookTest  {'code': 'US.AAPL', 'name': '蘋果', 'svr_recv_time_bid': '2025-04-07 05:00:52.266', 'svr_recv_time_ask': '2025-04-07 05:00:53.973', 'Bid': [(180.2, 15, 3, {}), (180.19, 1, 1, {}), (180.18, 11, 2, {}), (180.14, 200, 1, {}), (180.13, 3, 2, {}), (180.1, 99, 3, {}), (180.05, 3, 1, {}), (180.03, 400, 1, {}), (180.02, 10, 1, {}), (180.01, 100, 1, {}), (180.0, 441, 24, {})], 'Ask': [(180.3, 100, 1, {}), (180.38, 4, 2, {}), (180.4, 100, 1, {}), (180.42, 200, 1, {}), (180.46, 29, 1, {}), (180.5, 1019, 2, {}), (180.6, 1000, 1, {}), (180.8, 2001, 3, {}), (180.84, 15, 2, {}), (181.0, 2036, 4, {}), (181.2, 2000, 2, {}), (181.3, 3, 1, {}), (181.4, 2021, 3, {}), (181.5, 59, 2, {}), (181.79, 9, 1, {}), (181.8, 20, 1, {}), (181.9, 94, 4, {}), (181.98, 20, 1, {}), (182.0, 150, 7, {})]}

```

## **取消訂閲**  

`unsubscribe(code_list, subtype_list, unsubscribe_all=False)`  
* **介紹**

    取消訂閲   

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|取消訂閲的股票代碼列表  (list 中元素類型是 str)
    subtype_list|list|需要訂閲的數據類型列表  (list 中元素類型是 [SubType](./quote.md#7865))
    unsubscribe_all|bool|取消所有訂閲  (為 True 時忽略其他參數)


* **Return**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">err_message</td>
            <td>NoneType</td>
            <td>當 ret == RET_OK, 返回 None</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK, 返回錯誤描述</td>
        </tr>
    </table>

* **Example**

``` python
from futu import *
import time
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

print('current subscription status :', quote_ctx.query_subscription())  # 查詢初始訂閲狀態
ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.QUOTE, SubType.TICKER], subscribe_push=False, session=Session.ALL)
# 先訂閲了AAPL全時段 QUOTE 和 TICKER 兩個類型。訂閲成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:   # 訂閲成功
    print('subscribe successfully！current subscription status :', quote_ctx.query_subscription())  # 訂閲成功後查詢訂閲狀態
    time.sleep(60)  # 訂閲之後至少1分鐘才能取消訂閲
    ret_unsub, err_message_unsub = quote_ctx.unsubscribe(['US.AAPL'], [SubType.QUOTE])
    if ret_unsub == RET_OK:
        print('unsubscribe successfully！current subscription status:', quote_ctx.query_subscription())  # 取消訂閲後查詢訂閲狀態
    else:
        print('unsubscription failed！', err_message_unsub)
else:
    print('subscription failed', err_message)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

``` python
current subscription status : (0, {'total_used': 0, 'remain': 1000, 'own_used': 0, 'sub_list': {}})
subscribe successfully！current subscription status : (0, {'total_used': 2, 'remain': 998, 'own_used': 2, 'sub_list': {'QUOTE': ['US.AAPL'], 'TICKER': ['US.AAPL']}})
unsubscribe successfully！current subscription status: (0, {'total_used': 1, 'remain': 999, 'own_used': 1, 'sub_list': {'TICKER': ['US.AAPL']}})
```

## **取消所有訂閲**  

`unsubscribe_all()`  

* **介紹**

取消所有訂閲   


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">err_message</td>
            <td>NoneType</td>
            <td>當 ret == RET_OK, 返回 None</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK, 返回錯誤描述</td>
        </tr>
    </table>

* **Example** 

``` python
from futu import *
import time
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

print('current subscription status :', quote_ctx.query_subscription())  # 查詢初始訂閲狀態
ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.QUOTE, SubType.TICKER], subscribe_push=False, session=Session.None)
# 先訂閲了AAPL全時段 QUOTE 和 TICKER 兩個類型。訂閲成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:  # 訂閲成功
    print('subscribe successfully！current subscription status :', quote_ctx.query_subscription())  # 訂閲成功後查詢訂閲狀態
    time.sleep(60)  # 訂閲之後至少1分鐘才能取消訂閲
    ret_unsub, err_message_unsub = quote_ctx.unsubscribe_all()  # 取消所有訂閲
    if ret_unsub == RET_OK:
        print('unsubscribe all successfully！current subscription status:', quote_ctx.query_subscription())  # 取消訂閲後查詢訂閲狀態
    else:
        print('Failed to cancel all subscriptions！', err_message_unsub)
else:
    print('subscription failed', err_message)
quote_ctx.close()  # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

``` python
current subscription status : (0, {'total_used': 0, 'remain': 1000, 'own_used': 0, 'sub_list': {}})
subscribe successfully！current subscription status : (0, {'total_used': 2, 'remain': 998, 'own_used': 2, 'sub_list': {'QUOTE': ['US.AAPL'], 'TICKER': ['US.AAPL']}})
unsubscribe all successfully！current subscription status: (0, {'total_used': 0, 'remain': 1000, 'own_used': 0, 'sub_list': {}})
```

---



---

# 獲取訂閲狀態

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`query_subscription(is_all_conn=True)`

* **介紹**

    獲取訂閲資訊

* **參數**
    參數|類型|説明
    :-|:-|:-
    is_all_conn|bool|是否返回所有連線的訂閲狀態  (True：返回所有連線的訂閲狀態False：只返回當前連線的訂閲狀態)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>dict</td>
            <td>當 ret == RET_OK，返回訂閲資訊數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 訂閲資訊數據字典格式如下：
    
            {
                'total_used': 4,    # 所有連線已使用的訂閲額度
                'own_used': 0,       # 當前連線已使用的訂閲額度
                'remain': 496,       #  剩餘的訂閲額度
                'sub_list':          #  每種訂閲類型對應的股票列表
                {
                    '訂閲的類型': 該訂閲類型下所有已訂閲股票列表,
                    …
                }
            }
    
* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

quote_ctx.subscribe(['HK.00700'], [SubType.QUOTE])
ret, data = quote_ctx.query_subscription()
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
{'total_used': 1, 'remain': 999, 'own_used': 1, 'sub_list': {'QUOTE': ['HK.00700']}}
```

---



---

# 實時報價回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時報價回呼，非同步處理已訂閲股票的實時報價推送。  
    在收到實時報價數據推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateBasicQot_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回報價數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 報價數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        data_date|str|日期
        data_time|str|當前價更新時間  (格式：yyyy-MM-dd HH:mm:ss
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        last_price|float|最新價格
        open_price|float|今日開盤價
        high_price|float|最高價格
        low_price|float|最低價格
        prev_close_price|float|昨收盤價格
        volume|int|成交數量
        turnover|float|成交金額
        turnover_rate|float|換手率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        amplitude|int|振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        suspension|bool|是否停牌  (True：停牌)
        listing_date|str|上市日期  (格式：yyyy-MM-dd)
        price_spread|float|當前向上的價差  (即擺盤數據的賣檔的相鄰檔位的報價差)
        dark_status|[DarkStatus](./quote.md#6286)|暗盤交易狀態
        sec_status|[SecurityStatus](./quote.md#4506)|股票狀態
        strike_price|float|行權價
        contract_size|float|每份合約數
        open_interest|int|未平倉合約數
        implied_volatility|float|隱含波動率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        premium|float|溢價  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        delta|float|希臘值 Delta
        gamma|float|希臘值 Gamma
        vega|float|希臘值 Vega
        theta|float|希臘值 Theta
        rho|float|希臘值 Rho
        index_option_type|[IndexOptionType](./quote.md#1625)|指數期權類型
        net_open_interest|int|淨未平倉合約數  (僅港股期權適用)
        expiry_date_distance|int|距離到期日天數  (負數表示已過期)
        contract_nominal_value|float|合約名義金額  (僅港股期權適用)
        owner_lot_multiplier|float|相等正股手數  (指數期權無該欄位 ，僅港股期權適用)
        option_area_type|[OptionAreaType](./quote.md#9555)|期權類型（按行權時間）
        contract_multiplier|float|合約乘數
        pre_price|float|盤前價格
        pre_high_price|float|盤前最高價
        pre_low_price|float|盤前最低價
        pre_volume|int|盤前成交量
        pre_turnover|float|盤前成交額
        pre_change_val|float|盤前漲跌額
        pre_change_rate|float|盤前漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        pre_amplitude|float|盤前振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        after_price|float|盤後價格
        after_high_price|float|盤後最高價
        after_low_price|float|盤後最低價
        after_volume|int|盤後成交量  (科創板支援此數據)
        after_turnover|float|盤後成交額  (科創板支援此數據)
        after_change_val|float|盤後漲跌額
        after_change_rate|float|盤後漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        after_amplitude|float|盤後振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        overnight_price|float|夜盤價格
        overnight_high_price|float|夜盤最高價
        overnight_low_price|float|夜盤最低價
        overnight_volume|int|夜盤成交量
        overnight_turnover|float|夜盤成交額
        overnight_change_val|float|夜盤漲跌額
        overnight_change_rate|float|夜盤漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        overnight_amplitude|float|夜盤振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        last_settle_price|float|昨結  (期貨特有欄位)
        position|float|持倉量  (期貨特有欄位)
        position_change|float|日增倉  (期貨特有欄位)

* **Example**

```python
import time
from futu import *

class StockQuoteTest(StockQuoteHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(StockQuoteTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("StockQuoteTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("StockQuoteTest ", data) # StockQuoteTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = StockQuoteTest()
quote_ctx.set_handler(handler)  # 設定實時報價回呼
ret, data = quote_ctx.subscribe(['US.AAPL'], [SubType.QUOTE])  # 訂閲實時報價類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  #  設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲    	
```

* **Output**

```python
StockQuoteTest        code name data_date data_time  last_price  open_price  high_price  low_price  prev_close_price  volume  turnover  turnover_rate  amplitude  suspension listing_date  price_spread dark_status sec_status strike_price contract_size open_interest implied_volatility premium delta gamma vega theta  rho net_open_interest expiry_date_distance contract_nominal_value owner_lot_multiplier option_area_type contract_multiplier last_settle_price position position_change index_option_type pre_price pre_high_price pre_low_price pre_volume pre_turnover pre_change_val pre_change_rate pre_amplitude after_price after_high_price after_low_price after_volume after_turnover after_change_val after_change_rate after_amplitude overnight_price overnight_high_price overnight_low_price overnight_volume overnight_turnover overnight_change_val overnight_change_rate overnight_amplitude
0  US.AAPL   蘋果                             0.0         0.0         0.0        0.0               0.0       0       0.0            0.0        0.0       False                        0.0         N/A     NORMAL          N/A           N/A           N/A                N/A     N/A   N/A   N/A  N/A   N/A  N/A               N/A                  N/A                    N/A                  N/A              N/A                 N/A               N/A      N/A             N/A               N/A       N/A            N/A           N/A        N/A          N/A            N/A             N/A           N/A         N/A              N/A             N/A          N/A            N/A              N/A               N/A             N/A             N/A                  N/A                 N/A              N/A                N/A                  N/A                   N/A                 N/A
```

---



---

# 實時擺盤迴調

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時擺盤迴調，非同步處理已訂閲股票的實時擺盤推送。
    在收到實時擺盤資料推送後會回調到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateOrderBook_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>dict</td>
            <td>當 ret == RET_OK，返回擺盤資料</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 擺盤資料格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        svr_recv_time_bid|str|富途伺服器從交易所收到買盤資料的時間  (部分資料的接收時間為零，例如伺服器重啓或第一次推送的緩存資料)
        svr_recv_time_ask|str|富途伺服器從交易所收到賣盤資料的時間  (部分資料的接收時間為零，例如伺服器重啓或第一次推送的緩存資料)
        Bid|list|每個元祖包含如下資訊：委託價格，委託數量，委託訂單數，委託訂單明細  (委託訂單明細
  - 明細內容：交易所訂單 ID，單筆委託數量
  - 港股 SF 權限下最多支援 1000 筆委託訂單明細；其餘行情權限不支援獲取此類資料)
        Ask|list|每個元祖包含如下資訊：委託價格，委託數量，委託訂單數，委託訂單明細  (委託訂單明細
  - 明細內容：交易所訂單 ID，單筆委託數量
  - 港股 SF 權限下最多支援 1000 筆委託訂單明細；其餘行情權限不支援獲取此類資料)

        其中，Bid 和 Ask 字段的結構如下：  

          'Bid': [ (bid_price1, bid_volume1, order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }), (bid_price2, bid_volume2, order_num,  {'orderid1': order_volume1, 'orderid2': order_volume2, …… }),…]
          'Ask': [ (ask_price1, ask_volume1，order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }), (ask_price2, ask_volume2, order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }),…] 

* **Example**

```python
import time
from futu import *
class OrderBookTest(OrderBookHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(OrderBookTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("OrderBookTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("OrderBookTest ", data) # OrderBookTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = OrderBookTest()
quote_ctx.set_handler(handler)  # 設定實時擺盤迴調
ret, data = quote_ctx.subscribe(['US.AAPL'], [SubType.ORDER_BOOK])  # 訂閲買賣擺盤類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  #  設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()  # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲
```

* **Output**

```python
OrderBookTest  {'code': 'US.AAPL', 'name': '蘋果', 'svr_recv_time_bid': '', 'svr_recv_time_ask': '', 'Bid': [(179.77, 100, 1, {}), (179.68, 200, 1, {}), (179.65, 2, 2, {}), (179.64, 27, 1, {}), (179.6, 9, 2, {}), (179.58, 39, 2, {}), (179.5, 13, 4, {}), (179.48, 331, 2, {}), (179.4, 1002, 2, {}), (179.38, 330, 1, {}), (179.37, 2, 1, {}), (179.3, 47, 1, {}), (179.28, 330, 1, {}), (179.21, 2, 1, {}), (179.2, 1000, 1, {}), (179.18, 330, 1, {}), (179.17, 100, 1, {}), (179.16, 1, 1, {}), (179.13, 400, 1, {}), (179.1, 3000, 1, {}), (179.08, 330, 1, {}), (179.05, 125, 2, {}), (179.01, 17, 2, {}), (179.0, 81, 7, {})], 'Ask': [(179.95, 400, 2, {}), (180.0, 360, 2, {}), (180.05, 20, 1, {}), (180.1, 246, 4, {}), (180.18, 20, 1, {}), (180.2, 2030, 3, {}), (180.23, 20, 1, {}), (180.3, 23, 1, {}), (180.33, 15, 1, {}), (180.4, 2000, 2, {}), (180.49, 5, 1, {}), (180.59, 253, 1, {}), (180.6, 2000, 2, {}), (180.8, 2010, 3, {}), (181.0, 2018, 4, {}), (181.08, 1, 1, {}), (181.2, 1009, 2, {}), (181.3, 17, 3, {}), (181.4, 1, 1, {}), (181.5, 50, 1, {}), (181.79, 9, 1, {}), (181.9, 66, 2, {})]}
```

---



---

# 實時 K 線回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時 K 線回呼，非同步處理已訂閲股票的實時 K 線推送。

    在收到實時 K 線數據推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateKL_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回 K 線數據數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * K 線數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        time_key|str|時間  (格式：yyyy-MM-dd HH:mm:ss
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        open|float|開盤價
        close|float|收盤價
        high|float|最高價
        low|float|最低價
        volume|int|成交量
        turnover|float|成交額
        pe_ratio|float|市盈率
        turnover_rate|float|換手率  (該欄位為百分比欄位，預設返回小數，如 0.01 實際對應 1%)
        last_close|float|昨收價  (即前一個時間的收盤價出於效率原因，第一個數據的昨收價可能為 0)
        k_type|[KLType](./quote.md#9701)|K 線類型

* **Example**

```python
import time
from futu import *
class CurKlineTest(CurKlineHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(CurKlineTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("CurKlineTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("CurKlineTest ", data) # CurKlineTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = CurKlineTest()
quote_ctx.set_handler(handler)  # 設定實時K線回呼
ret, data = quote_ctx.subscribe(['US.AAPL'], [SubType.K_1M], session=Session.ALL)   # 訂閲 K 線數據類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  # 設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲    
```

* **Output**

```python
CurKlineTest        code name             time_key    open   close    high    low  volume   turnover k_type  last_close
0  US.AAPL   蘋果  2025-04-07 05:15:00  180.39  180.26  180.46  180.2    1322  238340.48   K_1M         0.0
```

---



---

# 實時分時回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時分時回呼，非同步處理已訂閲股票的實時分時推送。  
    在收到實時分時數據推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateRT_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回分時數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 分時數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        time|str|時間  (格式：yyyy-MM-dd HH:mm:ss 港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        is_blank|bool|數據狀態  (False：正常數據True：偽造數據)
        opened_mins|int|零點到當前多少分鐘
        cur_price|float|當前價格
        last_close|float|昨天收盤的價格
        avg_price|float|平均價格  (對於期權，該欄位為 None)
        volume|float|成交量
        turnover|float|成交金額

* **Example**

```python
import time
from futu import *

class RTDataTest(RTDataHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(RTDataTest, self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("RTDataTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("RTDataTest ", data) # RTDataTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = RTDataTest()
quote_ctx.set_handler(handler)  # 設定實時分時推送回呼
ret, data = quote_ctx.subscribe(['US.AAPL'], [SubType.RT_DATA], session=Session.ALL) # 訂閲分時類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  # 設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲    
```

* **Output**

```python
RTDataTest        code name                 time  is_blank  opened_mins  cur_price  last_close   avg_price   turnover  volume
0  US.AAPL   蘋果  2025-04-07 05:24:00     False          324     179.53      188.38  180.465762  651262.42    3624
```

---



---

# 實時逐筆回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時逐筆回呼，非同步處理已訂閲股票的實時逐筆推送。  
    在收到實時逐筆數據推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateTicker_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回逐筆數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 逐筆數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        sequence|int|逐筆序號
        time|str|成交時間  (格式：yyyy-MM-dd HH:mm:ss:xxx
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        price|float|成交價格
        volume|int|成交數量  (股數)
        turnover|float|成交金額
        ticker_direction|[TickerDirect](./quote.md#7136)|逐筆方向
        type|[TickerType](./quote.md#7400)|逐筆類型
        push_data_type|[PushDataType](./quote.md#817)|數據來源

* **Example**

```python
import time
from futu import *

class TickerTest(TickerHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, data = super(TickerTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("TickerTest: error, msg: %s" % data)
            return RET_ERROR, data
        print("TickerTest ", data) # TickerTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = TickerTest()
quote_ctx.set_handler(handler)  # 設定實時逐筆推送回呼
ret, data = quote_ctx.subscribe(['US.AAPL'], [SubType.TICKER], session=Session.ALL) # 訂閲逐筆類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  # 設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲	
```

* **Output**

```python
TickerTest        code name                     time   price  volume  turnover ticker_direction             sequence     type push_data_type
0  US.AAPL   蘋果  2025-04-07 05:25:44.116  179.81       9   1618.29          NEUTRAL  7490500033117159426  ODD_LOT          CACHE

```

---



---

# 實時經紀隊列回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    實時經紀隊列回呼，非同步處理已訂閲股票的實時經紀隊列推送。  
    在收到實時經紀隊列數據推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
	
* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdateBroker_pb2.Response|衍生類別中不需要直接處理該參數


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>tuple</td>
            <td>當 ret == RET_OK，返回經紀隊列數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 經紀隊列元組內容如下：
        欄位|類型|説明
        :-|:-|:-
        stock_code|str|股票
        bid_frame_table|pd.DataFrame|買盤數據
        ask_frame_table|pd.DataFrame|賣盤數據

        * bid_frame_table 格式如下：
            欄位|類型|説明
            :-|:-|:-
            code|str|股票代碼
            name|str|股票名稱
            bid_broker_id|int|經紀買盤 ID
            bid_broker_name|str|經紀買盤名稱
            bid_broker_pos|int|經紀檔位
            order_id|int|交易所訂單 ID  (- 不是下單介面返回的訂單 ID
  - 只有港股 SF 行情權限支援返回該欄位)
            order_volume|int|單筆委託數量  (只有港股 SF 行情權限支援返回該欄位)
        * ask_frame_table 格式如下：
            欄位|類型|説明
            :-|:-|:-
            code|str|股票代碼
            name|str|股票名稱
            ask_broker_id|int|經紀賣盤 ID
            ask_broker_name|str|經紀賣盤名稱
            ask_broker_pos|int|經紀檔位
            order_id|int|交易所訂單 ID  (- 不是下單介面返回的訂單 ID
  - 只有港股 SF 行情權限支援返回該欄位)
            order_volume|int|單筆委託數量  (只有港股 SF 行情權限支援返回該欄位)

* **Example**

```python
import time
from futu import *
    
class BrokerTest(BrokerHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, err_or_stock_code, data = super(BrokerTest, self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("BrokerTest: error, msg: {}".format(err_or_stock_code))
            return RET_ERROR, data
        print("BrokerTest: stock: {} data: {} ".format(err_or_stock_code, data))  # BrokerTest 自己的處理邏輯
        return RET_OK, data
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = BrokerTest()
quote_ctx.set_handler(handler)  # 設定實時經紀推送回呼
ret, data = quote_ctx.subscribe(['HK.00700'], [SubType.BROKER]) # 訂閲經紀類型，OpenD 開始持續收到伺服器的推送
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
time.sleep(15)  # 設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲
```

* **Output**

```python
BrokerTest: stock: HK.00700 data: [        code  name  bid_broker_id bid_broker_name  bid_broker_pos order_id order_volume
0   HK.00700  騰訊控股           5338          J.P.摩根               1      N/A          N/A
..       ...   ...            ...             ...             ...      ...          ...
36  HK.00700  騰訊控股           8305  富途證券國際(香港)有限公司               4      N/A          N/A

[37 rows x 7 columns],         code  name  ask_broker_id ask_broker_name  ask_broker_pos order_id order_volume
0   HK.00700  騰訊控股           1179  華泰金融控股(香港)有限公司               1      N/A          N/A
..       ...   ...            ...             ...             ...      ...          ...
39  HK.00700  騰訊控股           6996      中國投資信息有限公司               1      N/A          N/A

[40 rows x 7 columns]] 
```

---



---

# 獲取快照

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_market_snapshot(code_list)`

* **介紹**

    獲取快照數據

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|股票代碼列表  (每次最多可請求 400 個標的list 內元素類型為 str)


* **返回**
 
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回股票快照數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 股票快照數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        update_time|str|當前價更新時間  (格式：yyyy-MM-dd HH:mm:ss 港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        last_price|float|最新價格
        open_price|float|今日開盤價
        high_price|float|最高價格
        low_price|float|最低價格
        prev_close_price|float|昨收盤價格
        volume|int|成交數量
        turnover|float|成交金額
        turnover_rate|float|換手率  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        suspension|bool|是否停牌  (True：停牌)
        listing_date|str|上市日期  (格式：yyyy-MM-dd)
        equity_valid|bool|是否正股  (此欄位返回為 True 時，以下正股相關欄位才有合法數值)
        issued_shares|int|總股本
        total_market_val|float|總市值  (單位：元)
        net_asset|int|資產淨值
        net_profit|int|淨利潤
        earning_per_share|float|每股盈利
        outstanding_shares|int|流通股本
        net_asset_per_share|float|每股淨資產
        circular_market_val|float|流通市值  (單位：元)
        ey_ratio|float|收益率  (該欄位為比例欄位，預設不顯示 %)
        pe_ratio|float|市盈率  (該欄位為比例欄位，預設不顯示 %)
        pb_ratio|float|市淨率  (該欄位為比例欄位，預設不顯示 %)
        pe_ttm_ratio|float|市盈率 TTM  (該欄位為比例欄位，預設不顯示 %)
        dividend_ttm|float|股息 TTM，派息
        dividend_ratio_ttm|float|股息率 TTM  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        dividend_lfy|float|股息 LFY，上一年度派息
        dividend_lfy_ratio|float|股息率 LFY  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        stock_owner|str|窩輪所屬正股的代碼或期權的標的股代碼
        wrt_valid|bool|是否是窩輪  (此欄位返回為 True 時，以下窩輪相關欄位才有合法數值)
        wrt_conversion_ratio|float|換股比率
        wrt_type|[WrtType](./quote.md#9275)|窩輪類型
        wrt_strike_price|float|行使價格
        wrt_maturity_date|str|格式化窩輪到期時間
        wrt_end_trade|str|格式化窩輪最後交易時間
        wrt_leverage|float|槓桿比率  (單位：倍)
        wrt_ipop|float|價內/價外  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        wrt_break_even_point|float|打和點
        wrt_conversion_price|float|換股價
        wrt_price_recovery_ratio|float|正股距收回價  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        wrt_score|float|窩輪綜合評分
        wrt_code|str|窩輪對應的正股（此欄位已廢除，修改為 stock_owner）
        wrt_recovery_price|float|窩輪收回價
        wrt_street_vol|float|窩輪街貨量
        wrt_issue_vol|float|窩輪發行量
        wrt_street_ratio|float|窩輪街貨佔比  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        wrt_delta|float|窩輪對沖值
        wrt_implied_volatility|float|窩輪引伸波幅
        wrt_premium|float|窩輪溢價  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        wrt_upper_strike_price|float|上限價  (僅界內證支援該欄位)
        wrt_lower_strike_price|float|下限價  (僅界內證支援該欄位)
        wrt_inline_price_status|[PriceType](./quote.md#5665)|界內界外  (僅界內證支援該欄位)
        wrt_issuer_code|str|發行人代碼
        option_valid|bool|是否是期權  (此欄位返回為 True 時，以下期權相關欄位才有合法數值)
        option_type|[OptionType](./quote.md#7263)|期權類型
        strike_time|str|期權行權日  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        option_strike_price|float|行權價
        option_contract_size|float|每份合約數
        option_open_interest|int|總未平倉合約數
        option_implied_volatility|float|隱含波動率
        option_premium|float|溢價
        option_delta|float|希臘值 Delta
        option_gamma|float|希臘值 Gamma
        option_vega|float|希臘值 Vega
        option_theta|float|希臘值 Theta
        option_rho|float|希臘值 Rho
        index_option_type|[IndexOptionType](./quote.md#1625)|指數期權類型
        option_net_open_interest|int|淨未平倉合約數  (僅港股期權適用)
        option_expiry_date_distance|int|距離到期日天數  (負數表示已過期)
        option_contract_nominal_value|float|合約名義金額  (僅港股期權適用)
        option_owner_lot_multiplier|float|相等正股手數  (指數期權無該欄位，僅港股期權適用)
        option_area_type|[OptionAreaType](./quote.md#9555)|期權類型（按行權時間）
        option_contract_multiplier|float|合約乘數
        plate_valid|bool|是否為板塊類型  (此欄位返回為 True 時，以下板塊相關欄位才有合法數值)
        plate_raise_count|int|板塊類型上漲支數
        plate_fall_count|int|板塊類型下跌支數
        plate_equal_count|int|板塊類型平盤支數
        index_valid|bool|是否有指數類型  (此欄位返回為 True 時，以下指數相關欄位才有合法數值)
        index_raise_count|int|指數類型上漲支數
        index_fall_count|int|指數類型下跌支數
        index_equal_count|int|指數類型平盤支數
        lot_size|int|每手股數，股票期權表示每份合約的股數  (指數期權無該欄位)，期貨表示合約乘數
        price_spread|float|當前向上的擺盤價差  (即擺盤數據的賣一價相鄰檔位的報價差)
        ask_price|float|賣價
        bid_price|float|買價
        ask_vol|float|賣量
        bid_vol|float|買量
        enable_margin|bool|是否可融資（已廢棄）  (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        mortgage_ratio|float|股票抵押率（已廢棄）
        long_margin_initial_ratio|float|融資初始保證金率（已廢棄）  (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        enable_short_sell|bool|是否可賣空（已廢棄）  (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        short_sell_rate|float|賣空參考利率（已廢棄）  (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        short_available_volume|int|剩餘可賣空數量（已廢棄） (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        short_margin_initial_ratio|float|賣空（融券）初始保證金率（已廢棄）  (請使用 [獲取融資融券數據](../trade/get-margin-ratio.html) 介面獲取)
        sec_status|[SecurityStatus](./quote.md#4506)|股票狀態
        amplitude|float|振幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        avg_price|float|平均價
        bid_ask_ratio|float|委比  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        volume_ratio|float|量比
        highest52weeks_price|float|52 周最高價
        lowest52weeks_price|float|52 周最低價
        highest_history_price|float|歷史最高價
        lowest_history_price|float|歷史最低價
        pre_price|float|盤前價格
        pre_high_price|float|盤前最高價
        pre_low_price|float|盤前最低價
        pre_volume|int|盤前成交量
        pre_turnover|float|盤前成交額
        pre_change_val|float|盤前漲跌額
        pre_change_rate|float|盤前漲跌幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        pre_amplitude|float|盤前振幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        after_price|float|盤後價格
        after_high_price|float|盤後最高價
        after_low_price|float|盤後最低價
        after_volume|int|盤後成交量  (科創板支援該數據)
        after_turnover|float|盤後成交額  (科創板支援該數據)
        after_change_val|float|盤後漲跌額
        after_change_rate|float|盤後漲跌幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        after_amplitude|float|盤後振幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        overnight_price|float|夜盤價格
        overnight_high_price|float|夜盤最高價
        overnight_low_price|float|夜盤最低價
        overnight_volume|int|夜盤成交量
        overnight_turnover|float|夜盤成交額
        overnight_change_val|float|夜盤漲跌額
        overnight_change_rate|float|夜盤漲跌幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        overnight_amplitude|float|夜盤振幅  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        future_valid|bool|是否期貨
        future_last_settle_price|float|昨結
        future_position|float|持倉量
        future_position_change|float|日增倉
        future_main_contract|bool|是否主連合約
        future_last_trade_time|str|最後交易時間  (主連，當月，下月等期貨沒有該欄位)
        trust_valid|bool|是否基金
        trust_dividend_yield|float|股息率  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        trust_aum|float|資產規模  (單位：元)
        trust_outstanding_units|int|總發行量
        trust_netAssetValue|float|單位淨值
        trust_premium|float|溢價  (該欄位為百分比欄位，預設不顯示 %，如 20 實際對應 20%)
        trust_assetClass|[AssetClass](./quote.md#1879)|資產類別

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_market_snapshot(['HK.00700', 'US.AAPL'])
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
code  name              update_time  last_price  open_price  high_price  low_price  prev_close_price     volume      turnover  turnover_rate  suspension listing_date  lot_size  price_spread  stock_owner  ask_price  bid_price  ask_vol  bid_vol  enable_margin  mortgage_ratio  long_margin_initial_ratio  enable_short_sell  short_sell_rate  short_available_volume  short_margin_initial_ratio  amplitude  avg_price  bid_ask_ratio  volume_ratio  highest52weeks_price  lowest52weeks_price  highest_history_price  lowest_history_price  close_price_5min  after_volume  after_turnover sec_status  equity_valid  issued_shares  total_market_val     net_asset    net_profit  earning_per_share  outstanding_shares  circular_market_val  net_asset_per_share  ey_ratio  pe_ratio  pb_ratio  pe_ttm_ratio  dividend_ttm  dividend_ratio_ttm  dividend_lfy  dividend_lfy_ratio  wrt_valid  wrt_conversion_ratio wrt_type  wrt_strike_price  wrt_maturity_date  wrt_end_trade  wrt_recovery_price  wrt_street_vol  \
0  HK.00700  騰訊控股      2025-04-07 16:09:07      435.40      441.80      462.40     431.00            497.80  123364114  5.499476e+10          1.341       False   2004-06-16       100          0.20          NaN      435.4     435.20   281300    17300            NaN             NaN                        NaN                NaN              NaN                     NaN                         NaN      6.308    445.792        -68.499         5.627             547.00000           294.400000             706.100065            -13.202011            431.60             0    0.000000e+00     NORMAL          True     9202391012      4.006721e+12  1.051300e+12  2.095753e+11             22.774          9202391012         4.006721e+12              114.242     0.199    19.118     3.811        19.118          3.48                0.80          3.48               0.799      False                   NaN      N/A               NaN                NaN            NaN                 NaN             NaN   
1   US.AAPL    蘋果  2025-04-07 05:30:43.301      188.38      193.89      199.88     187.34            203.19  125910913  2.424473e+10          0.838       False   1980-12-12         1          0.01          NaN      180.8     180.48       29      400            NaN             NaN                        NaN                NaN              NaN                     NaN                         NaN      6.172    192.554         86.480         2.226             259.81389           163.300566             259.813890              0.053580            188.93       3151311    5.930968e+08     NORMAL          True    15022073000      2.829858e+12  6.675809e+10  9.133420e+10              6.080         15016677308         2.828842e+12                4.444     1.417    30.983    42.389        29.901          0.99                0.53          0.98               0.520      False                   NaN      N/A               NaN                NaN            NaN                 NaN             NaN   

   wrt_issue_vol  wrt_street_ratio  wrt_delta  wrt_implied_volatility  wrt_premium  wrt_leverage  wrt_ipop  wrt_break_even_point  wrt_conversion_price  wrt_price_recovery_ratio  wrt_score  wrt_upper_strike_price  wrt_lower_strike_price wrt_inline_price_status  wrt_issuer_code  option_valid option_type  strike_time  option_strike_price  option_contract_size  option_open_interest  option_implied_volatility  option_premium  option_delta  option_gamma  option_vega  option_theta  option_rho  option_net_open_interest  option_expiry_date_distance  option_contract_nominal_value  option_owner_lot_multiplier option_area_type  option_contract_multiplier index_option_type  index_valid  index_raise_count  index_fall_count  index_equal_count  plate_valid  plate_raise_count  plate_fall_count  plate_equal_count  future_valid  future_last_settle_price  future_position  future_position_change  future_main_contract  future_last_trade_time  trust_valid  trust_dividend_yield  trust_aum  \
0            NaN               NaN        NaN                     NaN          NaN           NaN       NaN                   NaN                   NaN                       NaN        NaN                     NaN                     NaN                     N/A              NaN         False         N/A          NaN                  NaN                   NaN                   NaN                        NaN             NaN           NaN           NaN          NaN           NaN         NaN                       NaN                          NaN                            NaN                          NaN              N/A                         NaN               N/A        False                NaN               NaN                NaN        False                NaN               NaN                NaN         False                       NaN              NaN                     NaN                   NaN                     NaN        False                   NaN        NaN   
1            NaN               NaN        NaN                     NaN          NaN           NaN       NaN                   NaN                   NaN                       NaN        NaN                     NaN                     NaN                     N/A              NaN         False         N/A          NaN                  NaN                   NaN                   NaN                        NaN             NaN           NaN           NaN          NaN           NaN         NaN                       NaN                          NaN                            NaN                          NaN              N/A                         NaN               N/A        False                NaN               NaN                NaN        False                NaN               NaN                NaN         False                       NaN              NaN                     NaN                   NaN                     NaN        False                   NaN        NaN   

   trust_outstanding_units  trust_netAssetValue  trust_premium trust_assetClass pre_price pre_high_price pre_low_price pre_volume pre_turnover pre_change_val pre_change_rate pre_amplitude after_price after_high_price after_low_price after_change_val after_change_rate after_amplitude overnight_price overnight_high_price overnight_low_price overnight_volume overnight_turnover overnight_change_val overnight_change_rate overnight_amplitude  
0                      NaN                  NaN            NaN              N/A       N/A            N/A           N/A        N/A          N/A            N/A             N/A           N/A         N/A              N/A             N/A              N/A               N/A             N/A             N/A                  N/A                 N/A              N/A                N/A                  N/A                   N/A                 N/A  
1                      NaN                  NaN            NaN              N/A    180.68         181.98        177.47     276016  49809244.83           -7.7          -4.087         2.394       186.6          188.639          186.44            -1.78            -0.944          1.1673          176.94                186.5               174.4           533115        94944250.56               -11.44                -6.072              6.4231  
HK.00700
['HK.00700', 'US.AAPL']

```

---



---

# 獲取即時報價

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_stock_quote(code_list)`

* **介紹**

    獲取已訂閱股票的即時報價，必須要先訂閱。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|股票代碼列表  (list 中元素類型是 str)
    


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回報價數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 報價數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        data_date|str|日期
        data_time|str|當前價更新時間  (格式：yyyy-MM-dd HH:mm:ss
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        last_price|float|最新價格
        open_price|float|今日開盤價
        high_price|float|最高價格
        low_price|float|最低價格
        prev_close_price|float|昨收盤價格
        volume|int|成交數量
        turnover|float|成交金額
        turnover_rate|float|換手率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        amplitude|int|振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        suspension|bool|是否停牌  (True：停牌)
        listing_date|str|上市日期  (格式：yyyy-MM-dd)
        price_spread|float|當前向上的價差  (即擺盤數據的賣檔的相鄰檔位的報價差)
        dark_status|[DarkStatus](./quote.md#6286)|暗盤交易狀態
        sec_status|[SecurityStatus](./quote.md#4506)|股票狀態
        strike_price|float|行權價
        contract_size|float|每份合約數
        open_interest|int|未平倉合約數
        implied_volatility|float|隱含波動率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        premium|float|溢價  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        delta|float|希臘值 Delta
        gamma|float|希臘值 Gamma
        vega|float|希臘值 Vega
        theta|float|希臘值 Theta
        rho|float|希臘值 Rho
        index_option_type|[IndexOptionType](./quote.md#1625)|指數期權類型
        net_open_interest|int|淨未平倉合約數  (僅港股期權適用)
        expiry_date_distance|int|距離到期日天數  (負數表示已過期)
        contract_nominal_value|float|合約名義金額  (僅港股期權適用)
        owner_lot_multiplier|float|相等正股手數  (指數期權無該欄位 ，僅港股期權適用)
        option_area_type|[OptionAreaType](./quote.md#9555)|期權類型（按行權時間）
        contract_multiplier|float|合約乘數
        pre_price|float|盤前價格
        pre_high_price|float|盤前最高價
        pre_low_price|float|盤前最低價
        pre_volume|int|盤前成交量
        pre_turnover|float|盤前成交額
        pre_change_val|float|盤前漲跌額
        pre_change_rate|float|盤前漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        pre_amplitude|float|盤前振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        after_price|float|盤後價格
        after_high_price|float|盤後最高價
        after_low_price|float|盤後最低價
        after_volume|int|盤後成交量  (科創板支援此數據)
        after_turnover|float|盤後成交額  (科創板支援此數據)
        after_change_val|float|盤後漲跌額
        after_change_rate|float|盤後漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        after_amplitude|float|盤後振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        overnight_price|float|夜盤價格
        overnight_high_price|float|夜盤最高價
        overnight_low_price|float|夜盤最低價
        overnight_volume|int|夜盤成交量
        overnight_turnover|float|夜盤成交額
        overnight_change_val|float|夜盤漲跌額
        overnight_change_rate|float|夜盤漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        overnight_amplitude|float|夜盤振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        last_settle_price|float|昨結  (期貨特有欄位)
        position|float|持倉量  (期貨特有欄位)
        position_change|float|日增倉  (期貨特有欄位)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.QUOTE], subscribe_push=False)
# 先訂閱 K 線類型。訂閱成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:  # 訂閱成功
    ret, data = quote_ctx.get_stock_quote(['US.AAPL'])  # 獲取訂閱股票報價的即時數據
    if ret == RET_OK:
        print(data)
        print(data['code'][0])   # 取第一條的股票代碼
        print(data['code'].values.tolist())   # 轉為 list
    else:
        print('error:', data)
else:
    print('subscription failed', err_message)
quote_ctx.close()  # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閱
```

* **Output**

```python
code name   data_date     data_time  last_price  open_price  high_price  low_price  prev_close_price     volume      turnover  turnover_rate  amplitude  suspension listing_date  price_spread dark_status sec_status strike_price contract_size open_interest implied_volatility premium delta gamma vega theta  rho net_open_interest expiry_date_distance contract_nominal_value owner_lot_multiplier option_area_type contract_multiplier last_settle_price position position_change index_option_type  pre_price  pre_high_price  pre_low_price  pre_volume  pre_turnover  pre_change_val  pre_change_rate  pre_amplitude  after_price  after_high_price  after_low_price  after_volume  after_turnover  after_change_val  after_change_rate  after_amplitude  overnight_price  overnight_high_price  overnight_low_price  overnight_volume  overnight_turnover  overnight_change_val  overnight_change_rate  overnight_amplitude
0  US.AAPL   蘋果  2025-04-07  05:37:21.794      188.38      193.89      199.88     187.34            203.19  125910913  2.424473e+10          0.838      6.172       False   1980-12-12          0.01         N/A     NORMAL          N/A           N/A           N/A                N/A     N/A   N/A   N/A  N/A   N/A  N/A               N/A                  N/A                    N/A                  N/A              N/A                 N/A               N/A      N/A             N/A               N/A     181.43          181.98         177.47      288853   52132735.18           -6.95           -3.689          2.394        186.6           188.639           186.44       3151311    5.930968e+08             -1.78             -0.944           1.1673           176.94                 186.5                174.4            533115         94944250.56                -11.44                 -6.072               6.4231
US.AAPL
['US.AAPL']
```

---



---

# 獲取實時擺盤

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_order_book(code, num=10)`

* **介紹**

    獲取已訂閱股票的實時擺盤，必須要先訂閱。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    name|str|股票名稱
    num|int|請求擺盤檔數  (擺盤檔數獲取上限請參見 [擺盤檔數明細](../qa/quote.md#6889)) 


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>dict</td>
            <td>當 ret == RET_OK，返回擺盤數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

   * 擺盤數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        svr_recv_time_bid|str|富途伺服器從交易所收到買盤數據的時間  (部分數據的接收時間為零，例如伺服器重啓或第一次推送的快取數據)
        svr_recv_time_ask|str|富途伺服器從交易所收到賣盤數據的時間  (部分數據的接收時間為零，例如伺服器重啓或第一次推送的快取數據)
        Bid|list|每個元祖包含如下資訊：委託價格，委託數量，委託訂單數，委託訂單明細  (委託訂單明細
  - 明細內容：交易所訂單 ID，單筆委託數量
  - 港股 SF 權限下最多支援 1000 筆委託訂單明細；其餘行情權限不支援獲取此類數據)
        Ask|list|每個元祖包含如下資訊：委託價格，委託數量，委託訂單數，委託訂單明細  (委託訂單明細
  - 明細內容：交易所訂單 ID，單筆委託數量
  - 港股 SF 權限下最多支援 1000 筆委託訂單明細；其餘行情權限不支援獲取此類數據)

     其中，Bid 和 Ask 欄位的結構如下：  

          'Bid': [ (bid_price1, bid_volume1, order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }), (bid_price2, bid_volume2, order_num,  {'orderid1': order_volume1, 'orderid2': order_volume2, …… }),…]
          'Ask': [ (ask_price1, ask_volume1，order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }), (ask_price2, ask_volume2, order_num, {'orderid1': order_volume1, 'orderid2': order_volume2, …… }),…] 

 
    
* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret_sub = quote_ctx.subscribe(['US.AAPL'], [SubType.ORDER_BOOK], subscribe_push=False)[0]
# 先訂閱買賣擺盤類型。訂閱成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:  # 訂閱成功
    ret, data = quote_ctx.get_order_book('US.AAPL', num=3)  # 獲取一次 3 檔實時擺盤數據
    if ret == RET_OK:
        print(data)
    else:
        print('error:', data)
else:
    print('subscription failed')
quote_ctx.close()  # 關閉當條連線，OpenD 會在 1 分鐘後自動取消相應股票相應類型的訂閱
```

* **Output**

```python
{'code': 'US.AAPL', 'name': '蘋果', 'svr_recv_time_bid': '2025-04-07 05:39:20.352', 'svr_recv_time_ask': '2025-04-07 05:39:20.352', 'Bid': [(181.17, 227, 2, {}), (181.15, 2, 2, {}), (181.12, 100, 1, {})], 'Ask': [(181.71, 200, 1, {}), (181.79, 9, 1, {}), (181.9, 616, 3, {})]}
```

---



---

# 獲取實時 K 線

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_cur_kline(code, num, ktype=KLType.K_DAY, autype=AuType.QFQ)`

* **介紹**

    獲取已訂閱股票的實時 K 線數據，必須要先訂閱。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    name|str|股票名稱
    num|int|K 線數據個數  (最多 1000 根)
    ktype|[KLType](./quote.md#9701)|K 線類型
    autype|[AuType](./quote.md#6706)|復權類型


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回 K 線數據數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * K 線數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        time_key|str|時間  (格式：yyyy-MM-dd HH:mm:ss
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        open|float|開盤價
        close|float|收盤價
        high|float|最高價
        low|float|最低價
        volume|int|成交量
        turnover|float|成交額
        pe_ratio|float|市盈率
        turnover_rate|float|換手率  (該欄位為百分比欄位，預設返回小數，如 0.01 實際對應 1%)
        last_close|float|昨收價  (即前一個時間的收盤價為了效率原因，第一個數據的昨收價可能為 0)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.K_DAY], subscribe_push=False, session=Session.ALL)
# 先訂閱 K 線類型。訂閱成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:  # 訂閱成功
    ret, data = quote_ctx.get_cur_kline('US.AAPL', 2, KLType.K_DAY, AuType.QFQ)  # 獲取美股AAPL最近2個 K 線數據
    if ret == RET_OK:
        print(data)
        print(data['turnover_rate'][0])   # 取第一條的換手率
        print(data['turnover_rate'].values.tolist())   # 轉為 list
    else:
        print('error:', data)
else:
    print('subscription failed', err_message)
quote_ctx.close()  # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閱
```

* **Output**

```python
code name             time_key    open   close    high     low     volume      turnover  pe_ratio  turnover_rate  last_close
0  US.AAPL   蘋果  2025-04-03 00:00:00  205.54  203.19  207.49  201.25  103419006  2.111773e+10    33.419        0.00689      223.89
1  US.AAPL   蘋果  2025-04-04 00:00:00  193.89  188.38  199.88  187.34  125910913  2.424473e+10    30.983        0.00838      203.19
0.00689
[0.00689, 0.00838]
```

---



---

# 獲取即時分時

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_rt_data(code)`

* **介紹**

    獲取已訂閱股票的即時分時資料，必須要先訂閱。

* **參數**

    參數|類型|説明
    :-|:-|:-
    code|str|股票


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回分時資料</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 分時資料格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        time|str|時間  (格式：yyyy-MM-dd HH:mm:ss 港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        is_blank|bool|資料狀態  (False：正常資料True：偽造資料)
        opened_mins|int|零點到當前多少分鐘
        cur_price|float|當前價格
        last_close|float|昨天收盤的價格
        avg_price|float|平均價格  (對於期權，該欄位為 N/A)
        volume|float|成交量
        turnover|float|成交金額

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.RT_DATA], subscribe_push=False, session=Session.ALL)
# 先訂閱分時資料類型。訂閱成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:   # 訂閱成功
    ret, data = quote_ctx.get_rt_data('US.AAPL')   # 獲取一次分時資料
    if ret == RET_OK:
        print(data)
    else:
        print('error:', data)
else:
    print('subscription failed', err_message)
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閱
```

* **Output**

```python
code  name                 time  is_blank  opened_mins  cur_price  last_close   avg_price   volume     turnover
0    US.AAPL   蘋果  2025-04-06 20:01:00     False         1201     183.00      188.38  181.643916    9463  1718896.38
..      ...    ...                  ...       ...          ...        ...         ...         ...      ...          ...
586  US.AAPL   蘋果  2025-04-07 05:47:00     False          347     181.26      188.38  180.555673     661   119859.75

[587 rows x 10 columns]
```

---



---

# 獲取實時逐筆

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_rt_ticker(code, num=500)`

* **介紹**

    獲取已訂閲股票的實時逐筆數據，必須要先訂閲。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    num|int|最近逐筆個數


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回逐筆數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 逐筆數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        sequence|int|逐筆序號
        time|str|成交時間  (格式：yyyy-MM-dd HH:mm:ss:xxx
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        price|float|成交價格
        volume|int|成交數量  (股數)
        turnover|float|成交金額
        ticker_direction|[TickerDirect](./quote.md#7136)|逐筆方向
        type|[TickerType](./quote.md#7400)|逐筆類型

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret_sub, err_message = quote_ctx.subscribe(['US.AAPL'], [SubType.TICKER], subscribe_push=False, session=Session.ALL)
# 先訂閲逐筆類型。訂閲成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:  # 訂閲成功
    ret, data = quote_ctx.get_rt_ticker('US.AAPL', 2)  # 獲取美股AAPL最近2個逐筆
    if ret == RET_OK:
        print(data)
        print(data['turnover'][0])   # 取第一條的成交金額
        print(data['turnover'].values.tolist())   # 轉為 list
    else:
        print('error:', data)
else:
    print('subscription failed', err_message)
quote_ctx.close()  # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲
```

* **Output**

```python
code name                     time   price  volume  turnover ticker_direction             sequence     type
0  US.AAPL   蘋果  2025-04-07 05:50:23.745  181.70       2    363.40          NEUTRAL  7490506385373790208  ODD_LOT
1  US.AAPL   蘋果  2025-04-07 05:50:24.170  181.73       1    181.73          NEUTRAL  7490506389668757504  ODD_LOT
363.4
[363.4, 181.73]
```

---



---

# 獲取實時經紀佇列

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_broker_queue(code)`

* **介紹**

    獲取已訂閱股票的實時經紀佇列數據，必須要先訂閱。

* **參數**

    參數|類型|說明
    :-|:-|:-
    code|str|股票代號


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">bid_frame_table</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，bid_frame_table 傳回買盤經紀佇列數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，bid_frame_table 傳回錯誤描述</td>
        </tr>
        <tr>
            <td rowspan="2">ask_frame_table</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，ask_frame_table 傳回賣盤經紀佇列數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，ask_frame_table 傳回錯誤描述</td>
        </tr>
    </table>

    * 買盤經紀佇列格式如下：
        欄位|類型|說明
        :-|:-|:-
        code|str|股票代號
        name|str|股票名稱
        bid_broker_id|int|經紀買盤 ID
        bid_broker_name|str|經紀買盤名稱
        bid_broker_pos|int|經紀檔位
        order_id|int|交易所訂單 ID  (- 不是下單介面傳回的訂單 ID
  - 只有港股 SF 行情權限支援傳回該欄位)
        order_volume|int|單筆委託數量  (只有港股 SF 行情權限支援傳回該欄位)
    * 賣盤經紀佇列格式如下：
        欄位|類型|說明
        :-|:-|:-
        code|str|股票代號
        name|str|股票名稱
        ask_broker_id|int|經紀賣盤 ID
        ask_broker_name|str|經紀賣盤名稱
        ask_broker_pos|int|經紀檔位
        order_id|int|交易所訂單 ID  (- 不是下單介面傳回的訂單 ID
  - 只有港股 SF 行情權限支援傳回該欄位)
        order_volume|int|單筆委託數量  (只有港股 SF 行情權限支援傳回該欄位)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret_sub, err_message = quote_ctx.subscribe(['HK.00700'], [SubType.BROKER], subscribe_push=False)
# 先訂閱經紀佇列類型。訂閱成功後 OpenD 將持續收到伺服器的推送，False 代表暫時不需要推送給腳本
if ret_sub == RET_OK:   # 訂閱成功
    ret, bid_frame_table, ask_frame_table = quote_ctx.get_broker_queue('HK.00700')   # 獲取一次經紀佇列數據
    if ret == RET_OK:
        print(bid_frame_table)
    else:
        print('error:', bid_frame_table)
else:
    print(err_message)
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閱
```

* **Output**

```python
        code  name  bid_broker_id bid_broker_name  bid_broker_pos order_id order_volume
0   HK.00700  騰訊控股           5338          J.P.摩根               1      N/A          N/A
..       ...   ...            ...             ...             ...      ...          ...
36  HK.00700  騰訊控股           8305  富途證券國際(香港)有限公司               4      N/A          N/A

[37 rows x 7 columns]
```

---



---

# 獲取標的市場狀態

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_market_state(code_list)`

* **介紹**

    獲取指定標的的市場狀態

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|需要查詢市場狀態的股票代碼列表  (list 中元素類型是 str)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回市場狀態數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 市場狀態數據
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        stock_name|str|股票名稱
        market_state|[MarketState](./quote.md#8754)|市場狀態

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_market_state(['SZ.000001', 'HK.00700'])
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code         stock_name   market_state
0  SZ.000001    平安銀行     AFTERNOON
1  HK.00700     騰訊控股     AFTERNOON
```

---



---

# 獲取資金流向

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_capital_flow(stock_code, period_type = PeriodType.INTRADAY, start=None, end=None)`

* **介紹**

    獲取個股資金流向

* **參數**
    參數|類型|說明
    :-|:-|:-
    stock_code|str|股票代號
    period_type|[PeriodType](./quote.md#4950)|週期類型
    start|str|開始時間  (格式：yyyy-MM-dd 
 例如：“2017-06-20”)
    end|str|結束時間  (格式：yyyy-MM-dd 
 例如：“2017-06-20”)


    - start 和 end 的組合如下  
        |start 類型 |end 類型 |說明 |
        |:--|:--|:--|
        |str |str |start 和 end 分別為指定的日期|
        |None |str |start 為 end 往前 365 天  |
        |str |None |end 為 start 往後 365 天 |
        |None |None |end 為 當前日期，start 往前 365 天 |


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，傳回資金流向數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，傳回錯誤描述</td>
        </tr>
    </table>

    * 資金流向數據格式如下：
        欄位|類型|說明
        :-|:-|:-
        in_flow|float|整體淨流入
        main_in_flow|float|主力大單淨流入  (僅歷史週期（日、周、月）有效)
        super_in_flow|float|特大單淨流入 
        big_in_flow|float|大單淨流入 
        mid_in_flow|float|中單淨流入 
        sml_in_flow|float|小單淨流入 
        capital_flow_item_time|str|開始時間  (格式：yyyy-MM-dd HH:mm:ss
精確到分鐘)
        last_valid_time|str|數據最後有效時間  (僅實時週期有效)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_capital_flow("HK.00700", period_type = PeriodType.INTRADAY)
if ret == RET_OK:
    print(data)
    print(data['in_flow'][0])    # 取第一條的淨流入的資金額度
    print(data['in_flow'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    last_valid_time       in_flow  ...  main_in_flow  capital_flow_item_time
0               N/A -1.857915e+08  ... -1.066828e+08     2021-06-08 00:00:00
..              ...           ...  ...           ...                     ...
245             N/A  2.179240e+09  ...  2.143345e+09     2022-06-08 00:00:00

[246 rows x 8 columns]
-185791500.0
[-185791500.0, -18315000.0, -672100100.0, -714394350.0, -698391950.0, -818886750.0, 304827400.0, 73026200.0, -2078217500.0, 
..                   ...           ...                    ...
2031460.0, 638067040.0, 622466600.0, -351788160.0, -328529240.0, 715415020.0, 76749700.0, 2179240320.0]
```

---



---

# 獲取資金分佈

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_capital_distribution(stock_code)`

* **介紹**

    獲取資金分佈

* **參數**
    參數|類型|說明
    :-|:-|:-
    stock_code|str|股票代號


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，傳回股票資金分佈數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，傳回錯誤描述</td>
        </tr>
    </table>

    * 資金分佈數據格式如下：
        欄位|類型|說明
        :-|:-|:-
        capital_in_super|float|流入資金額度，特大單
        capital_in_big|float|流入資金額度，大單
        capital_in_mid|float|流入資金額度，中單
        capital_in_small|float|流入資金額度，小單
        capital_out_super|float|流出資金額度，特大單
        capital_out_big|float|流出資金額度，大單
        capital_out_mid|float|流出資金額度，中單
        capital_out_small|float|流出資金額度，小單
        update_time|str|更新時間字串  (格式：yyyy-MM-dd HH:mm:ss)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_capital_distribution("HK.00700")
if ret == RET_OK:
    print(data)
    print(data['capital_in_big'][0])    # 取第一條的流入資金額度，大單
    print(data['capital_in_big'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
   capital_in_super  capital_in_big  ...  capital_out_small          update_time
0      2.261085e+09    2.141964e+09  ...       2.887413e+09  2022-06-08 15:59:59

[1 rows x 9 columns]
2141963720.0
[2141963720.0]
```

---



---

# 獲取股票所屬板塊

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_owner_plate(code_list)`

* **介紹**

    獲取單支或多支股票的所屬板塊資訊列表

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|股票代碼列表  (僅支援正股、指數list 中元素類型是 str)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回所屬板塊數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 所屬板塊數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|證券代碼
        name|str|股票名稱
        plate_code|str|板塊代碼
        plate_name|str|板塊名字
        plate_type|[Plate](./quote.md#2649)|板塊類型  (行業板塊或概念板塊)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

code_list = ['HK.00001']
ret, data = quote_ctx.get_owner_plate(code_list)
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['plate_code'].values.tolist())   # 板塊代碼轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
        code name          plate_code plate_name plate_type
0   HK.00001   長和  HK.HSI Constituent      恒指成份股      OTHER
..       ...  ...                 ...        ...        ...
8   HK.00001   長和           HK.BK1983    香港股票ADR      OTHER

[9 rows x 5 columns]
HK.00001
['HK.HSI Constituent', 'HK.GangGuTong', 'HK.BK1000', 'HK.BK1061', 'HK.BK1107', 'HK.BK1331', 'HK.BK1600', 'HK.BK1922', 'HK.BK1983']
```

---



---

# 獲取歷史 K 線

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`request_history_kline(code, start=None, end=None, ktype=KLType.K_DAY, autype=AuType.QFQ, fields=[KL_FIELD.ALL], max_count=1000, page_req_key=None, extended_time=False, session=Session.NONE)`

* **介紹**

    獲取歷史 K 線

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    start|str|開始時間  (格式：yyyy-MM-dd
例如：“2017-06-20”)
    end|str|結束時間  (格式：yyyy-MM-dd
例如：“2017-07-20”)
    ktype|[KLType](./quote.md#9701)|K 線類型
    autype|[AuType](./quote.md#6706)|復權類型
    fields|[KLFields](./quote.md#382)|需返回的欄位列表
    max_count|int|本次請求最大返回的 K 線根數  (- 傳 None 表示返回 start 和 end 之間所有的數據 
  - 注意：OpenD 接收到所有數據後才會下發給腳本，如果您要獲取的 K 線根數大於 1000 根，建議選擇分頁，防止出現超時)
    page_req_key|bytes|分頁請求  (如果 start 和 end 之間的 K 線根數多於 max_count：1. 首頁請求時應該傳 None 2. 後續頁請求時必須要傳入上次呼叫返回的參數 page_req_key)
    extended_time|bool|是否允許美股盤前盤後數據  (False：不允許True：允許)
    session|[Session](./quote.md#3103)|獲取美股分時段歷史K線  (- 僅用於獲取美股分時段歷史K線
  - 獲取美股歷史K線不支援入參OVERNIGHT
  - 最低OpenD版本要求：9.2.4207)


    * start 和 end 的組合如下
        Start 類型|End 類型|説明
        :-|:-|:-
        str|str|start 和 end 分別為指定的日期
        None|str|start 為 end 往前 365 天
        str|None|end 為 start 往後 365 天
        None|None|end 為當前日期，start 往前 365 天


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回歷史 K 線數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
        <tr>
            <td>page_req_key</td>
            <td>bytes</td>
            <td>下一頁請求的 key</td>
        </tr>
    </table>

    * 歷史 K 線數據格式如下:
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        time_key|str|K 線時間  (格式：yyyy-MM-dd HH:mm:ss
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        open|float|開盤價
        close|float|收盤價
        high|float|最高價
        low|float|最低價
        pe_ratio|float|市盈率  (該欄位為比例欄位，預設不展示 %)
        turnover_rate|float|換手率
        volume|int|成交量
        turnover|float|成交額
        change_rate|float|漲跌幅
        last_close|float|昨收價

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret, data, page_req_key = quote_ctx.request_history_kline('US.AAPL', start='2019-09-11', end='2019-09-18', max_count=5, session=Session.ALL)  # 每頁5個，請求第一頁
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['close'].values.tolist())   # 第一頁收盤價轉為 list
else:
    print('error:', data)
while page_req_key != None:  # 請求後面的所有結果
    print('*************************************')
    ret, data, page_req_key = quote_ctx.request_history_kline('US.AAPL', start='2019-09-11', end='2019-09-18', max_count=5, page_req_key=page_req_key, session=Session.ALL) # 請求翻頁後的數據
    if ret == RET_OK:
        print(data)
    else:
        print('error:', data)
print('All pages are finished!')
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
code  name             time_key       open      close       high        low  pe_ratio  turnover_rate    volume      turnover  change_rate  last_close
0  US.AAPL   蘋果  2019-09-11 00:00:00  52.631194  53.963447  53.992409  52.549135    18.773        0.01039  177158584  9.808562e+09     3.179511   52.300545
..       ...   ...                  ...        ...        ...        ...        ...       ...            ...       ...           ...          ...         ...
4  US.AAPL   蘋果  2019-09-17 00:00:00  53.087346  53.265945  53.294907  52.884612    18.530        0.00432   73545872  4.046314e+09     0.363802   53.072865

[5 rows x 13 columns]
US.AAPL
[53.9634465, 53.84156475, 52.7953125, 53.072865, 53.265945]
*************************************
       code  name             time_key       open      close       high        low  pe_ratio  turnover_rate   volume      turnover  change_rate  last_close
0  US.AAPL   蘋果  2019-09-18 00:00:00  53.352831  53.76554  53.784847  52.961844    18.704        0.00602  102572372  5.682068e+09     0.937925   53.265945
All pages are finished!
```

---



---

# 獲取復權因子

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_rehab(code)`

* **介紹**

    獲取股票的復權因子

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回復權數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 復權數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        ex_div_date|str|除權除息日
        split_base|float|拆股分子 (拆股比例=拆股分子/拆股分母)
        split_ert|float|拆股分母
        join_base|float|合股分子 (合股比例=合股分子/合股分母)
        join_ert|float|合股分母
        split_ratio|float|拆合股比例  (- 當公司出現合股，5股合1股時，合股分子=5，合股分母=1，拆合股比例=合股分子/合股分母=5/1- 當公司出現拆股，1股拆5股時，拆股分子=1，拆股分母=5，拆合股比例=拆股分子/拆股分母=1/5)
        per_cash_div|float|每股派現
        bonus_base|float|送股分子 (送股比例=送股分子/送股分母)
        bonus_ert|float|送股分母
        per_share_div_ratio|float|送股比例  (- 當公司出現送股，5股送1股時，送股分子=5，送股分母=1，送股比例=送股分子/送股分母=5/1)
        transfer_base|float|轉增股分子 (轉增股比例=轉增股分子/轉增股分母)
        transfer_ert|float|轉增股分母
        per_share_trans_ratio|float|轉增股比例  (- 當公司出現轉增股，10股轉增3股時，轉增股分子=10，轉增股分母=3，轉增股比例=轉增股分子/轉增股分母=10/3)
        allot_base|float|配股分子 (配股比例=配股分子/配股分母)
        allot_ert|float|配股分母
        allotment_ratio|float|配股比例  (- 當公司出現配股，5股配1股時，配股分子=5，配股分母=1，配股比例=配股分子/配股分母=5/1)
        allotment_price|float|配股價
        add_base|float|增發股分子 (增發股比例=增發股分子/增發股分母)
        add_ert|float|增發股分母
        stk_spo_ratio|float|增發比例  (- 當公司出現增發股，1股增發5股時，增發股分子=1，增發股分母=5，增發股比例=增發股分子/增發股分母=1/5)
        stk_spo_price|float|增發價格
        spin_off_base|float|分立分子
        spin_off_ert|float|分立分母
        spin_off_ratio|float|分立比例
        forward_adj_factorA|float|前復權因子 A
        forward_adj_factorB|float|前復權因子 B
        backward_adj_factorA|float|後復權因子 A
        backward_adj_factorB|float|後復權因子 B

        前復權價格 = 不復權價格 × 前復權因子 A + 前復權因子 B  
        後復權價格 = 不復權價格 × 後復權因子 A + 後復權因子 B

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_rehab("HK.00700")
if ret == RET_OK:
    print(data)
    print(data['ex_div_date'][0])    # 取第一條的除權除息日
    print(data['ex_div_date'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    ex_div_date  split_ratio  per_cash_div  per_share_div_ratio  per_share_trans_ratio  allotment_ratio  allotment_price  stk_spo_ratio  stk_spo_price  spin_off_base   spin_off_ert   spin_off_ratio   forward_adj_factorA  forward_adj_factorB  backward_adj_factorA  backward_adj_factorB
0   2005-04-19          NaN          0.07                  NaN                    NaN              NaN              NaN            NaN            NaN         NaN         NaN          NaN         1.0                -0.07                   1.0                  0.07
..         ...          ...           ...                  ...                    ...              ...              ...            ...            ...                  ...                  ...                   ...                   ...
15  2019-05-17          NaN          1.00                  NaN                    NaN              NaN              NaN            NaN            NaN         NaN        NaN        NaN           1.0                -1.00                   1.0                  1.00

[16 rows x 16 columns]
2005-04-19
['2005-04-19', '2006-05-15', '2007-05-09', '2008-05-06', '2009-05-06', '2010-05-05', '2011-05-03', '2012-05-18', '2013-05-20', '2014-05-15', '2014-05-16', '2015-05-15', '2016-05-20', '2017-05-19', '2018-05-18', '2019-05-17']
```

---



---

# 獲取期權鏈到期日

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_option_expiration_date(code, index_option_type=IndexOptionType.NORMAL)`

* **介紹**

    透過標的股票，查詢期權鏈的所有到期日。如需獲取完整期權鏈，請配合 [獲取期權鏈](../quote/get-option-chain.md) 介面使用。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|標的股票代碼
    index_option_type|[IndexOptionType](../quote/quote.md#1625)|指數期權類型  (僅對港股指數期權篩選有效，正股、ETFs、美股指數期權可忽略此參數)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回期權鏈到期日相關數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 期權鏈到期日數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        strike_time|str|期權鏈行權日  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        option_expiry_date_distance|int|距離到期日天數  (負數表示已過期)
        expiration_cycle|[ExpirationCycle](./quote.md#5835)|交割週期  (支援香港指數期權、美股指數期權)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret, data = quote_ctx.get_option_expiration_date(code='HK.00700')
if ret == RET_OK:
    print(data)
    print(data['strike_time'].values.tolist())  # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
  strike_time  option_expiry_date_distance expiration_cycle
0  2021-04-29                            4              N/A
1  2021-05-28                           33              N/A
2  2021-06-29                           65              N/A
3  2021-07-29                           95              N/A
4  2021-09-29                          157              N/A
5  2021-12-30                          249              N/A
6  2022-03-30                          339              N/A
['2021-04-29', '2021-05-28', '2021-06-29', '2021-07-29', '2021-09-29', '2021-12-30', '2022-03-30']
```

---



---

# 獲取期權鏈

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_option_chain(code, index_option_type=IndexOptionType.NORMAL, start=None, end=None, option_type=OptionType.ALL, option_cond_type=OptionCondType.ALL, data_filter=None)`

* **介紹**

    透過標的股票查詢期權鏈。此介面僅返回期權鏈的靜態資訊，如需獲取報價或擺盤等動態資訊，請用此介面返回的股票代碼，自行 [訂閱](../quote/sub.md) 所需要的類型。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|標的股票代碼
    index_option_type|[IndexOptionType](./quote.md#1625)|指數期權類型  (僅對港股指數期權篩選有效，正股、ETFs、美股指數期權可忽略此參數)
    start|str|開始日期，該日期指到期日  (例如：“2017-08-01”)
    end|str|結束日期（包括這一天），該日期指到期日  (例如：“2017-08-30”)
    option_type|[OptionType](./quote.md#7263)|期權看漲看跌類型  (預設為全部)
    option_cond_type|[OptionCondType](./quote.md#7482)|期權價內外類型  (預設為全部)
    data_filter|OptionDataFilter|數據篩選條件  (預設為不篩選)
    * start 和 end 的組合如下：  
        Start 類型|End 類型|説明
        :-|:-|:-
        str|str|start 和 end 分別為指定的日期
        None|str|start 為 end 往前 30 天
        str|None|end 為 start 往後30天
        None|None|start 為當前日期，end 往後 30 天

    * OptionDataFilter 欄位如下
        欄位|類型|説明
        :-|:-|:-
        implied_volatility_min|float|隱含波動率過濾起點  (精確到小數點後 0 位，超出部分會被捨棄該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        implied_volatility_max|float|隱含波動率過濾終點  (精確到小數點後 0 位，超出部分會被捨棄該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        delta_min|float|希臘值 Delta 過濾起點  (精確到小數點後 3 位，超出部分會被捨棄)
        delta_max|float|希臘值 Delta 過濾終點  (精確到小數點後 3 位，超出部分會被捨棄)
        gamma_min|float|希臘值 Gamma 過濾起點  (精確到小數點後 3 位，超出部分會被捨棄)
        gamma_max|float|希臘值 Gamma 過濾終點  (精確到小數點後 3 位，超出部分會被捨棄)
        vega_min|float|希臘值 Vega 過濾起點  (精確到小數點後 3 位，超出部分會被捨棄)
        vega_max|float|希臘值 Vega 過濾終點  (精確到小數點後 3 位，超出部分會被捨棄)
        theta_min|float|希臘值 Theta 過濾起點  (精確到小數點後 3 位，超出部分會被捨棄)
        theta_max|float|希臘值 Theta 過濾終點  (精確到小數點後 3 位，超出部分會被捨棄)
        rho_min|float|希臘值 Rho 過濾起點  (精確到小數點後 3 位，超出部分會被捨棄)
        rho_max|float|希臘值 Rho 過濾終點  (精確到小數點後 3 位，超出部分會被捨棄)
        net_open_interest_min|float|淨未平倉合約數過濾起點  (精確到小數點後 0 位，超出部分會被捨棄)
        net_open_interest_max|float|淨未平倉合約數過濾終點  (精確到小數點後 0 位，超出部分會被捨棄)
        open_interest_min|float|未平倉合約數過濾起點  (精確到小數點後 0 位，超出部分會被捨棄)
        open_interest_max|float|未平倉合約數過濾終點  (精確到小數點後 0 位，超出部分會被捨棄)
        vol_min|float|成交量過濾起點  (精確到小數點後 0 位，超出部分會被捨棄)
        vol_max|float|成交量過濾終點  (精確到小數點後 0 位，超出部分會被捨棄)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回期權鏈數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 期權鏈數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|名字
        lot_size|int|每手股數，期權表示每份合約股數  (指數期權無該欄位)
        stock_type|[SecurityType](./quote.md#9111)|股票類型
        option_type|[OptionType](./quote.md#7263)|期權類型
        stock_owner|str|標的股
        strike_time|str|行權日  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        strike_price|float|行權價
        suspension|bool|是否停牌  (True：停牌False：未停牌)
        stock_id|int|股票 ID
        index_option_type|[IndexOptionType](./quote.md#1625)|指數期權類型
        expiration_cycle|[ExpirationCycle](./quote.md#5835)|交割週期
        option_standard_type|[OptionStandardType](./quote.md#3451)|期權標準類型
        option_settlement_mode|[OptionSettlementMode](./quote.md#3944)|期權結算方式

* **Example**

```python
from futu import *
import time
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret1, data1 = quote_ctx.get_option_expiration_date(code='HK.00700')

filter1 = OptionDataFilter()
filter1.delta_min = 0
filter1.delta_max = 0.1

if ret1 == RET_OK:
    expiration_date_list = data1['strike_time'].values.tolist()
    for date in expiration_date_list:
        ret2, data2 = quote_ctx.get_option_chain(code='HK.00700', start=date, end=date, data_filter=filter1)
        if ret2 == RET_OK:
            print(data2)
            print(data2['code'][0])  # 取第一條的股票代碼
            print(data2['code'].values.tolist())  # 轉為 list
        else:
            print('error:', data2)
        time.sleep(3)
else:
    print('error:', data1)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
                     code                 name  lot_size stock_type option_type stock_owner strike_time  strike_price  suspension  stock_id index_option_type expiration_cycle option_standard_type option_settlement_mode
0     HK.TCH210429C350000   騰訊 210429 350.00 購       100       DRVT        CALL    HK.00700  2021-04-29         350.0       False  80235167               N/A        WEEK        STANDARD			N/A        
1     HK.TCH210429P350000   騰訊 210429 350.00 沽       100       DRVT         PUT    HK.00700  2021-04-29         350.0       False  80235247               N/A        WEEK        STANDARD			N/A        
2     HK.TCH210429C360000   騰訊 210429 360.00 購       100       DRVT        CALL    HK.00700  2021-04-29         360.0       False  80235163               N/A        WEEK        STANDARD			N/A        
3     HK.TCH210429P360000   騰訊 210429 360.00 沽       100       DRVT         PUT    HK.00700  2021-04-29         360.0       False  80235246               N/A        WEEK        STANDARD			N/A        
4     HK.TCH210429C370000   騰訊 210429 370.00 購       100       DRVT        CALL    HK.00700  2021-04-29         370.0       False  80235165               N/A        WEEK        STANDARD			N/A        
5     HK.TCH210429P370000   騰訊 210429 370.00 沽       100       DRVT         PUT    HK.00700  2021-04-29         370.0       False  80235248               N/A        WEEK        STANDARD			N/A        
HK.TCH210429C350000
['HK.TCH210429C350000', 'HK.TCH210429P350000', 'HK.TCH210429C360000', 'HK.TCH210429P360000', 'HK.TCH210429C370000', 'HK.TCH210429P370000']
...
                   code                name  lot_size stock_type option_type stock_owner strike_time  strike_price  suspension  stock_id index_option_type expiration_cycle option_standard_type option_settlement_mode
0   HK.TCH220330C490000  騰訊 220330 490.00 購       100       DRVT        CALL    HK.00700  2022-03-30         490.0       False  80235143               N/A        WEEK        STANDARD			N/A            
1   HK.TCH220330P490000  騰訊 220330 490.00 沽       100       DRVT         PUT    HK.00700  2022-03-30         490.0       False  80235193               N/A        WEEK        STANDARD			N/A            
2   HK.TCH220330C500000  騰訊 220330 500.00 購       100       DRVT        CALL    HK.00700  2022-03-30         500.0       False  80233887               N/A        WEEK        STANDARD			N/A            
3   HK.TCH220330P500000  騰訊 220330 500.00 沽       100       DRVT         PUT    HK.00700  2022-03-30         500.0       False  80233912               N/A        WEEK        STANDARD			N/A            
4   HK.TCH220330C510000  騰訊 220330 510.00 購       100       DRVT        CALL    HK.00700  2022-03-30         510.0       False  80233747               N/A        WEEK        STANDARD 			N/A           
5   HK.TCH220330P510000  騰訊 220330 510.00 沽       100       DRVT         PUT    HK.00700  2022-03-30         510.0       False  80233766               N/A        WEEK        STANDARD 			N/A           
HK.TCH220330C490000
['HK.TCH220330C490000', 'HK.TCH220330P490000', 'HK.TCH220330C500000', 'HK.TCH220330P500000', 'HK.TCH220330C510000', 'HK.TCH220330P510000']
```

---



---

# 篩選窩輪

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_warrant(stock_owner='', req=None)`

* **介紹**

    篩選窩輪（僅用於篩選香港市場的窩輪、牛熊證、界內證）

* **參數**
    參數|類型|説明
    :-|:-|:-
    stock_owner|str|所屬正股的股票代碼
    req|WarrantRequest|篩選參數組合
    * WarrantRequest 類型欄位説明如下： 
        欄位|類型|説明
        :-|:-|:-
        begin|int|數據起始點
        num|int|請求數據個數  (最大 200)
        sort_field|[SortField](./quote.md#382)|根據哪個欄位排序
        ascend|bool|排序方向  (True：升序False：降序)
        type_list|list|窩輪類型過濾列表  (list 中元素類型是 [WrtType](./quote.md#9275))
        issuer_list|list|發行人過濾列表  (list 中元素類型是 [Issuer](./quote.md#2860))
        maturity_time_min|str|到期日過濾範圍的開始時間
        maturity_time_max|str|到期日過濾範圍的結束時間
        ipo_period|[IpoPeriod](./quote.md#6566)|上市時段
        price_type|[PriceType](./quote.md#382)|價內/價外  (暫不支援界內證的界內外篩選)
        status|[WarrantStatus](./quote.md#1344)|窩輪狀態
        cur_price_min|float|最新價的過濾下限  (閉區間不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        cur_price_max|float|最新價的過濾上限  (閉區間不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        strike_price_min|float|行使價的過濾下限  (閉區間不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        strike_price_max|float|行使價的過濾上限  (閉區間不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        street_min|float|街貨佔比的過濾下限  (閉區間不傳代表下限為 -∞該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。精確到小數點後 3 位，超出部分會被捨棄)
        street_max|float|街貨佔比的過濾上限  (閉區間不傳代表上限為 +∞該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。精確到小數點後 3 位，超出部分會被捨棄)
        conversion_min|float|換股比率的過濾下限  (閉區間不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        conversion_max|float|換股比率的過濾上限  (閉區間不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        vol_min|int|成交量的過濾下限  (閉區間不傳代表下限為 -∞)
        vol_max|int|成交量的過濾上限  (閉區間不傳代表上限為 +∞)
        premium_min|float|溢價的過濾下限  (閉區間不傳代表下限為 -∞該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。精確到小數點後 3 位，超出部分會被捨棄)
        premium_max|float|溢價的過濾上限  (閉區間不傳代表上限為 +∞該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。精確到小數點後 3 位，超出部分會被捨棄)
        leverage_ratio_min|float|槓桿比率的過濾下限  (閉區間不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        leverage_ratio_max|float|槓桿比率的過濾上限  (閉區間不傳代表上限為 +∞)
        delta_min|float|對沖值的過濾下限  (閉區間僅認購認沽支援此欄位過濾不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        delta_max|float|對沖值的過濾上限  (閉區間僅認購認沽支援此欄位過濾不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        implied_min|float|引伸波幅的過濾下限  (閉區間僅認購認沽支援此欄位過濾不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        implied_max|float|引伸波幅的過濾上限  (閉區間僅認購認沽支援此欄位過濾不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        recovery_price_min|float|收回價的過濾下限  (閉區間僅牛熊證支援此欄位過濾不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        recovery_price_max|float|收回價的過濾上限  (閉區間僅牛熊證支援此欄位過濾不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)
        price_recovery_ratio_min|float|正股距收回價的過濾下限  (閉區間僅牛熊證支援此欄位過濾該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%不傳代表下限為 -∞精確到小數點後 3 位，超出部分會被捨棄)
        price_recovery_ratio_max|float|正股距收回價的過濾上限  (閉區間僅牛熊證支援此欄位過濾該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%不傳代表上限為 +∞精確到小數點後 3 位，超出部分會被捨棄)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>tuple</td>
            <td>當 ret == RET_OK，返回窩輪數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 窩輪數據組成如下：
        欄位|類型|説明
        :-|:-|:-
        warrant_data_list|pd.DataFrame|篩選後的窩輪數據
        last_page|bool|是否是最後一頁  (True：是最後一頁False：不是最後一頁)
        all_count|int|篩選結果中的窩輪總數量

        - warrant_data_list 返回的 pd dataframe 數據格式：
            欄位|類型|説明
            :-|:-|:-
            stock|str|窩輪代碼
            stock_owner|str|所屬正股
            type|[WrtType](./quote.md#9275)|窩輪類型
            issuer|[Issuer](./quote.md#2860)|發行人
            maturity_time|str|到期日  (格式：yyyy-MM-dd)
            list_time|str|上市時間  (格式：yyyy-MM-dd)
            last_trade_time|str|最後交易日  (格式：yyyy-MM-dd)
            recovery_price|float|收回價  (僅牛熊證支援此欄位)
            conversion_ratio|float|換股比率
            lot_size|int|每手數量
            strike_price|float|行使價
            last_close_price|float|昨收價
            name|str|名稱
            cur_price|float|當前價
            price_change_val|float|漲跌額
            status|[WarrantStatus](./quote.md#1344)|窩輪狀態
            bid_price|float|買入價
            ask_price|float|賣出價
            bid_vol|int|買量
            ask_vol|int|賣量
            volume|int|成交量
            turnover|float|成交額
            score|float|綜合評分
            premium|float|溢價  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            break_even_point|float|打和點
            leverage|float|槓桿比率  (單位：倍)
            ipop|float|價內/價外  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            price_recovery_ratio|float|正股距收回價  (僅牛熊證支援此欄位該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            conversion_price|float|換股價
            street_rate|float|街貨佔比  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            street_vol|int|街貨量
            amplitude|float|振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            issue_size|int|發行量
            high_price|float|最高價
            low_price|float|最低價
            implied_volatility|float|引伸波幅  (僅認購認沽支援此欄位)
            delta|float|對沖值  (僅認購認沽支援此欄位)
            effective_leverage|float|有效槓桿 (僅認購認沽支援此欄位)
            upper_strike_price|float|上限價  (僅界內證支援此欄位)
            lower_strike_price|float|下限價  (僅界內證支援此欄位)
            inline_price_status|[PriceType](./quote.md#382)|界內界外  (僅界內證支援此欄位)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

req = WarrantRequest()
req.sort_field = SortField.TURNOVER
req.type_list = WrtType.CALL
req.cur_price_min = 0.1
req.cur_price_max = 0.2
ret, ls = quote_ctx.get_warrant("HK.00700", req)
if ret == RET_OK:  # 先判斷介面返回是否正常，再取數據
    warrant_data_list, last_page, all_count = ls
    print(len(warrant_data_list), all_count, warrant_data_list)
    print(warrant_data_list['stock'][0])    # 取第一條的窩輪代碼
    print(warrant_data_list['stock'].values.tolist())   # 轉為 list
else:
    print('error: ', ls)
    
req = WarrantRequest()
req.sort_field = SortField.TURNOVER
req.issuer_list = ['UB','CS','BI']
ret, ls = quote_ctx.get_warrant(Market.HK, req)
if ret == RET_OK: 
    warrant_data_list, last_page, all_count = ls
    print(len(warrant_data_list), all_count, warrant_data_list)
else:
    print('error: ', ls)

quote_ctx.close()  # 所有介面結尾加上這條 close，防止連線條數用盡
```

* **Output**

```python
2 2 
    stock        name stock_owner  type issuer maturity_time   list_time last_trade_time  recovery_price  conversion_ratio  lot_size  strike_price  last_close_price  cur_price  price_change_val  change_rate  status  bid_price  ask_price   bid_vol  ask_vol    volume   turnover   score  premium  break_even_point  leverage    ipop  price_recovery_ratio  conversion_price  street_rate  street_vol  amplitude  issue_size  high_price  low_price  implied_volatility  delta  effective_leverage  list_timestamp  last_trade_timestamp  maturity_timestamp  upper_strike_price  lower_strike_price  inline_price_status
0   HK.20306  騰訊麥銀零乙購A.C    HK.00700  CALL     MB    2020-12-01  2019-06-27      2020-11-25             NaN              50.0      5000        588.88             0.188      0.188             0.000     0.000000  NORMAL      0.000      0.188         0     10000           0          0.0   0.198    2.008            598.28    62.393  -0.404                   NaN              9.40        4.400     1584000      0.000    36000000       0.000      0.000              31.751  0.479              29.886    1.561565e+09          1.606234e+09        1.606752e+09                 NaN                 NaN                  NaN
1   HK.16545  騰訊法興一二購B.C    HK.00700  CALL     SG    2021-02-26  2020-07-14      2021-02-22             NaN             100.0     10000        700.00             0.147      0.144            -0.003    -2.040816  NORMAL      0.141      0.144  28000000  28000000           0          0.0  81.506   21.807            714.40    40.729 -16.214                   NaN             14.40        1.420     2130000      0.000   150000000       0.000      0.000              40.643  0.226               9.204    1.594656e+09          1.613923e+09        1.614269e+09                 NaN                 NaN                  NaN
HK.20306
['HK.20306', 'HK.16545']

200 358
    stock        name stock_owner    type issuer maturity_time   list_time last_trade_time  recovery_price  conversion_ratio  lot_size  strike_price  last_close_price  cur_price  price_change_val  change_rate      status  bid_price  ask_price   bid_vol   ask_vol  volume  turnover   score  premium  break_even_point  leverage     ipop  price_recovery_ratio  conversion_price  street_rate  street_vol  amplitude  issue_size  high_price  low_price  implied_volatility  delta  effective_leverage  list_timestamp  last_trade_timestamp  maturity_timestamp  upper_strike_price  lower_strike_price inline_price_status
0    HK.19839  平安瑞銀零乙購A.C    HK.02318    CALL     UB    2020-12-31  2017-12-11      2020-12-24             NaN             100.0     50000         83.88             0.057      0.046            -0.011   -19.298246      NORMAL      0.043      0.046  30000000  30000000       0       0.0  39.641    1.642            88.480    18.923    3.779                   NaN             4.600         1.25     6250000        0.0   500000000         0.0        0.0              25.129  0.692              13.094    1.512922e+09          1.608739e+09        1.609344e+09                 NaN                 NaN                 NaN
1    HK.20084  平安中銀零乙購A.C    HK.02318    CALL     BI    2020-12-31  2017-12-19      2020-12-24             NaN             100.0     50000         83.88             0.059      0.050            -0.009   -15.254237      NORMAL      0.044      0.050  10000000  10000000       0       0.0   0.064    2.102            88.880    17.410    3.779                   NaN             5.000         0.07      350000        0.0   500000000         0.0        0.0              29.174  0.672              11.699    1.513613e+09          1.608739e+09        1.609344e+09                 NaN                 NaN                 NaN
......
198  HK.56886  恒指瑞銀三一牛F.C   HK.800000    BULL     UB    2023-01-30  2020-03-24      2023-01-27         21200.0           20000.0     10000      21100.00             0.230      0.232             0.002     0.869565      NORMAL      0.232      0.233  30000000  30000000       0       0.0  46.627   -2.884         25740.000     5.712   25.613             25.021179          4640.000         0.01       40000        0.0   400000000         0.0        0.0                 NaN    NaN               5.712    1.584979e+09          1.674749e+09        1.675008e+09                 NaN                 NaN                 NaN
199  HK.56895  小米瑞銀零乙牛D.C    HK.01810    BULL     UB    2020-12-30  2020-03-24      2020-12-29             8.0              10.0      2000          7.60             2.010      1.930            -0.080    -3.980100      NORMAL      1.910      1.930   6000000   6000000       0       0.0   0.040    0.938            26.900     1.380  250.657            233.125000            19.300         0.10       60000        0.0    60000000         0.0        0.0                 NaN    NaN               1.380    1.584979e+09          1.609171e+09        1.609258e+09                 NaN                 NaN                 NaN

```

---



---

# 獲取窩輪和期貨列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_referencestock_list(code, reference_type)`

* **介紹**

    獲取證券的關聯數據，如：獲取正股相關窩輪、獲取期貨相關合約

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|證券代碼
    reference_type|[SecurityReferenceType](./quote.md#817)|要獲得的相關數據


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回證券的關聯數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 證券的關聯數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|證券代碼
        lot_size|int|每手股數，期貨表示合約乘數
        stock_type|[SecurityType](./quote.md#9111)|證券類型
        stock_name|str|證券名字
        list_time|str|上市時間  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        wrt_valid|bool|是否是窩輪  (若為 True，下面 wrt 開頭的欄位有效)
        wrt_type|[WrtType](./quote.md#9275)|窩輪類型
        wrt_code|str|所屬正股
        future_valid|bool|是否是期貨  (若為 True，以下 future 開頭的欄位有效)
        future_main_contract|bool|是否主連合約  (期貨特有欄位)
        future_last_trade_time|str|最後交易時間  (期貨特有欄位主連，當月，下月等無該欄位)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

# 獲取正股相關的窩輪
ret, data = quote_ctx.get_referencestock_list('HK.00700', SecurityReferenceType.WARRANT)
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
print('******************************************')
# 港期相關合約
ret, data = quote_ctx.get_referencestock_list('HK.A50main', SecurityReferenceType.FUTURE)
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
        code  lot_size stock_type stock_name   list_time  wrt_valid wrt_type  wrt_code  future_valid  future_main_contract  future_last_trade_time
0     HK.24719      1000    WARRANT    騰訊東亞九四沽A  2018-07-20       True      PUT  HK.00700         False                   NaN                     NaN
..         ...       ...        ...                ...       ...        ...       ...       ...           ...                   ...                    ...
1617  HK.63402     10000    WARRANT    騰訊高盛一八牛Y  2020-11-26       True     BULL  HK.00700         False                   NaN                     NaN

[1618 rows x 11 columns]
HK.24719
['HK.24719', 'HK.27886', 'HK.28621', 'HK.14339', 'HK.27952', 'HK.18693', 'HK.20306', 'HK.53635', 'HK.47269', 'HK.27227', 
...        ...       ...        ...        ...         ...        ...      ...       ... 
'HK.63402']
******************************************
        code  lot_size stock_type         stock_name list_time  wrt_valid  wrt_type  wrt_code  future_valid  future_main_contract future_last_trade_time
0  HK.A50main      5000     FUTURE      安碩富時 A50 ETF主連(2012)                False       NaN       NaN          True                  True                       
..         ...       ...        ...                ...       ...        ...       ...       ...           ...                   ...                    ...
5  HK.A502106      5000     FUTURE      安碩富時 A50 ETF2106                False       NaN       NaN          True                 False             2021-06-29

[6 rows x 11 columns]
HK.A50main
['HK.A50main', 'HK.A502011', 'HK.A502012', 'HK.A502101', 'HK.A502103', 'HK.A502106']
```

---



---

# 獲取期貨合約資料

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_future_info(code_list)`

* **介紹**

    獲取期貨合約資料

* **參數**
    參數|類型|說明
    :-|:-|:-
    code_list|list|股票代號列表  (list 中元素類型是 str)


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，傳回期貨合約資料數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，傳回錯誤描述</td>
        </tr>
    </table>

    * 期貨合約資料數據格式如下：
        欄位|類型|說明
        :-|:-|:-
        code|str|股票代號
        name|str|股票名稱
        owner|str|標的
        exchange|str|交易所
        type|str|合約類型
        size|float|合約規模
        size_unit|str|合約規模單位
        price_currency|str|報價貨幣
        price_unit|str|報價單位
        min_change|float|最小變動
        min_change_unit|str|最小變動的單位 (該欄位已廢棄)
        trade_time|str|交易時間
        time_zone|str|時區
        last_trade_time|str|最後交易時間  (主連，當月，下月等期貨沒有該欄位)
        exchange_format_url|str|交易所規格連結 url
        origin_code|str|實際合約代號

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_future_info(["HK.MPImain", "HK.HAImain"])
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代號
    print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code      name       owner exchange  type     size size_unit price_currency price_unit  min_change min_change_unit                        trade_time time_zone last_trade_time                                exchange_format_url           origin_code
0  HK.MPImain   內房期貨主連  恒生中國內地地產指數      港交所  股指期貨     50.0    指數點×港元             港元        指數點        0.50                (09:15 - 12:00), (13:00 - 16:30)       CCT                  https://sc.hkex.com.hk/TuniS/www.hkex.com.hk/P...           HK.MPI2112
1  HK.HAImain   海通證券期貨主連    HK.06837      港交所  股票期貨  10000.0         股             港元      每股/港元        0.01                   (09:30 - 12:00), (13:00 - 16:00)       CCT                  https://sc.hkex.com.hk/TuniS/www.hkex.com.hk/P...           HK.HAI2112
HK.MPImain
['HK.MPImain', 'HK.HAImain']
```

---



---

# 條件選股

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_stock_filter(market, filter_list, plate_code=None, begin=0, num=200)`

* **介紹**

    條件選股

* **參數**
    參數|類型|説明
    :-|:-|:-
    market|[Market](./quote.md#8744)|市場標識  (不區分滬股和深股，傳入滬股或者深股都會返回滬深市場的股票)
    filter_list|list|篩選條件的列表  (參考下面的表格，列表中元素類型為 SimpleFilter 或 AccumulateFilter 或 FinancialFilter)
    plate_code|str|板塊代碼
    begin|int|數據起始點
    num|int|請求數據個數
    * SimpleFilter 對象相關參數如下：  

        欄位|類型|説明
        :-|:-|:-
        stock_field|[StockField](./quote.md#9850)|簡單屬性
        filter_min|float|區間下限  (閉區間不傳預設為 -∞)
        filter_max|float|區間上限  (閉區間不傳預設為 +∞)
        is_no_filter|bool|該欄位是否不需要篩選  (True：不篩選False：篩選不傳預設不篩選)
        sort|[SortDir](./quote.md#5471)|排序方向  (不傳預設為不排序)

    * AccumulateFilter 對象相關參數如下：

        欄位|類型|説明
        :-|:-|:-
        stock_field|[StockField](./quote.md#9843)|累積屬性
        filter_min|float|區間下限  (閉區間不傳預設為 -∞)
        filter_max|float|區間上限  (閉區間不傳預設為 +∞)
        is_no_filter|bool|該欄位是否不需要篩選  (True：不篩選False：篩選不傳預設不篩選)
        sort|[SortDir](./quote.md#5471)|排序方向  (不傳預設為不排序)
        days|int|所篩選的數據的累計天數

    * FinancialFilter 對象相關參數如下：

        欄位|類型|説明
        :-|:-|:-
        stock_field|[StockField](./quote.md#9553)|財務屬性
        filter_min|float|區間下限  (閉區間不傳預設為 -∞)
        filter_max|float|區間上限  (閉區間不傳預設為 +∞)
        is_no_filter|bool|該欄位是否不需要篩選  (True：不篩選False：篩選不傳預設不篩選)
        sort|[SortDir](./quote.md#5471)|排序方向  (不傳預設為不排序)
        quarter|[FinancialQuarter](./quote.md#9850)|財報累積時間

    * CustomIndicatorFilter 對象相關參數如下：

        欄位|類型|説明
        :-|:-|:-
        stock_field1|[StockField](./quote.md#9850)|自定義技術指標屬性
        stock_field1_para|list|自定義技術指標屬性參數  (根據指標類型進行傳參：1. MA：[平均移動週期] 2.EMA：[指數移動平均週期] 3.RSI：[RSI 指標週期] 4.MACD：[快速平均線值, 慢速平均線值, DIF值] 5.BOLL：[均線週期, 偏移值] 6.KDJ：[RSV 週期, K 值計算週期, D 值計算週期]) 
        relative_position|[RelativePosition](./quote.md#954)|相對位置
        stock_field2|[StockField](./quote.md#9850)|自定義技術指標屬性
        stock_field2_para|list|自定義技術指標屬性參數  (根據指標類型進行傳參：1. MA：[平均移動週期] 2.EMA：[指數移動平均週期] 3.RSI：[RSI 指標週期] 4.MACD：[快速平均線值, 慢速平均線值, DIF值] 5.BOLL：[均線週期, 偏移值] 6.KDJ：[RSV 週期, K 值計算週期, D 值計算週期]) 
        value|float|自定義數值  (當 stock_field2 在 [StockField](./quote.md#9850) 中選擇自定義數值時，value 為必傳參數) 
        ktype|[KLType](./quote.md#9701)|K線類型 KLType   (僅支援K_60M，K_DAY，K_WEEK，K_MON 四種時間週期)
        consecutive_period|int|篩選連續週期（consecutive_period）都符合條件的數據  (填寫範圍為[1,12]) 
        is_no_filter|bool|該欄位是否不需要篩選  (True：不篩選False：篩選不傳預設不篩選)
 
    * PatternFilter 對象相關參數如下：

        欄位|類型|説明
        :-|:-|:-
        stock_field|[StockField](./quote.md#2923)|形態技術指標屬性
        ktype|[KLType](./quote.md#9701)|K線類型 KLType （僅支援K_60M，K_DAY，K_WEEK，K_MON 四種時間週期）
        consecutive_period|int|篩選連續週期（consecutive_period）都符合條件的數據  (填寫範圍為[1,12]) 
        is_no_filter|bool|該欄位是否不需要篩選  (True：不篩選False：篩選不傳預設不篩選)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>tuple</td>
            <td>當 ret == RET_OK，返回選股數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 選股數據元組組成如下：
        欄位|類型|説明
        :-|:-|:-
        last_page|bool|是否是最後一頁
        all_count|int|列表總數量
        stock_list|list|選股數據  (list 中元素類型是 FilterStockData)
        
        - FilterStockData 類型的欄位格式：

            欄位|類型|説明
            :-|:-|:-
            stock_code|str|股票代碼
            stock_name|str|股票名字
            cur_price|float|最新價
            cur_price_to_highest_52weeks_ratio|float|(現價 - 52周最高)/52周最高  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            cur_price_to_lowest_52weeks_ratio|float|(現價 - 52周最低)/52周最低  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            high_price_to_highest_52weeks_ratio|float|(今日最高 - 52周最高)/52周最高  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            low_price_to_lowest_52weeks_ratio|float|(今日最低 - 52周最低)/52周最低  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            volume_ratio|float|量比
            bid_ask_ratio|float|委比  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            lot_price|float|每手價格
            market_val|float|市值
            pe_annual|float|市盈率
            pe_ttm|float|市盈率 TTM
            pb_rate|float|市淨率
            change_rate_5min|float|五分鐘價格漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            change_rate_begin_year|float|年初至今價格漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            ps_ttm|float|市銷率 TTM  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            pcf_ttm|float|市現率 TTM  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            total_share|float|總股數  (單位：股)
            float_share|float|流通股數  (單位：股)
            float_market_val|float|流通市值  (單位：元)
            change_rate|float|漲跌幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            amplitude|float|振幅  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            volume|float|日均成交量
            turnover|float|日均成交額
            turnover_rate|float|換手率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            net_profit|float|淨利潤
            net_profix_growth|float|淨利潤增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            sum_of_business|float|營業收入
            sum_of_business_growth|float|營業同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            net_profit_rate|float|淨利率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            gross_profit_rate|float|毛利率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            debt_asset_rate|float|資產負債率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            return_on_equity_rate|float|淨資產收益率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            roic|float|投入資本回報率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            roa_ttm|float|資產回報率 TTM  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報)
            ebit_ttm|float|息税前利潤 TTM  (單位：元。僅適用於年報)
            ebitda|float|税息折舊及攤銷前利潤  (單位：元)
            operating_margin_ttm|float|營業利潤率 TTM  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報)
            ebit_margin|float|EBIT 利潤率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            ebitda_margin|float|EBITDA 利潤率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            financial_cost_rate|float|財務成本率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            operating_profit_ttm|float|營業利潤 TTM  (單位：元。僅適用於年報)
            shareholder_net_profit_ttm|float|歸屬於母公司的淨利潤  (單位：元。僅適用於年報)
            net_profit_cash_cover_ttm|float|盈利中的現金收入比例  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報)
            current_ratio|float|流動比率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            quick_ratio|float|速動比率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            current_asset_ratio|float|流動資產率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            current_debt_ratio|float|流動負債率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            equity_multiplier|float|權益乘數 
            property_ratio|float|產權比率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            cash_and_cash_equivalents|float|現金和現金等價  (單位：元)
            total_asset_turnover|float|總資產週轉率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            fixed_asset_turnover|float|固定資產週轉率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            inventory_turnover|float|存貨週轉率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            operating_cash_flow_ttm|float|經營活動現金流 TTM   (單位：元。僅適用於年報)
            accounts_receivable|float|應收賬款淨額  (單位：元)
            ebit_growth_rate|float|EBIT 同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            operating_profit_growth_rate|float|營業利潤同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            total_assets_growth_rate|float|總資產同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            profit_to_shareholders_growth_rate|float|歸母淨利潤同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            profit_before_tax_growth_rate|float|總利潤同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            eps_growth_rate|float|EPS 同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            roe_growth_rate|float|ROE 同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            roic_growth_rate|float|ROIC 同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            nocf_growth_rate|float|經營現金流同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            nocf_per_share_growth_rate|float|每股經營現金流同比增長率  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            operating_revenue_cash_cover|float|經營現金收入比  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            operating_profit_to_total_profit|float|營業利潤佔比  (該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
            basic_eps|float|基本每股收益  (單位：元)
            diluted_eps|float|稀釋每股收益  (單位：元)
            nocf_per_share|float|每股經營現金淨流量  (單位：元)
            price|float|最新價格
            ma|float|簡單均線  (根據 MA 參數返回具體的數值)
            ma5|float|5日簡單均線
            ma10|float|10日簡單均線
            ma20|float|20日簡單均線
            ma30|float|30日簡單均線
            ma60|float|60日簡單均線
            ma120|float|120日簡單均線
            ma250|float|250日簡單均線
            rsi|float|RSI的值  (根據 RSI 參數返回具體的數值，RSI 預設參數為12)
            ema|float|指數移動均線  (根據 EMA 參數返回具體的數值) 
            ema5|float|5日指數移動均線 
            ema10|float|10日指數移動均線
            ema20|float|20日指數移動均線
            ema30|float|30日指數移動均線
            ema60|float|60日指數移動均線
            ema120|float|120日指數移動均線
            ema250|float|250日指數移動均線
            kdj_k|float|KDJ 指標的 K 值  (根據 KDJ 參數返回具體的數值，KDJ 預設參數為[9,3,3]) 
            kdj_d|float|KDJ 指標的 D 值  (根據 KDJ 參數返回具體的數值，KDJ 預設參數為[9,3,3]) 
            kdj_j|float|KDJ 指標的 J 值  (根據 KDJ 參數返回具體的數值，KDJ 預設參數為[9,3,3]) 
            macd_diff|float|MACD 指標的 DIFF 值  (根據 MACD 參數返回具體的數值，MACD 預設參數為[12,26,9]) 
            macd_dea|float|MACD 指標的 DEA 值  (根據 MACD 參數返回具體的數值，MACD 預設參數為[12,26,9]) 
            macd|float|MACD 指標的 MACD 值  (根據 MACD 參數返回具體的數值，MACD 預設參數為[12,26,9]) 
            boll_upper|float|BOLL 指標的 UPPER 值  (根據 BOLL 參數返回具體的數值，BOLL 預設參數為[20.2]) 
            boll_middler|float|BOLL 指標的 MIDDLER 值  (根據 BOLL 參數返回具體的數值，BOLL 預設參數為[20.2])
            boll_lower|float|BOLL 指標的 LOWER 值  (根據 BOLL 參數返回具體的數值，BOLL 預設參數為[20.2])


* **Example**

```python
from futu import *
import time

quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
simple_filter = SimpleFilter()
simple_filter.filter_min = 2
simple_filter.filter_max = 1000
simple_filter.stock_field = StockField.CUR_PRICE
simple_filter.is_no_filter = False
# simple_filter.sort = SortDir.ASCEND

financial_filter = FinancialFilter()
financial_filter.filter_min = 0.5
financial_filter.filter_max = 50
financial_filter.stock_field = StockField.CURRENT_RATIO
financial_filter.is_no_filter = False
financial_filter.sort = SortDir.ASCEND
financial_filter.quarter = FinancialQuarter.ANNUAL

custom_filter = CustomIndicatorFilter()
custom_filter.ktype = KLType.K_DAY
custom_filter.stock_field1 = StockField.KDJ_K
custom_filter.stock_field1_para = [10,4,4]
custom_filter.stock_field2 = StockField.KDJ_K
custom_filter.stock_field2_para = [9,3,3]
custom_filter.relative_position = RelativePosition.MORE
custom_filter.is_no_filter = False

nBegin = 0
last_page = False
ret_list = list()
while not last_page:
    nBegin += len(ret_list)
    ret, ls = quote_ctx.get_stock_filter(market=Market.HK, filter_list=[simple_filter, financial_filter, custom_filter], begin=nBegin)  # 對香港市場的股票做簡單、財務和指標篩選
    if ret == RET_OK:
        last_page, all_count, ret_list = ls
        print('all count = ', all_count)
        for item in ret_list:
            print(item.stock_code)  # 取股票代碼
            print(item.stock_name)  # 取股票名稱
            print(item[simple_filter])   # 取 simple_filter 對應的變量值
            print(item[financial_filter])   # 取 financial_filter 對應的變量值
            print(item[custom_filter])  # 獲取 custom_filter 的數值
    else:
        print('error: ', ls)
    time.sleep(3)  # 加入時間間隔，避免觸發限頻

quote_ctx.close()  # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
39 39 [ stock_code:HK.08103  stock_name:HMVOD視頻  cur_price:2.69  current_ratio(annual):4.413 ,  stock_code:HK.00376  stock_name:雲鋒金融  cur_price:2.96  current_ratio(annual):12.585 ,  stock_code:HK.09995  stock_name:榮昌生物-B  cur_price:92.65  current_ratio(annual):16.054 ,  stock_code:HK.80737  stock_name:灣區發展-R  cur_price:2.8  current_ratio(annual):17.249 ,  stock_code:HK.00737  stock_name:灣區發展  cur_price:3.25  current_ratio(annual):17.249 ,  stock_code:HK.03939  stock_name:萬國國際礦業  cur_price:2.22  current_ratio(annual):17.323 ,  stock_code:HK.01055  stock_name:中國南方航空股份  cur_price:5.17  current_ratio(annual):17.529 ,  stock_code:HK.02638  stock_name:港燈-SS  cur_price:7.68  current_ratio(annual):21.255 ,  stock_code:HK.00670  stock_name:中國東方航空股份  cur_price:3.53  current_ratio(annual):25.194 ,  stock_code:HK.01952  stock_name:雲頂新耀-B  cur_price:69.5  current_ratio(annual):26.029 ,  stock_code:HK.00089  stock_name:大生地產  cur_price:4.22  current_ratio(annual):26.914 ,  stock_code:HK.00728  stock_name:中國電信  cur_price:2.81  current_ratio(annual):27.651 ,  stock_code:HK.01372  stock_name:比速科技  cur_price:5.1  current_ratio(annual):28.303 ,  stock_code:HK.00753  stock_name:中國國航  cur_price:6.38  current_ratio(annual):31.828 ,  stock_code:HK.01997  stock_name:九龍倉置業  cur_price:43.75  current_ratio(annual):33.239 ,  stock_code:HK.02158  stock_name:醫渡科技  cur_price:39.0  current_ratio(annual):34.046 ,  stock_code:HK.02588  stock_name:中銀航空租賃  cur_price:77.0  current_ratio(annual):34.531 ,  stock_code:HK.01330  stock_name:綠色動力環保  cur_price:3.36  current_ratio(annual):35.028 ,  stock_code:HK.01525  stock_name:建橋教育  cur_price:6.28  current_ratio(annual):36.989 ,  stock_code:HK.09908  stock_name:嘉興燃氣  cur_price:10.02  current_ratio(annual):37.848 ,  stock_code:HK.06078  stock_name:海吉亞醫療  cur_price:49.8  current_ratio(annual):39.0 ,  stock_code:HK.01071  stock_name:華電國際電力股份  cur_price:2.16  current_ratio(annual):39.507 ,  stock_code:HK.00357  stock_name:美蘭空港  cur_price:34.15  current_ratio(annual):39.514 ,  stock_code:HK.00762  stock_name:中國聯通  cur_price:5.15  current_ratio(annual):40.74 ,  stock_code:HK.01787  stock_name:山東黃金  cur_price:15.56  current_ratio(annual):41.604 ,  stock_code:HK.00902  stock_name:華能國際電力股份  cur_price:2.66  current_ratio(annual):42.919 ,  stock_code:HK.00934  stock_name:中石化冠德  cur_price:2.96  current_ratio(annual):43.361 ,  stock_code:HK.01117  stock_name:現代牧業  cur_price:2.3  current_ratio(annual):45.037 ,  stock_code:HK.00177  stock_name:江蘇寧滬高速公路  cur_price:8.78  current_ratio(annual):45.93 ,  stock_code:HK.01379  stock_name:温嶺工量刃具  cur_price:5.71  current_ratio(annual):46.774 ,  stock_code:HK.01876  stock_name:百威亞太  cur_price:22.5  current_ratio(annual):46.917 ,  stock_code:HK.01907  stock_name:中國旭陽集團  cur_price:4.38  current_ratio(annual):47.129 ,  stock_code:HK.02160  stock_name:心通醫療-B  cur_price:15.54  current_ratio(annual):47.384 ,  stock_code:HK.00293  stock_name:國泰航空  cur_price:7.1  current_ratio(annual):47.983 ,  stock_code:HK.00694  stock_name:北京首都機場股份  cur_price:6.34  current_ratio(annual):47.985 ,  stock_code:HK.09922  stock_name:九毛九  cur_price:26.65  current_ratio(annual):48.278 ,  stock_code:HK.01083  stock_name:港華燃氣  cur_price:3.39  current_ratio(annual):49.2 ,  stock_code:HK.00291  stock_name:華潤啤酒  cur_price:58.0  current_ratio(annual):49.229 ,  stock_code:HK.00306  stock_name:冠忠巴士集團  cur_price:2.29  current_ratio(annual):49.769 ]
HK.08103
HMVOD視頻
2.69
2.69
4.413
...
HK.00306
冠忠巴士集團
2.29
2.29
49.769
```

---



---

# 獲取板塊內股票列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_plate_stock(plate_code, sort_field=SortField.CODE, ascend=True)`

* **介紹**

    獲取指定板塊內的股票列表，獲取股指的成分股

* **參數**
    參數|類型|説明
    :-|:-|:-
    plate_code|str|板塊代碼  (先利用 [獲取板塊列表](../quote/get-plate-list.md) 獲取板塊代碼例如：“SH.BK0001”，“SH.BK0002”)
    sort_field|[SortField](./quote.md#382)|排序欄位
    ascend|bool|排序方向  (True：升序False：降序)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回板塊股票數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 板塊股票數據
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        lot_size|int|每手股數，期貨表示合約乘數
        stock_name|str|股票名稱
        stock_type|[SecurityType](./quote.md#9111)|股票類型
        list_time|str|上市時間  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        stock_id|int|股票 ID
        main_contract|bool|是否主連合約  (期貨特有欄位)
        last_trade_time|str|最後交易時間  (期貨特有欄位主連，當月，下月等期貨沒有該欄位)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_plate_stock('HK.BK1001')
if ret == RET_OK:
    print(data)
    print(data['stock_name'][0])    # 取第一條的股票名稱
    print(data['stock_name'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code  lot_size stock_name  stock_owner  stock_child_type stock_type   list_time        stock_id  main_contract last_trade_time
0   HK.00462      4000       天然乳品          NaN               NaN      STOCK  2005-06-10  55589761712590          False                
..       ...       ...        ...          ...               ...        ...         ...             ...            ...             ...
9   HK.06186      1000       中國飛鶴          NaN               NaN      STOCK  2019-11-13  78159814858794          False               

[10 rows x 10 columns]
天然乳品
['天然乳品', '現代牧業', '雅士利國際', '原生態牧業', '中國聖牧', '中地乳業', '莊園牧場', '澳優', '蒙牛乳業', '中國飛鶴']
```

---



---

# 獲取板塊列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_plate_list(market, plate_class)`

* **介紹**

    獲取板塊列表

* **參數**
    參數|類型|説明
    :-|:-|:-
    market|[Market](./quote.md#8744)|市場標識  (注意：這裡不區分滬和深，輸入滬或者深都會返回滬深市場的子板塊)
    plate_class|[Plate](./quote.md#8056)|板塊分類


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回板塊列表數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 板塊列表數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|板塊代碼
        plate_name|str|板塊名字
        plate_id|str|板塊 ID

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_plate_list(Market.HK, Plate.CONCEPT)
if ret == RET_OK:
    print(data)
    print(data['plate_name'][0])    # 取第一條的板塊名稱
    print(data['plate_name'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code plate_name plate_id
0   HK.BK1000      做空集合股   BK1000
..        ...        ...      ...
77  HK.BK1999       殯葬概念   BK1999

[78 rows x 3 columns]
做空集合股
['做空集合股', '阿里概念股', '雄安概念股', '蘋果概念', '一帶一路', '5G概念', '夜店股', '粵港澳大灣區', '特斯拉概念股', '啤酒', '疑似財技股', '體育用品', '稀土概念', '人民幣升值概念', '抗疫概念', '新股與次新股', '騰訊概念', '雲辦公', 'SaaS概念', '在線教育', '汽車經銷商', '挪威政府全球養老基金持倉', '武漢本地概念股', '核電', '內地醫藥股', '化妝美容股', '科網股', '公用股', '石油股', '電訊設備', '電力股', '手遊股', '嬰兒及小童用品股', '百貨業股', '收租股', '港口運輸股', '電信股', '環保', '煤炭股', '汽車股', '電池', '物流', '內地物業管理股', '農業股', '黃金股', '奢侈品股', '電力設備股', '連鎖快餐店', '重型機械股', '食品股', '內險股', '紙業股', '水務股', '奶製品股', '光伏太陽能股', '內房股', '內地教育股', '家電股', '風電股', '藍籌地產股', '內銀股', '航空股', '石化股', '建材水泥股', '中資券商股', '高鐵基建股', '燃氣股', '公路及鐵路股', '鋼鐵金屬股', '華為概念', 'OLED概念', '工業大麻', '香港本地股', '香港零售股', '區塊鏈', '豬肉概念', '節假日概念', '殯葬概念']
```

---



---

# 獲取靜態數據

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_stock_basicinfo(market, stock_type=SecurityType.STOCK, code_list=None)`

* **介紹**

    獲取靜態數據

* **參數**
    參數|類型|説明
    :-|:-|:-
    market|[Market](./quote.md#8744)|市場類型
    stock_type|[SecurityType](./quote.md#9111)|股票類型，但不支援傳入 SecurityType.DRVT
    code_list|list|股票列表  (- 預設為 None，代表獲取全市場股票的靜態資訊
  - 若傳入股票列表，只返回指定股票的資訊
  - list 中元素類型是 str)
    注：當 market 和 code_list 同時存在時，會忽略 market，僅對 code_list 進行查詢。


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回股票靜態數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 股票靜態數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        lot_size|int|每手股數，期權表示每份合約股數  (指數期權無該欄位)，期貨表示合約乘數
        stock_type|[SecurityType](./quote.md#9111)|股票類型
        stock_child_type|[WrtType](./quote.md#9275)|窩輪子類型
        stock_owner|str|窩輪所屬正股的代碼，或期權標的股的代碼
        option_type|[OptionType](./quote.md#7263)|期權類型
        strike_time|str|期權行權日  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間)
        strike_price|float|期權行權價
        suspension|bool|期權是否停牌  (True：停牌False：未停牌)
        listing_date|str|上市時間  (此欄位停止維護，不建議使用
格式：yyyy-MM-dd)
        stock_id|int|股票 ID
        delisting|bool|是否退市
        index_option_type|str|指數期權類型
        main_contract|bool|是否主連合約
        last_trade_time|str|最後交易時間  (主連，當月，下月等期貨沒有該欄位)
        exchange_type|[ExchType](./quote.html#6592)|所屬交易所

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
ret, data = quote_ctx.get_stock_basicinfo(Market.HK, SecurityType.STOCK)
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
print('******************************************')
ret, data = quote_ctx.get_stock_basicinfo(Market.HK, SecurityType.STOCK, ['HK.06998', 'HK.00700'])
if ret == RET_OK:
    print(data)
    print(data['name'][0])  # 取第一條的股票名稱
    print(data['name'].values.tolist())  # 轉為 list
else:
    print('error:', data)
quote_ctx.close()  # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
        code             name  lot_size stock_type stock_child_type stock_owner option_type strike_time strike_price suspension listing_date        stock_id  delisting index_option_type  main_contract last_trade_time exchange_type
0      HK.00001               長和       500      STOCK              N/A                     N/A                      N/A        N/A   2015-03-18   4440996184065      False               N/A          False                  HK_MAINBOARD  
...         ...              ...       ...        ...              ...         ...         ...         ...          ...        ...          ...             ...        ...               ...            ...             ...
2592   HK.09979     綠城管理控股      1000      STOCK              N/A                                              N/A        N/A   2020-07-10  79203491915515      False               N/A          False                  HK_MAINBOARD                

[2593 rows x 16 columns]
******************************************
        code            name  lot_size stock_type stock_child_type stock_owner option_type strike_time strike_price suspension listing_date        stock_id  delisting index_option_type  main_contract last_trade_time exchange_type
0  HK.06998     嘉和生物-B       500      STOCK              N/A                                              N/A        N/A   2020-10-07  79572859099990      False               N/A          False                  HK_MAINBOARD                
1  HK.00700     騰訊控股         100      STOCK              N/A                                              N/A        N/A   2004-06-16  54047868453564      False               N/A          False                  HK_MAINBOARD               
嘉和生物-B
['嘉和生物-B', '騰訊控股']
```

---



---

# 獲取 IPO 資訊

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>


`get_ipo_list(market)`

* **介紹**

    獲取指定市場的 IPO 資訊

* **參數**
    參數|類型|説明
    :-|:-|:-
    market|[Market](./quote.md#8744)|市場標識  (注意：這裡不區分滬和深，輸入滬或者深都會返回滬深市場的股票)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回 IPO 數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * IPO 數據
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        list_time|str|上市日期，美股是預計上市日期 (格式：yyyy-MM-dd)
        list_timestamp|float|上市日期時間戳記，美股是預計上市日期時間戳記
        apply_code|str|申購代碼（A 股適用）
        issue_size|int|發行總數（A 股適用）；發行量（美股適用）
        online_issue_size|int|網上發行量（A 股適用）
        apply_upper_limit|int|申購上限（A 股適用）
        apply_limit_market_value|int|頂格申購需配市值（A 股適用）
        is_estimate_ipo_price|bool|是否預估發行價（A 股適用）
        ipo_price|float|發行價  (預估值會因為募集資金、發行數量、發行費用等數據變動而變動，僅供參考。實際數據公佈後會第一時間更新)（A 股適用）
        industry_pe_rate|float|行業市盈率（A 股適用）
        is_estimate_winning_ratio|bool|是否預估中籤率（A 股適用）
        winning_ratio|float|中籤率  (- 預估值會因為募集資金、發行數量、發行費用等數據變動而變動，僅供參考。實際數據公佈後會第一時間更新
  - 該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)（A 股適用）
        issue_pe_rate|float|發行市盈率（A 股適用）
        apply_time|str|申購日期字串 (格式：yyyy-MM-dd)（A 股適用）
        apply_timestamp|float|申購日期時間戳記（A 股適用）
        winning_time|str|公佈中籤日期字串 (格式：yyyy-MM-dd)（A 股適用）
        winning_timestamp|float|公佈中籤日期時間戳記（A 股適用）
        is_has_won|bool|是否已經公佈中籤號（A 股適用）
        winning_num_data|str|中籤號（A 股適用）  (格式類似：末"五"位數：12345，12346末"六"位數：123456)
        ipo_price_min|float|最低發售價（港股適用）；最低發行價（美股適用）
        ipo_price_max|float|最高發售價（港股適用）；最高發行價（美股適用）
        list_price|float|上市價（港股適用）
        lot_size|int|每手股數
        entrance_price|float|入場費（港股適用）
        is_subscribe_status|bool|是否為認購狀態  (True：認購中False：待上市)
        apply_end_time|str|截止認購日期字串 (格式：yyyy-MM-dd)（港股適用）
        apply_end_timestamp|float|截止認購日期時間戳記|因需處理認購手續，富途認購截止時間會早於交易所公佈的日期（港股適用）

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_ipo_list(Market.HK)
if ret == RET_OK:
    print(data)
    print(data['code'][0])    # 取第一條的股票代碼
    print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code      name   list_time  list_timestamp apply_code issue_size online_issue_size apply_upper_limit apply_limit_market_value is_estimate_ipo_price ipo_price industry_pe_rate is_estimate_winning_ratio winning_ratio issue_pe_rate apply_time apply_timestamp winning_time winning_timestamp is_has_won winning_num_data  ipo_price_min  ipo_price_max  list_price  lot_size  entrance_price  is_subscribe_status apply_end_time  apply_end_timestamp
0  HK.06666  恒大物業  2020-12-02    1.606838e+09        N/A        N/A               N/A               N/A                      N/A                   N/A       N/A              N/A                       N/A           N/A           N/A        N/A             N/A          N/A               N/A        N/A              N/A          8.500           9.75         0.0       500         4924.12                 True     2020-11-26         1.606352e+09
1  HK.02110  裕勤控股  2020-12-07    1.607270e+09        N/A        N/A               N/A               N/A                      N/A                   N/A       N/A              N/A                       N/A           N/A           N/A        N/A             N/A          N/A               N/A        N/A              N/A          0.225           0.27         0.0     10000         2727.21                 True     2020-11-27         1.606439e+09
HK.06666
['HK.06666', 'HK.02110']
```

---



---

# 獲取全域市場狀態

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>


`get_global_state()`  

* **介紹**

    獲取全域狀態


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>dict</td>
            <td>當 ret == RET_OK 時，返回全域狀態</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 全域狀態字典格式如下：
        欄位|類型|説明
        :-|:-|:-
        market_sz|[MarketState](./quote.md#8754)|深圳市場狀態
        market_sh|[MarketState](./quote.md#8754)|上海市場狀態
        market_hk|[MarketState](./quote.md#8754)|香港市場狀態
        market_hkfuture|[MarketState](./quote.md#8754)|香港期貨市場狀態  (不同品種的交易時間存在差異，建議使用 [get_market_state](../quote/get-market-state.md) 介面獲取指定品種的市場狀態)
        market_usfuture|[MarketState](./quote.md#8754)|美國期貨市場狀態  (不同品種的交易時間存在差異，建議使用 [get_market_state](../quote/get-market-state.md) 介面獲取指定品種的市場狀態)
        market_us|[MarketState](./quote.md#8754)|美國市場狀態  (不同品種的交易時間存在差異，建議使用 [get_market_state](../quote/get-market-state.md) 介面獲取指定品種的市場狀態)
        market_sgfuture|[MarketState](./quote.md#8754)|新加坡期貨市場狀態  (不同品種的交易時間存在差異，建議使用 [get_market_state](../quote/get-market-state.md) 介面獲取指定品種的市場狀態)
        market_jpfuture|[MarketState](./quote.md#8754)|日本期貨市場狀態
        server_ver|str|OpenD 版本號
        trd_logined|bool|True：已登入交易伺服器，False：未登入交易伺服器
        qot_logined|bool|True：已登入行情伺服器，False：未登入行情伺服器
        timestamp|str|當前格林威治時間戳記  (單位：秒)
        local_timestamp|float|OpenD 運行機器的當前時間戳記  (單位：秒)
        program_status_type|[ProgramStatusType](../ftapi/common.md#3053)|當前狀態
        program_status_desc|str|額外描述
    

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
print(quote_ctx.get_global_state())
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
(0, {'market_sz': 'MORNING', 'market_us': 'AFTER_HOURS_END', 'market_sh': 'MORNING', 'market_hk': 'MORNING', 'market_hkfuture': 'FUTURE_DAY_OPEN', 'market_usfuture': 'FUTURE_OPEN', 'market_sgfuture': 'FUTURE_DAY_OPEN', 'market_jpfuture': 'FUTURE_DAY_OPEN', 'server_ver': '504', 'trd_logined': True, 'timestamp': '1620962951', 'qot_logined': True, 'local_timestamp': 1620962951.047128, 'program_status_type': 'READY', 'program_status_desc': ''})
```

---



---

# 獲取交易日曆

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`request_trading_days(market=None, start=None, end=None, code=None)`

* **介紹**

    請求指定市場 / 指定標的的交易日曆。  
    注意：該交易日是透過自然日剔除週末和節假日得到，未剔除臨時休市數據。  

* **參數**
    參數|類型|説明
    :-|:-|:-
    market|[TradeDateMarket](./quote.md#4509)|市場類型
    start|str|起始日期  (格式：yyyy-MM-dd
例如：“2018-01-01”)
    end|str|結束日期  (格式：yyyy-MM-dd
例如：“2018-01-01”)
    code| str | 股票代碼
    注：當 market 和 code 同時存在時，會忽略 market，僅對 code 進行查詢。

    * start 和 end 的組合如下
        Start 類型|End 類型|説明
        :-|:-|:-
        str|str|start 和 end 分別為指定的日期
        None|str|start 為 end 往前 365 天
        str|None|end 為 start 往後 365 天
        None|None|start 為往前 365 天，end 當前日期


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>list</td>
            <td>當 ret == RET_OK 時，返回交易日數據。list 中元素類型為 dict</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 交易日數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        time|str|時間 (格式：yyyy-MM-dd)
        trade_date_type|[TradeDateType](./quote.md#4509)|交易日類型

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.request_trading_days(market=TradeDateMarket.HK, start='2020-04-01', end='2020-04-10')
if ret == RET_OK:
    print('HK market calendar:', data)
else:
    print('error:', data)
print('******************************************')
ret, data = quote_ctx.request_trading_days(start='2020-04-01', end='2020-04-10', code='HK.00700')
if ret == RET_OK:
    print('HK.00700 calendar:', data)
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
HK market calendar: [{'time': '2020-04-01', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-02', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-03', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-06', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-07', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-08', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-09', 'trade_date_type': 'WHOLE'}]
******************************************
HK.00700 calendar: [{'time': '2020-04-01', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-02', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-03', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-06', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-07', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-08', 'trade_date_type': 'WHOLE'}, {'time': '2020-04-09', 'trade_date_type': 'WHOLE'}]
```

---



---

# 獲取歷史 K 線額度使用明細

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_history_kl_quota(get_detail=False)`

* **介紹**

    獲取歷史 K 線額度使用明細

* **參數**
    參數|類型|説明
    :-|:-|:-
    get_detail|bool|是否返回下載 / 讀取歷史 K 線的詳細紀錄  (True：返回False：不返回)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>tuple</td>
            <td>當 ret == RET_OK，返回歷史 K 線額度數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 歷史 K 線額度數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        used_quota|int|已用額度  (即當前週期內已經下載過多少隻股票)
        remain_quota|int|剩餘額度
        detail_list|list|下載 / 讀取歷史 K 線的詳細紀錄，含股票代碼和下載 / 讀取時間  (list 中元素類型是 dict)

        - detail_list 數據列格式如下
            欄位|類型|説明
            :-|:-|:-
            code|str|股票代碼
            name|str|股票名稱
            request_time|str|最後一次下載 / 讀取的時間字串  (格式：yyyy-MM-dd HH:mm:ss)

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_history_kl_quota(get_detail=True)  # 設定 true 代表需要返回詳細的下載 / 讀取歷史 K 線的記錄
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
(2, 98, {'code': 'HK.00123', 'name': '越秀地產', 'request_time': '2023-06-20 19:59:00'}, {'code': 'HK.00700', 'name': '騰訊控股', 'request_time': '2023-07-19 17:48:16'}])
```

---



---

# 設定到價提醒

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`set_price_reminder(code, op, key=None, reminder_type=None, reminder_freq=None, value=None, note=None, reminder_session_list=NONE)`

* **介紹**

    新增、刪除、修改、啟用、禁用指定股票的到價提醒

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    op|[SetPriceReminderOp](./quote.md#1977)|操作類型
    key|int|標識，新增和刪除全部的情況不需要填
    reminder_type|[PriceReminderType](./quote.md#28)|到價提醒的類型，刪除、啟用、禁用的情況下會忽略該入參
    reminder_freq|[PriceReminderFreq](./quote.md#9426)|到價提醒的頻率，刪除、啟用、禁用的情況下會忽略該入參
    value|float|提醒值，刪除、啟用、禁用的情況下會忽略該入參  (精確到小數點後 3 位，超出部分會被捨棄)
    note|str|用戶設定的備註，僅支援 20 個以內的中文字元，刪除、啟用、禁用的情況下會忽略該入參
    reminder_session_list|list|美股到價提醒的時段列表，刪除、啟用、禁用的情況下會忽略該入參  (- list中元素類型是[PriceReminderMarketStatus](./quote.md#7145)
  - 美股預設到價提醒時段：盤中+盤前盤後)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">key</td>
            <td>int</td>
            <td>當 ret == RET_OK 時，返回操作的到價提醒 key</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>


* **Example**

```python
from futu import *
import time
class PriceReminderTest(PriceReminderHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, content = super(PriceReminderTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("PriceReminderTest: error, msg: %s" % content)
            return RET_ERROR, content
        print("PriceReminderTest ", content) # PriceReminderTest 自己的處理邏輯
        return RET_OK, content
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = PriceReminderTest()
quote_ctx.set_handler(handler)
ret, data = quote_ctx.get_market_snapshot(['US.AAPL'])
if ret == RET_OK:
    bid_price = data['bid_price'][0]  # 獲取實時買一價
    ask_price = data['ask_price'][0]  # 獲取實時賣一價
    # 設定當AAPL全時段賣一價低於（ask_price-1）時提醒
    ret_ask, ask_data = quote_ctx.set_price_reminder(code='US.AAPL', op=SetPriceReminderOp.ADD, key=None, reminder_type=PriceReminderType.ASK_PRICE_DOWN, reminder_freq=PriceReminderFreq.ALWAYS, value=(ask_price-1), note='123', reminder_session_list=[PriceReminderMarketStatus.US_PRE, PriceReminderMarketStatus.OPEN, PriceReminderMarketStatus.US_AFTER, PriceReminderMarketStatus.US_OVERNIGHT])
    if ret_ask == RET_OK:
        print('賣一價低於（ask_price-1）時提醒設定成功：', ask_data)
    else:
        print('error:', ask_data)
    # 設定當AAPL全時段買一價高於（bid_price+1）時提醒
    ret_bid, bid_data = quote_ctx.set_price_reminder(code='US.AAPL', op=SetPriceReminderOp.ADD, key=None, reminder_type=PriceReminderType.BID_PRICE_UP, reminder_freq=PriceReminderFreq.ALWAYS, value=(bid_price+1), note='456', reminder_session_list=[PriceReminderMarketStatus.US_PRE, PriceReminderMarketStatus.OPEN, PriceReminderMarketStatus.US_AFTER, PriceReminderMarketStatus.US_OVERNIGHT])
    if ret_bid == RET_OK:
        print('買一價高於（bid_price+1）時提醒設定成功：', bid_data)
    else:
        print('error:', bid_data)
time.sleep(15)
quote_ctx.close()
```

* **Output**

```python
賣一價低於（ask_price-1）時提醒設定成功： 1744022257023211123
買一價高於（bid_price+1）時提醒設定成功： 1744022257052794489
```

---



---

# 獲取到價提醒列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_price_reminder(code=None, market=None)`

* **介紹**

    獲取對指定股票 / 指定市場設定的到價提醒列表

* **參數**
    參數|類型|説明
    :-|:-|:-
    code|str|股票代碼
    market|[Market](./quote.md#8744)|市場類型  (輸入滬股市場和深股市場，都會認為是 A 股市場) 
    注：code 和 market 都存在的情況下，code 優先。


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回到價提醒數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 到價提醒數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        key|int|標識，用於修改到價提醒
        reminder_type|[PriceReminderType](./quote.md#28)|到價提醒的類型
        reminder_freq|[PriceReminderFreq](./quote.md#9426)|到價提醒的頻率
        value|float|提醒值
        enable|bool|是否啓用
        note|str|備註  (僅支援 20 個以內的中文字符) 
        reminder_session_list|list|美股到價提醒時段列表  (list中元素類型是[PriceReminderMarketStatus](./quote.md#7145))

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_price_reminder(code='US.AAPL')
if ret == RET_OK:
    print(data)
    print(data['key'].values.tolist())   # 轉為 list
else:
    print('error:', data)
print('******************************************')
ret, data = quote_ctx.get_price_reminder(code=None, market=Market.US)
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果到價提醒列表不為空
        print(data['code'][0])    # 取第一條的股票代碼
        print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
code name                  key   reminder_type reminder_freq   value  enable note                   reminder_session_list
0  US.AAPL   蘋果  1744021708234288125    BID_PRICE_UP        ALWAYS  184.37    True  456                              [US_AFTER]
1  US.AAPL   蘋果  1744022257052794489    BID_PRICE_UP        ALWAYS  185.50    True  456  [OPEN, US_PRE, US_AFTER, US_OVERNIGHT]
2  US.AAPL   蘋果  1744021708211891867  ASK_PRICE_DOWN        ALWAYS  182.54    True  123                              [US_AFTER]
3  US.AAPL   蘋果  1744022257023211123  ASK_PRICE_DOWN        ALWAYS  183.70    True  123  [OPEN, US_PRE, US_AFTER, US_OVERNIGHT]
[1744021708234288125, 1744022257052794489, 1744021708211891867, 1744022257023211123]
******************************************
      code name                  key   reminder_type reminder_freq   value  enable note                   reminder_session_list
0  US.AAPL   蘋果  1744021708234288125    BID_PRICE_UP        ALWAYS  184.37    True  456                              [US_AFTER]
1  US.AAPL   蘋果  1744022257052794489    BID_PRICE_UP        ALWAYS  185.50    True  456  [OPEN, US_PRE, US_AFTER, US_OVERNIGHT]
2  US.AAPL   蘋果  1744021708211891867  ASK_PRICE_DOWN        ALWAYS  182.54    True  123                              [US_AFTER]
3  US.AAPL   蘋果  1744022257023211123  ASK_PRICE_DOWN        ALWAYS  183.70    True  123  [OPEN, US_PRE, US_AFTER, US_OVERNIGHT]
4  US.NVDA  英偉達  1739697581665326308      PRICE_DOWN        ALWAYS  102.00    True       [OPEN, US_PRE, US_AFTER, US_OVERNIGHT]
US.AAPL
['US.AAPL', 'US.AAPL', 'US.AAPL', 'US.AAPL', 'US.NVDA']
```

---



---

# 獲取自選股列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_user_security(group_name)`

* **介紹**

    獲取指定分組的自選股列表

* **參數**

    參數|類型|説明
    :-|:-|:-
    group_name|str|需要查詢的自選股分組名稱


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回自選股數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 自選股數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|名字
        lot_size|int|每手股數，期權表示每份合約股數，期貨表示合約乘數
        stock_type|[SecurityType](./quote.md#9111)|股票類型
        stock_child_type|[WrtType](./quote.md#9275)|窩輪子類型
        stock_owner|str|窩輪所屬正股的代碼，或期權標的股的代碼
        option_type|[OptionType](./quote.md#7263)|期權類型
        strike_time|str|期權行權日  (格式：yyyy-MM-dd
港股和 A 股市場預設是北京時間，美股市場預設是美東時間) 
        strike_price|float|期權行權價
        suspension|bool|期權是否停牌  (True：停牌) 
        listing_date|str|上市時間  (格式：yyyy-MM-dd)
        stock_id|int|股票 ID
        delisting|bool|是否退市
        main_contract|bool|是否主連合約
        last_trade_time|str|最後交易時間  (主連，當月，下月等期貨沒有此欄位) 

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_user_security("A")
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果自選股列表不為空
        print(data['code'][0])    # 取第一條的股票代碼
        print(data['code'].values.tolist())   # 轉為 list
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
    code    name  lot_size stock_type stock_child_type stock_owner option_type strike_time strike_price suspension listing_date        stock_id  delisting  main_contract last_trade_time
0  HK.HSImain  恒指期貨主連        50     FUTURE              N/A                                              N/A        N/A                     71000662      False           True                
1  HK.00700    騰訊控股       100      STOCK              N/A                                              N/A        N/A   2004-06-16  54047868453564      False          False                
HK.HSImain
['HK.HSImain', 'HK.00700']
```

---



---

# 獲取自選股分組

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_user_security_group(group_type = UserSecurityGroupType.ALL)`

* **介紹**

    獲取自選股分組列表

* **參數**
    參數|類型|説明
    :-|:-|:-
    group_type|[UserSecurityGroupType](./quote.md#6296)|分組類型


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK，返回自選股分組數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 自選股分組數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        group_name|str|分組名
        group_type|[UserSecurityGroupType](./quote.md#6296)|分組類型

* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.get_user_security_group(group_type = UserSecurityGroupType.ALL)
if ret == RET_OK:
    print(data)
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連接，防止連接條數用盡
```

* **Output**

```python
        group_name group_type
0          期權     SYSTEM
..         ...        ...
12          C     CUSTOM

[13 rows x 2 columns]
```

---



---

# 修改自選股列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`modify_user_security(group_name, op, code_list)`

* **介紹**

    修改指定分組的自選股列表（系統分組不支援修改）

* **參數**
    參數|類型|説明
    :-|:-|:-
    group_name|str|需要修改的自選股分組名稱
    op|[ModifyUserSecurityOp](./quote.md#1977)|操作類型
    code_list|list|股票列表  (list 中元素類型是 str) 


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">msg</td>
            <td rowspan="2">str</td>
            <td>當 ret == RET_OK，返回“success”</td>
        </tr>
        <tr>
            <td>當 ret != RET_OK，msg 返回錯誤描述</td>
        </tr>
    </table>


* **Example**

```python
from futu import *
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)

ret, data = quote_ctx.modify_user_security("A", ModifyUserSecurityOp.ADD, ['HK.00700'])
if ret == RET_OK:
    print(data) # 返回 success
else:
    print('error:', data)
quote_ctx.close() # 結束後記得關閉當條連線，防止連線條數用盡
```

* **Output**

```python
success
```

---



---

# 到價提醒回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    到價提醒通知回呼，非同步處理已設定到價提醒的通知推送。  
    在收到實時到價提醒通知推送後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  


* **參數**

    參數|類型|説明
    :-|:-|:-
    rsp_pb|Qot_UpdatePriceReminder_pb2.Response|衍生類別中不需要直接處理該參數


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面呼叫結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>dict</td>
            <td>當 ret == RET_OK，返回到價提醒</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 到價提醒
        欄位|類型|説明
        :-|:-|:-
        code|str|股票代碼
        name|str|股票名稱
        price|float|當前價格
        change_rate|str|當前漲跌幅
        market_status|[PriceReminderMarketStatus](./quote.md#7145)|觸發的時間段
        content|str|到價提醒文字內容
        note|str|備註  (僅支援 20 個以內的中文字元) 
        key|int|到價提醒標識
        reminder_type|[PriceReminderType](./quote.md#28)|到價提醒的類型
        set_value|float|用戶設定的提醒值
        cur_value|float|提醒觸發時的值

* **Example**

```python
import time
from futu import *

class PriceReminderTest(PriceReminderHandlerBase):
    def on_recv_rsp(self, rsp_pb):
        ret_code, content = super(PriceReminderTest,self).on_recv_rsp(rsp_pb)
        if ret_code != RET_OK:
            print("PriceReminderTest: error, msg: %s" % content)
            return RET_ERROR, content
        print("PriceReminderTest ", content) # PriceReminderTest 自己的處理邏輯
        return RET_OK, content
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
handler = PriceReminderTest()
quote_ctx.set_handler(handler)  # 設定到價提醒通知回呼
time.sleep(15)  # 設定腳本接收 OpenD 的推送持續時間為15秒
quote_ctx.close()   # 關閉當條連線，OpenD 會在1分鐘後自動取消相應股票相應類型的訂閲
```

* **Output**

```python
PriceReminderTest  {'code': 'US.AAPL', 'name': '蘋果', 'price': 185.750, 'change_rate': 0.11, 'market_status': 'US_PRE', 'content': '買一價高於185.500', 'note': '', 'key': 1744022257052794489, 'reminder_type': 'BID_PRICE_UP', 'set_value': 185.500, 'cur_value': 185.750}
```

---



---

# 行情定義

## 累積過濾屬性

**AccumulateField**

```protobuf
enum AccumulateField
{
    AccumulateField_Unknown = 0; // 未知
    AccumulateField_ChangeRate = 1; // 漲跌幅（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-10.2,20.4]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    AccumulateField_Amplitude = 2; // 振幅（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[0.5,20.6]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    AccumulateField_Volume = 3; // 日均成交量（精確到小數點後 0 位，超出部分會被捨棄）例如填寫[2000,70000]值區間
    AccumulateField_Turnover = 4; // 日均成交額（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[1400,890000]值區間
    AccumulateField_TurnoverRate = 5; // 換手率（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[2,30]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
}
```

## 資產類別

**AssetClass**

```protobuf
enum AssetClass
{
    AssetClass_Unknow = 0; //未知
    AssetClass_Stock = 1; //股票
    AssetClass_Bond = 2; //債券
    AssetClass_Commodity = 3; //商品
    AssetClass_CurrencyMarket = 4; //貨幣市場
    AssetClass_Future = 5; //期貨
    AssetClass_Swap = 6; //掉期（互換）
}
```

## 公司行動

**CompanyAct**

```protobuf
enum CompanyAct
{
    CompanyAct_None = 0; //無
    CompanyAct_Split = 1; //拆股
    CompanyAct_Join = 2; //合股
    CompanyAct_Bonus = 4; //送股
    CompanyAct_Transfer = 8; //轉贈股
    CompanyAct_Allot = 16; //配股    
    CompanyAct_Add = 32; //增發股
    CompanyAct_Dividend = 64; //現金分紅
    CompanyAct_SPDividend = 128; //特別股息    
}
```

## 暗盤狀態

**DarkStatus**

```protobuf
enum DarkStatus
{
    DarkStatus_None = 0; //無暗盤交易
    DarkStatus_Trading = 1; //暗盤交易中
    DarkStatus_End = 2; //暗盤交易結束
}
```

## 財務過濾屬性

**FinancialField**

```protobuf
enum FinancialField
{
    // 基礎財務屬性
    FinancialField_Unknown = 0; // 未知
    FinancialField_NetProfit = 1; // 淨利潤 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[100000000,2500000000]值區間
    FinancialField_NetProfitGrowth = 2; // 淨利潤增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-10,300]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_SumOfBusiness = 3; // 營業收入 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[100000000,6400000000]值區間
    FinancialField_SumOfBusinessGrowth = 4; // 營收按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-5,200]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_NetProfitRate = 5; // 淨利率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[10,113]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_GrossProfitRate = 6; // 毛利率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[4,65]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_DebtAssetsRate = 7; // 資產負債率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[5,470]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ReturnOnEquityRate = 8; // 淨資產收益率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[20,230]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    
    // 盈利能力屬性
    FinancialField_ROIC = 9; // 投入資本回報率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ROATTM = 10; // 資產回報率(TTM) （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報。）
    FinancialField_EBITTTM = 11; // 息稅前利潤(TTM) （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間（單位：元。僅適用於年報。）
    FinancialField_EBITDA = 12; // 稅息折舊及攤銷前利潤 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間（單位：元）
    FinancialField_OperatingMarginTTM = 13; // 營業利潤率(TTM) （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報。）
    FinancialField_EBITMargin = 14; // EBIT 利潤率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_EBITDAMargin  = 15; // EBITDA 利潤率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_FinancialCostRate = 16; // 財務成本率（精確到小數點後 3 位，超出部分會被捨棄） 例如填寫 [1.0,10.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_OperatingProfitTTM  = 17; // 營業利潤(TTM) （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間 （單位：元。僅適用於年報。）
    FinancialField_ShareholderNetProfitTTM = 18; // 歸屬於母公司的淨利潤 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間 （單位：元。僅適用於年報。）
    FinancialField_NetProfitCashCoverTTM = 19; // 盈利中的現金收入比例 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,60.0] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%。僅適用於年報。）
    
    // 償債能力屬性
    FinancialField_CurrentRatio = 20; // 流動比率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [100,250] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_QuickRatio = 21; // 速動比率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [100,250] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）    
    
    // 清債能力屬性
    FinancialField_CurrentAssetRatio = 22; // 流動資產率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [10,100] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_CurrentDebtRatio = 23; // 流動負債率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [10,100] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_EquityMultiplier = 24; // 權益乘數 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [100,180] 值區間
    FinancialField_PropertyRatio = 25; // 產權比率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [50,100] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%） 
    FinancialField_CashAndCashEquivalents = 26; // 現金和現金等價 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間（單位：元）    
    
    // 營運能力屬性
    FinancialField_TotalAssetTurnover = 27; // 總資產週轉率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [50,100] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_FixedAssetTurnover = 28; // 固定資產週轉率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [50,100] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_InventoryTurnover = 29; // 存貨週轉率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [50,100] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_OperatingCashFlowTTM = 30; // 經營活動現金流(TTM) （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間（單位：元。僅適用於年報。）
    FinancialField_AccountsReceivable = 31; // 應收帳款淨額 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間 例如填寫 [1000000000,1000000000] 值區間 （單位：元）    
    
    // 成長能力屬性
    FinancialField_EBITGrowthRate = 32 ; // EBIT 按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_OperatingProfitGrowthRate = 33; // 營業利潤按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_TotalAssetsGrowthRate = 34; // 總資產按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ProfitToShareholdersGrowthRate = 35; // 歸母淨利潤按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ProfitBeforeTaxGrowthRate = 36; // 總利潤按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_EPSGrowthRate = 37; // EPS 按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ROEGrowthRate = 38; // ROE 按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_ROICGrowthRate = 39; // ROIC 按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_NOCFGrowthRate = 40; // 經營現金流按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_NOCFPerShareGrowthRate = 41; // 每股經營現金流按年增長率 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1.0,10.0] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）

    // 現金流屬性
    FinancialField_OperatingRevenueCashCover = 42; // 經營現金收入比 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [10,100] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    FinancialField_OperatingProfitToTotalProfit = 43; // 營業利潤佔比 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [10,100] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）     

    // 市場表現屬性
    FinancialField_BasicEPS = 44; // 基本每股收益 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [0.1,10] 值區間（單位：元）
    FinancialField_DilutedEPS = 45; // 稀釋每股收益 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [0.1,10] 值區間（單位：元）
    FinancialField_NOCFPerShare = 46; // 每股經營現金淨流量 （精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [0.1,10] 值區間（單位：元）    
}
```

## 財務過濾屬性週期

**FinancialQuarter**

```protobuf
enum FinancialQuarter
{
    FinancialQuarter_Unknown = 0; // 未知
    FinancialQuarter_Annual = 1; // 年報
    FinancialQuarter_FirstQuarter = 2; // 一季報
    FinancialQuarter_Interim = 3; // 中報
    FinancialQuarter_ThirdQuarter = 4; // 三季報
    FinancialQuarter_MostRecentQuarter = 5; // 最近季報
}
```

## 自定義技術指標屬性

**CustomIndicatorField**

```protobuf
enum CustomIndicatorField
{
    CustomIndicatorField_Unknown = 0; // 未知
    CustomIndicatorField_Price = 1; // 最新價格 
    CustomIndicatorField_MA5 = 2; // 5日簡單均線（不建議使用）
    CustomIndicatorField_MA10 = 3; // 10日簡單均線（不建議使用）
    CustomIndicatorField_MA20 = 4; // 20日簡單均線（不建議使用） 
    CustomIndicatorField_MA30 = 5; // 30日簡單均線（不建議使用） 
    CustomIndicatorField_MA60 = 6; // 60日簡單均線（不建議使用） 
    CustomIndicatorField_MA120 = 7; // 120日簡單均線（不建議使用）
    CustomIndicatorField_MA250 = 8; // 250日簡單均線（不建議使用）
    CustomIndicatorField_RSI = 9; // RSI 指標參數的預設值為[12]
    CustomIndicatorField_EMA5 = 10; // 5日指數移動均線（不建議使用） 
    CustomIndicatorField_EMA10 = 11; // 10日指數移動均線（不建議使用） 
    CustomIndicatorField_EMA20 = 12; // 20日指數移動均線（不建議使用） 
    CustomIndicatorField_EMA30 = 13; // 30日指數移動均線（不建議使用） 
    CustomIndicatorField_EMA60 = 14; // 60日指數移動均線（不建議使用） 
    CustomIndicatorField_EMA120 = 15; // 120日指數移動均線（不建議使用）
    CustomIndicatorField_EMA250 = 16; // 250日指數移動均線（不建議使用）
    CustomIndicatorField_Value = 17; // 自定義數值（stock_field1不支援此欄位）
    CustomIndicatorField_MA = 30; // 簡單均線
	CustomIndicatorField_EMA = 40; // 指數移動均線
	CustomIndicatorField_KDJ_K = 50; // KDJ 指標的 K 值。指標參數需要根據 KDJ 進行傳參。不傳則預設為 [9,3,3]
	CustomIndicatorField_KDJ_D = 51; // KDJ 指標的 D 值。指標參數需要根據 KDJ 進行傳參。不傳則預設為 [9,3,3]
	CustomIndicatorField_KDJ_J = 52; // KDJ 指標的 J 值。指標參數需要根據 KDJ 進行傳參。不傳則預設為 [9,3,3]	
	CustomIndicatorField_MACD_DIFF = 60; // MACD 指標的 DIFF 值。指標參數需要根據 MACD 進行傳參。不傳則預設為 [12,26,9]
	CustomIndicatorField_MACD_DEA = 61; // MACD 指標的 DEA 值。指標參數需要根據 MACD 進行傳參。不傳則預設為 [12,26,9]
	CustomIndicatorField_MACD = 62; // MACD 指標的 MACD 值。指標參數需要根據 MACD 進行傳參。不傳則預設為 [12,26,9]
	CustomIndicatorField_BOLL_UPPER = 70; // BOLL 指標的 UPPER 值。指標參數需要根據 BOLL 進行傳參。不傳則預設為 [20,2]
	CustomIndicatorField_BOLL_MIDDLER = 71; // BOLL 指標的 MIDDLER 值。指標參數需要根據 BOLL 進行傳參。不傳則預設為 [20,2]
	CustomIndicatorField_BOLL_LOWER = 72; // BOLL 指標的 LOWER 值。指標參數需要根據 BOLL 進行傳參。不傳則預設為 [20,2]
}
```

## 相對位置

**RelativePosition**

```protobuf
enum RelativePosition
{
    RelativePosition_Unknown = 0; // 未知
    RelativePosition_More = 1; // 大於，first位於second的上方
    RelativePosition_Less = 2; // 小於，first位於second的下方
    RelativePosition_CrossUp = 3; // 升穿，first從下往上穿second
    RelativePosition_CrossDown = 4; // 跌穿，first從上往下穿second
}
```

## 形態技術指標屬性

**PatternField**

```protobuf
enum PatternField
{
    PatternField_Unknown = 0 ; // 未知
    PatternField_MAAlignmentLong = 1 ; // MA多頭排列（連續兩天MA5>MA10>MA20>MA30>MA60，且當日收盤價大於前一天收盤價）
    PatternField_MAAlignmentShort = 2 ; // MA空頭排列（連續兩天MA5 <MA10 <MA20 <MA30 <MA60，且當日收盤價小於前一天收盤價）
    PatternField_EMAAlignmentLong = 3 ; // EMA多頭排列（連續兩天EMA5>EMA10>EMA20>EMA30>EMA60，且當日收盤價大於前一天收盤價）
    PatternField_EMAAlignmentShort = 4 ; // EMA空頭排列（連續兩天EMA5 <EMA10 <EMA20 <EMA30 <EMA60，且當日收盤價小於前一天收盤價）
    PatternField_RSIGoldCrossLow = 5 ; // RSI低位金叉（50以下，短線RSI上穿長線RSI（前一日短線RSI小於長線RSI，當日短線RSI大於長線RSI）） 
    PatternField_RSIDeathCrossHigh = 6 ; // RSI高位死叉（50以上，短線RSI下穿長線RSI（前一日短線RSI大於長線RSI，當日短線RSI小於長線RSI）） 
    PatternField_RSITopDivergence = 7 ; // RSI頂背離（相鄰的兩個K線波峯，後面的波峯對應的CLOSE>前面的波峯對應的CLOSE，後面波峯的RSI12值 <前面波峯的RSI12值）
    PatternField_RSIBottomDivergence = 8 ; // RSI底背離（相鄰的兩個K線波谷，後面的波谷對應的CLOSE <前面的波谷對應的CLOSE，後面波谷的RSI12值>前面波谷的RSI12值） 
    PatternField_KDJGoldCrossLow = 9 ; // KDJ低位金叉（KDJ的值都小於或等於30，且前一日K,J值分別小於D值，當日K,J值分別大於D值）
    PatternField_KDJDeathCrossHigh = 10 ; // KDJ高位死叉（KDJ的值都大於或等於70，且前一日K,J值分別大於D值，當日K,J值分別小於D值）
    PatternField_KDJTopDivergence = 11 ; // KDJ頂背離（相鄰的兩個K線波峯，後面的波峯對應的CLOSE>前面的波峯對應的CLOSE，後面波峯的J值 <前面波峯的J值）
    PatternField_KDJBottomDivergence = 12 ; // KDJ底背離（相鄰的兩個K線波谷，後面的波谷對應的CLOSE <前面的波谷對應的CLOSE，後面波谷的J值>前面波谷的J值）
    PatternField_MACDGoldCrossLow = 13 ; // MACD低位金叉（DIFF上穿DEA（前一日DIFF小於DEA，當日DIFF大於DEA））
    PatternField_MACDDeathCrossHigh = 14 ; // MACD高位死叉（DIFF下穿DEA（前一日DIFF大於DEA，當日DIFF小於DEA））
    PatternField_MACDTopDivergence = 15 ; // MACD頂背離（相鄰的兩個K線波峯，後面的波峯對應的CLOSE>前面的波峯對應的CLOSE，後面波峯的macd值 <前面波峯的macd值）
    PatternField_MACDBottomDivergence = 16 ; // MACD底背離（相鄰的兩個K線波谷，後面的波谷對應的CLOSE <前面的波谷對應的CLOSE，後面波谷的macd值>前面波谷的macd值）
    PatternField_BOLLBreakUpper = 17 ; // BOLL突破上軌（前一日股價低於上軌值，當日股價大於上軌值） 
    PatternField_BOLLLower = 18 ; // BOLL突破下軌（前一日股價高於下軌值，當日股價小於下軌值）
    PatternField_BOLLCrossMiddleUp = 19 ; // BOLL向上破中軌（前一日股價低於中軌值，當日股價大於中軌值）
    PatternField_BOLLCrossMiddleDown = 20 ; // BOLL向下破中軌（前一日股價大於中軌值，當日股價小於中軌值）
}
```

## 自選股分組類型

**GroupType**

```protobuf
enum GroupType
{
    GroupType_Unknown = 0; // 未知
    GroupType_Custom = 1; // 自定義分組
    GroupType_System = 2; // 系統分組
    GroupType_All = 3; // 全部分組
}
```

## 指數期權類別

**IndexOptionType**

```protobuf
enum IndexOptionType
{
    IndexOptionType_Unknown = 0; //未知
    IndexOptionType_Normal = 1; //普通的指數期權
    IndexOptionType_Small = 2; //小型指數期權
}
```

## 上市時段

**IpoPeriod**

```protobuf
enum IpoPeriod
{
    IpoPeriod_Unknow = 0; //未知
    IpoPeriod_Today = 1; //今日上市
    IpoPeriod_Tomorrow = 2; //明日上市
    IpoPeriod_Nextweek = 3; //未來一週上市
    IpoPeriod_Lastweek = 4; //過去一週上市
    IpoPeriod_Lastmonth = 5; //過去一月上市
}
```

## 窩輪發行商

**Issuer**

```protobuf
enum Issuer
{
    Issuer_Unknow = 0; //未知
    Issuer_SG = 1; //法興
    Issuer_BP = 2; //法巴
    Issuer_CS = 3; //瑞信
    Issuer_CT = 4; //花旗    
    Issuer_EA = 5; //東亞
    Issuer_GS = 6; //高盛
    Issuer_HS = 7; //滙豐
    Issuer_JP = 8; //摩通    
    Issuer_MB = 9; //麥銀    
    Issuer_SC = 10; //渣打
    Issuer_UB = 11; //瑞銀
    Issuer_BI = 12; //中銀
    Issuer_DB = 13; //德銀
    Issuer_DC = 14; //大和
    Issuer_ML = 15; //美林
    Issuer_NM = 16; //野村
    Issuer_RB = 17; //荷合
    Issuer_RS = 18; //蘇皇    
    Issuer_BC = 19; //巴克萊
    Issuer_HT = 20; //海通
    Issuer_VT = 21; //瑞通
    Issuer_KC = 22; //比聯
    Issuer_MS = 23; //摩利
    Issuer_GJ = 24; //國君
    Issuer_XZ = 25; //星展
    Issuer_HU = 26; //華泰
    Issuer_KS = 27; //韓投
    Issuer_CI = 28; //信證
}
```

## K 線欄位

**KLFields**

```protobuf
enum KLFields
{
    KLFields_None = 0; //
    KLFields_High = 1; //最高價
    KLFields_Open = 2; //開盤價
    KLFields_Low = 4; //最低價
    KLFields_Close = 8; //收盤價
    KLFields_LastClose = 16; //昨收價
    KLFields_Volume = 32; //成交量
    KLFields_Turnover = 64; //成交額
    KLFields_TurnoverRate = 128; //換手率
    KLFields_PE = 256; //市盈率
    KLFields_ChangeRate = 512; //漲跌幅
}
```

## K 線類型

**KLType**

```protobuf
enum KLType
{
    KLType_Unknown = 0; //未知
    KLType_1Min = 1; //1分 K
    KLType_Day = 2; //日 K
    KLType_Week = 3; //周 K (期權暫不支援)
    KLType_Month = 4; //月 K (期權暫不支援)    
    KLType_Year = 5; //年 K (期權暫不支援)
    KLType_5Min = 6; //5分 K
    KLType_15Min = 7; //15分 K
    KLType_30Min = 8; //30分 K (期權暫不支援)
    KLType_60Min = 9; //60分 K        
    KLType_3Min = 10; //3分 K (期權暫不支援)
    KLType_Quarter = 11; //季 K (期權暫不支援)
}
```

## 週期類型

**PeriodType**

```protobuf
enum PeriodType
{
    PeriodType_INTRADAY = 0; //實時
    PeriodType_DAY = 1; //日
    PeriodType_WEEK = 2; //周
    PeriodType_MONTH = 3; //月
}
```


## 到價提醒市場狀態

**MarketStatus**

```protobuf
enum MarketStatus
{
    MarketStatus_Unknow = 0;
    MarketStatus_Open = 1; // 盤中
    MarketStatus_USPre = 2;  // 美股盤前
    MarketStatus_USAfter = 3; // 美股盤後
}
```

## 自選股操作

**ModifyUserSecurityOp**

```protobuf
enum ModifyUserSecurityOp
{
    ModifyUserSecurityOp_Unknown = 0;
    ModifyUserSecurityOp_Add = 1; //新增
    ModifyUserSecurityOp_Del = 2; //刪除自選
    ModifyUserSecurityOp_MoveOut = 3; //移出分組
}
```

## 期權類型（按行權時間）

**OptionAreaType**

```protobuf
enum OptionAreaType
{
    OptionAreaType_Unknown = 0; //未知
    OptionAreaType_American = 1; //美式
    OptionAreaType_European = 2; //歐式
    OptionAreaType_Bermuda = 3; //百慕大
}
```

## 期權價內/外

**OptionCondType**

```protobuf
enum OptionCondType
{
    OptionCondType_Unknow = 0;
    OptionCondType_WithIn = 1; //價內
    OptionCondType_Outside = 2; //價外
}
```

## 期權類型（按方向）

**OptionType**

```protobuf
enum OptionType
{
    OptionType_Unknown = 0; //未知
    OptionType_Call = 1; //認購期權
    OptionType_Put = 2; //認沽期權
};
```

## 板塊集合類型

**PlateSetType**

```protobuf
enum PlateSetType
{
    PlateSetType_All = 0; //所有板塊
    PlateSetType_Industry = 1; //行業板塊
    PlateSetType_Region = 2; //地域板塊,港美股市場的地域分類資料暫為空
    PlateSetType_Concept = 3; //概念板塊
    PlateSetType_Other = 4; //其他板塊, 僅用於3207（獲取股票所屬板塊）協議返回,不可作為其他協議的請求參數
}
```

## 到價提醒頻率

**PriceReminderFreq**

```protobuf
enum PriceReminderFreq
{
    PriceReminderFreq_Unknown = 0; // 未知
    PriceReminderFreq_Always = 1; // 持續提醒
    PriceReminderFreq_OnceADay = 2; // 每日一次
    PriceReminderFreq_OnlyOnce = 3; // 僅提醒一次
}
```

## 到價提醒類型

**PriceReminderType**

```protobuf
enum PriceReminderType
{
    PriceReminderType_Unknown = 0; // 未知
    PriceReminderType_PriceUp = 1; // 價格漲到
    PriceReminderType_PriceDown = 2; // 價格跌到
    PriceReminderType_ChangeRateUp = 3; // 日漲幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_ChangeRateDown = 4; // 日跌幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_5MinChangeRateUp = 5; // 5 分鐘漲幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_5MinChangeRateDown = 6; // 5 分鐘跌幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_VolumeUp = 7; // 成交量超過
    PriceReminderType_TurnoverUp = 8; // 成交額超過
    PriceReminderType_TurnoverRateUp = 9; // 換手率超過（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_BidPriceUp = 10; // 買一價高於
    PriceReminderType_AskPriceDown = 11; // 賣一價低於
    PriceReminderType_BidVolUp = 12; // 買一量高於    
    PriceReminderType_AskVolUp = 13; // 賣一量高於
    PriceReminderType_3MinChangeRateUp = 14; // 3 分鐘漲幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
    PriceReminderType_3MinChangeRateDown = 15; // 3 分鐘跌幅超（該欄位為百分比欄位，設置時填 20 表示 20%）
}
```

## 窩輪價內/外

**PriceType**

```protobuf
enum PriceType
{
    PriceType_Unknow = 0;
    PriceType_Outside = 1; //價外，界內證表示界外
    PriceType_WithIn = 2; //價內，界內證表示界內
}
```

## 逐筆推送類型

**PushDataType**

```protobuf
enum PushDataType
{
    PushDataType_Unknow = 0;
    PushDataType_Realtime = 1; //實時推送的資料
    PushDataType_ByDisConn = 2; //對後台行情連線斷開期間拉取補充的資料（最多50個）
    PushDataType_Cache = 3; //非實時非連線斷開補充資料
}
```

## 行情市場

**QotMarket**

```protobuf
enum QotMarket
{
    QotMarket_Unknown = 0; //未知市場
    QotMarket_HK_Security = 1; //香港市場
    QotMarket_HK_Future = 2; //港期貨（已廢棄，使用 QotMarket_HK_Security 即可）
    QotMarket_US_Security = 11; //美國市場
    QotMarket_CNSH_Security = 21; //滬股市場
    QotMarket_CNSZ_Security = 22; //深股市場
    QotMarket_SG_Security = 31; //新加坡市場
    QotMarket_JP_Security = 41; //日本市場
    QotMarket_AU_Security = 51; //澳大利亞市場
    QotMarket_MY_Security = 61; //馬來西亞市場
    QotMarket_CA_Security = 71; //加拿大市場
    QotMarket_FX_Security = 81; //外匯市場
}
```

## 市場狀態

**QotMarketState**

各市場狀態的對應時段：[點擊這裡](../qa/quote.md#3454)瞭解更多

```protobuf
enum QotMarketState
{
    QotMarketState_None = 0; // 無交易
    QotMarketState_Auction = 1; // 盤前競價 
    QotMarketState_WaitingOpen = 2; // 等待開盤
    QotMarketState_Morning = 3; // 早盤 
    QotMarketState_Rest = 4; // 午間休市 
    QotMarketState_Afternoon = 5; // 午盤 / 美股持續交易時段
    QotMarketState_Closed = 6; // 收盤
    QotMarketState_PreMarketBegin = 8; // 美股盤前交易時段
    QotMarketState_PreMarketEnd = 9; // 美股盤前交易結束 
    QotMarketState_AfterHoursBegin = 10; // 美股盤後交易時段
    QotMarketState_AfterHoursEnd = 11; // 美股收盤 
    QotMarketState_NightOpen = 13; // 夜市交易時段
    QotMarketState_NightEnd = 14; // 夜市收盤 
    QotMarketState_FutureDayOpen = 15; // 日市交易時段
    QotMarketState_FutureDayBreak = 16; // 日市休市 
    QotMarketState_FutureDayClose = 17; // 日市收盤 
    QotMarketState_FutureDayWaitForOpen = 18; // 期貨待開盤 
    QotMarketState_HkCas = 19; // 盤後競價,港股市場增加 CAS 機制對應的市場狀態    
    QotMarketState_FutureNightWait = 20; // 夜市等待開盤（已廢棄）
    QotMarketState_FutureAfternoon = 21; // 期貨下午開盤（已廢棄）
    //美國期貨新增加狀態
    QotMarketState_FutureSwitchDate = 22; // 美期待開盤
    QotMarketState_FutureOpen = 23; // 美期交易時段
    QotMarketState_FutureBreak = 24; // 美期中盤休息
    QotMarketState_FutureBreakOver = 25; // 美期休息後交易時段
    QotMarketState_FutureClose = 26; // 美期收盤
    //科創板新增狀態
    QotMarketState_StibAfterHoursWait = 27; // 科創板的盤後撮合時段（已廢棄）
    QotMarketState_StibAfterHoursBegin = 28; // 科創板的盤後交易開始（已廢棄）
    QotMarketState_StibAfterHoursEnd = 29; // 科創板的盤後交易結束（已廢棄）
    //美指期權新增加狀態
    QotMarketState_NIGHT = 32; // 美指期權夜市交易時段
    QotMarketState_TRADE_AT_LAST = 35; // 美指期權盤尾交易時段
    QotMarketState_OVERNIGHT = 37;  // 美股夜盤交易時段
}
```

## 美股時段

> **Session**

```protobuf
enum Session
{
	Session_NONE = 0; // 未知
	Session_RTH = 1; // 盤中
	Session_ETH = 2; // 盤中+盤前盤後
	Session_ALL = 3; // 全時段
	Session_OVERNIGHT = 4; // 夜盤
}
```

## 行情權限

**QotRight**

```protobuf
enum QotRight
{
    QotRight_Unknow = 0; //未知
    QotRight_Bmp = 1; //BMP（此權限不支援訂閲）
    QotRight_Level1 = 2; //Level1
    QotRight_Level2 = 3; //Level2
    QotRight_SF = 4; //SF 高級行情
    QotRight_No = 5; //無權限
}
```

## 關聯資料類型

**ReferenceType**

```protobuf
enum ReferenceType
{
    ReferenceType_Unknow = 0; 
    ReferenceType_Warrant = 1; //正股相關的窩輪
    ReferenceType_Future = 2; //期貨主連的相關合約
}
```

## K 線復權類型

**RehabType**

```protobuf
enum RehabType
{
    RehabType_None = 0; //不復權
    RehabType_Forward = 1; //前復權
    RehabType_Backward = 2; //後復權
}
```

## 股票狀態

**SecurityStatus**

```protobuf
enum SecurityStatus
{
    SecurityStatus_Unknown = 0; //未知
    SecurityStatus_Normal = 1; //正常狀態
    SecurityStatus_Listing = 2; //待上市
    SecurityStatus_Purchasing = 3; //申購中
    SecurityStatus_Subscribing = 4; //認購中
    SecurityStatus_BeforeDrakTradeOpening = 5; //暗盤開盤前
    SecurityStatus_DrakTrading = 6; //暗盤交易中
    SecurityStatus_DrakTradeEnd = 7; //暗盤已收盤
    SecurityStatus_ToBeOpen = 8; //待開盤
    SecurityStatus_Suspended = 9; //停牌
    SecurityStatus_Called = 10; //已收回
    SecurityStatus_ExpiredLastTradingDate = 11; //已過最後交易日
    SecurityStatus_Expired = 12; //已過期
    SecurityStatus_Delisted = 13; //已退市
    SecurityStatus_ChangeToTemporaryCode = 14; //公司行動中，交易關閉，轉至臨時代碼交易
    SecurityStatus_TemporaryCodeTradeEnd = 15; //臨時買賣結束，交易關閉
    SecurityStatus_ChangedPlateTradeEnd = 16; //已轉板，舊代碼交易關閉
    SecurityStatus_ChangedCodeTradeEnd = 17; //已換代碼，舊代碼交易關閉
    SecurityStatus_RecoverableCircuitBreaker = 18; //可恢復性熔斷
    SecurityStatus_UnRecoverableCircuitBreaker = 19; //不可恢復性熔斷
    SecurityStatus_AfterCombination = 20; //盤後撮合
    SecurityStatus_AfterTransation = 21; //盤後交易
}
```

## 股票類型

**SecurityType**

```protobuf
enum SecurityType
{
    SecurityType_Unknown = 0; //未知
    SecurityType_Bond = 1; //債券
    SecurityType_Bwrt = 2; //一攬子權證
    SecurityType_Eqty = 3; //正股
    SecurityType_Trust = 4; //信託,基金
    SecurityType_Warrant = 5; //窩輪
    SecurityType_Index = 6; //指數
    SecurityType_Plate = 7; //板塊
    SecurityType_Drvt = 8; //期權
    SecurityType_PlateSet = 9; //板塊集
    SecurityType_Future = 10; //期貨
}
```

## 設置到價提醒操作類型

**SetPriceReminderOp**

```protobuf
enum SetPriceReminderOp
{
    SetPriceReminderOp_Unknown = 0;
    SetPriceReminderOp_Add = 1; //新增
    SetPriceReminderOp_Del = 2; //刪除
    SetPriceReminderOp_Enable = 3; //啟用
    SetPriceReminderOp_Disable = 4; //禁用
    SetPriceReminderOp_Modify = 5; //修改
    SetPriceReminderOp_DelAll = 6; //刪除全部（刪除指定股票下的所有到價提醒）
}
```

## 排序方向

**SortDir**

```protobuf
enum SortDir
{
    SortDir_No = 0; // 不排序
    SortDir_Ascend = 1; // 升序
    SortDir_Descend = 2; // 降序
}
```

## 排序欄位

**SortField**

```protobuf
enum SortField
{
    SortField_Unknow = 0;
    SortField_Code = 1; //代碼
    SortField_CurPrice = 2; //最新價
    SortField_PriceChangeVal = 3; //漲跌額
    SortField_ChangeRate = 4; //漲跌幅%
    SortField_Status = 5; //狀態
    SortField_BidPrice = 6; //買入價
    SortField_AskPrice = 7; //賣出價
    SortField_BidVol = 8; //買量
    SortField_AskVol = 9; //賣量
    SortField_Volume = 10; //成交量
    SortField_Turnover = 11; //成交額
    SortField_Amplitude = 30; //振幅%

    //以下排序欄位只支援用於 Qot_GetWarrant 協議
    SortField_Score = 12; //綜合評分
    SortField_Premium = 13; //溢價%
    SortField_EffectiveLeverage = 14; //有效槓桿
    SortField_Delta = 15; //對沖值,僅認購認沽支援該欄位
    SortField_ImpliedVolatility = 16; //引伸波幅,僅認購認沽支援該欄位
    SortField_Type = 17; //類型
    SortField_StrikePrice = 18; //行權價
    SortField_BreakEvenPoint = 19; //打和點
    SortField_MaturityTime = 20; //到期日
    SortField_ListTime = 21; //上市日期
    SortField_LastTradeTime = 22; //最後交易日
    SortField_Leverage = 23; //槓桿比率
    SortField_InOutMoney = 24; //價內/價外%
    SortField_RecoveryPrice = 25; //收回價,僅牛熊證支援該欄位
    SortField_ChangePrice = 26; // 換股價
    SortField_Change = 27; //換股比率
    SortField_StreetRate = 28; //街貨比%
    SortField_StreetVol = 29; //街貨量
    SortField_WarrantName = 31; // 窩輪名稱
    SortField_Issuer = 32; //發行人
    SortField_LotSize = 33; // 每手
    SortField_IssueSize = 34; //發行量
    SortField_UpperStrikePrice = 45; //上限價，僅用於界內證
    SortField_LowerStrikePrice = 46; //下限價，僅用於界內證
    SortField_InLinePriceStatus = 47; //界內界外，僅用於界內證

    //以下排序欄位只支援用於 Qot_GetPlateSecurity 協議，並僅支援美股
    SortField_PreCurPrice = 35; //盤前最新價
    SortField_AfterCurPrice = 36; //盤後最新價
    SortField_PrePriceChangeVal = 37; //盤前漲跌額
    SortField_AfterPriceChangeVal = 38; //盤後漲跌額
    SortField_PreChangeRate = 39; //盤前漲跌幅%
    SortField_AfterChangeRate = 40; //盤後漲跌幅%
    SortField_PreAmplitude = 41; //盤前振幅%
    SortField_AfterAmplitude = 42; //盤後振幅%
    SortField_PreTurnover = 43; //盤前成交額
    SortField_AfterTurnover = 44; //盤後成交額

    //以下排序欄位只支援用於 Qot_GetPlateSecurity 協議，並僅支援期貨
    SortField_LastSettlePrice = 48; //昨結
    SortField_Position = 49; //持倉量
    SortField_PositionChange = 50; //日增倉
}
```

## 簡單過濾屬性

**StockField**

```protobuf
enum StockField  
{
    StockField_Unknown = 0; // 未知
    StockField_StockCode = 1; // 股票代碼，不能填區間上下限值。
    StockField_StockName = 2; // 股票名稱，不能填區間上下限值。
    StockField_CurPrice = 3; // 最新價（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[10,20]值區間
    StockField_CurPriceToHighest52WeeksRatio = 4; // (現價 - 52周最高)/52周最高，對應 PC 端離52周高點百分比（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-30,-10]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_CurPriceToLowest52WeeksRatio = 5; // (現價 - 52周最低)/52周最低，對應 PC 端離52周低點百分比（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[20,40]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_HighPriceToHighest52WeeksRatio = 6; // (今日最高 - 52周最高)/52周最高（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-3,-1]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_LowPriceToLowest52WeeksRatio = 7; // (今日最低 - 52周最低)/52周最低（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[10,70]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_VolumeRatio = 8; // 量比（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[0.5,30]值區間
    StockField_BidAskRatio = 9; // 委比（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-20,80.5]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_LotPrice = 10; // 每手價格（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[40,100]值區間
    StockField_MarketVal = 11; // 市值（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[50000000,3000000000]值區間
    StockField_PeAnnual = 12; // 市盈率(靜態)（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-8,65.3]值區間
    StockField_PeTTM = 13; // 市盈率 TTM （精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-10,20.5]值區間
    StockField_PbRate = 14; // 市淨率（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[0.5,20]值區間
    StockField_ChangeRate5min = 15; // 五分鐘價格漲跌幅（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-5,6.3]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_ChangeRateBeginYear = 16; // 年初至今價格漲跌幅（精確到小數點後 3 位，超出部分會被捨棄）例如填寫[-50.1,400.7]值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）    
        
    // 基礎量價屬性
    StockField_PSTTM = 17; // 市銷率(TTM)（精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [100, 500] 值區間（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%） 
    StockField_PCFTTM = 18; // 市現率(TTM)（精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [100, 1000] 值區間 （該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    StockField_TotalShare = 19; // 總股數（精確到小數點後 0 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間 (單位：股)
    StockField_FloatShare = 20; // 流通股數（精確到小數點後 0 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間 (單位：股)
    StockField_FloatMarketVal = 21; // 流通市值（精確到小數點後 3 位，超出部分會被捨棄）例如填寫 [1000000000,1000000000] 值區間（單位：元）
}
```

## 訂閲類型

**SubType**

```protobuf
enum SubType
{
    SubType_None = 0;
    SubType_Basic = 1; //基礎報價
    SubType_OrderBook = 2; //擺盤
    SubType_Ticker = 4; //逐筆
    SubType_RT = 5; //分時
    SubType_KL_Day = 6; //日 K
    SubType_KL_5Min = 7; //5分 K
    SubType_KL_15Min = 8; //15分 K
    SubType_KL_30Min = 9; //30分 K
    SubType_KL_60Min = 10; //60分 K
    SubType_KL_1Min = 11; //1分 K
    SubType_KL_Week = 12; //周 K
    SubType_KL_Month = 13; //月 K
    SubType_Broker = 14; //經紀隊列
    SubType_KL_Qurater = 15; //季 K
    SubType_KL_Year = 16; //年 K
    SubType_KL_3Min = 17; //3分 K
}
```

## 逐筆成交方向

**TickerDirection**

```protobuf
enum TickerDirection
{
    TickerDirection_Unknown = 0; //未知
    TickerDirection_Bid = 1; //外盤（主動買入），即以賣一價或更高的價格成交股票
    TickerDirection_Ask = 2; //內盤（主動賣出），即以買一價或更低的價格成交股票
    TickerDirection_Neutral = 3; //中性盤，即以買一價與賣一價之間的價格撮合成交
}
```

## 逐筆成交類型

**TickerType**

```protobuf
enum TickerType
{
    TickerType_Unknown = 0; //未知
    TickerType_Automatch = 1; //自動對盤
    TickerType_Late = 2; //開市前成交盤
    TickerType_NoneAutomatch = 3; //非自動對盤
    TickerType_InterAutomatch = 4; //同一證券商自動對盤
    TickerType_InterNoneAutomatch = 5; //同一證券商非自動對盤
    TickerType_OddLot = 6; //碎股交易
    TickerType_Auction = 7; //競價交易    
    TickerType_Bulk = 8; //批量交易
    TickerType_Crash = 9; //現金交易
    TickerType_CrossMarket = 10; //跨市場交易
    TickerType_BulkSold = 11; //批量賣出
    TickerType_FreeOnBoard = 12; //離價交易
    TickerType_Rule127Or155 = 13; //第127條交易（紐交所規則）或第155條交易
    TickerType_Delay = 14; //延遲交易
    TickerType_MarketCenterClosePrice = 15; //中央收市價
    TickerType_NextDay = 16; //隔日交易
    TickerType_MarketCenterOpening = 17; //中央開盤價交易
    TickerType_PriorReferencePrice = 18; //前參考價
    TickerType_MarketCenterOpenPrice = 19; //中央開盤價
    TickerType_Seller = 20; //賣方
    TickerType_T = 21; //T 類交易(盤前和盤後交易)
    TickerType_ExtendedTradingHours = 22; //延長交易時段
    TickerType_Contingent = 23; //合單交易
    TickerType_AvgPrice = 24; //平均價成交
    TickerType_OTCSold = 25; //場外售出
    TickerType_OddLotCrossMarket = 26; //碎股跨市場交易
    TickerType_DerivativelyPriced = 27; //衍生工具定價
    TickerType_ReOpeningPriced = 28; //再開盤定價
    TickerType_ClosingPriced = 29; //收盤定價
    TickerType_ComprehensiveDelayPrice = 30; //綜合延遲價格
    TickerType_Overseas = 31; //交易的一方不是香港交易所的成員，屬於場外交易
}
```

## 交易日查詢市場

**TradeDateMarket**

```protobuf
enum TradeDateMarket
{
    TradeDateMarket_Unknown = 0; //未知
    TradeDateMarket_HK = 1; //香港市場（含股票、ETFs、窩輪、牛熊、期權、非假期交易期貨；不含假期交易期貨）
    TradeDateMarket_US = 2; //美國市場（含股票、ETFs、期權；不含期貨）
    TradeDateMarket_CN = 3; //A 股市場
    TradeDateMarket_NT = 4; //深（滬）股通
    TradeDateMarket_ST = 5; //港股通（深、滬）
    TradeDateMarket_JP_Future = 6; //日本期貨
    TradeDateMarket_SG_Future = 7; //新加坡期貨
}
```

## 交易日類型

**TradeDateType**

```protobuf
enum TradeDateType
{
    TradeDateType_Whole = 0; //全天交易
    TradeDateType_Morning = 1; //上午交易，下午休市
    TradeDateType_Afternoon = 2; //下午交易，上午休市
}            
```

## 窩輪狀態

**WarrantStatus**

```protobuf
enum WarrantStatus
{
    WarrantStatus_Unknow = 0; //未知
    WarrantStatus_Normal = 1; //正常狀態
    WarrantStatus_Suspend = 2; //停牌
    WarrantStatus_StopTrade = 3; //終止交易
    WarrantStatus_PendingListing = 4; //等待上市
}
```

## 窩輪類型

**WarrantType**

```protobuf
enum WarrantType
{
    WarrantType_Unknown = 0; //未知
    WarrantType_Buy = 1; //認購窩輪
    WarrantType_Sell = 2; //認沽窩輪
    WarrantType_Bull = 3; //牛證
    WarrantType_Bear = 4; //熊證
    WarrantType_InLine = 5; //界內證
};
```

## 所屬交易所

**ExchType**

```protobuf
enum ExchType
{
    ExchType_Unknown = 0; //未知
    ExchType_HK_MainBoard = 1; // 港交所·主板
    ExchType_HK_GEMBoard = 2; //港交所·創業板
    ExchType_HK_HKEX = 3; //港交所
    ExchType_US_NYSE = 4; //紐交所
    ExchType_US_Nasdaq = 5; //納斯達克
    ExchType_US_Pink = 6; //OTC 市場
    ExchType_US_AMEX = 7; //美交所 
    ExchType_US_Option = 8; //美國（僅美股期權適用）
    ExchType_US_NYMEX = 9; //NYMEX 
    ExchType_US_COMEX = 10; //COMEX
    ExchType_US_CBOT = 11; //CBOT
    ExchType_US_CME = 12; //CME
    ExchType_US_CBOE = 13; //CBOE
    ExchType_CN_SH = 14; //上交所  
    ExchType_CN_SZ = 15; //深交所
    ExchType_CN_STIB = 16; //科創板
    ExchType_SG_SGX = 17; //新交所
    ExchType_JP_OSE = 18; //大阪交易所 
};
```

## 證券標識

**Security**

```protobuf
message Security
{
    required int32 market = 1; //QotMarket，行情市場
    required string code = 2; //代碼
}
```

## K 線資料

**KLine**

```protobuf
message KLine
{
    required string time = 1; //時間戳字符串（格式：yyyy-MM-dd HH:mm:ss）
    required bool isBlank = 2; //是否是空內容的點,若為 true 則只有時間資訊
    optional double highPrice = 3; //最高價
    optional double openPrice = 4; //開盤價
    optional double lowPrice = 5; //最低價
    optional double closePrice = 6; //收盤價
    optional double lastClosePrice = 7; //昨收價
    optional int64 volume = 8; //成交量
    optional double turnover = 9; //成交額
    optional double turnoverRate = 10; //換手率（該欄位為百分比欄位，展示為小數表示）
    optional double pe = 11; //市盈率
    optional double changeRate = 12; //漲跌幅（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    optional double timestamp = 13; //時間戳
}
```

## 基礎報價的期權特有欄位

**OptionBasicQotExData**

```protobuf
message OptionBasicQotExData
{
    required double strikePrice = 1; //行權價
    required int32 contractSize = 2; //每份合約數(整型資料)
    optional double contractSizeFloat = 17; //每份合約數（浮點型資料）
    required int32 openInterest = 3; //未平倉合約數
    required double impliedVolatility = 4; //隱含波動率（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    required double premium = 5; //溢價（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    required double delta = 6; //希臘值 Delta
    required double gamma = 7; //希臘值 Gamma
    required double vega = 8; //希臘值 Vega
    required double theta = 9; //希臘值 Theta
    required double rho = 10; //希臘值 Rho
    optional int32 netOpenInterest = 11; //淨未平倉合約數，僅港股期權適用
    optional int32 expiryDateDistance = 12; //距離到期日天數，負數表示已過期
    optional double contractNominalValue = 13; //合約名義金額，僅港股期權適用
    optional double ownerLotMultiplier = 14; //相等正股手數，指數期權無該欄位，僅港股期權適用
    optional int32 optionAreaType = 15; //OptionAreaType，期權類型（按行權時間）
    optional double contractMultiplier = 16; //合約乘數
    optional int32 indexOptionType = 18; //IndexOptionType，指數期權類型
}    
```

## 基礎報價的期貨特有欄位

**FutureBasicQotExData**

```protobuf
message FutureBasicQotExData
{
    required double lastSettlePrice = 1; //昨結
    required int32 position = 2; //持倉量
    required int32 positionChange = 3; //日增倉
    optional int32 expiryDateDistance = 4; //距離到期日天數
}    
```

## 基礎報價

**BasicQot**

```protobuf
message BasicQot
{
    required Security security = 1; //股票
    optional string name = 24; // 股票名稱
    required bool isSuspended = 2; //是否停牌
    required string listTime = 3; //上市日期字符串（此欄位停止維護，不建議使用，格式：yyyy-MM-dd）
    required double priceSpread = 4; //價差
    required string updateTime = 5; //最新價的更新時間字符串（格式：yyyy-MM-dd HH:mm:ss），對其他欄位不適用
    required double highPrice = 6; //最高價
    required double openPrice = 7; //開盤價
    required double lowPrice = 8; //最低價
    required double curPrice = 9; //最新價
    required double lastClosePrice = 10; //昨收價
    required int64 volume = 11; //成交量
    required double turnover = 12; //成交額
    required double turnoverRate = 13; //換手率（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    required double amplitude = 14; //振幅（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    optional int32 darkStatus = 15; //DarkStatus, 暗盤交易狀態	
    optional OptionBasicQotExData optionExData = 16; //期權特有欄位
    optional double listTimestamp = 17; //上市日期時間戳（此欄位停止維護，不建議使用）
    optional double updateTimestamp = 18; //最新價的更新時間戳，對其他欄位不適用
    optional PreAfterMarketData preMarket = 19; //盤前資料
    optional PreAfterMarketData afterMarket = 20; //盤後資料
    optional int32 secStatus = 21; //SecurityStatus, 股票狀態
    optional FutureBasicQotExData futureExData = 22; //期貨特有欄位
}
```

## 盤前盤後資料

**PreAfterMarketData**
 
```protobuf
//美股支援盤前盤後資料
//科創板僅支援盤後資料：成交量，成交額
message PreAfterMarketData
{
    optional double price = 1;  // 盤前或盤後## 價格
    optional double highPrice = 2;  // 盤前或盤後## 最高價
    optional double lowPrice = 3;  // 盤前或盤後## 最低價
    optional int64 volume = 4;  // 盤前或盤後## 成交量
    optional double turnover = 5;  // 盤前或盤後## 成交額
    optional double changeVal = 6;  // 盤前或盤後## 漲跌額
    optional double changeRate = 7;  // 盤前或盤後## 漲跌幅（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    optional double amplitude = 8;  // 盤前或盤後## 振幅（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
}
```

## 分時資料

**TimeShare**

```protobuf
message TimeShare
{
    required string time = 1; //時間字串（格式：yyyy-MM-dd HH:mm:ss）
    required int32 minute = 2; //距離0點過了多少分鐘
    required bool isBlank = 3; //是否是空內容的點,若為 true 則只有時間資訊
    optional double price = 4; //當前價
    optional double lastClosePrice = 5; //昨收價
    optional double avgPrice = 6; //均價
    optional int64 volume = 7; //成交量
    optional double turnover = 8; //成交額
    optional double timestamp = 9; //時間戳
}
```

## 證券基本靜態資訊

**SecurityStaticBasic**

```protobuf

message SecurityStaticBasic
{
    required Qot_Common.Security security = 1; //股票
    required int64 id = 2; //股票 ID
    required int32 lotSize = 3; //每手數量,期權類型表示一份合約的股數
    required int32 secType = 4; //Qot_Common.SecurityType,股票類型
    required string name = 5; //股票名字
    required string listTime = 6; //上市時間字串（此欄位停止維護，不建議使用，格式：yyyy-MM-dd）
    optional bool delisting = 7; //是否退市
    optional double listTimestamp = 8; //上市時間戳（此欄位停止維護，不建議使用）
    optional int32 exchType = 9; //Qot_Common.ExchType,所屬交易所
}
```

## 窩輪額外靜態資訊
**WarrantStaticExData**

```protobuf
message WarrantStaticExData
{
    required int32 type = 1; //Qot_Common.WarrantType,窩輪類型
    required Qot_Common.Security owner = 2; //所屬正股
}    
```
## 期權額外靜態資訊

**OptionStaticExData**

```protobuf
message OptionStaticExData
{
    required int32 type = 1; //Qot_Common.OptionType,期權
    required Qot_Common.Security owner = 2; //標的股
    required string strikeTime = 3; //行權日（格式：yyyy-MM-dd）
    required double strikePrice = 4; //行權價
    required bool suspend = 5; //是否停牌
    required string market = 6; //發行市場名字
    optional double strikeTimestamp = 7; //行權日時間戳
    optional int32 indexOptionType = 8; //Qot_Common.IndexOptionType, 指數期權的類型，僅在指數期權有效
	optional int32 expirationCycle = 9; // ExpirationCycle，交割週期
    optional int32 optionStandardType = 10; // OptionStandardType，標準期權
    optional int32 optionSettlementMode = 11; // OptionSettlementMode，結算方式
}
```

## 期貨額外靜態資訊

**FutureStaticExData**

```protobuf
message FutureStaticExData
{
    required string lastTradeTime = 1; //最後交易日，只有非主連期貨合約才有該欄位
    optional double lastTradeTimestamp = 2; //最後交易日時間戳，只有非主連期貨合約才有該欄位
    required bool isMainContract = 3; //是否主連合約
}    
```

## 證券靜態資訊

**SecurityStaticInfo**

```protobuf
message SecurityStaticInfo
{
    required SecurityStaticBasic basic = 1; //證券基本靜態資訊
    optional WarrantStaticExData warrantExData = 2; //窩輪額外靜態資訊
    optional OptionStaticExData optionExData = 3; //期權額外靜態資訊
    optional FutureStaticExData futureExData = 4; //期貨額外靜態資訊
}
```

## 買賣經紀

**Broker**

```protobuf
message Broker
{
    required int64 id = 1; //經紀 ID
    required string name = 2; //經紀名稱
    required int32 pos = 3; //經紀檔位
    
    //以下為港股 SF 行情特有欄位
    optional int64 orderID = 4; //交易所訂單 ID，與交易介面返回的訂單 ID 並不一樣
    optional int64 volume = 5; //訂單股數
}
```

## 逐筆成交

**Ticker**

```protobuf
message Ticker
{
    required string time = 1; //時間字串（格式：yyyy-MM-dd HH:mm:ss）
    required int64 sequence = 2; // 唯一標識
    required int32 dir = 3; //TickerDirection, 買賣方向
    required double price = 4; //價格
    required int64 volume = 5; //成交量
    required double turnover = 6; //成交額
    optional double recvTime = 7; //收到推送資料的本地時間戳，用於定位延遲
    optional int32 type = 8; //TickerType, 逐筆類型
    optional int32 typeSign = 9; //逐筆類型符號
    optional int32 pushDataType = 10; //用於區分推送情況，僅推送時有該欄位
    optional double timestamp = 11; //時間戳
}	
```
## 買賣檔明細

**OrderBookDetail**

```protobuf
message OrderBookDetail
{
    required int64 orderID = 1; //交易所訂單 ID，與交易介面返回的訂單 ID 並不一樣
    required int64 volume = 2; //訂單股數
}
```

## 買賣檔

**OrderBook**

```protobuf
message OrderBook
{
    required double price = 1; //委託價格
    required int64 volume = 2; //委託數量
    required int32 orederCount = 3; //委託訂單個數
    repeated OrderBookDetail detailList = 4; //訂單資訊，港股 SF，美股深度擺盤特有
}
```

## 持股變動

**ShareHoldingChange**

```protobuf
message ShareHoldingChange
{
    required string holderName = 1; //持有者名稱（機構名稱 或 基金名稱 或 高管姓名）
    required double holdingQty = 2; //當前持股數量
    required double holdingRatio = 3; //當前持股百分比（該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%）
    required double changeQty = 4; //較上一次變動數量
    required double changeRatio = 5; //較上一次變動百分比（該欄位為百分比欄位，預設不展示 %，如20實際對應20%。是相對於自身的比例，而不是總的。如總股本1萬股，持有100股，持股百分比是1%，賣掉50股，變動比例是50%，而不是0.5%）
    required string time = 6; //發佈時間（格式：yyyy-MM-dd HH:mm:ss）
    optional double timestamp = 7; //時間戳
}
```

## 單個訂閲類型資訊

**SubInfo**

```protobuf
message SubInfo
{
    required int32 subType = 1;  //Qot_Common.SubType,訂閲類型
    repeated Qot_Common.Security securityList = 2; 	//訂閲該類型行情的證券
}	
```

## 單條連線訂閲資訊

**ConnSubInfo**

```protobuf
message ConnSubInfo
{
    repeated SubInfo subInfoList = 1; //該連線訂閲資訊
    required int32 usedQuota = 2; //該連線已經使用的訂閲額度
    required bool isOwnConnData = 3; //用於區分是否是自己連線的資料
}
```

## 板塊資訊

**PlateInfo**

```protobuf
message PlateInfo
{
    required Qot_Common.Security plate = 1; //板塊
    required string name = 2; //板塊名字
    optional int32 plateType = 3; //PlateSetType 板塊類型, 僅3207（獲取股票所屬板塊）協議返回該欄位
}
```

## 復權資訊

**Rehab**

```protobuf
message Rehab
{
    required string time = 1; //時間字串（格式：yyyy-MM-dd）
    required int64 companyActFlag = 2; //公司行動(CompanyAct)組合標誌位,指定某些欄位值是否有效
    required double fwdFactorA = 3; //前復權因子 A
    required double fwdFactorB = 4; //前復權因子 B
    required double bwdFactorA = 5; //後復權因子 A
    required double bwdFactorB = 6; //後復權因子 B
    optional int32 splitBase = 7; //拆股(例如，1拆5，Base 為1，Ert 為5)
    optional int32 splitErt = 8;	
    optional int32 joinBase = 9; //合股(例如，50合1，Base 為50，Ert 為1)
    optional int32 joinErt = 10;	
    optional int32 bonusBase = 11; //送股(例如，10送3, Base 為10,Ert 為3)
    optional int32 bonusErt = 12;	
    optional int32 transferBase = 13; //轉贈股(例如，10轉3, Base 為10,Ert 為3)
    optional int32 transferErt = 14;	
    optional int32 allotBase = 15; //配股(例如，10送2, 配股價為6.3元, Base 為10, Ert 為2, Price 為6.3)
    optional int32 allotErt = 16;	
    optional double allotPrice = 17;	
    optional int32 addBase = 18; //增發股(例如，10送2, 增發股價為6.3元, Base 為10, Ert 為2, Price 為6.3)
    optional int32 addErt = 19;	
    optional double addPrice = 20;	
    optional double dividend = 21; //現金分紅(例如，每10股派現0.5元,則該欄位值為0.05)
    optional double spDividend = 22; //特別股息(例如，每10股派特別股息0.5元,則該欄位值為0.05)
    optional double timestamp = 23; //時間戳
}
```

> - 公司行動組合標誌位參見 [CompanyAct](./quote.html#2729)

## 交割週期
**ExpirationCycle**

```protobuf
enum ExpirationCycle
{
    ExperationCycle_Unknow = 0; //未知
    ExperationCycle_Week = 1; //週期權
    ExperationCycle_Month = 2; //月期權
}
```


## 期權標準類型
**OptionStandardType**

```protobuf
enum OptionStandardType
{
    OptionStandardType_Unknown = 0; //未知
    OptionStandardType_Standard = 1; // 標準
    OptionStandardType_NonStandard = 2; // 非標準
}
```


## 期權結算方式
**OptionSettlementMode**

```protobuf
enum OptionSettlementMode
{
    OptionSettlementMode_Unknown = 0; //未知
    OptionSettlementMode_AM = 1; // AM
    OptionSettlementMode_PM = 2; // PM
}
```

## 股票持有者（已廢棄）

**HolderCategory**

```protobuf
enum HolderCategory
{
    HolderCategory_Unknow = 0; //未知
    HolderCategory_Agency = 1; //機構
    HolderCategory_Fund = 2; //基金
    HolderCategory_SeniorManager = 3; //高管
}
```

---



---

# 交易介面總覽

<table>
    <tr>
        <th>模組</th>
        <th>協定 ID</th>
        <th>Protobuf 文件</th>
        <th>說明</th>
    </tr>
    <tr>
        <td rowspan="2">帳戶</td>
        <td>2001</td>
	    <td><a href="../trade/get-acc-list.html">Trd_GetAccList</a></td>
	    <td>讀取交易業務帳戶列表</td>
    </tr>
    <tr>
        <td>2005</td>
	    <td><a href="../trade/unlock.html">Trd_UnlockTrade</a></td>
	    <td>解鎖或鎖定交易</td>
    </tr>
    <tr>
        <td rowspan="5">資產持倉</td>
        <td>2101</td>
	    <td><a href="../trade/get-funds.html">Trd_GetFunds</a></td>
	    <td>讀取帳戶資金</td>
    </tr>
    <tr>
        <td>2111</td>
	    <td><a href="../trade/get-max-trd-qtys.html">Trd_GetMaxTrdQtys</a></td>
	    <td>讀取最大交易數量</td>
    </tr>
    <tr>
        <td>2102</td>
	    <td><a href="../trade/get-position-list.html">Trd_GetPositionList</a></td>
	    <td>讀取帳戶持倉</td>
    </tr>
    <tr>
	    <td>2223</td>
	    <td><a href="../trade/get-margin-ratio.html">Trd_GetMarginRatio</a></td>
	    <td>讀取融資融券數據</td>
    </tr>
    <tr>
	    <td>2226</td>
        <td><a href="../trade/get-acc-cash-flow.html">Trd_FlowSummary</a></td>
	    <td>查詢帳戶現金流紀錄 (最低版本要求：9.1.5108)</td>
    </tr>
    <tr>
        <td rowspan="7">訂單</td>
        <td>2202</td>
	    <td><a href="../trade/place-order.html">Trd_PlaceOrder</a></td>
	    <td>下單</td>
    </tr>
    <tr>
        <td>2205</td>
	    <td><a href="../trade/modify-order.html">Trd_ModifyOrder</a></td>
	    <td>修改訂單</td>
    </tr>
    <tr>
        <td>2201</td>
	    <td><a href="../trade/get-order-list.html">Trd_GetOrderList</a></td>
	    <td>讀取訂單列表</td>
    </tr>
	<tr>
        <td>2225</td>
	    <td><a href="../trade/order-fee-query.html">Trd_GetOrderFee</a></td>
	    <td>讀取訂單費用 (最低版本要求：8.2.4218)</td>
    </tr>
    <tr>
        <td>2221</td>
	    <td><a href="../trade/get-history-order-list.html">Trd_GetHistoryOrderList</a></td>
	    <td>讀取歷史訂單列表</td>
    </tr>
    <tr>
        <td>2208</td>
	    <td><a href="../trade/update-order.html">Trd_UpdateOrder</a></td>
	    <td>推送訂單狀態變動通知</td>
    </tr>
    <tr>
        <td>2008</td>
	    <td><a href="../trade/sub-acc-push.html">Trd_SubAccPush</a></td>
	    <td>訂閱業務帳戶的交易推送數據</td>
    </tr>
    <tr>
        <td rowspan="3">成交</td>
        <td>2211</td>
	    <td><a href="../trade/get-order-fill-list.html">Trd_GetOrderFillList</a></td>
	    <td>讀取成交列表</td>
    </tr>
    <tr>
        <td>2222</td>
	    <td><a href="../trade/get-history-order-fill-list.html">Trd_GetHistoryOrderFillList</a></td>
	    <td>讀取歷史成交列表</td>
    </tr>
    <tr>
        <td>2218</td>
	    <td><a href="../trade/update-order-fill.html">Trd_UpdateOrderFill</a></td>
	    <td>推送成交通知</td>
    </tr>
</table>

---



---

# 交易物件

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>

## 建立連接

`OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, is_encrypt=None, security_firm=SecurityFirm.FUTUSECURITIES)`  
  
`OpenFutureTradeContext(host='127.0.0.1', port=11111, is_encrypt=None, security_firm=SecurityFirm.FUTUSECURITIES)` 


* **介紹**

    根據交易品類，選擇帳戶，並創建對應的交易物件。
    實例|帳戶
    :-|:-
    OpenSecTradeContext|證券帳戶  (股票、ETFs、窩輪牛熊、股票及指數的期權使用此帳戶)
    OpenFutureTradeContext|期貨帳戶   (期貨、期貨期權使用此帳戶)

* **參數**
    參數|類型|說明
    :-|:-|:-
    filter_trdmarket|[TrdMarket](./trade.html#1256)|篩選對應交易市場權限的帳戶  (- 此參數僅對 OpenSecTradeContext 適用
  - 此參數僅用於篩選帳戶，不影響交易連接)
    host|str|OpenD 監聽的 IP 地址
    port|int|OpenD 監聽的 IP 端口
    is_encrypt|bool|是否啟用加密  (預設 None 表示：使用 [enable_proto_encrypt](../ftapi/init.md#1542) 的設置)
    security_firm|[SecurityFirm](./trade.md#6222)|所屬券商

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, is_encrypt=None, security_firm=SecurityFirm.FUTUSECURITIES)
trd_ctx.close() # 結束後記得關閉當條連接，防止連接條數用盡
```


## 關閉連接

`close()`  

* **介紹**

    關閉交易物件。預設情況下，Futu API 內部創建的執行緒會阻止程序退出，只有當所有 Context 都 close 後，程序才能正常退出。但通過 [set_all_thread_daemon](../ftapi/init.md#7809) 可以設置所有內部執行緒為 daemon 執行緒，這時即使沒有調用 Context 的 close，程序也可以正常退出。

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111)
trd_ctx.close()  # 結束後記得關閉當條連接，防止連接條數用盡
```

---



---

# 讀取交易業務賬户列表

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_acc_list()`

* **介紹**

    讀取交易業務賬户列表。  
    要執行其他交易介面前，請先讀取此列表，確認要操作的交易業務賬户無誤。

* **參數**
    


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>接口調用結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回交易業務賬户列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 交易業務賬户列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        acc_id|int|交易業務賬户
        trd_env|[TrdEnv](./trade.md#9175)|交易環境
        acc_type|[TrdAccType](./trade.md#5161)|賬户類型
        uni_card_num|str|綜合賬户卡號，同流動裝置內的展示
        card_num|str|業務賬户卡號  (綜合賬户下包含一個或多個業務賬户（綜合證券賬户、綜合期貨賬户等等），與交易品種有關)
        security_firm|[SecurityFirm](./trade.md#6222)|所屬券商
        sim_acc_type|[SimAccType](./trade.md#6090)|模擬賬户類型  (僅模擬賬户適用) 
        trdmarket_auth|list|交易市場權限  (list 中元素類型是 [TrdMarket](./trade.html#1256)) 
        acc_status|[TrdAccStatus](./trade.md#3606)|賬户狀態
        acc_role|[TrdAccRole](./trade.md#4604)|賬户結構  (用於區分主子賬户結構
  - MASTER: 主賬户
  - NORMAL: 普通賬户
  - IPO: 馬來西亞 IPO 賬户)
        jp_acc_type|list|日本賬户類型  (list 中元素類型是[SubAccType](./trade.md#6090)，僅對日本券商生效)


* **說明**

    獲取港股模擬交易帳戶，需要指定 filter_trdmarket 爲 TrdMarket.HK，此時會返回2個模擬交易賬號。其中 sim_acc_type = STOCK 爲港股模擬帳戶，sim_acc_type = OPTION 爲港股期權模擬帳戶，sim_acc_type = FUTURES 爲港股期貨模擬帳戶。   
    獲取美股模擬交易帳戶，需要指定 filter_trdmarket 爲 TrdMarket.US，sim_acc_type = STOCK_AND_OPTION 代表美股融資融券模擬帳戶，可以模擬交易股票和期權。sim_acc_type = FUTURES 爲美國期貨模擬帳戶。


* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_acc_list()
if ret == RET_OK:
    print(data)
    print(data['acc_id'][0])  # 取第一個賬號
    print(data['acc_id'].values.tolist())  # 轉為 list
else:
    print('get_acc_list error: ', data)
trd_ctx.close()
```

* **Output**

```python
               acc_id   trd_env acc_type       uni_card_num           card_num    security_firm   sim_acc_type                           trdmarket_auth    acc_status    acc_role    jp_acc_type
0  281756479345015383      REAL   MARGIN   1001289516908051   1001329805025007   FUTUSECURITIES            N/A    [HK, US, HKCC, SG, HKFUND, USFUND, JP]       ACTIVE      NORMAL             []
1             8377516  SIMULATE     CASH                N/A                N/A              N/A          STOCK                                      [HK]       ACTIVE         N/A             []
2            10741586  SIMULATE   MARGIN                N/A                N/A              N/A         OPTION                                      [HK]       ACTIVE         N/A             []
3  281756455983234027      REAL   MARGIN                N/A   1001100321720699   FUTUSECURITIES            N/A                                      [HK]     DISABLED      NORMAL             []
281756479345015383
[281756479345015383, 8377516, 10741586, 281756455983234027]
```

---



---

# 解鎖交易

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`unlock_trade(password=None, password_md5=None, is_unlock=True)`

* **介紹**

    解鎖或鎖定交易

* **參數**
    
    參數|類型|說明
    :-|:-|:-
    password|str|交易密碼  (如果 password_md5 不為空，就使用傳入的 password_md5 解鎖；否則使用 password 轉 MD5 得到 password_md5 再解鎖)
    password_md5|str|交易密碼的 32 位 MD5 加密（全小寫） (解鎖交易必須要填密碼，鎖定交易忽略)
    is_unlock|bool|解鎖或鎖定  (True：解鎖False：鎖定)


* **回傳**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">msg</td>
            <td>NoneType</td>
            <td>當 ret == RET_OK 時，回傳 None</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，回傳錯誤描述</td>
        </tr>
    </table>

        

* **Example**

```python
from futu import *
pwd_unlock = '123456'
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.unlock_trade(pwd_unlock)
if ret == RET_OK:
    print('unlock success!')
else:
    print('unlock_trade failed: ', data)
trd_ctx.close()
```

* **Output**

```python
unlock success!
```

---



---

# 查詢帳戶資金

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`accinfo_query(trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, refresh_cache=False, currency=Currency.HKD, asset_category=AssetCategory.NONE)`

* **介紹**

    查詢交易業務帳戶的資產淨值、證券市值、現金、購買力等資金數據。

* **參數**
    參數|類型|說明
    :-|:-|:-
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務帳戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳入 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)
    refresh_cache|bool|是否更新快取  (- True：立即向富途伺服器重新請求數據，不使用 OpenD 的緩存，此時會受到介面頻率限制的限制
  - False：使用 OpenD 的緩存（特殊情況導致緩存沒有及時更新才需要刷新）)
    currency|[Currency](./trade.md#3345)|計價貨幣  (- 僅期貨帳戶、綜合證券帳戶適用，其它帳戶類型會忽略此參數
  - 返回的 DataFrame 中，除了明確指明瞭貨幣的欄位，其它資金相關字段都以此參數換算)
    asset_category|[AssetCategory](./trade.md#1879)|資產類別  (僅對日本證券商生效)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回資金數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 資金數據格式如下：
        字段|類型|說明
        :-|:-|:-
        power|float|最大購買力  (- 此字段是按照 50% 的融資初始保證金率計算得到的 **近似值**。但事實上，每個標的的融資初始保證金率並不相同。我們建議您使用 [查詢最大可買可賣](./get-max-trd-qtys.md) 接口返回的 **最大可買** 字段，來判斷實際可買入的最大數量。)
        max_power_short|float|賣空購買力  (- 此字段是按照 60% 的融券保證金率計算得到的 **近似值**。但事實上，每個標的的融券保證金率並不相同。我們建議您使用 [查詢最大可買可賣](./get-max-trd-qtys.md) 接口返回的 **可賣空** 字段，來判斷實際可賣空的最大數量。)
        net_cash_power|float|現金購買力 (已廢棄，請使用usd_net_cash_power等字段獲取分幣種的現金購買力)
        total_assets|float|總資產淨值 (總資產淨值 = 證券資產淨值 + 基金資產淨值 + 債券資產淨值) 
        securities_assets|float|證券資產淨值 (最低 OpenD 版本要求：8.2.4218) 
        fund_assets|float|基金資產淨值 (- 綜合帳戶返回結果為總基金資產淨值，暫時不支援查詢港元基金資產和美元基金資產
  - 最低 OpenD 版本要求：8.2.4218)
        bond_assets|float|債券資產淨值 (最低 OpenD 版本要求：8.2.4218) 
        cash|float|現金 (已廢棄，請使用us_cash等字段獲取分幣種的現金)
        market_val|float|證券市值  (僅證券帳戶適用)
        long_mv|float|長倉市值  
        short_mv|float|短倉市值  
        pending_asset|float|未交收資產  
        interest_charged_amount|float|計息金額 
        frozen_cash|float|凍結資金
        avl_withdrawal_cash|float|現金可提  (僅證券帳戶適用)
        max_withdrawal|float|最大可提  (僅富途證券（香港）的證券帳戶適用) 
        currency|[Currency](./trade.md#3345)|計價貨幣  (僅綜合證券帳戶、期貨帳戶適用)
        available_funds|float|可用資金  (僅期貨帳戶適用)
        unrealized_pl|float|未實現盈虧  (僅期貨帳戶適用)
        realized_pl|float|已實現盈虧  (僅期貨帳戶適用)
        risk_level|[CltRiskLevel](./trade.md#9117)|風險管理狀態  (僅期貨帳戶適用。建議統一使用 risk_status 字段獲取證券、期貨帳戶的風險狀態)
        risk_status|[CltRiskStatus](./trade.md#9117)|風險狀態  (- 證券帳戶和期貨帳戶均適用
  - 共分 9 個等級， `LEVEL1`是最安全，`LEVEL9`是最危險)
        initial_margin|float|初始保證金 
        margin_call_margin|float|Margin Call 保證金 
        maintenance_margin|float|維持保證金 
        hk_cash|float|港元現金  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        hk_avl_withdrawal_cash|float|港元可提  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        hkd_net_cash_power|float|港元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        hkd_assets|float|港股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        us_cash|float|美元現金  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        us_avl_withdrawal_cash|float|美元可提  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        usd_net_cash_power|float|美元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        usd_assets|float|美股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        cn_cash|float|人民幣現金  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        cn_avl_withdrawal_cash|float|人民幣可提  (此字段表示該幣種實際的值，而不是以該幣種計價的值)
        cnh_net_cash_power|float|人民幣現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        cnh_assets|float|A股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        jp_cash|float|日元現金  (- 僅期貨帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低 Futu API 版本要求：5.8.2008)
        jp_avl_withdrawal_cash|float|日元可提  (- 僅期貨帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低 Futu API 版本要求：5.8.2008)
        jpy_net_cash_power|float|日元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        jpy_assets|float|日股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        sg_cash|float|新元現金  (- 僅期貨帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值)
        sg_avl_withdrawal_cash|float|新元可提  (- 僅期貨帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值)
        sgd_net_cash_power|float|新元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        sgd_assets|float|新股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        au_cash|float|澳元現金  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低 Futu API 版本要求：5.8.2008)
        au_avl_withdrawal_cash|float|澳元可提  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低 Futu API 版本要求：5.8.2008)
        aud_net_cash_power|float|澳元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：8.7)
        aud_assets|float|澳股資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：9.0.5008)
        ca_cash|float|加元現金  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        ca_avl_withdrawal_cash|float|加元可提  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        cad_net_cash_power|float|加元現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        cad_assets|float|加元資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        my_cash|float|令吉現金  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        my_avl_withdrawal_cash|float|令吉可提  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        myr_net_cash_power|float|令吉現金購買力  (- 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        myr_assets|float|令吉資產淨值  (- 僅綜合證券帳戶適用
  - 此字段表示該幣種實際的值，而不是以該幣種計價的值
  - 最低版本要求：10.0.6008)
        is_pdt|bool|是否為 PDT 帳戶  (True：是 PDT 帳戶，False：不是 PDT 帳戶僅moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)
        pdt_seq|string|剩餘日內交易次數  (僅moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)   
        beginning_dtbp|float|初始日內交易購買力  (僅被標記為 PDT 的moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)
        remaining_dtbp|float|剩餘日內交易購買力  (僅被標記為 PDT 的moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)
        dt_call_amount|float|日內交易待繳金額  (僅被標記為 PDT 的moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)
        dt_status|[DtStatus](./trade.html#9841)|日內交易限制情況  (僅被標記為 PDT 的moomoo證券(美國)帳戶適用最低 OpenD 版本要求：5.8.2008)

        
* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.accinfo_query()
if ret == RET_OK:
    print(data)
    print(data['power'][0])  # 取第一行的購買力
    print(data['power'].values.tolist())  # 轉為 list
else:
    print('accinfo_query error: ', data)
trd_ctx.close()  # 關閉當條連接
```

* **Output**

 ```python
power  max_power_short  net_cash_power  total_assets  securities_assets  fund_assets  bond_assets   cash   market_val      long_mv   short_mv  pending_asset  interest_charged_amount  frozen_cash  avl_withdrawal_cash  max_withdrawal currency available_funds unrealized_pl realized_pl risk_level risk_status  initial_margin  margin_call_margin  maintenance_margin  hk_cash  hk_avl_withdrawal_cash  hkd_net_cash_power  hkd_assets  us_cash  us_avl_withdrawal_cash  usd_net_cash_power  usd_assets  cn_cash  cn_avl_withdrawal_cash  cnh_net_cash_power  cnh_assets  jp_cash  jp_avl_withdrawal_cash  jpy_net_cash_power jpy_assets  sg_cash sg_avl_withdrawal_cash sgd_net_cash_power sgd_assets  au_cash au_avl_withdrawal_cash aud_net_cash_power aud_assets  ca_cash ca_avl_withdrawal_cash cad_net_cash_power cad_assets  my_cash my_avl_withdrawal_cash myr_net_cash_power myr_assets  is_pdt pdt_seq beginning_dtbp remaining_dtbp dt_call_amount dt_status
0  465453.903307    465453.903307             0.0   289932.0404        197028.2204     92903.82          0.0  25.18  197003.0448  211960.7568 -14957.712            0.0                      0.0    25.930845                  0.0             0.0      HKD             N/A           N/A         N/A        N/A      LEVEL3   219346.648525       288656.787955       181250.967601      0.0                     0.0          13225.7955     0.0   3.24                     0.0           9656.4365      0.0    0.0                     0.0                 0.0    0.0      0.0                     0.0                 0.0     0.0    N/A                    N/A                N/A     0.0    N/A                    N/A                N/A    0.0    N/A                    N/A                N/A    0.0    N/A                    N/A                N/A    0.0        N/A     N/A            N/A            N/A            N/A       N/A
465453.903307
[465453.903307]
```

---



---

# 查詢最大可買可賣

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`acctradinginfo_query(order_type, code, price, order_id=None, adjust_limit=0, trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, session=Session.NONE, jp_acc_type=SubAccType.JP_GENERAL, position_id=NONE)`

* **介紹**

    查詢指定交易業務帳戶下的最大可買賣數量，亦可查詢指定交易業務帳戶下指定訂單的最大可改成的數量。

    現金帳戶請求期權不適用。

* **參數**
    參數|類型|說明
    :-|:-|:-
    order_type|[OrderType](./trade.md#379)|訂單類型
    code|str|證券代碼  (如果是期貨交易，且 code 為期貨主連代碼，則會自動轉為對應的實際合約代碼)
    price|float|報價  (證券帳戶精確到小數點後 3 位，超出部分會被捨棄期貨帳戶精確到小數點後 9 位，超出部分會被捨棄)
    order_id|str|訂單號  (- 預設傳 None，查詢的是新下單的最大可買可賣數量
  - 如果是改單則要傳訂單號，此時計算最大可買可賣時，會回傳此訂單可改成的最大數量
  - 如果通過此參數，查詢某筆訂單最大可改成的數量，需要在下單之後，間隔 0.5 秒以上再執行此介面)
    adjust_limit|float|價格微調幅度  (OpenD 會對傳入價格自動調整到合法價位上（期貨會忽略此參數）
  - 正數代表向上調整，負數代表向下調整
  - 例如：0.015 代表向上調整且幅度不超過 1.5%；-0.01 代表向下調整且幅度不超過 1%。預設 0 表示不調整)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務帳戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)
    session|[Session](../quote/quote.md#8417)|美股交易時段  (僅對美股生效，支援傳入RTH、ETH、OVERNIGHT、ALL)
    jp_acc_type|[SubAccType](./trade.md#5161)|日本帳戶類型  (僅日本券商適用)
    position_id|int|持倉ID  (- 適用於日本衍生品帳戶查詢持倉可賣和平倉需買回
  - 可通過[查詢持倉](./get-position-list.md)介面獲取)
    


* **回傳**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，回傳賬號列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，回傳錯誤描述</td>
        </tr>
    </table>

    * 賬號列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        max_cash_buy|float|現金可買  (-  期權的單位是“張”
  - 期貨帳戶不適用)
        max_cash_and_margin_buy|float|最大可買  (-  期權的單位是“張”
  - 期貨帳戶不適用)
        max_position_sell|float|持倉可賣  (期權的單位是"張")
        max_sell_short|float|可賣空  (-  期權的單位是“張”
  - 期貨帳戶不適用)
        max_buy_back|float|平倉需買入  (- 當持有淨空倉時，必須先買回空頭持倉的股數，才能再繼續買多
  -  期貨、期權的單位是“張”)
        long_required_im|float|買 1 張合約所帶來的初始保證金變動  (-  當前僅期貨和期權適用。
  - 無持倉時，回傳 **買入** 1 張的初始保證金佔用（正數）。 
  - 有多倉時，回傳 **買入** 1 張的初始保證金佔用（正數）。
  - 有空倉時，回傳 **買回** 1 張的初始保證金釋放（負數）。)
        short_required_im|float|賣 1 張合約所帶來的初始保證金變動  (-  當前僅期貨和期權適用。
  - 無持倉時，回傳 **賣空** 1 張的初始保證金佔用（正數）。 
  - 有多倉時，回傳 **賣出** 1 張的初始保證金釋放（負數）。
  -  有空倉時，回傳 **賣空** 1 張的初始保證金釋放（正數）。)
        session|[Session](../quote/quote.md#8417)|交易訂單時段（僅用於美股）

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.acctradinginfo_query(order_type=OrderType.NORMAL, code='HK.00700', price=400)
if ret == RET_OK:
    print(data)
    print(data['max_cash_and_margin_buy'][0])  # 最大融資可買數量
else:
    print('acctradinginfo_query error: ', data)
trd_ctx.close()  # 關閉當條連接
```

* **Output**

```python
    max_cash_buy  max_cash_and_margin_buy  max_position_sell  max_sell_short  max_buy_back long_required_im short_required_im    session
0           0.0                   1500.0                0.0             0.0           0.0              N/A               N/A             N/A
1500.0
```

---



---

# 查詢持倉

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`position_list_query(code='', position_market=TrdMarket.NONE, pl_ratio_min=None, pl_ratio_max=None, trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, refresh_cache=False, asset_category=AssetCategory.NONE)`

* **介紹**

    查詢交易業務賬户的持倉列表

* **參數**
    參數|類型|說明
    :-|:-|:-
    code|str|代碼過濾  (- 只回傳此代碼對應的持倉數據。不傳則回傳所有
  - 注意：期貨持倉的代碼過濾，需要傳入含具體月份的合約代碼，無法通過主連合約代碼進行過濾)
    position_market| [TrdMarket](./trade.md#1256)|持倉所屬市場過濾 (- 回傳指定市場的持倉數據
  - 預設狀態時，回傳所有市場持倉數據)
    pl_ratio_min|float|當前盈虧比例下限過濾，僅回傳高於此比例的持倉  (證券賬户使用攤薄成本價的盈虧比例，期貨賬户使用平均成本價的盈虧比例例如：傳入 10，則回傳盈虧比例大於 +10% 的持倉)
    pl_ratio_max|float|當前盈虧比例上限過濾，低於此比例的會回傳  (證券賬户使用攤薄成本價的盈虧比例，期貨賬户使用平均成本價的盈虧比例例如：傳入 10，回傳盈虧比例小於 +10% 的持倉)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務賬户 ID  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的賬户為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的賬户為準)
    acc_index|int|交易業務賬户列表中的賬户序號  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷賬户時發生變動，導致您指定的賬户與實際交易賬户不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務賬户)
    refresh_cache|bool|是否更新快取  (- True：立即向富途伺服器重新請求數據，不使用 OpenD 的快取，此時會受到介面頻率限制的限制
  - False：使用 OpenD 的快取（特殊情況導致快取沒有及時更新才需要更新）)
    asset_category|[AssetCategory](./trade.md#1879)|資產類別  (僅對日本券商生效)
    


* **回傳**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>接口執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，回傳持倉列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，回傳錯誤描述</td>
        </tr>
    </table>

    * 持倉列表
        欄位|類型|說明
        :-|:-|:-
        position_side|[PositionSide](./trade.md#5184)|持倉方向
        code|str|股票編號
        stock_name|str|股票名稱
        position_market|[TrdMarket](./trade.md#1256)|持倉所屬市場
        qty|float|持有數量  (期權和期貨的單位是“張”)
        can_sell_qty|float|可用數量  (可用數量，是指持有的可平倉的數量。可用數量=持有數量-凍結數量期權和期貨的單位是“張”。)
        currency|[Currency](./trade.md#3345)|交易貨幣
        nominal_price|float|市價  (精確到小數點後 3 位，超出部分四捨五入)
        cost_price|float|攤薄成本價（證券賬户），平均開倉價（期貨賬户）  (建議使用 average_cost，diluted_cost 欄位獲取持倉成本價)
        cost_price_valid|bool|成本價是否有效  (True：有效False：無效)
        average_cost|float|平均成本價  (模擬證券賬户不適用最低OpenD版本要求：9.2.5208)
        diluted_cost|float|攤薄成本價  (期貨賬户不適用最低OpenD版本要求：9.2.5208)
        market_val|float|市值  (精度：3 位小數（A 股 2 位小數，期貨 0 位小數）)
        pl_ratio|float|盈虧比例（攤薄成本價模式）  (期貨不適用該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%)
        pl_ratio_valid|bool|盈虧比例是否有效  (True：有效False：無效)
        pl_ratio_avg_cost|float|盈虧比例（平均成本價模式）  (模擬證券賬户不適用該欄位為百分比欄位，預設不展示 %，如 20 實際對應 20%最低OpenD版本要求：9.2.5208)
        pl_val|float|盈虧金額  (精度：3 位小數（A 股 2 位小數）)
        pl_val_valid|bool|盈虧金額是否有效  (True：有效False：無效)
        today_pl_val|float|今日盈虧金額  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數，期貨 2 位小數）)
        today_trd_val|float|今日交易金額  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數）期貨不適用)
        today_buy_qty|float|今日買入總量  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數）期貨不適用)
        today_buy_val|float|今日買入總額  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數）期貨不適用)
        today_sell_qty|float|今日賣出總量  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數）期貨不適用)
        today_sell_val|float|今日賣出總額  (只在真實交易環境下有效精度：3 位小數（A 股 2 位小數）期貨不適用)
        unrealized_pl|float|未實現盈虧  (模擬證券賬户不適用綜合證券賬户，回傳平均成本價模式下的未實現盈虧金額)
        realized_pl|float|已實現盈虧  (模擬證券賬户不適用綜合證券賬户，回傳平均成本價模式下的已實現盈虧金額)
        position_id|int|持倉ID

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.position_list_query()
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果持倉列表不為空
        print(data['stock_name'][0])  # 獲取持倉第一個股票名稱
        print(data['stock_name'].values.tolist())  # 轉為 list
else:
    print('position_list_query error: ', data)
trd_ctx.close()  # 關閉當條連接
```

* **Output**

```python
       code stock_name position_market    qty  can_sell_qty  cost_price  cost_price_valid average_cost  diluted_cost  market_val  nominal_price  pl_ratio  pl_ratio_valid pl_ratio_avg_cost  pl_val  pl_val_valid today_buy_qty today_buy_val today_pl_val today_trd_val today_sell_qty today_sell_val position_side unrealized_pl realized_pl currency asset_category position_id
0  HK.01810     小米集團-W              HK  400.0         400.0      53.975              True          53.975        53.975     19820.0          49.55  -8.19824            True            -8.19824    -1770.0          True           0.0           0.0          0.0           0.0            0.0            0.0          LONG           0.0         0.0      HKD      N/A      6596101776329286054
小米集團-W
['小米集團-W']
```

---



---

# 讀取融資融券數據

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_margin_ratio(code_list)`

* **介紹**

    查詢股票的融資融券數據。

* **參數**
    參數|類型|説明
    :-|:-|:-
    code_list|list|股票代碼列表  (每次最多可請求 100 個標的list 內元素類型為 str)
    


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，傳回融資融券數據</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，傳回錯誤描述</td>
        </tr>
    </table>

    * 融資融券數據格式如下：
        欄位|類型|説明
        :-|:-|:-
        code| str| 股票代碼
        is_long_permit|bool|是否允許融資
        is_short_permit | bool | 是否允許融券
        short_pool_remain | float | 賣空池剩餘  (單位：股)
        short_fee_rate | float | 融券參考利率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        alert_long_ratio | float | 融資預警比率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        alert_short_ratio | float | 融券預警比率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        im_long_ratio | float | 融資初始保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        im_short_ratio | float | 融券初始保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        mcm_long_ratio | float | 融資 margin call 保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        mcm_short_ratio | float  | 融券 margin call 保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        mm_long_ratio |float | 融資維持保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)
        mm_short_ratio |float | 融券維持保證金率  (該字段為百分比字段，預設不展示 %，如 20 實際對應 20%)

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_margin_ratio(code_list=['HK.00700','HK.09988'])  
if ret == RET_OK:
    print(data)
    print(data['is_long_permit'][0])  # 取第一條的是否允許融資
    print(data['im_short_ratio'].values.tolist())  # 轉為 list
else:
    print('error:', data)
trd_ctx.close()  # 結束後記得關閉當條連接，防止連接條數用盡
```

* **Output**

```python
       code  is_long_permit  is_short_permit  short_pool_remain  short_fee_rate  alert_long_ratio  alert_short_ratio  im_long_ratio  im_short_ratio  mcm_long_ratio  mcm_short_ratio  mm_long_ratio  mm_short_ratio
0  HK.00700            True             True          1826900.0            0.89              33.0               56.0           35.0            60.0            32.0             53.0           25.0            40.0
1  HK.09988            True             True          1150600.0            0.95              48.0               46.0           50.0            50.0            47.0             43.0           40.0            30.0
True
[60.0, 50.0]
```

---



---

# 查詢帳戶現金流紀錄

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`get_acc_cash_flow(clearing_date='', trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, cashflow_direction=CashFlowDirection.NONE)`

* **介紹**

    查詢交易業務帳戶在指定日期的現金流紀錄數據。數據覆蓋存款及提款、調撥、貨幣兌換、買賣金融資產、融資融券利息等所有導致現金變動的事項。

* **參數**
    
    參數|類型|說明
    :-|:-|:-
    clearing_date|str|清算日期 (- 如需查詢多日，需逐日請求
  - 格式：yyyy-MM-dd，例如：“2017-06-20”)
    trd_env|TrdEnv|交易環境
    acc_id|int|交易業務帳戶 ID   (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，建議使用 acc_id。
  - 當 acc_id 傳入 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號
    cashflow_direction|[CashFlowDirection](./trade.md#7263)|篩選現金流方向

* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>API 執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回交易業務帳戶現金流紀錄列表格式</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，傳回錯誤描述</td>
        </tr>
    </table>

    * 交易業務帳戶現金流紀錄列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        cashflow_id|int|現金流ID
        clearing_date|str|清算日期
        settlement_date|str|交收日期
        currency|[Currency](./trade.md#5161)|貨幣種類
        cashflow_type|str|現金流類型
        cashflow_direction|[CashFlowDirection](./trade.md#7263)|現金流方向
        cashflow_amount|float|金額（正數表示流入，負數表示流出）
        cashflow_remark|str|備註


* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_acc_cash_flow(clearing_date='2025-02-18', trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, cashflow_direction=CashFlowDirection.NONE)
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果現金流紀錄列表不為空
        print(data['cashflow_type'][0])  # 獲取第一條流水的現金流類型
        print(data['cashflow_amount'].values.tolist())  # 轉為 list
else:
    print('get_acc_cash_flow error: ', data)
trd_ctx.close()

```

* **Output**

```python
   cashflow_id     clearing_date     settlement_date     currency     cashflow_type     cashflow_direction     cashflow_amount     cashflow_remark
0  16308           2025-02-27        2025-02-28          HKD             其他                 N/A                   0.00      Opt ASS-P-JXC250227P13000-20250227
1  16357           2025-02-27        2025-03-03          HKD             其他                 OUT               -104000.00
2  16360           2025-02-27        2025-02-27          USD            基金贖回               IN                 23000.00     Fund Redemption#Taikang Kaitai US Dollar Money...
3  16384           2025-02-27        2025-02-27          HKD            基金贖回               IN                104108.96     Fund Redemption#Taikang Kaitai Hong Kong Dolla...
其他
[0.00, -104000.00, 23000.00, 104108.96]
```

---



---

# 下單

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`place_order(price, qty, code, trd_side, order_type=OrderType.NORMAL, adjust_limit=0, trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, remark=None, time_in_force=TimeInForce.DAY,  fill_outside_rth=False, aux_price=None, trail_type=None, trail_value=None, trail_spread=None, session=Session.NONE, jp_acc_type=SubAccType.JP_GENERAL, position_id=NONE)`

* **介紹**

    下單 
    :::tip 提示
    Python API 是同步的，但網絡收發是非同步的。當 place_order 對應的應答數據包與 [響應成交推送回呼](../trade/update-order-fill.md) 或 [響應訂單推送回呼](../trade/update-order.md) 間隔很短時，就可能出現 place_order 的數據包先返回，但回呼函數先被執行的情況。例如：可能先執行了 [響應訂單推送回呼](../trade/update-order.md)，然後 place_order 這個介面才返回。
    :::

* **參數**

    參數|類型|說明
    :-|:-|:-
    price|float|訂單價格  (- 當訂單是市價單或競價單類型，仍需對 price 傳遞參數，price 可以傳入任意值
  - 精度：
  - 期貨：整數8位，小數9位，支援負數價格
  - 美股期權：小數2位
  - 美股：不超過$1，允許小數4位
  - 其他：小數3位，超出部分四捨五入)
    qty|float|訂單數量  (期權期貨單位是"張")
    code|str|標的代碼  (如果 code 為期貨主連代碼，則會自動轉為實際對應的合約代碼)
    trd_side|[TrdSide](./trade.md#5815)|交易方向
    order_type|[OrderType](./trade.md#379)|訂單類型
    adjust_limit|float|價格微調幅度  (OpenD 會對傳入價格自動調整到合法價位上
  - 正數代表向上調整，負數代表向下調整
  - 例如：0.015 代表向上調整且幅度不超過 1.5%；-0.01 代表向下調整且幅度不超過 1%。預設 0 表示不調整)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務賬户 ID  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的賬户為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的賬户為準)
    acc_index|int|交易業務賬户列表中的賬户序號  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷賬户時發生變動，導致您指定的賬户與實際交易賬户不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務賬户)
    remark|str|備註  (- 訂單會帶上此備註欄位，方便您標記訂單
  - 轉成 utf8 後的長度上限為 64 位元組)
    time_in_force|[TimeInForce](./trade.md#114)|有效期限  (香港市場、A 股市場和環球期貨的市價單，僅支援當日有效)
    fill_outside_rth|bool|是否允許盤前盤後  (用於港股盤前競價與美股盤前盤後，且盤前盤後時段不支援市價單)
    aux_price|float|觸發價格  (- 當訂單是止損市價單、止損限價單、觸及限價單（止盈）、觸及市價單（止盈） 時，aux_price 為必傳遞參數數
  - 同price精度，超過部分四捨五入)
    trail_type|[TrailType](./trade.md#589)|跟蹤類型  (當訂單是跟蹤止損市價單、跟蹤止損限價單時，trail_type 為必傳遞參數數)
    trail_value|float|跟蹤金額/百分比  (- 當訂單是跟蹤止損市價單、跟蹤止損限價單時，trail_value 為必傳遞參數數
  - 當跟蹤類型為比例時，該欄位為百分比欄位，傳入 20 實際對應 20%
  - 當跟蹤類型為金額時，整數部分同price；小數部分美股期權固定2位，美股4位，其他同price；超過部分四捨五入
  - 當跟蹤類型為比例時，精確到小數點後 2 位，整數部分同price，超過部分四捨五入)
    trail_spread|float|指定價差  (- 當訂單是跟蹤止損限價單時，trail_spread 為必傳遞參數數
  - 證券賬户精確到小數點後 3 位，期貨賬户精確到小數點後 9 位，超過部分四捨五入)
    session|[Session](../quote/quote.md#8417)|美股交易時段  (僅對美股生效，支援傳入RTH、ETH、OVERNIGHT、ALL)
    jp_acc_type|[SubAccType](./trade.md#5752)|日本賬户類型  (僅日本券商適用)
    position_id|int|持倉ID  (- 日本券商平倉時需要填寫
  - 可通過[查詢持倉](./get-position-list.md)介面獲取)


* **返回**
    
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回訂單列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 訂單列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        order_type|[OrderType](./trade.md#379)|訂單類型
        order_status|[OrderStatus](./trade.md#3177)|訂單狀態
        order_id|str|訂單號
        code|str|股票代碼
        stock_name|str|股票名稱
        qty|float|訂單數量  (期權期貨單位是"張")
        price|float|訂單價格  (精確到小數點後 3 位，超出部分四捨五入)
        create_time|str|創建時間  (格式：yyyy-MM-dd HH:mm:ss
期貨時區指定，請參見 [FutuOpenD 配置](../quick/opend-base.md#3795))
        updated_time|str|最後更新時間  (格式：yyyy-MM-dd HH:mm:ss
期貨時區指定，請參見 [FutuOpenD 配置](../quick/opend-base.md#3795))
        dealt_qty|float|成交數量  (期權期貨單位是"張")
        dealt_avg_price|float|成交均價  (無精度限制)
        last_err_msg|str|最後的錯誤描述  (如果有錯誤，會返回最後一次錯誤的原因如果無錯誤，返回空字符串)
        remark|str|下單時備註的標記  (詳見 [place_order](./place-order.md) 介面參數中的 remark)
        time_in_force|[TimeInForce](./trade.md#114)|有效期限
        fill_outside_rth|bool|是否允許盤前盤後（用於港股盤前競價與美股盤前盤後）  (True：允許False：不允許)
        session|[Session](../quote/quote.md#8417)|交易訂單時段（僅用於美股）
        aux_price|float|觸發價格
        trail_type|[TrailType](./trade.md#589)|跟蹤類型
        trail_value|float|跟蹤金額/百分比
        trail_spread|float|指定價差
        

* **Example**

```python
from futu import *
pwd_unlock = '123456'
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.unlock_trade(pwd_unlock)  # 若使用真實賬户下單，需先對賬户進行解鎖。此處示例為模擬賬户下單，也可省略解鎖。
if ret == RET_OK:
    ret, data = trd_ctx.place_order(price=510.0, qty=100, code="HK.00700", trd_side=TrdSide.BUY, trd_env=TrdEnv.SIMULATE, session=Session.NONE)
    if ret == RET_OK:
        print(data)
        print(data['order_id'][0])  # 獲取下單的訂單號
        print(data['order_id'].values.tolist())  # 轉為 list
    else:
        print('place_order error: ', data)
else:
    print('unlock_trade failed: ', data)
trd_ctx.close()
```

* **Output**

```python

       code stock_name trd_side order_type order_status           order_id    qty  price          create_time         updated_time  dealt_qty  dealt_avg_price last_err_msg remark time_in_force fill_outside_rth session aux_price trail_type trail_value trail_spread currency
0  HK.00700       騰訊控股      BUY     NORMAL   SUBMITTING  38196006548709500  100.0  420.0  2021-11-04 11:38:19  2021-11-04 11:38:19        0.0              0.0                               DAY              N/A       N/A    N/A      N/A         N/A          N/A      HKD
38196006548709500
['38196006548709500']
```

---



---

# 改單撤單

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`modify_order(modify_order_op, order_id, qty, price, adjust_limit=0, trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, aux_price=None, trail_type=None, trail_value=None, trail_spread=None)`

* **介紹**

    修改訂單的價格和數量、撤單、操作訂單的失效和生效、刪除訂單等。  
	如果是 A 股通市場，將不支援改單。可撤單。刪除訂單是 OpenD 本地操作。

* **參數**
    參數|類型|說明
    :-|:-|:-
    modify_order_op|[ModifyOrderOp](./trade.md#8243)|改單操作類型
    order_id|str|訂單號
    qty|float|訂單改單後的數量  (期權和期貨單位是“張”精確到小數點後 0 位，超出部分會被捨棄)
    price|float|訂單改單後的價格  (證券帳戶精確到小數點後 3 位，超出部分會被捨棄期貨帳戶精確到小數點後 9 位，超出部分會被捨棄)
    adjust_limit|float|價格微調幅度  (OpenD 會對傳入價格自動調整到合法價位上（期貨忽略此參數）
  - 正數代表向上調整，負數代表向下調整
  - 例如：0.015 代表向上調整且幅度不超過 1.5%；-0.01 代表向下調整且幅度不超過 1%。預設 0 表示不調整)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務帳戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)
    aux_price|float|觸發價格  (- 當訂單是止損市價單、止損限價單、觸及限價單（止盈）、觸及市價單（止盈） 時，aux_price 為必傳參數
  - 證券帳戶精確到小數點後 3 位，期貨帳戶精確到小數點後 9 位，超過部分四捨五入)
    trail_type|[TrailType](./trade.md#589)|跟蹤類型  (當訂單是跟蹤止損市價單、跟蹤止損限價單時，trail_type 為必傳參數)
    trail_value|float|跟蹤金額/百分比  (- 當訂單是跟蹤止損市價單、跟蹤止損限價單時，trail_value 為必傳參數
  - 當跟蹤類型為比例時，該欄位為百分比欄位，傳入 20 實際對應 20%
  - 當跟蹤類型為金額時，證券帳戶精確到小數點後 3 位，期貨帳戶精確到小數點後 9 位，超過部分四捨五入
  - 當跟蹤類型為比例時，精確到小數點後 2 位，超過部分四捨五入)
    trail_spread|float|指定價差  (- 當訂單是跟蹤止損限價單時，trail_spread 為必傳參數
  - 證券帳戶精確到小數點後 3 位，期貨帳戶精確到小數點後 9 位，超過部分四捨五入)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回改單資訊</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 改單資訊格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_env|[TrdEnv](./trade.md#9175)|交易環境
        order_id|str|訂單號

* **Example**

```python
from futu import *
pwd_unlock = '123456'
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.unlock_trade(pwd_unlock)  # 若使用真實帳戶改單/撤單，需先對帳戶進行解鎖。此處示例為模擬帳戶撤單，也可省略解鎖。
if ret == RET_OK:
    order_id = "8851102695472794941"
    ret, data = trd_ctx.modify_order(ModifyOrderOp.CANCEL, order_id, 0, 0)
    if ret == RET_OK:
        print(data)
        print(data['order_id'][0])  # 獲取改單的訂單號
        print(data['order_id'].values.tolist())  # 轉為 list
    else:
        print('modify_order error: ', data)
else:
    print('unlock_trade failed: ', data)
trd_ctx.close()
```

* **Output**

```python
    trd_env             order_id
0    REAL      8851102695472794941
8851102695472794941
['8851102695472794941']
```


`cancel_all_order(trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, trdmarket=TrdMarket.NONE)`

* **介紹**

    撤消全部訂單。模擬交易以及 A 股通帳戶暫不支援全部撤單。

* **參數**
    參數|類型|說明
    :-|:-|:-
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務帳戶 ID  (當 acc_id 傳 0 時， 以 acc_index 指定的帳戶為準當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)
    trdmarket|[TrdMarket](./trade.html#1256)|指定交易市場  (撤銷指定帳戶指定市場的訂單預設狀態時，撤銷指定帳戶全部市場的訂單)


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td>str</td>
            <td>介面執行結果。ret == RET_OK 代表介面執行正常，ret != RET_OK 代表介面執行失敗</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td rowspan="2">str</td>
            <td>當 ret == RET_OK，返回"success"</td>
        </tr>
        <tr>
            <td>當 ret != RET_OK，返回錯誤描述</td>
        </tr>
    </table>

    * 全部撤單資訊格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_env|[TrdEnv](./trade.md#9175)|交易環境
        order_id|str|訂單號

* **Example**

```python
from futu import *
pwd_unlock = '123456'
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.unlock_trade(pwd_unlock)  # 若使用真實帳戶改單/撤單，需先對帳戶進行解鎖。此處示例為模擬帳戶全部撤單，也可省略解鎖。
if ret == RET_OK:
    ret, data = trd_ctx.cancel_all_order()
    if ret == RET_OK:
        print(data)
    else:
        print('cancel_all_order error: ', data)
else:
    print('unlock_trade failed: ', data)
trd_ctx.close()
```

* **Output**

```python
success
```

---



---

# 查詢未完成訂單

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`order_list_query(order_id="", order_market=TrdMarket.NONE, status_filter_list=[], code='', start='', end='', trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, refresh_cache=False)`

* **介紹**

    查詢指定交易業務賬户的未完成訂單列表

* **參數**
    參數|類型|說明
    :-|:-|:-
    order_id|str|訂單號過濾  (- 回傳指定訂單號的數據
  - 預設狀態時，回傳所有數據)
    order_market|[TrdMarket](./trade.md#9498)|訂單標的所屬市場過濾 (- 訂單標的市場過濾，會回傳該市場下的標的訂單
  - 預設值為NONE，會回傳賬户下所有市場的訂單數據)
    status_filter_list|list|訂單狀態過濾  (- 回傳指定狀態的訂單數據
  - 預設狀態時，回傳所有數據
  - list 中元素類型是 [OrderStatus](./trade.md#3177))
    code|str|代碼過濾  (- 回傳指定代碼的數據
  - 預設狀態時，回傳所有數據)
    start|str|開始時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 配置](../quick/opend-base.md#3795))
    end|str|結束時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 配置](../quick/opend-base.md#3795))
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務賬户 ID  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的賬户為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的賬户為準)
    acc_index|int|交易業務賬户列表中的賬户序號  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷賬户時發生變動，導致您指定的賬户與實際交易賬户不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務賬户)
    refresh_cache|bool|是否更新快取  (- True：立即向富途伺服器重新請求數據，不使用 OpenD 的快取，此時會受到介面限頻的限制
  - False：使用 OpenD 的快取（特殊情況導致快取沒有及時更新才需要更新）)
    


* **回傳**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面調用結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，回傳訂單列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，回傳錯誤描述</td>
        </tr>
    </table>

    * 訂單列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        order_type|[OrderType](./trade.md#379)|訂單類型
        order_status|[OrderStatus](./trade.md#3177)|訂單狀態
        order_id|str|訂單號
        code|str|股票編號
        stock_name|str|股票名稱
        order_market|[TrdMarket](./trade.md#9498)|訂單標的所屬市場
        qty|float|訂單數量  (期權期貨單位是"張")
        price|float|訂單價格  (精確到小數點後 3 位，超出部分四捨五入)
        currency|[Currency](./trade.md#3345)|交易貨幣
        create_time|str|創建時間  (期貨時區指定，請參見 [OpenD 配置](../quick/opend-base.md#3795))
        updated_time|str|最後更新時間  (期貨時區指定，請參見 [OpenD 配置](../quick/opend-base.md#3795))
        dealt_qty|float|成交數量  (期權期貨單位是"張")
        dealt_avg_price|float|成交均價  (無精度限制)
        last_err_msg|str|最後的錯誤描述  (如果有錯誤，會回傳最後一次錯誤的原因如果無錯誤，回傳空字符串)
        remark|str|下單時備註的標識  (詳見 [place_order](./place-order.md) 介面參數中的 remark)
        time_in_force|[TimeInForce](./trade.md#114)|有效期限
        fill_outside_rth|bool|是否允許盤前盤後（用於港股盤前競價與美股盤前盤後）  (True：允許False：不允許)
        session|[Session](../quote/quote.md#8417)|交易訂單時段（僅用於美股）
        aux_price|float|觸發價格
        trail_type|[TrailType](./trade.md#589)|跟蹤類型
        trail_value|float|跟蹤金額/百分比
        trail_spread|float|指定價差
        jp_acc_type|[SubAccType](./trade.md#5752)|日本賬户類型  (僅對日本券商生效)

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.order_list_query()
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果訂單列表不為空
        print(data['order_id'][0])  # 獲取未完成訂單的第一個訂單號
        print(data['order_id'].values.tolist())  # 轉為 list
else:
    print('order_list_query error: ', data)
trd_ctx.close()
```

* **Output**

```python
        code stock_name  order_market   trd_side           order_type   order_status             order_id    qty  price              create_time             updated_time  dealt_qty  dealt_avg_price last_err_msg      remark time_in_force fill_outside_rth session aux_price trail_type trail_value trail_spread currency jp_acc_type
0   HK.00700        HK         BUY           NORMAL  CANCELLED_ALL  6644468615272262086  100.0  520.0  2021-09-06 10:17:52.465  2021-09-07 16:10:22.806        0.0              0.0               asdfg+=@@@           GTC        N/A      N/A       560        N/A         N/A          N/A      HKD        N/A
6644468615272262086
['6644468615272262086']
```

---



---

# 查詢歷史訂單

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`history_order_list_query(status_filter_list=[], code='', order_market=TrdMarket.NONE, start='', end='', trd_env=TrdEnv.REAL, acc_id=0, acc_index=0)`

* **介紹**

    查詢指定交易業務帳戶的歷史訂單列表

* **參數**
    參數|類型|說明
    :-|:-|:-
    status_filter_list|list|訂單狀態過濾  (- 返回指定狀態的訂單數據
  - 預設狀態時，返回所有數據
  - list 中元素類型是 [OrderStatus](./trade.md#5161))
    code|str|代碼過濾  (- 返回指定代碼的數據
  - 預設狀態時，返回所有數據)
    order_market|[TrdMarket](./trade.md#9498)|訂單標的所屬市場過濾 (- 訂單標的市場過濾，會返回該市場下的標的訂單
  - 預設值為NONE，會返回帳戶下所有市場的訂單數據)
    start|str|開始時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
    end|str|結束時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務帳戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)

    * start 和 end 的組合如下
        Start 類型|End 類型|說明
        :-|:-|:-
        str|str|start 和 end 分別為指定的日期
        None|str|start 為 end 往前 90 天
        str|None|end 為 start 往後 90 天
        None|None|start 為往前 90 天，end 當前日期

* **返回**
    
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回訂單列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 訂單列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        order_type|[OrderType](./trade.md#379)|訂單類型
        order_status|[OrderStatus](./trade.md#5161)|訂單狀態
        order_id|str|訂單號
        code|str|股票編號
        stock_name|str|股票名稱
        order_market|[TrdMarket](./trade.md#9498)|訂單標的所屬市場
        qty|float|訂單數量  (期權期貨單位是"張")
        price|float|訂單價格  (精確到小數點後 3 位，超出部分四捨五入)
        currency|[Currency](./trade.md#3345)|交易貨幣
        create_time|str|建立時間  (期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
        updated_time|str|最後更新時間  (期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
        dealt_qty|float|成交數量  (期權期貨單位是"張")
        dealt_avg_price|float|成交均價  (無精度限制)
        last_err_msg|str|最後的錯誤描述  (如果有錯誤，會返回最後一次錯誤的原因如果無錯誤，返回空字符串)
        remark|str|下單時備註的標記  (詳見 [place_order](./place-order.md) 接口參數中的 remark)
        time_in_force|[TimeInForce](./trade.md#114)|有效期限
        fill_outside_rth|bool|是否允許盤前盤後（用於港股盤前競價與美股盤前盤後）  (True：允許False：不允許)
        session|[Session](../quote/quote.md#8417)|交易訂單時段（僅用於美股）
        aux_price|float|觸發價格
        trail_type|[TrailType](./trade.md#589)|跟蹤類型
        trail_value|float|跟蹤金額/百分比
        trail_spread|float|指定價差
        jp_acc_type|[SubAccType](./trade.md#5752)|日本帳戶類型  (僅對日本券商生效)

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.US, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUINC)
ret, data = trd_ctx.history_order_list_query()
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果訂單列表不為空
        print(data['order_id'][0])  # 獲取持倉第一個訂單號
        print(data['order_id'].values.tolist())  # 轉為 list
else:
    print('history_order_list_query error: ', data)
trd_ctx.close()
```

* **Output**

```python
        code stock_name order_market    trd_side           order_type   order_status             order_id    qty  price              create_time             updated_time  dealt_qty  dealt_avg_price last_err_msg      remark time_in_force fill_outside_rth session aux_price trail_type trail_value trail_spread currency jp_acc_type
0   US.AAPL        US          BUY           NORMAL  CANCELLED_ALL  6644468615272262086  100.0  520.0  2021-09-06 10:17:52.465  2021-09-07 16:10:22.806        0.0              0.0               asdfg+=@@@           GTC      N/A        N/A       560        N/A         N/A          N/A      USD        N/A
6644468615272262086
['6644468615272262086']
```

---



---

# 回應訂單推送回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    回應訂單推送，非同步處理 OpenD 推送過來的訂單狀態資訊。  
    在收到 OpenD 推送過來的訂單狀態資訊後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。

* **參數**
    
    參數|類型|説明
    :-|:-|:-
    rsp_pb|Trd_UpdateOrder_pb2.Response|衍生類別中不需要直接處理該參數

* **傳回**
    
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，傳回訂單列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，傳回錯誤描述</td>
        </tr>
    </table>

    * 訂單列表格式如下：
        欄位|類型|説明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        order_type|[OrderType](./trade.md#379)|訂單類型
        order_status|[OrderStatus](./trade.md#3177)|訂單狀態
        order_id|str|訂單號
        code|str|股票編號
        stock_name|str|股票名稱
        qty|float|訂單數量  (期權期貨單位是"張")
        price|float|訂單價格  (精確到小數點後 3 位，超出部分四捨五入)
        currency|[Currency](./trade.md#3345)|交易貨幣
        create_time|str|建立時間  (期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
        updated_time|str|最後更新時間  (期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
        dealt_qty|float|成交數量  (期權期貨單位是"張")
        dealt_avg_price|float|成交均價  (無精度限制)
        last_err_msg|str|最後的錯誤描述  (如果有錯誤，會傳回最後一次錯誤的原因如果無錯誤，傳回空字串)
        remark|str|下單時備註的標識  (詳見 [place_order](./place-order.md) 介面參數中的 remark)
        time_in_force|[TimeInForce](./trade.md#114)|有效期限
        fill_outside_rth|bool|是否允許盤前盤後（僅用於美股）  (True：允許False：不允許)
        session|[Session](../quote/quote.md#8417)|交易訂單時段（僅用於美股）
        aux_price|float|觸發價格
        trail_type|[TrailType](./trade.md#589)|跟蹤類型
        trail_value|float|跟蹤金額/百分比
        trail_spread|float|指定價差

* **Example**

```python
from futu import *
from time import sleep
class TradeOrderTest(TradeOrderHandlerBase):
    """ order update push"""
    def on_recv_rsp(self, rsp_pb):
        ret, content = super(TradeOrderTest, self).on_recv_rsp(rsp_pb)
        if ret == RET_OK:
            print("* TradeOrderTest content={}\n".format(content))
        return ret, content

trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
trd_ctx.set_handler(TradeOrderTest())
print(trd_ctx.place_order(price=518.0, qty=100, code="HK.00700", trd_side=TrdSide.SELL))

sleep(15)
trd_ctx.close()
```

* **Output**

```python
* TradeOrderTest content=  trd_env      code stock_name  dealt_avg_price  dealt_qty    qty           order_id order_type  price order_status          create_time         updated_time trd_side last_err_msg trd_market remark time_in_force fill_outside_rth session aux_price trail_type trail_value trail_spread currency
0    REAL  HK.00700       騰訊控股              0.0        0.0  100.0  72625263708670783     NORMAL  518.0   SUBMITTING  2021-11-04 11:26:27  2021-11-04 11:26:27      BUY                      HK                  DAY      N/A        N/A       N/A        N/A         N/A          N/A      HKD
```

---



---

# 查詢訂單費用

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`order_fee_query(order_id_list=[], acc_id=0, acc_index=0, trd_env=TrdEnv.REAL)`

* **介紹**

    查詢指定訂單的收費明細（最低版本要求：8.2.4218）

* **參數**
    參數|類型|說明
    :-|:-|:-
    order_id_list|list|訂單號列表 (- 每次請求最多查詢 400 筆訂單
  - list 內元素類型為 str)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境
    acc_id|int|交易業務賬戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務賬戶，二選一即可，建議使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的賬戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的賬戶為準)
    acc_index|int|交易業務賬戶列表中的賬戶序號  (- acc_id 和 acc_index 都可用於指定交易業務賬戶，二選一即可，建議使用 acc_id。acc_index 會在新開立/關閉帳戶時發生變動，導致您指定的賬戶與實際交易賬戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務賬戶)
    


* **傳回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，傳回訂單費用列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，傳回錯誤描述</td>
        </tr>
    </table>

    * 訂單列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        order_id|str|訂單號
        fee_amount|float|總費用
        fee_details|list|收費明細 (- 格式：[('收費項1', 收費項1的金額), ('收費項2', 收費項2的金額), ('收費項3', 收費項3的金額)……]
  - 常見的收費項包括：佣金、平台使用費、期權監管費、期權清算費、期權交收費、交收費、證監會徵費、交易活動費)

        
* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.US, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret1, data1 = trd_ctx.history_order_list_query(status_filter_list=[OrderStatus.FILLED_ALL])
if ret1 == RET_OK:
    if data1.shape[0] > 0:  # 如果訂單列表不為空
        ret2, data2 = trd_ctx.order_fee_query(data1['order_id'].values.tolist())  # 將訂單 id 轉為 list，查詢訂單費用
        if ret2 == RET_OK:
            print(data2)
            print(data2['fee_details'][0])  # 打印第一筆訂單的收費明細
        else:
            print('order_fee_query error: ', data2)
else:
    print('order_list_query error: ', data1)
trd_ctx.close()
```

* **Output**

```python
                                            order_id  fee_amount                                        fee_details
0  v3_20240314_12345678_MTc4NzA5NzY5OTA3ODAzMzMwN       10.46  [(佣金, 5.85), (平台使用費, 2.7), (期權監管費, 0.11), (期權清...
1  v3_20240318_12345678_MTM5Nzc5MDYxNDY1NDM1MDI1M        2.25  [(佣金, 0.99), (平台使用費, 1.0), (交收費, 0.15), (證監會徵費...
[('佣金', 5.85), ('平台使用費', 2.7), ('期權監管費', 0.11), ('期權清算費', 0.18), ('期權交收費', 1.62)]
```

---



---

# 訂閱交易推送

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>
    Python 不需要訂閱交易推送

---



---

# 查詢當日成交

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`deal_list_query(code="", deal_market= TrdMarket.NONE, trd_env=TrdEnv.REAL, acc_id=0, acc_index=0, refresh_cache=False)`

* **介紹**
    
	查詢指定交易業務賬户的當日成交列表。  
    該介面只支援實盤交易，不支援模擬交易。

* **參數**
    參數|類型|說明
    :-|:-|:-
    code|str|代碼過濾  (只返回此代碼對應的成交數據不傳則返回所有)
    deal_market|[TrdMarket](./trade.md#1253)|成交標的所屬市場過濾  (- 成交標的市場過濾，會返回該市場下的成交數據
  - 預設值為NONE，會返回賬户下所有市場的成交數據)
    trd_env|[TrdEnv](./trade.md#9175)|交易環境  (僅支援 TrdEnv.REAL（真實環境），模擬環境暫不支援查詢成交數據)
    acc_id|int|交易業務賬户 ID  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的賬户為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的賬户為準)
    acc_index|int|交易業務賬户列表中的賬户序號  (- acc_id 和 acc_index 都可用於指定交易業務賬户，二選一即可，推薦使用 acc_id。acc_index 會在新開立/註銷賬户時發生變動，導致您指定的賬户與實際交易賬户不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務賬户)
    refresh_cache|bool|是否更新快取  (- True：立即向富途伺服器重新請求數據，不使用 OpenD 的緩存，此時會受到介面限頻的限制
  - False：使用 OpenD 的緩存（特殊情況導致緩存沒有及時更新才需要刷新）)
    


* **返回**

    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面調用結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回交易成交列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 交易成交列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        deal_id|str|成交號
        order_id|str|訂單號
        code|str|股票代碼
        stock_name|str|股票名稱
        deal_market|[TrdMarket](./trade.md#1253)|成交標的所屬市場
        qty|float|成交數量  (期權期貨單位是"張")
        price|float|成交價格  (精確到小數點後 3 位，超出部分四捨五入)
        create_time|str|創建時間  (期貨時區指定，請參見 [OpenD 配置](../quick/opend-base.md#3795))
        counter_broker_id|int|對手經紀號  (僅港股有效)
        counter_broker_name|str|對手經紀名稱  (僅港股有效)
        status|[DealStatus](./trade.md#7204)|成交狀態
        jp_acc_type|[SubAccType](./trade.md#5752)|日本賬户類型  (僅對日本券商生效)

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.deal_list_query()
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果成交列表不為空
        print(data['order_id'][0])  # 獲取當日成交的第一個訂單號
        print(data['order_id'].values.tolist())  # 轉為 list
else:
    print('deal_list_query error: ', data)
trd_ctx.close()
```

* **Output**

```python
    code stock_name     deal_market         deal_id             order_id        qty  price    trd_side     create_time      counter_broker_id counter_broker_name status jp_acc_type
0  HK.00388      香港交易所     HK    5056208452274069375  4665291631090960915  100.0  370.0      BUY  2020-09-17 21:15:59.979           5         OK       N/A
4665291631090960915
['4665291631090960915']
```

---



---

# 查詢歷史成交

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">
<template v-slot:py>


`history_deal_list_query(code='', deal_market=TrdMarket.NONE, start='', end='', trd_env=TrdEnv.REAL, acc_id=0, acc_index=0)`

* **介紹**

    查詢指定交易業務帳戶的歷史成交列表。  
    該介面只支援實盤交易，不支援模擬交易。

* **參數**

    參數|類型|說明
    :-|:-|:-
    code|str|代碼過濾  (只回傳此代碼對應的成交數據不傳則回傳所有)
    deal_market|[TrdMarket](./trade.md#1253)|成交標的所屬市場過濾  (- 成交標的市場過濾，會回傳該市場下的成交數據
  - 預設值為NONE，會回傳帳戶下所有市場的成交數據)
    start|str|開始時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
    end|str|結束時間  (- 嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
  - 期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
    trd_env|[TrdEnv](./trade.md#9175)|交易環境  (僅支援 TrdEnv.REAL（真實環境），模擬環境暫不支援查詢成交數據)
    acc_id|int|交易業務帳戶 ID  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，建議使用 acc_id。
  - 當 acc_id 傳 0 時， 以 acc_index 指定的帳戶為準
  - 當 acc_id 傳 ID 號時（不為 0 ），以 acc_id 指定的帳戶為準)
    acc_index|int|交易業務帳戶列表中的帳戶序號  (- acc_id 和 acc_index 都可用於指定交易業務帳戶，二選一即可，建議使用 acc_id。acc_index 會在新開立/註銷帳戶時發生變動，導致您指定的帳戶與實際交易帳戶不一致。
  - acc_index 預設為 0，表示指定第 1 個交易業務帳戶)
    
    * start 和 end 的組合如下
        Start 類型|End 類型|說明
        :-|:-|:-
        str|str|start 和 end 分別為指定的日期
        None|str|start 為 end 往前 90 天
        str|None|end 為 start 往後 90 天
        None|None|start 為往前 90 天，end 當前日期

* **回傳**
    
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>說明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面調用結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，回傳交易成交列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，回傳錯誤描述</td>
        </tr>
    </table>

    * 交易成交列表格式如下：
        欄位|類型|說明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        deal_id|str|成交號
        order_id|str|訂單號
        code|str|股票編號
        stock_name|str|股票名稱
        deal_market|[TrdMarket](./trade.md#1253)|成交標的所屬市場
        qty|float|成交數量  (期權期貨單位是"張")
        price|float|成交價格  (精確到小數點後 3 位，超過部分四捨五入)
        create_time|str|建立時間  (期貨時區指定，請參見 [OpenD 設定](../quick/opend-base.md#3795))
        counter_broker_id|int|對手經紀號  (僅港股有效)
        counter_broker_name|str|對手經紀名稱  (僅港股有效)
        status|[DealStatus](./trade.md#7204)|成交狀態
        jp_acc_type|[SubAccType](./trade.md#5752)|日本帳戶類型  (僅對日本券商生效)

* **Example**

```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.history_deal_list_query()
if ret == RET_OK:
    print(data)
    if data.shape[0] > 0:  # 如果成交列表不為空
        print(data['deal_id'][0])  # 獲取歷史成交的第一個成交號
        print(data['deal_id'].values.tolist())  # 轉為 list
else:
    print('history_deal_list_query error: ', data)
trd_ctx.close()
```

* **Output**

```python
    code stock_name     deal_market         deal_id             order_id    qty  price trd_side              create_time  counter_broker_id counter_broker_name status jp_acc_type
0  HK.00388      香港交易所    HK  5056208452274069375  4665291631090960915  100.0  370.0      BUY  2020-09-17 21:15:59.979                  5                         OK        N/A
5056208452274069375
['5056208452274069375']
```

---



---

# 回應成交推送回呼

<FtSwitcher :languages="{py:'Python', pb:'Proto', cs:'C#', java:'Java', cpp:'C++', js:'JavaScript'}">

<template v-slot:py>


`on_recv_rsp(self, rsp_pb)`

* **介紹**

    回應成交推送，非同步處理 OpenD 推送過來的成交狀態資訊。  
    在收到 OpenD 推送過來的成交狀態資訊後會回呼到該函數，您需要在衍生類別中覆寫 on_recv_rsp。  
    該接口只支援實盤交易，不支援模擬交易。
 
* **參數**
    
    參數|類型|説明
    :-|:-|:-
    rsp_pb|Trd_UpdateOrderFill_pb2.Response|衍生類別中不需要直接處理該參數

* **返回**
    
    <table>
        <tr>
            <th>參數</th>
            <th>類型</th>
            <th>説明</th>
        </tr>
        <tr>
            <td>ret</td>
            <td><a href="../ftapi/common.html#3835"> RET_CODE</a></td>
            <td>介面執行結果</td>
        </tr>
        <tr>
            <td rowspan="2">data</td>
            <td>pd.DataFrame</td>
            <td>當 ret == RET_OK 時，返回交易成交列表</td>
        </tr>
        <tr>
            <td>str</td>
            <td>當 ret != RET_OK 時，返回錯誤描述</td>
        </tr>
    </table>

    * 交易成交列表格式如下：
        欄位|類型|説明
        :-|:-|:-
        trd_side|[TrdSide](./trade.md#5815)|交易方向
        deal_id|str|成交號
        order_id|str|訂單號
        code|str|股票編號
        stock_name|str|股票名稱
        qty|float|成交數量  (期權期貨單位是"張")
        price|float|成交價格
        create_time|str|建立時間  (期貨時區指定，請參見 [FutuOpenD 配置](../quick/opend-base.md#3795))
        counter_broker_id|int|對手經紀號  (僅港股有效)
        counter_broker_name|str|對手經紀名稱  (僅港股有效)
        status|[DealStatus](./trade.md#7204)|成交狀態

* **Example**

```python
from futu import *
from time import sleep
class TradeDealTest(TradeDealHandlerBase):
    """ order update push"""
    def on_recv_rsp(self, rsp_pb):
        ret, content = super(TradeDealTest, self).on_recv_rsp(rsp_pb)
        if ret == RET_OK:
            print("TradeDealTest content={}".format(content))
        return ret, content

trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
trd_ctx.set_handler(TradeDealTest())
print(trd_ctx.place_order(price=595.0, qty=100, code="HK.00700", trd_side=TrdSide.BUY))

sleep(15)
trd_ctx.close()
```

* **Output**

```python
TradeDealTest content=  trd_env      code stock_name              deal_id             order_id    qty  price trd_side              create_time  counter_broker_id counter_broker_name trd_market status
0    REAL  HK.00700       騰訊控股  2511067564122483295  8561504228375901919  100.0  518.0      BUY  2021-11-04 11:29:41.595                  5                   5         HK     OK
```

---



---

# 交易定義

## 賬戶風控狀態

**CltRiskLevel**

```protobuf
enum CltRiskLevel
{
    CltRiskLevel_Unknown = -1;        // 未知
    CltRiskLevel_Safe = 0;          // 安全
    CltRiskLevel_Warning = 1;       // 預警
    CltRiskLevel_Danger = 2;        // 危險
    CltRiskLevel_AbsoluteSafe = 3;  // 絕對安全
    CltRiskLevel_OptDanger = 4;     // 危險（期權相關）
}
```

## 貨幣類型

**Currency**

```protobuf
enum Currency
{
    Currency_Unknown = 0;  //未知貨幣
    Currency_HKD = 1;   // 港元
    Currency_USD = 2;   // 美元
    Currency_CNH = 3;   // 離岸人民幣
    Currency_JPY = 4;   // 日元
    Currency_SGD = 5;   // 新元
	  Currency_AUD = 6;   // 澳元
    Currency_CAD = 7; // 加拿大元
    Currency_MYR = 8; // 馬幣
}
```

## 追蹤類型

```protobuf
enum TrailType
{
	TrailType_Unknown = 0; //未知類型
	TrailType_Ratio = 1; //比例
	TrailType_Amount = 2; //金額
}
```

## 修改訂單操作

**ModifyOrderOp**

```protobuf
enum ModifyOrderOp
{
    //港股支援全部操作，美股目前僅支援 ModifyOrderOp_Normal 和 ModifyOrderOp_Cancel
    ModifyOrderOp_Unknown = 0; //未知操作
    ModifyOrderOp_Normal = 1; //修改訂單的價格、數量，即以前的改單
    ModifyOrderOp_Cancel = 2; //撤單。未成交訂單將直接從交易所撮合隊列中撤銷。
    ModifyOrderOp_Disable = 3; //使失效。對交易所來説，「失效」的效果等同於 「撤單」。訂單「失效」後，未成交訂單將直接從交易所撮合隊列中撤出，但訂單資訊（如價格和數量）會繼續保留在富途伺服器，您隨時可以重新使它生效。
    ModifyOrderOp_Enable = 4; //使生效。對交易所來説，「生效」等同於下一筆新訂單。訂單重新「生效」後，將按照原來的價格數量重新提交到交易所，並按照價格優先、時間優先順序重新排隊。
    ModifyOrderOp_Delete = 5; //刪除。指對已撤單/下單失敗的訂單進行隱藏操作。
}
```

## 成交狀態

**OrderFillStatus**

```protobuf
enum OrderFillStatus
{
    OrderFillStatus_OK = 0; //正常
    OrderFillStatus_Cancelled = 1; //成交被取消
    OrderFillStatus_Changed = 2; //成交被更改
}
```

## 訂單狀態

**OrderStatus**

```protobuf
enum OrderStatus
{
    OrderStatus_Unknown = -1; //未知狀態
    OrderStatus_WaitingSubmit = 1; //待提交
    OrderStatus_Submitting = 2; //提交中
    OrderStatus_Submitted = 5; //已提交，等待成交
    OrderStatus_Filled_Part = 10; //部分成交
    OrderStatus_Filled_All = 11; //全部已成
    OrderStatus_Cancelled_Part = 14; //部分成交，剩餘部分已撤單
    OrderStatus_Cancelled_All = 15; //全部已撤單，無成交
    OrderStatus_Failed = 21; //下單失敗，服務拒絕
    OrderStatus_Disabled = 22; //已失效
    OrderStatus_Deleted = 23; //已刪除，無成交的訂單才能刪除
    OrderStatus_FillCancelled = 24; //成交被撤銷（一般遇不到，意思是已經成交的訂單被回滾撤銷，成交無效變為廢單）
};
```

## 訂單類型

:::tip 提示
* [實盤交易中，各個品類支援的訂單類型](../qa/trade.md#9863)
* 模擬交易中，僅支援限價單(NORMAL)和市價單(MARKET)。
:::

**OrderType**

```protobuf
enum OrderType
{
    OrderType_Unknown = 0; //未知類型
    OrderType_Normal = 1; //限價單
    OrderType_Market = 2; //市價單
    OrderType_AbsoluteLimit = 5; //絕對限價訂單（僅港股），只有價格完全匹配才成交，否則下單失敗。舉例：下一筆價格為 5 元的絕對限價買單，賣方的價格必須也是5元才能成交，賣方即使低於 5 元也不能成交，下單失敗。賣出同理
    OrderType_Auction = 6; //競價訂單（僅港股），僅港股早盤競價和收盤競價有效
    OrderType_AuctionLimit = 7; //競價限價訂單（僅港股），僅早盤競價和收盤競價有效，參與競價，且要求滿足指定價格才會成交
    OrderType_SpecialLimit = 8; //特別限價訂單（僅港股），成交規則同增強限價訂單，且部分成交後，交易所自動撤銷訂單
    OrderType_SpecialLimit_All = 9; //特別限價且要求全部成交訂單（僅港股）。全部成交，否則自動撤單
    OrderType_Stop = 10; // 止損市價單
    OrderType_StopLimit = 11; // 止損限價單
    OrderType_MarketifTouched = 12; // 觸及市價單（止盈）
    OrderType_LimitifTouched = 13; // 觸及限價單（止盈）
    OrderType_TrailingStop = 14; // 追蹤止損市價單
    OrderType_TrailingStopLimit = 15; // 追蹤止損限價單
    OrderType_TWAP  = 16; // 時間加權市價算法單（僅美股）
    OrderType_TWAP_LIMIT = 17; // 時間加權限價算法單 （港股和美股）
    OrderType_VWAP  = 18; // 成交量加權市價算法單（僅美股）
    OrderType_VWAP_LIMIT  = 19; // 成交量加權限價算法單（港股和美股）
}
```

## 持倉方向

**PositionSide**

```protobuf
enum PositionSide
{
    PositionSide_Long = 0; //長倉，預設情況是長倉
    PositionSide_Unknown = -1; //未知方向
    PositionSide_Short = 1; //短倉
};
```

## 賬戶類型

**TrdAccType**

```protobuf
enum TrdAccType
{
    TrdAccType_Unknown = 0; //未知類型
    TrdAccType_Cash = 1;    //現金賬戶
    TrdAccType_Margin = 2;  //保證金賬戶
    TrdAccType_TFSA = 3;    //加拿大免税賬戶
    TrdAccType_RRSP = 4;    //加拿大註冊退休賬戶
    TrdAccType_SRRSP = 5;    //加拿大配偶退休賬戶
    TrdAccType_Derivatives = 6;    //日本衍生品賬戶
};
```

## 交易環境

**TrdEnv**

```protobuf
enum TrdEnv
{
    TrdEnv_Simulate = 0; //模擬環境
    TrdEnv_Real = 1; //真實環境
}
```

## 交易市場

**TrdMarket**

```protobuf
enum TrdMarket
{
    TrdMarket_Unknown = 0; //未知市場
    TrdMarket_HK = 1; //香港市場（證券、期權）
    TrdMarket_US = 2; //美國市場（證券、期權）
    TrdMarket_CN = 3; //A 股市場（僅用於模擬交易）
    TrdMarket_HKCC = 4; //A 股通市場（股票）
    TrdMarket_Futures = 5; //期貨市場（環球期貨）
    TrdMarket_SG = 6; //新加坡市場
    TrdMarket_AU = 8; //澳洲市場
    TrdMarket_Futures_Simulate_HK = 10; //香港期貨模擬市場
    TrdMarket_Futures_Simulate_US = 11; //美國期貨模擬市場
    TrdMarket_Futures_Simulate_SG = 12; //新加坡期貨模擬市場
    TrdMarket_Futures_Simulate_JP = 13; //日本期貨模擬市場
    TrdMarket_JP = 15; //日本市場
    TrdMarket_MY = 111; //馬來西亞市場
    TrdMarket_CA = 112; //加拿大市場
    TrdMarket_HK_Fund = 113; //香港基金市場
    TrdMarket_US_Fund = 123; //美國基金市場
}
```


## 賬戶狀態

**TrdAccStatus**

```protobuf
enum TrdAccStatus
{
    TrdAccStatus_Active = 0; //生效賬戶
    TrdAccStatus_Disabled = 1; //失效賬戶
}
```


## 賬戶結構

**TrdAccRole**

```protobuf
enum TrdAccRole
{
    TrdAccRole_Unknown = 0; //未知
    TrdAccRole_Normal = 1; //普通賬戶
    TrdAccRole_Master = 2; //主賬戶
    TrdAccRole_IPO = 3; //馬來西亞IPO賬戶
}
```


## 交易證券市場

**TrdSecMarket**

```protobuf
enum TrdSecMarket
{
    TrdSecMarket_Unknown = 0; //未知市場
    TrdSecMarket_HK = 1; //香港市場（股票、窩輪、牛熊、期權、期貨等）
    TrdSecMarket_US = 2; //美國市場（股票、期權、期貨等）
    TrdSecMarket_CN_SH = 31; //滬股市場（股票）
    TrdSecMarket_CN_SZ = 32; //深股市場（股票）
    TrdSecMarket_SG = 41;  //新加坡市場（期貨）  
    TrdSecMarket_JP = 51;  //日本市場（期貨）  
    TrdSecMarket_AU = 61; // 澳大利亞  
    TrdSecMarket_MY = 71; // 馬來西亞  
    TrdSecMarket_CA = 81; // 加拿大  
    TrdSecMarket_FX = 91; // 外匯  
}
```

## 交易方向

**TrdSide**

```protobuf
enum TrdSide
{
    //客戶端下單隻傳 Buy 或 Sell 即可，SellShort 是美股訂單時伺服器返回有此方向，BuyBack 目前不存在，但也不排除伺服器會傳
    TrdSide_Unknown = 0; //未知方向
    TrdSide_Buy = 1; //買入
    TrdSide_Sell = 2; //賣出
    TrdSide_SellShort = 3; //賣空
    TrdSide_BuyBack = 4; //買回
}
```

:::tip 提示
**下單** 接口的交易方向 ，建議僅使用 `買入` 和 `賣出` 兩個方向作為入參。  
`賣空` 和 `買回` 僅適用於日本券商，其他券商僅用於 **查詢今日訂單** ，**查詢歷史訂單** ，**響應訂單推送回調** ，**查詢當日成交** ，**查詢歷史成交** ，**響應成交推送回調** 接口的返回欄位展示。
:::

## 訂單有效期

**TimeInForce**

```protobuf
enum TimeInForce
{
    TimeInForce_DAY = 0;       // 當日有效
    TimeInForce_GTC = 1;       // 撤單前有效，最多持續90自然日。
}
```

## 賬戶所屬券商

**SecurityFirm**

```protobuf
enum SecurityFirm
{
    SecurityFirm_Unknown = 0;        //未知
    SecurityFirm_FutuSecurities = 1; //富途證券（香港）
    SecurityFirm_FutuInc = 2;        //moomoo證券(美國)
    SecurityFirm_FutuSG = 3;        //moomoo證券(新加坡)
    SecurityFirm_FutuAU = 4;         //moomoo證券(澳大利亞)
    SecurityFirm_FutuCA = 5;         //富途證券（加拿大）
    SecurityFirm_FutuMY = 6;         //富途證券（馬來西亞）
    SecurityFirm_FutuJP = 7;         //富途證券（日本）
}
```

## 模擬交易賬戶類型

**SimAccType**

```protobuf
enum SimAccType
{
    SimAccType_Unknown = 0;		//未知
    SimAccType_Stock = 1;		//股票模擬賬戶（僅用於交易證券類產品，不支援交易期權）
    SimAccType_Option = 2;      //期權模擬賬戶（僅用於交易期權，不支援交易股票證券類產品）
    SimAccType_Futures = 3;      //期貨模擬賬戶
    SimAccType_StockAndOption = 4;      //美股融資融券模擬帳戶

}
```

## 風險狀態

**CltRiskStatus**

```protobuf
enum CltRiskStatus
{
  CltRiskStatus_Level1 = 0;  //非常安全
  CltRiskStatus_Level2 = 1;  //安全
  CltRiskStatus_Level3 = 2;  //較安全
  CltRiskStatus_Level4 = 3;  //較低風險
  CltRiskStatus_Level5 = 4;  //中等風險
  CltRiskStatus_Level6 = 5;  //較高風險
  CltRiskStatus_Level7 = 6;  //預警
  CltRiskStatus_Level8 = 7;  //預警
  CltRiskStatus_Level9 = 8;  //預警
}
```

## 日內交易限制情況

**DTStatus**

```protobuf
enum DTStatus
{
	DTStatus_Unknown = 0; 		//未知
	DTStatus_Unlimited = 1;		//無限次(當前可以無限次日內交易，注意留意剩餘日內交易購買力)
	DTStatus_EMCall = 2;		//EM Call(當前狀態不能新建倉位，需要補充資產淨值至$25000以上，否則會被禁止新建倉位90天)
	DTStatus_DTCall = 3;		//DT Call(當前狀態有未補平的日內交易追繳金額（DTCall），需要在5個交易日內足額入金來補平 DTCall，否則會被禁止新建倉位，直到足額存入資金才會解禁)
}
```

## 現金流方向

**TrdCashFlowDirection**

```protobuf
enum TrdCashFlowDirection
{
	TrdCashFlowDirection_Unknown = 0; //未知
	TrdCashFlowDirection_In = 1; //現金流入
	TrdCashFlowDirection_Out = 2; //現金流出
}
```

## 日本子賬戶類型

**TrdSubAccType**

```protobuf
enum TrdSubAccType
{
	TrdSubAccType_None = 0; //未知
	TrdSubAccType_JP_GENERAL = 1; // 一般-Long
	TrdSubAccType_JP_TOKUTEI = 2; // 特定-Long
	TrdSubAccType_JP_NISA_GENERAL = 3; // 一般NISA
	TrdSubAccType_JP_NISA_TSUMITATE = 4; // 累計NISA

	TrdSubAccType_JP_GENERAL_SHORT = 5; // 一般-short
	TrdSubAccType_JP_TOKUTEI_SHORT = 6; // 特定-short
	TrdSubAccType_JP_HONPO_GENERAL = 7; // 本國信用交易抵押品-一般
	TrdSubAccType_JP_GAIKOKU_GENERAL = 8; // 外國信用交易抵押品-一般
	TrdSubAccType_JP_HONPO_TOKUTEI = 9; // 本國信用交易抵押品-特定
	TrdSubAccType_JP_GAIKOKU_TOKUTEI = 10; // 外國信用交易抵押品-特定

	TrdSubAccType_JP_DERIVATIVE_LONG = 11; // 衍生品子賬戶-Long
	TrdSubAccType_JP_DERIVATIVE_SHORT = 12; // 衍生品子賬戶-Short
	TrdSubAccType_JP_HONPO_DERIVATIVE_GENERAL = 13; // 本國衍生品證據金子賬戶-一般
	TrdSubAccType_JP_GAIKOKU_DERIVATIVE_GENERAL = 14; // 外國衍生品證據金子賬戶-一般
	TrdSubAccType_JP_HONPO_DERIVATIVE_TOKUTEI = 15; // 本國衍生品證據金子賬戶-特定
	TrdSubAccType_JP_GAIKOKU_DERIVATIVE_TOKUTEI = 16; // 外國衍生品證據金子賬戶-特定
}
```

## 資產類別

**TrdAssetCategory**

```protobuf
enum TrdAssetCategory
{
	TrdAssetCategory_Unknown = 0; 	//未知
	TrdAssetCategory_JP = 1;	    //本國
	TrdAssetCategory_US = 2;	    //外國
}
```

## 交易品類

**TrdCategory**

```protobuf
enum TrdCategory
{
    TrdCategory_Unknown = 0; //未知品類
    TrdCategory_Security = 1; //證券
    TrdCategory_Future = 2; //期貨
}
```

## 賬戶現金資訊

**AccCashInfo**

```protobuf
message AccCashInfo
{
    optional int32 currency = 1;        // 貨幣類型，取值參考 Currency
    optional double cash = 2;           // 現金結餘
    optional double availableBalance = 3;   // 現金可提金額
    optional double netCashPower = 4;		// 現金購買力
}
```

## 分市場資產資訊

**AccMarketInfo**

```protobuf
message AccCashInfo
{
    optional int32 trdMarket = 1;        // 交易市場, 參見TrdMarket的枚舉定義
    optional double assets = 2;          // 分市場資產資訊
}
```


## 交易協議公共參數頭

**TrdHeader**

```protobuf
message TrdHeader
{
  required int32 trdEnv = 1; //交易環境, 參見 TrdEnv 的枚舉定義
  required uint64 accID = 2; //業務賬號, 業務賬號與交易環境、市場權限需要匹配，否則會返回錯誤
  required int32 trdMarket = 3; //交易市場, 參見 TrdMarket 的枚舉定義
  optional int32 jpAccType = 4; //JP子賬戶類型，取值見 TrdSubAccType
}
```

## 交易業務賬戶

**TrdAcc**

```protobuf
message TrdAcc
{
  required int32 trdEnv = 1; //交易環境，參見 TrdEnv 的枚舉定義
  required uint64 accID = 2; //業務賬號
  repeated int32 trdMarketAuthList = 3; //業務賬戶支援的交易市場權限，即此賬戶能交易那些市場, 可擁有多個交易市場權限，目前僅單個，取值參見 TrdMarket 的枚舉定義
  optional int32 accType = 4;   //賬戶類型，取值見 TrdAccType
  optional string cardNum = 5;  //卡號
  optional int32 securityFirm = 6; //所屬券商，取值見SecurityFirm
  optional int32 simAccType = 7; //模擬交易賬號類型，取值見SimAccType
  optional string uniCardNum = 8;  //所屬綜合賬戶卡號
  optional int32 accStatus = 9; //賬號狀態，取值見TrdAccStatus
  optional int32 accRole = 10; //賬號分類，是不是主賬號，取值見TrdAccRole
  repeated int32 jpAccType = 11; //JP子賬戶類型，取值見 TrdSubAccType
}
```


## 賬戶資金

**Funds**

```protobuf
message Funds
{
  required double power = 1; //最大購買力（此欄位是按照 50% 的融資初始保證金率計算得到的 近似值。但事實上，每個標的的融資初始保證金率並不相同。我們建議您使用 查詢最大可買可賣 接口返回的 最大可買 欄位，來判斷實際可買入的最大數量）
  required double totalAssets = 2; //資產淨值
  required double cash = 3; //現金（僅單幣種賬戶使用此欄位，綜合賬戶請使用 cashInfoList 獲取分幣種現金）
  required double marketVal = 4; //證券市值, 僅證券賬戶適用
  required double frozenCash = 5; //凍結資金
  required double debtCash = 6; //計息金額
  required double avlWithdrawalCash = 7; //現金可提（僅單幣種賬戶使用此欄位，綜合賬戶請使用 cashInfoList 獲取分幣種現金可提）

  optional int32 currency = 8;            //幣種，本結構體資金相關的貨幣類型，取值參見 Currency，期貨和綜合證券賬戶適用
  optional double availableFunds = 9;     //可用資金，期貨適用
  optional double unrealizedPL = 10;      //未實現盈虧，期貨適用
  optional double realizedPL = 11;        //已實現盈虧，期貨適用
  optional int32 riskLevel = 12;           //風控狀態，參見 CltRiskLevel, 期貨適用。建議統一使用 riskStatus 欄位獲取證券、期貨賬戶的風險狀態
  optional double initialMargin = 13;      //初始保證金
  optional double maintenanceMargin = 14;  //維持保證金
  repeated AccCashInfo cashInfoList = 15;  //分幣種的現金、現金可提和現金購買力（僅綜合賬戶適用）
  optional double maxPowerShort = 16; //賣空購買力（此欄位是按照 60% 的融券保證金率計算得到的近似值。但事實上，每個標的的融券保證金率並不相同。我們建議您使用 查詢最大可買可賣 接口返回的 可賣空 欄位，來判斷實際可賣空的最大數量。）
  optional double netCashPower = 17;  //現金購買力（僅單幣種賬戶使用此欄位，綜合賬戶請使用 cashInfoList 獲取分幣種現金購買力）
  optional double longMv = 18;        //多頭市值
  optional double shortMv = 19;       //空頭市值
  optional double pendingAsset = 20;  //在途資產
  optional double maxWithdrawal = 21;          //融資可提，僅證券賬戶適用
  optional int32 riskStatus = 22;              //風險狀態，參見 CltRiskStatus，共分 9 個等級，LEVEL1是最安全，LEVEL9是最危險
  optional double marginCallMargin = 23;       //	Margin Call 保證金

  optional bool isPdt = 24;				//是否PDT賬戶，僅moomoo證券(美國)賬戶適用
  optional string pdtSeq = 25;			//剩餘日內交易次數，僅被標記為 PDT 的moomoo證券(美國)賬戶適用
  optional double beginningDTBP = 26;		//初始日內交易購買力，僅被標記為 PDT 的moomoo證券(美國)賬戶適用
  optional double remainingDTBP = 27;		//剩餘日內交易購買力，僅被標記為 PDT 的moomoo證券(美國)賬戶適用
  optional double dtCallAmount = 28;		//日內交易待繳金額，僅被標記為 PDT 的moomoo證券(美國)賬戶適用
  optional int32 dtStatus = 29;				//日內交易限制情況，取值見 DTStatus。僅被標記為 PDT 的moomoo證券(美國)賬戶適用
  
  optional double securitiesAssets = 30; // 證券資產淨值
  optional double fundAssets = 31; // 基金資產淨值
  optional double bondAssets = 32; // 債券資產淨值

  repeated AccMarketInfo marketInfoList = 33; //分市場資產資訊
}
```

## 賬戶持倉

**Position**

```protobuf
message Position
{
    required uint64 positionID = 1;     //持倉 ID，一條持倉的唯一標識
    required int32 positionSide = 2;    //持倉方向，參見 PositionSide 的枚舉定義
    required string code = 3;           //代碼
    required string name = 4;           //名稱
    required double qty = 5;            //持有數量，2位精度，期權單位是"張"，下同
    required double canSellQty = 6;     //可用數量，是指持有的可平倉的數量。可用數量=持有數量-凍結數量。期權和期貨的單位是“張”。
    required double price = 7;          //市價，3位精度，期貨為2位精度
    optional double costPrice = 8;      //攤薄成本價（證券賬戶），平均開倉價（期貨賬戶）。證券無精度限制，期貨為2位精度，如果沒傳，代表此時此值無效
    required double val = 9;            //市值，3位精度, 期貨此欄位值為0
    required double plVal = 10;         //盈虧金額，3位精度，期貨為2位精度
    optional double plRatio = 11;       //盈虧百分比(平均成本價模式)，無精度限制，如果沒傳，代表此時此值無效
    optional int32 secMarket = 12;      //證券所屬市場，參見 TrdSecMarket 的枚舉定義
    
	//以下是此持倉今日統計
    optional double td_plVal = 21;      //今日盈虧金額，3位精度，下同, 期貨為2位精度
    optional double td_trdVal = 22;     //今日交易額，期貨不適用
    optional double td_buyVal = 23;     //今日買入總額，期貨不適用
    optional double td_buyQty = 24;     //今日買入總量，期貨不適用
    optional double td_sellVal = 25;    //今日賣出總額，期貨不適用
    optional double td_sellQty = 26;    //今日賣出總量，期貨不適用

    optional double unrealizedPL = 28;       //未實現盈虧（僅期貨賬戶適用）
    optional double realizedPL = 29;         //已實現盈虧（僅期貨賬戶適用）	
    optional int32 currency = 30;        // 貨幣類型，取值參考 Currency
    optional int32 trdMarket = 31;  //交易市場, 參見 TrdMarket 的枚舉定義

    optional double dilutedCostPrice = 32;      //攤薄成本價，僅支援證券賬戶使用
    optional double averageCostPrice = 33;      //平均成本價，模擬交易證券賬戶不適用
    optional double averagePlRatio = 34;        //盈虧百分比(平均成本價模式)，無精度限制，如果沒傳，代表此時此值無效
}
```

## 訂單

**Order**

```protobuf
message Order
{
    required int32 trdSide = 1; //交易方向, 參見 TrdSide 的枚舉定義
    required int32 orderType = 2; //訂單類型, 參見 OrderType 的枚舉定義
    required int32 orderStatus = 3; //訂單狀態, 參見 OrderStatus 的枚舉定義
    required uint64 orderID = 4; //訂單號
    required string orderIDEx = 5; //擴展訂單號(僅查問題時備用)
    required string code = 6; //代碼
    required string name = 7; //名稱
    required double qty = 8; //訂單數量，2位精度，期權單位是"張"
    optional double price = 9; //訂單價格，3位精度
    required string createTime = 10; //建立時間，嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
    required string updateTime = 11; //最後更新時間，嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
    optional double fillQty = 12; //成交數量，2位精度，期權單位是"張"
    optional double fillAvgPrice = 13; //成交均價，無精度限制
    optional string lastErrMsg = 14; //最後的錯誤描述，如果有錯誤，會有此描述最後一次錯誤的原因，無錯誤為空
    optional int32 secMarket = 15; //證券所屬市場，參見 TrdSecMarket 的枚舉定義
    optional double createTimestamp = 16; //建立時間戳
    optional double updateTimestamp = 17; //最後更新時間戳
    optional string remark = 18; //用戶備註字串，最大長度64字節
    optional double auxPrice = 21; //觸發價格
    optional int32 trailType = 22; //追蹤類型, 參見Trd_Common.TrailType的枚舉定義
    optional double trailValue = 23; //追蹤金額/百分比
    optional double trailSpread = 24; //指定價差
    optional int32 currency = 25;        // 貨幣類型，取值參考 Currency
    optional int32 trdMarket = 26;  //交易市場, 參見TrdMarket的枚舉定義
    optional int32 session = 27; //美股訂單時段, 參見Common.Session的枚舉定義
    optional int32 jpAccType = 28; //JP子賬戶類型，取值見 TrdSubAccType
}
```

## 訂單費用條目

**OrderFeeItem**

```protobuf
message OrderFeeItem
{
    optional string title = 1; //費用名字
    optional double value = 2; //費用金額
}
```

## 訂單費用

**OrderFee**

```protobuf
message OrderFee
{
    required string orderIDEx = 1; //擴展訂單號
    optional double feeAmount = 2; //費用總額
    repeated OrderFeeItem feeList = 3; //費用明細
}
```

## 成交

**OrderFill**

```protobuf
message OrderFill
{
	required int32 trdSide = 1; //交易方向, 參見 TrdSide 的枚舉定義
    required uint64 fillID = 2; //成交號
    required string fillIDEx = 3; //擴展成交號(僅查問題時備用)
    optional uint64 orderID = 4; //訂單號
    optional string orderIDEx = 5; //擴展訂單號(僅查問題時備用)
    required string code = 6; //代碼
    required string name = 7; //名稱
    required double qty = 8; //成交數量，2位精度，期權單位是"張"
    required double price = 9; //成交價格，3位精度
    required string createTime = 10; //建立時間（成交時間），嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳
    optional int32 counterBrokerID = 11; //對手經紀號，港股有效
    optional string counterBrokerName = 12; //對手經紀名稱，港股有效
    optional int32 secMarket = 13; //證券所屬市場，參見 TrdSecMarket 的枚舉定義
    optional double createTimestamp = 14; //建立時間戳
    optional double updateTimestamp = 15; //最後更新時間戳
    optional int32 status = 16; //成交狀態, 參見 OrderFillStatus 的枚舉定義
    optional int32 trdMarket = 17;  //交易市場, 參見TrdMarket的枚舉定義
    optional int32 jpAccType = 18; //JP子賬戶類型，取值見 TrdSubAccType
}
```

## 最大可交易數量

**MaxTrdQtys**

```protobuf
message MaxTrdQtys
{
	//因目前伺服器實現的問題，賣空需要先賣掉多頭持倉才能再賣空，是分開兩步賣的，買回來同樣是逆向兩步；而看多的買是可以現金加融資一起一步買的，請注意這個差異
	required double maxCashBuy = 1;             //現金可買（期權的單位是“張”，期貨賬戶不適用）
    optional double maxCashAndMarginBuy = 2;    //最大可買（期權的單位是“張”，期貨賬戶不適用）
    required double maxPositionSell = 3;        //持倉可賣（期權的單位是“張”）
    optional double maxSellShort = 4;           //可賣空（期權的單位是“張”，期貨賬戶不適用）
    optional double maxBuyBack = 5;             //平倉需買入（當持有淨短倉時，必須先買回空頭持倉的股數，才能再繼續買多。期貨、期權的單位是“張”）
    optional double longRequiredIM = 6;         //買 1 張合約所帶來的初始保證金變動。僅期貨和期權適用。無持倉時，返回 買入 1 張的初始保證金佔用（正數）。有長倉時，返回 買入1 張的初始保證金佔用（正數）。有短倉時，返回 買回 1 張的初始保證金釋放（負數）。
    optional double shortRequiredIM = 7;        //賣 1 張合約所帶來的初始保證金變動。僅期貨和期權適用。無持倉時，返回 賣空 1 張的初始保證金佔用（正數）。 有長倉時，返回賣出1 張的初始保證金佔用（正數）。有短倉時，返回 賣空1 張的初始保證金釋放（正數）。
}
```

## 現金流水數據

**FlowSummaryInfo**

```protobuf
message FlowSummaryInfo
{
	optional string clearingDate = 1; //清算日期
	optional string settlementDate = 2; //結算日期
	optional int32 currency = 3; //幣種
	optional string cashFlowType = 4; //現金流類型
	optional int32 cashFlowDirection = 5; //現金流方向 TrdCashFlowDirection
	optional double cashFlowAmount = 6; //金額
	optional string cashFlowRemark = 7; //備註
	optional uint64 cashFlowID = 8; //現金流 ID
}
```

## 過濾條件

**TrdFilterConditions**

```protobuf
message TrdFilterConditions
{
  repeated string codeList = 1; //代碼過濾，只返回包含這些代碼的數據，沒傳不過濾
  repeated uint64 idList = 2; //ID 主鍵過濾，只返回包含這些 ID 的數據，沒傳不過濾，訂單是 orderID、成交是 fillID、持倉是 positionID
  optional string beginTime = 3; //開始時間，嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳，對持倉無效，拉歷史數據必須填
  optional string endTime = 4; //結束時間，嚴格按 YYYY-MM-DD HH:MM:SS 或 YYYY-MM-DD HH:MM:SS.MS 格式傳，對持倉無效，拉歷史數據必須填
  repeated string orderIDExList = 5; // 伺服器訂單ID列表，可以用來替代orderID列表，二選一
  optional int32 filterMarket = 6; //指定交易市場, 參見TrdMarket的枚舉定義
}
```

---



---

# 基礎功能

協議 ID|Protobuf 文件|説明
:-|:-|:-
1001|[InitConnect](../quote/base.md)|初始化連線
1002|[GetGlobalState](../quote/get-global-state.md)|取得全域狀態
1003|[Notify](./init.md#1951)|事件通知推送
1004|[KeepAlive](./protocol.md#2603)|保持連線

## 設定 API 資訊

InitConnect.proto  

```protobuf
message C2S
{
    required int32 clientVer = 1; //客户端版本號，clientVer = "."以前的數 * 100 + "."以後的，舉例：1.1版本的 clientVer 為1 * 100 + 1 = 101，2.21版本為2 * 100 + 21 = 221
    required string clientID = 2; //客户端唯一標識，無生具體生成規則，客户端自己保證唯一性即可
    optional bool recvNotify = 3; //此連接是否接收市場狀態、交易需要重新解鎖等等事件通知，true 代表接收，OpenD 就會向此連接推送這些通知，反之 false 代表不接收不推送
    //如果通信要加密，首先得在 OpenD 和客户端都配置 RSA 密鑰，不配置始終不加密
    //如果配置了 RSA 密鑰且指定的加密算法不為 PacketEncAlgo_None 則加密(即便這裏不設置，配置了 RSA 密鑰，也會採用默認加密方式)，默認採用 FTAES_ECB 算法
    optional int32 packetEncAlgo = 4; //指定包加密算法，參見 Common.PacketEncAlgo 的枚舉定義
    optional int32 pushProtoFmt = 5; //指定這條連接上的推送協議格式，若不指定則使用 push_proto_type 配置項
}
```

* **介紹**

    在初始化連線協議的 clientVer 及 clientID 欄位中設定相關資訊。

## 設置協議格式

InitConnect.proto  

```protobuf
message C2S
{
    required int32 clientVer = 1; //客户端版本號，clientVer = "."以前的數 * 100 + "."以後的，舉例：1.1版本的 clientVer 為1 * 100 + 1 = 101，2.21版本為2 * 100 + 21 = 221
    required string clientID = 2; //客户端唯一標識，無生具體生成規則，客户端自己保證唯一性即可
    optional bool recvNotify = 3; //此連接是否接收市場狀態、交易需要重新解鎖等等事件通知，true 代表接收，OpenD 就會向此連接推送這些通知，反之 false 代表不接收不推送
    //如果通信要加密，首先得在 OpenD 和客户端都配置 RSA 密鑰，不配置始終不加密
    //如果配置了 RSA 密鑰且指定的加密算法不為 PacketEncAlgo_None 則加密(即便這裏不設置，配置了 RSA 密鑰，也會採用默認加密方式)，默認採用 FTAES_ECB 算法
    optional int32 packetEncAlgo = 4; //指定包加密算法，參見 Common.PacketEncAlgo 的枚舉定義
    optional int32 pushProtoFmt = 5; //指定這條連接上的推送協議格式，若不指定則使用 push_proto_type 配置項
}
```


* **介紹**

    可在初始化連線協議的 pushProtoFmt 欄位中，指定該連線的推送數據格式。
    請求的數據格式，請參考[協議頭](../ftapi/protocol.md#2260) 中的 nProtoFmtType 欄位。

## 對所有連接設置協議加密


## 設置私鑰路徑

InitConnect.proto  

```protobuf
message C2S
{
        required int32 clientVer = 1; //客户端版本號，clientVer = "."以前的數 * 100 + "."以後的，舉例：1.1版本的 clientVer 為1 * 100 + 1 = 101，2.21版本為2 * 100 + 21 = 221
        required string clientID = 2; //客户端唯一標識，無生具體生成規則，客户端自己保證唯一性即可
        optional bool recvNotify = 3; //此連接是否接收市場狀態、交易需要重新解鎖等等事件通知，true 代表接收，OpenD 就會向此連接推送這些通知，反之 false 代表不接收不推送
        //如果通信要加密，首先得在 OpenD 和客户端都配置 RSA 密鑰，不配置始終不加密
        //如果配置了 RSA 密鑰且指定的加密算法不為 PacketEncAlgo_None 則加密(即便這裏不設置，配置了 RSA 密鑰，也會採用默認加密方式)，默認採用 FTAES_ECB 算法
        optional int32 packetEncAlgo = 4; //指定包加密算法，參見 Common.PacketEncAlgo 的枚舉定義
        optional int32 pushProtoFmt = 5; //指定這條連接上的推送協議格式，若不指定則使用 push_proto_type 配置項
}
```


* **介紹**

    在初始化連線協議中 packetEncAlgo 字段指定該連接上加密算法。  
    協議加密詳情，參見[加密通信流程](./protocol.md#5455)

## 設置線程模式


## 設置回調


## 獲取連接 ID

InitConnect.proto

```protobuf
message S2C
{
	required int32 serverVer = 1; //OpenD 的版本號
	required uint64 loginUserID = 2; //OpenD 登錄的用户 ID
	required uint64 connID = 3; //此連接的連接 ID，連接的唯一標識
	required string connAESKey = 4; //此連接後續 AES 加密通信的 Key，固定為16字節長字符串
	required int32 keepAliveInterval = 5; //心跳保活間隔
	optional string aesCBCiv = 6; //AES 加密通信 CBC 加密模式的 iv，固定為16字節長字符串
}
```

* **介紹**

    即 InitConnect 協議回傳資料中的 connID 欄位

## 事件通知回調

Notify.proto

```protobuf
message S2C
{
	required int32 type = 1; //通知類型
	optional GtwEvent event = 2; //事件通息
	optional ProgramStatus programStatus = 3; //程序狀態
	optional ConnectStatus connectStatus = 4; //連接狀態
	optional QotRight qotRight = 5; //行情權限
	optional APILevel apiLevel = 6; //用户等級，已在2.10版本之後廢棄
	optional APIQuota apiQuota = 7; //API 額度
}

message Response
{
	required int32 retType = 1 [default = -400]; //RetType,返回結果
	optional string retMsg = 2;
	optional int32 errCode = 3;
	
	optional S2C s2c = 4;
}
```

---



---

# 通用定義

## API 調用結果

**RetType**  

```protobuf
enum RetType
{
    RetType_Succeed = 0; //成功
    RetType_Failed = -1; //失敗
    RetType_TimeOut = -100; //超時
    RetType_Unknown = -400; //未知結果
}
```

## 協議格式

**ProtoFmt**  

```protobuf
enum ProtoFmt
{
    ProtoFmt_Protobuf = 0; //Google Protobuf 格式
    ProtoFmt_Json = 1; //Json 格式
}
```

## 封包加密演算法

**PacketEncAlgo**  

```protobuf
enum PacketEncAlgo
{
    PacketEncAlgo_FTAES_ECB = 0; //修改過的 AES 的 ECB 加密模式
    PacketEncAlgo_None = -1; //不加密
    PacketEncAlgo_AES_ECB = 1; //標準的 AES 的 ECB 加密模式
    PacketEncAlgo_AES_CBC = 2; //標準的 AES 的 CBC 加密模式
}
```

## 程式狀態類型

**ProgramStatusType**  

```protobuf
enum ProgramStatusType
{
    ProgramStatusType_None = 0;
    ProgramStatusType_Loaded = 1; //已完成類似加載配置,啟動服務器等操作,服務器啟動之前的狀態無需返回

    ProgramStatusType_Loging = 2; //登錄中
    ProgramStatusType_NeedPicVerifyCode = 3; //需要圖形驗證碼
    ProgramStatusType_NeedPhoneVerifyCode = 4; //需要手機驗證碼
    ProgramStatusType_LoginFailed = 5; //登錄失敗,詳細原因在描述返回
    ProgramStatusType_ForceUpdate = 6; //客端戶版本過舊

    ProgramStatusType_NessaryDataPreparing = 7; //正在拉取類似免責聲明等一些必要資訊
    ProgramStatusType_NessaryDataMissing = 8; //缺少必要資訊
    ProgramStatusType_UnAgreeDisclaimer = 9; //未同意免責聲明
    ProgramStatusType_Ready = 10; //可以接收業務協議收發,正常可用狀態
	
    //OpenD 登入後被強制登出，會導致連接全部斷開,需要重連後才能得到以下該狀態（並且需要在 ui 模式下）
    ProgramStatusType_ForceLogout = 11; //被強制退出登錄,例如修改了登錄密碼,中途打開設備鎖等,詳細原因在描述返回

    ProgramStatusType_DisclaimerPullFailed = 12; //拉取免責聲明標誌失敗
}
```

## 網關事件通知類型

**GtwEventType**  

```protobuf
enum GtwEventType
{
    GtwEventType_None = 0; //正常無錯
    GtwEventType_LocalCfgLoadFailed	= 1; //加載本地配置失敗
    GtwEventType_APISvrRunFailed = 2; //服務器啟動失敗
	  GtwEventType_ForceUpdate = 3; //客端戶版本過舊
	  GtwEventType_LoginFailed = 4; //登錄失敗
	  GtwEventType_UnAgreeDisclaimer = 5; //未同意免責聲明
	  GtwEventType_NetCfgMissing = 6; //缺少必要網絡配置資訊;例如控制訂閲額度 //已優化，不會再出現該情況
	  GtwEventType_KickedOut = 7; //牛牛帳號在別處登錄
	  GtwEventType_LoginPwdChanged = 8; //登錄密碼被修改
	  GtwEventType_BanLogin = 9; //用户被禁止登錄
	  GtwEventType_NeedPicVerifyCode = 10; //需要圖形驗證碼
	  GtwEventType_NeedPhoneVerifyCode = 11; //需要手機驗證碼
	  GtwEventType_AppDataNotExist = 12; //程式自帶數據不存在
	  GtwEventType_NessaryDataMissing = 13; //缺少必要數據
	  GtwEventType_TradePwdChanged = 14; //交易密碼被修改
	  GtwEventType_EnableDeviceLock = 15; //啟用設備鎖
}
```


## 系統通知類型

**NotifyType**  

```protobuf
enum NotifyType
{
    NotifyType_None = 0; //無
	  NotifyType_GtwEvent = 1; //OpenD 運行事件通知
	  NotifyType_ProgramStatus = 2; //程式狀態
	  NotifyType_ConnStatus = 3; //連接狀態
	  NotifyType_QotRight = 4; //行情權限
	  NotifyType_APILevel = 5; //用户等級，已在2.10版本之後廢棄
	  NotifyType_APIQuota = 6; //API 額度
}
```

## 封包唯一識別碼

**PacketID** 

```protobuf
message PacketID
{
	  required uint64 connID = 1; //當前 TCP 連接的連接 ID，一條連接的唯一標識，InitConnect 協議會返回
	  required uint32 serialNo = 2; //自增序列號
}
```

## 程式狀態

**ProgramStatus**

```protobuf
message ProgramStatus
{
	  required ProgramStatusType type = 1; //當前狀態
	  optional string strExtDesc = 2; // 額外描述
}
```

---



---

# 底層協議介紹

Futu API 是富途為主流的程式語言（Python、Java、C#、C++、JavaScript）封裝的 API SDK，以方便您調用，降低策略開發難度。  
這部分主要介紹策略腳本與 OpenD 服務之間通訊的底層協議，適用於非上述 5 種程式語言用戶，自行對接實現底層裸協議。

:::tip 提示
* 如果您使用的程式語言在上述的 5 種主流程式語言之內，可以直接跳過這部分內容。
:::

## 協議請求流程
* 建立連接
* 初始化連接
* 請求數據或接收推送數據
* 定時發送 KeepAlive 保持連接

![proto-process](../img/proto-process.png)


## 協議設計
協議數據包括協議頭以及協議體，協議頭固定欄位，協議體根據具體協議決定。

### 協議頭

```
struct APIProtoHeader
{
    u8_t szHeaderFlag[2];
    u32_t nProtoID;
    u8_t nProtoFmtType;
    u8_t nProtoVer;
    u32_t nSerialNo;
    u32_t nBodyLen;
    u8_t arrBodySHA1[20];
    u8_t arrReserved[8];
};
```
欄位|説明
:-|:-
szHeaderFlag|包頭起始標誌，固定為“FT”
nProtoID|協議 ID
nProtoFmtType|協議格式類型，0 為 Protobuf 格式，1 為 Json 格式
nProtoVer|協議版本，用於迭代兼容，目前填 0
nSerialNo|包序列號，用於對應請求包和回傳封包 / 回傳資料，要求遞增
nBodyLen|包體長度
arrBodySHA1|包體原始數據(解密後)的 SHA1 雜湊值 (Hash)
arrReserved|保留 8 位元組 (Byte)擴展

::: tip 提示
* u8_t 表示 8 位無符號整數，u32_t 表示 32 位無符號整數
* OpenD 內部處理使用 Protobuf，因此協議格式建議使用 Protobuf，減少 Json 轉換開銷
* nProtoFmtType 欄位指定了包體的數據類型，回傳封包 / 回傳資料會回對應類型的數據；推送協議數據類型由 OpenD 配置檔案指定
* **arrBodySHA1 用於驗證請求數據在網絡傳輸前後的一致性，必須正確填入**
* **協議頭的二進制流使用的是小端位元組 (Byte)序，即一般不需要使用 ntohl 等相關函數轉換數據**
:::

### 協議體
#### Protobuf 協議請求包體結構
```
message C2S
{
    required int64 req = 1;
}

message Request
{
    required C2S c2s = 1;
}
```

#### Protobuf 協議回應包體結構
```
message S2C
{
    required int64 data = 1;
}

message Response
{
    required int32 retType = 1 [default = -400]; //RetType，返回結果
    optional string retMsg = 2;
    optional int32 errCode = 3;
    optional S2C s2c = 4;
}
```

欄位|説明
:-|:-
c2s|請求參數結構
req|請求參數，實際根據協議定義
retType|請求結果
retMsg|若請求失敗，説明失敗原因
errCode|若請求失敗對應錯誤碼
s2c|回應數據結構，部分協議不返回數據則無該欄位
data|回應數據，實際根據協議定義

::: tip 提示
* 包體格式類型請求包由協議頭 nProtoFmtType 指定，OpenD 主動推送格式在 [InitConnect](../ftapi/init.md#1061) 設置。
* 原始協議檔案格式是以 Protobuf 格式定義，若需要 json 格式傳輸，建議使用 protobuf3 的介面 / API直接轉換成 json。
* 枚舉值欄位定義使用有符號整形，註解指明對應枚舉，枚舉一般定義於 Common.proto，Qot_Common.proto，Trd_Common.proto 檔案中。
* 協議中價格、百分比等數據用浮點類型來傳輸，直接使用會有精度問題，需要根據精度（如協議中未指明，預設小數點後三位）做四捨五入之後再使用。
:::

## 心跳保活

```protobuf
syntax = "proto2";
package KeepAlive;
option java_package = "com.futu.openapi.pb";
option go_package = "github.com/futuopen/ftapi4go/pb/keepalive";

import "Common.proto";

message C2S
{
	required int64 time = 1; //客户端發包時的格林威治時間戳，單位秒
}

message S2C
{
	required int64 time = 1; //伺服器回傳封包 / 回傳資料時的格林威治時間戳，單位秒
}

message Request
{
	required C2S c2s = 1;
}

message Response
{
	required int32 retType = 1 [default = -400]; //RetType,返回結果
	optional string retMsg = 2;
	optional int32 errCode = 3;
	
	optional S2C s2c = 4;
}
```

* **介紹**

    心跳保活

* **協議 ID**

    1004

* **使用**

    根據[初始化連結](./init.md#1061)返回的心跳保活間隔時間，向 OpenD 發送保活協議

## 加密通訊流程

* 若 OpenD 配置了加密，[InitConnect](../ftapi/init.md#1061) 初始化連接協議必須使用 [RSA](../qa/other.md#9747) 公鑰加密，後續其他協議使用 InitConnect 返回的隨機密鑰進行 AES 加密通訊。
* OpenD 的加密流程借鑑了 SSL 協議，但考慮到一般是本地部署服務和應用，簡化了相關流程，OpenD 與接入 Client 共用了同一個 [RSA](../qa/other.md#9747) 私鑰檔案，請妥善保存和分發私鑰檔案。
* 可到這個 [網址](http://web.chacuo.net/netrsakeypair) 在線生成隨機 [RSA](../qa/other.md#9747) 密鑰對，密鑰格式必須為 PCKS#1，密鑰長度 512，1024 都可以，不要設置密碼，將生成的私鑰複製保存到檔案中，然後將私鑰檔案路徑配置到 [OpenD 配置](../opend/opend-cmd.md#1028) 約定的 **rsa_private_key** 配置項中。
*  **建議有實盤交易的用戶配置加密，避免賬户和交易資訊泄露。**

![encrypt](../img/encrypt.png)


## RSA 加解密
* [OpenD 配置](../opend/opend-cmd.md#1028) 約定 **rsa_private_key** 為私鑰檔案路徑
* OpenD 與接入客户端共用相同的私鑰檔案
* RSA 加解密僅用於 InitConnect 請求，用於安全獲取其它請求協議的對稱加密 Key
* OpenD 的 [RSA](../qa/other.md#9747) 密鑰為 1024 位，填充方式 PKCS1，公鑰加密，私鑰解密，公鑰可通過私鑰生成
* Python API 參考實現：[RsaCrypt](https://github.com/FutunnOpen/py-futu-api/tree/master/futu/common/sys_config.py) 類的 encrypt / decrypt 介面 / API

### 發送數據加密
* RSA 加密規則:若密鑰位數是 key_size，單次加密串的最大長度為 (key_size)/8 - 11，目前位數 1024，一次加密長度可定為 100。
* 將明文數據分成一個或數個最長 100 位元組 (Byte)的小段進行加密，拼接分段加密數據即為最終的 Body 加密數據。

### 接收數據解密
* RSA 解密同樣遵循分段規則，對於 1024 位密鑰，每小段待解密數據長度為 128 位元組 (Byte)。
* 將密文數據分成一個或數個 128 位元組 (Byte)長的小段進行解密，拼接分段解密數據即為最終的 Body 解密數據。

## AES 加解密
* 加密 key 由 InitConnect 協議返回
* 預設使用的是 AES 的 ecb 加密模式。
* Python API 參考實現: [ConnMng](https://github.com/FutunnOpen/py-futu-api/tree/master/futu/common/conn_mng.py) 類的 encrypt_conn_data / decrypt_conn_data 介面 / API

### 發送數據加密

* AES 加密要求源數據長度必須是 16 的整數倍，故需補‘0’對齊後再加密，記錄 mod_len 為源數據長度與 16 取模值。
* 因加密前有可能對源數據作修改，故需在加密後的數據尾再增加一個 16 位元組 (Byte)的填充數據塊，其最後一個位元組 (Byte)賦值 mod_len，其餘位元組 (Byte)賦值‘0’，將加密數據和額外的填充數據塊拼接作為最終要發送協議的 body 數據。

### 接收數據解密

* 協議 body 數據，先將最後一個位元組 (Byte)取出，記為 mod_len，然後將 body 截掉尾部 16 位元組 (Byte)填充數據塊後再解密（與加密填充額外數據塊邏輯對應）。
* mod_len 為 0 時，上述解密後的數據即為協議返回的 body 數據，否則需截掉尾部(16 - mod_len)長度的用於填充對齊的數據。

![aes](../img/aes.png)

---



---

# OpenD 相關


## Q1：OpenD 因未完成“問卷評估及協議確認”自動退出

A: 您需要進行相關問卷評估及協議確認，才可以使用 OpenD，請先 [前往完成](https://www.futunn.com/about/api-disclaimer?lang=zh-CN)。

## Q2：OpenD 因”程式內置數據不存在“退出

A: 一般因權限問題導致自帶數據拷貝失敗，可以嘗試將程式目錄下 <font color=Gray> __*Appdata.dat*__ </font> 解壓後的檔案複製到程式資料目錄下。

* windows 程式資料目錄:`%appdata%/com.futunn.FutuOpenD/F3CNN`
* 非 windows 程式資料目錄:`~/.com.futunn.FutuOpenD/F3CNN`

## Q3：OpenD 服務啟動失敗

A: 請檢查：
1. 是否有其他程式佔用所配置的端口；
2. 是否已經有配置了相同端口的 OpenD 在運行。

## Q4：如何驗證手機驗證碼？

A: 在 OpenD 界面上或遠程連接至 Telnet 端口，輸入命令`input_phone_verify_code -code=123456`。

::: tip 提示
* 123456 是收到的手機驗證碼
* -code=123456 前有空格
:::

## Q5：是否支持其他編程語言？

A: OpenD 有對外提供基於 socket 的協議，目前我們提供並維護 Python，C++，Java，C# 和 JavaScript 接口，[下載入口](https://www.futunn.com/download/OpenAPI)。

如果上述語言仍不能滿足您的需求，您可以自行整合 Protobuf 協議。

## Q6：在同一設備多次驗證設備鎖 

A: 設備標識隨機生成並存放於 

windows: %appdata%/com.futunn.FutuOpenD/F3CNN/Device.dat 檔案中。
非windows: ~/.com.futunn.FutuOpenD/F3CNN/Device.dat

::: tip 提示
1. 如果檔案被刪除或損壞，OpenD 會重新生成新設備標識，然後驗證設備鎖。  
2. 另外映像檔複製部署的用戶需要注意，如果多台機器的 Device.dat 內容相同，也會導致這些機器多次驗證設備鎖。刪除 Device.dat 檔案即可解決。
:::

## Q7：OpenD 是否有提供 Docker 鏡像？

A: 目前沒有提供。

## Q8：一個賬號可以登入多個 OpenD 嗎？

A: 一個賬號可以在多台機器上登入 OpenD 或者其他客戶端，最多允許 10 個 OpenD 終端同時登入。同時有“行情互踢”的限制，只能有一個 OpenD 獲得最高權限行情。例如：兩個終端登入同一個賬號，只能有一個港股 LV2 行情，另一個是港股 BMP 行情。

## Q9：如何控制 OpenD 和其他客户端（桌面版和流動版）的行情權限？

A: 應交易所的規定，多個終端同時在線會有“行情互踢”的限制，只能有一個終端獲得最高權限行情。OpenD 命令行版本的啟動參數中，內置了 [auto_hold_quote_right](../opend/opend-cmd.md#1028) 參數，用於靈活配置行情權限。當該參數選項開啟時，OpenD 在行情權限被搶後，會自動搶回。如果 10 秒內再次被搶，則其他終端獲得最高行情權限（OpenD 不會再搶）。

## Q10：如何優先保證 OpenD 行情權限？

A: 
1. 將 OpenD 啟動參數 [auto_hold_quote_right](../opend/opend-cmd.md#1028) 配置為 1；
2. 保證不要在流動版或桌面版富途牛牛上在 10 秒內連續兩次搶最高權限（登入算一次，點擊“重啟行情”算第二次）。

![quote-right-kick](../img/quote-right-kick.png)

## Q11：如何優先保證流動版（或桌面版）的行情權限？

A: OpenD 啟動參數 [auto_hold_quote_right](../opend/opend-cmd.md#1028) 設置為 0，流動版或桌面版富途牛牛在 OpenD 之後登入即可。 

## Q12：使用可視化 OpenD 記住密碼登入，長時間掛機後提示連接斷開，需要重新登入？

A: 使用可視化 OpenD，如果選擇記住密碼登入，用的是記錄在本地的令牌。由於令牌有時間限制，當令牌過期後，如果出現網絡波動或富途後台發佈，就可能導致與後台斷開連接後無法自動連接上的情況。因此，可視化 OpenD 如果希望長時間掛機，建議手動輸入密碼登入，由 OpenD 自動處理該情況。


## Q13：遇到產品缺陷，如何請富途的研發工程師排查日誌？

A: 
1. 與客服溝通問題表現，詳述：發生錯誤的時間、OpenD 版本號、 API 版本號、程式語言名稱、接口名或通訊協定編號、含詳細輸入參數和回傳的短代碼或截圖。

2. 客服確認是產品缺陷後，如需進一步日誌排查，研發工程師會主動聯繫。

3. 部分問題須提供 OpenD 日誌，方便定位確認問題。交易類問題需要 info 日誌級別，行情類問題需要 debug 日誌級別。日誌級別 log_level 可以在 <font color=Gray> __*OpenD.xml*__ </font> 中 [配置](../opend/opend-cmd.md#2092) ，配置後需要重啟 OpenD 方能生效，待問題重現後，將該段日誌打包發給富途研發工程師。

:::tip 提示
日誌路徑如下：  
windows：`%appdata%/com.futunn.FutuOpenD/Log`

非 windows：`~/.com.futunn.FutuOpenD/Log`
:::

## Q14：程式連接不上 OpenD

A: 請先嘗試檢查：
1. 程式連接的端口與 OpenD 配置的端口是否一致。
2. 由於 OpenD 連接上限為 128，是否有無用連接未關閉。
3. 檢查監聽地址是否正確，如果腳本和 OpenD 不在同一機器，OpenD 監聽地址需要設置成 0.0.0.0 。

## Q15：連接上一段時間後斷開

A: 如果是自己整合協議，檢查下是否有定時發送心跳維持連接。


## Q16：Linux 下通過 multiprocessing 模塊以多進程方式運行 Python 腳本，連不上 OpenD？

A: Linux/Mac 環境下以預設方式創建進程後，父進程中 py-futu-api 內部創建的線程將會在子進程中消失，導致程式內部狀態錯誤。  
可以用 spawn 方式來啟動進程：

```python
import multiprocessing as mp
mp.set_start_method('spawn')
p = mp.Process(target=func)
```


## Q17：如何在一台電腦同時登入兩個 OpenD?

A: 可視化 OpenD 不支持，命令行 OpenD 支持。

1. 解壓從官網下載的檔案，複製整個命令行 OpenD 資料夾（如 OpenD_5.2.1408_Windows）得到副本（此處以 Windows 為例，其他系統可採取相同操作）。

![file-page](../img/nnfile-page.png)

2. 分別打開兩個命令行 OpenD 資料夾配置好兩份 OpenD.xml 檔案。

第一份配置檔案參數：api_port = 11111，login_account = 登入賬號1，login_pwd = 登入密碼1

第二份配置檔案參數：api_port = 11112，login_account = 登入賬號2，login_pwd = 登入密碼2

![order-page](../img/nnorder-page.png)

3. 配置完成後，分別打開兩個 OpenD 程式運行。

![fod-page](../img/nnfod-page.png)

4. 調用接口時，注意接口的參數`port`（OpenD 監聽端口）與 OpenD.xml 檔案中的參數`api_port`為對應關係  
例如：

```python
from futu import *

# 向賬號1登入的 OpenD 進行請求
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111, is_encrypt=False)
quote_ctx.close() # 結束後記得關閉當條連接，防止連接條數用盡

# 向賬號2登入的 OpenD 進行請求
quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11112, is_encrypt=False)
quote_ctx.close() # 結束後記得關閉當條連接，防止連接條數用盡
```

## Q18：行情權限被其他客户端踢掉，如何通過腳本執行搶權限的運維命令？
A：
1. 在OpenD啟動參數中，配置好 Telnet 地址和 Telnet 端口。
![telnet_GUI](../img/telnet_GUI.jpg)
![telnet_CMD](../img/telnet_CMD.jpg)
2. 啟動 OpenD（會同時啟動 Telnet）。
3. 當發現行情權限被搶之後，您可以參考如下代碼示例，通過 Telnet，向 OpenD 發送 `request_highest_quote_right` 命令。
```python
from telnetlib import Telnet
with Telnet('127.0.0.1', 22222) as tn:  # Telnet 地址為：127.0.0.1，Telnet 端口為：22222
    tn.write(b'request_highest_quote_right\r\n')
    reply = b''
    while True:
        msg = tn.read_until(b'\r\n', timeout=0.5)
        reply += msg
        if msg == b'':
            break
    print(reply.decode('gb2312'))
```

<span id="update-failed-qa"></span>

## Q19：OpenD 自動升級失敗

A：
通過`update`命令執行 OpenD 自動更新失敗，可能的原因：
- 檔案被其他進程佔用：可以嘗試關閉其他 OpenD 進程，或者重啟系統後，再次執行 `update`
如果以上仍無法解決，可以通過[官網](https://www.futunn.com/download/OpenAPI?lang=zh-CN)自行下載更新。

## Q20：ubuntu22無法啟動可視化 OpenD？
A：
在有些Linux發行版（例如Ubuntu 22.04）運行可視化OpenD時，可能會提示：`dlopen(): error loading libfuse.so.2`。
這是因為這些系統沒有預設安裝libfuse。通常可以手動安裝來解決，例如對於Ubuntu22.04，可以在命令行運行：
```
sudo apt update
sudo apt install -y libfuse2
```
安裝成功後就可以正常運行可視化OpenD了。詳細資訊請參考：[https://docs.appimage.org/user-guide/troubleshooting/fuse.html](https://docs.appimage.org/user-guide/troubleshooting/fuse.html)。

## Q21：Linux上如何在後台運行命令行OpenD？


A：先切到 FutuOpenD 所在目錄，配置好 FutuOpenD.xml 之後，執行如下命令
```
nohup ./FutuOpenD &
```

---



---

# 行情相關


## Q1：訂閲失敗

A: 訂閲接口返回錯誤，有以下兩類常見情況：
* 訂閲額度不足：

  訂閲額度規則參見 [訂閲額度 & 歷史 K 線額度](../intro/authority.md#4499)

* 訂閲權限不足：

  支持訂閲的行情權限見下表
  <table>
    <tr>
      <th> 市場 </th>
      <th> 品種 </th>
      <th> 支持訂閲的行情權限 </th>
    </tr>
    <tr>
      <td rowspan="3"> 香港市場 </td>
      <td > 股票 </td>
      <td > LV1, LV2, SF </td>
    </tr>
    <tr>
	    <td> 期權</td>
      <td> LV1, LV2</td>
    </tr>
    <tr>
	    <td> 期貨</td>
      <td> LV1, LV2</td>
    </tr>
    <tr>
      <td rowspan="3"> 美國市場 </td>
      <td > 股票 </td>
      <td > LV1, LV2 </td>
    </tr>
    <tr>
	    <td> 期權</td>
      <td> LV1</td>
    </tr>
    <tr>
	    <td> 期貨</td>
      <td> LV1, LV2</td>
    </tr>
    <tr>
      <td > A 股市場 </td>
      <td > 股票 </td>
      <td > LV1 </td>
    </tr>  
</table>

  獲取行情權限的方式參見 [行情權限](../intro/authority.html#5731) 

  注意：若賬號擁有上述權限，但仍訂閲失敗，可能存在被其他終端 [踢掉行情權限](./opend.html#7124) 的情況。

## Q2：取消訂閱失敗

A: 訂閲至少一分鐘後才能取消訂閱。

## Q3：取消訂閱成功但沒釋放額度

A: 所有連接都對該行情取消訂閱，才會釋放額度。

舉例：A 連接和 B 連接都在訂閲 HK.00700 的掛盤，當 A 連接取消訂閱後，由於 B 連接仍在調用騰訊的掛盤數據，因此 OpenD 的額度不會釋放，直至所有連接都取消訂閱 HK.00700 的掛盤。


## Q4：訂閲不足一分鐘關閉程式連接，會釋放額度嗎？

A: 不會。連接關閉後，訂閲時長不足一分鐘的標的類型，會在達到一分鐘後才自動取消訂閱，並釋放相應的訂閲額度。


## Q5：請求頻率限制的具體限制邏輯是怎樣？

A: 30 秒內最多 n 次，是指第 1 次和第 n+1 次請求間隔需要大於 30 秒。

## Q6：自選股添加不上是什麼原因？

A: 請先檢查是否有超出上限，或者刪除一部分自選。

## Q7：為什麼 API 端的美股報價和牛牛顯示端的全美綜合報價有不同？

A: 由於美股交易分散在很多家交易所，富途有提供兩種美股基本報價行情，一種是 Nasdaq Basic（Nasdaq 交易所的報價），另一種是全美綜合報價（全美13家交易所的報價）。而 Futu API 的美股正股行情目前僅支持通過行情卡購買的方式獲取 Nasdaq Basic，不支持全美綜合報價。因此，如果您同時購買了顯示端的全美綜合報價行情卡，和僅用於 Futu API 的 Nasdaq Basic 行情卡，確實有可能出現牛牛顯示端和 Futu API 端的報價差異。   
因此，如果您發現美股當天開市價與客戶端顯示不一致，這是因為Futu API實時上游行情僅會獲取 Nasdaq Basic 數據。


## Q8：API 行情卡在哪裏購買？

A:  
* 港股市場
  * [港股 LV2 高級行情（僅港澳台及海外 IP）](https://qtcardfthk.futufin.com/buy?market_id=1&amp;channel=2&amp;good_type=1#/)
  * [港股期權期貨 LV2高級行情（僅港澳台及海外 IP）](https://qtcardfthk.futufin.com/buy?market_id=1&amp;channel=2&amp;good_type=8#/)
  * [港股 LV2 + 期權期貨 LV2 行情（僅港澳台及海外 IP）](https://qtcardfthk.futufin.com/buy?market_id=1&amp;channel=2&amp;good_type=9#/)
  * [港股高級全盤行情（SF 行情）](https://qtcardfthk.futufin.com/buy?market_id=1&amp;channel=2&amp;good_type=10#/)
  
* 美股市場
  * [Nasdaq Basic](https://qtcardfthk.futufin.com/buy?market_id=2&amp;channel=2&amp;good_type=12#/)
  * [Nasdaq Basic+TotalView (Non-Pro)](https://qtcardfthk.futufin.com/buy?market_id=2&good_type=18&channel=2#/)
  * [Nasdaq Basic+TotalView (Pro)](https://qtcardfthk.futufin.com/buy?market_id=2&good_type=19&channel=2#/)
  * [期權 OPRA 實時行情](https://qtcardfthk.futufin.com/buy?market_id=2&good_type=16&qtcard_channel=2#/)


## Q9：為什麼有時候，獲取實時數據的 get 接口響應比較慢？

A: 因為獲取實時數據的 get 接口需要先訂閲，並依賴後台給 OpenD 的推送。如果用戶剛訂閲就立刻用 get 接口請求，OpenD 有可能尚未收到後台推送。為了防止這種情況的發生，get 接口內置了等待邏輯，3 秒內收到推送會立刻返回給程式，超過 3 秒仍未收到後台推送，才會給程式返回空數據。  
涉及的 get 接口包括：get_rt_ticker、get_rt_data、get_cur_kline、get_order_book、get_broker_queue、get_stock_quote。因此，當發現獲取實時數據的 get 接口響應比較慢時，可以先檢查一下是否是無成交數據的原因。


## Q10：購買 API 美股 Nasdaq Basic 行情卡後，可以獲取哪些數據？

A: Nasdaq Basic 行情卡購買啟用後，可以獲取的品類涵蓋 Nasdaq、NYSE、NYSE MKT 交易所上市證券（包括美股正股和 ETF，不包括美股期貨和美股期權）。  
支持的數據接口包括：快照，歷史 K 線，實時逐筆訂閲，實時一檔掛盤訂閲，實時 K 線訂閲，實時報價訂閲，實時分時訂閲，到價提醒。

## Q11：各個行情品類的掛盤支持多少檔？

A: 
行情品類|LV1|LV2|SF
:-|:-|:-|:-
港股（含正股、窩輪、牛熊、界內證）|/|10|全盤+千筆明細
港股期權期貨|1|10|/
美股（含 ETF）|1|60檔|/
美股期權|1|/|/
美股期貨 |/|40檔|/
A 股|5|/|/

## Q12：為什麼我購買啟用了行情卡之後，OpenD 仍然沒有行情權限？

A:   
1. 由於 Futu API 的行情權限跟 APP 的行情權限不完全一樣，部分行情卡僅適用於 APP 端（例如：Futu API美股行情卡需單獨購買）。請先確認您所購買的行情卡是否是 OpenD 適用的。   
我們已將 Futu API 適用的 **所有** 行情卡列在《權限與限制》一節，請點擊 [這裏](/intro/authority.html#4440) 查看。
2. 行情卡購買啟用成功後，是立即生效的。請 **重新啟動 OpenD** 後，再次查看權限狀態。


## Q13：如何通過訂閲接口獲取實時行情？
**第一步：訂閲**  

將標的的代碼和數據類型傳入 [訂閲接口](../quote/sub.md)，完成訂閲。  

訂閲接口支持了實時報價、實時掛盤、實時逐筆、實時分時、實時 K 線、實時經紀隊列數據的獲取。訂閲成功後，OpenD 會持續收到富途服務器的實時數據推送。

注意：訂閲額度會根據您的總資產、交易筆數和交易量，來進行分配，具體規則參見 [訂閲額度 & 歷史 K 線額度](../intro/authority.md#4499)。所以，如果您的訂閲額度不足，可以先檢查一下是否有無用的訂閲在佔用額度，及時 [取消訂閱](../quote/sub.md) 即可釋放已佔用的訂閲額度。

**第二步：取得數據**  

如何將訂閲推送的數據從 OpenD 取回程式呢？我們提供瞭如下兩種方式：

**方式 1：實時數據回調**  
設置相應的回調函數，來異步處理 OpenD 收到的數據推送。  

設置好回調函數後，OpenD 會將收到的實時數據，立即推給程式的回調函數進行處理。  

如果所訂閲的標的比較活躍，此時的推送數據可能數據量較大且頻率較高。如果您希望適當降低 OpenD 給程式的推送頻率，建議在 [OpenD 啟動參數](../opend/opend-cmd.md#1028) 中設定 API 推送頻率（`qot_push_frequency`）。  

方式 1 涉及的接口包括：[實時報價回調](../quote/update-stock-quote.md)、[實時掛盤迴調](../quote/update-order-book.md)、[實時 K 線回調](../quote/update-kl.md)、[實時分時回調](../quote/update-rt.md)、[實時逐筆回調](../quote/update-ticker.md)、[實時經紀隊列回調](../quote/update-broker.md)。

**方式 2：獲取實時數據**  
通過獲取實時數據接口，可以將 OpenD 收到的最新的數據，取回程式。這種方式更加靈活，程式不需要處理海量的推送。只要 OpenD 在持續接收富途服務器的推送，程式可以隨用隨取，不用不取。  

由於是從 OpenD 接收的推送數據中取，所以這類接口沒有頻率限制。  

方式 2 涉及的接口包括：[獲取實時報價](../quote/get-stock-quote.md)、[獲取實時掛盤](../quote/get-order-book.md)、[獲取實時 K 線](../quote/get-kl.md)、[獲取實時分時](../quote/get-rt.md)、[獲取實時逐筆](../quote/get-ticker.md)、[獲取實時經紀隊列](../quote/get-broker.md)。

## Q14：各個市場狀態對應什麼時間段？
A: 
<table>
    <tr>
        <th>市場</th>
        <th>品類</th>
        <th>市場狀態</th>
        <th>時間段（當地時間）</th>
    </tr>
    <tr>
        <td rowspan="19" width = "15%">香港市場</td>
	    <td rowspan="8" width = "15%">證券類產品（含股票、ETFs、窩輪、牛熊、界內證）</td>
	    <td> * NONE：無交易</td>
      <td> CST 08:55 - 09:00</td>
    </tr>
    <tr>
	    <td >* AUCTION：盤前競價</td>
      <td> CST 09:00 - 09:20</td>
    </tr>
    <tr>
	    <td >* WAITING_OPEN：等待開市</td>
      <td> CST 09:20 - 09:30</td>
    </tr>
    <tr>
	    <td>* MORNING：早盤</td>
      <td> CST 09:30 - 12:00</td>
    </tr>
    <tr>
      <td>* REST: 午間休市</td>
	    <td>CST 12:00 - 13:00</td>
    </tr>
    <tr>
	    <td>* AFTERNOON：午盤</td>
      <td>CST 13:00 - 16:00</td>
    </tr>
    <tr>
	    <td>* HK_CAS：港股盤後競價（港股市場增加 CAS 機制對應的市場狀態）</td>
      <td>CST 16:00 - 16:08</td>
    </tr>
    <tr>
	    <td>* CLOSED：收市</td>
      <td>CST 16:08 - 08:55（T+1）</td>
    </tr>
    <tr>
	    <td rowspan="5">期權、期貨（僅日市）</td>
      <td>* NONE：期權待開市</td>
      <td> CST 08:55 - 09:30</td>
    </tr>
    <tr>
	    <td>* MORNING：早盤</td>
      <td>CST 09:30 - 12:00</td>
    </tr>
    <tr>
      <td>* REST: 午間休市</td>
	    <td>CST 12:00 - 13:00</td>
    </tr>
    <tr>
	    <td>* AFTERNOON：午盤</td>
      <td>CST 13:00 - 16:00</td>
    </tr>
    <tr>
	    <td>* CLOSED：收市</td>
      <td>CST 16:00 - 08:55（T+1）</td>
    </tr>
    <tr>
	    <td rowspan="6">期貨（日夜市）</td>
      <td>* FUTURE_DAY_WAIT_FOR_OPEN：期貨待開市</td>
      <td rowspan="6"> 不同品種交易時間不同</td>
    </tr>
    <tr>
	    <td>* NIGHT_OPEN: 夜市交易時段</td>
    </tr>
    <tr>
	    <td>* NIGHT_END：夜市收市</td>
    </tr>
    <tr>
	    <td>* FUTURE_DAY_WAIT_FOR_OPEN：期貨待開市</td>
    </tr>
    <tr>
	    <td>* FUTURE_DAY_OPEN：日市交易時段</td>
    </tr>
    <tr>
	    <td>* FUTURE_DAY_CLOSE：日市收市</td>
    </tr>
  <tr>
        <td rowspan="16">美國市場</td>
	    <td rowspan="5">證券類產品（含股票、ETFs）</td>
	    <td>* PRE_MARKET_BEGIN：美股盤前交易時段</td>
      <td>EST 04:00 - 09:30</td>
    </tr>
    <tr>
	    <td>* AFTERNOON：美股持續交易時段</td>
      <td>EST 09:30 - 16:00</td>
    </tr>
    <tr>
	    <td>* AFTER_HOURS_BEGIN：美股盤後交易時段</td>
      <td>EST 16:00 - 20:00</td>
    </tr>
    <tr>
	    <td>* AFTER_HOURS_END：美股盤後收市</td>
      <td>EST 20:00 - 04:00（T+1）</td>
    </tr>
    <tr>
	    <td>* OVERNIGHT：美股夜盤交易時段</td>
      <td>EST 20:00 - 04:00（T+1）</td>
    </tr>
    <tr>
	    <td rowspan="6">期權</td>
      <td>* NONE：期權待開市</td>
      <td rowspan="6"> 不同品種交易時間不同</td>
    </tr>
    <tr>
	    <td>* REST：美指期權午間休市</td>
    </tr>
    <tr>
	    <td>* AFTERNOON：美股持續交易時段</td>
    </tr>
    <tr>
	    <td>* TRADE_AT_LAST：美指期權盤尾交易時段</td>
    </tr>
    <tr>
	    <td>* NIGHT：美指期權夜市交易時段</td>
    </tr>
    <tr>
	    <td>* CLOSED：收市</td>
    </tr>
    <tr>
	    <td rowspan="5">期貨</td>
      <td>* FUTURE_SWITCH_DATE：美期待開市</td>
      <td rowspan="5"> 不同品種交易時間不同</td>
    </tr>
    <tr>
	    <td>* FUTURE_OPEN：美期交易時段</td>
     </tr>
     <tr>
	    <td>* FUTURE_BREAK：美期中盤休息</td>
     </tr>
     <tr>
	    <td>* FUTRUE_BREAK_OVER：美期休息後交易時段</td>
     </tr>
     <tr>
	    <td>* FUTURE_CLOSE：美期收市</td>
     </tr>
    <tr>
        <td rowspan="7">A股市場</td>
	    <td rowspan="7">證券類產品（含股票、ETFs）</td>
	    <td>* NONE：無交易</td>
      <td>CST 08:55 - 09:15</td>
    </tr>
    <tr>
	    <td>* Auction：盤前競價</td>
      <td>CST 09:15 - 09:25</td>
    </tr>
    <tr>
	    <td>* WAITING_OPEN：等待開市</td>
      <td> CST 09:25 - 09:30</td>
    </tr>
    <tr>
	    <td>* MORNING：早盤</td>
      <td>CST 09:30 - 11:30</td>
    </tr>
    <tr>
	    <td>* REST：午間休市</td>
      <td>CST 11:30 - 13:00</td>
    </tr>
    <tr>
	    <td>* AFTERNOON：午盤</td>
      <td>CST 13:00 - 15:00</td>
    </tr>
    <tr>
	    <td>* CLOSED：收市</td>
      <td>CST 15:00 - 08:55（T+1）</td>
    </tr>
    <tr>
        <td rowspan="5">新加坡市場</td>
	    <td rowspan="5">期貨</td>
	    <td>* FUTURE_DAY_WAIT_FOR_OPEN：期貨待開市</td>
      <td rowspan="5">不同品種交易時間不同</td>
    </tr>
     <tr>
	    <td>* NIGHT_OPEN：夜市交易時段</td>
    </tr>
     <tr>
	    <td>* NIGHT_END：夜市收市</td>
    </tr>
     <tr>
	    <td>* FUTURE_DAY_OPEN：日市交易時段</td>
    </tr>
     <tr>
	    <td>* FUTURE_DAY_CLOSE：日市收市</td>
    </tr>
    <tr>
        <td rowspan="5">日本市場</td>
	    <td rowspan="5">期貨</td>
	    <td>* FUTURE_DAY_WAIT_FOR_OPEN：期貨待開市</td>
      <td>JST 16:25（T-1）- 16:30（T-1）</td>
    </tr>
     <tr>
	    <td>* NIGHT_OPEN：夜市交易時段</td>
      <td>JST 16:30（T-1） - 05:30</td>
    </tr>
     <tr>
	    <td>* NIGHT_END：夜市收市</td>
      <td>JST 05:30 - 08:45</td>
    </tr>
     <tr>
	    <td>* FUTURE_DAY_OPEN：日市交易時段</td>
      <td>JST 08:45 - 15:15</td>
    </tr>
     <tr>
	    <td>* FUTURE_DAY_CLOSE：日市收市</td>
      <td>JST 15:15 - 16:25</td>
    </tr>
</table>
\* CST, EST, JST 分別表示中國時間，美東時間，日本時間

## Q15：接口參數股票代碼的格式

A：  
* 使用不同程式語言的用戶，需要的股票代碼的格式不同：
   * **Python 用戶**  
    標的代碼code 使用`exchange_market.symbol`格式,`exchange_market`表示交易所市場，`symbol`表示標的代碼。支援訂閱的標的如下：    

<table>
    <tr>
        <th>市場</th>
        <th>標的類別</th>
        <th>exchange_market</th>
        <th>example</th>
    </tr>
    <tr>
        <td rowspan="5">香港市場</td>
        <td>證券類產品（含股票、ETFs、窩輪、牛熊、界內證）</td>
        <td>HK</td>
        <td>騰訊控股：HK.00700</td>
    </tr>
    <tr>
        <td>指數</td>
        <td>HK</td>
        <td>恒生指數：HK.800000</td>
    </tr>  
    <tr>
        <td>期貨</td>
        <td>HK</td>
        <td>恒指期貨2606：HK.HSI2606</td>
    </tr>
    <tr>
        <td>期權</td>
        <td>HK</td>
        <td>* 股票期權 騰訊 260330 450.00購：HK.TCH260330C450000 <br> * 指數期权 恒指 260330 24000.00購：HK.HSI260330C24000000</td>
    </tr>
    <tr>
        <td>板塊  (建議使用[get_plate_list](../quote/get-plate-list.html) 先取得板塊列表) </td>
        <td>HK</td>
        <td>AI應用股：HK.LIST24037</td>
    </tr>
    <tr>
        <td rowspan="5">美國市場</td>
        <td>證券類產品（含紐交所、美交所、納斯達克上市的股票、ETFs）</td>
        <td>US</td>
        <td>英偉達：US.NVDA</td>
    </tr>
    <tr>
        <td>期權</td>
        <td>US</td>
        <td>* 股票期權 NVDA 260330 160.00C：US.NVDA260330C160000 <br> * 指數期權 SPXW 260330 6330.00C: US..SPXW260330C6330000</td>
    </tr>
    <tr>
        <td>期貨</td>
        <td>US</td>
        <td>標普500指數期貨2606：US.ES2606</td>
    </tr>
    <tr>
        <td>板塊  (建議使用[get_plate_list](../quote/get-plate-list.html) 先取得板塊列表) </td>
        <td>US</td>
        <td>半導體精選：US.LIST20077</td>
    </tr>
    <tr>
        <td>指數（暫不支援取得）</td>
        <td>US</td>
        <td>標普500指數：US..SPX</td>
    </tr>
    <tr>
        <td rowspan="3">A 股市場</td>
        <td>證券類產品（含股票、ETFs）</td>
        <td>SH/SZ</td>
        <td>貴州茅台：SH.600519</td>
    </tr>
    <tr>
        <td>指數</td>
        <td>SH/SZ</td>
        <td>上證指數：SH.000001</td>
    </tr>
    <tr>
        <td>板塊   (建議使用[get_plate_list](../quote/get-plate-list.html) 先取得板塊列表) </td>
        <td>SH/SZ</td>
        <td>汽車電子概念：SH.LIST0301</td>
    </tr>
    <tr>
        <td rowspan="1">新加坡市場（暫不支援取得）</td>
        <td>期貨</td>
        <td>SG</td>
        <td>A50指數期貨2606：SG.CN2606</td>
    </tr>
    <tr>
        <td rowspan="1">日本市場（暫不支援取得）</td>
        <td>期貨</td>
        <td>JP</td>
        <td>大阪日經指數期貨2606：JP.NK2252606</td>
    </tr>
    </table>

  

   * **非 Python 用戶**   
    股票結構參見 [Security](../quote/quote.html#3103)。   
    例如：騰訊控股，參數 market 傳入 QotMarket_HK_Security，參數 code 傳入'00700'。

* 查詢方式：  
   通過 APP 查看代碼和行情市場：行情 > 自選 > 全部。  
   行情市場定義，請參考 [這裏](../quote/quote.html#8744)。  
    ![code](../img/code.png)    


## Q16：復權因子相關
A：  
### 概述
所謂 [復權](../quote/get-rehab.html#4123) 就是對股價和成交量進行權息修復，按照股票的實際漲跌繪製股價走勢圖，並把成交量調整為相同的股本口徑。  
公司行動（如：拆股、合股、送股、轉增股、配股、增發股、分紅）均可能對股價產生影響，而復權計算可對量價進行調整，剔除公司行動的影響，保持股價走勢的連續性。   

### 名詞解釋
- 公司行動：上市公司進行一些股權、股票等影響公司股價和股東持倉變化的行為。
- 前復權：保持現有的股價不變，以當前的股價為基準，對以前的股價進行復權計算。
- 後復權：保持先前的股價不變，以過去的股價為基準，對以後的股價進行復權計算。
- 復權因子：即權息修複比例，用於計算復權後的價格及持倉數量。
- 除權除息日：即股權登記日下一個交易日。在股票的除權除息日，證券交易所都要計算出股票的除權除息價，以作為股民在除權除息日開市的參考。其意義是股票股利分配給股東的日期。

### 復權方法
主流的復權計算方法分為兩種：事件法和連乘法；而 Futu API 針對不同市場使用不同的計算方法。
- 事件復權法：通過還原除權除息的各類事件進行復權；存在兩個復權因子（復權因子 A 和 復權因子 B），復權因子 B 主要調整現金分紅對股價的影響，而復權因子 A 調整其他公司行動對股價的影響。
- 連乘復權法：通過復權因子連乘的方式進行復權，只保留 復權因子 A（或將 復權因子 B 置為0），復權因子 A 為 除權除息日前收市價/該日經權息調整後的前收市價。

::: tip 提示
*  API 對美股前復權使用連乘法，即將 復權因子 B 置為0。  
*  API 對除美股以外的標的（A股、港股、新加坡股票等）及美股後復權使用事件法。  
:::

### 計算公式
#### 單次復權
- 前復權：  
前復權價格 = 不復權價格 × 前復權因子 A + 前復權因子 B   
- 後復權：  
後復權價格 = 不復權價格 × 後復權因子 A + 後復權因子 B

#### 多次復權
- 前復權：按照時間順序，篩選出大於計算日期的復權因子，優先使用時間較早的復權因子進行復權計算。以兩次復權為例： 

  ![code](../img/forward_fomula.png)    
- 後復權：按照時間倒序，篩選出小於等於計算日期的復權因子，優先使用時間較晚的復權因子進行復權計算。以兩次復權為例： 

  ![code](../img/backward_fomula.png)    

### 示例
#### 單次前復權示例
以牧原股份為例：
- 篩選復權因子如下：  

除權除息日|股票代碼|方案説明|前復權因子 A |前復權因子 B 
:-|:-|:-|:-|:-
2021/06/03|SZ.002714|10轉4.0股派14.61元（含税）|0.71429|-1.04357

- 不復權數據如下：  

日期|股票代碼|不復權收市價
:-|:-|:-
2021/06/02|SZ.002714|93.11
2021/06/03|SZ.002714|66.25

- 前復權數據如下：  

日期|股票代碼|前復權收市價
:-|:-|:-
2021/06/02|SZ.002714|65.4639719
2021/06/03|SZ.002714|66.25

- 前復權數據計算方法：  
牧原股份在 2021/06/03 進行拆股及現金分紅行動（10轉4.0股派14.61元），根據前復權計算公式對 2021/06/02 的收市價進行調整計算，則：前復權價格（65.4639719） = 不復權價格（93.11） × 前復權因子 A（0.71429） + 前復權因子 B（-1.04357）   

  ![code](../img/forward_example.png)    

#### 多次後復權示例
接上一個例子，計算牧原股份在 2021/06/02 的後復權價格：
- 篩選復權因子如下：  

除權除息日|股票代碼|方案説明|後復權因子 A |後復權因子 B 
:-|:-|:-|:-|:-|
2014/07/04|SZ.002714|10派2.34元（含税）|1|0.234
2015-06-10|SZ.002714|10轉10.0股派0.61元（含税）|2|0.061
2016-07-08|SZ.002714|10轉10.0股派3.53元（含税）|2|0.353
2017-07-11|SZ.002714|10轉8.0股派6.9元（含税）|1.8|0.69
2018-07-03|SZ.002714|10派6.91元（含税）|1|0.691
2019-07-04|SZ.002714|10派0.5元（含税）|1|0.05
2020-06-04|SZ.002714|10轉7.0股派5.5元（含税）|1.7|0.55

- 不復權數據如下：  

日期|股票代碼|不復權收市價
:-|:-|:-
2021/06/02|SZ.002714|93.11

- 後復權數據如下：  

日期|股票代碼|後復權收市價
:-|:-|:-
2021/06/02|SZ.002714|1152.7226

- 後復權數據計算方法：  
為了計算牧原股份在 2021/06/02 的後復權價格，需要將早於 2021/06/02 的復權事件進行一一復權，得到最後的後復權價格，具體計算如下：

  ![code](../img/backward_example.jpg)

---



---

# 交易相關

## Q1：模擬交易相關

A:
### 概述
模擬交易是在真實的市場環境中，用虛擬資金做交易，不會對您的真實帳戶的資產造成影響。

#### 交易時間
模擬交易僅支持在常規交易時段交易，支持美股盤中交易時段、美股盤前盤後時段，不支持美股夜盤、全時段交易和A股港股盤前盤後競價時段交易。詳情可點擊 [模擬交易規則](https://support.futunn.com/topic692)。

#### 支持品類
Futu API 支持模擬交易的品類請參考 [這裏](../intro/intro.md#1396)。

#### 解鎖
與真實交易不同，模擬交易無需對帳戶進行解鎖，即可下單或修改訂單取消訂單。


#### 訂單
1. 訂單類型：限價單和市價單。  
2. 修改訂單操作類型：模擬交易不支持使生效、使失效、刪除，僅支持修改訂單、 取消訂單。  
3. 成交：模擬交易不支持成交相關操作，包括 [查詢今日成交](../trade/get-order-fill-list.md#7660)、[查詢歷史成交](../trade/get-history-order-fill-list.md#2782)、[響應成交推送回調](../trade/update-order-fill.md#7852)。
4. 有效期限：模擬交易有效期限僅支持當日有效。
5. 賣空：期權和期貨支持賣空。股票僅美股支持賣空。
6. 模擬交易帳戶不支持查詢訂單費用。
7. 模擬交易帳戶不支持查詢現金流水。
8. 在組合期權訂單場景下，支持持倉查詢，暫不支持組合訂單查詢。


#### 操作平台
1. 手機版：我的 — 模擬交易  

![sim-page](../img/sim-page.png)

2. 桌面版：左側模擬 tab  

![sim-page](../img/create-sim-account.png)


3. 網頁版：[模擬交易界面](https://m-match.futunn.com/simulate/)

4. Futu API：在調用接口時，設置參數交易環境為模擬環境即可。詳見 [如何使用 Futu API 進行模擬交易](../qa/trade.md#3142)。

::: tip 提示
* 以上四種方式只是操作平台不同，四種方式操作的模擬帳戶是共通的。  
:::


### 如何使用 Futu API 進行模擬交易？

#### 創建連接
先根據交易品種 [創建相應的連接](../trade/base.md#9291) 。當交易品種是股票或期權時，請使用 `OpenSecTradeContext`。當交易品種是期貨時，請使用 `OpenFutureTradeContext`。

#### 獲取交易業務帳戶列表
使用 [獲取交易業務帳戶列表](../trade/get-acc-list.md#4991) 查看交易帳戶（包括模擬帳戶、真實帳戶）。以 Python 為例：返回欄位交易環境 `trd_env` 為 `SIMULATE`，表示模擬帳戶。   
獲取港股模擬交易帳戶，需要指定 filter_trdmarket 爲 TrdMarket.HK，此時會返回2個模擬交易賬號。其中 sim_acc_type = STOCK 爲港股模擬帳戶，sim_acc_type = OPTION 爲港股期權模擬帳戶，sim_acc_type = FUTURES 爲港股期貨模擬帳戶。    
獲取美股模擬交易帳戶，需要指定 filter_trdmarket 爲 TrdMarket.US，sim_acc_type = STOCK_AND_OPTION 代表美股融資融券模擬帳戶，可以模擬交易股票和期權。sim_acc_type = FUTURES 爲美國期貨模擬帳戶。   

* **Example：Stocks and Options**
```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
#trd_ctx = OpenFutureTradeContext(host='127.0.0.1', port=11111, is_encrypt=None, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_acc_list()
if ret == RET_OK:
    print(data)
    print(data['acc_id'][0])  # get the first account id
    print(data['acc_id'].values.tolist())  # convert to list format
else:
    print('get_acc_list error: ', data)
trd_ctx.close()
```

* **Output**
```python
               acc_id   trd_env acc_type          card_num   security_firm  \
0  281756480572583411      REAL   MARGIN  1001318721909873  FUTUSECURITIES   
1             9053218  SIMULATE     CASH               N/A             N/A   
2             9048221  SIMULATE   MARGIN               N/A             N/A   

  sim_acc_type  trdmarket_auth  
0          N/A  [HK, US, HKCC]  
1        STOCK            [HK]  
2       OPTION            [HK] 
```
::: tip 提示
* 模擬交易中，區分股票帳戶和期權帳戶，股票帳戶只能交易股票，期權帳戶只能交易期權；以 Python 為例：返回欄位中模擬帳戶類型 `sim_acc_type` 為 `STOCK`，表示股票帳戶；為`OPTION`，表示期權帳戶。
::: 

* **Example: Futures**
```python
from futu import *
#trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
trd_ctx = OpenFutureTradeContext(host='127.0.0.1', port=11111, is_encrypt=None, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_acc_list()
if ret == RET_OK:
    print(data)
    print(data['acc_id'][0])  # get the first account id
    print(data['acc_id'].values.tolist())  # convert to list format
else:
    print('get_acc_list error: ', data)
trd_ctx.close()
```

* **Output**
```python
    acc_id   trd_env acc_type card_num security_firm sim_acc_type  \
0  9497808  SIMULATE   MARGIN      N/A           N/A      FUTURES   
1  9497809  SIMULATE   MARGIN      N/A           N/A      FUTURES   
2  9497810  SIMULATE   MARGIN      N/A           N/A      FUTURES   
3  9497811  SIMULATE   MARGIN      N/A           N/A      FUTURES   

          trdmarket_auth  
0  [FUTURES_SIMULATE_HK]  
1  [FUTURES_SIMULATE_US]  
2  [FUTURES_SIMULATE_SG]  
3  [FUTURES_SIMULATE_JP]  
```  

#### 下單
使用 [下單接口](../trade/place-order.md) 時，設置交易環境為模擬環境即可。以 Python 為例：`trd_env = TrdEnv.SIMULATE`。

* **Example**
```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.place_order(price=510.0, qty=100, code="HK.00700", trd_side=TrdSide.BUY, trd_env=TrdEnv.SIMULATE)
if ret == RET_OK:
    print(data)
else:
    print('place_order error: ', data)
trd_ctx.close()
```
* **Output**
```python
	code	stock_name	trd_side	order_type	order_status	order_id	qty	price	create_time	updated_time	dealt_qty	dealt_avg_price	last_err_msg	remark	time_in_force	fill_outside_rth
0	HK.00700	騰訊控股	BUY	NORMAL	SUBMITTING	4642000476506964749	100.0	510.0	2021-10-09 11:34:54	2021-10-09 11:34:54	0.0	0.0			DAY	N/A
```

#### 取消訂單修改訂單
使用 [取消訂單接口](../trade/modify-order.md) 時，設置交易環境為模擬環境即可。以 Python 為例： `trd_env = TrdEnv.SIMULATE`。

* **Example**
```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
order_id = "4642000476506964749"
ret, data = trd_ctx.modify_order(ModifyOrderOp.CANCEL, order_id, 0, 0, trd_env=TrdEnv.SIMULATE)
if ret == RET_OK:
    print(data)
else:
    print('modify_order error: ', data)
trd_ctx.close()
```
* **Output**
```python
    trd_env             order_id
0  SIMULATE  4642000476506964749
```

#### 查詢歷史訂單
使用 [查詢歷史訂單接口](../trade/get-history-order-list.md) 時，設置交易環境為模擬環境即可。以 Python 為例：`trd_env = TrdEnv.SIMULATE`。

* **Example**
```python
from futu import *
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK, host='127.0.0.1', port=11111, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.history_order_list_query(trd_env=TrdEnv.SIMULATE)
if ret == RET_OK:
    print(data)
else:
    print('history_order_list_query error: ', data)
trd_ctx.close()
```
* **Output**
```python
	code	stock_name	trd_side	order_type	order_status	order_id	qty	price	create_time	updated_time	dealt_qty	dealt_avg_price	last_err_msg	remark	time_in_force	fill_outside_rth
0	HK.00700	騰訊控股	BUY	ABSOLUTE_LIMIT	CANCELLED_ALL	4642000476506964749	100.0	510.0	2021-10-09 11:34:54	2021-10-09 11:37:08	0.0	0.0			DAY	N/A
```

### 如何重置模擬帳戶？
目前 Futu API 不支持重置模擬帳戶，您可在手機版使用復活卡重置指定模擬帳戶，重置後帳戶資金將恢復至初始值，歷史訂單將會被清空。

#### 具體操作
手機版：我的 — 模擬交易 — 我的頭像 — 我的道具 — 復活卡。
![sim-page](../img/sim-reset.png)


## Q2：是否支持 A 股交易？

A: 模擬交易支持 A 股交易。但真實交易僅可通過 A 股通交易部分 A 股，具體詳見 [A 股通名單](https://www.hkex.com.hk/Mutual-Market/Stock-Connect/Eligible-Stocks/View-All-Eligible-Securities?sc_lang=zh-HK)。

## Q3：各市場支持的交易方向

A: 除了期貨，其他股票都只支持傳入 BUY 和 SELL 兩個交易方向。在空倉情況下傳入 SELL，產生的訂單交易方向是賣空。

## Q4：真實交易中，各市場支持的訂單類型

A: 
<table style="font-size:14px;">
    <tr>
        <th>市場</th>
        <th>品種</th>
        <th>限價單</th>
        <th>市價單</th>
        <th>競價限價單</th>
        <th>競價市價單</th>
        <th>絕對限價單</th>
        <th>特別限價單</th>
        <th>特別限價且要求<br/>全部成交訂單</th>
        <th>止損市價單</th>
        <th>止損限價單</th>
        <th>觸及市價單（止盈）</th>
        <th>觸及限價單（止盈）</th>
        <th>跟蹤止損市價單</th>
        <th>跟蹤止損限價單</th>
    </tr>
    <tr>
        <td rowspan="3">香港市場</td>
        <td>證券類產品（含股票、ETFs、<br/>窩輪、牛熊、界內證）</td>
        <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td>期權</td>
        <td>✓</td> <td>X</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>X</td> <td>✓</td> <td>X</td> <td>✓</td> <td>X</td> <td>✓</td>
    </tr>
    <tr>
        <td>期貨</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td rowspan="3">美國市場</td>
        <td>證券類產品（含股票、ETFs）</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td>期權</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td>期貨</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td>A 股通市場</td>
        <td>證券類產品（含股票、ETFs）</td>
        <td>✓</td> <td>X</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>X</td> <td>✓</td> <td>X</td> <td>✓</td> <td>X</td> <td>✓</td>
    </tr>
    <tr>
        <td>新加坡市場</td>
        <td>期貨</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
    <tr>
        <td>日本市場</td>
        <td>期貨</td>
        <td>✓</td> <td>✓</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>-</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td> <td>✓</td>
    </tr>
</table>


## Q5：各市場支持的訂單操作

A: 
* 港股支持修改訂單、取消訂單、生效、失效、刪除
* 美股僅支持修改訂單和取消訂單
* A 股通僅支持取消訂單
* 期貨支持修改訂單、取消訂單、刪除

## Q6：OpenD 啟動參數 future_trade_api_time_zone 如何使用？

A：由於期貨帳戶支持交易的品種分佈在全球多個交易所，交易所的所屬時區各有不同，因此期貨交易 API 的時間顯示就成為了一個問題。  
OpenD 啟動參數中新增了 future_trade_api_time_zone 這一參數，供全球不同地區的期貨交易者靈活指定時區。預設時區為 UTC+8，如果您更習慣美東時間，只需將此參數設定為 UTC-5 即可。
::: tip  提示
+ 此參數僅會對期貨交易接口類對象生效。港股交易、美股交易、A 股通交易接口類對象的時區，仍然按照交易所所在的時區進行顯示。
+ 此參數會影響的接口包括：響應訂單推送回調，響應成交推送回調，查詢今日訂單，查詢歷史訂單，查詢當日成交，查詢歷史成交，下單。
:::

## Q7：通過 API 下的訂單，能在 APP 上面看到嗎？
A：可以看到。  
通過 Futu API 成功發出下單指示後，您可以在 APP 的 **交易** 頁面，查看今日訂單、訂單狀態、成交情況等等，也可以在 **消息—訂單消息** 中收到成交提醒的通知。

## Q8：哪些品類支持在非交易時段下單？
A：所有的訂單，都需要在開市期間才能夠成交。  
Futu API 僅對一部分品類，支持了 **非交易時段下單** 的功能（APP 上支持更多品類的非交易時段下單功能）。具體請參考下表：
<table>
    <tr>
        <th rowspan="2">市場</th>
        <th rowspan="2">標的類型</th>
        <th rowspan="2">模擬交易</th>
        <th colspan="7">真實交易</th>
    </tr>
    <tr>
        <th>Futu HK</th>
        <th>Moomoo US</th>
        <th>Moomoo SG</th>
        <th>Moomoo AU</th>
        <th>Moomoo MY</th>
        <th>Moomoo CA</th>
        <th>Moomoo JP</th>
    </tr>
    <tr>
        <td rowspan="3">香港市場</td>
	    <td>股票、ETFs、窩輪、牛熊、界內證</td>
	    <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
   <tr>
	    <td>期權 (含指數期權，需使用期貨帳戶交易)</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="3">美國市場</td>
	    <td>股票、ETFs</td>
	    <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
    </tr>
    <tr>
        <td>期權</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
    </tr>
   <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td rowspan="2">A 股市場</td>
	    <td>A 股通股票</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
     <tr>
	    <td>非 A 股通股票</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
   <tr>
        <td rowspan="2">新加坡市場</td>
	    <td>股票、ETFs、窩輪、REITs、DLCs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="2">日本市場</td>
        <td>股票、ETFs、REITs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
        <td>期貨</td>
        <td align="center">✓</td>
        <td align="center">✓</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="1">澳大利亞市場</td>
        <td>股票、ETFs</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
    <tr>
	    <td rowspan="1">加拿大市場</td>
        <td>股票</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
        <td align="center">X</td>
    </tr>
</table>
::: tip 提示
- ✓：支持非交易時段下單
- X：暫不支持非交易時段下單（或暫不支持交易）
:::

## Q9：對於下單接口，各訂單類型對應的必要參數以及券商對單筆訂單的下單限制
A1: 各訂單類型對應的必要參數

<table style="font-size:14px;">
    <tr>
        <th>參數</th>
        <th>限價單</th>
        <th>市價單</th>
        <th>競價限價單</th>
        <th>競價市價單</th>
        <th>絕對限價單</th>
        <th>特別限價單</th>
        <th>特別限價且要求<br/>全部成交訂單</th>
        <th>止損市價單</th>
        <th>止損限價單</th>
        <th>觸及市價單（止盈）</th>
        <th>觸及限價單（止盈）</th>
        <th>跟蹤止損市價單</th>
        <th>跟蹤止損限價單</th>
    </tr>
    <tr>
        <td>price</td>
        <td>✓</td> <td></td> <td>✓</td> <td> </td> <td>✓</td> <td>✓</td> <td>✓</td>  <td></td><td>✓</td> <td></td> <td>✓</td><td> </td><td> </td>
    </tr>
    <tr>
        <td>qty</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>code</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trd_side</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>order_type</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trd_env</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>aux_price</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td> </td><td> </td>
    </tr>
    <tr>
        <td>trail_type</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trail_value</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trail_spread</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td> </td><td>✓</td>
    </tr>
</table>

`Python 用户` 注意，[place_order](../trade/place-order.html#8194) 並未對 price 設置預設值，對於上述五類訂單類型，仍需對 price 傳參，price 可以傳入任意值。

A2：各券商對單筆訂單的股數及金額限制
<table style="font-size:14px;">
    <tr>
        <th>券商</th>
        <th>品類</th>
        <th>單筆訂單股數上限</th>
        <th>單筆訂單金額上限</th>
    </tr>
    <tr>
        <td rowspan="3">FUTU HK</td>
        <td>A股通</td>
        <td>1,000,000 股</td>
        <td>￥5,000,000</td>
    </tr>
    <tr>
        <td>美股</td>
        <td>500,000 股</td>
        <td>$5,000,000</td>
    </tr>
    <tr>
        <td>香港股票期貨/期權</td>
        <td>3,000 手</td>
        <td>無限制</td>
    </tr>
    <tr>
        <td>moomoo US</td>
        <td>美股</td>
        <td>500,000 股</td>
        <td>$10,000,000</td>
    </tr>
    <tr>
        <td>moomoo SG</td>
        <td>美股</td>
        <td>500,000 股</td>
        <td>$5,000,000</td>
    </tr>
    <tr>
        <td>moomoo AU</td>
        <td>美股</td>
        <td>無限制</td>
        <td>無限制</td>
    </tr>
</table>


## Q10：對於修改訂單接口，修改訂單時，各訂單類型對應的必要參數
A: 

<table style="font-size:14px;">
    <tr>
        <th>參數</th>
        <th>限價單</th>
        <th>市價單</th>
        <th>競價限價單</th>
        <th>競價市價單</th>
        <th>絕對限價單</th>
        <th>特別限價單</th>
        <th>特別限價且要求<br/>全部成交訂單</th>
        <th>止損市價單</th>
        <th>止損限價單</th>
        <th>觸及市價單（止盈）</th>
        <th>觸及限價單（止盈）</th>
        <th>跟蹤止損市價單</th>
        <th>跟蹤止損限價單</th>
    </tr>
    <tr>
        <td>modify_order_op</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>order_id</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>price</td>
        <td>✓</td> <td></td> <td>✓</td> <td> </td> <td>✓</td> <td>✓</td> <td>✓</td>  <td></td><td>✓</td> <td></td> <td>✓</td><td> </td><td> </td>
    </tr>
    <tr>
        <td>qty</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trd_env</td>
        <td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>aux_price</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td>✓</td><td>✓</td><td>✓</td><td>✓</td> <td> </td><td> </td>
    </tr>
    <tr>
        <td>trail_type</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trail_value</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td>✓</td><td>✓</td>
    </tr>
    <tr>
        <td>trail_spread</td>
        <td></td> <td></td> <td></td> <td></td> <td></td> <td></td> <td> </td><td> </td><td> </td><td> </td><td> </td> <td> </td><td>✓</td>
    </tr>
</table>

`Python 用户` 注意，[modify_order](../trade/modify-order.html#5781) 並未對 price 設置預設值，對於上述五類訂單類型，仍需對 price 傳參，price 可以傳入任意值。

## Q11：交易接口返回“當前證券業務帳戶尚未同意免責協議”？
A：  
點擊下方連結完成協議確認，重啟 OpenD 即可正常使用交易功能。
所屬券商|協議確認
:-|:-|:-
FUTU HK|[點擊這裏](https://risk-disclosure.futuhk.com/index?agreementNo=HKOT0015)
Moomoo US|[點擊這裏](https://risk-disclosure.us.moomoo.com/index?agreementNo=USOT0027)
Moomoo SG|[點擊這裏](https://risk-disclosure.sg.moomoo.com/index?agreementNo=SGOT0015)
Moomoo AU|[點擊這裏](https://risk-disclosure.au.moomoo.com/index?agreementNo=AUOT0025)
Moomoo CA|[點擊這裏](https://risk-disclosure.ca.moomoo.com/index?agreementNo=CAOT0117)
Moomoo MY|[點擊這裏](https://risk-disclosure.my.moomoo.com/index?agreementNo=MYOT0066)
Moomoo JP|[點擊這裏](https://risk-disclosure.jp.moomoo.com/index?agreementNo=JPOT0140)


## Q12：典型日內交易者（PDT）相關

### 概述

客户使用moomoo證券(美國) 帳戶進行日內交易時，會受到美國 FINRA 的監管限制（此為美國券商受到的監管要求，與交易股票的所屬市場無關。其他國家或地區的券商  (如：富途證券(香港)、moomoo證券(新加坡)) 的交易帳戶則不受此限制）。若用户在任意連續的5個交易日內，進行日內交易 3 次以上，則會被標記為典型日內交易者（PDT）。  
更多詳情，[點擊這裏](https://fastsupport.fututrade.com/hans/category11014/scid11017)

### 進行日內交易的流程圖
![PDT_process](../img/PDT_process.png) 

### 我願意被標記為 PDT，且不希望程式交易被打斷，如何關閉“防止被標記為 PDT”？
A：  
當您在連續的 5 個交易日內，進行第 4 次日內交易時，為了防止您被無意識地標記為 PDT，服務器會對此交易進行攔截。若您主動想被標記為 PDT，並且不希望服務器攔截，可以採取以下措施：  
在 [命令行 OpenD 中設定參數](../opend/opend-cmd.html#1028)，將啟動參數 `pdt_protection` 的值修改為 0，以關閉“防止被標記為日內交易者”的功能。

![US_para](../img/US_para.png)  
注意：若您被標記 PDT，當您的帳戶權益小於$25000時，您將無法開倉。

### 如何關閉 DTCall 預警提醒？
A：  
您被標記為 PDT 後，需要留意帳戶的日內交易購買力（DTBP），日內交易超出 DTBP 時將收到日內交易保證金追繳（DTCall）。服務器會在您即將開倉下單超出剩餘日內交易購買力前，阻止您的下單。若您仍然希望進行下單，並且不希望服務器攔截，可以採取以下措施：    
在 [命令行 OpenD 中設定參數](../opend/opend-cmd.html#1028)，將啟動參數 `dtcall_confirmation` 的值修改為 0，以關閉“日內交易保證金追繳預警”的功能。

![US_para2](../img/US_para2.png)  
注意：若您開倉訂單的市值大於您的剩餘日內交易購買力，並且在今日平倉當前標的，您將會收到日內交易保證金追繳通知（Day-Trading Call），只能通過存入資金才能解除。

### 如何查看 DTBP 的值？
A：  
通過 [查詢帳戶資金](../trade/get-funds.html#1465) 接口，可以獲取日內交易相關的返回值，如：剩餘日內交易次數、初始日內交易購買力、剩餘日內交易購買力等。


## Q13：如何跟蹤訂單成交狀態
A:
下單後，可使用以下接口跟蹤訂單成交狀態：
<table>
    <tr>
      <th> 交易環境 </th>
      <th> 接口 </th>
    </tr>
    <tr>
      <td > 真實交易 </td>
      <td > [響應訂單推送回調](../trade/update-order.html)，[響應成交推送回調](../trade/update-order-fill.html) </td>
    </tr>
    <tr>
	  <td> 模擬交易</td>
      <td> [響應訂單推送回調](../trade/update-order.html)</td>
    </tr>
</table>

注意：對於非 python 語言用户，在使用上述兩個接口之前，需要先進行 [訂閲交易推送](../trade/sub-acc-push.html)

#### 響應訂單推送回調 的特點：
反饋 整個訂單 的資訊變動。當以下 8 個字段發生變化時，會觸發訂單推送：  
`訂單狀態`，`訂單價格`，`訂單數量`，`成交數量`，`觸發價格`，`跟蹤類型`，`跟蹤金額/百分比`，`指定價差`  

因此，當您進行下單、修改訂單，取消訂單、使生效、使失效操作，或者訂單在市場中發生了高級訂單被觸發、有成交變動的情況，都會觸發訂單推送。您只需要調用 [響應成交推送回調](../trade/update-order-fill.html)，即可監聽這些資訊。

#### 響應成交推送回調 的特點：
只反饋 單筆成交 的資訊。當以下 1 個字段發生變化時，會觸發訂單推送：  
`成交狀態`  

舉例：假設一筆限價單訂單 900 股，分成了 3 次才完全成交，每次成交分別是：200、300、400 股。  
![example](../img/example.png)


## Q14：下單接口返回“此產品最小單位為 xxx，請調整至最小單位的整數倍後再提交”？
A:  
對於不同市場的標的，交易所有着不同的最小變動單位要求。如果提交的訂單價格不符合要求，訂單將會被拒絕。各市場價位規則如下：  

### 價位規則
#### 香港市場

以港交所官方説明為準，點擊 [這裏](https://www.futufin.com/hans/support/topic605?lang=zh-cn)。


#### A 股市場
股票價位：0.01。

#### 美國市場
股票價位：
<table>
    <tr>
      <th> 合約價格 </th>
      <th> 價位 </th>
    </tr>
    <tr>
      <td > $1 以下 </td>
      <td > $0.0001 </td>
    </tr>
    <tr>
	  <td> $1 以上</td>
      <td> $0.01 </td>
    </tr>
</table>

期權價位：
<table>
    <tr>
      <th> 合約價格 </th>
      <th> 價位 </th>
    </tr>
    <tr>
      <td > $0.10 - $3.00 </td>
      <td > $0.01 或者 $0.05</td>
    </tr>
    <tr>
	  <td> $3.00 以上</td>
      <td> $0.05 或者 $0.10</td>
    </tr>
</table>

期貨價位：不同合約價位規則不同。可以通過 [獲取期貨合約資料](../quote/get-future-info.html#5784) 接口的返回欄位 `最小變動的單位` 查看。

### 怎麼避免訂單價格不在價位上？
* 方法一：通過 [獲取實時擺盤](../quote/get-order-book.html) 接口，獲取合法的交易價格。交易所擺盤上的價位一定是合法的價位。  
* 方法二：通過 [下單](../trade/place-order.html) 接口的參數 `價格微調幅度`，將傳入價格自動調整到合法的交易價格上。  

   例如：假設騰訊控股當前市價為 359.600，根據價位規則，對應的最小變動價位為 0.200。  

   假設您的下單傳入訂單價格為 359.678，價格微調幅度為 0.0015，代表接受 OpenD 對傳入價格自動向上調整到最近的合法價位，且不能超過 0.15%。此情景下，向上最近的合法價格為 359.800，價格實際需要調整的幅度為 0.034%，符合價格微調幅度的要求，因此最終提交的訂單價格為 359.800。  

   若價格微調幅度設置數值小於實際需要調整的幅度，OpenD 自動調整價位失敗，訂單仍會返回報錯“訂單價格不在價位上”。


## Q15：我的購買力足夠，為什麼下市價單會返回“購買力不足”？
A：
### 為什麼市價單會提示購買力不足  
- 出於風控考量，系統給了市價單較高的購買力系數。在所有訂單參數都相同的情況下，選擇市價單會比限價單佔用更多的購買力。  
- 而且對於不同的品種，和不同的市場情況，風控系統會對市價單的購買力系數做動態調整。所以在下市價單時，若您通過最大購買力去計算最大可買數量，計算的結果很可能是不準確的。  
### 如何計算正確的可買數量  
不建議自己計算，您可以通過 [查詢最大可買可賣](../trade/get-max-trd-qtys.html) 接口獲取正確的可買數量。  
### 如何儘可能買更多  
您可以用價格為對價的限價單，替代市價單進行交易。  
其中，對價：買1價（下賣單時）或 賣1價（下買單時）  


## Q16：API模擬交易下單，支持美股融資融券模擬帳戶接入
A：  
API模擬交易下單，已經支持美股融資融券模擬帳戶接入，交易功能更全面。  
原API接口稍後階段陸續停止美股模擬交易服務，為保障更優質的使用體驗，建議您盡快切換至新接口，暢享專業的美股模擬交易服務。


## Q17：交易接口參數使用説明
### 1. 什麼是交易對象？
您的平台賬號下一般會開設一個保證金綜合帳戶，其中有多個交易子帳戶（正常有兩個，一個綜合證券帳戶，一個綜合期貨帳戶；根據需要還可能有綜合外匯帳戶等其他子帳戶）。一些特殊用户或機構客户可能會在多個券商下開設多個綜合帳戶。  
創建交易對象，是初步篩選子帳戶的過程。
- 使用 OpenSecTradeContext 創建的交易對象，調用 get_acc_list 時只會返回**證券交易帳戶**
- 使用 OpenFutureTradeContext 創建的交易對象，調用 get_acc_list 時只會返回**期貨交易帳戶**  

參數 security_firm 用來篩選對應歸屬券商的帳戶，參數 filter_trdmarket  用來篩選對應交易市場權限的帳戶。
#### 1.1 security_firm 券商參數
Futu API 目前支持的券商有 [這些](../trade/trade.html#6090)。  
創建的交易對象，在調用 get_acc_list 時，會返回 security_firm 對應券商的真實帳戶和所有模擬交易帳戶（這是因為模擬交易沒有券商的概念，所以無論 security_firm 傳什麼，都會返回所有的模擬帳戶）。  
security_firm 的預設值是 FUTUSECURITIES，FUTU HK 券商帳戶可以不填此參數，但需要獲取其他券商的帳戶時，需要修改券商參數。  
* **Example 1**
```python
trd_ctx = OpenSecTradeContext(security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.get_acc_list()
print(data)
```
* **Output**
```python
               acc_id   trd_env acc_type      uni_card_num          card_num   security_firm sim_acc_type                  trdmarket_auth acc_status
0  281756478396547854      REAL   MARGIN  1001200163530138  1001369091153722  FUTUSECURITIES          N/A  [HK, US, HKCC, HKFUND, USFUND]     ACTIVE
1             3450309  SIMULATE     CASH               N/A               N/A             N/A        STOCK                            [HK]     ACTIVE
2             3548731  SIMULATE   MARGIN               N/A               N/A             N/A       OPTION                            [HK]     ACTIVE
3  281756455998014447      REAL   MARGIN               N/A  1001100320482767  FUTUSECURITIES          N/A                            [HK]   DISABLED
```

* **Example 2**
```python
trd_ctx = OpenSecTradeContext(security_firm=SecurityFirm.FUTUSG)
ret, data = trd_ctx.get_acc_list()
print(data)
```
* **Output**
```python
    acc_id   trd_env acc_type uni_card_num card_num security_firm sim_acc_type trdmarket_auth acc_status
0  3450309  SIMULATE     CASH          N/A      N/A           N/A        STOCK           [HK]     ACTIVE
1  3548731  SIMULATE   MARGIN          N/A      N/A           N/A       OPTION           [HK]     ACTIVE
```


#### 1.2 filter_trdmarket 交易市場參數
Futu API 目前支持的交易市場有 [這些](../trade/trade.html#1256)。
創建的交易對象，在調用 get_acc_list 時，會返回所有擁有 filter_trdmarket 市場交易權限的帳戶；當 filter_trdmarket 入參傳 NONE 時，不過濾市場，返回所有的帳戶。  
filter_trdmarket 的預設參數是 HK，在綜合帳戶體系下，這個參數用來篩選不同市場下的模擬交易帳戶。  
* **Example 1**
```python
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.US)
ret, data = trd_ctx.get_acc_list()
print(data)
```
* **Output**
```python
               acc_id   trd_env acc_type      uni_card_num          card_num   security_firm sim_acc_type                  trdmarket_auth acc_status
0  281756478396547854      REAL   MARGIN  1001200163530138  1001369091153722  FUTUSECURITIES          N/A  [HK, US, HKCC, HKFUND, USFUND]     ACTIVE
1             3450310  SIMULATE   MARGIN               N/A               N/A             N/A        STOCK                            [US]     ACTIVE
2             3548732  SIMULATE   MARGIN               N/A               N/A             N/A       OPTION                            [US]     ACTIVE
3  281756460292981743      REAL   MARGIN               N/A  1001100520714263  FUTUSECURITIES          N/A                            [US]   DISABLED
```

* **Example 2**
```python
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.NONE)
ret, data = trd_ctx.get_acc_list()
print(data)
```
* **Output**
```python
                acc_id   trd_env acc_type      uni_card_num          card_num   security_firm sim_acc_type                  trdmarket_auth acc_status
0   281756478396547854      REAL   MARGIN  1001200163530138  1001369091153722  FUTUSECURITIES          N/A  [HK, US, HKCC, HKFUND, USFUND]     ACTIVE
1              3450309  SIMULATE     CASH               N/A               N/A             N/A        STOCK                            [HK]     ACTIVE
2              3450310  SIMULATE   MARGIN               N/A               N/A             N/A        STOCK                            [US]     ACTIVE
3              3450311  SIMULATE     CASH               N/A               N/A             N/A        STOCK                            [CN]     ACTIVE
4              3548732  SIMULATE   MARGIN               N/A               N/A             N/A       OPTION                            [US]     ACTIVE
5              3548731  SIMULATE   MARGIN               N/A               N/A             N/A       OPTION                            [HK]     ACTIVE
6   281756455998014447      REAL   MARGIN               N/A  1001100320482767  FUTUSECURITIES          N/A                            [HK]   DISABLED
7   281756460292981743      REAL   MARGIN               N/A  1001100520714263  FUTUSECURITIES          N/A                            [US]   DISABLED
8   281756468882916335      REAL   MARGIN               N/A  1001100610464507  FUTUSECURITIES          N/A                          [HKCC]   DISABLED
9   281756507537621999      REAL     CASH               N/A  1001100910390035  FUTUSECURITIES          N/A                        [HKFUND]   DISABLED
10  281756550487294959      REAL     CASH               N/A  1001101010406844  FUTUSECURITIES          N/A                        [USFUND]   DISABLED
```
::: tip 提示  
當 filter_trdmarket 入參NONE時，可以返回所有的交易帳戶。其中第0行是真實帳戶，1~5行均為模擬交易帳戶，6~10行是已失效的真實帳戶。這些失效帳戶都是單市場帳戶，現已被綜合帳戶替代。但歷史訂單和歷史成交還在這些已失效的帳戶中，可以通過這些帳戶來查詢。  
OpenFutureTradeContext 對象中沒有 filter_trdmarket 參數，只有 security_firm 參數，功能與 OpenSecTradeContext  一樣。  
:::  

### 2. 交易接口參數
在使用具體的交易接口（如下單、查詢訂單列表）時，接口中的 `trd_env`, `acc_index` 和 `acc_id` 參數，會先篩選確認一個唯一的帳戶，對此帳戶實施對應的接口行為。
![acc-select](../img/acc-select.jpg)

::: tip 總結
1. 根據 trd_env 篩選出真實帳戶還是模擬帳戶
2. 在篩選結果中，優先選擇 acc_id 指定的帳戶
3. 如果 acc_id 為0，則通過 acc_index選取對應賬號
4. 報錯場景：指定的 acc_id 不存在，或 acc_index 超出範圍  
:::


### 3. 應用舉例
#### 3.1 綜合證券帳戶實盤下單
```python
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.NONE, security_firm=SecurityFirm.FUTUSECURITIES)
ret, data = trd_ctx.unlock_trade("123123")
if ret == RET_OK:
    print("解鎖成功")
    ret, data = trd_ctx.place_order(45, 200, 'HK.00700', TrdSide.BUY,
                                    order_type=OrderType.NORMAL,
                                    trd_env=TrdEnv.REAL,  # 和預設參數一樣，可以不填
                                    acc_id=0)  # 和預設參數一樣，可以不填
    print(data)
```

#### 3.2 綜合期貨帳戶查詢實盤訂單列表
```python
trd_ctx = OpenFutureTradeContext(security_firm=SecurityFirm.FUTUSECURITIES)

ret, data = trd_ctx.order_list_query(trd_env=TrdEnv.REAL,   # 和預設參數一樣，可以不填
                                     acc_id=0)  # 和預設參數一樣，可以不填
print(data)
```

#### 3.3 港股模擬現金帳戶查詢帳戶資金
```python
# filter_trdmarket 填 TrdMarket.HK
# trd_env 填 TrdEnv.SIMULATE
# acc_index 填 0
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.HK)
ret, data = trd_ctx.accinfo_query(trd_env=TrdEnv.SIMULATE, acc_index=0)
print(data)
```

#### 3.4 美股模擬保證金帳戶下單期權
```python
# 通過 filter_trdmarket 和 trd_env 篩選完之後只剩兩個帳戶
# 第0個是美股現金帳戶（交易股票）,第1個是美股保證金帳戶（交易期權）
# acc_index 填 1 指定美股保證金帳戶
trd_ctx = OpenSecTradeContext(filter_trdmarket=TrdMarket.US)
ret, data = trd_ctx.place_order(10, 1, code="US.AAPL250618P550000",trd_side=TrdSide.BUY,
                                trd_env=TrdEnv.SIMULATE,
                                acc_index=1)
print(data)
```

#### 3.5 日本期貨模擬帳戶查詢最大可買賣
```python
# 將 get_acc_list 的結果打印出來，可以看到日本期貨模擬帳戶的 acc_id 是 6271199
# 請求最大可買賣接口時傳入這個 acc_id 
trd_ctx = OpenFutureTradeContext()
ret, data = trd_ctx.acctradinginfo_query(order_type=OrderType.NORMAL,
                                         price=5000,
                                         trd_env=TrdEnv.SIMULATE,
                                         acc_id=6271199,
                                         code="JP.NK225main")
print(data)
```


### 4. API 中的帳戶如何與 APP/桌面版對應

![card-app](../img/card-app.png)
APP 上的帳戶僅顯示卡號後 4 位數字，我們將 [get_acc_list](../trade/get-acc-list.html) 的返回結果打印出來後，有 uni_card_num 列和 card_num 列，分別對應綜合帳戶的卡號，和單幣種帳戶（已廢棄）的卡號。通過卡號後 4 位數就能把 API 中獲取到的賬號與 APP 上對應起來了。

---



---

# 其他

## Q1：如何編譯C++ API？

A: 
futu api c++ SDK支持Windows/MacOS/Linux，每個系統提供了以下編譯環境生成的程式庫檔案：
操作系統|編譯工具
:-|:-
Windows |Visual Studio 2013
Centos 7|g++ 4.8.5
Ubuntu 16.04|g++ 5.4.0
MacOS | XCode 11

如果編譯器版本不同，或相依性項的protobuf版本不同，則可能需要自己使用原始碼重新編譯FTAPI和protobuf，原始碼位置見下圖目錄：

```
FTAPI目錄結構：
+---Bin                               存放各個系統預設編譯環境編譯出的相依性項庫
+---Include                           存放公共標頭檔，以及proto協議生成的.h/.cc文件
+---Sample                            示例專案
\---Src
    +---FTAPI                         FTAPI原始碼
    +---protobuf-all-3.5.1.tar.gz     protobuf原始碼
```

#### 編譯步驟：
1. 重新編譯protobuf：生成libprotobuf靜態程式庫
2. 從協議proto檔案中生成C++檔案
3. 重新編譯FTAPI: 原始碼在Src/FTAPI，生成libFTAPI靜態程式庫

#### 步驟1： 重新編譯protobuf：
- Windows：
  - 安裝CMake
  - 打開VS命令行工具，cd到protobuf/cmake目錄
  - 執行：cmake -G "Visual Studio 12 2019" -DCMAKE_INSTALL_PREFIX=install -Dprotobuf_BUILD_TESTS=OFF  這樣會生成Visual Studio 2019的項目檔案，其它版本Visual Studio請修改-G參數
  - 打開生成的Visual Studio項目檔案，平台工具組設置為v120_xp，編譯即可
- Linux（參考protobuf/src/README）
  - 執行 ./autogen.sh
  - 執行 CXXFLAGS="-std=gnu++11" ./configure --disable-shared
  - 執行 make
  - 將生成的libprotobuf.a放入Bin/Linux目錄
- MacOS（參考protobuf/src/README）
  - 使用brew安裝這些相依性項庫：autoconf automake libtool
  - 執行./configure CC=clang CXX="clang++ -std=gnu++11 -stdlib=libc++" --disable-shared

#### 步驟2: 重新生成proto代碼
- 上面編譯Protobuf後會同時生成可執行檔案protoc。用protoc將Include/Proto下面的.proto檔案生成對應的.h和.cc檔案。例如命令以下命令會從Common.proto生成對應的Common.pb.h和Common.pb.cc
  - protoc -I="FTAPI路徑/Include/Proto" --cpp_out="." FTAPI路徑/Include/Proto/Common.proto
- 將生成的.h和.cc檔案放到Include/Proto下面

#### 步驟3: 重新編譯FTAPI
- Windows：新建Visual Studio C++靜態程式庫專案，將Src/FTAPI和Include下的原始碼加入專案中，平台工具組設置為v120_xp，然後編譯
- Mac：新建XCode C++靜態程式庫專案，將Src/FTAPI和Include下的原始碼加入專案中，然後編譯
- Linux：使用CMake編譯FTAPI靜態程式庫，在FTAPI路徑/Src目錄下執行：
  - cmake -DTARGET_OS=Linux

## Q2：有沒有更完整的策略範例可以參考？

A:
* Python 策略範例在 /futu/examples/ 資料夾下。您可以通過執行如下命令，找到 Python API 的安裝路徑：
    ```
    import futu
    print(futu.__file__)
    ```
* C# 策略範例在 /FTAPI4NET/Sample/ 資料夾下
* Java 策略範例在 /FTAPI4J/sample/ 資料夾下
* C++ 策略範例在 /FTAPI4CPP/Sample/ 資料夾下
* JavaScript 策略範例在 /FTAPI4JS/sample/ 資料夾下


## Q3：使用 python API 匯入異常

A：

**場景一**：已經在 Python 環境中安裝了 futu 模組，仍然提示 No module named 'futu'？  
很可能是因為當前 IDE 所使用的 interpreter 並不是你裝過 futu 模組的 interpreter。也就是説，您的電腦可能裝了兩個以上的 Python 環境。
您可以操作如下兩步：
1. 在 Python 中運行如下代碼，得到當前 interpreter 的路徑：
```
import sys
print(sys.executable)
```
示例圖：  
 ![No module named 'futu'](../img/import-futu-error.png)

2. 在命令行中，執行 `$ D:\software\anaconda3\python.exe -m pip install futu-api`（其中前半部分的檔案路徑來自第 1 步打印的路徑）。
這樣就可以在當前的 interpreter 中也安裝一份 futu 模組。

## Q4： import 成功了，仍然調用不了相關接口？ 

A：通常遇到這種情況，需要確認一下：成功匯入的 futu，是不是真正的 Futu API 模組。以下幾種場景也可能 import 成功。

**場景一**：存在與“futu”重名的檔案

  1. 當前檔案名是 futu.py
  2. 當前檔案所在目錄下存在另一個名為 futu.py 的檔案
  3. 當前檔案所在目錄下存在名為 `/futu` 的資料夾    

因此，我們強烈建議您，在給檔案 / 資料夾 / 專案起名的時候，不要起名叫“futu”。重名一時爽，查 bug 兩行淚。

**場景二**：誤安裝了一個名為“futu”的第三方程式庫  

   Futu API 的正確名稱為`futu-api`，而非“futu”。   

   如果您安裝過名為“futu”的第三方程式庫，請將其卸載，並 [下載 futu-api](../quick/demo.md#2556)。
   
   以 PyCharm 為例：查看第三方程式庫的安裝情況。

   ![settings](../img/settings.png)  
   ![futuku](../img/futuku.png)


## Q5：協議加密相關

A：  
### 概述
您可以使用非對稱加密演算法 RSA，對策略程式（Futu API）與 OpenD 之間的請求和返回內容進行加密，以保證通信安全。  
如果您的策略程式（Futu API）與 FutuOpenD 在同一台電腦上，則通常無需加密。

### 協議加密流程
您可以嘗試通過以下步驟解決此問題：
1. 通過第三方 web 平台自動生成密鑰檔案。  
    - 具體方法：在 baidu 或 google 上搜索“RSA 在線生成”，**密鑰格式**設置為 PKCS#1，**密鑰長度**設置為 1024 bit，不需要設置私鑰密碼，點擊**生成密鑰對**。  
    ![ui-config](../img/create_rsa.png)  

2. 將生成的 **RSA 加密私鑰** 複製粘貼至 txt 記事本，並保存至 OpenD 所在電腦的指定路徑。
3. 在 OpenD 所在的電腦中，指定 **RSA 加密私鑰** 的路徑。  
    - 方式一：在 [可視化 OpenD](../quick/opend-base.md#3795) 啟動界面右側的“加密私鑰”一欄，指定上一步驟中放置 **RSA 加密私鑰** 的路徑。如下圖所示：  
    ![ui-config](../img/nnrsa_ui-config.png)  
    - 方式二：在 [命令行 OpenD](../opend/opend-cmd.md#1028) 啟動檔案 OpenD.xml 中，找到參數`rsa_private_key`，將其設定為第 2 步中 **RSA 加密私鑰** 的路徑。如下圖所示：  
    ![ui-config](../img/nnrsa_xml.png)  
4. 將第 2 步中 txt 檔案另存至策略程式（Futu API）所在電腦的指定路徑， 並在策略程式中將此路徑 [設置為私鑰路徑](../ftapi/init.md#7106)。
5. 在策略程式（Futu API）中啟用協議加密。 啟用協議加密的方式有兩種，其中方式二的優先級更高。
    - 方式一：對單條的連接加密（通用）。在對 [行情對象](../quote/base.md#9160) 或 [交易對象](../trade/base.md#4802) 創建連接時，通過 **是否啟用加密** 參數設置加密。
    - 方式二：對所有的連接加密（僅 Python）。通過`enable_proto_encrypt`接口設置加密，詳見 [這裏](../ftapi/init.md#1542)。


:::tip 提示
* 在 OpenD 或策略程式（Futu API）中指定 **RSA 加密私鑰** 路徑時，需指定至 txt 檔案本身。
* RSA 加密公鑰無需保存，可通過私鑰計算得到。
:::


## Q6：為什麼我獲取的 DataFrame 數據，只能展示一部分 ？

A：打印 pandas.DataFrame 數據的時候，如果行列數過多，pandas 預設會將數據摺疊，導致看起來顯示不全。  
因此，並不是接口返回數據真的不全。您只需要在 Python 程式前面加上如下代碼即可解決。

```
import pandas as pd
pd.options.display.max_rows=5000
pd.options.display.max_columns=5000
pd.options.display.width=1000
```

## Q7：Mac 機器使用 C++ 語言的 API，遇到 “無法打開 libFTAPIChannel.dylib” 的問題

A：在對應程式庫目錄中執行以下命令即可解決:`$ xattr -r -d com.apple.quarantine libFTAPIChannel.dylib`。


## Q8：Python 用戶，為什麼在 OpenD 設定檔中設置了日誌級別為 no 後，log 資料夾下仍然持續產生超大容量的日誌檔案？

A：OpenD 設定檔中的日誌級別參數，只用來控制 OpenD 產生的日誌。而 Python API 預設也會產生日誌，如果您不希望希望 Python API 產生日誌，可以在 Python 程式加上如下語句：

```
logger.file_level = logging.FATAL  # 用於關閉 Python API 日誌
logger.console_level = logging.FATAL  # 用於關閉 Python 運行時的控制台日誌
```


## Q9：對於 5.4 及以上的版本，Java API 的程式庫名稱和設定方式的變更

A:
* 如果您是 Java API 5.3 及以下版本的用戶，在更新版本時，請注意以下變更：

  **設定流程的變更**：

  1. 通過 [富途牛牛官網](https://www.futunn.com/download/OpenAPI)下載 Futu API。
  2. 解壓下載好的 FTAPI 檔案，`/FTAPI4J` 是 Java API 的目錄，將目錄結構中的 `/lib/futu-api-.x.y.z.jar` 添加到您的專案設置中。創建 futu-api 專案請參考 [這裏](../quick/demo.html#2556)。


  **目錄結構的變更**：
  1. Futu API 的 Java 版本，程式庫名稱由之前的 ftapi4j.jar 變更為 `futu-api-x.y.z.jar`，其中 “x.y.z” 表示版本編號。
  2. 第三方程式庫的引用中，去掉了 /lib/jna.jar 和 /lib/jna-platform.jar 相依性項，增加了 `/lib/bcprov-jdk15on-1.68.jar` 和 `/lib/bcpkix-jdk15on-1.68.jar` 相依性項。
    ```
    +---ftapi4j                      futu-api 原始碼，如果所用 JDK 版本不兼容可以用這裏的專案重新編譯出 futu-api.jar
    +---lib                          存放公共庫文件
    |    futu-api-x.y.z.jar          Futu API 的 Java 版本
    |    bcprov-jdk15on-1.68.jar     第三方程式庫，用於加解密
    |    bcpkix-jdk15on-1.68.jar     第三方程式庫，用於加解密
    |    protobuf-java-3.5.1.jar     第三方程式庫，用於解析 protobuf 數據
    +---sample                       示例專案
    +---resources                    maven 專案預設生成的目錄
    ```
* 如果您第一次接觸 Futu API，我們提供了更便捷的通過 maven 倉庫設定 Java API 的方式。設定流程請參考 [這裏](../quick/demo.html#518)。


## Q10：Python 用戶，使用 pyinstaller 打包程式時報錯：找不到 Common_pb2 模組

A：你可以嘗試通過以下步驟解決此問題：
1. 假設你需要對 main.py 進行打包。使用命令行語句，運行代碼：pyinstaller main.py，不要加參數 “- F”（path 為 main.py 的所在路徑）
  ```
  pyinstaller path\main.py
  ```
  打包成功後，main.py 所在目錄下的 /dist 中，會生成 /main 資料夾，main.exe 就在這個資料夾中。  
  ![dist](../img/dist.png)  
2. 運行以下代碼，找到 futu-api 的安裝目錄。  
  ```
  import futu
  print(futu.__file__)
  ```
  運行結果:  
  ```
  C:\Users\ceciliali\Anaconda3\lib\site-packages\futu\__init__.py
  ```
  ![path_futu](../img/path_futu.png)  

3. 打開上圖資料夾中的 /common/pb，將所有檔案全部複製到 /main 中。

4. 在 /main 中創建資料夾，命名為 futu，將上圖資料夾中的 `VERSION.txt` 檔案複製到 /main/futu 中。  
  ![main_futu](../img/main_futu.png) 
5. 再次嘗試運行 main.exe

## Q11：接口調用結果正常，但其返回表現不符合預期？
A:
* 接口調用結果正常，表示富途已經成功收到並響應了您的請求，但接口返回表現可能與您的預期不符。  

  例如：若您在非交易時段調用 [訂閲](../quote/sub.md) 接口，雖然您的請求可以被成功響應，並且接口調用結果正常，但在非交易時段下，交易所無行情數據變動，所以您將暫時無法收到行情數據推送，直至市場重新回到交易時段。  
* 接口調用結果可以通過返回欄位（定義參見：[接口調用結果](../ftapi/common.md#3835)）查看，返回欄位為 0 代表接口調用正常，非 0 代表接口調用失敗。  
  
  對於 Python 用戶，下面兩種寫法等價：
  ```
  if ret_code == RET_OK:
  ```
  ```
  if ret_code == 0:
  ```

## Q12：WebSocket相關
A：

### 概述

Futu API 中，WebSocket 主要用於以下兩方面：
* 可視化 OpenD 中，UI 界面跟底層的命令行 OpenD 的通信使用 WebSocket 方式。
* JavaScript API 跟 OpenD 之間的通信使用 WebSocket 方式。

![WebSocket-struct](../img/WebSocket-struct.png)  
* 當 WebSocket 啟動時，命令行 OpenD 會與 **FTWebSocket 中轉服務** 建立 Socket 連接（TCP），這一連接會用到預設的 **監聽地址** 和 **API 協議監聽連接埠**。
* 同時，JavaScript API 會與 **FTWebSocket 中轉服務** 建立 WebSocket 連接（HTTP），這一連接會用到 **WebSocket 監聽地址** 和 **WebSocket 連接埠**。

### 使用
為保證帳戶安全，當 WebSocket 監聽來自非本地請求時，我們強烈建議您啟用 SSL 並設定 **WebSocket 鑑權密鑰**。

SSL 通過在設定 **WebSocket 證書** 以及 **WebSocket 私鑰** 來啟用。  
命令行 OpenD 可通過設定 OpenD.xml 或設定命令行參數來設置檔案路徑。可視化 OpenD 點擊【更多選項】下拉菜單，可以看到設置項。

![ui-more-config](../img/ui-more-config.png)

::: tip 提示
如果證書是自籤的，則需要在調用 JavaScript 接口所在機器上安裝該證書，或者設置不驗證證書。
:::

#### 生成自簽證書
自簽證書生成詳細資料不便在此文檔展開，請自行查閲。  
在此提供較簡單可用的生成步驟：
1. 安裝 openssl。
2. 修改 openssl.cnf，在 alt_names 節點下加上 OpenD 所在機器 IP 地址或域名。  
例如：IP.2 = xxx.xxx.xxx.xxx, DNS.2 = www.xxx.com
3. 生成私鑰以及證書（PEM）。

**證書生成參數參考如下**：  
`openssl req -x509 -newkey rsa:2048 -out futu.cer -outform PEM -keyout futu.key -days 10000 -verbose -config openssl.cnf -nodes -sha256 -subj "/CN=Futu CA" -reqexts v3_req -extensions v3_req`

::: tip 提示
* openssl.cnf 需要放到系統路徑下，或在生成參數中指定絕對路徑。
* 注意生成私鑰需要指定不設置密碼（-nodes）。
:::

附上本地自簽證書以及生成證書的設定檔供測試：  
* [openssl.cnf](../file/openssl.cnf)  
* [futu.cer](../file/cer)  
* [futu.key](../file/key)

## Q13：API 的行情和交易服務分別部署在哪裏？
A：  
- 行情：  

平台賬號|行情伺服器所在地
:-|:-|:-
牛牛號|騰訊雲廣州和香港
moomoo 號|騰訊雲美國弗吉尼亞和新加坡

- 交易：  

所屬券商|交易伺服器所在地
:-|:-|:-
富途證券(香港)|香港
moomoo證券(美國)|騰訊雲美國弗吉尼亞
moomoo證券(新加坡) |騰訊雲新加坡
moomoo證券(澳大利亞)|騰訊雲新加坡
moomoo證券(馬來西亞)|阿里雲馬來西亞
moomoo證券(加拿大)|AWS加拿大
moomoo證券(日本)|騰訊雲日本


## Q14：關於綜合帳戶升級的過渡指引

### 1. [**綜合帳戶升級**](https://www.futuhk.com/hans/support/topic2_1734)
綜合帳戶支持以多種貨幣在同一個帳戶內交易不同市場品類。從單幣種帳戶升級到綜合帳戶，是在您原來的牛牛號下，進行帳戶遷移。主要包括：
- 創建新的綜合帳戶
- 將您原來單幣種業務帳戶裏的資產，轉移到綜合帳戶裏
- 關閉原來的單幣種帳戶

### 2. **OpenD版本更新**
我們會在 2024年9月14日、15日 集中為 Futu API 客户的帳戶做升級，請提前檢查 OpenD 和 API 版本編號：  
- **7.01 及以下版本**  
  OpenD 因版本過舊，將於 2024/09/14 停止服務。屆時，已登入的帳戶會被強制退出登入。我們建議您在 9/14 之前升級 [OpenD](../quick/opend-base.html#3795) 和 [API](../quick/demo.html#1999) 至最新版本，且不要在 9/14~9/15 期間跨週末運作策略。
- **7.02 ~ 8.2 版本**  
  OpenD 版本較舊，無法獲取綜合帳戶。我們建議您在 9/14 之前升級 [OpenD](../quick/opend-base.html#3795) 和 [API](../quick/demo.html#1999) 至最新版本，且不要在 9/14~9/15 期間跨週末運作策略。
- **8.3 及以上版本**  
  可以正常使用，我們建議您不要在 9/14~9/15 期間跨週末運作策略。  

綜合帳戶升級時，您的資產會轉移到新的綜合帳戶，如果策略指定舊的帳戶，可能會運行異常。同時，在實盤交易之前，建議您進行必要的檢查與測試，確保一切設置正常。

### 3. **帳戶升級後，Futu API有哪些表現？**
- Python API 將不再支持使用 OpenHKTradeContext,  OpenUSTradeContext, OpenHKCCTradeContext, OpenCNTradeContext 創建交易對象，請參考 [創建交易對象連接](../trade/base.html#9291) 改用 OpenSecTradeContext。  
- 非Python API用戶，在使用 Trd_GetAccList 接口時，需要將 needGeneralSecAccount 參數設為 true，才能獲取到綜合帳戶的相關資訊。  
- 帳戶新增 [帳戶狀態](../trade/trade.html#3606): 在使用 [獲取交易業務帳戶列表](../trade/get-acc-list.html#4991) 時，返回結果新增了帳戶狀態 。綜合帳戶標記為 `ACTIVE` 生效帳戶，被停用的單幣種帳戶標記為 `DISABLED` 失效帳戶。  
- [下單](../trade/place-order.html#6160)、[改單撤單](../trade/modify-order.html#6716)、[查詢最大可買可賣](../trade/get-max-trd-qtys.html#3188) 等交易接口表現  
    - 支持使用 `ACTIVE` 生效帳戶所對應的 acc_id或acc_index 進行購買力查詢與交易。
    - 不支持使用 `DISABLED` 失效帳戶所對應的 acc_id或acc_index 進行購買力查詢與交易，若使用，將會出現報錯資訊。
    - Python API用戶：在接口入參中，請指定 acc_id 為升級後的綜合帳戶。
    - 非Python API用戶：在 TrdHeader 中，請指定accID為升級後的綜合帳戶。

---

