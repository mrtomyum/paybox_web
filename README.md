# Paybox_Terminal 1.0 
"The Ticket vending machine terminal client."
Paybox Ticket Terminal software Service and Application written in Qt5 and interface with QML.

## API Document
https://docs.google.com/document/d/1OfTsGxqN9QS4xgPgSkZJoN2dK2tgUIa_Cke57Ye86uo/edit?usp=sharing

## Design
                 ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐
                 │               Paybox Cloud                │
                 └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘
                             ▲                      ▲
                             │                      │
                             ▼                      ▼
                 ┌───────────────────────┐ ┌─────────────────┐
                 │   Paybox Terminal     │ │  Paybox Mobile  │
                 └───────────────────────┘ └─────────────────┘
                 ┌──────────┐ ┌──────────┐
                 │  Web UI  │ │  Divices │    
                 └──────────┘ └──────────┘ 
                 ┌──────────┐ ┌──────────┐ 
                 │   User   │ │ Hardware │     
                 └──────────┘ └──────────┘ 


## MVP Feature
+ Customer Interface for Ticket/Item selection and Payment
  + Browse Item from Categories, pictures.
  + Or Short-cut select item by press on-screen numeric button. (Client must print a list of product with NUMBER for customer)
  + [Ux] User select less than 3 click must found target item.
  + User choose size and quantity then select other item or proceed to check-out.
  + Display 3 language in same page [Thai, Eng, Chinese].
  + Voice greeting and sound effect when user selecting.
 
+ Local Storage 
  + DB SQlite3 
  + Local pictures and media file.
  
+ Initialization on boot or manual or push command for "System Health Check-up" such as...
  + 3G Network Status.
  + Peripheral Status as order and coin acceptor, hopper, printer
  + Door magnatic sensor
  + Touch screen LCD Mornitor
  + Update time from Paybox NTP Server.
+ Call API from cloud service.
  + Sell Transaction.
  + Money/Payment Transaction.
  + Alert
    + Door Open
    + Paper near end
    + Paper out
    + Coin Full
    + Bill Full

## 2nd Phase

+ Secret Interface for system setup such as.
  + OpenBox authentication with password.
  + Setup Alarm etc.


+ Subscribe MQTT for Push command such as 
  + Push for update Item Menu.
  + Size update.
  + Send heartbeat/online Status every minute etc.

+ Menu/Item Status online update.
  + Mobile App. add and update Menus/Items. Name, Size, Picture.
  + Client can control Item sale Status as "OPEN_SALE", "CLOSE_SALE" prevent payment and no available products.
  
## What's Next.
+ Add first time activation and "token" authentication practice.
+ OpenBox request by UserID and use mobile app receive "password" before open Box
+ MQTT Push new "Menu" with QML file from cloud.
+ MQTT Push new "Item" from cloud.
+ Staff can use machine as "Time Attendance".
