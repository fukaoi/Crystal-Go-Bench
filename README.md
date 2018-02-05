# Crystal-Go-Bench

* benchmark pc spec
  * pc: MacBook Pro (Retina 13-inch、Early 2015)
  * cpu: 2.7 GHz Intel Core i5
  * memory: 16 GB 1867 MHz DDR3

#### sha256 
* loop 1000000 print(Output) hash string

| language | version|average sec/10counts | raw data| 
|:-----------|:-----------|------------:|------------:|
| Crystal|0.24.1 (2017-12-26)|9.191sec|real 91.97, user 47.71, sys 23.79
| Go| go1.9.2 darwin/amd64 |9.181sec|real 91.81, user 33.16, sys 20.57|


* ※ Crystal release build binary