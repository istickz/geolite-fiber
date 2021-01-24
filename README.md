# GeoIP server for MaxMind databases

### Prerequisites
Before you begin, you should download a GeoLite2-City and GeoLite2-Country databases from the MaxMind website - https://dev.maxmind.com/geoip/geoip2/geolite2/. 
To do this, you may need to register for a free account.

Place copy database files to geolite2_bases folder and run

```
go run geoip-maxmind
```

### Usage
Make a request to `http://127.0.0.1:3000/geocountry/46.88.89.152` or `http://127.0.0.1:3000/geocity/46.88.89.152` for example. 
You can omit an IP address to use your current IP address, or replace to use another. If the IP address is invalid, you will receive an HTTP 400 error.

The response fields can be modified from the `ipLookup` struct, found in the `handlers/handlers.go` file.

### Example response for geocountry request
```json
{
  "Country": {
    "IsoCode": "DE"
  }
}
```
### Example response for geocity request

```json
{
  "City": {
    "GeoNameID": 2950159,
    "Names": {
      "de": "Berlin",
      "en": "Berlin",
      "es": "Berlín",
      "fr": "Berlin",
      "ja": "ベルリン",
      "pt-BR": "Berlim",
      "ru": "Берлин",
      "zh-CN": "柏林"
    }
  },
  "Country": {
    "IsoCode": "DE"
  },
  "Location": {
    "AccuracyRadius": 1
  }
}
```

### Third-party library licenses
- [Examples for Fiber](https://github.com/gofiber/recipes/blob/master/LICENSE)
