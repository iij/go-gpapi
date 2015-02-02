//  http://manual.iij.jp/gp/gpapi/11454854.html

package protocol

type DisableIPv6Arg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	IPv6Address   string `gpapi:"IPv6Address,required"`   // 無効化するIPv6アドレス
}

type DisableIPv6Response struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	IPv6Address   string // 無効なったIPv6アドレス
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
