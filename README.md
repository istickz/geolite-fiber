# GeoIP server for MaxMind databases

### Prerequisites
Before you begin, you should download a GeoLite2-City or GeoLite2-Country databases from the MaxMind website - https://dev.maxmind.com/geoip/geoip2/geolite2/. 
To do this, you may need to register for a free account.

Note if you want to receive only countries by IP, you can download GeoLite2-Country. It's smaller than GeoLite2-City but doesn't provide cities.

Place database file to geolite2_bases folder and run

```bash
export APP_TOKEN=123
go run geoip-maxmind
```

### Usage
Make a request to `http://127.0.0.1:3000/geocountry/46.88.89.152` or `http://127.0.0.1:3000/geocity/46.88.89.152` for example.
If the IP address is invalid, you will receive an HTTP 400 error.

If you want to modify fields that be returned, you need to use `"github.com/oschwald/maxminddb-golang"` and the structures that you want.

### Example response for geocountry request
```bash
curl --location --request GET 'http://127.0.0.1:3000/geocountry/46.88.89.152' \
--header 'Authorization: Token 123'
```

```json
{
  "Continent": {
    "Code": "EU",
    "GeoNameID": 6255148,
    "Names": {
      "de": "Europa",
      "en": "Europe",
      "es": "Europa",
      "fr": "Europe",
      "ja": "ヨーロッパ",
      "pt-BR": "Europa",
      "ru": "Европа",
      "zh-CN": "欧洲"
    }
  },
  "Country": {
    "GeoNameID": 2921044,
    "IsInEuropeanUnion": true,
    "IsoCode": "DE",
    "Names": {
      "de": "Deutschland",
      "en": "Germany",
      "es": "Alemania",
      "fr": "Allemagne",
      "ja": "ドイツ連邦共和国",
      "pt-BR": "Alemanha",
      "ru": "Германия",
      "zh-CN": "德国"
    }
  },
  "RegisteredCountry": {
    "GeoNameID": 2921044,
    "IsInEuropeanUnion": true,
    "IsoCode": "DE",
    "Names": {
      "de": "Deutschland",
      "en": "Germany",
      "es": "Alemania",
      "fr": "Allemagne",
      "ja": "ドイツ連邦共和国",
      "pt-BR": "Alemanha",
      "ru": "Германия",
      "zh-CN": "德国"
    }
  },
  "RepresentedCountry": {
    "GeoNameID": 0,
    "IsInEuropeanUnion": false,
    "IsoCode": "",
    "Names": null,
    "Type": ""
  },
  "Traits": {
    "IsAnonymousProxy": false,
    "IsSatelliteProvider": false
  }
}
```
### Example response for geocity request

```bash
curl --location --request GET 'http://127.0.0.1:3000/geocity/46.88.89.152' \
--header 'Authorization: Token 123'
```

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
  "Continent": {
    "Code": "EU",
    "GeoNameID": 6255148,
    "Names": {
      "de": "Europa",
      "en": "Europe",
      "es": "Europa",
      "fr": "Europe",
      "ja": "ヨーロッパ",
      "pt-BR": "Europa",
      "ru": "Европа",
      "zh-CN": "欧洲"
    }
  },
  "Country": {
    "GeoNameID": 2921044,
    "IsInEuropeanUnion": true,
    "IsoCode": "DE",
    "Names": {
      "de": "Deutschland",
      "en": "Germany",
      "es": "Alemania",
      "fr": "Allemagne",
      "ja": "ドイツ連邦共和国",
      "pt-BR": "Alemanha",
      "ru": "Германия",
      "zh-CN": "德国"
    }
  },
  "Location": {
    "AccuracyRadius": 1,
    "Latitude": 52.5153,
    "Longitude": 13.2992,
    "MetroCode": 0,
    "TimeZone": "Europe/Berlin"
  },
  "Postal": {
    "Code": "10585"
  },
  "RegisteredCountry": {
    "GeoNameID": 2921044,
    "IsInEuropeanUnion": true,
    "IsoCode": "DE",
    "Names": {
      "de": "Deutschland",
      "en": "Germany",
      "es": "Alemania",
      "fr": "Allemagne",
      "ja": "ドイツ連邦共和国",
      "pt-BR": "Alemanha",
      "ru": "Германия",
      "zh-CN": "德国"
    }
  },
  "RepresentedCountry": {
    "GeoNameID": 0,
    "IsInEuropeanUnion": false,
    "IsoCode": "",
    "Names": null,
    "Type": ""
  },
  "Subdivisions": [
    {
      "GeoNameID": 2950157,
      "IsoCode": "BE",
      "Names": {
        "de": "Berlin",
        "en": "Land Berlin",
        "es": "Berlín",
        "fr": "Berlin",
        "ja": "ベルリン",
        "pt-BR": "Berlim",
        "ru": "Берлин",
        "zh-CN": "柏林"
      }
    }
  ],
  "Traits": {
    "IsAnonymousProxy": false,
    "IsSatelliteProvider": false
  }
}
```
