# dates-webservice

The **DatesAPI** webservice (accessible at https://api.datesapi.net/) answers a few date related questions. I wish to work with those values often in automation services like Zapier or Integromat. Because I haven't found a simple enough solution I quickly built my own service. Feel free to use this service.

## endpoints

The API offers these endpoints via HTTP GET requests and answers with a simple JSON struct.

- /
- /status
- /this-month
- /next-month
- /last-month
- /this-year
- /next-year
- /last-year
- /today
- /tomorrow
- /yesterday

## Formatting

The **day** endpoints (today, tomorrow, yesterday) accept an optional query parameter `format`.

The value of the `format` parameter must be compatible to the Go time package (see https://golang.org/pkg/time/#Time.Format for more details).

## Language

Coming soon: an option to demand the response in a specific language. Makes probably only sense for the month names (I want them to be returned in German). Since there is no built-in Go date localization API I will probably only deliver a few hand-selected languages.

