## Starport beer

หัดสร้าง Blockchain Application ตอนที่ 2 มาสร้าง ระบบจัดการพระเครื่องบน blockchain กันเถอะ

### Create beer
มาเริ่มสร้าง starport app กันก่อนเลยแบบรอบที่แล้ว

```bash
$starport app github.com/aofiee/beer

⭐️ Successfully created a Cosmos app 'beer'.
👉 Get started with the following commands:

 % cd beer
 % starport serve

NOTE: add --verbose flag for verbose (detailed) output.
```

ใครยังไม่ได้ลง [Starport อ่านวิธีติดตั้งได้ที่ลิ้งนี้เลยนะครับ](/1-work-from-home-application-blockchain-cosmos-network/)

เราจะมาสร้าง Type ของ พระเครื่อง กันก่อน โดยผมจะ Design ว่า Application beer ของเราจะมี fields ต่างๆดังนี้

1. Name ชื่อเจ้าของพระเครื่อง eg. นาย ออฟฟี่ จี้โคตรๆ
2. Identity No หมายเลขบัตรประจำตัวประชาชน eg. 3100532567813
3. Address ที่อยู่ eg. 666/66 หมู่ 3 ต.สวนหลวง จ. กทม 10250
4. Telephon No หมายเลขโทรศัพท์ eg.0925905444
5. beer Type ประเภทของพระเครื่อง eg.เหรียญ,กรุ,พระผง
6. beer Name ชื่อของหลวงพ่อ eg. ไม่ทราบวัด,วัดซั่บแหมน
7. From ชื่อวัด/กรุ eg. ไม่ทราบรุ่น, วัดซั่บแหมน, เทิดเทิง
8. Model รุ่นพิมพ์ eg.ไม่ทราบรุ่น, พรายใหม่ล่าสุด
9. Surface เนื้อ eg.ทองแดง, ไม่ทราบ, ว่าน
10. Year ปีพระ eg. 2559,2524
11. Province จังหวัดของพระ eg. ปทุมธานี, กรุงเทพ
12. Description รายละเอียดเพิ่มเติม eg. สภาพสวย,ขอข้อมูลพระเพิ่มเติม
13. Remark หมายเหตุ eg.ไม่ทราบชื่อพระ, มีตำหนิ

จะสร้างโดยใช้ command starport type

```bash
$starport type beer Name Identity Address Telephon beerType beerName From Model Surface Year Province Description Remark

🎉 Created a type `beer`.

```

หลังจากนั้นสั่ง 

```bash
$starport serve

📦 Installing dependencies...
🛠️  Building the app...
🙂 Created an account. Password (mnemonic): voice home upgrade void initial glory language noble exit illness duty slice galaxy crane exhibit waste olive enforce bubble phone reject soon siren august
🙂 Created an account. Password (mnemonic): ridge swallow trade antenna canvas humble apple hotel jacket pond asthma cook sunny athlete chair rapid force beauty kangaroo degree void hundred zebra riot
🌍 Running a Cosmos 'beer' app with Tendermint at http://localhost:26657.
🌍 Running a server at http://localhost:1317 (LCD)

🚀 Get started: http://localhost:12345/
```

เมื่อทดสอบ run จะได้ตามภาพด้านล่าง

![Blog](./01.png "Blog")

เราจะทำการแยกหน้า Post Content Blog ใหม่ กับหน้า List Blog ออกจากกัน ด้วยการไปสร้างไฟล์ **vue/src/views/Post.vue**

```html
<template>
  <div>
    <app-layout>
      <app-text type="h1">Post</app-text>
      <wallet />
      <post-form-group />
    </app-layout>
  </div>
</template>
```
และเราะจำไปทำการสร้าง component เพื่อสร้าง group form ขึ้นมาโดย list type มาใส่ **vue/src/components/PostFormGroup.vue**

```html
<template>
  <div>
    <router-link to="/">List All beers</router-link>
    <div class="type" v-for="type in typeList" :key="type.type">
      <post-form :value="type" />
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    typeList() {
      return this.$store.state.app.types || [];
    },
  },
};
</script>
```


