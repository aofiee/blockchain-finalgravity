# Hacker Craft Blockchain for Beer Brewer

## Design Blockchain

### 1. Setup Project
```bash
scaffold app lvl-1 aofiee finalgravity
```
goto **/app/app.go**

```go
const appName = "finalgravity"
```

หลังจากนั้นเราจะเข้าไปยัง directory **finalgravity/x**
และ run command ด้านล่าง

```bash
$scaffold module aofiee finalgravity beer
$scaffold module aofiee finalgravity brewer
$scaffold module aofiee finalgravity recipes
```
โดยแรกเริ่ม ระบบ file จะมีโครงสร้างดังนี้
```bash
├── Makefile
├── Makefile.ledger
├── app
│   ├── app.go
│   └── export.go
├── cmd
│   ├── appcli
│   │   └── main.go
│   └── appd
│       ├── genaccounts.go
│       └── main.go
├── go.mod
├── go.sum
└── x
    ├── beer
    │   ├── README.md
    │   ├── abci.go
    │   ├── client
    │   ├── genesis.go
    │   ├── handler.go
    │   ├── keeper
    │   ├── module.go
    │   ├── spec
    │   └── types
    ├── brewer
    │   ├── README.md
    │   ├── abci.go
    │   ├── client
    │   ├── genesis.go
    │   ├── handler.go
    │   ├── keeper
    │   ├── module.go
    │   ├── spec
    │   └── types
    └── recipes
        ├── README.md
        ├── abci.go
        ├── client
        ├── genesis.go
        ├── handler.go
        ├── keeper
        ├── module.go
        ├── spec
        └── types
```

โดยจะสร้าง module สำหรับ blockchain เราขึ้นมา 3 module ได้แก่ beer, brewer, recipes
โดยแบ่งหน้าที่การทำงานออกเป็น 

1. brewer ทำหน้าที่เก็บ profile ของผู้ผลิตเบียร์ craft 
2. beer คือ ชนิดของ profile ของเบียร์นั้นๆ
3. recipes คือสูตรของแต่ละ batch เพื่อนำ transaction hash มาออกเป็น qr code เพื่อทำการ trackback

โดยเราจะวางแผนว่าในระบบ จะมี สกุลเงิน

1. stake สำหรับยึดถือครองทำ validator
2. ferment สำหรับ brewer เพื่อใช้ในการเปิด Branding
3. hoppies สำหรับ batch ที่ต้องการ generate qr code ออกมา

## Stucture ของ Brewer
1. BrewerID
2. Name ชื่อ Brand ของ Brewer
3. Brewer Type
   1. Major Brewery
   2. Micro Brewery
   3. Brew Pub
   4. Home Brew
4. Address
5. Telephone
6. Email
7. Story
8. Logo Image
9. Cover Image
10. Founded Since
11. Founder
12. Site url

## Structure ของ Beer
1. BeerID
2. BrewerID
3. Beer Name ชื่อของเบียร์
4. Type ชนิดของเบียร์
   1. Ale ผลิตโดยใช้อุณหภูมิอยู่ที่ 18- 24 องศาเซลเซียส และหมักแบบ Top-Fermenting Yeast (ยีสต์หมักลอยผิว) คือการที่ยีสต์ จะลอยอยู่ที่ผิวหน้าของเบียร์เมื่อเสร็จสิ้นกระบวนการหมัก ความโดดเด่นของเบียร์เอลนั้น อยู่ที่รสชาติค่อนไปทางหวาน, มีสีหลากหลาย ตั้งแต่ทองสว่าง จนถึงสีน้ำตาล โดยแตกต่างกันไปตามเมล็ดข้าวที่นำมาใช้ในกระบวนการผลิต และแบ่งแยกย่อยลงไปได้อีก
   2. Lager ผลิตโดยใช้อุณหภูมิอยู่ที่ 7-12 องศาเซลเซียส (ต่ำกว่าเบียร์เอล) หมักแบบ Bottom-Fermenting Yeast (ยีสต์หมักนอนก้น) หรือยีสต์ที่จมอยู่ที่ก้นภาชนะ เมื่อเสร็จสิ้นการหมัก นั่นเอง เบียร์ลาเกอร์มีคาแรคเตอร์ที่สดชื่น, สะอาด และนุ่มนวล โดยจะแยกย่อยลงไปอีก
   3. Lambic ผลิตจากยีสต์ชนิดพิเศษ มีต้นกำเนิดมาจากประเทศเบลเยี่ยม และในปัจจุบัน กลายเป็นที่นิยมไปทั่วโลก มักจะมีรสชาติเปรี้ยวสดชื่น วิธีการผลิต Lambic Beer (เบียร์ลัมบิค) ตามประเพณีดั้งเดิม คือ ผลิตในช่วงฤดูหนาว ตั้งแต่สิงหาคม จนถึงเมษายน โดยการบ่มในถังไม้อีกหลายปีก่อนจะสามารถนำมาดื่มได้
