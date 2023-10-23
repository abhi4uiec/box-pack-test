# Software Engineering Challenge

Imagine for a moment that one of our product lines ships in various pack sizes:

•	250 Items

•	500 Items

•	1000 Items

•	2000 Items

•	5000 Items

Our customers can order any number of these items through our website, but they will always only be given complete packs.
1.	Only whole packs can be sent. Packs cannot be broken open.
2.	Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3.	Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

Write an application that can calculate the number of packs we need to ship to the customer.

## Deployment
#### The program is deployed on AWS and uses API Gateway + Lambda

## How to execute a program

#### Below curl command can be used to check the output of the program:

curl --location 'https://uakx3p554d.execute-api.us-east-1.amazonaws.com/default/packs' \
--header 'Content-Type: application/json' \
--data '{
"itemsOrdered": 1,
"boxSizes":
[
250,
500,
1000,
2000,
5000
]
}'

#### Aternatively postman can be used:

    Method : POST

    URL : https://uakx3p554d.execute-api.us-east-1.amazonaws.com/default/packs
    
    Body :
    {
        "itemsOrdered": 1,
        "boxSizes": 
        [
            250,
            500,
            1000,
            2000,
            5000
        ]
    }