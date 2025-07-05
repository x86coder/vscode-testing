from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from time import sleep

def test_title():
    print("Wal-mart automation testing")
    driver = webdriver.Chrome("")
    driver.get("https://www.walmart.com.mx")

    titulo = driver.title

    print("titulo:")
    print(titulo);

    sleep(6)
    driver.quit()

    assert titulo == "Compras en línea | Walmart online México"