หลังจากนั้นเราจะไปแก้ไขไฟล์ **vue/src/components/TypeItem.vue** เพื่อแยก Post ออกมาอีกไฟล์นึงก่อน โดยจะเก็บไว้แค่ List Post Items พอ

```html
<template>
  <div>
    <app-text type="h2">List of {{ value.type }} items</app-text>
    <div class="item" v-for="instance in instanceList" :key="instance.id">
      <div class="item__field" v-for="(value, key) in instance" :key="key">
        <div class="item__field__key">{{ key }}:</div>
        <div class="item__field__value">{{ value }}</div>
      </div>
    </div>
    <div
      class="card__empty"
      v-if="instanceList.length < 1"
    >There are no {{ value.type }} items yet. Create one using the form below.</div>
  </div>
</template>

<style scoped>
.item {
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow: hidden;
}
.item__field {
  display: grid;
  line-height: 1.5;
  grid-template-columns: 15% 1fr;
  grid-template-rows: 1fr;
  word-break: break-all;
}
.item__field__key {
  color: rgba(0, 0, 0, 0.25);
  word-break: keep-all;
  overflow: hidden;
}
.card__empty {
  margin-bottom: 1rem;
  border: 1px dashed rgba(0, 0, 0, 0.1);
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  border-radius: 8px;
  color: rgba(0, 0, 0, 0.25);
  text-align: center;
  min-height: 8rem;
}
@keyframes rotate {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(-360deg);
  }
}
@media screen and (max-width: 980px) {
  .narrow {
    padding: 0;
  }
}
</style>

<script>
export default {
  props: ["value"],
  data: function () {
    return {
      fields: {},
      flight: false,
    };
  },
  computed: {
    hasAddress() {
      return !!this.$store.state.account.address;
    },
    instanceList() {
      return this.$store.state.data[this.value.type] || [];
    },
    valid() {
      return Object.values(this.fields).every((el) => {
        return el.trim().length > 0;
      });
    },
  },
};
</script>
```

หลังจากนั้นเราจะไปสร้าง Form **vue/src/components/PostForm.vue** สำหรับ Post ไว้ตามนี้ 

```html
<template>
  <div>
    <app-text type="h2">New {{ value.type }}</app-text>
    <div v-for="field in value.fields" :key="field">
      <div v-if="field === 'Detail'">
        <textarea
          v-model="fields[field]"
          :placeholder="title(field)"
          :disabled="flight"
          rows="10"
          cols="60"
        ></textarea>
      </div>
      <div v-else>
        <app-input
          v-model="fields[field]"
          type="text"
          :placeholder="title(field)"
          :disabled="flight"
        />
      </div>
    </div>
    <button
      :class="['button', `button__valid__${!!valid && !flight && hasAddress}`]"
      @click="submit"
    >
      Create {{ value.type }}
      <div class="button__label" v-if="flight">
        <div class="button__label__icon">
          <icon-refresh />
        </div>Sending transaction...
      </div>
    </button>
  </div>
</template>

<style scoped>
button {
  background: none;
  border: none;
  color: rgba(0, 125, 255);
  padding: 0;
  font-size: inherit;
  font-weight: 800;
  font-family: inherit;
  text-transform: uppercase;
  margin-top: 0.5rem;
  cursor: pointer;
  transition: opacity 0.1s;
  letter-spacing: 0.03em;
  transition: color 0.25s;
  display: inline-flex;
  align-items: center;
}
button:focus {
  opacity: 0.85;
  outline: none;
}
.button.button__valid__true:active {
  opacity: 0.65;
}
.button__label {
  display: inline-flex;
  align-items: center;
}
.button__label__icon {
  height: 1em;
  width: 1em;
  margin: 0 0.5em 0 0.5em;
  fill: rgba(0, 0, 0, 0.25);
  animation: rotate linear 4s infinite;
}
.button.button__valid__false {
  color: rgba(0, 0, 0, 0.25);
  cursor: not-allowed;
}
@keyframes rotate {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(-360deg);
  }
}
@media screen and (max-width: 980px) {
  .narrow {
    padding: 0;
  }
}
</style>

<script>
export default {
  props: ["value"],
  data: function () {
    return {
      fields: {},
      flight: false,
    };
  },
  created() {
    (this.value.fields || []).forEach((field) => {
      this.$set(this.fields, field, "");
    });
  },
  computed: {
    hasAddress() {
      return !!this.$store.state.account.address;
    },
    instanceList() {
      return this.$store.state.data[this.value.type] || [];
    },
    valid() {
      return Object.values(this.fields).every((el) => {
        return el.trim().length > 0;
      });
    },
  },
  methods: {
    title(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    async submit() {
      if (this.valid && !this.flight && this.hasAddress) {
        this.flight = true;
        const payload = { type: this.value.type, body: this.fields };
        await this.$store.dispatch("entitySubmit", payload);
        await this.$store.dispatch("entityFetch", payload);
        this.flight = false;
        Object.keys(this.fields).forEach((f) => {
          this.fields[f] = "";
        });
      }
    },
  },
};
</script>
```

