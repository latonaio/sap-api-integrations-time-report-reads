# sap-api-integrations-time-report-reads  
sap-api-integrations-time-report-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API タイムレポートデータを取得するマイクロサービスです。  
sap-api-integrations-time-report-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-time-report-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/timereport/overview  

## 動作環境
sap-api-integrations-time-report-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-time-report-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-time-report-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/timereport/overview  
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-time-report-reads には、次の API をコールするためのリソースが含まれています。  

* TimeReportCollection（タイムレポート - タイムレポート）  
* TimeReportPartyCollection（タイムレポート - 関係者）  

## API への 値入力条件 の 初期値
sap-api-integrations-time-report-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.TimeReportCollection.ID（ID）  
* inoutSDC.TimeReportCollection.TimeReportCollection.ObjectID（対象ID）  
* inoutSDC.TimeReportCollection.TimeReportCollection.PartyID（関係者ID）  

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"TimeReportCollection" が指定されています。    
  
```
	"api_schema": "TimeReport",
	"accepter": ["TimeReportCollection"],
	"time_report_code": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "TimeReport",
	"accepter": ["All"],
	"time_report_code": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetTimeReport(iD, objectID, partyID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "TimeReportCollection":
			func() {
				c.TimeReportCollection(iD)
				wg.Done()
			}()
		case "TimeReportPartyCollection":
			func() {
				c.TimeReportPartyCollection(objectID, partyID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP タイムレポート の タイムレポートデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ETag" は、/SAP_API_Output_Formatter/type.go 内 の Type TimeReportCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-time-report-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-time-report-reads/SAP_API_Caller.(*SAPAPICaller).TimeReportCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1EE79C91FAF19C5B4BE1",
			"Description": "",
			"languageCode": "",
			"EmployeeUUID": "00163E05-AE66-1ED3-88BF-33E82FE7E463",
			"EndDate": "2017-07-24T09:00:00+09:00",
			"ID": "1",
			"InformationLifeCycleStatusCode": "",
			"RejectionReason": "",
			"languageCode1": "",
			"ReportName": "24 July 2017",
			"languageCode2": "",
			"StartDate": "2017-07-24T09:00:00+09:00",
			"Status": "3",
			"CreationDateTime": "2017-07-25T01:48:18+09:00",
			"CreationIdentityUUID": "00163E05-AE66-1ED3-88BF-3F0E7DE5A467",
			"LastChangeDateTime": "2017-07-25T13:29:37+09:00",
			"LastChangeIdentityUUID": "00163E03-A070-1EE2-88BA-3917457F10B3",
			"TotalDuration": "PT2H15M",
			"EntityLastChangedOn": "2017-07-25T13:29:37+09:00",
			"ETag": "2017-07-25T13:29:37+09:00"
		}
	],
	"time": "2022-08-10T10:30:29+09:00"
}
```