export const ID_NO_PASS_ERROR_MSG = 'ID_NO_PASS_ERROR_MSG'
export const ID_NO_APP_PASS_ERROR_MSG = 'ID_NO_APP_PASS_ERROR_MSG'
export const ID_SET_BUTTON_BUY = 'ID_SET_BUTTON_BUY'
export const ID_SET_BUTTON_SELL = 'ID_SET_BUTTON_SELL'
export const ID_OFF = 'ID_OFF'
export const ID_READY = 'ID_READY'
export const ID_LOCKED = 'ID_LOCKED'
export const ID_NOWALLET = 'ID_NOWALLET'
export const ID_WALLET_SYNC_PROGRESS = 'ID_WALLET_SYNC_PROGRESS'
export const ID_HIDE_ADDIIONAL_SETTINGS = 'ID_HIDE_ADDIIONAL_SETTINGS'
export const ID_SHOW_ADDIIONAL_SETTINGS = 'ID_SHOW_ADDIIONAL_SETTINGS'
export const ID_BUY = 'ID_BUY'
export const ID_SELL = 'ID_SELL'
export const ID_NOT_SUPPORTED = 'ID_NOT_SUPPORTED'
export const ID_CONNECTION_FAILED = 'ID_CONNECTION_FAILED'
export const ID_ORDER_PREVIEW = 'ID_ORDER_PREVIEW'
export const ID_CALCULATING = 'ID_CALCULATING'
export const ID_ESTIMATE_UNAVAILABLE = 'ID_ESTIMATE_UNAVAILABLE'
export const ID_NO_ZERO_RATE = 'ID_NO_ZERO_RATE'
export const ID_NO_ZERO_QUANTITY = 'ID_NO_ZERO_QUANTITY'
export const ID_TRADE = 'ID_TRADE'
export const ID_NO_ASSET_WALLET = 'ID_NO_ASSET_WALLET'
export const ID_EXECUTED = 'ID_EXECUTED'
export const ID_BOOKED = 'ID_BOOKED'
export const ID_CANCELING = 'ID_CANCELING'
export const ID_PASSWORD_NOT_MATCH = 'ID_PASSWORD_NOT_MATCH'
export const ID_DEX_ADDRESS_EMPTY = 'ID_DEX_ADDRESS_EMPTY'
export const ID_ACCT_UNDEFINED = 'ID_ACCT_UNDEFINED'
export const ID_URL_EMPTY = 'ID_URL_EMPTY'
export const ID_KEEP_WALLET_PASS = 'ID_KEEP_WALLET_PASS'
export const ID_NEW_WALLET_PASS = 'ID_NEW_WALLET_PASS'

export const templateKeys = {
  [ID_NO_PASS_ERROR_MSG]: 'password cannot be empty',
  [ID_NO_APP_PASS_ERROR_MSG]: 'app password cannot be empty',
  [ID_PASSWORD_NOT_MATCH]: 'passwords do not match',
  [ID_SET_BUTTON_BUY]: 'Place order to buy  {{ asset }}',
  [ID_SET_BUTTON_SELL]: 'Place order to sell {{ asset }}',
  [ID_OFF]: 'off',
  [ID_READY]: 'ready',
  [ID_LOCKED]: 'locked',
  [ID_NOWALLET]: 'no wallet',
  [ID_WALLET_SYNC_PROGRESS]: 'wallet is {{ syncProgress }}% synced',
  [ID_HIDE_ADDIIONAL_SETTINGS]: 'hide additional settings',
  [ID_SHOW_ADDIIONAL_SETTINGS]: 'show additional settings',
  [ID_BUY]: 'Buy',
  [ID_SELL]: 'Sell',
  [ID_NOT_SUPPORTED]: '{{ asset }} is not supported',
  [ID_CONNECTION_FAILED]: 'Connection to dex server failed. You can close dexc and try again later or wait for it to reconnect.',
  [ID_ORDER_PREVIEW]: 'Total: {{ total }} {{ asset }}',
  [ID_CALCULATING]: 'calculating...',
  [ID_ESTIMATE_UNAVAILABLE]: 'estimate unavailable',
  [ID_NO_ZERO_RATE]: 'zero rate not allowed',
  [ID_NO_ZERO_QUANTITY]: 'zero quantity not allowed',
  [ID_TRADE]: 'trade',
  [ID_NO_ASSET_WALLET]: 'No {{ asset }} wallet',
  [ID_EXECUTED]: 'executed',
  [ID_BOOKED]: 'booked',
  [ID_CANCELING]: 'cancelling',
  [ID_DEX_ADDRESS_EMPTY]: 'DEX address cannot be empty',
  [ID_ACCT_UNDEFINED]: 'Account undefined.',
  [ID_URL_EMPTY]: 'URL cannot be empty',
  [ID_KEEP_WALLET_PASS]: 'keep current wallet password',
  [ID_NEW_WALLET_PASS]: 'set a new wallet password'

}

export default class Locales {
  static getCurrentLocale () {
    return window.navigator.userLanguage || window.navigator.language
  }

  // get from webserver/locales, but how?
  getLocales () {
  }

  setLocale (locale) {
    // check if setted locale is available on locales.
    // if (this.getLocales().indexOf(locale) !== -1) {
    //   ...

  }

  // formatDetails will format the message to its locale.
  // need to add one for plurals
  static formatDetails (stringKey, args = undefined) {
    return this.stringTemplateParser(templateKeys[stringKey], args)
  }

  // reference: https://stackoverflow.com/questions/43261798/javascript-how-to-use-template-literals-with-json
  static stringTemplateParser (expression, valueObj) {
    const templateMatcher = /{{\s?([^{}\s]*)\s?}}/g
    const text = expression.replace(templateMatcher, (substring, value, index) => {
      value = valueObj[value]
      return value
    })
    return text
  }
}