ทำการสร้าง route ใน VueJs ให้ route ไปยัง post component
**vue/src/router/index.js**

```javascript
const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/post",
    component: Post,
  },
];
```
จะได้หน้าตา flow การทำงานคร่าวๆตามนี้

![Blog](./02.png "Blog")

![Blog](./03.png "Blog")

ต่อไปเราจะมาดู Data Type ของ **x/beer/types/Typebeer.go** ที่ถู starport generate ขึ้นมาให้เรา

```golang
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type beer struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Name string `json:"Name" yaml:"Name"`
  Identity string `json:"Identity" yaml:"Identity"`
  Address string `json:"Address" yaml:"Address"`
  Telephon string `json:"Telephon" yaml:"Telephon"`
  beerType string `json:"beerType" yaml:"beerType"`
  beerName string `json:"beerName" yaml:"beerName"`
  From string `json:"From" yaml:"From"`
  Model string `json:"Model" yaml:"Model"`
  Surface string `json:"Surface" yaml:"Surface"`
  Year string `json:"Year" yaml:"Year"`
  Province string `json:"Province" yaml:"Province"`
  Description string `json:"Description" yaml:"Description"`
  Remark string `json:"Remark" yaml:"Remark"`
}
```

beer structure จะมี 15 properties ตามที่เรา design ไว้เมื่อข้างต้น

ตอนนี้เรายังไม่ได้ไป modify อะไรมันปล่อยไว้ก่อน เราจะกลับมาที่ Terminal ของเราก่อน โดยใน base directory เราจะพิมพ์ commandline **appcli** ดูเพื่อจะลองสร้าง transaction กันครับ

```bash
$appcli -help   
Command line interface for interacting with appd

Usage:
  appcli [command]

Available Commands:
  status      Query remote node for status
  config      Create or query an application CLI configuration file
  query       Querying subcommands
  tx          Transactions subcommands
              
  rest-server Start LCD (light-client daemon), a local REST server
              
  keys        Add or view local private keys
              
  version     Print the app version
  help        Help about any command

Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
  -h, --help              help for appcli
      --home string       directory for config and data (default "/Users/skulltree/.appcli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

Use "appcli [command] --help" for more information about a command.
```

หากเราลองสั่งคำสั่งที่เราจะสร้างข้อมูลพระเครื่องใหม่ของเราก็สั่ง **-help** ตาม commandline ข้างล่างเลย ก็จะบอกว่าเราต้องระบุ options อะไรบ้าง จะเห็นว่าเหมือนที่เรา define ไว้ข้างต้นเลยครับ

```bash
$appcli tx beer create-beer -help
Creates a new beer

Usage:
  appcli tx beer create-beer [Name] [Identity] [Address] [Telephon] [beerType] [beerName] [From] [Model] [Surface] [Year] [Province] [Description] [Remark] [flags]
```

อ่ะ เรามาลองดูกัน

