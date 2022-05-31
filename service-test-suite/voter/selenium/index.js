const { Builder } = require("selenium-webdriver");
const chrome = require("/Users/garvitjethwani/go/src/github.com/roost-io/voting_app/service-test-suite/voter/selenium/node_modules/chromedriver/lib/chromedriver");

const driver = new Builder().forBrowser("chrome").build() ;




(async function helloSelenium() {
    try {
        await driver.get('http://www.google.com');
    }
    finally {
        await driver.quit();
    }
})();
  
