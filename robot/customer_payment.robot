*** Settings ***
Documentation    Suite description
Library    Selenium2Library
*** Variables ***
*** Test Cases ***
Customer Payment success
    [Tags]    DEBUG
    Provided precondition
    When select Item via Menu
    Host ask customer to make payment
    Money onhand is enough
    Then check expectations

*** Keywords ***
Provided precondition
    Setup system under test