```bash
$appcli tx beer create-beer "Khomkrid Lerdprasert" "3100503235431" "35/5 ซอยพระรามเก้า 57/1(วิเศษสุข 2) แขวงสวนหลวง เขตสวนหลวง กรุงเทพฯ 10250 เบอร์โทร 021539489" "0925905444" "พระเนื้อว่าน" "สมเด็จวัดระฆัง" "บางขุนพรหม เกศไชโย" "พิมพ์ใหญ่พิมพ์พระรอด" "เนื้อเขียวก้านมะลิ" "2497" "อยุธยา" "สมเด็จวัดระฆังพิมพ์ใหญ่เนื้อเขียวก้านมะลิ ลักษณะเนื้อพระองค์นี้เรียกกันว่าเนื้อปูนนุ่ม" "-" --from=aofiee

$appcli tx beer create-beer "Arnon Kidlerdpol" \
"3100503235431" \
"44/261"  \
"0925905444" \
"พระเนื้อว่าน" \
"สมเด็จวัดระฆัง" \
"บางขุนพรหม เกศไชโย" \
"พิมพ์ใหญ่พิมพ์พระรอด" \
"เนื้อเขียวก้านมะลิ" \
"2497" \
"อยุธยา" \
"สมเด็จวัดระฆังพิมพ์ใหญ่เนื้อเขียวก้านมะลิ ลักษณะเนื้อพระองค์นี้เรียกกันว่าเนื้อปูนนุ่ม" \
"-" \
--from=snappytux
```

เมื่อเรา enter ก็จะพบข้อมูลยืนยัน block ข้อมูลพระเครื่องที่เราเพิ่มลงไป

```bash
{
  "chain_id": "beer",
  "account_number": "2",
  "sequence": "3",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "beer/Createbeer",
      "value": {
        "ID": "637fc717-b850-43c4-867f-cc43720f9a06",
        "creator": "cosmos1r9fuvas8fvdvu3xa4gmyqz83c8hn8287wza6dq",
        "Name": "Khomkrid Lerdprasert",
        "Identity": "3100503235431",
        "Address": "35/5 ซอยพระรามเก้า 57/1(วิเศษสุข 2) แขวงสวนหลวง เขตสวนหลวง กรุงเทพฯ 10250 เบอร์โทร 021539489",
        "Telephon": "0925905444",
        "beerType": "พระเนื้อว่าน",
        "beerName": "สมเด็จวัดระฆัง",
        "From": "บางขุนพรหม เกศไชโย",
        "Model": "พิมพ์ใหญ่พิมพ์พระรอด",
        "Surface": "เนื้อเขียวก้านมะลิ",
        "Year": "2497",
        "Province": "อยุธยา",
        "Description": "สมเด็จวัดระฆังพิมพ์ใหญ่เนื้อเขียวก้านมะลิในตำนาน เนื้อที่ออกเขียวอ่อนๆมาจากส่วนผสมมวลสารที่มีส่วนประกอบของเหลวสีเขียวจากดอกไม้บูชาไม่ว่าจะเป็นส่วนใบหรือก้านที่นำมาตำบดให้ละเอียดเพื่อผสมกับเนื้อปูนที่ขาวใสจากเปลือกหอย  ลักษณะเนื้อพระองค์นี้เรียกกันว่าเนื้อปูนนุ่ม",
        "Remark": "-"
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]:
```

เมื่อเราทำการยืนยัน sign transaction แล้วเราจะได้ TxHash กลับมา

```bash
{
  "height": "0",
  "txhash": "A350DBB31FE3D3207723180A85DC299A237D7F103D74E9DB3CE7C03532118764",
  "raw_log": "[]"
}
```
เราสามารถตรวจสอบ Transaction ของเราว่าได้รับการประมลผลแล้วหรือยัง ด้วยการเข้าไปที่ url **http://localhost:26657/tx?hash=0x** แล้วตามด้วย **txhash** ที่เราได้มา แบบ link ด้านล่าง

