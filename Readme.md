# Hacker Craft Blockchain for Beer Brewer

## Design Blockchain

### 1. Setup Project
```bash
scaffold app lvl-1 aofiee hackercraft
```
goto **/app/app.go**

```go
const appName = "hackercraft"
```

หลังจากนั้นเราจะเข้าไปยัง directory **blockchain/x**
และ run command ด้านล่าง

```bash
$scaffold module aofiee hackercraft beer
$scaffold module aofiee hackercraft brewer
$scaffold module aofiee hackercraft recipes
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

หลังจากนั้นเราจะทำการสร้าง Type Structure ขึ้นมา
**x/brewer/types/types.go**

```go
package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("ferment", 1)}

type Brewer struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	BrewerID string         `json:"BrewerID" yaml:"BrewerID"`
	TypeOfBrewer string	`json:"TypeOfBrewer" yaml:"TypeOfBrewer"`
	Address string	`json:"Address" yaml:"Address"`
	Telephone string	`json:"Telephone" yaml:"Telephone"`
	Email string	`json:"Email" yaml:"Email"`
	Story string	`json:"Story" yaml:"Story"`
	LogoURL string	`json:"LogoURL" yaml:"LogoURL"`
	CoverURL string	`json:"CoverURL" yaml:"CoverURL"`
	Founded string	`json:"Founded" yaml:"Founded"`
	Founder string	`json:"Founder" yaml:"Founder"`
	SiteURL string	`json:"SiteURL" yaml:"SiteURL"`
}

func NewBrewer() Brewer {
	return Brewer {
		BrewerID: uuid.New().String(),
		TypeOfBrewer: "Home Brew",
		Address: "44/261 Passorn Onut Prawet Prawet Bangkok 10250 Thailand.",
		Telephone: "+6692 590 5444",
		Email: "aofiee666@gmail.com",
		Story: "Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish.This is the beer that started it all - and it’s not done yet...",
		LogoURL: "https://www.brewdog.com/static/version1600328468/frontend/Born/arcticFox/en_US/images/logo.svg",
		CoverURL: "https://www.brewdog.com/static/version1600328468/frontend/Born/arcticFox/en_US/images/logo.svg",
		Founded: "2020-09-18",
		Founder: "Khomkrid Lerdprasert",
		SiteURL: "https://www.aofiee.dev"
	}
}

func (b Brewer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`BrewerID: %s`, b.BrewerID))
}
```

หลังจากนั้นเราจะไปทำการสร้าง Msg ให้กับ Brewer โดยหลักการตั้งชื่อ ของ SDK จะเป็น Msg{.Action} เช่นเราจะตั้งว่า SetBrewer ใน **x/brewer/types/MsgBrewer.go** เราจะตั้งเป็น MsgSetBrewer แทน

```go
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgSetBrewer struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	BrewerID string         `json:"BrewerID" yaml:"BrewerID"`
	TypeOfBrewer string	`json:"TypeOfBrewer" yaml:"TypeOfBrewer"`
	Address string	`json:"Address" yaml:"Address"`
	Telephone string	`json:"Telephone" yaml:"Telephone"`
	Email string	`json:"Email" yaml:"Email"`
	Story string	`json:"Story" yaml:"Story"`
	LogoURL string	`json:"LogoURL" yaml:"LogoURL"`
	CoverURL string	`json:"CoverURL" yaml:"CoverURL"`
	Founded string	`json:"Founded" yaml:"Founded"`
	Founder string	`json:"Founder" yaml:"Founder"`
	SiteURL string	`json:"SiteURL" yaml:"SiteURL"`
}
```

หลังจากนั้นเราจะทำการสร้าง MsgSetBrewer เพื่อทำการ Set ค่าให้กับ Brewer โดยมี Attributes ตามนี้

```go
func NewMsgSetBrewer (creator sdk.AccAddress, brewerID string, typeofbrewer string, address string, telephone string, email string, story string, logourl string, coverurl string, founded string, founder string, siteurl string) MsgSetBrewer {
	return MsgSetBrewer {
		Creator: creator,
		BrewerID: uuid.New().String(),
		TypeOfBrewer: typeofbrewer,
		Address: address,
		Telephone: telephone,
		Email: email,
		Story: story,
		LogoURL: logourl,
		CoverURL: coverurl,
		Founded: founded,
		Founder: founder,
		SiteURL: siteurl,
	}
}
```

หลังจากนั้นเราจะทำการส่ง route Msg ไปยัง Module ที่จัดการ Msg ของเราอีกที และทำการส่งชื่อ Tag เก็บไว้ใน Database 

```go
func (msg MsgSetBrewer) Route() string  {
	return RouterKey
}

func (msg MsgSetBrewer) Type() string  {
	return "SetBrewer"
}
```

จากนั้นเราจะใช้ ValidateBasic เพื่อทำการ check และ validate Msg attribute ที่ส่งเข้ามาว่าถูกต้องหรือไม่

```go
func (msg MsgSetBrewer) ValidateBasic() error  {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,msg.Creator.String())
	}
	if len(msg.BrewerID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "BrewerID cannot be empty")
	}
	if len(msg.TypeOfBrewer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "TypeOfBrewer cannot be empty")
	}
	if len(msg.Address) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Address cannot be empty")
	}
	if len(msg.Telephone) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Telephone cannot be empty")
	}
	if len(msg.Email) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Email cannot be empty")
	}
	if len(msg.Story) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Story cannot be empty")
	}
	if len(msg.LogoURL) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "LogoURL cannot be empty")
	}
	if len(msg.CoverURL) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "CoverURL cannot be empty")
	}
	if len(msg.Founded) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Founded cannot be empty")
	}
	if len(msg.Founder) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Founder cannot be empty")
	}
	if len(msg.SiteURL) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SiteURL cannot be empty")
	}
	return nil
}
```

และหลังจากนั้นเราจะบอกว่าใครจะต้องเป็นผู้ sign transaction บน Tx ให้กับเรา ในที่นี้ก็คือ Creator นั่นเอง

```go
func (msg MsgSetBrewer) GetSignBytes() []byte  {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetBrewer) GetSigner() []sdk.AccAddress  {
	return []sdk.AccAddress{msg.Creator}
}
```

ต่อไปเราจะสร้าง Handler เพื่อสั่งให้มีการทำงาน เมื่อมีการรับ message event เข้ามา
โดยเริ่มสร้างไฟล์ที่ **x/brewer/handlerMsgSetBrewer.go**

```go
package brewer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/hackercraft/x/brewer/types"
	"github.com/aofiee/hackercraft/x/brewer/keeper"
	"github.com/tendermint/tendermint/crypto"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"fmt"
)
```

ต่อไปเราจะใช้ NewHandler เป็น router ที่รับ Message เข้ามาแล้ว Forward ไปยัง Handler ที่เราต้องการ

```go
func NewHandler(keeper Keeper)  sdk.Handler {
	return func (ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error)  {
		switch msg := msg.(type) {
		case MsgSetBrewer:
			return handleMsgSetBrewer(ctx, keerper, msg)
		}
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}
```