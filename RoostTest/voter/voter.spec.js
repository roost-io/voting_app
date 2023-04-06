const webdriver = require('selenium-webdriver');
const chrome = require('selenium-webdriver/chrome');
const ROOST_SVC_URL = 'https://www.example.com';

const options = new chrome.Options();
options.addArguments('start-maximized');

const driver = new webdriver.Builder()
  .forBrowser('chrome')
  .setChromeOptions(options)
  .build();

driver.get(ROOST_SVC_URL);

driver.getTitle().then((title) => {
  console.log(`The page title is: ${title}`);
});

driver.quit();