**http://localhost:26657/tx?hash=0xA350DBB31FE3D3207723180A85DC299A237D7F103D74E9DB3CE7C03532118764**

```json
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "hash": "A350DBB31FE3D3207723180A85DC299A237D7F103D74E9DB3CE7C03532118764",
    "height": "2577",
    "index": 0,
    "tx_result": {
      "code": 0,
      "data": null,
      "log": "[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"Createbeer\"}]}]}]",
      "info": "",
      "gasWanted": "200000",
      "gasUsed": "90871",
      "events": [
        {
          "type": "message",
          "attributes": [
            {
              "key": "YWN0aW9u",
              "value": "Q3JlYXRlQW11bGV0"
            }
          ]
        }
      ],
      "codespace": ""
    },
    "tx": "5wsoKBapCu4Kv+kDjgokNjM3ZmM3MTctYjg1MC00M2M0LTg2N2YtY2M0MzcyMGY5YTA2EhQZU8Z2B0sazkTdqjZACPHB7zOo/hoUS2hvbWtyaWQgTGVyZHByYXNlcnQiDTMxMDA1MDMyMzU0MzEq0AEzNS81IOC4i+C4reC4ouC4nuC4o+C4sOC4o+C4suC4oeC5gOC4geC5ieC4siA1Ny8xKOC4p+C4tOC5gOC4qOC4qeC4quC4uOC4giAyKSDguYHguILguKfguIfguKrguKfguJnguKvguKXguKfguIcg4LmA4LiC4LiV4Liq4Lin4LiZ4Lir4Lil4Lin4LiHIOC4geC4o+C4uOC4h+C5gOC4l+C4nuC4ryAxMDI1MCDguYDguJrguK3guKPguYzguYLguJfguKMgMDIxNTM5NDg5MgowOTI1OTA1NDQ0OiTguJ7guKPguLDguYDguJnguLfguYnguK3guKfguYjguLLguJlCKuC4quC4oeC5gOC4lOC5h+C4iOC4p+C4seC4lOC4o+C4sOC4huC4seC4h0o04Lia4Liy4LiH4LiC4Li44LiZ4Lie4Lij4Lir4LihIOC5gOC4geC4qOC5hOC4iuC5guC4olI84Lie4Li04Lih4Lie4LmM4LmD4Lir4LiN4LmI4Lie4Li04Lih4Lie4LmM4Lie4Lij4Liw4Lij4Lit4LiUWjbguYDguJnguLfguYnguK3guYDguILguLXguKLguKfguIHguYnguLLguJnguKHguLDguKXguLRiBDI0OTdqEuC4reC4ouC4uOC4mOC4ouC4snKMBuC4quC4oeC5gOC4lOC5h+C4iOC4p+C4seC4lOC4o+C4sOC4huC4seC4h+C4nuC4tOC4oeC4nuC5jOC5g+C4q+C4jeC5iOC5gOC4meC4t+C5ieC4reC5gOC4guC4teC4ouC4p+C4geC5ieC4suC4meC4oeC4sOC4peC4tOC5g+C4meC4leC4s+C4meC4suC4mSDguYDguJnguLfguYnguK3guJfguLXguYjguK3guK3guIHguYDguILguLXguKLguKfguK3guYjguK3guJnguYbguKHguLLguIjguLLguIHguKrguYjguKfguJnguJzguKrguKHguKHguKfguKXguKrguLLguKPguJfguLXguYjguKHguLXguKrguYjguKfguJnguJvguKPguLDguIHguK3guJrguILguK3guIfguYDguKvguKXguKfguKrguLXguYDguILguLXguKLguKfguIjguLLguIHguJTguK3guIHguYTguKHguYnguJrguLnguIrguLLguYTguKHguYjguKfguYjguLLguIjguLDguYDguJvguYfguJnguKrguYjguKfguJnguYPguJrguKvguKPguLfguK3guIHguYnguLLguJnguJfguLXguYjguJnguLPguKHguLLguJXguLPguJrguJTguYPguKvguYnguKXguLDguYDguK3guLXguKLguJTguYDguJ7guLfguYjguK3guJzguKrguKHguIHguLHguJrguYDguJnguLfguYnguK3guJvguLnguJnguJfguLXguYjguILguLLguKfguYPguKrguIjguLLguIHguYDguJvguKXguLfguK3guIHguKvguK3guKIgIOC4peC4seC4geC4qeC4k+C4sOC5gOC4meC4t+C5ieC4reC4nuC4o+C4sOC4reC4h+C4hOC5jOC4meC4teC5ieC5gOC4o+C4teC4ouC4geC4geC4seC4meC4p+C5iOC4suC5gOC4meC4t+C5ieC4reC4m+C4ueC4meC4meC4uOC5iOC4oXoBLRIEEMCaDBpqCibrWumHIQPn/NWGQxtOqBVpmeObvpg0k2aRcWKo7FtRrV0Fu7jpcxJAvCg/G1LQUUoUjUCETv7bbOWP5v7JqVKzh0parv1Z/5A2NYnve3QNK0fiLwwSqvGZBZWELd0mzASdpQm0pUSMLA=="
  }
}
```

