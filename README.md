# Bostats2

Scrapes hemnet.se to for apartment sale prices **per square meter**. Produces a table with 50th, 75th and 90th percentiles.
Like [bostats](https://github.com/krlvi/bostats) but not in Clojure but Go, faster, and it actually works.

# Usage
1) Have Go installed

2) On the [hemnet page for sold properties](https://www.hemnet.se/salda/bostader) search for with parameters that interest you. 

3) Copy the URL.

4) Do a `go run main.go '<URL>'`

#Example 
This url https://www.hemnet.se/salda/bostader?location_ids%5B%5D=925968&sold_age=all gives sold properties in Kungsholmen.

Running `go run main.go 'https://www.hemnet.se/salda/bostader?location_ids%5B%5D=925968&sold_age=all'` will produce a result like the one below:

```
|year|     month|vol|    50th|    75th|    90th|
|2018| September| 95| 84794.0| 92466.0|103173.0|
|2018|   October|155| 85526.0| 94813.0|102697.5|
|2018|  November|163| 83824.0| 91698.5|100000.0|
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
|2019|   October|237| 87353.0| 95300.5|102353.0|
|2019|  November|163| 86250.0| 98304.5|102728.5|
|2019|  December| 81| 89419.0| 98907.0|107977.5|
|2020|   January|  4| 92059.5| 93421.0| 97960.5|

```

# Caveats

Hemnet limits to 50 pages per query, thus the broader the query, the less historical data you will get.


