window.hostDomain = {
  "PCBILE": "https://www.hotcoin.com/",
  "MOBILE": "https://m.hotcoin.com/",
  "H5MOBILE": "https://m.hotcoin.com/",
  "H5LITE": "https://114.111.61.68/",
  "BASE_URL": "https://bin.hotcoins.cn/hk-web/",
  "OTC_URL": "https://bin.hotcoins.cn/otc-api/",
  "UPLOAD_URL": "https://bin.hotcoins.cn/oss-server/",
  "CHAT_URL": "https://bin.hotcoins.cn/",
  "CONTRACT_URL": "https://bin.hotcoins.cn/swap/",
  "CHAT_WS": "wss://clch.hotcoins.cn/ws",
  "COIN_WS": "wss://wsw.hotcoins.cn/trade/multiple",
  "CONTRACT_WS": "wss://wcws.hotcoins.cn",
  "GA_MEASUREMENT_ID": "G-KHZ9FEDDPX"
};
window.BASEURL = window.hostDomain.BASE_URL;
window.KLINEURL = "https://bin.hotcoins.cn/kline-history/";
window.hotcoin = 1;
window.MATTS_KEY = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApBqET8vPcO84Z7XOTx4FDlsVfRLfraE2miiVwCzlFacL50tX/MUZkTdJYEqDDF1Tcj80vNTjw6HwSgVShwWv8gSxRF77TyQdha/zH15I4Sakq784HIQHcVarcnvT4H/3xbO1SCF+PH+xgjEs0rdYSkiUngdhUufjr64Pgyq6ISflH4OW52JjtXUttRdh6ZjZnVFCuSfvSiJfLO7P2eo3p7g6HMt2xHwo3GgdEG7voDNvxUUtbu5kkFh8nuy/GV7bpIiCe/5sjhbCWMKykSlSYiI1NS50ZeTVBhTYe1484epzeBKQh38p5fen22Rqq1DYRxDtrZ52/mwKDzF8hRa5BQIDAQAB";

var endpoint = {};
if(localStorage.getItem('endpoint')){
  try {
    endpoint = JSON.parse(localStorage.getItem('endpoint'));
  } catch (error) {
    console.error(error);
  }
}

if(endpoint.baseURL){
  window.hostDomain.BASE_URL = `https://${endpoint.baseURL}/hk-web/`;
  window.hostDomain.OTC_URL = `https://${endpoint.baseURL}/otc-api/`;
  window.hostDomain.UPLOAD_URL = `https://${endpoint.baseURL}/oss-server/`;
  window.hostDomain.CHAT_URL = `https://${endpoint.baseURL}/`;
  window.hostDomain.CONTRACT_URL = `https://${endpoint.baseURL}/swap/`;
  window.hostDomain.CHAT_WS = `wss://${endpoint.chatWS}/ws`;
  window.hostDomain.COIN_WS = `wss://${endpoint.tradeWS}/trade/multiple`;
  window.hostDomain.CONTRACT_WS = `wss://${endpoint.contractWS}`;
  window.BASEURL = `https://${endpoint.baseURL}/hk-web/`;
  window.KLINEURL = `https://${endpoint.baseURL}/kline-history/`;
}