**http://localhost:26657/abci_query?path=%22custom/beer/list-beer%22**

ตอนที่เราสร้าง transaction จะมี params **--from=aofiee** เราสามารถตั้งต้นจำนวน User เริ่มต้น และทำการ initial coins เริ่มต้นลงใน genesis block ได้ จากไฟล์ **beer/config.yml**

```yml
version: 1
accounts:
  - name: aofiee
    coins: ["1000token", "100000000stake"]
  - name: champ
    coins: ["500token"]
  - name: eiang
    coins: ["200token"]
  - name: chao
    coins: ["100token"]
validator:
  name: aofiee
  staked: "100000000stake"
```

#### `accounts`

เราสามารถสร้าง account ได้จาก template นี้

| Key   | Required | Type            | Description          |
| ----- | -------- | --------------- | -------------------- |
| name  | Y        | String          | ชื่อผู้ใช้งาน             |
| coins | Y        | List of Strings | จำนวน coin พร้อมสกุลเงิน |

#### `validator`

validator ของระบบ เรา โดยใช้ได้เพียง user เดียวจาก accounts ที่เราลิสไว้ใน genesis block และจำนวน token ที่เราเก็บไว้ควรมีจำนวนมากกว่า 100 ล้าน token

| Key    | Required | Type   | Description                                                                         |
| ------ | -------- | ------ | ----------------------------------------------------------------------------------- |
| name   | Y        | String | Name of one the accounts                                                            |
| staked | Y        | String | Amount of coins staked by your validator, should be >= 10^6 (e.g. "100000000stake") |

หลังจากนั้นผมจะทำการ commit ขึ้น repo ไว้ก่อน แล้วมาทำ Makefile และ Makefile.ledger เพื่อเตรียม ทำการ deploy ขึ้นทดสอบ โดย Makefile จะมี code สำหรับ build go.mod ตามนี้

### Deploy beer Blockchain

```bash
PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NameService \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=nsd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=nscli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: install

install: go.sum
		@echo "--> Installing nsd & nscli"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/appd
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/appcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
```

และ Makefile.ledger 

```bash
LEDGER_ENABLED ?= true

build_tags =
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif
```

จากนั้นก็ไปทำ Dockerfile สำหรับ build container ไว้ทดสอบเล่นในเครื่อง

```Dockerfile
FROM ubuntu:focal

WORKDIR /blockchain
EXPOSE 12345 1317 26657 8080 26656
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH ${GOPATH}/bin:${GOROOT}/bin:${PATH}

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y make gcc vim curl wget git net-tools iputils-ping nmap netcat && \
    wget https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz && \
    tar -xvf go1.15.1.linux-amd64.tar.gz && \
    mv go /usr/local && \
    touch ~/.bashrc && \
    echo ${GOROOT} && \
    echo ${GOPATH} && \
    echo ${PATH} && \
    git clone https://github.com/aofiee/beer.git && \
    cd beer && \
    git checkout commit-transaction  && \
    make install && \
    rm -r /var/lib/apt/lists/*
```

