# dates-webservice

The **DatesAPI** webservice (accessible at https://api.datesapi.net/) answers a few date related questions. I wish to work with those values often in automation services like [Zapier](https://zapier.com/) or [Integromat](https://www.integromat.com/). Because I haven't found a simple enough solution I quickly built my own service. Feel free to use this service.

## endpoints

The API offers these endpoints via HTTP GET requests and answers with a simple JSON struct.

- [/](https://api.datesapi.net/)
- [/status](https://api.datesapi.net/status)
- [/this-month](https://api.datesapi.net/this-month)
- [/next-month](https://api.datesapi.net/next-month)
- [/last-month](https://api.datesapi.net/last-month)
- [/this-year](https://api.datesapi.net/this-year)
- [/next-year](https://api.datesapi.net/next-year)
- [/last-year](https://api.datesapi.net/last-year)
- [/today](https://api.datesapi.net/today)
- [/tomorrow](https://api.datesapi.net/tomorrow)
- [/yesterday](https://api.datesapi.net/yesterday)

## Formatting

The **day** endpoints (today, tomorrow, yesterday) accept an optional query parameter `format`.

The value of the `format` parameter must be compatible to the Go time package (see https://golang.org/pkg/time/#Time.Format for more details).

## Language

Coming soon: an option to demand the response in a specific language. Makes probably only sense for the month names (I want them to be returned in German). Since there is no built-in Go date localization API I will probably only deliver a few hand-selected languages.

