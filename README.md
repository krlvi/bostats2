# Bostats2

Scrapes hemnet.se for apartment sale prices **per square meter**. Produces the following output:
1. A table price per square meter 50th, 75th and 90th percentiles for each month.
2. A frequencies histogram of prices per square meter for the whole time period.

Like [bostats](https://github.com/krlvi/bostats) but in Go instead of Clojure, faster, and it actually works.

# Usage
1) Have Go installed

2) On the [hemnet page for sold properties](https://www.hemnet.se/salda/bostader) search with parameters that interest you. 

3) Copy the URL.

4) Do a `go run main.go '<URL>'`

# Example 
This url https://www.hemnet.se/salda/bostader?location_ids%5B%5D=925968&sold_age=all gives sold properties in Kungsholmen.

Running `go run main.go 'https://www.hemnet.se/salda/bostader?location_ids%5B%5D=925968&sold_age=all'` will produce a result like the one below:

```
Price/Sqm per month percentiles:
|year|     month|vol|    50th|    75th|    90th|
|2018| September| 76| 84332.0| 92432.0|102704.5|
|2018|   October|202| 84790.0| 94674.5|101317.0|
|2018|  November|166| 83810.5| 91479.0|100000.0|
|2018|  December| 76| 82756.0| 89873.0|101711.0|
|2019|   January|152| 84432.5| 94286.0|106222.5|
|2019|  February|137| 81818.0| 90000.0| 97786.5|
|2019|     March|184| 85054.5| 94271.0|101293.0|
|2019|     April|148| 84782.0| 94024.0| 98856.5|
|2019|       May|211| 82243.0| 90696.5| 99791.5|
|2019|      June|149| 84675.0| 93909.0|103101.5|
|2019|      July| 79| 85345.0| 99305.5|108153.0|
|2019|    August|190| 87659.0| 98245.0|110426.0|
|2019| September|224| 88721.0| 96944.0|105441.0|
|2019|   October|240| 87426.5| 95455.0|102439.0|
|2019|  November|163| 86333.0| 98304.5|102728.5|
|2019|  December| 81| 89419.0| 98907.0|107977.5|
|2020|   January| 20| 91823.5| 99412.0|107500.0|

Price/Sqm frequency histogram for the period September 2018 to January 2020
52267-59796    0.52%  ▌                      13
59796-67326    4.84%  ████▏                  121
67326-74856    13.5%  ███████████▋           338
74856-82385    21.1%  ██████████████████     526
82385-89914    23.5%  ████████████████████▏  586
89914-97444    17.7%  ███████████████▏       442
97444-104974   10.4%  ████████▉              260
104974-112503  4.76%  ████▏                  119
112503-120032  1.68%  █▌                     42
120032-127562  1.16%  █                      29
127562-135092  0.28%  ▎                      7
135092-142621  0.28%  ▎                      7
142621-150150  0.24%  ▎                      6
150150-157680  0%     ▏
157680-165210  0.04%  ▏                      1
165210-172739  0%     ▏
172739-180268  0%     ▏
180268-187798  0%     ▏
187798-195328  0%     ▏
195328-202857  0.04%  ▏                      1
```

# Caveats

Hemnet limits to 50 pages per query, thus the broader the query, the less historical data you will get.