หลังจากนั้นเราจะ build image จาก dockerfile ของเรากันก่อนด้วยคำสั่ง

```bash
docker build -t blockchain .
```

จากนั้นเราจะสร้าง container แรกเป็น master node ด้วยคำสั่ง

```bash
docker run -it -d --name node_1 blockchain

root@0fd579952e3e:/blockchain#
```

เมื่อเราเข้ามาอยู่ใน container แล้วเราจะทำการตั้งค่าต่างๆของ node แรกของเรากันก่อน

```bash
appd init beer_master_node --chain-id beer
appcli config chain-id beer
appcli config output json
appcli config indent true
appcli config trust-node true
appcli config keyring-backend test
appcli keys add aofiee
appcli keys add snappytux
```

จากนั้นเราจะได้ address และ mnemonic กลับมาให้เราเก็บไว้ให้ดี
```json
{
  "name": "aofiee",
  "type": "local",
  "address": "cosmos1exq0w75zlmmklz83krw6ppcr448a22mna9rsvd",
  "pubkey": "cosmospub1addwnpepqd4er6dqx6prh3tsgeueshrear658usd7xjcqc8nlj9glhngyrgmup29nkz",
  "mnemonic": "digital fox mask barrel gold maximum tuna job special match electric quote cook charge attract submit seek rally electric animal gallery violin device unhappy"
}
{
  "name": "snappytux",
  "type": "local",
  "address": "cosmos1lr6rzsglmvma8x3w3f2zt5f6kj5kfxrqy4jn6s",
  "pubkey": "cosmospub1addwnpepqg5eqysweewxh2drsg799m489fcj8yvcu29zp4vq4jh7ld90qngjxhnrqyn",
  "mnemonic": "still midnight razor title lion lyrics floor hub discover drama street easy either venue escape talk paddle melody solar easily absorb rotate unfold history"
}
```
เรามาสร้าง Stake ให้กับ Account เรากันต่อ สำหรับเป็น Validator

```bash
appd add-genesis-account $(appcli keys show aofiee -a) 100000000rune,100000000gold,100000000stake
appd add-genesis-account $(appcli keys show snappytux -a) 100000000rune,100000000gold,100000000stake
```

และทำการ validate genesis

```bash
appd gentx --name aofiee --keyring-backend test

appd collect-gentxs
appd validate-genesis
```

หลังจากนั้นเราทำการ start node แรกกันได้เลย ด้วยคำสั่ง 

```bash
appd start
```

ต่อไปเราจะทำการสร้าง node 2 และ node 3

```bash
docker run -it -d --name node_2 blockchain

root@0fd579952e3e:/blockchain#
```
และ 

```bash
docker run -it -d --name node_3 blockchain

root@0fd579952e3e:/blockchain#
```

โดยเราจะเปลี่ยนชื่อของแต่ละ node เป็น node_2 และ node_3 ตามลำดับ

```bash
appd init beer_node2 --chain-id beer
```

หลังจากนั้นเราจะทำการลบไฟล์ genesis.json ที่อยู่ใน directory **~/.appd/config/genesis.json**

```bash 
rm -rf ~/.appd/config/genesis.json
```

แล้วทำ ไฟล์ **~/.appd/config/genesis.json** จาก node_1 มาใส่แทนไว้ใน node_2 และ node_3

หลังจากนั้นเราจะมาทำการแก้ไข้ให้ node_2 และ node_3 รู้จัก node_1 กัน

ด้วยการไปแก้ไข ไฟล์ **vi ~/.appd/config/config.toml**

โดยเราจะหา id ของ node 1 จากคำสั่ง

```bash
appd tendermint show-node-id
```

```bash
persistent_peers = "id@first_node_ip:26656"
```

แล้วทำการ start node 2 และ 3 ด้วยคำสั่ง

```bash
appd start
```

เป็นอันว่าเราสามารถ run ทั้ง 3 node ขึ้นมาได้่แล้ว ซึ่งใน cosmos เราจะเรียกมันว่า zone ต่อไปเราจะทำให้ทั้ง 3 zone เชื่อมต่อกันบน cosmos hub

