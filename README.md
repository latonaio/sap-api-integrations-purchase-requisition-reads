# sap-api-integrations-purchase-requisition-reads
sap-api-integrations-purchase-requisition-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 購買依頼データを取得するマイクロサービスです。    
sap-api-integrations-purchase-requisition-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchase-requisition-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_PURCHASEREQ_PROCESS_SRV_0001/overview   

## 動作環境  
sap-api-integrations-purchase-requisition-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-purchase-requisition-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchase-requisition-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASEREQ_PROCESS_SRV_0001/overview    
* APIサービス名(=baseURL): API_PURCHASEREQ_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchase-requisition-reads には、次の API をコールするためのリソースが含まれています。  

* A_PurchaseRequisitionHeader（購買依頼 - ヘッダ）※購買依頼関連データを取得するために、ToItem、ToItemDeliveryAddress、ToItemAccount、と合わせて利用されます。  
* ToItem（購買依頼 - 明細  ※To）
* ToItemDeliveryAddress（購買依頼 - 納入先住所  ※To）
* ToItemAccount（購買依頼 - 勘定設定  ※To）
* A_PurchaseRequisitionItem（購買依頼 - 明細）※購買依頼関連データを取得するために、ToItemDeliveryAddress、ToItemAccount、と合わせて利用されます。  
* ToItemDeliveryAddress（購買依頼 - 納入先住所 ※To）
* ToItemAccount（購買依頼 - 明細 ※To）
* A_PurReqAddDelivery（購買依頼 - 納入先住所）
* A_PurReqnAcctAssgmt（購買依頼 - 勘定設定）

## API への 値入力条件 の 初期値
sap-api-integrations-purchase-requisition-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.PurchaseRequisition.PurchaseRequisition（購買依頼）
* inputSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchaseRequisitionItem（購買依頼明細）
* inputSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocument（購買発注）
* inputSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocumentItem（購買発注明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "/sap.s4.beh.purchaserequisition.v1.PurchaseRequisition.Created.v1",
	"accepter": ["Header"],
	"purchase_requisition": "10000010",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/sap.s4.beh.purchaserequisition.v1.PurchaseRequisition.Created.v1",
	"accepter": ["All"],
	"purchase_requisition": "10000010",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetPurchaseRequisition(purchaseRequisition, purchaseRequisitionItem, purchasingDocument, purchasingDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseRequisition)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "ItemDeliveryAddress":
			func() {
				c.ItemDeliveryAddress(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "ItemAccount":
			func() {
				c.ItemAccount(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "PurchasingDocument":
			func() {
				c.PurchasingDocument(purchasingDocument, purchasingDocumentItem)
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
以下の sample.json の例は、SAP 購買依頼  の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"PurchaseRequisition" ～ "to_PurchaseReqnItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchase-requisition-reads/SAP_API_Caller/caller.go#L73",
	"function": "sap-api-integrations-purchase-requisition-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"PurchaseRequisition": "10000010",
			"PurchaseRequisitionType": "NB",
			"SourceDetermination": false,
			"to_PurchaseReqnItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PURCHASEREQ_PROCESS_SRV/A_PurchaseRequisitionHeader('10000010')/to_PurchaseReqnItem"
		}
	],
	"time": "2022-01-28T10:51:32+09:00"
}
```