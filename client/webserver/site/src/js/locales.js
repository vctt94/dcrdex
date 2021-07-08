export const LOGIN_ERROR_MESSAGE = 'LOGIN_ERROR_MESSAGE'
export const SET_BUTTON_BUY = 'SET_BUTTON_BUY'

export const templateKeys = {
  [LOGIN_ERROR_MESSAGE]: 'password cannot be empty',
  [SET_BUTTON_BUY]: `Place order to buy  {{ asset }}`

};

export default class Locales {
  static getCurrentLocale() {
    return window.navigator.userLanguage || window.navigator.language
  }

  // get from webserver/locales, but how?
  getLocales() {
  }

  setLocale(locale) {
    // check if setted locale is available on locales.
    // if (this.getLocales().indexOf(locale) !== -1) {
    //   ...

  }
 
  // formatDetails will format the message to its locale.
  // need to add one for plurals
  static formatDetails(stringKey, args = undefined) {
    return this.stringTemplateParser(templateKeys[stringKey], args)
  }

  // reference: https://stackoverflow.com/questions/43261798/javascript-how-to-use-template-literals-with-json
  static stringTemplateParser(expression, valueObj) {
    const templateMatcher = /{{\s?([^{}\s]*)\s?}}/g;
    let text = expression.replace(templateMatcher, (substring, value, index) => {
      value = valueObj[value];
      return value;
    });
    return text
  }
}