### Installing gaiacli

เราจะมาทำ Dockerfile เพื่อ build image gaia กันก่อน

```Dockerfile
FROM ubuntu:focal

WORKDIR /blockchain
EXPOSE 12345 1317 26657 8080 26656
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH ${GOPATH}/bin:${GOROOT}/bin:${PATH}

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y make gcc vim curl wget git net-tools iputils-ping nmap netcat && \
    wget https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz && \
    tar -xvf go1.15.1.linux-amd64.tar.gz && \
    mv go /usr/local && \
    touch ~/.bashrc && \
    echo ${GOROOT} && \
    echo ${GOPATH} && \
    echo ${PATH} && \
    git clone -b v2.0.13 https://github.com/cosmos/gaia && \
    cd gaia && \
    make install && \
    rm -r /var/lib/apt/lists/*
```
แล้วทำการ build เพื่อสร้าง image cosmos hub ขึ้นมา

```bash
docker build -t beer_hub .
```

จากนั้น run container มันขึ้นมา และเข้าไปข้างในซะ

```bash
docker run -it -d --name beer_hub beer_hub
docker exec -it beer_hub bash
```

จากนั้นจะทดลอง run คำสั่งที่เรา build มันขึ้นมาจาก Dockerfile ตะกี้

```bash
gaiad version --long

name: gaia
server_name: gaiad
client_name: gaiacli
version: 2.0.13
commit: bc93dcc040570c71fad288ba5eeebe4033a5c5d8
build_tags: netgo,ledger
go: go version go1.15.1 linux/amd64
```

และ 

```bash
gaiacli version --long

name: gaia
server_name: gaiad
client_name: gaiacli
version: 2.0.13
commit: bc93dcc040570c71fad288ba5eeebe4033a5c5d8
build_tags: netgo,ledger
go: go version go1.15.1 linux/amd64
```

ต่อไปเราจะมาทำ 

### Deploy Your Own Gaia Testnet

โดยเราจะทำการ setup แบบ Single-node, Local, Manual Testnet

ขั้นแรกเราจะทำการ Create Genesis File และ Start the Network สำหรับ Cosmos  HUB กันนะครับ

```bash
gaiad init --chain-id=amulte testing
```

สร้าง account validator โดยผมทดลองใช้ **password**

```bash
gaiacli keys add beer_validator
override the existing name beer_validator [y/N]: y
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:

- name: beer_validator
  type: local
  address: cosmos1nudnvlwjfn60lrl8yxghyflccry577d2jhajav
  pubkey: cosmospub1addwnpepqgdce5d7c3ety2q6hrrvw48qg3encjy984hsk43ccwul8tjdpfmncjpvu0h
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

remind develop asthma imitate link fiscal web nasty path spider story actress whale exhibit bag there puzzle famous dolphin attack lounge spatial man token
```
ระบบจะให้เราจด key mnemonic ไว้เผื่อเราลืมรหัสผ่านของเราทีหลัง

เสร็จแล้วเรามาเพิ่ม stake ให้กับ validator เรากันต่อ

```bash
gaiad add-genesis-account $(gaiacli keys show beer_validator -a) 1000000000stake,1000000000token
```

และทำการสร้าง genesis transaction ด้วยคำสั่ง

```bash
gaiad gentx --name beer_validator
Password to sign with 'beer_validator':
Genesis transaction written to "/root/.gaiad/config/gentx/gentx-6053d713e08d72391a6626ff243524014ba53fed.json"
```

```bash
$(appd tendermint show-validator)

appcli tx staking create-validator \
  --amount=1000000uatom \
  --pubkey=cosmosvalconspub1zcjduepq3mhsjvh2nqg7pzsyp59at3kp6zlaayscsvh7q9np34cnhfw750lshk5yt8 \
  --moniker="node_2" \
  --chain-id="beer" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --gas-prices="0.025uatom" \
  --from=aofiee
```