5. Style สไตล์ของเบียร์ที่เราทำ ยกตัวอย่างเช่น ซึ่งในโลกนี้มีหลายพันชนิด ยกตัวอย่างเช่น
   1. Light Lager
   2. Plisner
   3. Amber lager
   4. Bock and Dark Lager
6. Tag line รายละเอียดย่อยของเบียร์
7. Appearance ลักษณะของเบียร์
8. Taste รสชาติของเบียร์
9. Aroma กลิ่นของเบียร์
10. Strength ความแรงของ Alc เช่น 2.8-4.2% ABV
11. Story เรื่องราวของเบียร์ตัวนี้
12. Image Label

## Structure ของ Recipes
1. RecipesID
2. BeerID
3. BatchNo
4. Properties
   1. Original Gravity
   2. Expected Final Gravity
   3. Final Gravity
   4. Total Liquor
   5. Makes (Litres)
   6. Ready To Drink (Weeks)
   7. Estimated ABV
   8. Bitterness Rating
   9. Colour Rating
5.  Mashing
    1.  Liquor (Litres)
    2.  Mash Time (Hr)
    3.  Temperature (C)
    4.  Grain bill Properties [Array]
        1. Name
        2. Quantity (Kg,Lb,Oz)
6.  Boil
    1.  Liquor (Litres)
    2.  Boil Time (mins)
    3.  Hops Properties [Array]
        1.  Name
        2.  Quantity
        3.  IBU
        4.  When to add
    4.  Other Properties [Array]
        1. Name
        2. Quantity
        3. When to add
7. Ferment
   1. Yeast Properties [Array]
      1. Name
      2. Quantity
   2. Fermentation
      1. Temperature
      2. Conditioning
   3. Hops Properties [Array]
        1.  Name
        2.  Quantity
        3.  When to add
   4. Other Properties [Array]
        1. Name
        2. Quantity
        3. When to add
8. Ages
   1. Best Before
   2. Expired Date
   3. Bottled Date
   4. Best After
9. Factory
10. Image Label

```bash
appcli tx brewer create-brewer "Home Brew" \
"44/261 Passorn Onnut Prawet Prawet Bangkok 10250" \
"+66925905444" \
"aofiee666@gmail.com" \
"Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish. This is the beer that started it all - and it’s not done yet... PUNK - Quintessential Empire with an anarchic twist." \
"https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg" \
"https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg" \
"2018" \
"Khomkrid Lerdprasert" \
"https://www.aofiee.dev" --from=aofiee
```

```bash
appcli q brewer list-brewer       
[
  {
    "creator": "cosmos184cut6ds3tpatcggzy7mdtxzglwug72day6z9a",
    "BrewerID": "5d2a483f-44a1-4bfe-a919-f40ab88f2dc5",
    "TypeOfBrewer": "Home Brew",
    "Address": "44/261 Passorn Onnut Prawet Prawet Bangkok 10250",
    "Telephone": "+66925905444",
    "Email": "aofiee666@gmail.com",
    "Story": "Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish. This is the beer that started it all - and it’s not done yet... PUNK - Quintessential Empire with an anarchic twist.",
    "LogoURL": "https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg",
    "CoverURL": "https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg",
    "Founded": "2018",
    "Founder": "Khomkrid Lerdprasert",
    "SiteURL": "https://www.aofiee.dev"
  }
]
```

```bash
appcli q brewer get-brewer 6983a152-e282-46dc-aabb-fe234b5c252d
{
  "creator": "cosmos1wtdy42wxv89jh0vndwk4a6qjz946g7lt2yt9ge",
  "BrewerID": "6983a152-e282-46dc-aabb-fe234b5c252d",
  "TypeOfBrewer": "Home Brew",
  "Address": "44/261 Passorn Onnut Prawet Prawet Bangkok 10250",
  "Telephone": "+66925905444",
  "Email": "aofiee666@gmail.com",
  "Story": "Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish. This is the beer that started it all - and it’s not done yet... PUNK - Quintessential Empire with an anarchic twist.",
  "LogoURL": "https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg",
  "CoverURL": "https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg",
  "Founded": "2018",
  "Founder": "Khomkrid Lerdprasert",
  "SiteURL": "https://www.aofiee.dev"
}
```