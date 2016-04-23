
curl -v  -X GET "https://api.flightstats.com/flex/flightstatus/rest/v2/json/airport/status/ABQ/arr/2016/04/21/12?appId=appid&appKey=appkeu&utc=false&numHours=1&maxFlights=5"

{
 "request": {
  "airport": {
   "requestedCode": "ABQ",
   "fsCode": "ABQ"
  },
  "date": {
   "year": "2016",
   "month": "4",
   "day": "21",
   "interpreted": "2016-04-21"
  },
  "hourOfDay": {
   "requested": "12",
   "interpreted": 12
  },
  "numHours": {
   "requested": "1",
   "interpreted": 1
  },
  "utc": {
   "requested": "false",
   "interpreted": false
  },
  "codeType": {},
  "maxFlights": {
   "requested": "5",
   "interpreted": 5
  },
  "extendedOptions": {},
  "url": "https://api.flightstats.com/flex/flightstatus/rest/v2/json/airport/status/ABQ/arr/2016/04/21/12"
 },
 "appendix": {
  "airports": []
 },
 "error": {
  "httpStatusCode": 403,
  "errorCode": "AUTH_FAILURE",
  "errorId": "0da271c0-1c6f-4f83-9b41-61bd5c646f6f",
  "errorMessage": "Authorization failed. application_not_found:application with id=\"appid\" was not found"
 },
 "flightStatuses": []
}
